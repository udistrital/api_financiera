package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"
	"github.com/astaxie/beego/orm"
	"strconv"
)

type Comprobante struct {
	Id                int                `orm:"column(id);pk;auto"`
	Secuencia         int                `orm:"column(secuencia)"`
	NumeroItems       int                `orm:"column(numero_items)"`
	RedondeoCifras    bool               `orm:"column(redondeo_cifras)"`
	Ano               int                `orm:"column(ano)"`
	Mes               int                `orm:"column(mes)"`
	FechaRegistro     time.Time          `orm:"column(fecha_registro);type(timestamp without time zone)"`
	TipoComprobante   *TipoComprobante   `orm:"column(tipo_comprobante);rel(fk)"`
	EstadoComprobante *EstadoComprobante `orm:"column(estado_comprobante);rel(fk)"`
	Observaciones     string 						 `orm:"column(observaciones)"`
}

func (t *Comprobante) TableName() string {
	return "comprobante"
}

func init() {
	orm.RegisterModel(new(Comprobante))
}

// AddComprobante insert a new Comprobante into database and returns
// last inserted Id on success.
func AddComprobante(m *Comprobante) (id int64, err error) {
	o := orm.NewOrm()
	m.FechaRegistro = time.Now()
	m.Mes = int(time.Now().Month())
	m.Ano  = time.Now().Year()
	id, err = o.Insert(m)
	return
}

// GetComprobanteById retrieves Comprobante by Id. Returns error if
// Id doesn't exist
func GetComprobanteById(id int) (v *Comprobante, err error) {
	o := orm.NewOrm()
	v = &Comprobante{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllComprobante retrieves all Comprobante matches certain condition. Returns empty list if
// no records exist
func GetAllComprobante(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Comprobante)).RelatedSel(1)
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

	var l []Comprobante
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

// UpdateComprobante updates Comprobante by Id and returns error if
// the record to be updated doesn't exist
func UpdateComprobanteById(m *Comprobante) (err error) {
	o := orm.NewOrm()
	v := Comprobante{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteComprobante deletes Comprobante by Id and returns error if
// the record to be deleted doesn't exist
func DeleteComprobante(id int) (err error) {
	o := orm.NewOrm()
	v := Comprobante{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Comprobante{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func CrearComprobanteOrdenPago(op OrdenPago){
	fmt.Println("hola soy la orden de pago creada", op)
	var consulta_homologacion = make(map[string]string);
	var consulta_movimiento_contable = make(map[string]string);
	var consulta_tercero = make(map[string]string);
	var fields []string
	var sortby []string
	var order []string
	var ObjetoHomologacion HomologacionComprobantes
	var ObjetoMovimientoContable MovimientoContable
	var ObjetoRP RegistroPresupuestal
	var valor float64

 	consulta_homologacion["TipoDocumentoAfectante.CodigoAbreviacion"] = "DA-OP"
	consulta_homologacion["TipoDocumentoAfectante.Activo"] = "true"

	//Buscar el tipo de comprobante por documento OP
	respuesta, err := GetAllHomologacionComprobantes(consulta_homologacion,fields,sortby,order,0,-1)
	ObjetoHomologacion = respuesta[0].(HomologacionComprobantes)

	//CREAR NUEVO COMPROBANTE
	nuevo_comprobante := &Comprobante{Secuencia: op.Consecutivo,NumeroItems: 250,RedondeoCifras: true,	Ano: time.Now().Year(),Mes: int(time.Now().Month()),FechaRegistro: time.Now(),TipoComprobante: &TipoComprobante{Id:ObjetoHomologacion.TipoComprobante.Id},	EstadoComprobante : &EstadoComprobante{Id:1},Observaciones: "Creada automáticamente para OP"}
	id_nuevo, err := AddComprobante(nuevo_comprobante)

  //BUSCAR TERCERO
	//consulta_tercero["OrdenPago.Id"] = strconv.Itoa(op.RegistroPresupuestal.Id)
	respuesta, _ = GetAllOrdenPagoRegistroPresupuestal(consulta_tercero,fields,sortby,order,0,-1)
	ObjetoRP = respuesta[0].(RegistroPresupuestal)

	if(id_nuevo != 0 && err == nil){
		consulta_movimiento_contable["TipoDocumentoAfectante.Id"] = "1"
		consulta_movimiento_contable["CodigoDocumentoAfectante"] = strconv.Itoa(op.Id)
		fmt.Println("comprobante creado exitosamente", consulta_movimiento_contable)
		respuesta, err := GetAllMovimientoContable(consulta_movimiento_contable,fields,sortby,order,0,-1)
		fmt.Println(err)
		for i, v := range respuesta {
			 ObjetoMovimientoContable = v.(MovimientoContable)
			 if(ObjetoMovimientoContable.CuentaContable.Naturaleza == "debito"){
				 valor = float64(ObjetoMovimientoContable.Debito)
				}else{
				 valor = float64(ObjetoMovimientoContable.Credito)
			 }
			 ObjetoRegistroComprobante := &RegistroComprobantes { Comprobante: &Comprobante{Id: int(id_nuevo)}, 	Movimiento: op.Id, Secuencia: i+1,	MovimientoContable: &MovimientoContable{Id:ObjetoMovimientoContable.Id }, CuentaContable: ObjetoMovimientoContable.CuentaContable.Id, TipoDocumentoAfectante: &TipoDocumentoAfectante{Id:1 }, Valor: valor , Tercero: ObjetoRP.Id}
			 _, err := AddRegistroComprobantes(ObjetoRegistroComprobante)
			 fmt.Println(err)
		 }
	}else{
		fmt.Println("error", err)
	}
}
