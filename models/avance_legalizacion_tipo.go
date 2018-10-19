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

type AvanceLegalizacionTipo struct {
	Id                           int                           `orm:"column(id);pk;auto"`
	TipoAvanceLegalizacion       *TipoAvanceLegalizacion       `orm:"column(tipo_avance_legalizacion);rel(fk)"`
	AvanceLegalizacion           *AvanceLegalizacion           `orm:"column(avance_legalizacion);rel(fk)"`
	Tercero                      string                        `orm:"column(tercero)"`
	FechaCompra                  time.Time                     `orm:"column(fecha_compra);type(date);null"`
	FechaCambioDivisa            time.Time                     `orm:"column(fecha_cambio_divisa);type(date);null"`
	Dias                         int                           `orm:"column(dias);null"`
	TrmFechaCompra               float64                       `orm:"column(trm_fecha_compra);null"`
	NumeroFactura                string                        `orm:"column(numero_factura);null"`
	Subtipo                      *AvanceLegalizacionSubTipo    `orm:"column(subtipo);rel(fk);null"`
	EntradaAlmacen               int                           `orm:"column(entrada_almacen)"`
	TipoDocumentoAfectante       *TipoDocumentoAfectante       `orm:"column(tipo_documento_afectante);rel(fk)"`
	EstadoAvanceLegalizacionTipo *EstadoAvanceLegalizacionTipo `orm:"column(estado);rel(fk)"`
}

func (t *AvanceLegalizacionTipo) TableName() string {
	return "avance_legalizacion_tipo"
}

func init() {
	orm.RegisterModel(new(AvanceLegalizacionTipo))
}

// AddAvanceLegalizacionTipo insert a new AvanceLegalizacionTipo into database and returns
// last inserted Id on success.
func AddAvanceLegalizacionTipo(m *AvanceLegalizacionTipo) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetAvanceLegalizacionTipoById retrieves AvanceLegalizacionTipo by Id. Returns error if
// Id doesn't exist
func GetAvanceLegalizacionTipoById(id int) (v *AvanceLegalizacionTipo, err error) {
	o := orm.NewOrm()
	v = &AvanceLegalizacionTipo{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllAvanceLegalizacionTipo retrieves all AvanceLegalizacionTipo matches certain condition. Returns empty list if
// no records exist
func GetAllAvanceLegalizacionTipo(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(AvanceLegalizacionTipo))
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

	var l []AvanceLegalizacionTipo
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				o.LoadRelated(&v, "TipoAvanceLegalizacion")
				o.LoadRelated(&v, "EstadoAvanceLegalizacionTipo")
				o.LoadRelated(&v, "Subtipo")
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

// UpdateAvanceLegalizacionTipo updates AvanceLegalizacionTipo by Id and returns error if
// the record to be updated doesn't exist
func UpdateAvanceLegalizacionTipoById(m *AvanceLegalizacionTipo) (err error) {
	o := orm.NewOrm()
	v := AvanceLegalizacionTipo{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteAvanceLegalizacionTipo deletes AvanceLegalizacionTipo by Id and returns error if
// the record to be deleted doesn't exist
func DeleteAvanceLegalizacionTipo(id int) (err error) {
	o := orm.NewOrm()
	v := AvanceLegalizacionTipo{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&AvanceLegalizacionTipo{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

// AddAllAvanceLegalizacionTipo insert a new AvanceLegalizacionTipo with all related information about it
// into database and returns
// last inserted Id on success.
func AddAllAvanceLegalizacionTipo(m map[string]interface{}) (avanceLegalizacionTipo AvanceLegalizacionTipo, err error) {

	var conceptoLegAvanceTipo ConceptoAvanceLegalizacionTipo
	var movimientosContables []MovimientoContable
	var solicitudAvance SolicitudAvance
	var legalizacionAvance AvanceLegalizacion
	var tipoDocAfectante TipoDocumentoAfectante
	var estadoMov EstadoMovimientoContable
	var idLegA int64
	var idavanceLegT int64
	var concepto Concepto
	var consecLeg int64
	var valorLeg float64
	var tipoDocumentoAfectante int64
	var estadoLegalizacionAvanceTipo EstadoAvanceLegalizacionTipo

	o := orm.NewOrm()

	err = formatdata.FillStruct(m["AvanceLegalizacionTipo"], &avanceLegalizacionTipo)
	err = formatdata.FillStruct(m["Movimientos"], &movimientosContables)
	err = formatdata.FillStruct(m["Avance"], &solicitudAvance)
	err = formatdata.FillStruct(m["Concepto"], &concepto)
	err = formatdata.FillStruct(m["Valor"], &valorLeg)
	err = formatdata.FillStruct(m["TipoDocAfectanteNO"], &tipoDocumentoAfectante)

	if err != nil {
		beego.Error(err.Error())
		return
	}
	conceptoLegAvanceTipo.Concepto = &concepto
	conceptoLegAvanceTipo.Valor = valorLeg
	o.Begin()

	err = o.QueryTable("avance_legalizacion").
		Filter("avance", solicitudAvance.Id).
		One(&legalizacionAvance)
	if err != nil {
		if err == orm.ErrNoRows {
			qb, _ := orm.NewQueryBuilder("mysql")
			qb.Select("COALESCE(MAX(al.legalizacion),0)+1").
				From("avance_legalizacion al")
			sql := qb.String()
			err = o.Raw(sql).QueryRow(&consecLeg)
			legalizacionAvance.Legalizacion = int(consecLeg)
			legalizacionAvance.Avance = &solicitudAvance
			idLegA, err = o.Insert(&legalizacionAvance)
			if err != nil {
				beego.Error(err.Error())
				o.Rollback()
				return
			}
			legalizacionAvance.Id = int(idLegA)
		} else {
			beego.Error("Error", err)
			o.Rollback()
			return
		}
	}
	err = o.QueryTable("tipo_documento_afectante").
		Filter("numeroOrden", tipoDocumentoAfectante).
		One(&tipoDocAfectante)
	if err != nil {
		beego.Error(err.Error())
		o.Rollback()
		return
	}

	avanceLegalizacionTipo.TipoDocumentoAfectante = &tipoDocAfectante

	avanceLegalizacionTipo.AvanceLegalizacion = &legalizacionAvance

	err = o.QueryTable("estado_avance_legalizacion_tipo").
		Filter("numeroOrden", 1).
		One(&estadoLegalizacionAvanceTipo)
	if err != nil {
		beego.Error(err.Error())
		o.Rollback()
		return
	}
	avanceLegalizacionTipo.EstadoAvanceLegalizacionTipo = &estadoLegalizacionAvanceTipo

	idavanceLegT, err = o.Insert(&avanceLegalizacionTipo)
	if err != nil {
		beego.Error(err.Error())
		o.Rollback()
		return
	} else {
		avanceLegalizacionTipo.Id = int(idavanceLegT)
	}
	conceptoLegAvanceTipo.AvanceLegalizacion = &avanceLegalizacionTipo
	_, err = o.Insert(&conceptoLegAvanceTipo)
	if err != nil {
		beego.Error(err.Error())
		o.Rollback()
		return
	}
	err = o.QueryTable("estado_movimiento_contable").
		Filter("numeroOrden", 1).
		One(&estadoMov)
	if err != nil {
		beego.Error(err.Error())
		o.Rollback()
		return
	}
	for i := 0; i < len(movimientosContables); i++ {
		movimientosContables[i].CodigoDocumentoAfectante = int(idavanceLegT)
		movimientosContables[i].TipoDocumentoAfectante = &tipoDocAfectante
		movimientosContables[i].Fecha = time.Now()
		movimientosContables[i].EstadoMovimientoContable = &estadoMov
	}
	_, err = o.InsertMulti(100, movimientosContables)
	if err != nil {
		beego.Error(err.Error())
		o.Rollback()
		return
	}
	beego.Error("avance legalziacion tipo", avanceLegalizacionTipo)
	o.Commit()
	return
}

//get the total legalized from a advance
func GetLegalizationValue(id int) (value float64, err error) {
	var midValue float64
	o := orm.NewOrm()

	qb, _ := orm.NewQueryBuilder("mysql")

	qb.Select("COALESCE(sum(cavt.valor),0)").
		From("avance_legalizacion_tipo avt").
		InnerJoin("concepto_avance_legalizacion_tipo cavt").On("cavt.avance_legalizacion = avt.id").
		Where("avt.avance_legalizacion = ?").And("avt.estado = 1")

	sql := qb.String()

	err = o.Raw(sql, id).QueryRow(&midValue)

	value = midValue

	qb, _ = orm.NewQueryBuilder("mysql")
	qb.Select("COALESCE(sum(ic.valor_agregado),0)").
		From("reintegro_avance_legalizacion rav").
		InnerJoin("reintegro r").On("rav.reintegro = r.id").
		InnerJoin("ingreso i").On("r.ingreso = i.id").
		InnerJoin("ingreso_concepto ic").On("ic.ingreso = i.id").
		Where("rav.avance_legalizacion = ?")

	sql = qb.String()

	err = o.Raw(sql, id).QueryRow(&midValue)

	value += midValue

	return
}

//get taxes from legalizacioTIpo according to movimientosContables
func GetTaxesLegalization(documento int, noTipoDocumento int) (response []interface{}, err error) {
	var movimientosContables []MovimientoContable
	var tipoDocAfectante TipoDocumentoAfectante
	o := orm.NewOrm()

	err = o.QueryTable("tipo_documento_afectante").
		Filter("numeroOrden", noTipoDocumento).
		One(&tipoDocAfectante)
	if err != nil {
		beego.Error("Error", err)
		return
	}

	qs := o.QueryTable(new(MovimientoContable))
	qs = qs.Filter("TipoDocumentoAfectante", tipoDocAfectante.Id)
	qs = qs.Filter("CodigoDocumentoAfectante", documento)
	qs = qs.Filter("CuentaEspecial__isnull", false)
	_, err = qs.All(&movimientosContables)
	if err == nil {
		for _, v := range movimientosContables {
			o.LoadRelated(&v, "CuentaEspecial", 5)
			val := reflect.ValueOf(v)
			m := val.FieldByName("CuentaEspecial").Interface()
			response = append(response, m)
		}
	} else {
		beego.Error("Error ", err)
		return nil, err
	}
	return response, nil
}

//get movs from legalizacioTipo according to movimientosContables
func GetMovsLegalization(documento int, noTipoDocumento int) (response []interface{}, err error) {
	var movimientosContables []MovimientoContable
	var tipoDocAfectante TipoDocumentoAfectante
	o := orm.NewOrm()

	err = o.QueryTable("tipo_documento_afectante").
		Filter("numeroOrden", noTipoDocumento).
		One(&tipoDocAfectante)
	if err != nil {
		beego.Error("Error", err)
		return
	}
	beego.Error("tipo doc afectante ", tipoDocAfectante.Id, "Codigo doc afectante ", documento)
	qs := o.QueryTable(new(MovimientoContable))
	qs = qs.Filter("TipoDocumentoAfectante", tipoDocAfectante.Id)
	qs = qs.Filter("CodigoDocumentoAfectante", documento)
	_, err = qs.All(&movimientosContables)
	for _, v := range movimientosContables {
		o.LoadRelated(&v, "CuentaContable")
		val := reflect.ValueOf(v).Interface()
		response = append(response, val)
	}
	if err != nil {
		beego.Error("Error", err)
		return
	}
	return response, nil
}
