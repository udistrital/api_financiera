package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/udistrital/api_financiera/utilidades"
)

type Ingreso struct {
	Id                int              `orm:"column(id);pk;auto"`
	Consecutivo       float64          `orm:"column(consecutivo)"`
	Vigencia          float64          `orm:"column(vigencia)"`
	FechaIngreso      time.Time        `orm:"column(fecha_ingreso);type(date)"`
	FechaConsignacion time.Time        `orm:"column(fecha_consignacion);type(date)"`
	Valor             float64          `orm:"column(valor)"`
	Observaciones     string           `orm:"column(observaciones);null"`
	OrigenIngreso     string           `orm:"column(origen_ingreso);null"`
	FormaIngreso      *FormaIngreso    `orm:"column(forma_ingreso);rel(fk)"`
	EstadoIngreso     *EstadoIngreso   `orm:"column(estado_ingreso);rel(fk)"`
	UnidadEjecutora   *UnidadEjecutora `orm:"column(unidad_ejecutora);rel(fk)"`
	Aportante         int              `orm:"column(aportante);null"`
	Reviso            int              `orm:"column(reviso);null"`
	Elaboro           int              `orm:"column(elaboro)"`
}

func (t *Ingreso) TableName() string {
	return "ingreso"
}

func init() {
	orm.RegisterModel(new(Ingreso))
}

// AddIngreso insert a new Ingreso into database and returns
// last inserted Id on success.
func AddIngresotr(m map[string]interface{}) (ingreso Ingreso, err error) {
	var id int64
	err = utilidades.FillStruct(m["Ingreso"], &ingreso)
	if err == nil {
		ingreso.EstadoIngreso = &EstadoIngreso{Id: 1}
		ingreso.FechaIngreso = time.Now()
		ingreso.Vigencia = float64(time.Now().Year())
		o := orm.NewOrm()
		o.Begin()
		id, err = o.Insert(&ingreso)
		if err != nil {
			o.Rollback()
			return
		} else {
			ingreso.Id = int(id)
			var ingresoslice []interface{}
			err = utilidades.FillStruct(m["IngresoBanco"], &ingresoslice)
			if err == nil {
				concepto := &Concepto{}
				err = utilidades.FillStruct(m["Concepto"], concepto)
				if err == nil {
					for _, ingresobanco := range ingresoslice {
						fmt.Println(ingresobanco)
						ingreso_concepto := &IngresoConcepto{ValorAgregado: 0,
							Ingreso:  &ingreso,
							Concepto: concepto}
						_, err = o.Insert(ingreso_concepto)
						if err != nil {
							o.Rollback()
							return
						}
					}
				} else {
					o.Rollback()
					return
				}

			} else {
				o.Rollback()
				return
			}

			o.Commit()
			return
		}
	} else {
		return
	}

}

// AddIngreso insert a new Ingreso into database and returns
// last inserted Id on success.
func AddIngreso(m *Ingreso) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetIngresoById retrieves Ingreso by Id. Returns error if
// Id doesn't exist
func GetIngresoById(id int) (v *Ingreso, err error) {
	o := orm.NewOrm()
	v = &Ingreso{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllIngreso retrieves all Ingreso matches certain condition. Returns empty list if
// no records exist
func GetAllIngreso(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Ingreso)).RelatedSel()
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

	var l []Ingreso
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

// UpdateIngreso updates Ingreso by Id and returns error if
// the record to be updated doesn't exist
func UpdateIngresoById(m *Ingreso) (err error) {
	o := orm.NewOrm()
	v := Ingreso{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteIngreso deletes Ingreso by Id and returns error if
// the record to be deleted doesn't exist
func DeleteIngreso(id int) (err error) {
	o := orm.NewOrm()
	v := Ingreso{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Ingreso{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
