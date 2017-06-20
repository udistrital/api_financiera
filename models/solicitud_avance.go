package models

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/udistrital/api_financiera/utilidades"
)

type SolicitudAvance struct {
	Id                       int     `orm:"column(id);pk;auto"`
	IdBeneficiario           int     `orm:"column(id_beneficiario)"`
	Vigencia                 string  `orm:"column(vigencia)"`
	Consecutivo              string  `orm:"column(consecutivo)"`
	Objetivo                 string  `orm:"column(objetivo)"`
	Justificacion            string  `orm:"column(justificacion)"`
	ValorTotal               float64 `orm:"column(valor_total)"`
	CodigoDependencia        string  `orm:"column(codigo_dependencia)"`
	Dependencia              string  `orm:"column(dependencia)"`
	CodigoFacultad           string  `orm:"column(codigo_facultad);null"`
	Facultad                 string  `orm:"column(facultad);null"`
	CodigoProyectoCurricular string  `orm:"column(codigo_proyecto_curricular)"`
	ProyectoCurricular       string  `orm:"column(proyecto_curricular)"`
	CodigoConvenio           string  `orm:"column(codigo_convenio);null"`
	Convenio                 string  `orm:"column(convenio);null"`
	CodigoProyectoInv        string  `orm:"column(codigo_proyecto_inv);null"`
	ProyectoInv              string  `orm:"column(proyecto_inv);null"`
	Estado                   string  `orm:"column(estado)"`
}

func (t *SolicitudAvance) TableName() string {
	return "solicitud_avance"
}

func init() {
	orm.RegisterModel(new(SolicitudAvance))
}

func TrSolicitudAvance(m map[string]interface{}) (solicitud SolicitudAvance, err error) {
	//var id int64
	err = utilidades.FillStruct(m["Solicitud"], &solicitud)
	if err == nil {
		fmt.Println("Solicitud: ", solicitud)
		solicitud.Estado = "A"
		//solicitud.Fecha = time.Now()
		solicitud.Vigencia = strconv.Itoa(time.Now().Year())
		o := orm.NewOrm()
		o.Begin()
		var consecutivo float64
		err = o.Raw(`SELECT COALESCE(MAX(consecutivo), 0)+1  as consecutivo
						FROM financiera.solicitud_avance WHERE vigencia = ?`, solicitud.Vigencia).QueryRow(&consecutivo)
		if err == nil {
			solicitud.Consecutivo = strconv.FormatFloat(consecutivo, 'E', -1, 64)
			//insert ingreso
			_, err = o.Insert(&solicitud)
			if err == nil {
				tipoAvance := SolicitudTipoAvance{}
				err = utilidades.FillStruct(m["TipoAvance"], &tipoAvance)
				if err == nil {
					tipoAvance.Estado = "A"
					tipoAvance.SolicitudAvance = &solicitud
					_, err = o.Insert(&tipoAvance)
					if err == nil {
						o.Commit()
						return
					} else {
						o.Rollback()
						return
					}
				} else {
					o.Rollback()
					return
				}
			} else {
				o.Rollback()
				return
			}
		} else {
			o.Rollback()
			return
		}
	} else {
		return
	}
}

// AddSolicitudAvance insert a new SolicitudAvance into database and returns
// last inserted Id on success.
func AddSolicitudAvance(m *SolicitudAvance) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSolicitudAvanceById retrieves SolicitudAvance by Id. Returns error if
// Id doesn't exist
func GetSolicitudAvanceById(id int) (v *SolicitudAvance, err error) {
	o := orm.NewOrm()
	v = &SolicitudAvance{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSolicitudAvance retrieves all SolicitudAvance matches certain condition. Returns empty list if
// no records exist
func GetAllSolicitudAvance(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SolicitudAvance))
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

	var l []SolicitudAvance
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

// UpdateSolicitudAvance updates SolicitudAvance by Id and returns error if
// the record to be updated doesn't exist
func UpdateSolicitudAvanceById(m *SolicitudAvance) (err error) {
	o := orm.NewOrm()
	v := SolicitudAvance{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSolicitudAvance deletes SolicitudAvance by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSolicitudAvance(id int) (err error) {
	o := orm.NewOrm()
	v := SolicitudAvance{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SolicitudAvance{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
