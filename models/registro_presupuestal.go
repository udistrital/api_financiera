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

type RegistroPresupuestal struct {
	Id                         int                         `orm:"column(id);pk;auto"`
	UnidadEjecutora            *UnidadEjecutora            `orm:"column(unidad_ejecutora);rel(fk)"`
	Vigencia                   float64                     `orm:"column(vigencia)"`
	FechaMovimiento            time.Time                   `orm:"column(fecha_movimiento);type(date);null"`
	Responsable                int                         `orm:"column(responsable);null"`
	Estado                     *EstadoRegistroPresupuestal `orm:"column(estado);rel(fk)"`
	NumeroRegistroPresupuestal int                         `orm:"column(numero_registro_presupuestal)"`
	Beneficiario               int                         `orm:"column(beneficiario);null"`
	Compromiso                 *Compromiso                 `orm:"column(compromiso);rel(fk)"`
}
type DatosRubroRegistroPresupuestal struct {
	Id             int
	Disponibilidad *Disponibilidad
	Apropiacion    *Apropiacion
	Valor          float64
	ValorAsignado  float64
}

type DatosRegistroPresupuestal struct { //estructura temporal para el registro con relacion a las apropiaciones
	Rp     *RegistroPresupuestal
	Rubros []DatosRubroRegistroPresupuestal
}
type DatosSaldoRp struct {
	Rp          *RegistroPresupuestal
	Apropiacion *Apropiacion
}
type Info_rp_a_anular struct {
	Anulacion      AnulacionRegistroPresupuestal
	Rp_apropiacion []RegistroPresupuestalDisponibilidadApropiacion
	Valor          float64
}

func (t *RegistroPresupuestal) TableName() string {
	return "registro_presupuestal"
}

func init() {
	orm.RegisterModel(new(RegistroPresupuestal))
}

// AddRegistroPresupuestal insert a new RegistroPresupuestal into database and returns
// last inserted Id on success.
func AddRegistoPresupuestal(m *DatosRegistroPresupuestal) (id int64, err error) {
	o := orm.NewOrm()
	o.Begin()
	id, err = o.Insert(m.Rp)
	if err == nil {
		m.Rp.Id = int(id)
		for _, data := range m.Rubros {
			registro := RegistroPresupuestalDisponibilidadApropiacion{
				RegistroPresupuestal:      m.Rp,
				DisponibilidadApropiacion: &DisponibilidadApropiacion{Id: data.Id},
				Valor: data.ValorAsignado,
			}
			_, err2 := o.Insert(&registro)
			if err2 != nil {
				o.Rollback()
			}
		}
	} else {
		fmt.Println("error registro rp: ", err.Error())
	}

	o.Commit()
	return
}

// GetRegistroPresupuestalById retrieves RegistroPresupuestal by Id. Returns error if
// Id doesn't exist
func GetRegistroPresupuestalById(id int) (v *RegistroPresupuestal, err error) {
	o := orm.NewOrm()
	v = &RegistroPresupuestal{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllRegistroPresupuestal retrieves all RegistroPresupuestal matches certain condition. Returns empty list if
// no records exist
func GetAllRegistroPresupuestal(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(RegistroPresupuestal))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
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

	var l []RegistroPresupuestal
	qs = qs.OrderBy(sortFields...).RelatedSel(5)
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

// UpdateRegistroPresupuestal updates RegistroPresupuestal by Id and returns error if
// the record to be updated doesn't exist
func UpdateRegistroPresupuestalById(m *RegistroPresupuestal) (err error) {
	o := orm.NewOrm()
	v := RegistroPresupuestal{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteRegistroPresupuestal deletes RegistroPresupuestal by Id and returns error if
// the record to be deleted doesn't exist
func DeleteRegistroPresupuestal(id int) (err error) {
	o := orm.NewOrm()
	v := RegistroPresupuestal{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&RegistroPresupuestal{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

//----------------------------------------
//funcion para obtener saldo restante del cdp
func SaldoRp(id_rp int, id_apropiacion int) (valor float64, err error) {
	o := orm.NewOrm()
	var maps []orm.Params
	o.Raw(`SELECT * FROM financiera.saldo_rp WHERE id = ? AND apropiacion = ? `, id_rp, id_apropiacion).Values(&maps)
	fmt.Println("maps: ", maps)
	if maps[0]["valor"] == nil {
		valor = 0
	} else {
		valor, err = strconv.ParseFloat(maps[0]["valor"].(string), 64)
	}

	return
}

//----------------------------------------

//----------------------------------------
//funcion para realizar anulacion total en el RP
func AnulacionTotalRp(m *Info_rp_a_anular) (alerta []string, err error) {
	o := orm.NewOrm()
	o.Begin()
	alerta = append(alerta, "success")
	m.Anulacion.FechaRegistro = time.Now()
	id_anulacion_rp, err1 := o.Insert(&m.Anulacion)
	fmt.Println("error")
	if err1 != nil {
		alerta[0] = "error"
		alerta = append(alerta, "No se pudo registrar el detalle de la anulacion")
		err = err1
		o.Rollback()
		return
	}
	for i := 0; i < len(m.Rp_apropiacion); i++ {

		saldoRp, err2 := SaldoRp(m.Rp_apropiacion[i].RegistroPresupuestal.Id, m.Rp_apropiacion[i].DisponibilidadApropiacion.Apropiacion.Id)
		if err2 != nil {
			alerta[0] = "error"
			alerta = append(alerta, "No se pudo cargar el saldo del RP N° "+strconv.Itoa(m.Rp_apropiacion[i].RegistroPresupuestal.NumeroRegistroPresupuestal)+" para la apropiacion del Rubro "+m.Rp_apropiacion[i].DisponibilidadApropiacion.Apropiacion.Rubro.Codigo)
			err = err2
			o.Rollback()
			return
		}
		anulacion_apropiacion := AnulacionRegistroPresupuestalDisponibilidadApropiacion{
			AnulacionRegistroPresupuestal:                 &AnulacionRegistroPresupuestal{Id: int(id_anulacion_rp)},
			RegistroPresupuestalDisponibilidadApropiacion: &m.Rp_apropiacion[i],
			Valor: saldoRp,
		}
		_, err3 := o.Insert(&anulacion_apropiacion)
		if err3 != nil {
			alerta[0] = "error"
			alerta = append(alerta, "No se pudo registrar la anulacion del RP N° "+strconv.Itoa(m.Rp_apropiacion[i].RegistroPresupuestal.NumeroRegistroPresupuestal)+" para la apropiacion del Rubro "+m.Rp_apropiacion[i].DisponibilidadApropiacion.Apropiacion.Rubro.Codigo)
			err = err3
			o.Rollback()
			return
		} else {
			alerta = append(alerta, "se anulo del RP N° "+strconv.Itoa(m.Rp_apropiacion[i].RegistroPresupuestal.NumeroRegistroPresupuestal)+" para la apropiacion del Rubro "+m.Rp_apropiacion[i].DisponibilidadApropiacion.Apropiacion.Rubro.Codigo+" la suma de "+strconv.FormatFloat(saldoRp, 'f', -1, 64))

		}
	}

	o.Commit()
	return
}

//--------------------------------------------------------
//funcion para realizar la anulacion parcial del RP

func AnulacionParcialRp(m *Info_rp_a_anular) (alerta []string, err error) {
	o := orm.NewOrm()
	o.Begin()
	alerta = append(alerta, "success")
	m.Anulacion.FechaRegistro = time.Now()
	id_anulacion_rp, err1 := o.Insert(&m.Anulacion)
	fmt.Println("error")
	if err1 != nil {
		alerta = append(alerta, "No se pudo registrar el detalle de la anulacion")
		alerta[0] = "error"
		err = err1
		o.Rollback()
		return
	}
	for i := 0; i < len(m.Rp_apropiacion); i++ {

		anulacion_apropiacion := AnulacionRegistroPresupuestalDisponibilidadApropiacion{
			AnulacionRegistroPresupuestal:                 &AnulacionRegistroPresupuestal{Id: int(id_anulacion_rp)},
			RegistroPresupuestalDisponibilidadApropiacion: &m.Rp_apropiacion[i],
			Valor: m.Valor,
		}
		_, err3 := o.Insert(&anulacion_apropiacion)
		if err3 != nil {
			alerta[0] = "error"
			alerta = append(alerta, "No se pudo registrar la anulacion del RP N° "+strconv.Itoa(m.Rp_apropiacion[i].RegistroPresupuestal.NumeroRegistroPresupuestal)+" para la apropiacion del Rubro "+m.Rp_apropiacion[i].DisponibilidadApropiacion.Apropiacion.Rubro.Codigo)
			err = err3
			o.Rollback()
			return
		} else {
			alerta = append(alerta, "se anulo del RP N° "+strconv.Itoa(m.Rp_apropiacion[i].RegistroPresupuestal.NumeroRegistroPresupuestal)+" para la apropiacion del Rubro "+m.Rp_apropiacion[i].DisponibilidadApropiacion.Apropiacion.Rubro.Codigo+" la suma de "+strconv.FormatFloat(m.Valor, 'f', -1, 64))
		}
	}

	o.Commit()
	return
}

//funcion GetValorTotalRp
func GetValorTotalRp(rp_id int) (total float64, err error){
	o := orm.NewOrm()
	var totalSql float64
	err = o.Raw("select sum(valor) from registro_presupuestal_disponibilidad_apropiacion where registro_presupuestal = ?", rp_id).QueryRow(&totalSql)
	if err == nil {
		return totalSql, nil
	}
	return totalSql, err
}
