package models

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"

	"github.com/astaxie/beego/orm"
	"github.com/udistrital/api_financiera/utilidades"
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
					err = utilidades.FillStruct(m, &res)
					resch := genChanMapStr(res...)
					var hijos []map[string]interface{}
					wg.Add(1)
					subdone := make(chan map[string]interface{}) // HLdone
					defer close(subdone)
					for hijo := range RamaApropiaciones(subdone, unidadEjecutora, Vigencia, resch) {
						hijos = append(hijos, hijo) //tomar valores del canal y agregarlos al array de hijos.
					}
					fork["Hijos"] = hijos
					//recorrer hijos sumando apropiaciones, si las tiene.
					if len(hijos) == 0 {
						query := make(map[string]string)
						var id string
						err = utilidades.FillStruct(fork["Id"], &id)
						query["Rubro.Id"] = id
						query["Vigencia"] = strconv.Itoa(Vigencia)
						v, err := GetAllApropiacion(query, nil, nil, nil, nil, 0, 1)
						if v != nil && err == nil {
							fork["Apropiacion"] = v[0]
							fork["Hijos"] = nil
						} else {
							fork["Apropiacion"] = nil
							fork["Hijos"] = nil
						}
					} else {
						ap := Apropiacion{}
						var valorPadre float64
						valorPadre = 0
						for _, hijo := range hijos {
							if hijo["Apropiacion"] != nil {
								ap = Apropiacion{}
								utilidades.FillStruct(hijo["Apropiacion"], &ap)
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
			  AND id not in (select DISTINCT rubro_padre from financiera.rubro_rubro))`).Values(&m)
	if err == nil {
		var res []map[string]interface{}
		err = utilidades.FillStruct(m, &res)
		resch := genChanMapStr(res...)
		done := make(chan map[string]interface{}) // HLdone
		defer close(done)                         // HLdone
		for padre := range RamaApropiaciones(done, unidadEjecutora, Vigencia, resch) {
			padres = append(padres, padre) //tomar valores del canal y agregarlos al array de hijos.
		}
	}
	return
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
		err = utilidades.FillStruct(m, &res)

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
			err = utilidades.FillStruct(fork["Apropiacion"], &ap)
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
			err = utilidades.FillStruct(fork["Hijos"], &hijos)
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
		utilidades.FillStruct(apropiacion, &ap)
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
