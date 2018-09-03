package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"github.com/udistrital/utils_oas/formatdata"
)

type ChequeEstadoCheque struct {
	Id            int           `orm:"column(id);pk;auto"`
	Cheque        *Cheque       `orm:"column(cheque);rel(fk)"`
	Activo        bool          `orm:"column(activo)"`
	FechaRegistro time.Time     `orm:"column(fecha_registro);auto_now_add;type(datetime)"`
	Estado        *EstadoCheque `orm:"column(estado);rel(fk)"`
	Usuario       int           `orm:"column(usuario);null"`
}

func (t *ChequeEstadoCheque) TableName() string {
	return "cheque_estado_cheque"
}

func init() {
	orm.RegisterModel(new(ChequeEstadoCheque))
}

// AddChequeEstadoCheque insert a new ChequeEstadoCheque into database and returns
// last inserted Id on success.
func AddChequeEstadoCheque(m *ChequeEstadoCheque) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetChequeEstadoChequeById retrieves ChequeEstadoCheque by Id. Returns error if
// Id doesn't exist
func GetChequeEstadoChequeById(id int) (v *ChequeEstadoCheque, err error) {
	o := orm.NewOrm()
	v = &ChequeEstadoCheque{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllChequeEstadoCheque retrieves all ChequeEstadoCheque matches certain condition. Returns empty list if
// no records exist
func GetAllChequeEstadoCheque(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ChequeEstadoCheque))
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

	var l []ChequeEstadoCheque
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

// UpdateChequeEstadoCheque updates ChequeEstadoCheque by Id and returns error if
// the record to be updated doesn't exist
func UpdateChequeEstadoChequeById(m *ChequeEstadoCheque) (err error) {
	o := orm.NewOrm()
	v := ChequeEstadoCheque{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteChequeEstadoCheque deletes ChequeEstadoCheque by Id and returns error if
// the record to be deleted doesn't exist
func DeleteChequeEstadoCheque(id int) (err error) {
	o := orm.NewOrm()
	v := ChequeEstadoCheque{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ChequeEstadoCheque{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}



//Insert active state for cheque, update another states
// as no active, change disponible cheque number on checker
func AddNewEstadoCheque(request map[string]interface{})(estadoCheque ChequeEstadoCheque, err error){
	var cheque Cheque
	err = formatdata.FillStruct(request["ChequeEstadoCheque"], &estadoCheque)
	err = formatdata.FillStruct(request["Cheque"], &cheque)

	if err != nil {
		beego.Error(err.Error())
		return
	}
	o := orm.NewOrm()
	o.Begin()
	_, err = o.QueryTable("cheque_estado_cheque").
		Filter("cheque", cheque.Id).
		Filter("activo", true).
		Update(orm.Params{
			"activo": "false",
		})
		if err != nil {
			beego.Error(err.Error())
			o.Rollback()
			return
		}
		estadoCheque.Cheque = &cheque

		_, err = o.Insert(&estadoCheque)
		if err != nil {
			beego.Error(err.Error())
			o.Rollback()
			return
		}
		o.Commit()
		return
}
