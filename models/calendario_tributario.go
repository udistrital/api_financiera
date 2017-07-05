package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/udistrital/core_api/utilidades"
)

type CalendarioTributario struct {
	Id               int               `orm:"column(id);pk;auto"`
	Descripcion      string            `orm:"column(descripcion)"`
	FechaInicio      time.Time         `orm:"column(fecha_inicio);type(date)"`
	FechaFin         time.Time         `orm:"column(fecha_fin);type(date)"`
	Vigencia         float64           `orm:"column(vigencia)"`
	ValorGirado      float64           `orm:"column(valor_girado);null"`
	Entidad          *Entidad          `orm:"column(entidad);rel(fk)"`
	EstadoCalendario *EstadoCalendario `orm:"column(estado_calendario);rel(fk)"`
	Responsable      int64             `orm:"column(responsable)"`
}

type MovsCalendario struct {
	Impuesto    CuentaEspecial
	Movimientos []MovimientoContable
}

func (t *CalendarioTributario) TableName() string {
	return "calendario_tributario"
}

func init() {
	orm.RegisterModel(new(CalendarioTributario))
}

// AddCalendarioTributario insert a new CalendarioTributario into database and returns
// last inserted Id on success.
func AddCalendarioTributario(m *CalendarioTributario) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetCalendarioTributarioById retrieves CalendarioTributario by Id. Returns error if
// Id doesn't exist
func GetCalendarioTributarioById(id int) (v *CalendarioTributario, err error) {
	o := orm.NewOrm()
	v = &CalendarioTributario{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllCalendarioTributario retrieves all CalendarioTributario matches certain condition. Returns empty list if
// no records exist
func GetAllCalendarioTributario(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(CalendarioTributario)).RelatedSel()
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

	var l []CalendarioTributario
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

// UpdateCalendarioTributario updates CalendarioTributario by Id and returns error if
// the record to be updated doesn't exist
func UpdateCalendarioTributarioById(m *CalendarioTributario) (err error) {
	o := orm.NewOrm()
	v := CalendarioTributario{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteCalendarioTributario deletes CalendarioTributario by Id and returns error if
// the record to be deleted doesn't exist
func DeleteCalendarioTributario(id int) (err error) {
	o := orm.NewOrm()
	v := CalendarioTributario{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&CalendarioTributario{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

//GetImpuestosCalendario informe  movimientos del calendario
func GetImpuestosCalendario(idcalendario int) (calendario interface{}, err error) {
	var info_calendario CalendarioTributario
	//var movimientos []MovimientoContable
	//var cuentas []CuentaContable
	info_calendario = CalendarioTributario{Id: idcalendario}
	//info_calendario.Id = idcalendario
	o := orm.NewOrm()
	var cuentas []int
	var movs []MovsCalendario
	var calmov MovsCalendario
	if err = o.Read(&info_calendario); err == nil {
		/*o.QueryTable(new(MovimientoContable)).Filter("fecha__gte", info_calendario.FechaInicio.Format("2006-01-2 ")+"23:59:59").Filter("fecha__lte", info_calendario.FechaFin.Format("2006-01-2 ")+"23:59:59").Filter("aprobado", true).RelatedSel().All(&movimientos)*/

		o.Raw(`SELECT M.cuenta_contable
		FROM financiera.movimiento_contable M inner join  (select * from financiera.cuenta_especial where tipo_cuenta_especial = 2) E
		on E.cuenta_contable = M.cuenta_contable where M.aprobado = true and
		M.fecha::DATE BETWEEN ?::DATE AND ?::DATE group by M.cuenta_contable`, info_calendario.FechaInicio, info_calendario.FechaFin).QueryRows(&cuentas)
		fmt.Println(cuentas)
		for index := range cuentas {
			//var idcuenta int
			//idcuenta := cuentas[index]["cuenta_contable"].(int)
			//fmt.Println("int:", cuentas[index]["cuenta_contable"])
			//err = utilidades.FillStruct(cuentas[index]["cuenta_contable"], &idcuenta)
			fmt.Println("cuenta:", cuentas[index])
			cm := calmov

			o.QueryTable(new(CuentaEspecial)).Filter("cuenta_contable", cuentas[index]).RelatedSel(5).All(&cm.Impuesto)
			//fmt.Println(&cm.CuentaContable)
			o.Raw(`SELECT M.*
			FROM financiera.movimiento_contable M
			where M.cuenta_contable = ? and M.aprobado = true and
			M.fecha::DATE BETWEEN ?::DATE AND ?::DATE `, cm.Impuesto.CuentaContable, info_calendario.FechaInicio, info_calendario.FechaFin).QueryRows(&cm.Movimientos)

			for index := range cm.Movimientos {
				o.QueryTable(new(MovimientoContable)).Filter("id", &cm.Movimientos[index].Id).RelatedSel("concepto", "tipo_documento_afectante").All(&cm.Movimientos[index])
			}

			movs = append(movs, cm)

			//o.QueryTable(new(MovimientoContable)).Filter("id", movimientos[index].Id).RelatedSel("concepto", "tipo_documento_afectante").All(&movimientos[index])
		}

		/*fmt.Println(movimientos)*/
		err = utilidades.FillStruct(movs, &calendario)

	}
	return
}
