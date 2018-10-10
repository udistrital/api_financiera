package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type ConceptoAvanceLegalizacionTipo struct {
	Id                 int                     `orm:"column(id);pk;auto"`
	AvanceLegalizacion *AvanceLegalizacionTipo `orm:"column(avance_legalizacion);rel(fk)"`
	Concepto           *Concepto               `orm:"column(concepto);rel(fk)"`
	Valor              float64                 `orm:"column(valor)"`
}

func (t *ConceptoAvanceLegalizacionTipo) TableName() string {
	return "concepto_avance_legalizacion_tipo"
}

func init() {
	orm.RegisterModel(new(ConceptoAvanceLegalizacionTipo))
}

// AddConceptoAvanceLegalizacionTipo insert a new ConceptoAvanceLegalizacionTipo into database and returns
// last inserted Id on success.
func AddConceptoAvanceLegalizacionTipo(m *ConceptoAvanceLegalizacionTipo) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetConceptoAvanceLegalizacionTipoById retrieves ConceptoAvanceLegalizacionTipo by Id. Returns error if
// Id doesn't exist
func GetConceptoAvanceLegalizacionTipoById(id int) (v *ConceptoAvanceLegalizacionTipo, err error) {
	o := orm.NewOrm()
	v = &ConceptoAvanceLegalizacionTipo{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllConceptoAvanceLegalizacionTipo retrieves all ConceptoAvanceLegalizacionTipo matches certain condition. Returns empty list if
// no records exist
func GetAllConceptoAvanceLegalizacionTipo(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64, groupby []string) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ConceptoAvanceLegalizacionTipo))
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
	var l []ConceptoAvanceLegalizacionTipo
	qs = qs.OrderBy(sortFields...)
	if len(groupby) != 0 {
		for i := 0; i < len(groupby); i++ {
			groupby[i] = strings.Replace(groupby[i], ".", "__", -1)
		}
		qs = qs.GroupBy(groupby...)
	}
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				o.LoadRelated(&v, "Concepto", 2)
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				beego.Error("Fields", fields, "tam", len(fields))
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

// UpdateConceptoAvanceLegalizacionTipo updates ConceptoAvanceLegalizacionTipo by Id and returns error if
// the record to be updated doesn't exist
func UpdateConceptoAvanceLegalizacionTipoById(m *ConceptoAvanceLegalizacionTipo) (err error) {
	o := orm.NewOrm()
	v := ConceptoAvanceLegalizacionTipo{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteConceptoAvanceLegalizacionTipo deletes ConceptoAvanceLegalizacionTipo by Id and returns error if
// the record to be deleted doesn't exist
func DeleteConceptoAvanceLegalizacionTipo(id int) (err error) {
	o := orm.NewOrm()
	v := ConceptoAvanceLegalizacionTipo{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ConceptoAvanceLegalizacionTipo{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

// GetConceptoAvanceLegalizacionTipoByIdAvanceLegalizacion retrieves grouped concept for AvanceLegalizacion
func GetConceptoAvanceLegalizacionTipoByIdAvance(id int) (ml []interface{}, err error) {
	o := orm.NewOrm()
	var conceptos []Concepto
	qb, _ := orm.NewQueryBuilder("mysql")

	qb.Select("capt.concepto as Id").
		From("financiera.avance_legalizacion al").
		InnerJoin("financiera.avance_legalizacion_tipo alt").On("alt.avance_legalizacion = al.id").
		And("alt.estado = 1").
		InnerJoin("financiera.concepto_avance_legalizacion_tipo capt").On("capt.avance_legalizacion = alt.id").
		Where("al.id = ?").
		GroupBy("capt.concepto")

	sql := qb.String()

	_, err = o.Raw(sql, id).QueryRows(&conceptos)

	if err != nil {
		return
	}
	for _, v := range conceptos {
		if err = o.Read(&v); err == nil {
			ml = append(ml, v)
		} else {
			beego.Error("error ", err.Error())
		}
	}
	return
}
