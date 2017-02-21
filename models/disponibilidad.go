package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Disponibilidad struct {
	Id                   int                   `orm:"column(id);pk"`
	UnidadEjecutora      *UnidadEjecutora      `orm:"column(unidad_ejecutora);rel(fk)"`
	Vigencia             float64               `orm:"column(vigencia)"`
	NumeroDisponibilidad float64               `orm:"column(numero_disponibilidad);null"`
	Responsable          int64                 `orm:"column(responsable);null"`
	Solicitante          int64                 `orm:"column(solicitante);null"`
	FechaRegistro        time.Time             `orm:"column(fecha_registro);type(date);null"`
	ModalidadGiro        int16                 `orm:"column(modalidad_giro);null"`
	Estado               *EstadoDisponibilidad `orm:"column(estado);rel(fk)"`
	NumeroOficio         string                `orm:"column(numero_oficio);null"`
	Objeto               string                `orm:"column(objeto);null"`
	VigenciaFutura       float64               `orm:"column(vigencia_futura);null"`
	Destino              int64                 `orm:"column(destino);null"`
	Solicitud            int64                 `orm:"column(solicitud)"`
}

func (t *Disponibilidad) TableName() string {
	return "disponibilidad"
}

func init() {
	orm.RegisterModel(new(Disponibilidad))
}

// AddDisponibilidad insert a new Disponibilidad into database and returns
// last inserted Id on success.
func AddDisponibilidad(m *Disponibilidad) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetDisponibilidadById retrieves Disponibilidad by Id. Returns error if
// Id doesn't exist
func GetDisponibilidadById(id int) (v *Disponibilidad, err error) {
	o := orm.NewOrm()
	v = &Disponibilidad{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllDisponibilidad retrieves all Disponibilidad matches certain condition. Returns empty list if
// no records exist
func GetAllDisponibilidad(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Disponibilidad))
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

	var l []Disponibilidad
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

// UpdateDisponibilidad updates Disponibilidad by Id and returns error if
// the record to be updated doesn't exist
func UpdateDisponibilidadById(m *Disponibilidad) (err error) {
	o := orm.NewOrm()
	v := Disponibilidad{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteDisponibilidad deletes Disponibilidad by Id and returns error if
// the record to be deleted doesn't exist
func DeleteDisponibilidad(id int) (err error) {
	o := orm.NewOrm()
	v := Disponibilidad{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Disponibilidad{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
