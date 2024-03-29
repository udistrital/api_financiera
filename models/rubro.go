package models

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/udistrital/utils_oas/formatdata"
	"github.com/udistrital/utils_oas/optimize"
)

type Rubro struct {
	Id              int              `orm:"auto;column(id);pk"`
	Entidad         int              `orm:"column(entidad)"`
	Codigo          string           `orm:"column(codigo)"`
	Descripcion     string           `orm:"column(descripcion);null"`
	UnidadEjecutora int16            `orm:"column(unidad_ejecutora)"`
	Nombre          string           `orm:"column(nombre);null"`
	Concepto        []*Concepto      `orm:"reverse(many)"`
	ProductoRubro   []*ProductoRubro `orm:"reverse(many)"`
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
				o.LoadRelated(&v, "Concepto", 5)
				//o.LoadRelated(&v, "ProductoRubro", 5, 0, 0, "-Activo")
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
	var apropiaciones []int
	var rubrorubro []int
	// ascertain id exists in the database
	o.Begin()
	if err = o.Read(&v); err == nil {
		var num int64
		qb, _ := orm.NewQueryBuilder("mysql")
		qb.Select("id").
			From("financiera.apropiacion").
			Where("rubro=?")
		if _, err = o.Raw(qb.String(), id).QueryRows(&apropiaciones); err == nil {

			if len(apropiaciones) == 0 {
				qb, _ = orm.NewQueryBuilder("mysql")
				qb.Select("id").
					From("financiera.rubro_rubro").
					//Where("rubro_padre=?").
					Where("rubro_hijo=?")
				if _, err = o.Raw(qb.String(), id).QueryRows(&rubrorubro); err == nil {
					for _, idx := range rubrorubro {
						if _, err = o.Delete(&RubroRubro{Id: idx}); err == nil {

						} else {
							o.Rollback()
						}
					}
				}
			} else {
				o.Rollback()
			}

		} else {
			fmt.Println("Error 1 ", err)
			o.Rollback()
		}
		if num, err = o.Delete(&Rubro{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
			o.Commit()
		} else {
			fmt.Println("Error 2 ", err)

			o.Rollback()
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
					err = formatdata.FillStruct(p["valor"], &val)
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
					err = formatdata.FillStruct(p["valor"], &val)
					proy = proy + val
					fmt.Println("proy: ", proy)
				}
			}

		}

	}
	res = proy / float64(nperiodos)

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
	err = formatdata.FillStruct(m, &res)
	if err != nil {
		return
	}
	return
}

func RubroIngresoCierre(inicio time.Time, fin time.Time, codigo string, vigencia int64) (res []interface{}, err error) {
	o := orm.NewOrm()
	var m []orm.Params
	_, err = o.Raw(`SELECT  apropiacion.id as idaprop,
							rubro.id as idrubro,
							rubro.codigo as codigoRub,
							rubro.nombre as descrubro,
							COALESCE(fuente.id,0) as idfuente,
							COALESCE(fuente.descripcion , 'Recursos Propios' ) as fdescrip,
							'Ingresos' as tipo,
							0 as Proyeccion,
							0 as Pvariacion,
							sum(COALESCE(ingresoconcepto.valor_agregado,0)) as Valor,
							0 as Variacion
					FROM financiera.apropiacion as apropiacion
					JOIN financiera.rubro as rubro ON rubro.id = apropiacion.rubro
					LEFT JOIN financiera.concepto_tesoral as concepto ON concepto.rubro = rubro.id
					LEFT JOIN financiera.ingreso_concepto as ingresoconcepto ON ingresoconcepto.concepto = concepto.id
					LEFT JOIN financiera.ingreso as ingreso  ON ingresoconcepto.ingreso = ingreso.id
					left JOIN financiera.ingreso_estado_ingreso estado_ing on estado_ing.ingreso = ingreso.id
												AND estado_ing.fecha_registro BETWEEN ? AND ?
												AND estado_ing.estado_ingreso = 4
					LEFT JOIN financiera.fuente_financiamiento_apropiacion as ffa ON apropiacion.id = ffa.apropiacion
					LEFT JOIN financiera.fuente_financiamiento as fuente ON fuente.id = ffa.fuente_financiamiento
					LEFT JOIN financiera.forma_ingreso as formaingreso ON formaingreso.id = ingreso.forma_ingreso
					WHERE
					rubro.id NOT IN (SELECT DISTINCT rubro_padre FROM financiera.rubro_rubro)
					AND rubro.codigo LIKE ?
					AND  apropiacion.vigencia =?
					GROUP BY
						apropiacion.id,
						rubro.id,
						rubro.codigo,
						idfuente,
						COALESCE(fuente.id,0),
						fuente.descripcion
					ORDER BY apropiacion.id`, inicio, fin, codigo, vigencia).Values(&m)
	err = formatdata.FillStruct(m, &res)
	return
}

func RubroEgresoCierre(inicio time.Time, fin time.Time, codigo string, vigencia int64) (res []interface{}, err error) {
	o := orm.NewOrm()
	var m []orm.Params
	_, err = o.Raw(`SELECT apropiacion.id as idaprop,
	    					rubro.id as idrubro,
	    					rubro.nombre as descrubro,
	    					COALESCE(fuente.id,0) as idfuente,
	    					COALESCE(fuente.descripcion , 'Recursos Propios' ) as fdescrip,
	    					'Egresos' as tipo,
	    					0 as Proyeccion,
	    					0 as Pvariacion,
        					SUM (COALESCE(orden_concepto.valor,0)) as valor ,
            						0 as Variacion
            		From financiera.apropiacion as apropiacion
            		    --ON apropiacion.id = disponibilidad.apropiacion
            		JOIN financiera.rubro as rubro
            		    ON apropiacion.rubro = rubro.id

            		LEFT JOIN financiera.fuente_financiamiento_apropiacion as ffa
							ON apropiacion.id = ffa.apropiacion
	    			LEFT JOIN financiera.fuente_financiamiento as fuente
							ON fuente.id = ffa.fuente_financiamiento
            		LEFT JOIN financiera.disponibilidad_apropiacion AS disp_apr
            		    ON disp_apr.apropiacion = apropiacion.id
            		LEFT JOIN financiera.registro_presupuestal_disponibilidad_apropiacion as rpda
            		    ON rpda.disponibilidad_apropiacion = disp_apr.id
            		LEFT JOIN financiera.concepto_orden_pago as orden_concepto
            		    ON  orden_concepto.registro_presupuestal_disponibilidad_apropiacion = rpda.id
            		LEFT JOIN financiera.orden_pago_estado_orden_pago ep
            		    ON ep.fecha_registro between ? and ?
            		    AND not exists (Select distinct ant.orden_pago
            		        			from  financiera.orden_pago_estado_orden_pago ant
            		        			where ant.orden_pago  = ep.orden_pago
            		            				and ant.estado_orden_pago = 5
            		            				and ant.fecha_registro
            		            				between ? and ?)
            		    AND ep.orden_pago = orden_concepto.orden_de_pago
            		LEFT JOIN financiera.orden_pago as orden
            		    ON ep.orden_pago = orden.id
        			WHERE rubro.id NOT IN (SELECT DISTINCT rubro_padre FROM financiera.rubro_rubro)
            								AND apropiacion.vigencia = ?
            								AND rubro.codigo LIKE ?
            		group by apropiacion.id ,
            				rubro.id ,
            				fuente.id,
            				rubro.codigo
							ORDER BY apropiacion.id`, inicio, fin, inicio, fin, codigo, vigencia).Values(&m)
	err = formatdata.FillStruct(m, &res)
	return
}

func RubroOrdenPagoP(ordenPago int) (res []interface{}, err error) {
	o := orm.NewOrm()
	var m []orm.Params
	_, err = o.Raw(`SELECT rubro.id as idrubro,
	 									COALESCE(ffa.fuente_financiamiento,0) as idfuente,
	 									ep.vigencia,
	 									ep.valor_base
         					From financiera.orden_pago ep
										LEFT JOIN financiera.concepto_orden_pago as orden_concepto
											ON ep.id = orden_concepto.orden_de_pago
										LEFT JOIN financiera.registro_presupuestal_disponibilidad_apropiacion as rpda
											ON orden_concepto.registro_presupuestal_disponibilidad_apropiacion = rpda.id
										LEFT JOIN financiera.disponibilidad_apropiacion AS disp_apr
											ON rpda.disponibilidad_apropiacion = disp_apr.id
										LEFT JOIN financiera.apropiacion as apropiacion
											ON apropiacion.id = disp_apr.apropiacion
										LEFT JOIN financiera.fuente_financiamiento_apropiacion as ffa
											ON apropiacion.id = ffa.apropiacion
										LEFT JOIN financiera.rubro as rubro
											ON apropiacion.rubro = rubro.id
									WHERE ep.id = ?`, ordenPago).Values(&m)
	err = formatdata.FillStruct(m, &res)
	return
}

func ValEjecutadoPac(vigencia int64, mes int, rubro string, fuente string) (res []interface{}, err error) {
	o := orm.NewOrm()
	var m []orm.Params
	fmt.Println("modelos ValEjecutadoPac")
	_, err = o.Raw(`Select valor_ejecutado_mes as valor
					from financiera.detalle_pac detalle
					join financiera.pac pac
    					on detalle.pac = pac.id
					where pac.vigencia = ?
						and detalle.mes = ?
						and detalle.rubro = ?
						and detalle.fuente_financiamiento = ?`, vigencia, mes, rubro, fuente).Values(&m)
	err = formatdata.FillStruct(m, &res)
	return
}

func GetSumbySource(vigencia int, mes int, fuente string, tipo string) (res []interface{}, err error) {
	o := orm.NewOrm()
	var m []orm.Params
	fmt.Println("modelos GetSumbySource")
	_, err = o.Raw(`Select sum(valor_ejecutado_mes) as ejecutado,
									sum(valor_proyectado_mes) as proyectado
										from financiera.detalle_pac detalle
											join financiera.pac pac
    										on detalle.pac = pac.id
											join financiera.rubro
												on rubro.id = detalle.rubro
										where pac.vigencia = ?
												and detalle.mes = ?
												and detalle.fuente_financiamiento = ?
												and rubro.Codigo like ?`, vigencia, mes, fuente, tipo).Values(&m)
	err = formatdata.FillStruct(m, &res)
	return
}

func GetSumbyTotal(vigencia int, mes int, tipo string) (res []interface{}, err error) {
	o := orm.NewOrm()
	var m []orm.Params
	fmt.Println("modelos GetSumbyTotal")
	_, err = o.Raw(`Select sum(valor_ejecutado_mes) as ejecutado,
									sum(valor_proyectado_mes) as proyectado
										from financiera.detalle_pac detalle
											join financiera.pac pac
    										on detalle.pac = pac.id
											join financiera.rubro
												on rubro.id = detalle.rubro
										where pac.vigencia = ?
												and detalle.mes = ?
												and rubro.Codigo like ?`, vigencia, mes, tipo).Values(&m)
	err = formatdata.FillStruct(m, &res)
	return
}

func GetRubroPac(vigencia int, mes int, fuente string, rubro string) (res []interface{}, err error) {
	o := orm.NewOrm()
	var m []orm.Params
	_, err = o.Raw(`Select valor_ejecutado_mes as ejecutado,
										valor_proyectado_mes as proyectado
										from financiera.detalle_pac detalle
											join financiera.pac pac
    										on detalle.pac = pac.id
										where pac.vigencia = ?
												and detalle.mes = ?
												and detalle.fuente_financiamiento = ?
												and detalle.rubro = ?`, vigencia, mes, fuente, rubro).Values(&m)
	err = formatdata.FillStruct(m, &res)
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
	err = formatdata.FillStruct(m, &res)
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
	err = formatdata.FillStruct(m, &res)
	return
}
func RamaRubros(forkin interface{}, params ...interface{}) (forkout interface{}) {
	fork := forkin.(map[string]interface{})
	o := orm.NewOrm()
	var m []orm.Params
	var res []interface{}
	//funcion para conseguir los hijos de los rubros padre.
	_, err := o.Raw(`SELECT rubro.id as "Id", rubro.codigo as "Codigo",rubro.nombre as "Nombre" ,rubro.descripcion as "Descripcion", rubro.unidad_ejecutora as "UnidadEjecutora"
	  from financiera.rubro
	  join financiera.rubro_rubro
		on  rubro_rubro.rubro_hijo = rubro.id
	  WHERE rubro_rubro.rubro_padre = ?
	  AND unidad_ejecutora in (?,0)`, fork["Id"], params).Values(&m)
	if err == nil {
		err = formatdata.FillStruct(m, &res)
		var hijos []map[string]interface{}
		done := make(chan interface{})
		defer close(done)
		resch := optimize.GenChanInterface(res...)
		charbolrubros := optimize.Digest(done, RamaRubros, resch, params)
		for hijo := range charbolrubros {
			if hijo != nil {
				hijos = append(hijos, hijo.(map[string]interface{})) //tomar valores del canal y agregarlos al array de hijos.

			}
		}
		fork["Hijos"] = hijos
		return fork
	}
	return
}

// Generar arbol de rubros.
func ArbolRubros(unidadEjecutora int, CodigoPadre int) (padres []map[string]interface{}, err error) {
	o := orm.NewOrm()
	var m []orm.Params
	searchparam := ""
	if CodigoPadre != 0 {
		searchparam = strconv.Itoa(CodigoPadre)
	}
	searchparam = searchparam + "%"

	//funcion para conseguir los rubros padre. OR (id not in (select DISTINCT rubro_padre from financiera.rubro_rubro))
	_, err = o.Raw(`  SELECT rubro.id as "Id", rubro.codigo as "Codigo",rubro.nombre as "Nombre" , rubro.descripcion as "Descripcion", rubro.unidad_ejecutora as "UnidadEjecutora"
	    from financiera.rubro
	      where (id  in (select DISTINCT rubro_padre from financiera.rubro_rubro)
			  AND id not in (select DISTINCT rubro_hijo from financiera.rubro_rubro))
			  OR (id not in (select DISTINCT rubro_padre from financiera.rubro_rubro)
					AND id not in (select DISTINCT rubro_hijo from financiera.rubro_rubro))
			  AND rubro.codigo LIKE ?
			  AND rubro.unidad_ejecutora IN (?,0)`, searchparam, unidadEjecutora).Values(&m)
	if err == nil {
		var res []interface{}
		err = formatdata.FillStruct(m, &res)
		done := make(chan interface{})
		defer close(done)
		resch := optimize.GenChanInterface(res...)
		var params []interface{}
		params = append(params, unidadEjecutora)
		charbolrubros := optimize.Digest(done, RamaRubros, resch, params)
		for padre := range charbolrubros {
			padres = append(padres, padre.(map[string]interface{})) //tomar valores del canal y agregarlos al array de hijos.
		}
	}
	return
}
