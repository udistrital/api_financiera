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

type OrdenPagoAvanceLegalizacion struct {
	Id        int              `orm:"column(id);pk;auto"`
	Avance    *SolicitudAvance `orm:"column(avance);rel(fk)"`
	OrdenPago *OrdenPago       `orm:"column(orden_pago);rel(fk)"`
}

func (t *OrdenPagoAvanceLegalizacion) TableName() string {
	return "orden_pago_avance_legalizacion"
}

func init() {
	orm.RegisterModel(new(OrdenPagoAvanceLegalizacion))
}

// AddOrdenPagoAvanceLegalizacion insert a new OrdenPagoAvanceLegalizacion into database and returns
// last inserted Id on success.
func AddOrdenPagoAvanceLegalizacion(m *OrdenPagoAvanceLegalizacion) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetOrdenPagoAvanceLegalizacionById retrieves OrdenPagoAvanceLegalizacion by Id. Returns error if
// Id doesn't exist
func GetOrdenPagoAvanceLegalizacionById(id int) (v *OrdenPagoAvanceLegalizacion, err error) {
	o := orm.NewOrm()
	v = &OrdenPagoAvanceLegalizacion{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllOrdenPagoAvanceLegalizacion retrieves all OrdenPagoAvanceLegalizacion matches certain condition. Returns empty list if
// no records exist
func GetAllOrdenPagoAvanceLegalizacion(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(OrdenPagoAvanceLegalizacion)).RelatedSel(4)
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

	var l []OrdenPagoAvanceLegalizacion
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

// UpdateOrdenPagoAvanceLegalizacion updates OrdenPagoAvanceLegalizacion by Id and returns error if
// the record to be updated doesn't exist
func UpdateOrdenPagoAvanceLegalizacionById(m *OrdenPagoAvanceLegalizacion) (err error) {
	o := orm.NewOrm()
	v := OrdenPagoAvanceLegalizacion{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteOrdenPagoAvanceLegalizacion deletes OrdenPagoAvanceLegalizacion by Id and returns error if
// the record to be deleted doesn't exist
func DeleteOrdenPagoAvanceLegalizacion(id int) (err error) {
	o := orm.NewOrm()
	v := OrdenPagoAvanceLegalizacion{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&OrdenPagoAvanceLegalizacion{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

//Add all OP avance relations returns error
//if any insert fails
func AddOPAvance(request map[string]interface{}) (successNums int64, err error) {
	var OPAvance []OrdenPagoAvanceLegalizacion

	var cnt int64
	var tamArray int
	err = formatdata.FillStruct(request["OPAvance"], &OPAvance)
	o := orm.NewOrm()
	if err == nil {

		o.Begin()
		tamArray = len(OPAvance)
		for i := 0; i < tamArray; i += 1 {

			qs := o.QueryTable(new(OrdenPagoAvanceLegalizacion))
			qs = qs.Filter("avance", OPAvance[i].Avance.Id)
			qs = qs.Filter("orden_pago", OPAvance[i].OrdenPago.Id)
			cnt, err = qs.Count()

			if err != nil {
				beego.Error(err.Error())
				return
			} else {
				if cnt > 0 {
					OPAvance = append(OPAvance[:i], OPAvance[i+1:]...)
					tamArray -= 1
				}
			}
		}
		successNums, err = o.InsertMulti(100, OPAvance)
		if err != nil {
			beego.Error(err.Error())
			o.Rollback()
			return
		}
	} else {
		beego.Error(err.Error())
		return
	}
	o.Commit()
	return
}
