package models

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type RegistroPresupuestal struct {
	Id                                            int                                              `orm:"column(id);pk;auto"`
	UnidadEjecutora                               *UnidadEjecutora                                 `orm:"column(unidad_ejecutora);rel(fk)"`
	Vigencia                                      float64                                          `orm:"column(vigencia)"`
	FechaMovimiento                               time.Time                                        `orm:"column(fecha_movimiento);type(date);null"`
	Responsable                                   int                                              `orm:"column(responsable);null"`
	Estado                                        *EstadoRegistroPresupuestal                      `orm:"column(estado);rel(fk)"`
	NumeroRegistroPresupuestal                    int                                              `orm:"column(numero_registro_presupuestal)"`
	Beneficiario                                  int                                              `orm:"column(beneficiario);null"`
	Compromiso                                    *Compromiso                                      `orm:"column(compromiso);rel(fk)"`
	Solicitud                                     int                                              `orm:"column(solicitud)"`
	RegistroPresupuestalDisponibilidadApropiacion []*RegistroPresupuestalDisponibilidadApropiacion `orm:"reverse(many)"`
}
type DatosRubroRegistroPresupuestal struct {
	Id                 int
	Disponibilidad     *Disponibilidad
	Apropiacion        *Apropiacion
	FuenteFinanciacion *FuenteFinanciamiento
	Valor              float64
	ValorAsignado      float64
}

type DatosRegistroPresupuestal struct { //estructura temporal para el registro con relacion a las apropiaciones
	Rp     *RegistroPresupuestal
	Rubros []DatosRubroRegistroPresupuestal
}
type DatosSaldoRp struct {
	Rp                 *RegistroPresupuestal
	Apropiacion        *Apropiacion
	FuenteFinanciacion *FuenteFinanciamiento
}
type Info_rp_a_anular struct {
	Anulacion      AnulacionRegistroPresupuestal
	Rp_apropiacion []RegistroPresupuestalDisponibilidadApropiacion
	Valor          float64
}

func (t *RegistroPresupuestal) TableName() string {
	return "registro_presupuestal"
}

func init() {
	orm.RegisterModel(new(RegistroPresupuestal))
}

// AddRegistroPresupuestal insert a new RegistroPresupuestal into database and returns
// last inserted Id on success.
func AddRegistoPresupuestal(m *DatosRegistroPresupuestal) (id int64, err error) {
	o := orm.NewOrm()
	o.Begin()
	var consecutivo int
	err = o.Raw(`SELECT COALESCE(MAX(numero_registro_presupuestal), 0)+1  as consecutivo
					FROM financiera.registro_presupuestal WHERE vigencia = ?`, m.Rp.Vigencia).QueryRow(&consecutivo)
	m.Rp.NumeroRegistroPresupuestal = consecutivo
	if err != nil {
		o.Rollback()
		return
	}
	id, err = o.Insert(m.Rp)
	if err == nil {
		m.Rp.Id = int(id)
		for _, data := range m.Rubros {
			registro := RegistroPresupuestalDisponibilidadApropiacion{
				RegistroPresupuestal:      m.Rp,
				DisponibilidadApropiacion: &DisponibilidadApropiacion{Id: data.Id},
				Valor: data.ValorAsignado,
			}
			_, err2 := o.Insert(&registro)
			if err2 != nil {
				o.Rollback()
				return
			}
		}
	} else {
		fmt.Println("error registro rp: ", err.Error())
		o.Rollback()
		return
	}

	o.Commit()
	return
}

// GetRegistroPresupuestalById retrieves RegistroPresupuestal by Id. Returns error if
// Id doesn't exist
func GetRegistroPresupuestalById(id int) (v *RegistroPresupuestal, err error) {
	o := orm.NewOrm()
	v = &RegistroPresupuestal{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllRegistroPresupuestal retrieves all RegistroPresupuestal matches certain condition. Returns empty list if
// no records exist
func GetAllRegistroPresupuestal(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(RegistroPresupuestal))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else if strings.Contains(k, "not_in") {
			k = strings.Replace(k, "__not_in", "", -1)
			qs = qs.Exclude(k, v)
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

	var l []RegistroPresupuestal
	qs = qs.OrderBy(sortFields...).RelatedSel(5)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				o.LoadRelated(&v, "RegistroPresupuestalDisponibilidadApropiacion", 5)
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

// UpdateRegistroPresupuestal updates RegistroPresupuestal by Id and returns error if
// the record to be updated doesn't exist
func UpdateRegistroPresupuestalById(m *RegistroPresupuestal) (err error) {
	o := orm.NewOrm()
	v := RegistroPresupuestal{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteRegistroPresupuestal deletes RegistroPresupuestal by Id and returns error if
// the record to be deleted doesn't exist
func DeleteRegistroPresupuestal(id int) (err error) {
	o := orm.NewOrm()
	v := RegistroPresupuestal{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&RegistroPresupuestal{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

//----------------------------------------
//funcion para obtener saldo restante del cdp
func SaldoRp(id_rp int, id_apropiacion int, id_fuente int) (saldo float64, comprometido float64, anulado float64, err error) {
	/*o := orm.NewOrm()
	var maps []orm.Params
	o.Raw(`SELECT * FROM financiera.saldo_rp WHERE id = ? AND apropiacion = ? `, id_rp, id_apropiacion).Values(&maps)
	fmt.Println("maps: ", maps)
	if maps[0]["valor"] == nil {
		valor = 0
	} else {
		valor, err = strconv.ParseFloat(maps[0]["valor"].(string), 64)
	}*/
	valorrp, err := ValorRp(id_rp, id_apropiacion, id_fuente)
	comprometidorp, err := ComprometidoRp(id_rp, id_apropiacion, id_fuente)
	anuladorp, err := AnuladoRp(id_rp, id_apropiacion, id_fuente)
	anulado = anuladorp
	comprometido = comprometidorp
	saldo = valorrp - anuladorp - comprometidorp
	return
}

//valor original rp.
func ValorRp(id_rp int, id_apropiacion int, id_fuente int) (valor float64, err error) {
	o := orm.NewOrm()
	var maps []orm.Params
	o.Raw(`SELECT * FROM ( SELECT registro_presupuestal.id,
            disponibilidad_apropiacion.apropiacion,
            COALESCE(disponibilidad_apropiacion.fuente_financiamiento, 0) as fuente_financiamiento,
            sum(registro_presupuestal_disponibilidad_apropiacion.valor) AS valor
           FROM financiera.registro_presupuestal_disponibilidad_apropiacion
             JOIN financiera.registro_presupuestal ON registro_presupuestal_disponibilidad_apropiacion.registro_presupuestal = registro_presupuestal.id
             JOIN financiera.disponibilidad_apropiacion ON disponibilidad_apropiacion.id = registro_presupuestal_disponibilidad_apropiacion.disponibilidad_apropiacion
          GROUP BY registro_presupuestal.id, disponibilidad_apropiacion.apropiacion,disponibilidad_apropiacion.fuente_financiamiento) as saldo
          WHERE id = ? AND apropiacion= ? AND fuente_financiamiento = ?;`, id_rp, id_apropiacion, id_fuente).Values(&maps)
	fmt.Println("maps: ", maps)
	if maps == nil {
		valor = 0
	} else {
		valor, err = strconv.ParseFloat(maps[0]["valor"].(string), 64)
	}

	return
}

//valor comprometido del rp.
func ComprometidoRp(id_rp int, id_apropiacion int, id_fuente int) (valor float64, err error) {
	o := orm.NewOrm()
	var maps []orm.Params
	o.Raw(`SELECT * FROM(SELECT registro_presupuestal.id,
            apropiacion.id AS apropiacion,
            COALESCE(disponibilidad_apropiacion.fuente_financiamiento, 0) as fuente_financiamiento,
            sum(concepto_orden_pago.valor) AS valor
           FROM financiera.concepto_orden_pago
             JOIN financiera.registro_presupuestal_disponibilidad_apropiacion ON registro_presupuestal_disponibilidad_apropiacion.id = concepto_orden_pago.registro_presupuestal_disponibilidad_apropiacion
             JOIN financiera.registro_presupuestal ON registro_presupuestal.id = registro_presupuestal_disponibilidad_apropiacion.registro_presupuestal
             JOIN financiera.disponibilidad_apropiacion ON disponibilidad_apropiacion.id = registro_presupuestal_disponibilidad_apropiacion.disponibilidad_apropiacion
             JOIN financiera.apropiacion ON financiera.apropiacion.id = disponibilidad_apropiacion.apropiacion
          GROUP BY registro_presupuestal.id, apropiacion.id, fuente_financiamiento) as comprometido
          WHERE id = ? AND apropiacion= ? AND fuente_financiamiento = ?`, id_rp, id_apropiacion, id_fuente).Values(&maps)
	fmt.Println("maps: ", maps)
	if maps == nil {
		valor = 0
	} else {
		valor, err = strconv.ParseFloat(maps[0]["valor"].(string), 64)
	}

	return
}

//valor anulado del rp por apropiaicon y fuente de financiamiento (si existe)
//valor anulaciones del cdp.
func AnuladoRp(id_rp int, id_apropiacion int, id_fuente int) (valor float64, err error) {
	o := orm.NewOrm()
	var maps []orm.Params
	o.Raw(`SELECT * FROM (SELECT registro_presupuestal.id,
            disponibilidad_apropiacion.apropiacion,
            COALESCE(disponibilidad_apropiacion.fuente_financiamiento,0) as fuente_financiamiento,
            sum(anulacion_registro_presupuestal_disponibilidad_apropiacion.valor) AS valor
           FROM financiera.anulacion_registro_presupuestal_disponibilidad_apropiacion
             JOIN financiera.registro_presupuestal_disponibilidad_apropiacion ON registro_presupuestal_disponibilidad_apropiacion.id = anulacion_registro_presupuestal_disponibilidad_apropiacion.registro_presupuestal_disponibilidad_apropiacion
             JOIN financiera.registro_presupuestal ON registro_presupuestal_disponibilidad_apropiacion.registro_presupuestal = registro_presupuestal.id
             JOIN financiera.disponibilidad_apropiacion ON disponibilidad_apropiacion.id = registro_presupuestal_disponibilidad_apropiacion.disponibilidad_apropiacion
          GROUP BY registro_presupuestal.id, disponibilidad_apropiacion.apropiacion, disponibilidad_apropiacion.fuente_financiamiento) as anulaciones
										WHERE id = ? AND apropiacion = ? AND fuente_financiamiento = ?;`, id_rp, id_apropiacion, id_fuente).Values(&maps)
	if maps == nil {
		valor = 0
	} else {
		valor, err = strconv.ParseFloat(maps[0]["valor"].(string), 64)
	}

	return
}

//----------------------------------------

//----------------------------------------
//funcion para realizar anulacion total en el RP
func AnulacionTotalRp(m *Info_rp_a_anular) (alerta []string, err error) {
	o := orm.NewOrm()
	o.Begin()
	alerta = append(alerta, "success")
	m.Anulacion.FechaRegistro = time.Now()
	id_anulacion_rp, err1 := o.Insert(&m.Anulacion)
	fmt.Println("error")
	if err1 != nil {
		alerta[0] = "error"
		alerta = append(alerta, "No se pudo registrar el detalle de la anulacion")
		err = err1
		o.Rollback()
		return
	}
	var acumRp float64
	acumRp = 0

	for i := 0; i < len(m.Rp_apropiacion); i++ {
		var saldoRp float64
		var err2 error
		if m.Rp_apropiacion[i].DisponibilidadApropiacion.FuenteFinanciamiento != nil {
			saldoRp, _, _, err2 = SaldoRp(m.Rp_apropiacion[i].RegistroPresupuestal.Id, m.Rp_apropiacion[i].DisponibilidadApropiacion.Apropiacion.Id, m.Rp_apropiacion[i].DisponibilidadApropiacion.FuenteFinanciamiento.Id)

		} else {
			saldoRp, _, _, err2 = SaldoRp(m.Rp_apropiacion[i].RegistroPresupuestal.Id, m.Rp_apropiacion[i].DisponibilidadApropiacion.Apropiacion.Id, 0)

		}
		if err2 != nil {
			alerta[0] = "error"
			alerta = append(alerta, "No se pudo cargar el saldo del RP N° "+strconv.Itoa(m.Rp_apropiacion[i].RegistroPresupuestal.NumeroRegistroPresupuestal)+" para la apropiacion del Rubro "+m.Rp_apropiacion[i].DisponibilidadApropiacion.Apropiacion.Rubro.Codigo)
			err = err2
			o.Rollback()
			return
		}
		acumRp = acumRp + saldoRp
		if saldoRp > 0 {
			anulacion_apropiacion := AnulacionRegistroPresupuestalDisponibilidadApropiacion{
				AnulacionRegistroPresupuestal:                 &AnulacionRegistroPresupuestal{Id: int(id_anulacion_rp)},
				RegistroPresupuestalDisponibilidadApropiacion: &m.Rp_apropiacion[i],
				Valor: saldoRp,
			}
			_, err3 := o.Insert(&anulacion_apropiacion)
			if err3 != nil {
				alerta[0] = "error"
				alerta = append(alerta, "No se pudo registrar la anulacion del RP N° "+strconv.Itoa(m.Rp_apropiacion[i].RegistroPresupuestal.NumeroRegistroPresupuestal)+" para la apropiacion del Rubro "+m.Rp_apropiacion[i].DisponibilidadApropiacion.Apropiacion.Rubro.Codigo)
				err = err3
				o.Rollback()
				return
			} else {
				alerta = append(alerta, "se anulo del RP N° "+strconv.Itoa(m.Rp_apropiacion[i].RegistroPresupuestal.NumeroRegistroPresupuestal)+" para la apropiacion del Rubro "+m.Rp_apropiacion[i].DisponibilidadApropiacion.Apropiacion.Rubro.Codigo+" la suma de "+strconv.FormatFloat(saldoRp, 'f', -1, 64))

			}
		} else {
			alerta[0] = "error"
			alerta = append(alerta, "El RP N° "+strconv.Itoa(m.Rp_apropiacion[i].RegistroPresupuestal.NumeroRegistroPresupuestal)+" para la apropiacion del Rubro "+m.Rp_apropiacion[i].DisponibilidadApropiacion.Apropiacion.Rubro.Codigo+" tiene saldo 0")

		}

	}
	if acumRp > 0 {
		m.Rp_apropiacion[0].RegistroPresupuestal.Estado = &EstadoRegistroPresupuestal{Id: 3}
		o.Update(m.Rp_apropiacion[0].RegistroPresupuestal)
		o.Commit()
	} else {
		o.Rollback()
	}

	return
}

//--------------------------------------------------------
//funcion para realizar la anulacion parcial del RP

func AnulacionParcialRp(m *Info_rp_a_anular) (alerta []string, err error) {
	o := orm.NewOrm()
	o.Begin()
	alerta = append(alerta, "success")
	m.Anulacion.FechaRegistro = time.Now()
	id_anulacion_rp, err1 := o.Insert(&m.Anulacion)
	fmt.Println("error")
	if err1 != nil {
		alerta = append(alerta, "No se pudo registrar el detalle de la anulacion")
		alerta[0] = "error"
		err = err1
		o.Rollback()
		return
	}
	var saldoRp float64
	var err2 error
	for i := 0; i < len(m.Rp_apropiacion); i++ {
		if m.Rp_apropiacion[i].DisponibilidadApropiacion.FuenteFinanciamiento != nil {
			saldoRp, _, _, err2 = SaldoRp(m.Rp_apropiacion[i].RegistroPresupuestal.Id, m.Rp_apropiacion[i].DisponibilidadApropiacion.Apropiacion.Id, m.Rp_apropiacion[i].DisponibilidadApropiacion.FuenteFinanciamiento.Id)

		} else {
			saldoRp, _, _, err2 = SaldoRp(m.Rp_apropiacion[i].RegistroPresupuestal.Id, m.Rp_apropiacion[i].DisponibilidadApropiacion.Apropiacion.Id, 0)

		}
		if err2 != nil {
			alerta[0] = "error"
			alerta = append(alerta, "No se pudo cargar el saldo del RP N° "+strconv.Itoa(m.Rp_apropiacion[i].RegistroPresupuestal.NumeroRegistroPresupuestal)+" para la apropiacion del Rubro "+m.Rp_apropiacion[i].DisponibilidadApropiacion.Apropiacion.Rubro.Codigo)
			err = err2
			o.Rollback()
			return
		}
		if saldoRp < m.Valor {
			alerta[0] = "error"
			alerta = append(alerta, "Valor a anular supera el saldo del RP N° "+strconv.Itoa(m.Rp_apropiacion[i].RegistroPresupuestal.NumeroRegistroPresupuestal)+" para la apropiacion del Rubro "+m.Rp_apropiacion[i].DisponibilidadApropiacion.Apropiacion.Rubro.Codigo)
			o.Rollback()
			return
		} else {
			anulacion_apropiacion := AnulacionRegistroPresupuestalDisponibilidadApropiacion{
				AnulacionRegistroPresupuestal:                 &AnulacionRegistroPresupuestal{Id: int(id_anulacion_rp)},
				RegistroPresupuestalDisponibilidadApropiacion: &m.Rp_apropiacion[i],
				Valor: m.Valor,
			}
			_, err3 := o.Insert(&anulacion_apropiacion)
			if err3 != nil {
				alerta[0] = "error"
				alerta = append(alerta, "No se pudo registrar la anulacion del RP N° "+strconv.Itoa(m.Rp_apropiacion[i].RegistroPresupuestal.NumeroRegistroPresupuestal)+" para la apropiacion del Rubro "+m.Rp_apropiacion[i].DisponibilidadApropiacion.Apropiacion.Rubro.Codigo)
				err = err3
				o.Rollback()
				return
			} else {

				alerta = append(alerta, "se anulo del RP N° "+strconv.Itoa(m.Rp_apropiacion[i].RegistroPresupuestal.NumeroRegistroPresupuestal)+" para la apropiacion del Rubro "+m.Rp_apropiacion[i].DisponibilidadApropiacion.Apropiacion.Rubro.Codigo+" la suma de "+strconv.FormatFloat(m.Valor, 'f', -1, 64))

			}
		}

	}
	o.Commit()
	var acumRP float64
	acumRP = 0

	for i := 0; i < len(m.Rp_apropiacion); i++ {
		var saldoRp float64
		if m.Rp_apropiacion[i].DisponibilidadApropiacion.FuenteFinanciamiento != nil {
			saldoRp, err = GetValorActualRp(m.Rp_apropiacion[i].RegistroPresupuestal.Id)

		} else {
			saldoRp, err = GetValorActualRp(m.Rp_apropiacion[i].RegistroPresupuestal.Id)

		}
		if err != nil {
			o.Rollback()
			alerta[0] = "error"
			alerta = append(alerta, "No se pudo registrar la anulacion del RP N° "+strconv.Itoa(m.Rp_apropiacion[i].RegistroPresupuestal.NumeroRegistroPresupuestal)+" para la apropiacion del Rubro "+m.Rp_apropiacion[i].DisponibilidadApropiacion.Apropiacion.Rubro.Codigo)
			fmt.Println("entro: ", saldoRp)
			return
		}
		fmt.Println("saldo: ", saldoRp)
		acumRP = acumRP + saldoRp
	}
	if acumRP == 0 {
		m.Rp_apropiacion[0].RegistroPresupuestal.Estado = &EstadoRegistroPresupuestal{Id: 3}
		o.Update(m.Rp_apropiacion[0].RegistroPresupuestal)

	}

	return
}

//funcion GetValorTotalRp
func GetValorTotalRp(rp_id int) (total float64, err error) {
	o := orm.NewOrm()
	var totalSql float64
	err = o.Raw("select sum(valor) from financiera.registro_presupuestal_disponibilidad_apropiacion where registro_presupuestal = ?", rp_id).QueryRow(&totalSql)
	if err == nil {
		return totalSql, nil
	}
	return 0, nil
}
func GetValorTotalComprometidoRp(rp_id int) (total float64, err error) {
	o := orm.NewOrm()
	var totalSql float64
	err = o.Raw(`SELECT valor FROM(SELECT registro_presupuestal.id,
            sum(concepto_orden_pago.valor) AS valor
           FROM financiera.orden_pago
             JOIN financiera.concepto_orden_pago ON concepto_orden_pago.orden_de_pago = orden_pago.id
             JOIN financiera.concepto ON concepto.id = concepto_orden_pago.concepto
             JOIN financiera.rubro ON concepto.rubro = rubro.id
             JOIN financiera.registro_presupuestal ON registro_presupuestal.id = orden_pago.registro_presupuestal
             JOIN financiera.apropiacion ON apropiacion.rubro = rubro.id AND apropiacion.vigencia = registro_presupuestal.vigencia
          GROUP BY registro_presupuestal.id) as comprometido
					WHERE id = ?`, rp_id).QueryRow(&totalSql)
	if err == nil {
		fmt.Println("total: ", totalSql)
		return totalSql, nil
	}
	return 0, nil

}

func GetValorTotalAnuladoRp(rp_id int) (total float64, err error) {
	o := orm.NewOrm()
	var totalSql float64
	err = o.Raw(`SELECT valor FROM (SELECT registro_presupuestal.id,
            sum(anulacion_registro_presupuestal_disponibilidad_apropiacion.valor) AS valor
           FROM financiera.anulacion_registro_presupuestal_disponibilidad_apropiacion
             JOIN financiera.registro_presupuestal_disponibilidad_apropiacion ON registro_presupuestal_disponibilidad_apropiacion.id = anulacion_registro_presupuestal_disponibilidad_apropiacion.registro_presupuestal_disponibilidad_apropiacion
             JOIN financiera.registro_presupuestal ON registro_presupuestal_disponibilidad_apropiacion.registro_presupuestal = registro_presupuestal.id
             JOIN financiera.disponibilidad_apropiacion ON disponibilidad_apropiacion.id = registro_presupuestal_disponibilidad_apropiacion.disponibilidad_apropiacion
          GROUP BY registro_presupuestal.id) as anulaciones
					WHERE id = ?`, rp_id).QueryRow(&totalSql)
	if err == nil {
		fmt.Println("total A: ", totalSql)
		return totalSql, nil
	}
	fmt.Println("total A: ", err)
	return 0, nil

}

func GetValorActualRp(rp_id int) (total float64, err error) {
	valor, err := GetValorTotalRp(rp_id)
	comprometido, err := GetValorTotalComprometidoRp(rp_id)
	anulado, err := GetValorTotalAnuladoRp(rp_id)
	total = valor - comprometido - anulado
	return
}
