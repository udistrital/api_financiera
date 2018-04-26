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

type InversionEstadoInversion struct {
	Id            int              `orm:"column(id);pk;auto"`
	Inversion     *Inversion       `orm:"column(inversion);rel(fk)"`
	Estado        *EstadoInversion `orm:"column(estado);rel(fk)"`
	FechaRegistro time.Time        `orm:"column(fecha_registro);auto_now_add;type(datetime)"`
	Usuario       int           		`orm:"column(usuario)"`
	Activo        bool             `orm:"column(activo)"`
}

func (t *InversionEstadoInversion) TableName() string {
	return "inversion_estado_inversion"
}

func init() {
	orm.RegisterModel(new(InversionEstadoInversion))
}

// AddInversionEstadoInversion insert a new InversionEstadoInversion into database and returns
// last inserted Id on success.
func AddInversionEstadoInversion(m *InversionEstadoInversion) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetInversionEstadoInversionById retrieves InversionEstadoInversion by Id. Returns error if
// Id doesn't exist
func GetInversionEstadoInversionById(id int) (v *InversionEstadoInversion, err error) {
	o := orm.NewOrm()
	v = &InversionEstadoInversion{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllInversionEstadoInversion retrieves all InversionEstadoInversion matches certain condition. Returns empty list if
// no records exist
func GetAllInversionEstadoInversion(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(InversionEstadoInversion)).RelatedSel()
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

	var l []InversionEstadoInversion
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

// UpdateInversionEstadoInversion updates InversionEstadoInversion by Id and returns error if
// the record to be updated doesn't exist
func UpdateInversionEstadoInversionById(m *InversionEstadoInversion) (err error) {
	o := orm.NewOrm()
	v := InversionEstadoInversion{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteInversionEstadoInversion deletes InversionEstadoInversion by Id and returns error if
// the record to be deleted doesn't exist
func DeleteInversionEstadoInversion(id int) (err error) {
	o := orm.NewOrm()
	v := InversionEstadoInversion{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&InversionEstadoInversion{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func AddEstadoInv(request map[string]interface{}) (invEstadoinversion InversionEstadoInversion, err error) {
	var invEstado EstadoInversion
	var invEstadoPadre EstadoInversion
	var inversion Inversion
	var idEstadoInv int64
	var num int64
	var usuario int

	o := orm.NewOrm()
	o.Begin()
	err = formatdata.FillStruct(request["Estado"], &invEstado)
	err = formatdata.FillStruct(request["EstadoPadre"], &invEstadoPadre)
	err = formatdata.FillStruct(request["Inversion"], &inversion)
	err = formatdata.FillStruct(request["Usuario"], &usuario)

	if err == nil {
		invEstadoinversion.Activo = true
		invEstadoinversion.Estado = &invEstado
		invEstadoinversion.Inversion = &inversion
		invEstadoinversion.Usuario = usuario

		num, err = o.QueryTable("inversion_estado_inversion").Filter("estado", invEstadoPadre.Id).Filter("inversion", inversion.Id).Update(orm.Params{
			"activo": "false",
		})

		fmt.Printf("Affected Num: %s, %s", num, err)

		if err != nil {
			beego.Error(err.Error())
			o.Rollback()
			return
		}

		if invEstado.Id == 4 {
			beego.Error("Aprobacion contable")
			num, err = o.QueryTable("movimiento_contable").Filter("tipo_documento_afectante", 3).Filter("codigo_documento_afectante", inversion.Id).Update(orm.Params{
				"estado_movimiento_contable": 2,
			})
			if err != nil {
				beego.Error(err.Error())
				o.Rollback()
				return
			}
		}

		idEstadoInv, err = o.Insert(&invEstadoinversion)
		invEstadoinversion.Id = int(idEstadoInv)

		if err != nil {
			beego.Error(err.Error())
			o.Rollback()
			return
		}
		o.Commit()
	}

	return
}
