package models

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/udistrital/api_financiera/utilidades"
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
	EntradaAlmacen       int                   `orm:"column(entrada_almacen);null"`
	Consecutivo          int                   `orm:"column(consecutivo)"`
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
func RegistrarOpProveedor(m *Data_OrdenPago_Concepto) (alerta Alert, err error, consecutivoOp int) {
	var idOrdenPago int64
	o := orm.NewOrm()
	o.Begin()
	// Inserta datos Orden de pago
	o.Raw(`SELECT COALESCE(MAX(consecutivo), 0)+1 as consecutivo
			FROM financiera.orden_pago`).QueryRow(&consecutivoOp)
	m.OrdenPago.Consecutivo = consecutivoOp
	m.OrdenPago.FechaCreacion = time.Now()
	m.OrdenPago.Nomina = "PROVEEDOR"
	m.OrdenPago.EstadoOrdenPago = &EstadoOrdenPago{Id: 1} //1 Elaborado

	idOrdenPago, err = o.Insert(&m.OrdenPago)
	if err != nil {
		alerta.Type = "error"
		alerta.Code = "E_OPP_01"
		alerta.Body = err.Error()
		o.Rollback()
		return
	}
	// Insertar data Conceptos
	for i := 0; i < len(m.ConceptoOrdenPago); i++ {
		m.ConceptoOrdenPago[i].OrdenDePago = &OrdenPago{Id: int(idOrdenPago)}
		_, err = o.Insert(&m.ConceptoOrdenPago[i])
		if err != nil {
			alerta.Type = "error"
			alerta.Code = "E_OPP_02"
			alerta.Body = err.Error()
			o.Rollback()
			return
		}
	}
	// Insertar data Movimientos Contables
	for i := 0; i < len(m.MovimientoContable); i++ {
		movimientoContable := MovimientoContable{
			Debito:                   m.MovimientoContable[i].Debito,
			Credito:                  m.MovimientoContable[i].Credito,
			Fecha:                    time.Now(),
			ConceptoTesoral:          m.MovimientoContable[i].ConceptoTesoral,
			CuentaContable:           m.MovimientoContable[i].CuentaContable,
			TipoDocumentoAfectante:   &TipoDocumentoAfectante{Id: 1}, //documento afectante tipo op
			CodigoDocumentoAfectante: int(idOrdenPago),
			Aprobado:                 false,
		}
		_, err = o.Insert(&movimientoContable)
		if err != nil {
			alerta.Type = "error"
			alerta.Code = "E_OPP_03"
			alerta.Body = err.Error()
			o.Rollback()
			return
		}
	}
	o.Commit()
	return
}

// personalizado Actualiza orden_pago, concepto_ordenpago y movimeintos contalbes
func ActualizarOpProveedor(m *Data_OrdenPago_Concepto) (alerta Alert, err error, consecutivoOp int) {
	o := orm.NewOrm()
	o.Begin()
	// Actualizar datos de la Orden
	orden := OrdenPago{Id: m.OrdenPago.Id}
	if o.Read(&orden) == nil {
		orden.Iva = m.OrdenPago.Iva
		orden.TipoOrdenPago = m.OrdenPago.TipoOrdenPago
		orden.ValorBase = m.OrdenPago.ValorBase
		if _, err = o.Update(&orden); err != nil {
			alerta.Type = "error"
			alerta.Code = "E_OPP_UPD_01"
			alerta.Body = err.Error()
			o.Rollback()
			return
		} else {
			consecutivoOp = orden.Consecutivo
		}
	}
	// Eliminar Conceptos Orden de Pagos y Movimientos contables
	if len(m.ConceptoOrdenPago) > 0 {
		_, err = o.Raw("DELETE FROM financiera.concepto_orden_pago where orden_de_pago = ?", m.OrdenPago.Id).Exec()
		if err != nil {
			alerta.Type = "error"
			alerta.Code = "E_OPP_UPD_02"
			alerta.Body = err.Error()
			o.Rollback()
			return
		}
	}
	if len(m.MovimientoContable) > 0 {
		_, err = o.Raw("DELETE FROM financiera.movimiento_contable where codigo_documento_afectante = ?", m.OrdenPago.Id).Exec()
		if err != nil {
			alerta.Type = "error"
			alerta.Code = "E_OPP_UPD_03"
			alerta.Body = err.Error()
			o.Rollback()
			return
		}
	}
	// Insertar Nueva Data Conceptos Orden de Pagos y Movimientos contables
	//Conceptos
	for i := 0; i < len(m.ConceptoOrdenPago); i++ {
		m.ConceptoOrdenPago[i].OrdenDePago = &OrdenPago{Id: int(m.OrdenPago.Id)}
		_, err = o.Insert(&m.ConceptoOrdenPago[i])
		if err != nil {
			alerta.Type = "error"
			alerta.Code = "E_OPP_UPD_04"
			alerta.Body = err.Error()
			o.Rollback()
			return
		}
	}
	//Movimientos
	for i := 0; i < len(m.MovimientoContable); i++ {
		movimientoContable := MovimientoContable{
			Debito:                   m.MovimientoContable[i].Debito,
			Credito:                  m.MovimientoContable[i].Credito,
			Fecha:                    time.Now(),
			ConceptoTesoral:          m.MovimientoContable[i].ConceptoTesoral,
			CuentaContable:           m.MovimientoContable[i].CuentaContable,
			TipoDocumentoAfectante:   &TipoDocumentoAfectante{Id: 1}, //documento afectante tipo op
			CodigoDocumentoAfectante: int(m.OrdenPago.Id),
			Aprobado:                 false,
		}
		_, err = o.Insert(&movimientoContable)
		if err != nil {
			alerta.Type = "error"
			alerta.Code = "E_OPP_UPD_05"
			alerta.Body = err.Error()
			o.Rollback()
			return
		}
	}
	o.Commit()
	return
}

// personalizado Registrar orden_pago nomina planta, homologa conceptos titan-kronos, concepto_ordenpago y transacciones
func RegistrarOpNomina(OrdenDetalle map[string]interface{}) (alerta Alert, err error, consecutivoOp int) {
	var idOrdenPago int64
	o := orm.NewOrm()
	o.Begin()
	newOrden := OrdenPago{}
	var detalle []interface{}
	err = utilidades.FillStruct(OrdenDetalle["OrdenPago"], &newOrden)
	err = utilidades.FillStruct(OrdenDetalle["DetalleLiquidacion"], &detalle)
	var allConceptoOrdenPago []ConceptoOrdenPago

	// Datos Orden de Pago Planta
	o.Raw(`SELECT COALESCE(MAX(consecutivo), 0)+1 as consecutivo
			FROM financiera.orden_pago`).QueryRow(&consecutivoOp)
	newOrden.Consecutivo = consecutivoOp
	newOrden.FechaCreacion = time.Now()
	newOrden.Nomina = "PLANTA"
	newOrden.EstadoOrdenPago = &EstadoOrdenPago{Id: 1} //1 Elaborado
	newOrden.Iva = &Iva{Id: 1}                         //1 iva del 0%
	newOrden.TipoOrdenPago = &TipoOrdenPago{Id: 2}     //2 cuenta de cobro

	// insertar OP Planta
	idOrdenPago, err = o.Insert(&newOrden)
	if err != nil {
		alerta.Type = "error"
		alerta.Code = "E_OPN_02"
		alerta.Body = err.Error()
		o.Rollback()
		return
	}

	// Agrupar valores por conceptos del detalle de la liquidacion y guardamos su homologado
	for i, element := range detalle {
		det := element.(map[string]interface{})
		var idconceptotitan int
		// data valorCalculado
		val, ok := det["ValorCalculado"]
		if !ok {
			alerta.Type = "error"
			alerta.Code = "E_OPN_02_1"
			alerta.Body = ""
			o.Rollback()
			return
		}
		valorcalculadoFloat := val.(float64)
		valorcalculado := int64(valorcalculadoFloat)
		// data concepto
		conc := det["Concepto"].(map[string]interface{})
		err = utilidades.FillStruct(conc["Id"], &idconceptotitan)
		if err != nil {
			alerta.Type = "error"
			alerta.Code = "E_OPN_02_2"
			alerta.Body = err
			o.Rollback()
			return
		}

		fmt.Println("****************************** ", strconv.Itoa(i), " ******************************")
		if i == 0 {
			// Buscamos concepto kronos homologado
			conceptoKronosHomologado := HomologacionConcepto{ConceptoTitan: idconceptotitan, Vigencia: newOrden.Vigencia}
			err = o.Read(&conceptoKronosHomologado, "ConceptoTitan", "Vigencia")
			if err != nil {
				alerta.Type = "error"
				alerta.Code = "E_OPN_02_3"
				alerta.Body = strconv.Itoa(idconceptotitan)
				o.Rollback()
				return
			}
			fmt.Println("***Priner append ***")
			fmt.Println("Concepto Titan: ", strconv.Itoa(idconceptotitan))
			fmt.Println(valorcalculado)
			fmt.Println("Concepto Kronos:", strconv.Itoa(conceptoKronosHomologado.ConceptoKronos.Id))
			newConceptoOrden := ConceptoOrdenPago{
				Valor:    valorcalculado,
				Concepto: &ConceptoTesoral{Id: conceptoKronosHomologado.ConceptoKronos.Id},
			}
			allConceptoOrdenPago = append(allConceptoOrdenPago, newConceptoOrden)
		} else {
			fmt.Println("***Sumar o Append***")
			// Buscamos concepto kronos homologado
			conceptoKronosHomologado := HomologacionConcepto{ConceptoTitan: idconceptotitan, Vigencia: newOrden.Vigencia}
			err = o.Read(&conceptoKronosHomologado, "ConceptoTitan", "Vigencia")
			if err != nil {
				alerta.Type = "error"
				alerta.Code = "E_OPN_02_3"
				alerta.Body = strconv.Itoa(idconceptotitan)
				o.Rollback()
				return
			}
			fmt.Println("Concepto Titan: ", strconv.Itoa(idconceptotitan))
			fmt.Println(valorcalculado)
			fmt.Println("Concepto Kronos:", strconv.Itoa(conceptoKronosHomologado.ConceptoKronos.Id))

			if esta, idlista := estaConcepto(conceptoKronosHomologado.ConceptoKronos.Id, allConceptoOrdenPago); esta == true {
				fmt.Println("---Sumar")
				sumaValor := allConceptoOrdenPago[idlista].Valor + valorcalculado
				allConceptoOrdenPago[idlista].Valor = sumaValor
			} else {
				fmt.Println("---Append")
				newConceptoOrden2 := ConceptoOrdenPago{
					Valor:    valorcalculado,
					Concepto: &ConceptoTesoral{Id: conceptoKronosHomologado.ConceptoKronos.Id},
				}
				allConceptoOrdenPago = append(allConceptoOrdenPago, newConceptoOrden2)
			}
		}
	}
	fmt.Println("\n***************** Totalizado **********************")
	for i := 0; i < len(allConceptoOrdenPago); i++ {
		fmt.Println("\n************ Concepto kronos", strconv.Itoa(allConceptoOrdenPago[i].Concepto.Id), " ************")
		fmt.Println(allConceptoOrdenPago[i].Valor)
		allConceptoOrdenPago[i].OrdenDePago = &OrdenPago{Id: int(idOrdenPago)}
		// insertar concepto_orden_pago
		_, err = o.Insert(&allConceptoOrdenPago[i])
		if err != nil {
			alerta.Type = "error"
			alerta.Code = "E_OPN_02_4"
			alerta.Body = err
			o.Rollback()
			return
		}
		// buscamos cuentas contables relacionadas al concepto para registrar movimientos
		qs := o.QueryTable(new(ConceptoTesoralCuentaContable)).RelatedSel()
		qs = qs.Filter("Concepto", allConceptoOrdenPago[i].Concepto.Id)
		qs = qs.RelatedSel()
		var l []ConceptoTesoralCuentaContable
		if _, err = qs.Limit(-1, 0).All(&l); err == nil {
			for _, v := range l {
				//registra movimientos
				fmt.Println("Data para Movimientos")
				newMovimientoContable := MovimientoContable{}
				if v.CuentaContable.Naturaleza == "debito" {
					fmt.Println(v.CuentaContable.Naturaleza)
					newMovimientoContable.Debito = allConceptoOrdenPago[i].Valor
					newMovimientoContable.Credito = 0
				} else {
					fmt.Println(v.CuentaContable.Naturaleza)
					newMovimientoContable.Debito = 0
					newMovimientoContable.Credito = allConceptoOrdenPago[i].Valor
				}
				newMovimientoContable.Fecha = time.Now()
				newMovimientoContable.ConceptoTesoral = v.ConceptoTesoral
				newMovimientoContable.CuentaContable = v.CuentaContable
				newMovimientoContable.TipoDocumentoAfectante = &TipoDocumentoAfectante{Id: 1} //documento afectante tipo op
				newMovimientoContable.CodigoDocumentoAfectante = int(idOrdenPago)
				newMovimientoContable.Aprobado = false
				// insertar OP Planta
				_, err = o.Insert(&newMovimientoContable)
				if err != nil {
					alerta.Type = "error"
					alerta.Code = "E_OPN_02_5"
					alerta.Body = err
					o.Rollback()
					return
				}
			}
		}
	}
	fmt.Println("\n*****************FIN Totalizado**********************")
	o.Commit()
	return
}

// Registra orden_pago nomina Seguridad Social, homologa conceptos titan-kronos, concepto_ordenpago y transacciones
func RegistrarOpSeguridadSocial(OrdenDetalle map[string]interface{}) (alerta Alert, err error, consecutivoOp int) {
	var idOrdenPago int64
	o := orm.NewOrm()
	o.Begin()
	newOrden := OrdenPago{}
	var PagosSeguridadSocial []interface{}
	err = utilidades.FillStruct(OrdenDetalle["OrdenPago"], &newOrden)
	err = utilidades.FillStruct(OrdenDetalle["PagosSeguridadSocial"], &PagosSeguridadSocial)
	var allConceptoOrdenPago []ConceptoOrdenPago

	// Datos Orden de Pago Planta
	o.Raw(`SELECT COALESCE(MAX(consecutivo), 0)+1 as consecutivo
			FROM financiera.orden_pago`).QueryRow(&consecutivoOp)
	newOrden.Consecutivo = consecutivoOp
	newOrden.FechaCreacion = time.Now()
	newOrden.Nomina = "SEGURIDAD SOCIAL"
	newOrden.EstadoOrdenPago = &EstadoOrdenPago{Id: 1} //1 Elaborado
	newOrden.Iva = &Iva{Id: 1}                         //1 iva del 0%
	newOrden.TipoOrdenPago = &TipoOrdenPago{Id: 2}     //2 cuenta de cobro

	// insertar OP Planta
	idOrdenPago, err = o.Insert(&newOrden)
	if err != nil {
		alerta.Type = "error"
		alerta.Code = "E_OPN_02"
		alerta.Body = err.Error()
		o.Rollback()
		return
	}
	// Agrupar valores por conceptos del detalle de la liquidacion y guardamos su homologado
	for i, element := range PagosSeguridadSocial {
		det := element.(map[string]interface{})
		// data valor
		val, ok := det["Valor"]
		if !ok {
			alerta.Type = "error"
			alerta.Code = "E_OPN_02_1"
			alerta.Body = ""
			o.Rollback()
			return
		}
		ValorFloat := val.(float64)
		Valor := int64(ValorFloat)
		// data concepto
		val2, ok2 := det["TipoPago"]
		if !ok2 {
			alerta.Type = "error"
			alerta.Code = "E_OPN_02_1"
			alerta.Body = ""
			o.Rollback()
			return
		}
		idconceptotitanFloat := val2.(float64)
		idconceptotitan := int(idconceptotitanFloat)
		// data periodo pago
		PeriodoPago := det["PeriodoPago"].(map[string]interface{})
		anioSting := fmt.Sprintf("%v", PeriodoPago["Anio"])
		anio, err1 := strconv.ParseFloat(anioSting, 64)
		if err1 != nil {
			alerta.Type = "error"
			alerta.Code = "E_OPN_02_3"
			alerta.Body = strconv.Itoa(idconceptotitan)
			o.Rollback()
			return
		}

		fmt.Println("****************************** ", strconv.Itoa(i), " ******************************")
		if i == 0 {
			// Buscamos concepto kronos homologado
			conceptoKronosHomologado := HomologacionConcepto{ConceptoTitan: idconceptotitan, Vigencia: anio}
			err = o.Read(&conceptoKronosHomologado, "ConceptoTitan", "Vigencia")
			if err != nil {
				alerta.Type = "error"
				alerta.Code = "E_OPN_02_3"
				alerta.Body = strconv.Itoa(idconceptotitan)
				o.Rollback()
				return
			}
			fmt.Println("***Priner append ***")
			fmt.Println("Concepto Titan: ", strconv.Itoa(idconceptotitan))
			fmt.Println(Valor)
			fmt.Println("Concepto Kronos:", strconv.Itoa(conceptoKronosHomologado.ConceptoKronos.Id))
			newConceptoOrden := ConceptoOrdenPago{
				Valor:    Valor,
				Concepto: &ConceptoTesoral{Id: conceptoKronosHomologado.ConceptoKronos.Id},
			}
			allConceptoOrdenPago = append(allConceptoOrdenPago, newConceptoOrden)
		} else {
			fmt.Println("***Sumar o Append***")
			// Buscamos concepto kronos homologado
			conceptoKronosHomologado := HomologacionConcepto{ConceptoTitan: idconceptotitan, Vigencia: anio}
			err = o.Read(&conceptoKronosHomologado, "ConceptoTitan", "Vigencia")
			if err != nil {
				alerta.Type = "error"
				alerta.Code = "E_OPN_02_3"
				alerta.Body = strconv.Itoa(idconceptotitan)
				o.Rollback()
				return
			}
			fmt.Println("Concepto Titan: ", strconv.Itoa(idconceptotitan))
			fmt.Println(Valor)
			fmt.Println("Concepto Kronos:", strconv.Itoa(conceptoKronosHomologado.ConceptoKronos.Id))

			if esta, idlista := estaConcepto(conceptoKronosHomologado.ConceptoKronos.Id, allConceptoOrdenPago); esta == true {
				fmt.Println("---Sumar")
				sumaValor := allConceptoOrdenPago[idlista].Valor + Valor
				allConceptoOrdenPago[idlista].Valor = sumaValor
			} else {
				fmt.Println("---Append")
				newConceptoOrden2 := ConceptoOrdenPago{
					Valor:    Valor,
					Concepto: &ConceptoTesoral{Id: conceptoKronosHomologado.ConceptoKronos.Id},
				}
				allConceptoOrdenPago = append(allConceptoOrdenPago, newConceptoOrden2)
			}
		}
	}
	fmt.Println("\n***************** Totalizado **********************")
	for i := 0; i < len(allConceptoOrdenPago); i++ {
		fmt.Println("\n************ Concepto kronos", strconv.Itoa(allConceptoOrdenPago[i].Concepto.Id), " ************")
		fmt.Println(allConceptoOrdenPago[i].Valor)
		allConceptoOrdenPago[i].OrdenDePago = &OrdenPago{Id: int(idOrdenPago)}
		// insertar concepto_orden_pago
		_, err3 := o.Insert(&allConceptoOrdenPago[i])
		if err3 != nil {
			alerta.Type = "error"
			alerta.Code = "E_OPN_02_4"
			alerta.Body = err3
			err = err3
			o.Rollback()
			return
		}
		// buscamos cuentas contables relacionadas al concepto para registrar movimientos
		qs := o.QueryTable(new(ConceptoTesoralCuentaContable)).RelatedSel()
		qs = qs.Filter("Concepto", allConceptoOrdenPago[i].Concepto.Id)
		qs = qs.RelatedSel()
		var l []ConceptoTesoralCuentaContable
		if _, err = qs.Limit(-1, 0).All(&l); err == nil {
			for _, v := range l {
				//registra movimientos
				fmt.Println("Data para Movimientos")
				newMovimientoContable := MovimientoContable{}
				if v.CuentaContable.Naturaleza == "debito" {
					fmt.Println(v.CuentaContable.Naturaleza)
					newMovimientoContable.Debito = allConceptoOrdenPago[i].Valor
					newMovimientoContable.Credito = 0
				} else {
					fmt.Println(v.CuentaContable.Naturaleza)
					newMovimientoContable.Debito = 0
					newMovimientoContable.Credito = allConceptoOrdenPago[i].Valor
				}
				newMovimientoContable.Fecha = time.Now()
				newMovimientoContable.ConceptoTesoral = v.ConceptoTesoral
				newMovimientoContable.CuentaContable = v.CuentaContable
				newMovimientoContable.TipoDocumentoAfectante = &TipoDocumentoAfectante{Id: 1} //documento afectante tipo op
				newMovimientoContable.CodigoDocumentoAfectante = int(idOrdenPago)
				newMovimientoContable.Aprobado = false
				// insertar OP Planta
				_, err4 := o.Insert(&newMovimientoContable)
				if err4 != nil {
					alerta.Type = "error"
					alerta.Code = "E_OPN_02_5"
					alerta.Body = err4
					err = err4
					o.Rollback()
					return
				}
			}
		}
	}
	fmt.Println("\n*****************FIN Totalizado**********************")
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
func FechaActual(formato string) (fechaActual string, err error) {
	hoy := time.Now()
	fechaActual = hoy.Format(formato)
	return
}
