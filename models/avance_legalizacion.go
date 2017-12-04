package models

import (
	"github.com/astaxie/beego"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"
	"github.com/astaxie/beego/orm"
)

type AvanceLegalizacion struct {
	Id                     int                     `orm:"column(id);pk;auto"`
	TipoAvanceLegalizacion *TipoAvanceLegalizacion `orm:"column(tipo_avance_legalizacion);rel(fk)"`
	Avance                 *SolicitudAvance        `orm:"column(avance);rel(fk)"`
	Tercero                string                  `orm:"column(tercero)"`
	Valor                  float64                 `orm:"column(valor)"`
	FechaCompra            time.Time               `orm:"column(fecha_compra);type(date);null"`
	FechaCambioDivisa      time.Time               `orm:"column(fecha_cambio_divisa);type(date);null"`
	TrmFechaCompra		   float64				   `orm:"column(trm_fecha_compra)"`
	NumeroFactura          string 				   `orm:"column(numero_factura)"`
	Dias                   int                     `orm:"column(dias);null"`
	AvanceLegalizacionCuentaEspecial          []*AvanceLegalizacionCuentaEspecial          `orm:"reverse(many)"`
}

func (t *AvanceLegalizacion) TableName() string {
	return "avance_legalizacion"
}

func init() {
	orm.RegisterModel(new(AvanceLegalizacion))
}

func AddAvanceLegalizacionCompra(legalizacionAvance AvanceLegalizacion) (alert Alert, err error) {
		o := orm.NewOrm()
		o.Begin()
		beego.Info("0",legalizacionAvance)
		_, err = o.Insert(&legalizacionAvance)
		if err == nil{
			for _,AvanceLegalizacionCuenta := range legalizacionAvance.AvanceLegalizacionCuentaEspecial{
				AvanceLegalizacionCuenta.AvanceLegalizacion = &legalizacionAvance
				beego.Info(AvanceLegalizacionCuenta)
				_, err = o.Insert(AvanceLegalizacionCuenta)
				if err != nil{
					beego.Info("1", err.Error())
					o.Rollback()
					return
				}
			}
		}else{
			beego.Info("2", err.Error())
			o.Rollback()
			return
		}
		o.Commit()
		alert = Alert{Type: "success", Code: "S_991", Body: legalizacionAvance}
	return
}
	
// AddAvanceLegalizacion insert a new AvanceLegalizacion into database and returns
// last inserted Id on success.
func AddAvanceLegalizacion(m *AvanceLegalizacion) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetAvanceLegalizacionById retrieves AvanceLegalizacion by Id. Returns error if
// Id doesn't exist
func GetAvanceLegalizacionById(id int) (v *AvanceLegalizacion, err error) {
	o := orm.NewOrm()
	v = &AvanceLegalizacion{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}


// GetAllAvanceLegalizacion retrieves all AvanceLegalizacion matches certain condition. Returns empty list if
// no records exist
func GetAllAvanceLegalizacion(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(AvanceLegalizacion)).RelatedSel()
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

	var l []AvanceLegalizacion
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				o.LoadRelated(&v, "AvanceLegalizacionCuentaEspecial", 5)
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

// UpdateAvanceLegalizacion updates AvanceLegalizacion by Id and returns error if
// the record to be updated doesn't exist
func UpdateAvanceLegalizacionById(m *AvanceLegalizacion) (err error) {
	o := orm.NewOrm()
	v := AvanceLegalizacion{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteAvanceLegalizacion deletes AvanceLegalizacion by Id and returns error if
// the record to be deleted doesn't exist
func DeleteAvanceLegalizacion(id int) (err error) {
	o := orm.NewOrm()
	v := AvanceLegalizacion{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&AvanceLegalizacion{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
