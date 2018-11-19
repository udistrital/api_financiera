package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/udistrital/utils_oas/formatdata"
)

type Apropiacion struct {
	Id       int                `orm:"column(id);pk;auto"`
	Vigencia float64            `orm:"column(vigencia)"`
	Rubro    *Rubro             `orm:"column(rubro);rel(fk)"`
	Valor    float64            `orm:"column(valor)"`
	Estado   *EstadoApropiacion `orm:"column(estado);rel(fk)"`
}

func (t *Apropiacion) TableName() string {
	return "apropiacion"
}

func init() {
	orm.RegisterModel(new(Apropiacion))
}

// AddApropiacion insert a new Apropiacion into database and returns
// last inserted Id on success.
func AddApropiacion(m *Apropiacion) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetApropiacionById retrieves Apropiacion by Id. Returns error if
// Id doesn't exist
func GetApropiacionById(id int) (v *Apropiacion, err error) {
	o := orm.NewOrm()
	v = &Apropiacion{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllApropiacion retrieves all Apropiacion matches certain condition. Returns empty list if
// no records exist
func GetAllApropiacion(query map[string]string, exclude map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Apropiacion))
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

	// exclude k=v
	for k, v := range exclude {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Exclude(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Exclude(k, v)
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

	var l []Apropiacion
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

// UpdateApropiacion updates Apropiacion by Id and returns error if
// the record to be updated doesn't exist
func UpdateApropiacionById(m *Apropiacion) (err error) {
	o := orm.NewOrm()
	v := Apropiacion{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteApropiacion deletes Apropiacion by Id and returns error if
// the record to be deleted doesn't exist
func DeleteApropiacion(id int) (err error) {
	o := orm.NewOrm()
	v := Apropiacion{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Apropiacion{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

//funcion para comprobar saldo de la apropiacion de un Rubro

func SaldoApropiacion(Id int) (saldo map[string]float64, err error) {
	var valor float64
	saldo = make(map[string]float64)
	valorapr, err := ValorApropiacion(Id)
	if err != nil {
		return
	}
	valorcdpapr, err := ValorCdpPorApropiacion(Id)
	if err != nil {
		return
	}
	valoranuladocdpapr, err := ValorAnuladoCdpPorApropiacion(Id)
	if err != nil {
		return
	}
	valorAdiciones, err := ValorMovimientosPorApropiacion(Id, 3, "cuenta_credito")
	if err != nil {
		return
	}
	valorAdicionesTraslados, err := ValorMovimientosPorApropiacion(Id, 1, "cuenta_credito")
	if err != nil {
		return
	}
	valorReducciones, err := ValorMovimientosPorApropiacion(Id, 2, "cuenta_credito")
	if err != nil {
		return
	}
	valorReduccionesTraslados, err := ValorMovimientosPorApropiacion(Id, 1, "cuenta_contra_credito")
	if err != nil {
		return
	}
	valor = valorapr - valorcdpapr + valoranuladocdpapr + valorAdiciones + valorAdicionesTraslados
	saldo["original"] = valorapr
	saldo["saldo"] = valor
	saldo["comprometido"] = valorcdpapr - valoranuladocdpapr
	saldo["adiciones"] = valorAdiciones
	saldo["traslados"] = valorAdicionesTraslados
	saldo["reducciones"] = valorReducciones + valorReduccionesTraslados
	saldo["comprometido_anulado"] = valoranuladocdpapr
	return
}

func VigenciaApropiacion() (ml []int, err error) {
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("DISTINCT vigencia").
		From("financiera.apropiacion")

	sql := qb.String()
	o := orm.NewOrm()
	o.Raw(sql).QueryRows(&ml)

	if len(ml) == 0 {
		return nil, err
	}
	return ml, nil
}

//funcion para determinar el valor con traslados de la apropiacion
func ValorApropiacion(Id int) (valor float64, err error) {
	o := orm.NewOrm()
	var maps_valor_tot []orm.Params
	_, err = o.Raw(`SELECT valor
				FROM financiera.apropiacion
				WHERE id= ?`, Id).Values(&maps_valor_tot)
	//fmt.Println("maps: ", len(maps_valor_tot))
	if len(maps_valor_tot) > 0 && err == nil {
		valor, _ = strconv.ParseFloat(maps_valor_tot[0]["valor"].(string), 64)
	} else {
		valor = 0
	}

	return
}

//funcion para determinar el total del valor de los cdp hechos a una apropiacion
func ValorMovimientosPorApropiacion(Id int, tipoMov int, cuenta string) (valor float64, err error) {
	o := orm.NewOrm()
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("COALESCE(sum(valor),0) as valor").
		From("financiera.movimiento_apropiacion_disponibilidad_apropiacion").
		InnerJoin("financiera.movimiento_apropiacion").
		On("movimiento_apropiacion.id = movimiento_apropiacion_disponibilidad_apropiacion.movimiento_apropiacion").
		Where(cuenta + " = ?").
		And("tipo_movimiento_apropiacion = ?").
		And("estado_movimiento_apropiacion = 2")
	err = o.Raw(qb.String(), Id, tipoMov).QueryRow(&valor)
	return
}

//funcion para determinar el total del valor de los cdp hechos a una apropiacion
func ValorCdpPorApropiacion(Id int) (valor float64, err error) {
	o := orm.NewOrm()
	var maps_valor_tot []orm.Params
	_, err = o.Raw(`SELECT  disponibilidad_apropiacion.apropiacion,
		COALESCE(sum(disponibilidad_apropiacion.valor),0) AS valor
	   FROM financiera.disponibilidad
		 JOIN financiera.disponibilidad_apropiacion ON disponibilidad_apropiacion.disponibilidad = disponibilidad.id
		 WHERE apropiacion= ?
		 GROUP BY disponibilidad_apropiacion.apropiacion
				`, Id).Values(&maps_valor_tot)
	//fmt.Println("maps: ", len(maps_valor_tot))
	if len(maps_valor_tot) > 0 && err == nil {
		valor, _ = strconv.ParseFloat(maps_valor_tot[0]["valor"].(string), 64)
	} else {
		valor = 0
	}

	return
}

//funcion para determinar el total del valor de los cdp hechos a una apropiacion
func ValorAnuladoCdpPorApropiacion(Id int) (valor float64, err error) {
	o := orm.NewOrm()
	var maps_valor_tot []orm.Params
	_, err = o.Raw(`SELECT anulacion_disponibilidad.estado_anulacion,
								disponibilidad_apropiacion.apropiacion,
								COALESCE(sum(anulacion_disponibilidad_apropiacion.valor),0) AS valor
	   						FROM financiera.anulacion_disponibilidad_apropiacion
		 					JOIN financiera.disponibilidad_apropiacion ON anulacion_disponibilidad_apropiacion.disponibilidad_apropiacion = disponibilidad_apropiacion.id
		 					JOIN financiera.disponibilidad ON disponibilidad_apropiacion.disponibilidad = disponibilidad.id
					 		JOIN financiera.anulacion_disponibilidad ON anulacion_disponibilidad.id = anulacion_disponibilidad_apropiacion.anulacion
							 WHERE apropiacion = ?  AND estado_anulacion = 3
							 GROUP BY  anulacion_disponibilidad.estado_anulacion, disponibilidad_apropiacion.apropiacion
							`, Id).Values(&maps_valor_tot)
	//fmt.Println("maps: ", len(maps_valor_tot))
	if len(maps_valor_tot) > 0 && err == nil {
		valor, _ = strconv.ParseFloat(maps_valor_tot[0]["valor"].(string), 64)
	} else {
		valor = 0
	}

	return
}

//----------------------------------------------------------
//funcion para generar canales de map[string]interface{}
func genChanMapStr(mp ...map[string]interface{}) <-chan map[string]interface{} {
	out := make(chan map[string]interface{})
	go func() {
		for _, ch := range mp {
			out <- ch
		}
		close(out)
	}()
	return out
}

//Generar ramas del arbol de rubros
func RamaApropiaciones(done <-chan map[string]interface{}, unidadEjecutora int, Vigencia int, forksin <-chan map[string]interface{}) (forksout <-chan map[string]interface{}) {
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
				_, err = o.Raw(`SELECT rubro.id as "Id", rubro.codigo as "Codigo",rubro.nombre as "Nombre", rubro.descripcion as "Descripcion", rubro.unidad_ejecutora as "UnidadEjecutora"
				  from financiera.rubro
				  join financiera.rubro_rubro
					on  rubro_rubro.rubro_hijo = rubro.id
				  WHERE rubro_rubro.rubro_padre = ?
				  AND unidad_ejecutora in (?,0)`, fork["Id"], unidadEjecutora).Values(&m)
				if err == nil {
					err = formatdata.FillStruct(m, &res)
					resch := genChanMapStr(res...)
					var hijos []map[string]interface{}
					wg.Add(1)
					subdone := make(chan map[string]interface{}) // HLdone
					defer close(subdone)
					for hijo := range RamaApropiaciones(subdone, unidadEjecutora, Vigencia, resch) {
						if hijo != nil {
							hijos = append(hijos, hijo) //tomar valores del canal y agregarlos al array de hijos.
						}
					}
					fork["Hijos"] = hijos
					//recorrer hijos sumando apropiaciones, si las tiene.
					if len(hijos) == 0 {
						query := make(map[string]string)
						var id string
						err = formatdata.FillStruct(fork["Id"], &id)
						query["Rubro.Id"] = id
						query["Vigencia"] = strconv.Itoa(Vigencia)
						v, err := GetAllApropiacion(query, nil, nil, nil, nil, 0, 1)
						if v != nil && err == nil {
							fork["Apropiacion"] = v[0]
							fork["Hijos"] = nil
						} else {
							// fork["Apropiacion"] = nil
							// fork["Hijos"] = nil
							fork = nil
						}
					} else {
						ap := Apropiacion{}
						var valorPadre float64
						valorPadre = 0
						for _, hijo := range hijos {
							if hijo["Apropiacion"] != nil {
								ap = Apropiacion{}
								formatdata.FillStruct(hijo["Apropiacion"], &ap)
								valorPadre = valorPadre + ap.Valor
							}
						}
						ap.Valor = valorPadre
						ap.Id = 0
						fork["Apropiacion"] = ap
					}

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
	// fmt.Println("rama generada...")
	return out
}

// Generar arbol de rubros.
func ArbolApropiaciones(unidadEjecutora int, Vigencia int) (padres []map[string]interface{}, err error) {
	o := orm.NewOrm()
	var m []orm.Params
	//funcion para conseguir los rubros padre.
	_, err = o.Raw(`  SELECT rubro.id as "Id", rubro.codigo as "Codigo",rubro.nombre as "Nombre", rubro.descripcion as "Descripcion", rubro.unidad_ejecutora as "UnidadEjecutora"
	    from financiera.rubro
	      where (id  in (select DISTINCT rubro_padre from financiera.rubro_rubro)
			  AND id not in (select DISTINCT rubro_hijo from financiera.rubro_rubro))
			  OR (id not in (select DISTINCT rubro_hijo from financiera.rubro_rubro)
			  AND id not in (select DISTINCT rubro_padre from financiera.rubro_rubro))
				AND rubro.unidad_ejecutora=?`, unidadEjecutora).Values(&m)
	if err == nil {
		fmt.Println("Generando arbol....")
		var res []map[string]interface{}
		err = formatdata.FillStruct(m, &res)
		resch := genChanMapStr(res...)
		done := make(chan map[string]interface{}) // HLdone
		defer close(done)                         // HLdone
		for padre := range RamaApropiaciones(done, unidadEjecutora, Vigencia, resch) {
			padres = append(padres, padre) //tomar valores del canal y agregarlos al array de hijos.
		}
	}
	fmt.Println("Arbol generado...")
	return
}

func EncapsuArbolApropiaciones(parameter ...interface{}) interface{} {
	unidadEjecutora := parameter[0].(int)
	vigencia := parameter[1].(int)

	if v, err := ArbolApropiaciones(unidadEjecutora, vigencia); err == nil {
		rankingsJson, _ := json.Marshal(v)
		err = ioutil.WriteFile("apropaciones_"+strconv.Itoa(unidadEjecutora)+"_"+strconv.Itoa(vigencia)+".json", rankingsJson, 0644)
		if err != nil {
			fmt.Println("Error escribiendo archivo....")
			return err
		}
		fmt.Println("Construyendo archivo....")
		return err
	} else {
		fmt.Println("error encapsula: ", err)
	}
	return nil
}

//SaldoRubroPadre... Funcion para determinar el saldo de un rubro padre a partir de sus hijos.
func SaldoRubroPadre(Id int, unidadEjecutora int, Vigencia int) (saldo map[string]float64, err error) {
	o := orm.NewOrm()
	var m []orm.Params
	var res []map[string]interface{}
	saldo = make(map[string]float64)
	//funcion para conseguir los hijos de los rubros padre.
	_, err = o.Raw(`SELECT rubro.id as "Id", rubro.codigo as "Codigo",rubro.nombre as "Nombre", rubro.descripcion as "Descripcion", rubro.unidad_ejecutora as "UnidadEjecutora"
	  from financiera.rubro
	  join financiera.rubro_rubro
		on  rubro_rubro.rubro_hijo = rubro.id
	  WHERE rubro_rubro.rubro_padre = ?
	  AND unidad_ejecutora in (?,0)`, Id, unidadEjecutora).Values(&m)
	if err == nil {
		err = formatdata.FillStruct(m, &res)

		resch := genChanMapStr(res...)
		done := make(chan map[string]interface{})
		defer close(done)
		for hijo := range RamaApropiaciones(done, unidadEjecutora, Vigencia, resch) {
			saldoaux, err := sumaApropiacionesHoja(hijo)
			if err == nil {
				saldo["original"] = saldo["original"] + saldoaux["original"]
				saldo["saldo"] = saldo["saldo"] + saldoaux["saldo"]
				saldo["comprometido"] = saldo["comprometido"] + saldoaux["comprometido"]
				saldo["comprometido_anulado"] = saldo["comprometido_anulado"] + saldoaux["comprometido_anulado"]
			} else {
				fmt.Println(err)
				return saldo, err
			}
		}

	}
	return
}

//sumaApropiacionesHoja... suma de los saldos de las apropiaciones hoja.
func sumaApropiacionesHoja(fork map[string]interface{}) (saldo map[string]float64, err error) {
	saldo = make(map[string]float64)
	ap := Apropiacion{}

	if fork == nil {
		return
	} else {
		if fork["Hijos"] == nil {
			err = formatdata.FillStruct(fork["Apropiacion"], &ap)
			if err == nil {
				saldo, err = SaldoApropiacion(ap.Id)
				if ap.Id == 240 {
					fmt.Println(ap)
				}
				if err != nil {
					fmt.Println("err 1 : ", err)
				}

				return
			} else {
				fmt.Println("err 1 : ", err)
				return
			}
		} else {
			var hijos []map[string]interface{}
			err = formatdata.FillStruct(fork["Hijos"], &hijos)
			if err == nil {
				for _, subfork := range hijos {
					saldoaux, err := sumaApropiacionesHoja(subfork)
					if err == nil {
						saldo["original"] = saldo["original"] + saldoaux["original"]
						saldo["saldo"] = saldo["saldo"] + saldoaux["saldo"]
						saldo["comprometido"] = saldo["comprometido"] + saldoaux["comprometido"]
						saldo["comprometido_anulado"] = saldo["comprometido_anulado"] + saldoaux["comprometido_anulado"]

					}

				}
				return
			} else {
				fmt.Println("err 2 : ", err)
				return
			}
		}
	}
}

//AprobarPresupuesto... Aprobacion de presupuesto (cambio de estado).
func AprobarPresupuesto(UnidadEjecutora int, Vigencia int) (err error) {
	query := make(map[string]string)
	o := orm.NewOrm()
	query["Rubro.UnidadEjecutora"] = strconv.Itoa(UnidadEjecutora)
	query["Vigencia"] = strconv.Itoa(Vigencia)
	fmt.Println(query)
	v, err := GetAllApropiacion(query, nil, nil, nil, nil, 0, -1)
	o.Begin()
	ap := Apropiacion{}
	for _, apropiacion := range v {
		formatdata.FillStruct(apropiacion, &ap)
		ap.Estado.Id = 2
		_, err = o.Update(&ap)
		if err != nil {
			o.Rollback()
			return
		}
	}
	o.Commit()
	return
}

//UpdateApropiacionValue... Actualiza la apropiacion inicial de un rubro dado un id
func UpdateApropiacionValue(id int, valor float64) (err error) {
	o := orm.NewOrm()
	apropiacion := &Apropiacion{Id: id, Valor: valor}
	_, err = o.Update(apropiacion, "Valor")
	if err != nil {
		panic(err.Error())
	}
	return
}

func ListaApropiacionesHijo(vigencia int, codigo string) (res []orm.Params, err error) {
	o := orm.NewOrm()
	//falta realizar proyeccion por cada rubro.
	_, err = o.Raw(`SELECT DISTINCT * FROM (SELECT apropiacion.id as Id ,rubro.id as idrubro, rubro.codigo, rubro.nombre as descripcion, apropiacion.vigencia, COALESCE( fuente.descripcion , 'Recursos Propios' ) as fdescrip, COALESCE( fuente.Id , 0 ) as idfuente
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
		order by idfuente,
							idrubro
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
			proy := 0.0
			aux := make(map[string]interface{})
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
				err = formatdata.FillStruct(fll["valor"], &ejstr)
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
				fll["proyeccion"] = proy
				var mp interface{}
				var mpp interface{}
				err = formatdata.FillStruct(variacion, &mp)
				err = formatdata.FillStruct(pvariacion, &mpp)
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
		if err != nil {
			fmt.Println("err1 ", err)
			return
		}

	}
	err = formatdata.FillStruct(m, &res)
	if err != nil {
		fmt.Println("err2 ", err)
		return
	}
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
				err = formatdata.FillStruct(fll["valor"], &ejstr)
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
				err = formatdata.FillStruct(variacion, &mp)
				err = formatdata.FillStruct(pvariacion, &mpp)
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
	err = formatdata.FillStruct(m, &res)
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
	err = formatdata.FillStruct(m, &res)
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
	err = formatdata.FillStruct(m, &res)
	return
}
