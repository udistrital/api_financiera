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

type CancelacionInversion struct {
	Id               int       `orm:"column(id);pk;auto"`
	FechaCancelacion time.Time `orm:"column(fecha_cancelacion);type(datetime)"`
	Observaciones    string    `orm:"column(observaciones)"`
	UnidadEjecutora  int       `orm:"column(unidad_ejecutora)"`
	Vigencia         int       `orm:"column(vigencia)"`
	UsuarioEjecucion int       `orm:"column(usuario_ejecucion)"`
}

func (t *CancelacionInversion) TableName() string {
	return "cancelacion_inversion"
}

func init() {
	orm.RegisterModel(new(CancelacionInversion))
}

// AddCancelacionInversion insert a new CancelacionInversion into database and returns
// last inserted Id on success.
func AddCancelacionInversion(m *CancelacionInversion) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetCancelacionInversionById retrieves CancelacionInversion by Id. Returns error if
// Id doesn't exist
func GetCancelacionInversionById(id int) (v *CancelacionInversion, err error) {
	o := orm.NewOrm()
	v = &CancelacionInversion{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllCancelacionInversion retrieves all CancelacionInversion matches certain condition. Returns empty list if
// no records exist
func GetAllCancelacionInversion(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(CancelacionInversion))
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

	var l []CancelacionInversion
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

// UpdateCancelacionInversion updates CancelacionInversion by Id and returns error if
// the record to be updated doesn't exist
func UpdateCancelacionInversionById(m *CancelacionInversion) (err error) {
	o := orm.NewOrm()
	v := CancelacionInversion{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteCancelacionInversion deletes CancelacionInversion by Id and returns error if
// the record to be deleted doesn't exist
func DeleteCancelacionInversion(id int) (err error) {
	o := orm.NewOrm()
	v := CancelacionInversion{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&CancelacionInversion{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

//Create a entire record for a cancelation
func CreateCancelacion(v map[string]interface{}) (cancelacionInversion CancelacionInversion, err error) {

	var estadoResp EstadoCancelacionInversion
	var tipoDocAfectante TipoDocumentoAfectante
	var movimientosContables []MovimientoContable
	var cancelacionConcepto CancelacionInversionConcepto
	var idCancInv int64

	o := orm.NewOrm()

	err = formatdata.FillStruct(v["cancelacionInversion"], &cancelacionInversion)
	err = formatdata.FillStruct(v["Movimientos"], &movimientosContables)
	err = formatdata.FillStruct(v["cancelacionConcepto"], &cancelacionConcepto)
	beego.Error("cancelacion concepto", cancelacionConcepto)
	beego.Error("cancelacion inversion", cancelacionInversion)
	beego.Error("movimientos ", movimientosContables)
	if err != nil {
		beego.Error(err.Error())
		return
	}

	o.Begin()

	idCancInv, err = o.Insert(&cancelacionInversion)

	if err == nil {
		cancelacionInversion.Id = int(idCancInv)
		cancelacionConcepto.Cancelacion = &cancelacionInversion

		_, err = o.Insert(&cancelacionConcepto)

		if err == nil {
			err = o.QueryTable("estado_cancelacion_inversion").
				Filter("numeroOrden", 1).
				One(&estadoResp)
			if err == nil {
				estadoCancInv := &CancelacionInversionEstadoCancelacion{CancelacionInversion: &cancelacionInversion, EstadoCancelacionInversion: &estadoResp, Activo: true, Usuario: cancelacionInversion.UsuarioEjecucion}
				_, err = o.Insert(estadoCancInv)
				if err == nil {
					err = o.QueryTable("tipo_documento_afectante").
						Filter("numeroOrden", 7).
						One(&tipoDocAfectante)
					if err == nil {
						for i, _ := range movimientosContables {
							movimientosContables[i].Fecha = time.Now()
							movimientosContables[i].CodigoDocumentoAfectante = cancelacionInversion.Id
							movimientosContables[i].TipoDocumentoAfectante = &tipoDocAfectante
							movimientosContables[i].EstadoMovimientoContable = &EstadoMovimientoContable{Id: 1}
						}

						_, err = AddMovimientoContableArray(&movimientosContables)
						if err == nil {
							o.Commit()
						} else {
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
				} else {
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
	} else {
		beego.Error(err.Error())
		o.Rollback()
		return
	}
	return
}
