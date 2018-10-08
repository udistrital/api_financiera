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

type DevolucionTributaria struct {
	Id               		 int               			 `orm:"column(id);pk;auto"`
	Vigencia         		 float64           			 `orm:"column(vigencia)"`
	UnidadEjecutora  		 *UnidadEjecutora  			 `orm:"column(unidad_ejecutora);rel(fk)"`
	Acta             		 int               			 `orm:"column(acta)"`
	Solicitante      		 int               			 `orm:"column(solicitante)"`
	FormaPago        		 *FormaPago        			 `orm:"column(forma_pago);rel(fk)"`
	CuentaBancariaEnte 		 *CuentaBancariaEnte 			 `orm:"column(cuenta_devolucion);rel(fk)"`
	Justificacion    		 string            			 `orm:"column(justificacion)"`
	DocumentoGenerador   *DocumentoGenerador     `orm:"column(documento_generador);rel(fk);null"`
	FechaRegistro    		 time.Time         			 `orm:"column(fecha_registro);auto_now_add;type(datetime)"`
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
	qs := o.QueryTable(new(DevolucionTributaria)).RelatedSel(2)
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

	var Id int64
	var idDevol int64
	var cuentaDevol CuentaBancariaEnte
  var estadoDevolucion EstadoDevolucion
	var movimientosAsoc  []DevolucionTributariaMovimientoAsociado
	var docGen DocumentoGenerador

	var concepto []map[string]interface{}
	var mov []MovimientoContable

	o := orm.NewOrm()

	err = formatdata.FillStruct(request["DevolucionTributaria"], &tributariaDevol)
	err = formatdata.FillStruct(request["Movimientos"], &mov)
	err = formatdata.FillStruct(request["MovimientosAsociados"], &movimientosAsoc)
	err = formatdata.FillStruct(request["Concepto"], &concepto)
	err = formatdata.FillStruct(request["DocumentoGenerador"], &docGen)

	if err != nil {
		beego.Error(err)
		return
	}
		o.Begin()

	idDocgenerador, err := o.Insert(&docGen)
			if err != nil {
				beego.Info(err)
				o.Rollback()
				return
			}

		err = o.QueryTable("cuenta_bancaria_ente").
			Filter("banco", tributariaDevol.CuentaBancariaEnte.Banco).
			Filter("tipo_cuenta", tributariaDevol.CuentaBancariaEnte.TipoCuenta).
			Filter("numero_cuenta", tributariaDevol.CuentaBancariaEnte.NumeroCuenta).
			One(&cuentaDevol)

			if err == nil {
				tributariaDevol.CuentaBancariaEnte.Id = cuentaDevol.Id
			} else if err == orm.ErrMultiRows {
				beego.Error("Returned Multi Rows Not One")
				return
			} else if err == orm.ErrNoRows {
				Id, err = o.Insert(tributariaDevol.CuentaBancariaEnte)
				if err != nil {
					beego.Error(err)
					o.Rollback()
					return
				} else {
					tributariaDevol.CuentaBancariaEnte.Id = int(Id)
				}
	}

		tributariaDevol.DocumentoGenerador = &DocumentoGenerador{Id: int(idDocgenerador)}

		idDevol, err = o.Insert(&tributariaDevol)
		if err != nil {
			beego.Error(err)
			o.Rollback()
			return
		}
		tributariaDevol.Id = int(idDevol)
		if err != nil {
			o.Rollback()
			return
		}

		err = o.QueryTable(new(EstadoDevolucion)).
			Filter("numeroOrden", 1).
			Filter("tipo", 3).
			One(&estadoDevolucion)

		if err != nil {
			beego.Error(err.Error())
			o.Rollback()
			return
		}
			devolucionEstadoDevolucion := &DevolucionTributariaEstadoDevolucion{Devolucion: &tributariaDevol, Activo: true, EstadoDevolucion: &estadoDevolucion}
			_, err = o.Insert(devolucionEstadoDevolucion)
			if err != nil {
				beego.Error(err.Error())
				o.Rollback()
				return
			}
			for i,_ := range movimientosAsoc {
				movimientosAsoc[i].Devolucion = &tributariaDevol;
			}
			_, err = o.InsertMulti(100, movimientosAsoc)

			if err != nil {
				beego.Error(err.Error())
				o.Rollback()
				return
			}

		for _, element := range concepto {
			conceptoDevol := &Concepto{Id: int(element["Id"].(float64))}
			devolucion_concepto := &DevolucionTributariaConcepto{ValorDevolucion: element["valorAfectacion"].(float64),
				DevolucionTributaria: &tributariaDevol,
				Concepto:             conceptoDevol,
			}
			_, err = o.Insert(devolucion_concepto)
			if err != nil {
				beego.Info(err.Error())
				o.Rollback()
				return
			}

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
	o.Commit()
	return
}

// GetRecordsNumberDevolucion retrieves quantity of records in tibutary devolutions table
// Id doesn't exist returns 0
func GetRecordsNumberDevolucion(query map[string]string) (cnt int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(DevolucionTributaria))

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
