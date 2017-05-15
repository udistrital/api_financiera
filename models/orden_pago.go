package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Data_OrdenPago_Concepto struct {
	OrdenPago          OrdenPago
	ConceptoOrdenPago  []ConceptoOrdenPago
	MovimientoContable []MovimientoContable
}

type OrdenPago struct {
	Id                   int                   `orm:"column(id);pk;auto"`
	Vigencia             float64               `orm:"column(vigencia)"`
	FechaCreacion        time.Time             `orm:"column(fecha_creacion);type(date)"`
	RegistroPresupuestal *RegistroPresupuestal `orm:"column(registro_presupuestal);rel(fk)"`
	ValorBase            float64               `orm:"column(valor_base)"`
	PersonaElaboro       int                   `orm:"column(persona_elaboro)"`
	Convenio             int                   `orm:"column(convenio);null"`
	TipoOrdenPago        *TipoOrdenPago        `orm:"column(tipo_orden_pago);rel(fk)"`
	UnidadEjecutora      *UnidadEjecutora      `orm:"column(unidad_ejecutora);rel(fk)"`
	EstadoOrdenPago      *EstadoOrdenPago      `orm:"column(estado_orden_pago);rel(fk)"`
	Iva                  *Iva                  `orm:"column(iva);rel(fk)"`
}

func (t *OrdenPago) TableName() string {
	return "orden_pago"
}

func init() {
	orm.RegisterModel(new(OrdenPago))
}

// AddOrdenPago insert a new OrdenPago into database and returns
// last inserted Id on success.
func AddOrdenPago(m *OrdenPago) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetOrdenPagoById retrieves OrdenPago by Id. Returns error if
// Id doesn't exist
func GetOrdenPagoById(id int) (v *OrdenPago, err error) {
	o := orm.NewOrm()
	v = &OrdenPago{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllOrdenPago retrieves all OrdenPago matches certain condition. Returns empty list if
// no records exist
func GetAllOrdenPago(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(OrdenPago)).RelatedSel()
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

	var l []OrdenPago
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

// UpdateOrdenPago updates OrdenPago by Id and returns error if
// the record to be updated doesn't exist
func UpdateOrdenPagoById(m *OrdenPago) (err error) {
	o := orm.NewOrm()
	v := OrdenPago{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteOrdenPago deletes OrdenPago by Id and returns error if
// the record to be deleted doesn't exist
func DeleteOrdenPago(id int) (err error) {
	o := orm.NewOrm()
	v := OrdenPago{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&OrdenPago{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

// personalizado Registrar orden_pago, concepto_ordenpago y transacciones
func RegistrarOp(m *Data_OrdenPago_Concepto) (alerta []string, err error, id_OrdenPago int64) {
	o := orm.NewOrm()
	o.Begin()
	m.OrdenPago.FechaCreacion = time.Now()
	id_OrdenPago, err1 := o.Insert(&m.OrdenPago)
	if err1 != nil {
		alerta = append(alerta, "Erro 1. No se puede registrar la Orden de Pago")
		err = err1
		o.Rollback()
		return
	}
	//fmt.Println(id_OrdenPago)
	//fmt.Println("recorrer")
	for i := 0; i < len(m.ConceptoOrdenPago); i++ {
		m.ConceptoOrdenPago[i].OrdenDePago = &OrdenPago{Id: int(id_OrdenPago)}
		_, err2 := o.Insert(&m.ConceptoOrdenPago[i])
		if err2 != nil {
			alerta = append(alerta, "Erro 2. No se puede registrar los Conceptos asociados a la Orden de Pago")
			err = err2
			o.Rollback()
		}
	}

	for i := 0; i < len(m.MovimientoContable); i++ {
		movimiento_contable := MovimientoContable{
			Debito:  m.MovimientoContable[i].Debito,
			Credito: m.MovimientoContable[i].Credito,
			Fecha:   time.Now(),
			ConceptoCuentaContable:   &ConceptoCuentaContable{Id: int(m.MovimientoContable[i].Id)},
			TipoDocumentoAfectante:   &TipoDocumentoAfectante{Id: 1}, //quemado
			CodigoDocumentoAfectante: int(id_OrdenPago),
			Aprobado:                 false,
		}
		_, err3 := o.Insert(&movimiento_contable)
		if err3 != nil {
			alerta = append(alerta, "Erro 3. No se puede registrar las Cuentas Contables Asociadas a los Concepto")
			err = err3
			o.Rollback()
		}
	}
	o.Commit()
	return
}

// personalizado Actualiza orden_pago, concepto_ordenpago y movimeintos contalbes
func ActualizarOpProveedor(m *Data_OrdenPago_Concepto) (alerta []string, err error, id_OrdenPago int64) {
	o := orm.NewOrm()
	o.Begin()
	fmt.Println("AAAAAAAA")
	fmt.Println(m.OrdenPago.Id)
	fmt.Println("AAAAAAAA")
	// Actualizar datos de la Orden
	orden := OrdenPago{Id: m.OrdenPago.Id}
	if o.Read(&orden) == nil {
	    orden.Iva = m.OrdenPago.Iva
			orden.TipoOrdenPago = m.OrdenPago.TipoOrdenPago
			orden.ValorBase = m.OrdenPago.ValorBase
	    if _, err1 := o.Update(&orden); err1 != nil {
				fmt.Println("Error 1")
				alerta = append(alerta, "Erro 1. No se puede actualizar los Campos de la Orden de Pago")
				err = err1
				o.Rollback()
	    }
	}
	// Eliminar Conceptos Orden de Pagos y Movimientos contables
	o.Commit()
	return
}
