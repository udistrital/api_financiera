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

type DetallePac struct {
	Id                   int     `orm:"column(id);pk;auto"`
	ValorProyectadoMes   float64 `orm:"column(valor_proyectado_mes)"`
	ValorEjecutadoMes    float64 `orm:"column(valor_ejecutado_mes)"`
	FuenteFinanciamiento int     `orm:"column(fuente_financiamiento);null"`
	Rubro                int     `orm:"column(rubro)"`
	Pac                  *Pac    `orm:"column(pac);rel(fk)"`
	Mes                  int     `orm:"column(mes)"`
}

func (t *DetallePac) TableName() string {
	return "detalle_pac"
}

func init() {
	orm.RegisterModel(new(DetallePac))
}

// AddDetallePac insert a new DetallePac into database and returns
// last inserted Id on success.
func AddDetallePac(m *DetallePac) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetDetallePacById retrieves DetallePac by Id. Returns error if
// Id doesn't exist
func GetDetallePacById(id int) (v *DetallePac, err error) {
	o := orm.NewOrm()
	v = &DetallePac{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllDetallePac retrieves all DetallePac matches certain condition. Returns empty list if
// no records exist
func GetAllDetallePac(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(DetallePac))
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

	var l []DetallePac
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

// UpdateDetallePac updates DetallePac by Id and returns error if
// the record to be updated doesn't exist
func UpdateDetallePacById(m *DetallePac) (err error) {
	o := orm.NewOrm()
	v := DetallePac{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteDetallePac deletes DetallePac by Id and returns error if
// the record to be deleted doesn't exist
func DeleteDetallePac(id int) (err error) {
	o := orm.NewOrm()
	v := DetallePac{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&DetallePac{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func AddIngresoPac(parameter ...interface{}) (err interface{}) {

	var IdFuente int
	var detPac DetallePac
	var TotalEjecutado float64
	var codigo string
	var errs error
	ingreso := parameter[0].(Ingreso)
	vigencia := ingreso.Vigencia
	mes := int(ingreso.FechaIngreso.Month())

	fuenteFinan := ingreso.FuenteFinanciamiento
	if fuenteFinan != nil {
		codigo = fuenteFinan.Codigo
	}

	IdFuente, errs = strconv.Atoi(codigo)

	if errs != nil {
		IdFuente = 0
		formatdata.FillStruct(errs.Error(), &err)
		beego.Error(errs)
		return
	}

	pac, errs := GetPacByVigencia(vigencia, mes)
	o := orm.NewOrm()
	o.Begin()
	for _, ingConcep := range ingreso.IngresoConcepto {

		concepto := ingConcep.Concepto
		rubro := concepto.Rubro

		beego.Info("rubro ", rubro.Id)
		beego.Info("valorAgregado ", ingConcep.ValorAgregado)

		errs = o.QueryTable("detalle_pac").
			Filter("mes", mes).
			Filter("pac", pac.Id).
			Filter("fuente_financiamiento", IdFuente).
			Filter("rubro", rubro.Id).
			One(&detPac)

		if errs == orm.ErrMultiRows {
			fmt.Println("Returned Multi Rows Not One")
			return
		}
		if errs == orm.ErrNoRows {

			detalle_pac := &DetallePac{ValorProyectadoMes: -1,
				ValorEjecutadoMes:    ingConcep.ValorAgregado,
				FuenteFinanciamiento: IdFuente,
				Rubro:                rubro.Id,
				Pac:                  &pac,
				Mes:                  mes}

			_, errs = o.Insert(detalle_pac)

			if errs != nil {
				formatdata.FillStruct(errs.Error(), &err)
				beego.Info(err)
				o.Rollback()
				return
			}

		}
		if errs == nil {
			TotalEjecutado = detPac.ValorEjecutadoMes + ingConcep.ValorAgregado
			detPac.ValorEjecutadoMes = TotalEjecutado
			var num int64
			if num, errs = o.Update(&detPac); errs == nil {
				fmt.Println("Number of records updated in database:", num)
			} else {
				beego.Info(errs.Error())
				o.Rollback()
				return
			}
		}
	}
	o.Commit()
	return
}

func AddPacCierre(v []map[string]interface{}, mes int, vigencia int) (detPac DetallePac, err error) {
	var proy string
	var ejec string
	var idFuente string
	var idRubro string
	var proyec float64
	var ejecu float64
	var idF int
	var idR int

	o := orm.NewOrm()
	o.Begin()
	//insert pac
	pac, errs := GetPacByVigencia(float64(vigencia), mes)
	if errs != nil {
		beego.Info(err.Error())
		o.Rollback()
		return
	}

	for _, registroInsertar := range v {
		formatdata.FillStruct(registroInsertar["Proyeccion"], &proy)
		formatdata.FillStruct(registroInsertar["Valor"], &ejec)
		formatdata.FillStruct(registroInsertar["Idfuente"], &idFuente)
		err = formatdata.FillStruct(registroInsertar["Idrubro"], &idRubro)
		if err != nil {
			beego.Info(err.Error())
		}
		fmt.Println("id rubro ", registroInsertar["Idrubro"])
		fmt.Println("id rubro  en estructura ", idRubro)
		proyec, err = strconv.ParseFloat(proy, 64)
		ejecu, err = strconv.ParseFloat(ejec, 64)
		idF, err = strconv.Atoi(idFuente)
		idR, err = strconv.Atoi(idRubro)
		detalle_pac := &DetallePac{ValorProyectadoMes: proyec,
			ValorEjecutadoMes:    ejecu,
			FuenteFinanciamiento: idF,
			Rubro:                idR,
			Pac:                  &pac,
			Mes:                  mes}
		_, err = o.Insert(detalle_pac)
		if err != nil {
			beego.Info(err.Error())
			o.Rollback()
			return
		}
	}
	o.Commit()
	return
}
