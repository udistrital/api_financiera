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

type SolicitudRequisitoTipoAvance struct {
	Id                  int                  `orm:"column(id);pk;auto"`
	RequisitoTipoAvance *RequisitoTipoAvance `orm:"column(requisito_tipo_avance);rel(fk)"`
	SolicitudTipoAvance *SolicitudTipoAvance `orm:"column(solicitud_tipo_avance);rel(fk)"`
	Valido              string               `orm:"column(valido);null"`
	Observaciones       string               `orm:"column(observaciones);null"`
	FechaRegistro       time.Time            `orm:"column(fecha_registro);type(timestamp with time zone)"`
	Documento           string               `orm:"column(documento);null"`
	Estado              string               `orm:"column(estado)"`
	UbicacionDoc        string               `orm:"column(ubicacion_doc);null"`
}

func (t *SolicitudRequisitoTipoAvance) TableName() string {
	return "solicitud_requisito_tipo_avance"
}

func init() {
	orm.RegisterModel(new(SolicitudRequisitoTipoAvance))
}

// AddSolicitudRequisitoTipoAvance insert a new SolicitudRequisitoTipoAvance into database and returns
// last inserted Id on success.
func AddSolicitudRequisitoTipoAvance(m *SolicitudRequisitoTipoAvance) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

func TrValidarAvance(m map[string]interface{}) (estado EstadoAvance, err error) {
	requisitos := []RequisitoTipoAvance{}
	solicitudRequisitoTipoAvance := []SolicitudRequisitoTipoAvance{}
	solicitud := SolicitudAvance{}
	estadoAvance := EstadoAvance{}
	estadoVerificado := Estados{}
	err = utilidades.FillStruct(m["Requisitos"], &requisitos)
	err = utilidades.FillStruct(m["Solicitud"], &solicitud)
	o := orm.NewOrm()
	o.Begin()
	if err == nil {
		for _, data := range solicitudRequisitoTipoAvance {
			data.Estado = "A"
			data.Valido = "SI"
			data.FechaRegistro = time.Now()
			fmt.Println("solicitud_requisito_tipo_avance: ", data)
			_, err = o.Insert(&data)
		}
		if err == nil {
			estadoVerificado.Id = 6
			estadoAvance.Estados = &estadoVerificado
			estadoAvance.SolicitudAvance = &solicitud
			estadoAvance.FechaRegistro = time.Now()
			estadoAvance.Usuario = "System"
			estadoAvance.Observaciones = "Requisitos Verificados"
			estadoAvance.Estado = "A"
			fmt.Println("solicitud: ", solicitud)
			_, err = o.Insert(&estadoAvance)
			if err == nil {
				o.Commit()
			} else {
				fmt.Println("Error", err)
			}
		}
	}
	return
}

// GetSolicitudRequisitoTipoAvanceById retrieves SolicitudRequisitoTipoAvance by Id. Returns error if
// Id doesn't exist
func GetSolicitudRequisitoTipoAvanceById(id int) (v *SolicitudRequisitoTipoAvance, err error) {
	o := orm.NewOrm()
	v = &SolicitudRequisitoTipoAvance{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSolicitudRequisitoTipoAvance retrieves all SolicitudRequisitoTipoAvance matches certain condition. Returns empty list if
// no records exist
func GetAllSolicitudRequisitoTipoAvance(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SolicitudRequisitoTipoAvance)).RelatedSel()
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

	var l []SolicitudRequisitoTipoAvance
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

// UpdateSolicitudRequisitoTipoAvance updates SolicitudRequisitoTipoAvance by Id and returns error if
// the record to be updated doesn't exist
func UpdateSolicitudRequisitoTipoAvanceById(m *SolicitudRequisitoTipoAvance) (err error) {
	o := orm.NewOrm()
	v := SolicitudRequisitoTipoAvance{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSolicitudRequisitoTipoAvance deletes SolicitudRequisitoTipoAvance by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSolicitudRequisitoTipoAvance(id int) (err error) {
	o := orm.NewOrm()
	v := SolicitudRequisitoTipoAvance{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SolicitudRequisitoTipoAvance{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
