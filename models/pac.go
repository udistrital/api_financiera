package models

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/udistrital/utils_oas/formatdata"
)

type Pac struct {
	Id          int           `orm:"column(id);pk;auto"`
	Descripcion string        `orm:"column(descripcion)"`
	Vigencia    int           `orm:"column(vigencia)"`
	DetallePac  []*DetallePac `orm:"reverse(many)"`
}

func (t *Pac) TableName() string {
	return "pac"
}

func init() {
	orm.RegisterModel(new(Pac))
}

// AddPac insert a new Pac into database and returns
// last inserted Id on success.
func AddPac(m *Pac) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetPacById retrieves Pac by Id. Returns error if
// Id doesn't exist
func GetPacById(id int) (v *Pac, err error) {
	o := orm.NewOrm()
	v = &Pac{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllPac retrieves all Pac matches certain condition. Returns empty list if
// no records exist
func GetAllPac(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Pac))
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

	var l []Pac
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

// UpdatePac updates Pac by Id and returns error if
// the record to be updated doesn't exist
func UpdatePacById(m *Pac) (err error) {
	o := orm.NewOrm()
	v := Pac{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeletePac deletes Pac by Id and returns error if
// the record to be deleted doesn't exist
func DeletePac(id int) (err error) {
	o := orm.NewOrm()
	v := Pac{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Pac{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func GetPacByVigencia(vigencia float64) (v Pac, err error) {
	o := orm.NewOrm()

	qs := o.QueryTable("pac").
		Filter("vigencia", vigencia)

	err = qs.One(&v)
	if err == orm.ErrMultiRows {
		fmt.Println("Returned Multi Rows Not One")
	}
	if err == orm.ErrNoRows {
		fmt.Println("Not row found")
		v.Descripcion = " PAC Vigencia " + strconv.FormatFloat(vigencia, 'f', 0, 64)
		v.Vigencia = int(vigencia)
		//insert pac
		_, err = o.Insert(&v)
		if err != nil {
			beego.Info(err.Error())
			return
		}
	}
	return
}

func GetPacProjection(vigencia int, mes int, fuente string, rubro string, nperiodos int) (res []interface{}, err error) {
	o := orm.NewOrm()
	var m []orm.Params
	vigenciaInicial := vigencia - nperiodos
	_, err = o.Raw(`Select COALESCE(valor_proyectado_mes,0) as pry,
       					ROW_NUMBER () OVER (ORDER BY  pac.vigencia desc) as nfila
					from financiera.pac pac
					left join financiera.detalle_pac detalle
    					on detalle.pac = pac.id
							and detalle.mes = ?
							and detalle.fuente_financiamiento = ?
							and detalle.rubro = ?
					where pac.vigencia >= ? and pac.vigencia < ?`,
		mes, fuente, rubro, vigenciaInicial, vigencia).Values(&m)
	err = formatdata.FillStruct(m, &res)
	return
}
