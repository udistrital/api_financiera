package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

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

func ListaFuentes() (res interface{}, err error) {
	o := orm.NewOrm()
	var fuentes []*FuenteFinanciamiento
	_, err = o.QueryTable("fuente_financiamiento").All(&fuentes)
	res = fuentes

	return
}

func ListaApropiacionesHijo(vigencia int, codigo string) (res []orm.Params, err error) {
	o := orm.NewOrm()
	//falta realizar proyeccion por cada rubro.
	_, err = o.Raw(`SELECT DISTINCT * FROM (SELECT apropiacion.id , rubro.codigo, rubro.descripcion, apropiacion.vigencia, COALESCE( fuente.descripcion , 'Recursos Propios' ) as fdescrip, fuente.id as idfuente
		FROM
		financiera.apropiacion as apropiacion
	JOIN
		financiera.rubro as rubro
	ON
		rubro.id = apropiacion.rubro
	LEFT JOIN
		financiera.fuente_financiamiento_apropiacion as ffa
	ON
		apropiacion.id = ffa.apropiacion
	LEFT JOIN
		financiera.fuente_financiamiento as fuente
	ON
		fuente.id = ffa.fuente_financiamiento
	WHERE
		rubro.id NOT IN (SELECT DISTINCT rubro_padre FROM financiera.rubro_rubro)
             ) as apropiacion

		WHERE vigencia = ?
		AND codigo LIKE ?
	`, vigencia, codigo).Values(&res)
	return
}

// RubroOrdenPago informe ordenes de pago y total por orden
func RubroReporteEgresos(inicio time.Time, fin time.Time) (res []interface{}, err error) {
	vigencia := int(inicio.Year())
	mesinicio := int(inicio.Month())
	mesfin := int(fin.Month())

	m, err := ListaApropiacionesHijo(vigencia, "3%")
	if err != nil {
		return
	}
	for i := 0; i < len(m); i++ {
		var fechas []map[string]interface{}
		for j := 0; j <= (mesfin - mesinicio); j++ {
			var ffin time.Time
			fmt.Println(ffin)
			finicio := inicio.AddDate(0, j, 0)
			if mesfin-mesinicio == 0 || j == mesfin-mesinicio {
				ffin = fin
			} else {
				ffin = inicio.AddDate(0, j+1, 0)
			}
			egresos, _ := RubroOrdenPago(m[i]["id"], m[i]["idfuente"])
			aux := make(map[string]interface{})
			fmt.Println("aux: ", aux["valores"])
			if egresos == nil {
				val := make(map[string]interface{})
				val["valor"] = "0"
				aux["valores"] = val
			} else {
				aux["valores"] = egresos[0]

			}
			if aux != nil {
				aux["mes"] = finicio.Format("Jan")
				fechas = append(fechas, aux)
			}

		}
		m[i]["reporte"] = fechas
		//m[i]["egresos"], err = RubroOrdenPago(m[i]["id"])
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
func RubroReporteIngresos(inicio time.Time, fin time.Time) (res []interface{}, err error) {
	vigencia := int(inicio.Year())
	mesinicio := int(inicio.Month())
	mesfin := int(fin.Month())

	m, err := ListaApropiacionesHijo(vigencia, "3")
	if err != nil {
		return
	}
	for i := 0; i < len(m); i++ {
		var fechas []map[string]interface{}
		for j := 0; j <= (mesfin - mesinicio); j++ {
			var ffin time.Time
			fmt.Println(ffin)
			finicio := inicio.AddDate(0, j, 0)
			if mesfin-mesinicio == 0 || j == mesfin-mesinicio {
				ffin = fin
			} else {
				ffin = inicio.AddDate(0, j+1, 0)
			}
			ingr, _ := RubroIngreso(m[i]["id"], finicio, ffin)
			aux := make(map[string]interface{})
			fmt.Println("aux: ", aux["valores"])
			if ingr == nil {
				val := make(map[string]interface{})
				val["valor"] = "0"
				aux["valores"] = val
			} else {
				aux["valores"] = ingr[0]

			}
			if aux != nil {
				aux["mes"] = finicio.Format("Jan")
				fechas = append(fechas, aux)
			}

		}
		m[i]["reporte"] = fechas
		//m[i]["egresos"], err = RubroOrdenPago(m[i]["id"])
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
func RubroReporte(inicio time.Time, fin time.Time) (res []interface{}, err error) {
	vigencia := int(inicio.Year())
	mesinicio := int(inicio.Month())
	mesfin := int(fin.Month())

	m, err := ListaApropiacionesHijo(vigencia, "")
	if err != nil {
		return
	}
	for i := 0; i < len(m); i++ {
		var fechas []map[string]interface{}
		for j := 0; j <= (mesfin - mesinicio); j++ {
			var ffin time.Time
			finicio := inicio.AddDate(0, j, 0)
			if mesfin-mesinicio == 0 || j == mesfin-mesinicio {
				ffin = fin
			} else {
				ffin = inicio.AddDate(0, j+1, 0)
			}
			ingr, _ := RubroIngreso(m[i]["id"], finicio, ffin)

			egresos, _ := RubroOrdenPago(m[i]["id"], m[i]["idfuente"])
			aux := make(map[string]interface{})
			fmt.Println("aux: ", aux)
			if ingr == nil {
				/*val := make(map[string]interface{})
				val["valor"] = "0"
				aux["ingresos"] = val*/
			} else {
				aux["ingresos"] = ingr[0]

			}
			if egresos == nil {
				/*val := make(map[string]interface{})
				val["valor"] = "0"
				aux["egresos"] = val*/
			} else {
				aux["egresos"] = egresos[0]

			}
			if aux != nil {
				aux["mes"] = finicio.Format("Jan")
				fechas = append(fechas, aux)
			}

		}
		m[i]["reporte"] = fechas
		//m[i]["egresos"], err = RubroOrdenPago(m[i]["id"])
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
func RubroOrdenPago(apropiacion interface{}, fuente interface{}) (res []interface{}, err error) {
	o := orm.NewOrm()
	var m []orm.Params
	_, err = o.Raw(`SELECT codigo,idfuente,SUM(valor) as valor FROM
		(SELECT orden.id , SUM(orden_concepto.valor) as valor , orden.estado_orden_pago , apropiacion.id as id_apr,rubro.codigo, fuente.id as idfuente,rp.numero_registro_presupuestal AS RP,
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
			financiera.fuente_financiamiento AS fuente
		ON
			disponibilidad.fuente_financiamiento = fuente.id
		GROUP BY
			apropiacion.rubro, orden.id, rubro.codigo, orden.estado_orden_pago, apropiacion.id, fuente.id, rp.numero_registro_presupuestal, cdp.numero_disponibilidad, fuente.descripcion) as rubro
		WHERE id_apr = ?
		AND
		idfuente = ?
		GROUP BY
		  codigo,
			idfuente`, apropiacion, fuente).Values(&m)
	err = utilidades.FillStruct(m, &res)
	return
}

// RubroOrdenPago informe ingresos
//falta filtro por fechas.
func RubroIngreso(apropiacion interface{}, inicio time.Time, fin time.Time) (res []interface{}, err error) {
	o := orm.NewOrm()
	var m []orm.Params
	_, err = o.Raw(`SELECT codigo, SUM(valor) as valor FROM
(
	SELECT
		ingreso.id ,ingreso.fecha_ingreso, estadoingreso.nombre as estado, estadoingreso.id as id_estado,formaingreso.nombre as forma_ingreso, rubro.codigo as codigo, ingresoconcepto.valor_agregado as valor , apropiacion.id as id_aprop
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
		ingreso.id , ingreso.fecha_ingreso,estadoingreso.nombre, estadoingreso.id,formaingreso.nombre, rubro.codigo , ingresoconcepto.valor_agregado,  apropiacion.id
) AS ingreso
WHERE id_aprop = ?
AND
ingreso.fecha_ingreso BETWEEN ? AND ?
GROUP BY
	codigo`, apropiacion, inicio, fin).Values(&m)
	err = utilidades.FillStruct(m, &res)
	return
}
