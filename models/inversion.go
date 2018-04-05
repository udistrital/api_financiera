package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Inversion struct {
	Id                  int       `orm:"column(id);pk"`
	Vendedor            string    `orm:"column(vendedor)"`
	Emisor              string    `orm:"column(emisor)"`
	NumOperacion        int       `orm:"column(num_operacion)"`
	Trm                 float64   `orm:"column(trm)"`
	TasaNominal         float64   `orm:"column(tasa_nominal);null"`
	ValorNomSaldo       float64   `orm:"column(valor_nom_saldo);null"`
	ValorNomSaldoMonNal float64   `orm:"column(valor_nom_saldo_mon_nal);null"`
	ValorActual         float64   `orm:"column(valor_actual);null"`
	ValorNetoGirar      float64   `orm:"column(valor_neto_girar);null"`
	FechaCompra         time.Time `orm:"column(fecha_compra);type(date);null"`
	FechaRedencion      time.Time `orm:"column(fecha_redencion);type(date);null"`
	FechaVencimiento    time.Time `orm:"column(fecha_vencimiento);type(date);null"`
	FechaEmision        time.Time `orm:"column(fecha_emision);type(date);null"`
	Comprador           string    `orm:"column(comprador);null"`
	ValorRecompra       float64   `orm:"column(valor_recompra);null"`
	FechaVenta          time.Time `orm:"column(fecha_venta);type(date);null"`
	FechaPacto          time.Time `orm:"column(fecha_pacto);type(date);null"`
	Observaciones       string    `orm:"column(observaciones);null"`
}

func (t *Inversion) TableName() string {
	return "inversion"
}

func init() {
	orm.RegisterModel(new(Inversion))
}

// AddInversion insert a new Inversion into database and returns
// last inserted Id on success.
func AddInversion(m *Inversion) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetInversionById retrieves Inversion by Id. Returns error if
// Id doesn't exist
func GetInversionById(id int) (v *Inversion, err error) {
	o := orm.NewOrm()
	v = &Inversion{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllInversion retrieves all Inversion matches certain condition. Returns empty list if
// no records exist
func GetAllInversion(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Inversion))
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

	var l []Inversion
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

// UpdateInversion updates Inversion by Id and returns error if
// the record to be updated doesn't exist
func UpdateInversionById(m *Inversion) (err error) {
	o := orm.NewOrm()
	v := Inversion{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteInversion deletes Inversion by Id and returns error if
// the record to be deleted doesn't exist
func DeleteInversion(id int) (err error) {
	o := orm.NewOrm()
	v := Inversion{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Inversion{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
