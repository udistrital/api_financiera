package models

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/udistrital/api_financiera/utilidades"
)

type Rubro struct {
	Id              int      `orm:"auto;column(id);pk"`
	Entidad         *Entidad `orm:"column(entidad);rel(fk)"`
	Codigo          string   `orm:"column(codigo)"`
	Descripcion     string   `orm:"column(descripcion);null"`
	UnidadEjecutora int16    `orm:"column(unidad_ejecutora)"`
	Nombre          string   `orm:"column(nombre);null"`
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
func GetAllRubro(query map[string]string, group []string, fields []string, sortby []string, order []string,
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
	_, err = o.Raw(`SELECT DISTINCT * FROM (SELECT apropiacion.id as Id ,rubro.id as idrubro, rubro.codigo, rubro.descripcion, apropiacion.vigencia, COALESCE( fuente.descripcion , 'Recursos Propios' ) as fdescrip, COALESCE( fuente.Id , 0 ) as idfuente
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

// ApropiacionReporteEgresos informe ordenes de pago y total por orden
func ApropiacionReporteEgresos(inicio time.Time, fin time.Time) (res []interface{}, err error) {
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
			var idfuente interface{}
			if m[i]["idfuente"] == nil {

				idfuente = 0
			} else {

				idfuente = m[i]["idfuente"]
				fmt.Println("fuente ", idfuente)
			}

			egresos, _ := ApropiacionOrdenPago(m[i]["id"], idfuente)
			proy := 0.0 //RubroReporteEgresosProyeccion(inicio, fin, 3, m[i]["idrubro"], m[i]["idfuente"])
			//fmt.Println("sss ", proy)
			aux := make(map[string]interface{})
			//fmt.Println("aux: ", finicio.AddDate(-1, 0, 0))
			if egresos == nil {
				val := make(map[string]interface{})
				val["valor"] = "0"
				val["proyeccion"] = "0"
				val["variacion"] = "0"
				val["pvariacion"] = "0"
				aux["valores"] = val
			} else {
				fll := egresos[0].(map[string]interface{})

				var ejstr string
				err = utilidades.FillStruct(fll["valor"], &ejstr)
				fmt.Println("err ", err)
				ej, err := strconv.ParseFloat(ejstr, 64)
				fmt.Println("err ", err)
				var variacion float64
				var pvariacion float64
				if proy <= 0 {
					variacion = ej - 0
					pvariacion = variacion / ej
				} else {
					variacion = ej - proy
					pvariacion = variacion / ej
				}

				//fmt.Println("vac ", variacion)
				fll["proyeccion"] = proy
				var mp interface{}
				var mpp interface{}
				err = utilidades.FillStruct(variacion, &mp)
				err = utilidades.FillStruct(pvariacion, &mpp)
				fll["variacion"] = mp
				fll["pvariacion"] = mpp
				aux["valores"] = fll

			}
			if aux != nil {
				aux["mes"] = finicio.Format("Jan")
				fechas = append(fechas, aux)
			}

		}
		m[i]["reporte"] = fechas
		//m[i]["egresos"], err = RubroOrdenPago(m[i]["id"])
		if err != nil {
			fmt.Println("err1 ", err)
			return
		}

	}
	err = utilidades.FillStruct(m, &res)
	if err != nil {
		fmt.Println("err2 ", err)
		return
	}
	//fmt.Println("err3 ", err)
	return
}

func RubroReporteEgresosProyeccion(inicio time.Time, fin time.Time, nperiodos int, idrubro interface{}, idfuente interface{}) (res float64, err error) {
	o := orm.NewOrm()
	proy := 0.0
	for i := 1; i <= nperiodos; i++ {
		apropiacion := []Apropiacion{}
		vigencia := int(inicio.AddDate(-i, 0, 0).Year())
		_, err = o.QueryTable("apropiacion").Filter("vigencia", vigencia).Filter("rubro", idrubro).All(&apropiacion)
		fmt.Println("ap ", apropiacion)
		if len(apropiacion) <= 0 {

		} else {
			aux, _ := ApropiacionOrdenPago(apropiacion[0].Id, idfuente)
			if aux != nil {
				for _, m := range aux {
					p := m.(map[string]interface{})
					var val float64
					err = utilidades.FillStruct(p["valor"], &val)
					proy = proy + val
					fmt.Println("proy: ", proy)
				}
			}

		}

	}
	res = proy / float64(nperiodos)

	return
}

func RubroReporteIngresosProyeccion(inicio time.Time, fin time.Time, nperiodos int, idrubro interface{}, idfuente interface{}) (res float64, err error) {
	o := orm.NewOrm()
	proy := 0.0
	for i := 1; i <= nperiodos; i++ {
		apropiacion := []Apropiacion{}
		vigencia := int(inicio.AddDate(-i, 0, 0).Year())
		_, err = o.QueryTable("apropiacion").Filter("vigencia", vigencia).Filter("rubro", idrubro).All(&apropiacion)
		fmt.Println("ap ", apropiacion)
		if len(apropiacion) <= 0 {

		} else {
			aux, _ := RubroIngreso(apropiacion[0].Id, idfuente, inicio, fin)
			if aux != nil {
				for _, m := range aux {
					p := m.(map[string]interface{})
					var val float64
					err = utilidades.FillStruct(p["valor"], &val)
					proy = proy + val
					fmt.Println("proy: ", proy)
				}
			}

		}

	}
	res = proy / float64(nperiodos)

	return
}

// RubroReporteIngresos informe ordenes de pago y total por orden
func ApropiacionReporteIngresos(inicio time.Time, fin time.Time) (res []interface{}, err error) {
	vigencia := int(inicio.Year())
	mesinicio := int(inicio.Month())
	mesfin := int(fin.Month())

	m, err := ListaApropiacionesHijo(vigencia, "2%")
	//fmt.Println("err: ", m)
	if err != nil {
		return
	}
	for i := 0; i < len(m); i++ {
		var fechas []map[string]interface{}
		for j := 0; j <= (mesfin - mesinicio); j++ {
			var ffin time.Time
			//fmt.Println(ffin)
			finicio := inicio.AddDate(0, j, 0)
			if mesfin-mesinicio == 0 || j == mesfin-mesinicio {
				ffin = fin
			} else {
				ffin = inicio.AddDate(0, j+1, 0)
			}
			var idfuente interface{}
			if m[i]["idfuente"] == nil {

				idfuente = 0
			} else {

				idfuente = m[i]["idfuente"]
			}
			ingr, _ := RubroIngreso(m[i]["id"], idfuente, finicio, ffin)
			proy := 0.0 //RubroReporteIngresosProyeccion(inicio, fin, 3, m[i]["idrubro"], m[i]["idfuente"])
			aux := make(map[string]interface{})
			//fmt.Println("aux: ", aux["valores"])
			if ingr == nil {
				val := make(map[string]interface{})
				val["valor"] = "0"
				val["proyeccion"] = "0"
				val["variacion"] = "0"
				val["pvariacion"] = "0"
				aux["valores"] = val
			} else {
				fll := ingr[0].(map[string]interface{})
				var ejstr string
				err = utilidades.FillStruct(fll["valor"], &ejstr)
				//fmt.Println("err ", err)
				ej, err := strconv.ParseFloat(ejstr, 64)
				fmt.Println("err ", err)
				var variacion float64
				var pvariacion float64
				if proy <= 0 {
					variacion = ej - 0
					pvariacion = variacion / ej
				} else {
					variacion = ej - proy
					pvariacion = variacion / ej
				}

				//fmt.Println("vac ", variacion)
				fll["proyeccion"] = proy
				var mp interface{}
				var mpp interface{}
				err = utilidades.FillStruct(variacion, &mp)
				err = utilidades.FillStruct(pvariacion, &mpp)
				fll["variacion"] = mp
				fll["pvariacion"] = mpp
				aux["valores"] = fll

			}
			if aux != nil {
				aux["mes"] = finicio.Format("Jan")
				aux["n_mes"] = int(finicio.Month())
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

// RubroReporte informe ordenes de pago y total por orden
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
			var idfuente interface{}
			if m[i]["idfuente"] == nil {
				fmt.Println("cero")
				idfuente = 0
			} else {
				fmt.Println("no cero")
				idfuente = m[i]["idfuente"]
			}
			ingr, _ := ApropiacionIngreso(m[i]["id"], idfuente, finicio, ffin)

			egresos, _ := ApropiacionOrdenPago(m[i]["id"], idfuente)
			aux := make(map[string]interface{})
			//fmt.Println("aux: ", aux)
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
func ApropiacionOrdenPago(apropiacion interface{}, fuente interface{}) (res []interface{}, err error) {
	o := orm.NewOrm()
	var m []orm.Params
	_, err = o.Raw(`SELECT id_apr,codigo,COALESCE( idfuente, 0 ) as idfuente,SUM(valor) as valor FROM
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
		 COALESCE( idfuente, 0 )  = ?
		GROUP BY
		id_apr,
		  codigo,
			idfuente`, apropiacion, fuente).Values(&m)
	err = utilidades.FillStruct(m, &res)
	return
}

// ApropiacionIngreso informe ingresos
//falta filtro por fechas.
func ApropiacionIngreso(apropiacion interface{}, fuente interface{}, inicio time.Time, fin time.Time) (res []interface{}, err error) {
	o := orm.NewOrm()
	var m []orm.Params
	_, err = o.Raw(`SELECT id_aprop,codigo,COALESCE( idfuente, 0 ) as idfuente, SUM(valor) as valor FROM
(
	SELECT
		ingreso.id ,ingreso.fecha_ingreso, estadoingreso.nombre as estado, estadoingreso.id as id_estado,formaingreso.nombre as forma_ingreso, rubro.codigo as codigo,fuente.id as idfuente, ingresoconcepto.valor_agregado as valor , apropiacion.id as id_aprop
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
		financiera.concepto_tesoral as concepto
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
	LEFT JOIN
		financiera.fuente_financiamiento AS fuente
	ON
		ingreso.fuente_financiamiento = fuente.id
	WHERE
		estadoingreso.nombre = 'Aprobado'
	GROUP BY
		ingreso.id ,fuente.id, ingreso.fecha_ingreso,estadoingreso.nombre, estadoingreso.id,formaingreso.nombre, rubro.codigo , ingresoconcepto.valor_agregado,  apropiacion.id
) AS ingreso
WHERE id_aprop = ?
AND
COALESCE( ingreso.idfuente, 0 )  = ?
AND
ingreso.fecha_ingreso BETWEEN ? AND ?
GROUP BY
id_aprop,
	codigo,
	idfuente`, apropiacion, fuente, inicio, fin).Values(&m)
	err = utilidades.FillStruct(m, &res)
	return
}

// RubroIngreso informe ingresos
//falta filtro por fechas.
func RubroIngreso(rubro interface{}, fuente interface{}, inicio time.Time, fin time.Time) (res []interface{}, err error) {
	o := orm.NewOrm()
	var m []orm.Params
	_, err = o.Raw(`SELECT idrubro,id_aprop,codigo,COALESCE( idfuente, 0 ) as idfuente, SUM(valor) as valor FROM
(
	SELECT
		ingreso.id ,ingreso.fecha_ingreso, estadoingreso.nombre as estado, estadoingreso.id as id_estado,formaingreso.nombre as forma_ingreso, rubro.codigo as codigo,rubro.id as idrubro,fuente.id as idfuente, ingresoconcepto.valor_agregado as valor , apropiacion.id as id_aprop
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
		financiera.concepto_tesoral as concepto
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
	LEFT JOIN
		financiera.fuente_financiamiento AS fuente
	ON
		ingreso.fuente_financiamiento = fuente.id
	WHERE
		estadoingreso.nombre = 'Aprobado'
	GROUP BY
		rubro.id,ingreso.id ,fuente.id, ingreso.fecha_ingreso,estadoingreso.nombre, estadoingreso.id,formaingreso.nombre, rubro.codigo , ingresoconcepto.valor_agregado,  apropiacion.id
) AS ingreso
WHERE idrubro = ?
AND
COALESCE( ingreso.idfuente, 0 )  = ?
AND
ingreso.fecha_ingreso BETWEEN ? AND ?
GROUP BY
id_aprop,
idrubro,
	codigo,
	idfuente`, rubro, fuente, inicio, fin).Values(&m)
	err = utilidades.FillStruct(m, &res)
	return
}

// RubroOrdenPago informe ordenes de pago y total por orden
func RubroOrdenPago(rubro interface{}, fuente interface{}) (res []interface{}, err error) {
	o := orm.NewOrm()
	var m []orm.Params
	_, err = o.Raw(`SELECT idrubro,id_apr,codigo,COALESCE( idfuente, 0 ) as idfuente,SUM(valor) as valor FROM
		(SELECT orden.id , SUM(orden_concepto.valor) as valor , orden.estado_orden_pago , apropiacion.id as id_apr,rubro.id as idrubro,rubro.codigo, fuente.id as idfuente,rp.numero_registro_presupuestal AS RP,
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
			idrubro,apropiacion.rubro, orden.id, rubro.codigo, orden.estado_orden_pago, apropiacion.id, fuente.id, rp.numero_registro_presupuestal, cdp.numero_disponibilidad, fuente.descripcion) as egresos
		WHERE idrubro = ?
		AND
		 COALESCE( idfuente, 0 )  = ?
		GROUP BY
		idrubro,
		id_apr,
		  codigo,
			idfuente`, rubro, fuente).Values(&m)
	err = utilidades.FillStruct(m, &res)
	return
}

//Generar ramas del arbol de rubros
func RamaRubros(done <-chan map[string]interface{}, unidadEjecutora int, forksin <-chan map[string]interface{}) (forksout <-chan map[string]interface{}) {
	out := make(chan map[string]interface{})
	var err error // HLdone
	go func() {   //creacion de gorutines por cada bifurcacion de ramas
		var wg sync.WaitGroup
		for fork := range forksin {
			if fork == nil { //condicion de final de recorrido del arbol.

			} else {
				o := orm.NewOrm()
				var m []orm.Params
				var res []map[string]interface{}
				//funcion para conseguir los hijos de los rubros padre.
				_, err = o.Raw(`SELECT rubro.id as "Id", rubro.codigo as "Codigo",rubro.nombre as "Nombre" ,rubro.descripcion as "Descripcion", rubro.unidad_ejecutora as "UnidadEjecutora"
				  from financiera.rubro
				  join financiera.rubro_rubro
					on  rubro_rubro.rubro_hijo = rubro.id
				  WHERE rubro_rubro.rubro_padre = ?`, fork["Id"]).Values(&m)
				if err == nil {
					err = utilidades.FillStruct(m, &res)
					resch := genChanMapStr(res...)
					var hijos []map[string]interface{}
					wg.Add(1)
					subdone := make(chan map[string]interface{}) // HLdone
					defer close(subdone)
					for hijo := range RamaRubros(subdone, unidadEjecutora, resch) {
						hijos = append(hijos, hijo) //tomar valores del canal y agregarlos al array de hijos.
					}
					fork["Hijos"] = hijos
					select {
					case out <- fork: // HL
					case <-done: // HL
					}
					wg.Done()

				}
			}
		}
		go func() { // HL
			wg.Wait()
			close(out) // HL
		}()
	}()
	return out
}

// Generar arbol de rubros.
func ArbolRubros(unidadEjecutora int) (padres []map[string]interface{}, err error) {
	o := orm.NewOrm()
	var m []orm.Params
	//funcion para conseguir los rubros padre.
	_, err = o.Raw(`  SELECT rubro.id as "Id", rubro.codigo as "Codigo",rubro.nombre as "Nombre" , rubro.descripcion as "Descripcion", rubro.unidad_ejecutora as "UnidadEjecutora"
	    from financiera.rubro
	      where (id  in (select DISTINCT rubro_padre from financiera.rubro_rubro)
			  AND id not in (select DISTINCT rubro_hijo from financiera.rubro_rubro))
			  OR (id not in (select DISTINCT rubro_hijo from financiera.rubro_rubro)
			  AND id not in (select DISTINCT rubro_padre from financiera.rubro_rubro))`).Values(&m)
	if err == nil {
		var res []map[string]interface{}
		err = utilidades.FillStruct(m, &res)
		resch := genChanMapStr(res...)
		done := make(chan map[string]interface{}) // HLdone
		defer close(done)                         // HLdone
		for padre := range RamaRubros(done, unidadEjecutora, resch) {
			padres = append(padres, padre) //tomar valores del canal y agregarlos al array de hijos.
		}
	}
	return
}
