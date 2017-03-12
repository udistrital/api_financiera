package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type RegistroPresupuestal struct {
	Id                         int                         `orm:"column(id);pk"`
	UnidadEjecutora            int16                       `orm:"column(unidad_ejecutora)"`
	Vigencia                   float64                     `orm:"column(vigencia)"`
	FechaMovimiento            time.Time                   `orm:"column(fecha_movimiento);type(date);null"`
	Responsable                int                         `orm:"column(responsable);null"`
	Estado                     *EstadoRegistroPresupuestal `orm:"column(estado);rel(fk)"`
	NumeroRegistroPresupuestal int                         `orm:"column(numero_registro_presupuestal)"`
	Beneficiario               int                         `orm:"column(beneficiario);null"`
	Compromiso                 *Compromiso                 `orm:"column(compromiso);rel(fk)"`
}
type DatosRubroRegistroPresupuestal struct {
	Id             int
	Disponibilidad *Disponibilidad
	Apropiacion    *Apropiacion
	Valor          float64
	ValorAsignado  float64
}

type DatosRegistroPresupuestal struct { //estructura temporal para el registro con relacion a las apropiaciones
	Rp     *RegistroPresupuestal
	Rubros []DatosRubroRegistroPresupuestal
}

func (t *RegistroPresupuestal) TableName() string {
	return "registro_presupuestal"
}

func init() {
	orm.RegisterModel(new(RegistroPresupuestal))
}

// AddRegistroPresupuestal insert a new RegistroPresupuestal into database and returns
// last inserted Id on success.
func AddRegistoPresupuestal(m *DatosRegistroPresupuestal) (id int64, err error) {
	o := orm.NewOrm()
	o.Begin()
	id, err = o.Insert(m.Rp)
	if err == nil {
		m.Rp.Id = int(id)
		for _, data := range m.Rubros {
			registro := RegistroPresupuestalDisponibilidadApropiacion{
				RegistroPresupuestal:      m.Rp,
				DisponibilidadApropiacion: &DisponibilidadApropiacion{Id: data.Id},
				Valor: data.ValorAsignado,
			}
			_, err2 := o.Insert(&registro)
			if err2 != nil {
				o.Rollback()
			}
		}
	}

	o.Commit()
	return
}

// GetRegistroPresupuestalById retrieves RegistroPresupuestal by Id. Returns error if
// Id doesn't exist
func GetRegistroPresupuestalById(id int) (v *RegistroPresupuestal, err error) {
	o := orm.NewOrm()
	v = &RegistroPresupuestal{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllRegistroPresupuestal retrieves all RegistroPresupuestal matches certain condition. Returns empty list if
// no records exist
func GetAllRegistroPresupuestal(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(RegistroPresupuestal))
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

	var l []RegistroPresupuestal
	qs = qs.OrderBy(sortFields...).RelatedSel(5)
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

// UpdateRegistroPresupuestal updates RegistroPresupuestal by Id and returns error if
// the record to be updated doesn't exist
func UpdateRegistroPresupuestalById(m *RegistroPresupuestal) (err error) {
	o := orm.NewOrm()
	v := RegistroPresupuestal{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteRegistroPresupuestal deletes RegistroPresupuestal by Id and returns error if
// the record to be deleted doesn't exist
func DeleteRegistroPresupuestal(id int) (err error) {
	o := orm.NewOrm()
	v := RegistroPresupuestal{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&RegistroPresupuestal{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
