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

type DevolucionTributariaEstadoDevolucion struct {
	Id               int                   `orm:"column(id);pk;auto"`
	Devolucion       *DevolucionTributaria `orm:"column(devolucion);rel(fk)"`
	Activo           bool                  `orm:"column(activo)"`
	FechaRegistro    time.Time             `orm:"column(fecha_registro);auto_now_add;type(datetime)"`
	EstadoDevolucion *EstadoDevolucion     `orm:"column(estado_devolucion);rel(fk)"`
}

func (t *DevolucionTributariaEstadoDevolucion) TableName() string {
	return "devolucion_tributaria_estado_devolucion"
}

func init() {
	orm.RegisterModel(new(DevolucionTributariaEstadoDevolucion))
}

// AddDevolucionTributariaEstadoDevolucion insert a new DevolucionTributariaEstadoDevolucion into database and returns
// last inserted Id on success.
func AddDevolucionTributariaEstadoDevolucion(m *DevolucionTributariaEstadoDevolucion) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetDevolucionTributariaEstadoDevolucionById retrieves DevolucionTributariaEstadoDevolucion by Id. Returns error if
// Id doesn't exist
func GetDevolucionTributariaEstadoDevolucionById(id int) (v *DevolucionTributariaEstadoDevolucion, err error) {
	o := orm.NewOrm()
	v = &DevolucionTributariaEstadoDevolucion{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllDevolucionTributariaEstadoDevolucion retrieves all DevolucionTributariaEstadoDevolucion matches certain condition. Returns empty list if
// no records exist
func GetAllDevolucionTributariaEstadoDevolucion(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(DevolucionTributariaEstadoDevolucion)).RelatedSel()
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

	var l []DevolucionTributariaEstadoDevolucion
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

// UpdateDevolucionTributariaEstadoDevolucion updates DevolucionTributariaEstadoDevolucion by Id and returns error if
// the record to be updated doesn't exist
func UpdateDevolucionTributariaEstadoDevolucionById(m *DevolucionTributariaEstadoDevolucion) (err error) {
	o := orm.NewOrm()
	v := DevolucionTributariaEstadoDevolucion{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteDevolucionTributariaEstadoDevolucion deletes DevolucionTributariaEstadoDevolucion by Id and returns error if
// the record to be deleted doesn't exist
func DeleteDevolucionTributariaEstadoDevolucion(id int) (err error) {
	o := orm.NewOrm()
	v := DevolucionTributariaEstadoDevolucion{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&DevolucionTributariaEstadoDevolucion{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func AddEstadoDevol(request map[string]interface{}) (devolucionEstado DevolucionTributariaEstadoDevolucion, err error) {

	var devolucionTributaria DevolucionTributaria
	o := orm.NewOrm()
	o.Begin()
	err = formatdata.FillStruct(request["EstadoDevolTribut"], &devolucionEstado)
	err = formatdata.FillStruct(request["DevolucionTributaria"], &devolucionTributaria)
	if err == nil {
		_, err = o.QueryTable("devolucion_tributaria_estado_devolucion").
			Filter("devolucion", devolucionTributaria.Id).
			Filter("activo", true).
			Update(orm.Params{
				"activo": "false",
			})
		if err != nil {
			beego.Error(err)
			o.Rollback()
			return
		}
		devolucionEstado.Devolucion = &devolucionTributaria
		devolucionEstado.Activo = true
		_, err = o.Insert(&devolucionEstado)
		if err != nil {
			beego.Error(err)
			o.Rollback()
			return
		}

		if devolucionEstado.EstadoDevolucion.Id == 11 {
			beego.Info("Aprobacion contable")
			_, err = o.QueryTable("movimiento_contable").Filter("tipo_documento_afectante", 6).Filter("codigo_documento_afectante", devolucionTributaria.Id).Update(orm.Params{
				"estado_movimiento_contable": 2,
			})
			if err != nil {
				beego.Error(err.Error())
				o.Rollback()
				return
			}

		}
		o.Commit()
	} else {
		beego.Error(err)
	}
	return
}
