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

type InversionesActaInversion struct {
	Id            int            `orm:"column(id);pk;auto"`
	Inversion     *Inversion     `orm:"column(inversion);rel(fk)"`
	ActaInversion *ActaInversion `orm:"column(acta_inversion);rel(fk)"`
	FechaRegistro time.Time      `orm:"column(fecha_registro);auto_now_add;type(datetime)"`
	Usuario       string         `orm:"column(usuario);null"`
	ActaPadre     *Inversion     `orm:"column(acta_padre);rel(fk);null"`
}

func (t *InversionesActaInversion) TableName() string {
	return "inversiones_acta_inversion"
}

func init() {
	orm.RegisterModel(new(InversionesActaInversion))
}

// AddInversionesActaInversion insert a new InversionesActaInversion into database and returns
// last inserted Id on success.
func AddInversionesActaInversion(m *InversionesActaInversion) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetInversionesActaInversionById retrieves InversionesActaInversion by Id. Returns error if
// Id doesn't exist
func GetInversionesActaInversionById(id int) (v *InversionesActaInversion, err error) {
	o := orm.NewOrm()
	v = &InversionesActaInversion{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllInversionesActaInversion retrieves all InversionesActaInversion matches certain condition. Returns empty list if
// no records exist
func GetAllInversionesActaInversion(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(InversionesActaInversion))
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

	var l []InversionesActaInversion
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

// UpdateInversionesActaInversion updates InversionesActaInversion by Id and returns error if
// the record to be updated doesn't exist
func UpdateInversionesActaInversionById(m *InversionesActaInversion) (err error) {
	o := orm.NewOrm()
	v := InversionesActaInversion{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteInversionesActaInversion deletes InversionesActaInversion by Id and returns error if
// the record to be deleted doesn't exist
func DeleteInversionesActaInversion(id int) (err error) {
	o := orm.NewOrm()
	v := InversionesActaInversion{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&InversionesActaInversion{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

// Insert an entire inversion in database returns error if record cant be inserted
func AddInver(request map[string]interface{}) (inversion Inversion, err error) {
	err = formatdata.FillStruct(request["Inversion"], &inversion)
	var idInversion int
	var valorHijos float64
	var tipoInversion int
	var usuario string
	var actapadre Inversion
	var inversionCompare Inversion
	var invActInv InversionesActaInversion
	var invEstadoInv InversionEstadoInversion
	var concepto Concepto
	var mov []MovimientoContable
	var totalInv float64
	o := orm.NewOrm()
	if err == nil {
		if qb, errq := orm.NewQueryBuilder("tidb"); errq == nil {
			qb.Select("COALESCE(MAX(id), 0)+1  as consecutivo").
				From("Inversion")
			sql := qb.String()
			o.Raw(sql).QueryRow(&idInversion)

			inversion.Id = idInversion

			o.Begin()

			_, err = o.Insert(&inversion)

			if err == nil {
				err = formatdata.FillStruct(request["tipoInversion"], &tipoInversion)
				err = formatdata.FillStruct(request["usuario"], &usuario)
				err = formatdata.FillStruct(request["actapadre"], &actapadre)
				err = formatdata.FillStruct(request["EstadoInversion"], &invEstadoInv)
				err = formatdata.FillStruct(request["Concepto"], &concepto)
				err = formatdata.FillStruct(request["Movimientos"], &mov)
				err = formatdata.FillStruct(request["TotalInversion"], &totalInv)

				if err != nil {
					beego.Info(err.Error())
					o.Rollback()
					return
				}

				inversion_concepto := &InversionConcepto{ValorAgregado: totalInv,
					Inversion: &inversion,
					Concepto:  &concepto}

				_, err = o.Insert(inversion_concepto)

				if err != nil {
					beego.Info(err.Error())
					o.Rollback()
					return
				}

				for _, element := range mov {
					element.Fecha = time.Now()
					element.TipoDocumentoAfectante = &TipoDocumentoAfectante{Id: 3}
					element.CodigoDocumentoAfectante = inversion.Id
					element.EstadoMovimientoContable = &EstadoMovimientoContable{Id: 1}
					_, err = o.Insert(&element)

					if err != nil {
						beego.Info(err.Error())
						o.Rollback()
						return
					}
				}

				if err != nil {
					beego.Info(err.Error())
					o.Rollback()
					return
				}

				actaInversion := ActaInversion{Id: tipoInversion}

				invActInv.Inversion = &inversion
				invActInv.ActaInversion = &actaInversion
				invActInv.Usuario = usuario

				if !reflect.DeepEqual(actapadre,inversionCompare ) {
					beego.Error("informacion acta padre")
					invActInv.ActaPadre = &actapadre

					qb.Select("coalesce(sum(ic.valor_agregado),0)").
								From("inversion_concepto ic").
								InnerJoin("inversiones_acta_inversion ac").On("ac.inversion = ic.inversion").
								Where("ac.acta_padre > ?")

					sql := qb.String()

				  o.Raw(sql,actapadre.Id).QueryRow(&valorHijos)

					if actapadre.ValorNetoGirar <= valorHijos + (totalInv) {

					}else{
						return
					}
				}

				_, err = o.Insert(&invActInv)

				if err != nil {
					beego.Error(err.Error())
					o.Rollback()
					return
				}
				invEstadoInv.Inversion = &inversion
				_, err = o.Insert(&invEstadoInv)

				if err != nil {
					beego.Error(err.Error())
					o.Rollback()
					return
				}
				o.Commit()
				return
			} else {
				o.Rollback()
				return
			}
		} else {
			beego.Info(errq.Error())
			return inversion, errq
		}
	}

	return
}
