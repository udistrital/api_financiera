package models

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/udistrital/utils_oas/formatdata"
)

type Giro struct {
	Id             int               `orm:"column(id);pk;auto"`
	Consecutivo    int               `orm:"column(consecutivo)"`
	ValorTotal     float64           `orm:"column(valor_total)"`
	CuentaBancaria *CuentaBancaria   `orm:"column(cuenta_bancaria);rel(fk)"`
	Vigencia       float64           `orm:"column(vigencia);null"`
	FechaRegistro  time.Time         `orm:"column(fecha_registro);type(date)"`
	FormaPago      *FormaPago        `orm:"column(forma_pago);rel(fk)"`
	GiroDetalle    []*GiroDetalle    `orm:"reverse(many)"`
	GiroEstadoGiro []*GiroEstadoGiro `orm:"reverse(many)"`
}

type GiroAlert struct {
	Type        string
	Code        string
	Body        interface{}
	IdGiro      int64
	OrdenesPago []map[string]interface{}
}

func (t *Giro) TableName() string {
	return "giro"
}

func init() {
	orm.RegisterModel(new(Giro))
}

// AddGiro insert a new Giro into database and returns
// last inserted Id on success.
func AddGiro(m *Giro) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetGiroById retrieves Giro by Id. Returns error if
// Id doesn't exist
func GetGiroById(id int) (v *Giro, err error) {
	o := orm.NewOrm()
	v = &Giro{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllGiro retrieves all Giro matches certain condition. Returns empty list if
// no records exist
func GetAllGiro(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Giro)).RelatedSel()
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

	var l []Giro
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				o.LoadRelated(&v, "GiroEstadoGiro", 5, 1, 0, "-Id")
				o.LoadRelated(&v, "GiroDetalle", 5)
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

// UpdateGiro updates Giro by Id and returns error if
// the record to be updated doesn't exist
func UpdateGiroById(m *Giro) (err error) {
	o := orm.NewOrm()
	v := Giro{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteGiro deletes Giro by Id and returns error if
// the record to be deleted doesn't exist
func DeleteGiro(id int) (err error) {
	o := orm.NewOrm()
	v := Giro{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Giro{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
func RegistrarGiroDescuentos(e []interface{}, idGiro int64, idCuenta int64, idOrdenPago int64) (alerta Alert) {
	var idCuentasEspeciales []int
	var giroDetalles []GiroDetalle
	var idTipoCuenta int
	var idNewCuentaTercero CuentaBancariaEnte
	var element map[string]interface{}
	o := orm.NewOrm()
	o.Begin()
	element = e[0].(map[string]interface{})
	nameTipoCuenta := element["TipoCuentaBancaria"].(string)
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("Id").
		From("financiera.tipo_cuenta_bancaria").
		Where("nombre = ?")
	err := o.Raw(qb.String(), strings.Title(strings.ToLower(nameTipoCuenta))).QueryRow(&idTipoCuenta)
	//fmt.Println("idTipoCuenta -> ", idTipoCuenta)
	if err != nil {
		alerta.Type = "error"
		alerta.Code = "E_GIRO_04"
		alerta.Body = idTipoCuenta
		o.Rollback()
		return
	}
	err = o.QueryTable("cuenta_bancaria_ente").
		Filter("banco", element["IdEntidadBancaria"]).
		Filter("tipo_cuenta", idTipoCuenta).
		Filter("numero_cuenta", element["NumCuentaBancaria"]).One(&idNewCuentaTercero)
	if err == nil {
		element["CuentaBancariaEnte"] = idNewCuentaTercero.Id
		// fmt.Println("Existe Cuenta", element["Proveedor"].(map[string]interface{})["CuentaBancariaEnte"])
	} else if err == orm.ErrMultiRows {
		beego.Error("Returned Multi Rows Not One")
		return
	} else if err == orm.ErrNoRows {
		// fmt.Println(reflect.TypeOf(element["Proveedor"].(map[string]interface{})["NumDocumento"]))
		titular, _ := strconv.Atoi(element["NumDocumento"].(string))
		idNewCuentaTercero := CuentaBancariaEnte{
			Banco:        int(element["IdEntidadBancaria"].(float64)),
			TipoCuenta:   int(idTipoCuenta),
			NumeroCuenta: element["NumCuentaBancaria"].(string),
			Titular:      titular,
		}
		// fmt.Println(idNewCuentaTercero)
		ID, err := o.Insert(&idNewCuentaTercero)
		if err != nil {
			fmt.Println(err)
			beego.Error(err)
			o.Rollback()
			return
		} else {
			// fmt.Println(ID)
			element["CuentaBancariaEnte"] = int(ID)
		}
	}
	rowGiroDetalle := GiroDetalle{
		Giro:               &Giro{Id: int(idGiro)},
		OrdenPago:          &OrdenPago{Id: int(idOrdenPago)},
		CuentaBancariaEnte: &CuentaBancariaEnte{Id: element["CuentaBancariaEnte"].(int)},
		CuentaEspecial:     &CuentaEspecial{Id: int(idCuenta)},
	}
	giroDetalles = append(giroDetalles, rowGiroDetalle)
	fmt.Println("cuentas_especiales", idCuentasEspeciales)

	// insertar giro_detalle
	_, err = o.InsertMulti(100, giroDetalles)
	if err != nil {
		alerta.Type = "error"
		alerta.Code = "E_GIRO_CUENTA_ESPECIAL_02"
		alerta.Body = err.Error()
		o.Rollback()
		return
	}

	o.Commit()
	return

}
func GetCuentasEspeciales(id int64) (cuentas []orm.Params, alerta Alert) {
	o := orm.NewOrm()
	o.Begin()
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("opce.cuenta_especial, ce.informacion_persona_juridica").
		From("financiera.orden_pago_cuenta_especial as opce").
		InnerJoin("financiera.cuenta_especial as ce").On("opce.cuenta_especial = ce.id").
		And("opce.orden_pago = ?")
	_, err := o.Raw(qb.String(), id).Values(&cuentas)
	if err != nil {
		alerta.Type = "error"
		alerta.Code = "E_OPCUENTA_ESPECIAL_01"
		alerta.Body = err.Error()
		o.Rollback()
		return
	}
	o.Commit()
	return

}

func RegistrarGiro(dataGiro map[string]interface{}) (alerta GiroAlert) {
	o := orm.NewOrm()
	o.Begin()
	newGiro := Giro{}
	var OrdenesPago []map[string]interface{}
	var idNewCuentaTercero CuentaBancariaEnte
	// var idCuentasEspeciales []int
	err1 := formatdata.FillStruct(dataGiro["Giro"], &newGiro)
	err2 := formatdata.FillStruct(dataGiro["OrdenPago"], &OrdenesPago)
	if err1 != nil || err2 != nil {
		alerta.Type = "error"
		alerta.Code = "E_GIRO_01" //error en parametros de entrada
		alerta.Body = "Error en parametros de entrada en RegistrarGiro()"
		o.Rollback()
		return
	}
	// consecutivo
	var consecutivo int
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("COALESCE(MAX(consecutivo), 0)+1 ").
		From("financiera.giro")
	err := o.Raw(qb.String()).QueryRow(&consecutivo)
	if err != nil {
		alerta.Type = "error"
		alerta.Code = "E_OPP_01"
		alerta.Body = consecutivo
		o.Rollback()
		return
	}
	newGiro.Consecutivo = consecutivo
	newGiro.FechaRegistro = time.Now()
	//insert giro
	idNewGiro, err := o.Insert(&newGiro)
	if err != nil {
		alerta.Type = "error"
		alerta.Code = "E_GIRO_01"
		alerta.Body = err.Error()
		o.Rollback()
		return
	}
	// Primer estado
	estadoNewGiro := EstadoGiro{CodigoAbreviatura: "EGI_01"}
	err = o.Read(&estadoNewGiro, "CodigoAbreviatura")
	if err != nil {
		alerta.Type = "error"
		alerta.Code = "E_GIRO_02" //en busqueda de estado
		alerta.Body = err.Error()
		o.Rollback()
		return
	}
	//insert giro_estado_giro
	newGiroEstadoGiro := GiroEstadoGiro{}
	newGiroEstadoGiro.Giro = &Giro{Id: int(idNewGiro)}
	newGiroEstadoGiro.FechaRegistro = time.Now()
	newGiroEstadoGiro.EstadoGiro = &EstadoGiro{Id: int(estadoNewGiro.Id)}
	_, err = o.Insert(&newGiroEstadoGiro)
	if err != nil {
		alerta.Type = "error"
		alerta.Code = "E_GIRO_03"
		alerta.Body = err.Error()
		o.Rollback()
		return
	}
	//insert giro_detalle and orden_pago_estado_ordenPago
	var giroDetalles []GiroDetalle
	var newEstadoOrdenPago []OrdenPagoEstadoOrdenPago
	newEstadoOP08, alertaAux := GetEstadoOrdenPago("EOP_08")
	alerta.Type = alertaAux.Type
	if alerta.Type == "error" {
		o.Rollback()
		return
	}
	// get or insert cuenta_bancaria_ente
	for _, element := range OrdenesPago {
		var idTipoCuenta int
		nameTipoCuenta := element["Proveedor"].(map[string]interface{})["TipoCuentaBancaria"].(string)
		qb, _ := orm.NewQueryBuilder("mysql")
		qb.Select("Id").
			From("financiera.tipo_cuenta_bancaria").
			Where("nombre = ?")
		err := o.Raw(qb.String(), strings.Title(strings.ToLower(nameTipoCuenta))).QueryRow(&idTipoCuenta)
		//fmt.Println("idTipoCuenta -> ", idTipoCuenta)
		if err != nil {
			alerta.Type = "error"
			alerta.Code = "E_GIRO_04"
			alerta.Body = idTipoCuenta
			o.Rollback()
			return
		}

		err = o.QueryTable("cuenta_bancaria_ente").
			Filter("banco", element["Proveedor"].(map[string]interface{})["IdEntidadBancaria"]).
			Filter("tipo_cuenta", idTipoCuenta).
			Filter("numero_cuenta", element["Proveedor"].(map[string]interface{})["NumCuentaBancaria"]).One(&idNewCuentaTercero)
		if err == nil {
			element["Proveedor"].(map[string]interface{})["CuentaBancariaEnte"] = idNewCuentaTercero.Id
			// fmt.Println("Existe Cuenta", element["Proveedor"].(map[string]interface{})["CuentaBancariaEnte"])
		} else if err == orm.ErrMultiRows {
			beego.Error("Returned Multi Rows Not One")
			return
		} else if err == orm.ErrNoRows {
			// fmt.Println(reflect.TypeOf(element["Proveedor"].(map[string]interface{})["NumDocumento"]))
			titular, _ := strconv.Atoi(element["Proveedor"].(map[string]interface{})["NumDocumento"].(string))
			idNewCuentaTercero := CuentaBancariaEnte{
				Banco:        int(element["Proveedor"].(map[string]interface{})["IdEntidadBancaria"].(float64)),
				TipoCuenta:   int(idTipoCuenta),
				NumeroCuenta: element["Proveedor"].(map[string]interface{})["NumCuentaBancaria"].(string),
				Titular:      titular,
			}
			// fmt.Println(idNewCuentaTercero)
			ID, err := o.Insert(&idNewCuentaTercero)
			if err != nil {
				fmt.Println(err)
				beego.Error(err)
				o.Rollback()
				return
			} else {
				// fmt.Println(ID)
				element["Proveedor"].(map[string]interface{})["CuentaBancariaEnte"] = int(ID)
			}
		}
		//giro detalle
		rowGiroDetalle := GiroDetalle{
			Giro:               &Giro{Id: int(idNewGiro)},
			OrdenPago:          &OrdenPago{Id: int(element["Id"].(float64))},
			CuentaBancariaEnte: &CuentaBancariaEnte{Id: element["Proveedor"].(map[string]interface{})["CuentaBancariaEnte"].(int)},
			CuentaEspecial:     &CuentaEspecial{Id: 0},
		}
		giroDetalles = append(giroDetalles, rowGiroDetalle)
		// estados orden pago
		rowEstadoOrdenPago := OrdenPagoEstadoOrdenPago{
			OrdenPago:       &OrdenPago{Id: int(element["Id"].(float64))},
			EstadoOrdenPago: &EstadoOrdenPago{Id: int(newEstadoOP08.Id)},
			FechaRegistro:   time.Now(),
			Usuario:         1, //entra por sesion
		}
		newEstadoOrdenPago = append(newEstadoOrdenPago, rowEstadoOrdenPago)

	}

	// insertar giro_detalle
	_, err = o.InsertMulti(100, giroDetalles)
	if err != nil {
		alerta.Type = "error"
		alerta.Code = "E_GIRO_01"
		alerta.Body = err.Error()
		o.Rollback()
		return
	}
	// insertar nuevo estado para las Ordenes de pago
	_, err = o.InsertMulti(100, newEstadoOrdenPago)
	if err != nil {
		alerta.Type = "error"
		alerta.Code = "E_GIRO_01"
		alerta.Body = err.Error()
		o.Rollback()
		return
	}
	alerta = GiroAlert{Type: "success", Code: "S_GIRO_01", Body: consecutivo, IdGiro: idNewGiro, OrdenesPago: OrdenesPago}
	o.Commit()
	return
}
