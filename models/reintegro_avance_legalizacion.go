package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/udistrital/utils_oas/formatdata"
)

type ReintegroAvanceLegalizacion struct {
	Id        int              `orm:"column(id);pk"`
	Reintegro *Reintegro       `orm:"column(reintegro);rel(fk)"`
	Avance    *SolicitudAvance `orm:"column(avance);rel(fk)"`
}

func (t *ReintegroAvanceLegalizacion) TableName() string {
	return "reintegro_avance_legalizacion"
}

func init() {
	orm.RegisterModel(new(ReintegroAvanceLegalizacion))
}

// AddReintegroAvanceLegalizacion insert a new ReintegroAvanceLegalizacion into database and returns
// last inserted Id on success.
func AddReintegroAvanceLegalizacion(m *ReintegroAvanceLegalizacion) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetReintegroAvanceLegalizacionById retrieves ReintegroAvanceLegalizacion by Id. Returns error if
// Id doesn't exist
func GetReintegroAvanceLegalizacionById(id int) (v *ReintegroAvanceLegalizacion, err error) {
	o := orm.NewOrm()
	v = &ReintegroAvanceLegalizacion{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllReintegroAvanceLegalizacion retrieves all ReintegroAvanceLegalizacion matches certain condition. Returns empty list if
// no records exist
func GetAllReintegroAvanceLegalizacion(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ReintegroAvanceLegalizacion)).RelatedSel()
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

	var l []ReintegroAvanceLegalizacion
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

// UpdateReintegroAvanceLegalizacion updates ReintegroAvanceLegalizacion by Id and returns error if
// the record to be updated doesn't exist
func UpdateReintegroAvanceLegalizacionById(m *ReintegroAvanceLegalizacion) (err error) {
	o := orm.NewOrm()
	v := ReintegroAvanceLegalizacion{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteReintegroAvanceLegalizacion deletes ReintegroAvanceLegalizacion by Id and returns error if
// the record to be deleted doesn't exist
func DeleteReintegroAvanceLegalizacion(id int) (err error) {
	o := orm.NewOrm()
	v := ReintegroAvanceLegalizacion{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ReintegroAvanceLegalizacion{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

//Add all reinintegro avance relations returns error
//if any insert fails
func AddReintegroAvance(request map[string]interface{}) (successNums int64, err error) {
	var reintegrosAvance []ReintegroAvanceLegalizacion
	err = formatdata.FillStruct(request["reintegroAvance"], &reintegrosAvance)
	o := orm.NewOrm()
	if err == nil {

		o.Begin()

		for _, element := range reintegrosAvance {
			_, err = o.QueryTable("reintegro").Filter("id", element.Reintegro.Id).Update(orm.Params{
				"disponible": false,
			})
			if err != nil {
				beego.Error(err.Error())
				o.Rollback()
				return
			}
			successNums, err = o.InsertMulti(100, reintegrosAvance)
			if err != nil {
				beego.Error(err.Error())
				o.Rollback()
				return
			}
		}

	} else {
		beego.Error(err.Error())
		return
	}
	o.Commit()
	return
}
