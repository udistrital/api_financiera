package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"
	"strconv"
	"github.com/astaxie/beego/orm"
)

type Apropiacion struct {
	Id              int                `orm:"column(id);pk"`
	Vigencia        float64            `orm:"column(vigencia);null"`
	Rubro           *Rubro             `orm:"column(rubro);rel(fk)"`
	UnidadEjecutora int                `orm:"column(unidad_ejecutora);null"`
	ValorRezago     float64            `orm:"column(valor_rezago);null"`
	Valor           float64            `orm:"column(valor);null"`
	TipoDocumento   string             `orm:"column(tipo_documento);null"`
	DocumentoNumero string             `orm:"column(documento_numero);null"`
	DocumentoFecha  time.Time          `orm:"column(documento_fecha);type(date);null"`
	Estado          *EstadoApropiacion `orm:"column(estado);rel(fk)"`
}

func (t *Apropiacion) TableName() string {
	return "apropiacion"
}

func init() {
	orm.RegisterModel(new(Apropiacion))
}

// AddApropiacion insert a new Apropiacion into database and returns
// last inserted Id on success.
func AddApropiacion(m *Apropiacion) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetApropiacionById retrieves Apropiacion by Id. Returns error if
// Id doesn't exist
func GetApropiacionById(id int) (v *Apropiacion, err error) {
	o := orm.NewOrm()
	v = &Apropiacion{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllApropiacion retrieves all Apropiacion matches certain condition. Returns empty list if
// no records exist
func GetAllApropiacion(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Apropiacion))
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

	var l []Apropiacion
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

// UpdateApropiacion updates Apropiacion by Id and returns error if
// the record to be updated doesn't exist
func UpdateApropiacionById(m *Apropiacion) (err error) {
	o := orm.NewOrm()
	v := Apropiacion{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteApropiacion deletes Apropiacion by Id and returns error if
// the record to be deleted doesn't exist
func DeleteApropiacion(id int) (err error) {
	o := orm.NewOrm()
	v := Apropiacion{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Apropiacion{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

//funcion para comprobar saldo de la apropiacion de un Rubro

func SaldoApropiacion(Id int) (valor float64) {
	o := orm.NewOrm()
	var maps_valor_tot []orm.Params
	var maps_valor_compr []orm.Params
	var maps_valor_anulaciones []orm.Params
	var valor_tot float64
	valor = 0
	o.Raw("SELECT valor as valor from financiera.apropiacion WHERE id = ? and estado = ? ", Id, 2).Values(&maps_valor_tot)
	fmt.Println("maps: ", len(maps_valor_tot))
	if len(maps_valor_tot) > 0 {
		valor_tot, _ = strconv.ParseFloat(maps_valor_tot[0]["valor"].(string), 64)
	} else {
		valor_tot = 0
	}
	o.Raw("SELECT COALESCE( SUM(valor), 0 ) as valor from financiera.disponibilidad_apropiacion WHERE apropiacion = ?", Id).Values(&maps_valor_compr)
	if len(maps_valor_compr) > 0 {
		if valor_compr, err := strconv.ParseFloat(maps_valor_compr[0]["valor"].(string), 64); err == nil {
			valor = valor_tot - valor_compr
		} else {
			valor = valor_tot - 0
		}

	}
	o.Raw(`SELECT  COALESCE(SUM(b.valor),0) as valor

				FROM financiera.anulacion_disponibilidad_apropiacion AS b
				LEFT JOIN
				     financiera.disponibilidad_apropiacion AS a
				ON a.id = b.disponibilidad_apropiacion

				WHERE  a.apropiacion = ?`, Id).Values(&maps_valor_anulaciones)
	if maps_valor_anulaciones[0]["valor"] != nil {
		if valor_anulaciones, err := strconv.ParseFloat(maps_valor_anulaciones[0]["valor"].(string), 64); err == nil {
			valor = valor + valor_anulaciones
		}

	}
	return
}

//----------------------------------------------------------
