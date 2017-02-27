package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"
	"strconv"
	"github.com/astaxie/beego/orm"
)
type Info_disponibilidad_a_anular struct {
	Anulacion                  AnulacionDisponibilidad
	Disponibilidad_apropiacion []DisponibilidadApropiacion
	Valor                      float64
}
type Disponibilidad struct {
	Id                   int                   `orm:"column(id);pk"`
	UnidadEjecutora      *UnidadEjecutora      `orm:"column(unidad_ejecutora);rel(fk)"`
	Vigencia             float64               `orm:"column(vigencia)"`
	NumeroDisponibilidad float64               `orm:"column(numero_disponibilidad);null"`
	Responsable          int                 `orm:"column(responsable);null"`
	Solicitante          int                 `orm:"column(solicitante);null"`
	FechaRegistro        time.Time             `orm:"column(fecha_registro);type(date);null"`
	ModalidadGiro        int16                 `orm:"column(modalidad_giro);null"`
	Estado               *EstadoDisponibilidad `orm:"column(estado);rel(fk)"`
	NumeroOficio         string                `orm:"column(numero_oficio);null"`
	Objeto               string                `orm:"column(objeto);null"`
	VigenciaFutura       float64               `orm:"column(vigencia_futura);null"`
	Destino              int                 `orm:"column(destino);null"`
	Solicitud            int                 `orm:"column(solicitud)"`
}

func (t *Disponibilidad) TableName() string {
	return "disponibilidad"
}

func init() {
	orm.RegisterModel(new(Disponibilidad))
}

// AddDisponibilidad insert a new Disponibilidad into database and returns
// last inserted Id on success.
func AddDisponibilidad(m *Disponibilidad) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetDisponibilidadById retrieves Disponibilidad by Id. Returns error if
// Id doesn't exist
func GetDisponibilidadById(id int) (v *Disponibilidad, err error) {
	o := orm.NewOrm()
	v = &Disponibilidad{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllDisponibilidad retrieves all Disponibilidad matches certain condition. Returns empty list if
// no records exist
func GetAllDisponibilidad(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Disponibilidad))
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

	var l []Disponibilidad
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

// UpdateDisponibilidad updates Disponibilidad by Id and returns error if
// the record to be updated doesn't exist
func UpdateDisponibilidadById(m *Disponibilidad) (err error) {
	o := orm.NewOrm()
	v := Disponibilidad{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteDisponibilidad deletes Disponibilidad by Id and returns error if
// the record to be deleted doesn't exist
func DeleteDisponibilidad(id int) (err error) {
	o := orm.NewOrm()
	v := Disponibilidad{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Disponibilidad{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}


func AnulacionTotal(m *Info_disponibilidad_a_anular) (alerta []string, err error) {
	o := orm.NewOrm()
	o.Begin()
	m.Anulacion.FechaRegistro = time.Now()
	id_anulacion_cdp, err1 := o.Insert(&m.Anulacion)
	fmt.Println("error")
	if err1 != nil {
		alerta = append(alerta, "No se pudo registrar el detalle de la anulacion")
		err = err1
		o.Rollback()
		return
	}
	for i := 0; i < len(m.Disponibilidad_apropiacion); i++ {

		saldoCDP, err2 := SaldoCdp(m.Disponibilidad_apropiacion[i].Disponibilidad.Id, m.Disponibilidad_apropiacion[i].Apropiacion.Id)
		if err2 != nil {
			alerta = append(alerta, "No se pudo cargar el saldo del CDP N° "+strconv.FormatFloat(m.Disponibilidad_apropiacion[i].Disponibilidad.NumeroDisponibilidad, 'f', -1, 64)+" para la apropiacion del Rubro "+m.Disponibilidad_apropiacion[i].Apropiacion.Rubro.Codigo)
			err = err2
			o.Rollback()
			return
		}
		anulacion_apropiacion := AnulacionDisponibilidadApropiacion{
			DisponibilidadApropiacion: &m.Disponibilidad_apropiacion[i],
			Anulacion:                 &AnulacionDisponibilidad{Id: int(id_anulacion_cdp)},
			Valor:                     saldoCDP,
		}
		_, err3 := o.Insert(&anulacion_apropiacion)
		if err3 != nil {
			alerta = append(alerta, "No se pudo registrar la anulacion del CDP N° "+strconv.FormatFloat(m.Disponibilidad_apropiacion[i].Disponibilidad.NumeroDisponibilidad, 'f', -1, 64)+" para la apropiacion del Rubro "+m.Disponibilidad_apropiacion[i].Apropiacion.Rubro.Codigo)
			err = err3
			o.Rollback()
			return
		} else {
			alerta = append(alerta, "se anulo del CDP N° "+strconv.FormatFloat(m.Disponibilidad_apropiacion[i].Disponibilidad.NumeroDisponibilidad, 'f', -1, 64)+" para la apropiacion del Rubro "+m.Disponibilidad_apropiacion[i].Apropiacion.Rubro.Codigo+" la suma de "+strconv.FormatFloat(saldoCDP, 'f', -1, 64))

		}
	}

	o.Commit()
	return
}

func AnulacionParcial(m *Info_disponibilidad_a_anular) (alerta []string, err error) {
	o := orm.NewOrm()
	o.Begin()
	m.Anulacion.FechaRegistro = time.Now()
	id_anulacion_cdp, err1 := o.Insert(&m.Anulacion)
	if err1 != nil {
		alerta = append(alerta, "No se pudo registrar el detalle de la anulacion")
		err = err1
		o.Rollback()
		return
	}
	for i := 0; i < len(m.Disponibilidad_apropiacion); i++ {

		saldoCDP, err2 := SaldoCdp(m.Disponibilidad_apropiacion[i].Disponibilidad.Id, m.Disponibilidad_apropiacion[i].Apropiacion.Id)
		if err2 != nil {
			alerta = append(alerta, "No se pudo cargar el saldo del CDP N° "+strconv.FormatFloat(m.Disponibilidad_apropiacion[i].Disponibilidad.NumeroDisponibilidad, 'f', -1, 64)+" para la apropiacion del Rubro "+m.Disponibilidad_apropiacion[i].Apropiacion.Rubro.Codigo)
			err = err2
			o.Rollback()
			return
		}
		fmt.Println("anulacion: ", m.Valor)
		if saldoCDP < m.Valor {
			alerta = append(alerta, "Valor a anular supera el saldo del CDP N° "+strconv.FormatFloat(m.Disponibilidad_apropiacion[i].Disponibilidad.NumeroDisponibilidad, 'f', -1, 64)+" para la apropiacion del Rubro "+m.Disponibilidad_apropiacion[i].Apropiacion.Rubro.Codigo)
			o.Rollback()
			return
		} else {
			anulacion_apropiacion := AnulacionDisponibilidadApropiacion{
				DisponibilidadApropiacion: &m.Disponibilidad_apropiacion[i],
				Anulacion:                 &AnulacionDisponibilidad{Id: int(id_anulacion_cdp)},
				Valor:                     m.Valor,
			}
			_, err3 := o.Insert(&anulacion_apropiacion)
			if err3 != nil {
				alerta = append(alerta, "No se pudo registrar la anulacion del CDP N° "+strconv.FormatFloat(m.Disponibilidad_apropiacion[i].Disponibilidad.NumeroDisponibilidad, 'f', -1, 64)+" para la apropiacion del Rubro "+m.Disponibilidad_apropiacion[i].Apropiacion.Rubro.Codigo)
				err = err3
				o.Rollback()
				return
			} else {
				alerta = append(alerta, "se anulo del CDP N° "+strconv.FormatFloat(m.Disponibilidad_apropiacion[i].Disponibilidad.NumeroDisponibilidad, 'f', -1, 64)+" para la apropiacion del Rubro "+m.Disponibilidad_apropiacion[i].Apropiacion.Rubro.Codigo+" la suma de "+strconv.FormatFloat(m.Valor, 'f', -1, 64))

			}
		}

	}

	o.Commit()
	return
}

//funcion para obtener valor total del cdp
func ValorCdp(id_cdp int, id_apropiacion int) (valor float64, err error) {
	o := orm.NewOrm()
	var maps []orm.Params
	o.Raw("SELECT SUM(valor) as valor from financiera.disponibilidad_apropiacion WHERE disponibilidad = ? AND apropiacion = ? ", id_cdp, id_apropiacion).Values(&maps)

	if maps == nil {
		valor = 0
	} else {
		valor, err = strconv.ParseFloat(maps[0]["valor"].(string), 64)
	}

	//fmt.Println("valor: ", valor)
	return
}

//----------------------------------------
//funcion para obtener saldo restante del cdp
func SaldoCdp(id_cdp int, id_apropiacion int) (valor float64, err error) {
	o := orm.NewOrm()
	var maps []orm.Params
	var maps_valor_anulaciones []orm.Params
	var valorDesagregado float64
	var valorAnulado float64
	valorCDP, _ := ValorCdp(id_cdp, id_apropiacion)
	o.Raw(`SELECT SUM(a.valor) as valor from financiera.registro_presupuestal_disponibilidad_apropiacion as a ,
				 financiera.disponibilidad_apropiacion as b
				WHERE b.disponibilidad = ?  AND b.apropiacion = ? AND a.disponibilidad_apropiacion = b.id;`, id_cdp, id_apropiacion).Values(&maps)
	fmt.Println("maps: ", maps)
	if maps[0]["valor"] == nil {
		valorDesagregado = 0
	} else {
		valorDesagregado, err = strconv.ParseFloat(maps[0]["valor"].(string), 64)
	}
	o.Raw(`SELECT  COALESCE(SUM(b.valor),0) as valor

				FROM financiera.anulacion_disponibilidad_apropiacion AS b
				LEFT JOIN
				     financiera.disponibilidad_apropiacion AS a
				ON a.id = b.disponibilidad_apropiacion

				WHERE  a.apropiacion = ? AND a.disponibilidad = ?`, id_apropiacion, id_cdp).Values(&maps_valor_anulaciones)
	if maps_valor_anulaciones[0]["valor"] == nil {
		valorAnulado = 0
	} else {
		valorAnulado, err = strconv.ParseFloat(maps_valor_anulaciones[0]["valor"].(string), 64)
	}
	valor = valorCDP - valorDesagregado - valorAnulado
	fmt.Println("valor: ", valor)
	return
}

//----------------------------------------
