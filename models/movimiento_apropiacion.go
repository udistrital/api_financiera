package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/udistrital/api_financiera/utilidades"
	"reflect"
	"strings"
	"time"
)

type MovimientoApropiacion struct {
	Id              int       `orm:"auto;column(id);pk"`
	FechaMovimiento time.Time `orm:"column(fecha_movimiento);type(date)"`
	Noficio         int       `orm:"column(n_oficio)"`
	Foficio         time.Time `orm:"column(f_oficio);type(date)"`
	Descripcion     string    `orm:"column(descripcion);null"`
}

func (t *MovimientoApropiacion) TableName() string {
	return "movimiento_apropiacion"
}

func init() {
	orm.RegisterModel(new(MovimientoApropiacion))
}

// AddMovimientoApropiacion insert a new MovimientoApropiacion into database and returns
// last inserted Id on success.
func AddMovimientoApropiacion(m *MovimientoApropiacion) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// Registra un MovimientoApropiacion
// retorna structura de alerta
func RegistrarMovimietnoApropiaciontr(movimiento map[string]interface{}) (alert Alert, err error) {
	var movimientoapr MovimientoApropiacion
	var desgrMovimientoApr []MovimientoApropiacionDisponibilidadApropiacion
	if err = utilidades.FillStruct(movimiento["MovimientoApropiacion"], &movimientoapr); err == nil {
		if err = utilidades.FillStruct(movimiento["MovimientoApropiacionDisponibilidadApropiacion"], &desgrMovimientoApr); err == nil {
			o := orm.NewOrm()
			o.Begin()
			movimientoapr.FechaMovimiento = time.Now().Local() //asignacion de la fecha de la solicitud del movimiento
			_, err = o.Insert(&movimientoapr)
			if err != nil {
				o.Rollback()
				return
			}

			for _, datDesgrMov := range desgrMovimientoApr {
				datDesgrMov.MovimientoApropiacion = &movimientoapr
				datDesgrMov.EstadoMovimientoApropiacion = &EstadoMovimientoApropiacion{Id: 1}
				_, err = o.Insert(&datDesgrMov)
				if err != nil {
					o.Rollback()
					return
				}

			}

			o.Commit()

		} else {
			alert.Code = "E_0458"
			alert.Body = err
			alert.Type = "error"
			return
		}
	} else {
		alert.Code = "E_0458"
		alert.Body = err
		alert.Type = "error"
		return
	}
	alert.Code = "S_MODP001"
	alert.Body = map[string]interface{}{"MovimientoApropiacion": movimientoapr, "MovimientoApropiacionDisponibilidadApropiacion": desgrMovimientoApr}
	alert.Type = "success"
	return
}

// GetMovimientoApropiacionById retrieves MovimientoApropiacion by Id. Returns error if
// Id doesn't exist
func GetMovimientoApropiacionById(id int) (v *MovimientoApropiacion, err error) {
	o := orm.NewOrm()
	v = &MovimientoApropiacion{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllMovimientoApropiacion retrieves all MovimientoApropiacion matches certain condition. Returns empty list if
// no records exist
func GetAllMovimientoApropiacion(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(MovimientoApropiacion))
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

	var l []MovimientoApropiacion
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

// UpdateMovimientoApropiacion updates MovimientoApropiacion by Id and returns error if
// the record to be updated doesn't exist
func UpdateMovimientoApropiacionById(m *MovimientoApropiacion) (err error) {
	o := orm.NewOrm()
	v := MovimientoApropiacion{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteMovimientoApropiacion deletes MovimientoApropiacion by Id and returns error if
// the record to be deleted doesn't exist
func DeleteMovimientoApropiacion(id int) (err error) {
	o := orm.NewOrm()
	v := MovimientoApropiacion{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&MovimientoApropiacion{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
