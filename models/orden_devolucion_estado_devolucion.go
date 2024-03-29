package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/udistrital/utils_oas/formatdata"
)

type OrdenDevolucionEstadoDevolucion struct {
	Id               int               `orm:"column(id);pk;auto"`
	Devolucion       *OrdenDevolucion  `orm:"column(devolucion);rel(fk)"`
	Activo           bool              `orm:"column(activo)"`
	FechaRegistro    time.Time         `orm:"column(fecha_registro);auto_now_add;type(datetime)"`
	EstadoDevolucion *EstadoDevolucion `orm:"column(estado_devolucion);rel(fk)"`
}

func (t *OrdenDevolucionEstadoDevolucion) TableName() string {
	return "orden_devolucion_estado_devolucion"
}

func init() {
	orm.RegisterModel(new(OrdenDevolucionEstadoDevolucion))
}

// AddOrdenDevolucionEstadoDevolucion insert a new OrdenDevolucionEstadoDevolucion into database and returns
// last inserted Id on success.
func AddOrdenDevolucionEstadoDevolucion(m *OrdenDevolucionEstadoDevolucion) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetOrdenDevolucionEstadoDevolucionById retrieves OrdenDevolucionEstadoDevolucion by Id. Returns error if
// Id doesn't exist
func GetOrdenDevolucionEstadoDevolucionById(id int) (v *OrdenDevolucionEstadoDevolucion, err error) {
	o := orm.NewOrm()
	v = &OrdenDevolucionEstadoDevolucion{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllOrdenDevolucionEstadoDevolucion retrieves all OrdenDevolucionEstadoDevolucion matches certain condition. Returns empty list if
// no records exist
func GetAllOrdenDevolucionEstadoDevolucion(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(OrdenDevolucionEstadoDevolucion)).RelatedSel()
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

	var l []OrdenDevolucionEstadoDevolucion
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

// UpdateOrdenDevolucionEstadoDevolucion updates OrdenDevolucionEstadoDevolucion by Id and returns error if
// the record to be updated doesn't exist
func UpdateOrdenDevolucionEstadoDevolucionById(m *OrdenDevolucionEstadoDevolucion) (err error) {
	o := orm.NewOrm()
	v := OrdenDevolucionEstadoDevolucion{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteOrdenDevolucionEstadoDevolucion deletes OrdenDevolucionEstadoDevolucion by Id and returns error if
// the record to be deleted doesn't exist
func DeleteOrdenDevolucionEstadoDevolucion(id int) (err error) {
	o := orm.NewOrm()
	v := OrdenDevolucionEstadoDevolucion{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&OrdenDevolucionEstadoDevolucion{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func AddEstadoOrden(request map[string]interface{}) (ordenEstado OrdenDevolucionEstadoDevolucion, err error) {
	var orden OrdenDevolucion
	var solicitudesOrden []*OrdenDevolucionSolicitudDevolucion
	o := orm.NewOrm()
	o.Begin()
	err = formatdata.FillStruct(request["estadoOrdenDevol"], &ordenEstado)
	err = formatdata.FillStruct(request["ordenDevolucion"], &orden)
	if err == nil {
		_, err = o.QueryTable("orden_devolucion_estado_devolucion").
			Filter("devolucion", orden.Id).
			Filter("activo", true).
			Update(orm.Params{
				"activo": "false",
			})
		if err != nil {
			beego.Error(err)
			o.Rollback()
			return
		}
		ordenEstado.Devolucion = &orden
		ordenEstado.Activo = true
		_, err = o.Insert(&ordenEstado)
		if err != nil {
			beego.Error(err)
			o.Rollback()
			return
		}

		if ordenEstado.EstadoDevolucion.Id == 4 {
			beego.Info("Aprobacion contable")
			_, err = o.QueryTable(new(OrdenDevolucionSolicitudDevolucion)).
				Filter("OrdenDevolucion", orden.Id).
				All(&solicitudesOrden)
			for _, v := range solicitudesOrden {
				_, err = o.QueryTable("movimiento_contable").Filter("tipo_documento_afectante", 5).Filter("codigo_documento_afectante", v.SolicitudDevolucion.Id).Update(orm.Params{
					"estado_movimiento_contable": 2,
				})
				if err != nil {
					beego.Error(err.Error())
					o.Rollback()
					return
				}
			}

		}
		o.Commit()
	} else {
		beego.Error(err)
	}
	return
}
