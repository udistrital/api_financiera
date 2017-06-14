package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"
	"github.com/udistrital/api_financiera/utilidades"
	"github.com/astaxie/beego/orm"
	"strconv"
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
	Nomina               string                `orm:"column(nomina)"`
	Liquidacion          int                   `orm:"column(liquidacion);null"`
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
func RegistrarOpProveedor(m *Data_OrdenPago_Concepto) (alerta []string, err error, id_OrdenPago int64) {
	o := orm.NewOrm()
	o.Begin()
	// Inserta datos Orden de pago
	m.OrdenPago.FechaCreacion = time.Now()
	m.OrdenPago.Nomina = "PROVEEDOR"
	m.OrdenPago.EstadoOrdenPago = &EstadoOrdenPago{Id: 1} //1 Elaborado

	id_OrdenPago, err1 := o.Insert(&m.OrdenPago)
	if err1 != nil {
		alerta = append(alerta, "ERROR_1 [RegistrarOpProveedor] No se puede registrar la Orden de Pago")
		err = err1
		o.Rollback()
		return
	}
	// Insertar data Conceptos
	for i := 0; i < len(m.ConceptoOrdenPago); i++ {
		m.ConceptoOrdenPago[i].OrdenDePago = &OrdenPago{Id: int(id_OrdenPago)}
		_, err2 := o.Insert(&m.ConceptoOrdenPago[i])
		if err2 != nil {
			alerta = append(alerta, "ERROR_2 [RegistrarOpProveedor] No se puede registrar los Conceptos asociados a la Orden de Pago")
			err = err2
			o.Rollback()
		}
	}
	// Insertar data Movimientos Contables
	for i := 0; i < len(m.MovimientoContable); i++ {
		movimiento_contable := MovimientoContable{
			Debito:                   m.MovimientoContable[i].Debito,
			Credito:                  m.MovimientoContable[i].Credito,
			Fecha:                    time.Now(),
			Concepto:                 m.MovimientoContable[i].Concepto,
			CuentaContable:           m.MovimientoContable[i].CuentaContable,
			TipoDocumentoAfectante:   &TipoDocumentoAfectante{Id: 1}, //documento afectante tipo op
			CodigoDocumentoAfectante: int(id_OrdenPago),
			Aprobado:                 false,
		}
		_, err3 := o.Insert(&movimiento_contable)
		if err3 != nil {
			alerta = append(alerta, "ERROR_3 [RegistrarOpProveedor] No se puede registrar las Cuentas Contables Asociadas a los Concepto")
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
	// Actualizar datos de la Orden
	orden := OrdenPago{Id: m.OrdenPago.Id}
	if o.Read(&orden) == nil {
		orden.Iva = m.OrdenPago.Iva
		orden.TipoOrdenPago = m.OrdenPago.TipoOrdenPago
		orden.ValorBase = m.OrdenPago.ValorBase
		if _, err1 := o.Update(&orden); err1 != nil {
			fmt.Println("Error 1")
			alerta = append(alerta, "ERRRO_1 [ActualizarOpProveedor] No se puede actualizar los Campos de la Orden de Pago")
			err = err1
			o.Rollback()
		}
	}
	// Eliminar Conceptos Orden de Pagos y Movimientos contables
	if len(m.ConceptoOrdenPago) > 0 {
		_, err2 := o.Raw("DELETE FROM financiera.concepto_orden_pago where orden_de_pago = ?", m.OrdenPago.Id).Exec()
		if err2 != nil {
			alerta = append(alerta, "ERROR_02 [ActualizarOpProveedor] No se puede Eliminar los conceptos relacionados a la orden de pago")
			err = err2
			o.Rollback()
		}
	}
	if len(m.MovimientoContable) > 0 {
		_, err3 := o.Raw("DELETE FROM financiera.movimiento_contable where codigo_documento_afectante = ?", m.OrdenPago.Id).Exec()
		if err3 != nil {
			alerta = append(alerta, "ERROR_03 [ActualizarOpProveedor] No se puede Eliminar los movimientos contables relacionados a la orden de pago")
			err = err3
			o.Rollback()
		}
	}
	// Insertar Nueva Data Conceptos Orden de Pagos y Movimientos contables
	//Conceptos
	for i := 0; i < len(m.ConceptoOrdenPago); i++ {
		m.ConceptoOrdenPago[i].OrdenDePago = &OrdenPago{Id: int(m.OrdenPago.Id)}
		_, err4 := o.Insert(&m.ConceptoOrdenPago[i])
		if err4 != nil {
			alerta = append(alerta, "ERROR_04 [ActualizarOpProveedor] No se puede registrar los Conceptos asociados a la Orden de Pago")
			err = err4
			o.Rollback()
		}
	}
	//Movimientos
	for i := 0; i < len(m.MovimientoContable); i++ {
		movimiento_contable := MovimientoContable{
			Debito:                   m.MovimientoContable[i].Debito,
			Credito:                  m.MovimientoContable[i].Credito,
			Fecha:                    time.Now(),
			Concepto:                 m.MovimientoContable[i].Concepto,
			CuentaContable:           m.MovimientoContable[i].CuentaContable,
			TipoDocumentoAfectante:   &TipoDocumentoAfectante{Id: 1}, //documento afectante tipo op
			CodigoDocumentoAfectante: int(m.OrdenPago.Id),
			Aprobado:                 false,
		}
		_, err5 := o.Insert(&movimiento_contable)
		if err5 != nil {
			alerta = append(alerta, "ERROR_05 [ActualizarOpProveedor] No se puede registrar las Cuentas Contables Asociadas a los Concepto")
			err = err5
			o.Rollback()
		}
	}
	o.Commit()
	return
}

// personalizado Registrar orden_pago nomina planta, homologa conceptos titan-kronos, concepto_ordenpago y transacciones
func RegistrarOpPlanta(OrdenDetalle map[string]interface{} ) (alerta []string, err error, id_OrdenPago int64) {
	fmt.Println("Models Registro OP Planta")
	o := orm.NewOrm()
	o.Begin()
	new_orden := OrdenPago{}
	var detalle []interface{}
	err = utilidades.FillStruct(OrdenDetalle["OrdenPago"], &new_orden)
	err = utilidades.FillStruct(OrdenDetalle["DetalleLiquidacion"], &detalle)
	//homologacion := HomologacionConcepto{}
	var all_concepto_orden_pago []ConceptoOrdenPago

	// Datos Orden de Pago Planta
	new_orden.FechaCreacion = time.Now()
	new_orden.Nomina = "PLANTA"
	new_orden.EstadoOrdenPago = &EstadoOrdenPago{Id: 1} //1 Elaborado
	new_orden.Iva = &Iva{Id: 1} //1 iva del 0%
	new_orden.TipoOrdenPago = &TipoOrdenPago{Id: 2} //2 cuenta de cobro

	// insertar OP Planta
	id_OrdenPago, err1 := o.Insert(&new_orden)
	if err1 != nil {
			alerta = append(alerta, "ERROR_1 [RegistrarOpPlanta] No se puede registrar la Orden de Pago")
			err = err1
			o.Rollback()
			return
	}

	// Agrupar valores por conceptos del detalle de la liquidacionssss
	for i,element := range detalle{
		det := element.(map[string]interface{})
		var idconceptotitan int
		var valorcalculado int64
		err = utilidades.FillStruct(det["ValorCalculado"], &valorcalculado)
		conc := det["Concepto"].(map[string]interface{})
		err = utilidades.FillStruct(conc["Id"], &idconceptotitan)

		fmt.Println("****************************** ",strconv.Itoa(i), " ******************************" )
		if i == 0 {
			// Buscamos concepto kronos homologado
			concepto_kronos_homologado := HomologacionConcepto{ConceptoTitan: idconceptotitan, Vigencia: new_orden.Vigencia}
			err2 := o.Read(&concepto_kronos_homologado, "ConceptoTitan", "Vigencia")
			if err2 != nil {
				alerta = append(alerta, "ERROR_02 [RegistrarOpPlanta] No se Encontro Concepto Homoloago")
				err = err2
				o.Rollback()
			}
			fmt.Println("***Priner append ***")
			fmt.Println("Concepto Titan: ", strconv.Itoa(idconceptotitan))
			fmt.Println(valorcalculado)
			fmt.Println("Concepto Kronos:", strconv.Itoa(concepto_kronos_homologado.ConceptoKronos.Id))
			new_concepto_orden := ConceptoOrdenPago {
				Valor: valorcalculado,
				Concepto: &Concepto{Id: concepto_kronos_homologado.ConceptoKronos.Id},
			}
			all_concepto_orden_pago = append(all_concepto_orden_pago, new_concepto_orden)
		}else{
			fmt.Println("***Sumar o Append***")
			// Buscamos concepto kronos homologado
			concepto_kronos_homologado := HomologacionConcepto{ConceptoTitan: idconceptotitan, Vigencia: new_orden.Vigencia}
			err2 := o.Read(&concepto_kronos_homologado, "ConceptoTitan", "Vigencia")
			if err2 != nil {
				alerta = append(alerta, "ERROR_02 [RegistrarOpPlanta] No se Encontro Concepto Homoloago")
				err = err2
				o.Rollback()
			}
			fmt.Println("Concepto Titan: ", strconv.Itoa(idconceptotitan))
			fmt.Println(valorcalculado)
			fmt.Println("Concepto Kronos:", strconv.Itoa(concepto_kronos_homologado.ConceptoKronos.Id))

			if esta, idlista := estaConcepto(concepto_kronos_homologado.ConceptoKronos.Id, all_concepto_orden_pago); esta == true {
				fmt.Println("---Sumar")
				suma_valor := all_concepto_orden_pago[idlista].Valor + valorcalculado
				all_concepto_orden_pago[idlista].Valor = suma_valor
			}else{
				fmt.Println("---Append")
				new_concepto_orden2 := ConceptoOrdenPago {
					Valor: valorcalculado,
					Concepto: &Concepto{Id: concepto_kronos_homologado.ConceptoKronos.Id},
				}
				all_concepto_orden_pago = append(all_concepto_orden_pago, new_concepto_orden2)
			}
		}
		// consulta tabla de homologacion
		//homologacion = {}
		//homologacion.Vigencia = m.Vigencia
		//homologacion.ConceptoTitan = idconceptotitan
		//err = o.Read(&homologacion, "Vigencia, ConceptoTitan")
	}
	fmt.Println("*****************Totalizado**********************")
	fmt.Println(len(all_concepto_orden_pago))
	for i:=0; i< len(all_concepto_orden_pago); i++{
		fmt.Println("************ ", strconv.Itoa(all_concepto_orden_pago[i].Concepto.Id), " ************")
		fmt.Println(all_concepto_orden_pago[i].Valor)
		all_concepto_orden_pago[i].OrdenDePago = &OrdenPago{Id: int(id_OrdenPago)}
		// ¿se tendrá que validar el saldo del rubro??
		// insertar concepto_orden_pago
		_, err3 := o.Insert(&all_concepto_orden_pago[i])
		if err3 != nil {
				alerta = append(alerta, "ERROR_3 [RegistrarOpPlanta] No se puede registrar concepto_orden_pago")
				err = err3
				o.Rollback()
				return
		}
		// buscamos cuentas contables relacionadas al concepto para registrar movimientos
		qs := o.QueryTable(new(ConceptoCuentaContable)).RelatedSel()
		qs = qs.Filter("Concepto", all_concepto_orden_pago[i].Concepto.Id)
		qs = qs.RelatedSel()
		var l []ConceptoCuentaContable
		if _, err = qs.Limit(-1, 0).All(&l); err == nil {
			for _,v := range l {
				//registra movimientos
				fmt.Println("Data para Movimientos")
				//fmt.Println(v.CuentaContable.Naturaleza)
				//fmt.Println(v.Id)
				new_movimiento_contable := MovimientoContable{}
				if v.CuentaContable.Naturaleza == "debito" {
					fmt.Println(v.CuentaContable.Naturaleza)
					new_movimiento_contable.Debito = all_concepto_orden_pago[i].Valor
					new_movimiento_contable.Credito = 0
				}else{
					fmt.Println(v.CuentaContable.Naturaleza)
					new_movimiento_contable.Debito = 0
					new_movimiento_contable.Credito = all_concepto_orden_pago[i].Valor
				}
				new_movimiento_contable.Fecha = time.Now()
				new_movimiento_contable.Concepto = v.Concepto
				new_movimiento_contable.CuentaContable = v.CuentaContable
				new_movimiento_contable.TipoDocumentoAfectante = &TipoDocumentoAfectante{Id: 1} //documento afectante tipo op
				new_movimiento_contable.CodigoDocumentoAfectante = int(id_OrdenPago)
				new_movimiento_contable.Aprobado = false
				// insertar OP Planta
				_, err4 := o.Insert(&new_movimiento_contable)
				if err4 != nil {
						alerta = append(alerta, "ERROR_4 [RegistrarOpPlanta] No se puede registrar los Movimeitos Contables")
						err = err4
						fmt.Println("****ERRPR")
						fmt.Println(err4.Error())
						o.Rollback()
						return
				}
			}
		}
		fmt.Println("----------------------------")
	}
	fmt.Println("*****************FIN Totalizado**********************")
	o.Commit()
	return
}
//
func estaConcepto(idConcepto int, lista []ConceptoOrdenPago) (esta bool, idlista int) {
	for or := 0; or < len(lista); or++ {
		if lista[or].Concepto.Id == idConcepto {
			return true, or
		}
	}
	return false, 0
}

// personalizado Retrona la fecha actual del servidor
func FechaActual(formato string)(fechaActual string, err error){
	hoy := time.Now()
	fechaActual = hoy.Format(formato)
	return
}
