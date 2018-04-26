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

type Inversion struct {
	Id                  int              `orm:"column(id);pk;auto"`
	Vendedor            int              `orm:"column(vendedor)"`
	Emisor              int              `orm:"column(emisor)"`
	NumeroTransaccion   int              `orm:"column(numero_transaccion)"`
	Trm                 float64          `orm:"column(trm)"`
	TasaNominal         float64          `orm:"column(tasa_nominal);null"`
	ValorNominalSaldo   float64          `orm:"column(valor_nominal_saldo);null"`
	ValorNomSaldoMonNal float64          `orm:"column(valor_nom_saldo_mon_nal);null"`
	ValorActual         float64          `orm:"column(valor_actual);null"`
	ValorNetoGirar      float64          `orm:"column(valor_neto_girar);null"`
	FechaCompra         time.Time        `orm:"column(fecha_compra);type(date);null"`
	FechaRedencion      time.Time        `orm:"column(fecha_redencion);type(date);null"`
	FechaVencimiento    time.Time        `orm:"column(fecha_vencimiento);type(date);null"`
	FechaEmision        time.Time        `orm:"column(fecha_emision);type(date);null"`
	Comprador           int              `orm:"column(comprador);null"`
	ValorRecompra       float64          `orm:"column(valor_recompra);null"`
	FechaVenta          time.Time        `orm:"column(fecha_venta);type(date);null"`
	FechaPacto          time.Time        `orm:"column(fecha_pacto);type(date);null"`
	Observaciones       string           `orm:"column(observaciones);null"`
	TituloInversion     *TituloInversion `orm:"column(titulo_inversion);rel(fk)"`
	UnidadEjecutora     *UnidadEjecutora `orm:"column(unidad_ejecutora);rel(fk)"`
	Vigencia            float64          `orm:"column(vigencia);null"`
}

func (t *Inversion) TableName() string {
	return "inversion"
}

func init() {
	orm.RegisterModel(new(Inversion))
}

// AddInversion insert a new Inversion into database and returns
// last inserted Id on success.
func AddInversion(m *Inversion) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetInversionById retrieves Inversion by Id. Returns error if
// Id doesn't exist
func GetInversionById(id int) (v *Inversion, err error) {
	o := orm.NewOrm()
	v = &Inversion{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllInversion retrieves all Inversion matches certain condition. Returns empty list if
// no records exist
func GetAllInversion(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Inversion)).RelatedSel()
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

	var l []Inversion
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

// UpdateInversion updates Inversion by Id and returns error if
// the record to be updated doesn't exist
func UpdateInversionById(m *Inversion) (err error) {
	o := orm.NewOrm()
	v := Inversion{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteInversion deletes Inversion by Id and returns error if
// the record to be deleted doesn't exist
func DeleteInversion(id int) (err error) {
	o := orm.NewOrm()
	v := Inversion{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Inversion{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

// Insert an entire inversion in database returns error if record cant be inserted
func AddInver(request map[string]interface{}) (inversion Inversion, err error) {

	var tipoInversion int
	var usuario int
	var actapadre Inversion
	var inversionCompare Inversion
	var invEstadoInv InversionEstadoInversion
	var concepto Concepto
	var mov []MovimientoContable
	var totalInv float64
	var idInvNueva int64
	var historicoInversiones HistoricoInversion
	o := orm.NewOrm()
	err = formatdata.FillStruct(request["Inversion"], &inversion)
	err = formatdata.FillStruct(request["tipoInversion"], &tipoInversion)
	err = formatdata.FillStruct(request["actapadre"], &actapadre)
	beego.Error(request)
	if err == nil {

		o.Begin()

		idInvNueva, err = o.Insert(&inversion)
		inversion.Id = int(idInvNueva)

		if err == nil {
			if !reflect.DeepEqual(actapadre, inversionCompare) {
				beego.Error("informacion acta padre")
				historicoInversiones.InversionAntigua = &actapadre
				historicoInversiones.InversionNueva = &inversion
				_, err = o.Insert(&historicoInversiones)
				if err != nil {
					beego.Error(err.Error())
					o.Rollback()
					return
				}
			}
		} else {
			beego.Error(err.Error())
			o.Rollback()
			return
		}

		if err == nil {

			err = formatdata.FillStruct(request["usuario"], &usuario)
			err = formatdata.FillStruct(request["EstadoInversion"], &invEstadoInv)
			err = formatdata.FillStruct(request["Concepto"], &concepto)
			err = formatdata.FillStruct(request["Movimientos"], &mov)
			err = formatdata.FillStruct(request["TotalInversion"], &totalInv)

			if err != nil {
				beego.Info(err.Error())
				o.Rollback()
				return
			}

			inversion_concepto := &InversionConcepto{ValorAgregado: totalInv,
				Inversion: &inversion,
				Concepto:  &concepto}

			_, err = o.Insert(inversion_concepto)

			if err != nil {
				beego.Info(err.Error())
				o.Rollback()
				return
			}

			for _, element := range mov {
				element.Fecha = time.Now()
				element.TipoDocumentoAfectante = &TipoDocumentoAfectante{Id: 3}
				element.CodigoDocumentoAfectante = inversion.Id
				element.EstadoMovimientoContable = &EstadoMovimientoContable{Id: 1}
				_, err = o.Insert(&element)

				if err != nil {
					beego.Info(err.Error())
					o.Rollback()
					return
				}
			}
		if err != nil {
				beego.Error(err.Error())
				o.Rollback()
				return
			}
			beego.Error("inserta estado inversion")
			beego.Error(usuario)
			invEstadoInv.Inversion = &inversion
			invEstadoInv.Usuario = usuario
			beego.Error(invEstadoInv)
			_, err = o.Insert(&invEstadoInv)

			if err != nil {
				beego.Error(err.Error())
				o.Rollback()
				return
			}
			o.Commit()
			return
		} else {
			o.Rollback()
			return
		}
	}

	return
}
