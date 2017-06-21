package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
	"github.com/udistrital/api_financiera/utilidades"
)

type Rubro struct {
	Id             int      `orm:"auto;column(id);pk"`
	Entidad        *Entidad `orm:"column(entidad);rel(fk)"`
	Codigo         string   `orm:"column(codigo)"`
	Vigencia       float64  `orm:"column(vigencia)"`
	Descripcion    string   `orm:"column(descripcion);null"`
	TipoPlan       int16    `orm:"column(tipo_plan);null"`
	Administracion string   `orm:"column(administracion);null"`
	Estado         int16    `orm:"column(estado);null"`
}

func (t *Rubro) TableName() string {
	return "rubro"
}

func init() {
	orm.RegisterModel(new(Rubro))
}

// AddRubro insert a new Rubro into database and returns
// last inserted Id on success.
func AddRubro(m *Rubro) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetRubroById retrieves Rubro by Id. Returns error if
// Id doesn't exist
func GetRubroById(id int) (v *Rubro, err error) {
	o := orm.NewOrm()
	v = &Rubro{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllRubro retrieves all Rubro matches certain condition. Returns empty list if
// no records exist
func GetAllRubro(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Rubro))
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

	var l []Rubro
	qs = qs.OrderBy(sortFields...).RelatedSel(5)
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

// UpdateRubro updates Rubro by Id and returns error if
// the record to be updated doesn't exist
func UpdateRubroById(m *Rubro) (err error) {
	o := orm.NewOrm()
	v := Rubro{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteRubro deletes Rubro by Id and returns error if
// the record to be deleted doesn't exist
func DeleteRubro(id int) (err error) {
	o := orm.NewOrm()
	v := Rubro{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Rubro{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func ListaApropiacionesHijo(vigencia int) (res []orm.Params, err error) {
	o := orm.NewOrm()
	//falta realizar proyeccion por cada rubro.
	_, err = o.Raw(`SELECT* FROM (SELECT apropiacion.id , rubro.codigo, apropiacion.vigencia, apropiacion.valor
		FROM
		financiera.apropiacion as apropiacion
	JOIN
		financiera.rubro as rubro
	ON
		rubro.id = apropiacion.rubro
	JOIN
		financiera.rubro_rubro as gerarquia
	ON
		gerarquia.rubro_hijo = rubro.id
	AND
		rubro.id NOT IN (SELECT rubro_padre FROM financiera.rubro_rubro)) as apropiacion
		WHERE vigencia = ?`, vigencia).Values(&res)
	return
}

// RubroOrdenPago informe ordenes de pago y total por orden
func RubroReporte(vigencia int) (res []interface{}, err error) {
	m, err := ListaApropiacionesHijo(vigencia)
	if err != nil {
		return
	}
	for i := 0; i < len(m); i++ {
		m[i]["egresos"], err = RubroOrdenPago(m[i]["id"])
		if err != nil {
			return
		}
		m[i]["ingresos"], err = RubroIngreso(m[i]["id"])
		if err != nil {
			return
		}
	}
	err = utilidades.FillStruct(m, &res)
	if err != nil {
		return
	}
	return
}

// RubroOrdenPago informe ordenes de pago y total por orden
func RubroOrdenPago(apropiacion interface{}) (res []interface{}, err error) {
	o := orm.NewOrm()
	var m []orm.Params
	_, err = o.Raw(`SELECT * FROM
		(SELECT orden.id , SUM(orden_concepto.valor) as valor , orden.estado_orden_pago , apropiacion.id as id_apr,rubro.codigo, rp.numero_registro_presupuestal AS RP,
		cdp.numero_disponibilidad AS CDP, fuente.descripcion AS fuente

		FROM
			financiera.orden_pago as orden
		JOIN
			financiera.concepto_orden_pago as orden_concepto
		ON
			orden_concepto.orden_de_pago = orden.id
		JOIN
			financiera.registro_presupuestal_disponibilidad_apropiacion as rpda
		ON
			rpda.id = orden_concepto.registro_presupuestal_disponibilidad_apropiacion
		JOIN
			financiera.disponibilidad_apropiacion as disponibilidad
		ON
			disponibilidad.id = rpda.disponibilidad_apropiacion
		JOIN
			financiera.apropiacion as apropiacion
		ON
			apropiacion.id = disponibilidad.apropiacion
		JOIN
			financiera.rubro as rubro
		ON      apropiacion.rubro = rubro.id
		JOIN
			financiera.estado_orden_pago as estado_ord
		ON
			estado_ord.id = orden.estado_orden_pago
		JOIN
			financiera.registro_presupuestal as rp
		ON
			rp.id = rpda.registro_presupuestal
		JOIN
			financiera.disponibilidad_apropiacion AS disp_apr
		ON
		  disp_apr.id = rpda.disponibilidad_apropiacion
		JOIN
			financiera.disponibilidad as cdp
		ON
			cdp.id = disp_apr.disponibilidad
		LEFT JOIN
			financiera.fuente_financiacion AS fuente
		ON
			disponibilidad.fuente_financiamiento = fuente.id
		GROUP BY
			apropiacion.rubro, orden.id, rubro.codigo, orden.estado_orden_pago, apropiacion.id, rp.numero_registro_presupuestal, cdp.numero_disponibilidad, fuente.descripcion) as rubro
		WHERE id_apr = ?`, apropiacion).Values(&m)
	err = utilidades.FillStruct(m, &res)
	return
}

// RubroOrdenPago informe ingresos
//falta filtro por fechas.
func RubroIngreso(apropiacion interface{}) (res []interface{}, err error) {
	o := orm.NewOrm()
	var m []orm.Params
	_, err = o.Raw(`SELECT * FROM
(
	SELECT
		ingreso.id , estadoingreso.nombre as estado, estadoingreso.id as id_estado,formaingreso.nombre as forma_ingreso, rubro.codigo as rubro, ingresoconcepto.valor_agregado as valor , apropiacion.id as id_aprop
	FROM
		financiera.ingreso as ingreso
	JOIN
		financiera.estado_ingreso as estadoingreso
	ON
		ingreso.estado_ingreso = estadoingreso.id
	JOIN
		financiera.forma_ingreso as formaingreso
	ON
		formaingreso.id = ingreso.forma_ingreso
	JOIN
		financiera.ingreso_concepto as ingresoconcepto
	ON
		ingresoconcepto.ingreso = ingreso.id
	JOIN
		financiera.concepto as concepto
	ON
		concepto.id = ingresoconcepto.concepto
	JOIN
		financiera.rubro as rubro
	ON
		rubro.id = concepto.rubro
	JOIN
		financiera.apropiacion as apropiacion
	ON
		apropiacion.rubro = rubro.id AND apropiacion.vigencia = ingreso.vigencia
	JOIN
		financiera.movimiento_contable as mov
	ON
		mov.tipo_documento_afectante = 2 AND mov.codigo_documento_afectante = ingreso.id
	GROUP BY
		ingreso.id , estadoingreso.nombre, estadoingreso.id,formaingreso.nombre, rubro.codigo, ingresoconcepto.valor_agregado,  apropiacion.id
) AS ingreso
WHERE id_aprop = ?`, apropiacion).Values(&m)
	err = utilidades.FillStruct(m, &res)
	return
}
