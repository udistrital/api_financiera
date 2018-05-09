package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/udistrital/utils_oas/formatdata"
)

type SolicitudDevolucion struct {
	Id               int                  `orm:"column(id);pk;auto"`
	Solicitante      *DocumentoDevolucion `orm:"column(solicitante);rel(fk)"`
	Beneficiario     *DocumentoDevolucion `orm:"column(beneficiario);rel(fk)"`
	FormaPago        *FormaPago           `orm:"column(forma_pago);rel(fk)"`
	RazonDevolucion  *RazonDevolucion     `orm:"column(razon_devolucion);rel(fk)"`
	Vigencia         float64              `orm:"column(vigencia)"`
	UnidadEjecutora  *UnidadEjecutora     `orm:"column(unidad_ejecutora);rel(fk)"`
	CuentaDevolucion *CuentaDevolucion    `orm:"column(cuenta_devolucion);rel(fk)"`
	Observaciones    string               `orm:"column(observaciones)"`
	FechaRegistro    time.Time            `orm:"column(fecha_registro);auto_now_add;type(datetime)"`
	Soporte          *ActaDevolucion      `orm:"column(soporte);rel(fk)"`
}

func (t *SolicitudDevolucion) TableName() string {
	return "solicitud_devolucion"
}

func init() {
	orm.RegisterModel(new(SolicitudDevolucion))
}

// AddSolicitudDevolucion insert a new SolicitudDevolucion into database and returns
// last inserted Id on success.
func AddSolicitudDevolucion(m *SolicitudDevolucion) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSolicitudDevolucionById retrieves SolicitudDevolucion by Id. Returns error if
// Id doesn't exist
func GetSolicitudDevolucionById(id int) (v *SolicitudDevolucion, err error) {
	o := orm.NewOrm()
	v = &SolicitudDevolucion{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSolicitudDevolucion retrieves all SolicitudDevolucion matches certain condition. Returns empty list if
// no records exist
func GetAllSolicitudDevolucion(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SolicitudDevolucion)).RelatedSel()
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

	var l []SolicitudDevolucion
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
// UpdateSolicitudDevolucion updates SolicitudDevolucion by Id and returns error if
// the record to be updated doesn't exist
func UpdateSolicitudDevolucionById(m *SolicitudDevolucion) (err error) {
	o := orm.NewOrm()
	v := SolicitudDevolucion{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSolicitudDevolucion deletes SolicitudDevolucion by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSolicitudDevolucion(id int) (err error) {
	o := orm.NewOrm()
	v := SolicitudDevolucion{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SolicitudDevolucion{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

// add a devolution, state,a id type returns error
//if any insert fails
func AddDevolution(request map[string]interface{}) (err error) {

	var solicitudDevol SolicitudDevolucion
	var documentoSol *DocumentoDevolucion
	var documentoBen *DocumentoDevolucion
	var documentoBusqeda DocumentoDevolucion
	var Id int64
	var idDevol int64
	var estadoDevol *EstadoDevolucion
	var solicitudEstado SolicitudDevolucionEstadoDevolucion
	var cuentaDevol CuentaDevolucion
	o := orm.NewOrm()

	err = formatdata.FillStruct(request["SolicitudDevolucion"], &solicitudDevol)
	err = formatdata.FillStruct(request["EstadoDevolucion"], &estadoDevol)

	if err == nil {
		o.Begin()

		solicitudEstado.EstadoDevolucion = estadoDevol
		solicitudEstado.Activo = true
		documentoBen = solicitudDevol.Beneficiario
		documentoSol = solicitudDevol.Solicitante

		err = o.QueryTable("cuenta_devolucion").
			Filter("banco", solicitudDevol.CuentaDevolucion.Banco).
			Filter("tipo_cuenta", solicitudDevol.CuentaDevolucion.TipoCuenta).
			Filter("numero_cuenta", solicitudDevol.CuentaDevolucion.NumeroCuenta).
			One(&cuentaDevol)
		beego.Error(err)
		if err == orm.ErrMultiRows {
			beego.Error("Returned Multi Rows Not One")
			return
		}

		if err == orm.ErrNoRows {
			Id, err = o.Insert(solicitudDevol.CuentaDevolucion)
			beego.Error("id", Id, "error", err)
			if err != nil {
				beego.Error(err)
				o.Rollback()
				return
			} else {
				solicitudDevol.CuentaDevolucion.Id = int(Id)
			}
		}
		if err == nil {
			solicitudDevol.CuentaDevolucion.Id = cuentaDevol.Id
		}

		err = o.QueryTable("documento_devolucion").
			Filter("Origen", documentoBen.Origen).
			Filter("tipo_identificacion", documentoBen.TipoIdentificacion).
			Filter("identificacion", documentoBen.Identificacion).
			One(&documentoBusqeda)

		if err == orm.ErrMultiRows {
			beego.Error("Returned Multi Rows Not One")
			return
		}
		if err == orm.ErrNoRows {

			Id, err = o.Insert(documentoBen)
			documentoBen.Id = int(Id)
			if err != nil {
				o.Rollback()
				return
			}
		}

		if err == nil {
			documentoBen.Id = documentoBusqeda.Id
		}

		err = o.QueryTable("documento_devolucion").
			Filter("Origen", documentoSol.Origen).
			Filter("tipo_identificacion", documentoSol.TipoIdentificacion).
			Filter("identificacion", documentoSol.Identificacion).
			One(&documentoBusqeda)

		if err == orm.ErrMultiRows {
			beego.Error("Error consultado documento solicitante")
			return
		}

		if err == orm.ErrNoRows {
			Id, err = o.Insert(documentoSol)
			documentoSol.Id = int(Id)
			if err != nil {
				o.Rollback()
				return
			}
		}
		if err == nil {
			documentoSol.Id = documentoBusqeda.Id
		}

		idDevol, err = o.Insert(&solicitudDevol)
		if err != nil {
			o.Rollback()
			return
		}
		beego.Error("id devolucion", idDevol)
		beego.Error("Solicitud estado", solicitudEstado)
		solicitudDevol.Id = int(idDevol)
		solicitudEstado.Devolucion = &solicitudDevol
		_, err = o.Insert(&solicitudEstado)
		if err != nil {
			o.Rollback()
			return
		}
	} else {
		return
	}
	o.Commit()
	return
}
