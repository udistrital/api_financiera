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

type ConceptoAvanceLegalizacion struct {
	Id       int              `orm:"column(id);pk;auto"`
	Valor    int64            `orm:"column(valor)"`
	Avance   *SolicitudAvance `orm:"column(avance);rel(fk)"`
	Concepto *Concepto        `orm:"column(concepto);rel(fk)"`
}

func (t *ConceptoAvanceLegalizacion) TableName() string {
	return "concepto_avance_legalizacion"
}

func init() {
	orm.RegisterModel(new(ConceptoAvanceLegalizacion))
}

// AddConceptoAvanceLegalizacion insert a new ConceptoAvanceLegalizacion into database and returns
// last inserted Id on success.
func AddConceptoAvanceLegalizacion(m *ConceptoAvanceLegalizacion) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetConceptoAvanceLegalizacionById retrieves ConceptoAvanceLegalizacion by Id. Returns error if
// Id doesn't exist
func GetConceptoAvanceLegalizacionById(id int) (v *ConceptoAvanceLegalizacion, err error) {
	o := orm.NewOrm()
	v = &ConceptoAvanceLegalizacion{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllConceptoAvanceLegalizacion retrieves all ConceptoAvanceLegalizacion matches certain condition. Returns empty list if
// no records exist
func GetAllConceptoAvanceLegalizacion(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ConceptoAvanceLegalizacion)).RelatedSel()
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

	var l []ConceptoAvanceLegalizacion
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

// UpdateConceptoAvanceLegalizacion updates ConceptoAvanceLegalizacion by Id and returns error if
// the record to be updated doesn't exist
func UpdateConceptoAvanceLegalizacionById(m *ConceptoAvanceLegalizacion) (err error) {
	o := orm.NewOrm()
	v := ConceptoAvanceLegalizacion{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteConceptoAvanceLegalizacion deletes ConceptoAvanceLegalizacion by Id and returns error if
// the record to be deleted doesn't exist
func DeleteConceptoAvanceLegalizacion(id int) (err error) {
	o := orm.NewOrm()
	v := ConceptoAvanceLegalizacion{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ConceptoAvanceLegalizacion{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

// add a Accountant information about an advance, state,a id type returns error
//if any insert fails
func CreateLegalizacionAccountantInfo(request map[string]interface{}) (conceptoAvance ConceptoAvanceLegalizacion, err error) {

	var mov []MovimientoContable
	var tipoDocAfectante TipoDocumentoAfectante
	var cnt int64
	var idConceptoAvnc int64
	o := orm.NewOrm()

	err = formatdata.FillStruct(request["ConceptoAvance"], &conceptoAvance)
	err = formatdata.FillStruct(request["Movimientos"], &mov)

	if err == nil {
		o.Begin()

		qs := o.QueryTable(new(ConceptoAvanceLegalizacion))
		qs = qs.Filter("avance", conceptoAvance.Avance.Id)
		cnt, err = qs.Count()

		if err != nil {
			beego.Error(err.Error())
			return
		}
		if cnt > 0 {
			_, err = qs.Update(orm.Params{
				"concepto": conceptoAvance.Concepto.Id,
			})
			if err != nil {
				beego.Info(err.Error())
				o.Rollback()
				return
			}
		} else {
			idConceptoAvnc, err = o.Insert(&conceptoAvance)
			if err != nil {
				beego.Info(err.Error())
				o.Rollback()
				return
			}
			conceptoAvance.Id = int(idConceptoAvnc)
		}
		err = o.QueryTable("tipo_documento_afectante").
			Filter("numeroOrden", 8).
			One(&tipoDocAfectante)

		if err == nil {

			for _, element := range mov {
				element.Fecha = time.Now()
				element.TipoDocumentoAfectante = &tipoDocAfectante
				element.CodigoDocumentoAfectante = conceptoAvance.Avance.Id
				element.EstadoMovimientoContable = &EstadoMovimientoContable{Id: 1}
				_, err = o.Insert(&element)

				if err != nil {
					beego.Info(err.Error())
					o.Rollback()
					return
				}
			}
		} else {
			beego.Error(err.Error())
			o.Rollback()
			return
		}

	} else {
		return
	}
	o.Commit()
	return
}
