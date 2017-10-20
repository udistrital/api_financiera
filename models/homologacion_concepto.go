package models

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type ConceptoValor struct {
	Concepto *Concepto
	Valor    int64
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
func HomolgacionConceptosTitan(DataOpProveedor []interface{}) (alerta Alert, err error, consecutivoOp int) {
	fmt.Println("Model HomolgacionConceptosTitan")
	o := orm.NewOrm()
	var allConceptoValor []ConceptoValor

	for i, row := range DataOpProveedor {
		m := row.(map[string]interface{})
		//fmt.Println(m["Id"])
		// data valorCalculado
		valorCalculadoFloat := m["ValorCalculado"].(float64)
		valorCalculado := int64(valorCalculadoFloat)
		// data concepto
		concepto, e := m["Concepto"].(map[string]interface{})
		if e != true {
			alerta.Type = "error"
			alerta.Code = "E_OPN_01_2"
			alerta.Body = err.Error()
			return
		}
		idConceptoTitanFloat := concepto["Id"].(float64)
		idConceptoTitan := int(idConceptoTitanFloat)
		fmt.Println("****************************** ", strconv.Itoa(i), " ******************************")
		if i == 0 {
			// Buscamos concepto kronos homologado
			conceptoKronosHomologado := HomologacionConcepto{ConceptoTitan: idConceptoTitan, Vigencia: 2017} //parametro
			err = o.Read(&conceptoKronosHomologado, "ConceptoTitan", "Vigencia")
			if err != nil {
				alerta.Type = "error"
				alerta.Code = "E_OPN_02_3"
				alerta.Body = strconv.Itoa(idConceptoTitan)
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
			conceptoKronosHomologado := HomologacionConcepto{ConceptoTitan: idConceptoTitan, Vigencia: 2017} //parametro
			err = o.Read(&conceptoKronosHomologado, "ConceptoTitan", "Vigencia")
			if err != nil {
				alerta.Type = "error"
				alerta.Code = "E_OPN_02_3"
				alerta.Body = strconv.Itoa(idConceptoTitan)
				o.Rollback()
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
	} //for
	// fin  data retornar
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
