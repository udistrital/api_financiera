package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type ConceptoValor struct {
	Concepto *Concepto
	Valor    int64
}

type RpCdpRubroConceptoValor struct {
	RegistroPresupuestalDisponibilidadApropiacion RegistroPresupuestalDisponibilidadApropiacion
	ConceptoValor                                 []ConceptoValor
}

type HomologacionConcepto struct {
	Id             int       `orm:"column(id);pk;auto"`
	Vigencia       float64   `orm:"column(vigencia)"`
	FechaCreacion  time.Time `orm:"column(fecha_creacion);type(date)"`
	ConceptoKronos *Concepto `orm:"column(concepto_kronos);rel(fk)"`
	ConceptoTitan  int       `orm:"column(concepto_titan)"`
	NominaTitan    int       `orm:"column(nomina_titan)"`
}

func (t *HomologacionConcepto) TableName() string {
	return "homologacion_concepto"
}

func init() {
	orm.RegisterModel(new(HomologacionConcepto))
}

// AddHomologacionConcepto insert a new HomologacionConcepto into database and returns
// last inserted Id on success.
func AddHomologacionConcepto(m *HomologacionConcepto) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetHomologacionConceptoById retrieves HomologacionConcepto by Id. Returns error if
// Id doesn't exist
func GetHomologacionConceptoById(id int) (v *HomologacionConcepto, err error) {
	o := orm.NewOrm()
	v = &HomologacionConcepto{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllHomologacionConcepto retrieves all HomologacionConcepto matches certain condition. Returns empty list if
// no records exist
func GetAllHomologacionConcepto(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(HomologacionConcepto)).RelatedSel(5)
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
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

	var l []HomologacionConcepto
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				o.LoadRelated(v.ConceptoKronos, "ConceptoCuentaContable", 5)
				o.LoadRelated(v.ConceptoKronos, "ConceptoTesoralFacultadProyecto", 5)
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

// UpdateHomologacionConcepto updates HomologacionConcepto by Id and returns error if
// the record to be updated doesn't exist
func UpdateHomologacionConceptoById(m *HomologacionConcepto) (err error) {
	o := orm.NewOrm()
	v := HomologacionConcepto{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteHomologacionConcepto deletes HomologacionConcepto by Id and returns error if
// the record to be deleted doesn't exist
func DeleteHomologacionConcepto(id int) (err error) {
	o := orm.NewOrm()
	v := HomologacionConcepto{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&HomologacionConcepto{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
func validarExistenciaHomologacionConcepto(inputHomologacion map[string]interface{}) (outputRegistrarHomologacion, conFacultad bool) {
	var proyectoC int
	//var conFacultad bool
	var existeHomologacionConcepto bool
	var existeConceptoTesoralFacultadProyecto bool
	o := orm.NewOrm()
	o.Begin()

	homologacion := HomologacionConcepto{
		Vigencia:       inputHomologacion["Vigencia"].(float64),
		NominaTitan:    int(inputHomologacion["NominaTitan"].(float64)),
		ConceptoKronos: &Concepto{Id: int(inputHomologacion["ConceptoKronos"].(float64))},
		ConceptoTitan:  int(inputHomologacion["ConceptoTitan"].(float64)),
	}
	err := o.Read(&homologacion, "Vigencia", "NominaTitan", "ConceptoKronos", "ConceptoTitan")
	if err == orm.ErrNoRows {
		existeHomologacionConcepto = false
	} else {
		existeHomologacionConcepto = true
	}

	if inputHomologacion["Facultad"] != nil {
		conFacultad = true
		facultad := int(inputHomologacion["Facultad"].(float64))
		if inputHomologacion["ProyectoCurricular"] == nil {
			proyectoC = 0
		} else {
			proyectoC = int(inputHomologacion["ProyectoCurricular"].(float64))
		}
		conceptoFacultad := ConceptoTesoralFacultadProyecto{
			ConceptoTesoral:    &Concepto{Id: int(inputHomologacion["ConceptoKronos"].(float64))},
			Facultad:           facultad,
			ProyectoCurricular: proyectoC,
		}
		err := o.Read(&conceptoFacultad, "ConceptoTesoral", "Facultad", "ProyectoCurricular")
		if err == orm.ErrNoRows {
			existeConceptoTesoralFacultadProyecto = false
		} else {
			existeConceptoTesoralFacultadProyecto = true
		}
	} else {
		conFacultad = false
	}

	if conFacultad {
		if existeHomologacionConcepto == false && existeConceptoTesoralFacultadProyecto == false {
			outputRegistrarHomologacion = true
		} else {
			outputRegistrarHomologacion = false
		}
	} else {
		if existeHomologacionConcepto == false {
			outputRegistrarHomologacion = true
		} else {
			outputRegistrarHomologacion = false
		}
	}
	return
}

// RegistrarHomologacionConcepto
func RegistrarHomologacionConcepto(dataHomologacionConcepto map[string]interface{}) (alerta Alert) {
	// validar existencia
	fmt.Println("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
	registrarHomologacion, conFacultad := validarExistenciaHomologacionConcepto(dataHomologacionConcepto)
	println("------------------")
	var proyectoC int
	o := orm.NewOrm()
	o.Begin()
	if registrarHomologacion == true {
		if conFacultad == true {
			println("dos registro ")
			// 2 Registros: registra tabla ConceptoTesoralFacultadProyecto y HomologacionConcepto
			facultad := int(dataHomologacionConcepto["Facultad"].(float64))
			if dataHomologacionConcepto["ProyectoCurricular"] == nil {
				proyectoC = 0
			} else {
				proyectoC = int(dataHomologacionConcepto["ProyectoCurricular"].(float64))
			}
			conceptoFacultad := ConceptoTesoralFacultadProyecto{
				ConceptoTesoral:    &Concepto{Id: int(dataHomologacionConcepto["ConceptoKronos"].(float64))},
				Facultad:           facultad,
				ProyectoCurricular: proyectoC,
			}
			_, err := o.Insert(&conceptoFacultad)
			if err != nil {
				alerta.Type = "error"
				alerta.Code = "E_OPP_01"
				alerta.Body = err.Error()
				o.Rollback()
				return
			}
			homologacion := HomologacionConcepto{
				Vigencia:       dataHomologacionConcepto["Vigencia"].(float64),
				NominaTitan:    int(dataHomologacionConcepto["NominaTitan"].(float64)),
				ConceptoKronos: &Concepto{Id: int(dataHomologacionConcepto["ConceptoKronos"].(float64))},
				ConceptoTitan:  int(dataHomologacionConcepto["ConceptoTitan"].(float64)),
			}
			idHomologacion, err := o.Insert(&homologacion)
			if err != nil {
				alerta.Type = "error"
				alerta.Code = "E_OPP_01"
				alerta.Body = err.Error()
				o.Rollback()
				return
			}
			alerta = Alert{Type: "success", Code: "S_HOMO_01", Body: idHomologacion}
			o.Commit()
			return
		} else {
			println("un registro")
			// 1 Registros: registra tabla HomologacionConcepto
			homologacion := HomologacionConcepto{
				Vigencia:       dataHomologacionConcepto["Vigencia"].(float64),
				NominaTitan:    int(dataHomologacionConcepto["NominaTitan"].(float64)),
				ConceptoKronos: &Concepto{Id: int(dataHomologacionConcepto["ConceptoKronos"].(float64))},
				ConceptoTitan:  int(dataHomologacionConcepto["ConceptoTitan"].(float64)),
			}
			idHomologacion, err := o.Insert(&homologacion)
			if err != nil {
				alerta.Type = "error"
				alerta.Code = "E_OPP_01"
				alerta.Body = err.Error()
				o.Rollback()
				return
			}
			alerta = Alert{Type: "success", Code: "S_HOMO_01", Body: idHomologacion}
			o.Commit()
			return
		}
	} else {
		println("error que ya existe homologacion")
		alerta = Alert{Type: "success", Code: "S_HOMO_01", Body: "Ya existe homologacion"}
		return
	}

}
