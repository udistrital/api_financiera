package models

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/udistrital/api_financiera/utilidades"
)

type ConceptoValor struct {
	Concepto *Concepto
	Valor    int64
}

type RpCdpRubroConceptoValor struct {
	RegistroPresupuestalDisponibilidadApropiacion RegistroPresupuestalDisponibilidadApropiacion
	ConceptoValor                                 []ConceptoValor
}

type HomologacionConcepto struct {
	Id             int       `orm:"column(id);pk;auto"`
	Vigencia       float64   `orm:"column(vigencia)"`
	FechaCreacion  time.Time `orm:"column(fecha_creacion);type(date)"`
	ConceptoKronos *Concepto `orm:"column(concepto_kronos);rel(fk)"`
	ConceptoTitan  int       `orm:"column(concepto_titan)"`
}

func (t *HomologacionConcepto) TableName() string {
	return "homologacion_concepto"
}

func init() {
	orm.RegisterModel(new(HomologacionConcepto))
}

// AddHomologacionConcepto insert a new HomologacionConcepto into database and returns
// last inserted Id on success.
func AddHomologacionConcepto(m *HomologacionConcepto) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetHomologacionConceptoById retrieves HomologacionConcepto by Id. Returns error if
// Id doesn't exist
func GetHomologacionConceptoById(id int) (v *HomologacionConcepto, err error) {
	o := orm.NewOrm()
	v = &HomologacionConcepto{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllHomologacionConcepto retrieves all HomologacionConcepto matches certain condition. Returns empty list if
// no records exist
func GetAllHomologacionConcepto(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(HomologacionConcepto))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []HomologacionConcepto
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateHomologacionConcepto updates HomologacionConcepto by Id and returns error if
// the record to be updated doesn't exist
func UpdateHomologacionConceptoById(m *HomologacionConcepto) (err error) {
	o := orm.NewOrm()
	v := HomologacionConcepto{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteHomologacionConcepto deletes HomologacionConcepto by Id and returns error if
// the record to be deleted doesn't exist
func DeleteHomologacionConcepto(id int) (err error) {
	o := orm.NewOrm()
	v := HomologacionConcepto{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&HomologacionConcepto{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

// Homologacion conceptos de titan
func HomolgacionConceptosTitan(DataOpProveedor map[string]interface{}) (alerta Alert, err error, allRpCdpRubroConceptoValor []RpCdpRubroConceptoValor) {
	o := orm.NewOrm()
	var allConceptoValor []ConceptoValor
	var detalleLiquidacion []interface{}
	var registroPresupuestal RegistroPresupuestal

	err = utilidades.FillStruct(DataOpProveedor["DetalleLiquidacion"], &detalleLiquidacion)
	err = utilidades.FillStruct(DataOpProveedor["RegistroPresupuestal"], &registroPresupuestal)
	if err != nil {
		alerta.Type = "error"
		alerta.Code = "E_TRANS_01" //error en parametros de entrada
		alerta.Body = err.Error()
		return
	}

	for i, row := range detalleLiquidacion {
		m := row.(map[string]interface{})
		// data valorCalculado
		valorCalculadoFloat := m["ValorCalculado"].(float64)
		valorCalculado := int64(valorCalculadoFloat)
		// data concepto
		concepto, e := m["Concepto"].(map[string]interface{})
		if e != true {
			alerta.Type = "error"
			alerta.Code = "E_HO_CONC_TITAN"
			alerta.Body = err.Error()
			return
		}
		idConceptoTitanFloat := concepto["Id"].(float64)
		idConceptoTitan := int(idConceptoTitanFloat)
		fmt.Println("****************************** ", strconv.Itoa(i), " ******************************")
		if i == 0 {
			// Buscamos concepto kronos homologado
			conceptoKronosHomologado := HomologacionConcepto{ConceptoTitan: idConceptoTitan, Vigencia: registroPresupuestal.Vigencia}
			err = o.Read(&conceptoKronosHomologado, "ConceptoTitan", "Vigencia")
			if err != nil {
				alerta.Type = "error"
				alerta.Code = "E_OPN_02_3"
				alerta.Body = strconv.Itoa(idConceptoTitan)
				return
			}
			fmt.Println("***Priner append ***")
			fmt.Println("Concepto Titan: ", strconv.Itoa(idConceptoTitan))
			fmt.Println(valorCalculado)
			fmt.Println("Concepto Kronos:", strconv.Itoa(conceptoKronosHomologado.ConceptoKronos.Id))
			newConceptoValor := ConceptoValor{
				Valor:    valorCalculado,
				Concepto: &Concepto{Id: conceptoKronosHomologado.ConceptoKronos.Id},
			}
			allConceptoValor = append(allConceptoValor, newConceptoValor)
		} else {
			fmt.Println("***Sumar o Append***")
			// Buscamos concepto kronos homologado
			conceptoKronosHomologado := HomologacionConcepto{ConceptoTitan: idConceptoTitan, Vigencia: registroPresupuestal.Vigencia}
			err = o.Read(&conceptoKronosHomologado, "ConceptoTitan", "Vigencia")
			if err != nil {
				alerta.Type = "error"
				alerta.Code = "E_OPN_02_3"
				alerta.Body = strconv.Itoa(idConceptoTitan)
				return
			}
			fmt.Println("Concepto Titan: ", strconv.Itoa(idConceptoTitan))
			fmt.Println(valorCalculado)
			fmt.Println("Concepto Kronos:", strconv.Itoa(conceptoKronosHomologado.ConceptoKronos.Id))

			if esta, idlista := estaConceptoValor(conceptoKronosHomologado.ConceptoKronos.Id, allConceptoValor); esta == true {
				fmt.Println("---Sumar")
				sumaValor := allConceptoValor[idlista].Valor + valorCalculado
				allConceptoValor[idlista].Valor = sumaValor
			} else {
				fmt.Println("---Append")
				newConceptoValor2 := ConceptoValor{
					Valor:    valorCalculado,
					Concepto: &Concepto{Id: conceptoKronosHomologado.ConceptoKronos.Id},
				}
				allConceptoValor = append(allConceptoValor, newConceptoValor2)
			}
		}
	}
	//**
	//trabajar con toda la estructura de RegistroPresupuestalDisponibilidadApropiacion
	// **

	// buscamos los registroPresupuestalDisponibilidadApropiacion asociados al RP
	qs := o.QueryTable(new(RegistroPresupuestalDisponibilidadApropiacion)).RelatedSel(5)
	qs = qs.Filter("RegistroPresupuestal", registroPresupuestal.Id)
	qs = qs.RelatedSel(5)
	var l []RegistroPresupuestalDisponibilidadApropiacion
	if _, err = qs.Limit(-1, 0).All(&l); err == nil {
		for _, v := range l {
			fmt.Println("*********** Id que necesita miguel ", v.Id)
			registroPresupuestalDisponibilidadApropiacion := RegistroPresupuestalDisponibilidadApropiacion{}
			disponibilidadApropiacion := DisponibilidadApropiacion{}
			apropiacion := Apropiacion{}
			rubro := Rubro{}
			err = utilidades.FillStruct(v, &registroPresupuestalDisponibilidadApropiacion)
			err = utilidades.FillStruct(v.DisponibilidadApropiacion, &disponibilidadApropiacion)
			err = utilidades.FillStruct(disponibilidadApropiacion.Apropiacion, &apropiacion)
			err = utilidades.FillStruct(apropiacion.Rubro, &rubro)
			if err != nil {
				alerta.Type = "error"
				alerta.Code = "E_OPN_02_3"
				alerta.Body = err.Error()
				return
			}
			fmt.Println(rubro)
			//conceptos
			qsc := o.QueryTable(new(Concepto)).RelatedSel(5)
			qsc = qsc.Filter("Rubro", rubro.Id)
			qsc = qsc.RelatedSel(5)
			var lc []Concepto
			var allConceptoPorRubro []ConceptoValor
			var add_concepto bool
			if _, err = qsc.Limit(-1, 0).All(&lc); err == nil {
				for _, vc := range lc {
					fmt.Println(vc.Codigo)
					// comparar
					if esta, idlista := estaConceptoValor(vc.Id, allConceptoValor); esta == true {
						fmt.Println("Esta")
						add_concepto = true
						allConceptoPorRubro = append(allConceptoPorRubro, allConceptoValor[idlista])
					}
				}
				fmt.Println("resultado: ", add_concepto)
				if add_concepto {
					newRpCdpRubroConceptoValor := RpCdpRubroConceptoValor{
						RegistroPresupuestalDisponibilidadApropiacion: registroPresupuestalDisponibilidadApropiacion,
						ConceptoValor:                                 allConceptoPorRubro,
					}
					//
					allRpCdpRubroConceptoValor = append(allRpCdpRubroConceptoValor, newRpCdpRubroConceptoValor)
				}
			}
		}
	}
	//
	return
}

func estaConceptoValor(idConcepto int, lista []ConceptoValor) (esta bool, idlista int) {
	for or := 0; or < len(lista); or++ {
		if lista[or].Concepto.Id == idConcepto {
			return true, or
		}
	}
	return false, 0
}
