package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/udistrital/utils_oas/formatdata"
)

type Cheque struct {
	Id               int        `orm:"column(id);pk;auto"`
	Consecutivo      int        `orm:"column(consecutivo)"`
	OrdenPago        *OrdenPago `orm:"column(orden_pago);rel(fk)"`
	Chequera         *Chequera  `orm:"column(chequera);rel(fk)"`
	Observaciones    string     `orm:"column(observaciones)"`
	Beneficiario     int        `orm:"column(beneficiario)"`
	FechaVencimiento time.Time  `orm:"column(fecha_vencimiento);type(date)"`
	Valor            float64    `orm:"column(valor)"`
}

func (t *Cheque) TableName() string {
	return "cheque"
}

func init() {
	orm.RegisterModel(new(Cheque))
}

// AddCheque insert a new Cheque into database and returns
// last inserted Id on success.
func AddCheque(m *Cheque) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetChequeById retrieves Cheque by Id. Returns error if
// Id doesn't exist
func GetChequeById(id int) (v *Cheque, err error) {
	o := orm.NewOrm()
	v = &Cheque{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllCheque retrieves all Cheque matches certain condition. Returns empty list if
// no records exist
func GetAllCheque(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Cheque)).RelatedSel()
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

	var l []Cheque
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

// UpdateCheque updates Cheque by Id and returns error if
// the record to be updated doesn't exist
func UpdateChequeById(m *Cheque) (err error) {
	o := orm.NewOrm()
	v := Cheque{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteCheque deletes Cheque by Id and returns error if
// the record to be deleted doesn't exist
func DeleteCheque(id int) (err error) {
	o := orm.NewOrm()
	v := Cheque{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Cheque{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

//Gets the value of sum for all cheque related to a pay order
//return zero value if not exists
func GetChequeSumaOP(idOP int) (value int64, err error) {
	o := orm.NewOrm()

	qb, _ := orm.NewQueryBuilder("mysql")

	qb.Select("COALESCE(sum(ch.valor),0)").
		From("cheque ch").
		InnerJoin("cheque_estado_cheque cec").On("ch.id = cec.cheque and cec.activo = true").
		InnerJoin("estado_cheque ec").On("ec.id = cec.estado and ec.numero_orden <= 5").
		Where("ch.orden_pago = ?")

	sql := qb.String()

	err = o.Raw(sql, idOP).QueryRow(&value)

	return

}

//counts  cheque number and returns one more; given a checker
//return one if not exists
func GetNextChequeNumber(idChequera int) (value int64, err error) {
	o := orm.NewOrm()

	qs := o.QueryTable(new(Cheque))
	qs = qs.Filter("chequera", idChequera)
	value, err = qs.Count()

	return

}

// AddCheque insert a new Cheque into database and returns
// last inserted Id on success.
func AddChequeEstado(m map[string]interface{}) (id int64, err error) {
	var cheque Cheque
	var estadoCheque EstadoCheque
	var usuario float64
	var idCheque int64
	o := orm.NewOrm()
	beego.Error("arrives get cheque estado")
	err = formatdata.FillStruct(m["Cheque"], &cheque)
	err = formatdata.FillStruct(m["Usuario"], &usuario)

	if err != nil {
		beego.Error(err.Error())
		return
	}
	o.Begin()

	idCheque, err = o.Insert(&cheque)
	if err == nil {
		cheque.Id = int(idCheque)
		m["Cheque"] = cheque
		err = o.QueryTable(new(EstadoCheque)).
			Filter("numeroOrden", 1).
			One(&estadoCheque)
		if err == nil {
			chequeEstadoCheque := &ChequeEstadoCheque{Cheque: &cheque, Activo: true, Estado: &estadoCheque, Usuario: int(usuario)}
			_, err = o.Insert(chequeEstadoCheque)
			if err != nil {
				beego.Error(err.Error())
				o.Rollback()
				return
			}
			_, err = o.QueryTable("chequera").
				Filter("Id", cheque.Chequera.Id).
				Update(orm.Params{
					"cheques_disponibles": orm.ColValue(orm.ColMinus, 1),
				})
			if err != nil {
				beego.Error(err.Error())
				o.Rollback()
				return
			}
		} else {
			beego.Error(err.Error())
			o.Rollback()
			return
		}
	} else {
		beego.Error(err.Error())
		o.Rollback()
		return
	}
	o.Commit()
	return
}

// GetRecordsCheque retrieves quantity of records in Cheque s table
// returns zero value Id doesn't exist
func GetRecordsCheque(query map[string]string) (cnt int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Cheque))

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
	cnt, err = qs.Count()
	return
}
