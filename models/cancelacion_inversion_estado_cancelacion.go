package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/udistrital/utils_oas/formatdata"
	"github.com/astaxie/beego"
)

type CancelacionInversionEstadoCancelacion struct {
	Id                         int                         `orm:"column(id);pk;auto"`
	CancelacionInversion       *CancelacionInversion       `orm:"column(cancelacion_inversion);rel(fk)"`
	Activo                     bool                        `orm:"column(activo)"`
	FechaRegistro              time.Time                   `orm:"column(fecha_registro);auto_now_add;type(datetime)"`
	EstadoCancelacionInversion *EstadoCancelacionInversion `orm:"column(estado_cancelacion_inversion);rel(fk)"`
	Usuario                    int                         `orm:"column(usuario);null"`
}

func (t *CancelacionInversionEstadoCancelacion) TableName() string {
	return "cancelacion_inversion_estado_cancelacion"
}

func init() {
	orm.RegisterModel(new(CancelacionInversionEstadoCancelacion))
}

// AddCancelacionInversionEstadoCancelacion insert a new CancelacionInversionEstadoCancelacion into database and returns
// last inserted Id on success.
func AddCancelacionInversionEstadoCancelacion(m *CancelacionInversionEstadoCancelacion) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetCancelacionInversionEstadoCancelacionById retrieves CancelacionInversionEstadoCancelacion by Id. Returns error if
// Id doesn't exist
func GetCancelacionInversionEstadoCancelacionById(id int) (v *CancelacionInversionEstadoCancelacion, err error) {
	o := orm.NewOrm()
	v = &CancelacionInversionEstadoCancelacion{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllCancelacionInversionEstadoCancelacion retrieves all CancelacionInversionEstadoCancelacion matches certain condition. Returns empty list if
// no records exist
func GetAllCancelacionInversionEstadoCancelacion(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(CancelacionInversionEstadoCancelacion)).RelatedSel()
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

	var l []CancelacionInversionEstadoCancelacion
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

// UpdateCancelacionInversionEstadoCancelacion updates CancelacionInversionEstadoCancelacion by Id and returns error if
// the record to be updated doesn't exist
func UpdateCancelacionInversionEstadoCancelacionById(m *CancelacionInversionEstadoCancelacion) (err error) {
	o := orm.NewOrm()
	v := CancelacionInversionEstadoCancelacion{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteCancelacionInversionEstadoCancelacion deletes CancelacionInversionEstadoCancelacion by Id and returns error if
// the record to be deleted doesn't exist
func DeleteCancelacionInversionEstadoCancelacion(id int) (err error) {
	o := orm.NewOrm()
	v := CancelacionInversionEstadoCancelacion{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&CancelacionInversionEstadoCancelacion{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}


func AddEstadoCancelacionCancInversion(request map[string]interface{}) (estadoCancelacion CancelacionInversionEstadoCancelacion,err error) {
	var cancelacionInversion CancelacionInversion
	var estadoInversion EstadoInversion
	var requestEstadoInv map[string]interface{}
	var inversionPadre Inversion
	o := orm.NewOrm()
	o.Begin()
	err = formatdata.FillStruct(request["EstadoCancelacion"], &estadoCancelacion)
	err = formatdata.FillStruct(request["CancelacionInversion"], &cancelacionInversion)
	err = formatdata.FillStruct(request["inversionPadre"],&inversionPadre)
	if err == nil {
		_, err = o.QueryTable("cancelacion_inversion_estado_cancelacion").
			Filter("cancelacion_inversion", cancelacionInversion.Id).
			Filter("activo", true).
			Update(orm.Params{
				"activo": "false",
			})
			if err != nil {
				beego.Error(err.Error())
				o.Rollback()
				return
			}
			estadoCancelacion.CancelacionInversion = &cancelacionInversion
			estadoCancelacion.Activo = true
			_, err = o.Insert(&estadoCancelacion)
			if err != nil {
				beego.Error(err.Error())
				o.Rollback()
				return
			}

			if estadoCancelacion.EstadoCancelacionInversion.Id == 3 {
				beego.Info("Aprobacion contable")
				_, err = o.QueryTable("movimiento_contable").Filter("tipo_documento_afectante", 7).Filter("codigo_documento_afectante", cancelacionInversion.Id).Update(orm.Params{
					"estado_movimiento_contable": 2,
				})
				err = o.QueryTable("estado_inversion").
					Filter("numeroOrden", 7).
					One(&estadoInversion)
				if err==nil{
					requestEstadoInv["Estado"]=estadoInversion
					requestEstadoInv["Inversion"]= inversionPadre
					requestEstadoInv["Usuario"]= estadoCancelacion.Usuario
					_,err = AddEstadoInv(requestEstadoInv)
					if err != nil {
							o.Commit()
					}else{
						beego.Error(err.Error())
						o.Rollback()
						return
					}

				} else if err == orm.ErrMultiRows {
					beego.Error("Returned Multi Rows Not One")
					o.Rollback()
					return
				} else if err != nil {
					beego.Error(err.Error())
					o.Rollback()
					return
				}
			}
		}
		return
}


// GetCancelationQuantity retrieves number of records in active state. Returns error if
// Id doesn't exist
func GetCancelationQuantity() (cnt int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(CancelacionInversionEstadoCancelacion))
	qs = qs.Filter("activo", true)
	cnt, err = qs.Count()
	return
}
