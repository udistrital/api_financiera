package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"github.com/udistrital/utils_oas/formatdata"
)

type Chequera struct {
	Id                  int             `orm:"column(id);pk;auto"`
	Consecutivo         int             `orm:"column(consecutivo)"`
	UnidadEjecutora     int             `orm:"column(unidad_ejecutora)"`
	Responsable         int             `orm:"column(responsable)"`
	Vigencia            int             `orm:"column(vigencia)"`
	Observaciones       string          `orm:"column(observaciones)"`
	NumeroChequeInicial int             `orm:"column(numero_cheque_inicial)"`
	NumeroChequeFinal   int             `orm:"column(numero_cheque_final)"`
	CuentaBancaria      *CuentaBancaria `orm:"column(cuenta_bancaria);rel(fk)"`
	ChequesDisponibles	int							`orm:"column(cheques_disponibles)"`
}

func (t *Chequera) TableName() string {
	return "chequera"
}

func init() {
	orm.RegisterModel(new(Chequera))
}

// AddChequera insert a new Chequera into database and returns
// last inserted Id on success.
func AddChequera(m *Chequera) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetChequeraById retrieves Chequera by Id. Returns error if
// Id doesn't exist
func GetChequeraById(id int) (v *Chequera, err error) {
	o := orm.NewOrm()
	v = &Chequera{Id: id}
	if err = o.Read(v); err == nil {
		_, err = o.LoadRelated(v, "cuenta_bancaria")
		return v, nil
	}
	return nil, err
}

// GetAllChequera retrieves all Chequera matches certain condition. Returns empty list if
// no records exist
func GetAllChequera(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Chequera)).RelatedSel(1)
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

	var l []Chequera
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


// GetRecordsChequera retrieves quantity of records in Chequera s table
// Id doesn't exist
func GetRecordsChequera(query map[string]string) (cnt int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Chequera))

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

// validates if checker hasn't aviable items if not
// changes checker state
func ValidateSpentChecker(parameter ...interface{}) (err interface{}) {
	var estadoChequera EstadoChequera
	beego.Error("parameters ",parameter[0])
	beego.Error("parameters ",parameter[1])

	chequera := parameter[0].(Chequera)
	usuario := parameter[1].(int)

	o := orm.NewOrm()
	o.Begin()
	if (chequera.ChequesDisponibles - 1) <= 0 {
		err = o.QueryTable("estado_chequera").
			Filter("numeroOrden", 1).
			One(&estadoChequera)
			if err != nil {
				beego.Error(err)
				return
			}
			chequeraEstadoChequera := &ChequeraEstadoChequera{Chequera:&chequera,Activo:true,Estado:&estadoChequera,Usuario:usuario}
			_,err = o.Insert(chequeraEstadoChequera)
			if err != nil {
				beego.Error(err)
				o.Rollback()
				return
			}
			_, err = o.QueryTable("chequera_estado_chequera").
				Filter("chequera", chequera.Id).
				Filter("activo", true).
				Update(orm.Params{
					"activo": "false",
				})
				if err != nil {
					beego.Error(err)
					o.Rollback()
					return
				}
				o.Commit()
	}
				return
}

// UpdateChequera updates Chequera by Id and returns error if
// the record to be updated doesn't exist
func UpdateChequeraById(m *Chequera) (err error) {
	o := orm.NewOrm()
	v := Chequera{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteChequera deletes Chequera by Id and returns error if
// the record to be deleted doesn't exist
func DeleteChequera(id int) (err error) {
	o := orm.NewOrm()
	v := Chequera{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Chequera{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}


// AddChequera insert a new Chequera into database and returns
// last inserted Id on success.
func AddChequeraEstado(m map[string]interface{}) (id int64, err error) {
	var chequera Chequera
	var estadoChequera EstadoChequera
	var usuario float64
	var idChequera int64
	var consec float64
	o := orm.NewOrm()
	err = formatdata.FillStruct(m["Chequera"], &chequera)
	err = formatdata.FillStruct(m["Usuario"], &usuario)

	if err != nil {
		beego.Error(err.Error())
		return
	}
	o.Begin()

	qb, _ := orm.NewQueryBuilder("mysql")

	qb.Select("COALESCE(MAX(c.consecutivo),0)+1").
		From("chequera c").
		Where("c.vigencia = ?")

	sql := qb.String()

	err = o.Raw(sql, chequera.Vigencia).QueryRow(&consec)

	if err != nil {
		o.Rollback()
	}
	chequera.Consecutivo = int(consec)

	idChequera, err = o.Insert(&chequera)
	if err == nil {
		chequera.Id = int(idChequera)
		m["Chequera"]=chequera
		err = o.QueryTable("estado_chequera").
			Filter("numeroOrden", 1).
			One(&estadoChequera)
		if err == nil {
			chequeraEstadoChequera := &ChequeraEstadoChequera{Chequera:&chequera,Activo:true,Estado:&estadoChequera,Usuario:int(usuario)}
			_,err = o.Insert(chequeraEstadoChequera)
			if err != nil {
				beego.Error(err.Error())
				o.Rollback()
				return
			}
		}else{
			beego.Error(err.Error())
			o.Rollback()
			return
		}
	}else{
		beego.Error(err.Error())
		o.Rollback()
		return
	}
	o.Commit()
	return
}
