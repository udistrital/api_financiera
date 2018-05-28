package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/udistrital/utils_oas/formatdata"
)

type DevolucionTributaria struct {
	Id               int                  `orm:"column(id);pk;auto"`
	Vigencia         float64              `orm:"column(vigencia)"`
	UnidadEjecutora  *UnidadEjecutora     `orm:"column(unidad_ejecutora);rel(fk)"`
	Acta             *ActaDevolucion      `orm:"column(acta);rel(fk)"`
	Oficio           int                  `orm:"column(oficio)"`
	FechaOficio      time.Time            `orm:"column(fecha_oficio);type(date)"`
	Solicitante      *DocumentoDevolucion `orm:"column(solicitante);rel(fk)"`
	FormaPago        *FormaPago           `orm:"column(forma_pago);rel(fk)"`
	CuentaDevolucion *CuentaDevolucion    `orm:"column(cuenta_devolucion);rel(fk)"`
	Justificacion    string               `orm:"column(justificacion)"`
	FechaRegistro    time.Time            `orm:"column(fecha_registro);auto_now_add;type(datetime)"`
}

func (t *DevolucionTributaria) TableName() string {
	return "devolucion_tributaria"
}

func init() {
	orm.RegisterModel(new(DevolucionTributaria))
}

// AddDevolucionTributaria insert a new DevolucionTributaria into database and returns
// last inserted Id on success.
func AddDevolucionTributaria(m *DevolucionTributaria) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetDevolucionTributariaById retrieves DevolucionTributaria by Id. Returns error if
// Id doesn't exist
func GetDevolucionTributariaById(id int) (v *DevolucionTributaria, err error) {
	o := orm.NewOrm()
	v = &DevolucionTributaria{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllDevolucionTributaria retrieves all DevolucionTributaria matches certain condition. Returns empty list if
// no records exist
func GetAllDevolucionTributaria(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(DevolucionTributaria)).RelatedSel()
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

	var l []DevolucionTributaria
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

// UpdateDevolucionTributaria updates DevolucionTributaria by Id and returns error if
// the record to be updated doesn't exist
func UpdateDevolucionTributariaById(m *DevolucionTributaria) (err error) {
	o := orm.NewOrm()
	v := DevolucionTributaria{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteDevolucionTributaria deletes DevolucionTributaria by Id and returns error if
// the record to be deleted doesn't exist
func DeleteDevolucionTributaria(id int) (err error) {
	o := orm.NewOrm()
	v := DevolucionTributaria{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&DevolucionTributaria{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

// add a devolution, state,a id type returns error
//if any insert fails
func AddDevolucionTr(request map[string]interface{}) (tributariaDevol DevolucionTributaria, err error) {

	var documentoSol *DocumentoDevolucion
	var documentoBusqeda DocumentoDevolucion
	var Id int64
	var idDevol int64
	var cuentaDevol CuentaDevolucion

	var concepto Concepto
	var mov []MovimientoContable
	var totalInv float64

	o := orm.NewOrm()

	err = formatdata.FillStruct(request["DevolucionTributaria"], &tributariaDevol)
	err = formatdata.FillStruct(request["Movimientos"], &mov)
	err = formatdata.FillStruct(request["TotalInversion"], &totalInv)
	err = formatdata.FillStruct(request["Concepto"], &concepto)

	if err == nil {
		o.Begin()

		documentoSol = tributariaDevol.Solicitante

		err = o.QueryTable("cuenta_devolucion").
			Filter("banco", tributariaDevol.CuentaDevolucion.Banco).
			Filter("tipo_cuenta", tributariaDevol.CuentaDevolucion.TipoCuenta).
			Filter("numero_cuenta", tributariaDevol.CuentaDevolucion.NumeroCuenta).
			One(&cuentaDevol)
		beego.Error(err)
		if err == orm.ErrMultiRows {
			beego.Error("Returned Multi Rows Not One")
			return
		}

		if err == orm.ErrNoRows {
			Id, err = o.Insert(tributariaDevol.CuentaDevolucion)
			beego.Error("id", Id, "error", err)
			if err != nil {
				beego.Error(err)
				o.Rollback()
				return
			} else {
				tributariaDevol.CuentaDevolucion.Id = int(Id)
			}
		}
		if err == nil {
			tributariaDevol.CuentaDevolucion.Id = cuentaDevol.Id
		}

		err = o.QueryTable("documento_devolucion").
			Filter("Origen", documentoSol.Origen).
			Filter("tipo_identificacion", documentoSol.TipoIdentificacion).
			Filter("identificacion", documentoSol.Identificacion).
			One(&documentoBusqeda)

		if err == orm.ErrMultiRows {
			beego.Error("Error consultado documento solicitante")
			return
		}

		if err == orm.ErrNoRows {
			Id, err = o.Insert(documentoSol)
			documentoSol.Id = int(Id)
			if err != nil {
				beego.Error(err)
				o.Rollback()
				return
			}
		}
		if err == nil {
			documentoSol.Id = documentoBusqeda.Id
		}
		lll, _ := json.Marshal(&tributariaDevol)
		beego.Info(string(lll))
		idDevol, err = o.Insert(&tributariaDevol)
		if err != nil {
			beego.Error(err)
			o.Rollback()
			return
		}
		beego.Error("id devolucion", idDevol)
		tributariaDevol.Id = int(idDevol)
		if err != nil {
			o.Rollback()
			return
		}

		devolucion_concepto := &DevolucionTributariaConcepto{ValorDevolucion: totalInv,
			DevolucionTributaria: &tributariaDevol,
			Concepto:             &concepto}

		_, err = o.Insert(devolucion_concepto)

		if err != nil {
			beego.Info(err.Error())
			o.Rollback()
			return
		}

		for _, element := range mov {
			element.Fecha = time.Now()
			element.TipoDocumentoAfectante = &TipoDocumentoAfectante{Id: 6}
			element.CodigoDocumentoAfectante = tributariaDevol.Id
			element.EstadoMovimientoContable = &EstadoMovimientoContable{Id: 1}
			_, err = o.Insert(&element)

			if err != nil {
				beego.Info(err.Error())
				o.Rollback()
				return
			}
		}

	} else {
		return
	}
	o.Commit()
	return
}