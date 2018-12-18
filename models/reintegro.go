package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/udistrital/utils_oas/formatdata"
	"github.com/udistrital/utils_oas/optimize"
)

type Reintegro struct {
	Id          int              `orm:"column(id);pk;auto"`
	Consecutivo int              `orm:"column(consecutivo)"`
	Causal      *CausalReintegro `orm:"column(causal);rel(fk)"`
	Ingreso     *Ingreso         `orm:"column(ingreso);rel(fk)"`
	OrdenPago   *OrdenPago       `orm:"column(orden_pago);rel(fk)"`
	Disponible  bool             `orm:"column(disponible);null"`
}

func (t *Reintegro) TableName() string {
	return "reintegro"
}

func init() {
	orm.RegisterModel(new(Reintegro))
}

// AddReintegro insert a new Reintegro into database and returns
// last inserted Id on success.
func AddReintegro(m *Reintegro) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// AddReintegro insert a new Reintegro setting consecutivo into database and returns
// last inserted Id on success.
func AddReintegroConsec(v map[string]interface{}) (id int64, err error) {
	o := orm.NewOrm()
	var consec float64
	var ingreso Ingreso
	var reintegro Reintegro
	var ingresoRes Ingreso
	var formaIngreso FormaIngreso

	err = formatdata.FillStruct(v["Reintegro"], &reintegro)
	err = formatdata.FillStruct(v["Ingreso"], &ingreso)

	beego.Error("valor v", v)

	if err != nil {
		return
	}

	o.Begin()

	err = o.QueryTable("forma_ingreso").
		Filter("nombre", "REINTEGROS").
		One(&formaIngreso)
	if err != nil {
		o.Rollback()
	}
	ingreso.FormaIngreso = &formaIngreso
	qb, _ := orm.NewQueryBuilder("mysql")

	qb.Select("COALESCE(MAX(r.consecutivo),0)+1").
		From("ingreso i").
		InnerJoin("reintegro r").On("r.ingreso = i.id").
		Where("i.vigencia > ?")

	sql := qb.String()

	err = o.Raw(sql, ingreso.Vigencia).QueryRow(&consec)

	if err != nil {
		o.Rollback()
	}
	reintegro.Consecutivo = int(consec)
	v["Ingreso"] = ingreso
	if ingresoRes, err = AddIngresotr(v); err != nil {
		beego.Error(err)
		o.Rollback()
	} else {
		v["Ingreso"] = ingresoRes
		reintegro.Ingreso = &ingresoRes
		if id, err = o.Insert(&reintegro); err == nil {
			reintegro.Id = int(id)
			v["Reintegro"] = reintegro
			o.Commit()
			return
		} else {
			beego.Error(err)
			o.Rollback()
			return
		}
	}
	o.Rollback()
	return
}

// GetReintegroById retrieves Reintegro by Id. Returns error if
// Id doesn't exist
func GetReintegroById(id int) (v *Reintegro, err error) {
	o := orm.NewOrm()
	v = &Reintegro{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllReintegro retrieves all Reintegro matches certain condition. Returns empty list if
// no records exist
func GetAllReintegro(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Reintegro)).RelatedSel()
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

	var l []Reintegro
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

//get reintegros which his state belongs to a certain number
func GetAllReintegroDisponible(query map[string]string, offset int, limit int) (ml map[string]interface{}, err error) {
	var reintegrosDisp []Reintegro
	var parametros []string
	var where string
	var cnt int64

	ml = make(map[string]interface{})

	qb, _ := orm.NewQueryBuilder("mysql")

	for k, v := range query {
		if strings.Contains(k, "mayor") {
			where = where + strings.Replace(k, "mayor", " > ?", -1)
		} else {
			where = where + k + " = ?"
		}
		where = where + " and "
		parametros = append(parametros, v)
	}
	where = strings.TrimRight(where, " and ")

	qb.Select("distinct r.*").
		From("financiera.reintegro r").
		InnerJoin("financiera.ingreso i").On("i.id = r.ingreso").
		InnerJoin("financiera.ingreso_estado_ingreso iei").On("iei.ingreso = i.id").
		Where(where).
		Limit(limit).Offset(offset)

	sql := qb.String()

	o := orm.NewOrm()
	o.Raw(sql, parametros).QueryRows(&reintegrosDisp)

	qb, _ = orm.NewQueryBuilder("mysql")

	qb.Select("count(1)").
		From("financiera.reintegro r").
		InnerJoin("financiera.ingreso i").On("i.id = r.ingreso").
		InnerJoin("financiera.ingreso_estado_ingreso iei").On("iei.ingreso = i.id").
		Where(where)

	sql = qb.String()
	o.Raw(sql, parametros).QueryRow(&cnt)

	reintegroInt := make([]interface{}, len(reintegrosDisp))
	for i, v := range reintegrosDisp {
		reintegroInt[i] = v
	}
	if reintegrosDisp != nil {
		reintegroInt = optimize.ProccDigest(reintegroInt, getReintegroList, nil, 3)
	}
	if err != nil {
		beego.Error(" error  load related sel ", err.Error())
	}
	ml["reintegros"] = reintegroInt
	ml["cantidadReg"] = cnt
	return
}

func getReintegroList(rpintfc interface{}, params ...interface{}) (res interface{}) {
	var reintegro Reintegro
	err := formatdata.FillStruct(rpintfc, &reintegro)
	o := orm.NewOrm()
	_, err = o.LoadRelated(&reintegro, "Causal")
	if err != nil {
		return err
	}
	rpintfc = reintegro
	return rpintfc
}

// UpdateReintegro updates Reintegro by Id and returns error if
// the record to be updated doesn't exist
func UpdateReintegroById(m *Reintegro) (err error) {
	o := orm.NewOrm()
	v := Reintegro{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteReintegro deletes Reintegro by Id and returns error if
// the record to be deleted doesn't exist
func DeleteReintegro(id int) (err error) {
	o := orm.NewOrm()
	v := Reintegro{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Reintegro{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
