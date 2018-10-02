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

type EstadoLegalizacionAvanceLegalizacion struct {
	Id                 int                 `orm:"column(id);pk;auto"`
	AvanceLegalizacion *AvanceLegalizacion `orm:"column(avance_legalizacion);rel(fk)"`
	Activo             bool                `orm:"column(activo)"`
	FechaRegistro      time.Time           `orm:"column(fecha_registro);auto_now_add;type(datetime)"`
	Estado             *EstadoLegalizacion `orm:"column(estado);rel(fk)"`
	Usuario            int                 `orm:"column(usuario);null"`
}

func (t *EstadoLegalizacionAvanceLegalizacion) TableName() string {
	return "estado_legalizacion_avance_legalizacion"
}

func init() {
	orm.RegisterModel(new(EstadoLegalizacionAvanceLegalizacion))
}

// AddEstadoLegalizacionAvanceLegalizacion insert a new EstadoLegalizacionAvanceLegalizacion into database and returns
// last inserted Id on success.
func AddEstadoLegalizacionAvanceLegalizacion(m *EstadoLegalizacionAvanceLegalizacion) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetEstadoLegalizacionAvanceLegalizacionById retrieves EstadoLegalizacionAvanceLegalizacion by Id. Returns error if
// Id doesn't exist
func GetEstadoLegalizacionAvanceLegalizacionById(id int) (v *EstadoLegalizacionAvanceLegalizacion, err error) {
	o := orm.NewOrm()
	v = &EstadoLegalizacionAvanceLegalizacion{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllEstadoLegalizacionAvanceLegalizacion retrieves all EstadoLegalizacionAvanceLegalizacion matches certain condition. Returns empty list if
// no records exist
func GetAllEstadoLegalizacionAvanceLegalizacion(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(EstadoLegalizacionAvanceLegalizacion))
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

	var l []EstadoLegalizacionAvanceLegalizacion
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				_, err = o.LoadRelated(&v, "estado")
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

// UpdateEstadoLegalizacionAvanceLegalizacion updates EstadoLegalizacionAvanceLegalizacion by Id and returns error if
// the record to be updated doesn't exist
func UpdateEstadoLegalizacionAvanceLegalizacionById(m *EstadoLegalizacionAvanceLegalizacion) (err error) {
	o := orm.NewOrm()
	v := EstadoLegalizacionAvanceLegalizacion{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteEstadoLegalizacionAvanceLegalizacion deletes EstadoLegalizacionAvanceLegalizacion by Id and returns error if
// the record to be deleted doesn't exist
func DeleteEstadoLegalizacionAvanceLegalizacion(id int) (err error) {
	o := orm.NewOrm()
	v := EstadoLegalizacionAvanceLegalizacion{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&EstadoLegalizacionAvanceLegalizacion{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func CreateEstadoIniAvanceLegalizacion(request map[string]interface{}) (err error) {
	var avanceLegalizacion AvanceLegalizacion
	var usuario int64
	var estadoLegalizacion EstadoLegalizacion
	var estadoLegalizacionAvance EstadoLegalizacionAvanceLegalizacion
	o := orm.NewOrm()
	err = formatdata.FillStruct(request["AvanceLegalizacion"], &avanceLegalizacion)
	beego.Error("request ", request)
	usuario = request["Usuario"].(int64)
	beego.Error("usuario  ", usuario)
	if err != nil {
		beego.Error(err.Error())
		return
	}
	o.Begin()
	_, err = o.QueryTable("estado_legalizacion_avance_legalizacion").
		Filter("avance_legalizacion", avanceLegalizacion.Id).
		Filter("activo", true).
		Update(orm.Params{
			"activo": "false",
		})
	if err != nil {
		beego.Error(err)
		o.Rollback()
		return
	}
	err = o.QueryTable(new(EstadoLegalizacion)).
		Filter("NumeroOrden", 1).
		One(&estadoLegalizacion)
	if err != nil {
		beego.Error(err.Error())
		o.Rollback()
		return
	}

	estadoLegalizacionAvance.AvanceLegalizacion = &avanceLegalizacion
	estadoLegalizacionAvance.Activo = true
	estadoLegalizacionAvance.Usuario = int(usuario)
	estadoLegalizacionAvance.Estado = &estadoLegalizacion
	_, err = o.Insert(&estadoLegalizacionAvance)
	if err != nil {
		beego.Error(err)
		o.Rollback()
		return
	}
	o.Commit()
	return
}

func AddEstadoLegalizacionTipo(request map[string]interface{}) (legalizacionEstado EstadoLegalizacionAvanceLegalizacion, err error) {
	var usuario float64
	var avanceLegalizacion AvanceLegalizacion
	var legalizacionesTipo []AvanceLegalizacionTipo
	var noDocAfectante int
	var tipoDocAfectante TipoDocumentoAfectante
	o := orm.NewOrm()

	err = formatdata.FillStruct(request["EstadoLegalizacion"], &legalizacionEstado)
	err = formatdata.FillStruct(request["AvanceLegalizacion"], &avanceLegalizacion)
	usuario = request["Usuario"].(float64)
	if err != nil {
		beego.Error(err)
		return
	}
	o.Begin()
	_, err = o.QueryTable("estado_legalizacion_avance_legalizacion").
		Filter("avance_legalizacion", avanceLegalizacion.Id).
		Filter("activo", true).
		Update(orm.Params{
			"activo": "false",
		})
	if err != nil {
		beego.Error(err)
		o.Rollback()
		return
	}
	legalizacionEstado.AvanceLegalizacion = &avanceLegalizacion
	legalizacionEstado.Activo = true
	legalizacionEstado.Usuario = int(usuario)
	_, err = o.Insert(&legalizacionEstado)
	if err != nil {
		beego.Error(err)
		o.Rollback()
		return
	}

	if legalizacionEstado.Estado.Id == 2 {
		_, err = o.QueryTable("avance_legalizacion_tipo").
			Filter("avance_legalizacion", avanceLegalizacion.Id).
			All(&legalizacionesTipo)

		if err != nil {
			beego.Error(err)
			o.Rollback()
			return
		}
		for _, avanceLegTipo := range legalizacionesTipo {
			switch avanceLegTipo.TipoAvanceLegalizacion.Id {
			case 1:
				noDocAfectante = 9
			case 2:
				noDocAfectante = 8
			}

			err = o.QueryTable("tipo_documento_afectante").
				Filter("numeroOrden", noDocAfectante).
				One(&tipoDocAfectante)
			if err != nil {
				beego.Error("Error", err)
				return
			}
			beego.Info("Aprobacion contable")
			_, err = o.QueryTable("movimiento_contable").
				Filter("tipo_documento_afectante", tipoDocAfectante.Id).
				Filter("codigo_documento_afectante", avanceLegTipo.Id).Update(orm.Params{
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
	return
}
