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

type OrdenDevolucion struct {
	Id              int              `orm:"column(id);pk;auto"`
	Observaciones   string           `orm:"column(observaciones)"`
	ValorTotal      float64          `orm:"column(valor_total)"`
	UnidadEjecutora *UnidadEjecutora `orm:"column(unidad_ejecutora);rel(fk)"`
	FechaRegistro   time.Time        `orm:"column(fecha_registro);auto_now_add;type(datetime)"`
	Vigencia        float64          `orm:"column(vigencia)"`
}

func (t *OrdenDevolucion) TableName() string {
	return "orden_devolucion"
}

func init() {
	orm.RegisterModel(new(OrdenDevolucion))
}

// AddOrdenDevolucion insert a new OrdenDevolucion into database and returns
// last inserted Id on success.
func AddOrdenDevolucion(m *OrdenDevolucion) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetOrdenDevolucionById retrieves OrdenDevolucion by Id. Returns error if
// Id doesn't exist
func GetOrdenDevolucionById(id int) (v *OrdenDevolucion, err error) {
	o := orm.NewOrm()
	v = &OrdenDevolucion{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllOrdenDevolucion retrieves all OrdenDevolucion matches certain condition. Returns empty list if
// no records exist
func GetAllOrdenDevolucion(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(OrdenDevolucion)).RelatedSel()
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

	var l []OrdenDevolucion
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

// UpdateOrdenDevolucion updates OrdenDevolucion by Id and returns error if
// the record to be updated doesn't exist
func UpdateOrdenDevolucionById(m *OrdenDevolucion) (err error) {
	o := orm.NewOrm()
	v := OrdenDevolucion{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteOrdenDevolucion deletes OrdenDevolucion by Id and returns error if
// the record to be deleted doesn't exist
func DeleteOrdenDevolucion(id int) (err error) {
	o := orm.NewOrm()
	v := OrdenDevolucion{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&OrdenDevolucion{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

//Add devolution order if fails returns error
func AddDevolutionOrder(request map[string]interface{}) (orden OrdenDevolucion, err error) {
	var ordenSolicitudes []*OrdenDevolucionSolicitudDevolucion
	var Id int64
	var ordenestado OrdenDevolucionEstadoDevolucion
	var estado EstadoDevolucion
	var solicitudEstado SolicitudDevolucionEstadoDevolucion
	var estadoSolicitud EstadoDevolucion

	o := orm.NewOrm()

	err = formatdata.FillStruct(request["ordenSolicitud"], &ordenSolicitudes)
	if err != nil {
		beego.Error(err)
		return
	}
	err = formatdata.FillStruct(request["ordenDevolucion"], &orden)
	if err != nil {
		beego.Error(err)
		return
	}
	err = formatdata.FillStruct(request["estadoOrdenDevol"], &estado)

	if err != nil {
		beego.Error(err)
		return
	} else {
		ordenestado.EstadoDevolucion = &estado
	}
	o.Begin()
	Id, err = o.Insert(&orden)
	if err != nil {
		o.Rollback()
		beego.Error(err)
		return
	} else {
		orden.Id = int(Id)
	}
	ordenestado.Devolucion = &orden
	for _, v := range ordenSolicitudes {

		v.OrdenDevolucion = &orden

		_, err = o.QueryTable("solicitud_devolucion_estado_devolucion").
			Filter("devolucion", v.SolicitudDevolucion.Id).
			Filter("activo", true).
			Update(orm.Params{
				"activo": "false",
			})

		if err != nil {
			o.Rollback()
			beego.Error(err)
			return
		}
		solicitudEstado.Devolucion = v.SolicitudDevolucion
		estadoSolicitud.Id = 7
		solicitudEstado.EstadoDevolucion = &estadoSolicitud
		solicitudEstado.Activo = true

		_, err = o.Insert(&solicitudEstado)
		if err != nil {
			beego.Error(err)
			o.Rollback()
			return
		}

	}
	_, err = o.InsertMulti(10, ordenSolicitudes)
	if err != nil {
		o.Rollback()
		beego.Error(err)
		return
	}
	o.Commit()
	return
}
