package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

type TrConcepto struct {
	Concepto      *Concepto
	ConceptoPadre *Concepto
	Afectaciones  *[]AfectacionConcepto
	Cuentas       *[]CuentaContable
	DelCuentas    *[]ConceptoCuentaContable
}

//AddTransaccionConcepto funcion para la transaccion de conceptos -Agregar un concepto
func AddTransaccionConcepto(m *TrConcepto) (err error) {
	o := orm.NewOrm()
	//alerta = append(alerta, "success")
	m.Concepto.FechaCreacion = time.Now()
	fmt.Println("Concepto!!!!", m.Concepto.FechaCreacion)
	o.Begin()

	if id, err := o.Insert(m.Concepto); err == nil {
		m.Concepto.Id = int(id)
		fmt.Print("entro aqui")

		fmt.Println("Concepto", m.Concepto)
		if m.Afectaciones != nil {
			for _, v := range *m.Afectaciones {
				v.Concepto = m.Concepto
				if _, err = o.Insert(&v); err != nil {
					fmt.Println("Afectacion", &v)
					o.Rollback()
					//alerta := models.Alert{Type: "error", Code: "E_3542", Body: nil}
					return err
					//alerta[0] = "error"
					//alerta = append(alerta, "Ocurrio un error al insertar las afectaciones!")
				}
				fmt.Println("Afectacion", &v)
			}
		}

		fmt.Println("padre", m.ConceptoPadre)

		if m.Cuentas != nil {
			for _, c := range *m.Cuentas {
				var concepto_cuentas ConceptoCuentaContable
				concepto_cuentas.Concepto = m.Concepto
				concepto_cuentas.CuentaContable = &c
				if string(c.Codigo[0]) == "9" {
					fmt.Println("CODIGO:", string(c.Codigo[0]))
					concepto_cuentas.CuentaAcreedora = true
				}
				if _, err = o.Insert(&concepto_cuentas); err != nil {
					fmt.Println("error concepto_cuentas", err)
					o.Rollback()
					return err
					//alerta[0] = "error"
					//alerta = append(alerta, "Ocurrio un error al insertar las cuentas al concepto!")
				}
				fmt.Println("exito concepto_cuentas", &concepto_cuentas)
			}
		}

		if m.ConceptoPadre.Id != 0 {
			if m.ConceptoPadre.Clasificador != true {
				o.Rollback()
				err = errors.New("C92011")
				return err
			}
			var conceptoestructura = new(ConceptoConcepto)
			conceptoestructura.ConceptoHijo = m.Concepto
			conceptoestructura.ConceptoPadre = m.ConceptoPadre

			fmt.Println("padre estructura", conceptoestructura.ConceptoPadre)
			if _, err = o.Insert(conceptoestructura); err != nil {
				o.Rollback()
				return err
				//alerta[0] = "error"
				//alerta = append(alerta, "Ocurrio un error al insertar el concepto en la estructura!")
			}
		}
		o.Commit()
		return nil
		//alerta = append(alerta, "El concepto se agrego exitosamente!")
	} else {
		fmt.Println("murio", err)
		o.Rollback()
		return err
	}
	//alerta[0] = "error"
	//alerta = append(alerta, "Ocurrio un error al insertar el concepto!")

}

func UpdateConceptoTr(m *TrConcepto) (err error) {
	o := orm.NewOrm()
	v := Concepto{Id: m.Concepto.Id}
	o.Begin()
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if o.Raw(`select count(concepto) from
						(select concepto_tesoral concepto from financiera.movimiento_contable where concepto_tesoral=?
						union
						select concepto_kronos concepto from financiera.homologacion_concepto
						where concepto_kronos = ?
						union
						select concepto_padre concepto from financiera.estructura_conceptos_tesorales where concepto_padre =?
						 ) a`, v.Id, v.Id, v.Id).QueryRow(&num); num > 0 {
			err = errors.New("04566")
			return
		} else {
			uc := &Concepto{
				Id:              v.Id,
				Codigo:          m.Concepto.Codigo,
				Clasificador:    m.Concepto.Clasificador,
				TipoConcepto:    m.Concepto.TipoConcepto,
				Rubro:           m.Concepto.Rubro,
				FechaCreacion:   v.FechaCreacion,
				Nombre:          m.Concepto.Nombre,
				Descripcion:     m.Concepto.Descripcion,
				FechaExpiracion: m.Concepto.FechaExpiracion,
			}
			//var num int64

			if m.ConceptoPadre.Id != 0 {
				if m.ConceptoPadre.Clasificador != true || m.ConceptoPadre.Id == m.Concepto.Id {
					o.Rollback()
					err = errors.New("C92011")
					return err
				}
				var conceptoestructura = new(ConceptoConcepto)
				conceptoestructura.ConceptoHijo = m.Concepto
				conceptoestructura.ConceptoPadre = m.ConceptoPadre
				var numc int
				if o.Raw(`select id from financiera.estructura_conceptos_tesorales where concepto_hijo =?`, v.Id).QueryRow(&numc); numc > 0 {
					fmt.Print("numero estructura", numc)
					conceptoestructura.Id = numc
					if _, err = o.Update(conceptoestructura); err != nil {
						o.Rollback()
						return err
					}
				} else {
					if _, err = o.Insert(conceptoestructura); err != nil {
						o.Rollback()
						return err
					}
				}
			}

			if _, err = o.Update(uc); err == nil {
				//fmt.Println("Number of records updated in database:", num)

				if uc.Clasificador {
					if _, err := o.Raw(`delete from financiera.afectacion_concepto_tesoral
													where concepto_tesoral =  ?`, uc.Id).Exec(); err != nil {
						o.Rollback()
						return err
					}
					if _, err := o.Raw(`delete from financiera.concepto_tesoral_cuenta_contable
															where concepto_tesoral =  ?`, uc.Id).Exec(); err != nil {
						o.Rollback()
						return err
					}
				} else {

					if m.Afectaciones != nil {
						for _, af := range *m.Afectaciones {
							if af.Id != 0 {
								vaf := AfectacionConcepto{Id: af.Id}
								// ascertain id exists in the database
								if err = o.Read(&vaf); err == nil {
									//var num int64
									if _, err = o.Update(&af); err != nil {
										o.Rollback()
										return
										//fmt.Println("Number of records updated in database:", num)
									}
								} else {
									o.Rollback()
									return
								}
							} else {
								if _, err = o.Insert(&af); err != nil {
									fmt.Println("Afectacion", &v)
									o.Rollback()
									//alerta := models.Alert{Type: "error", Code: "E_3542", Body: nil}
									return err
									//alerta[0] = "error"
									//alerta = append(alerta, "Ocurrio un error al insertar las afectaciones!")
								}
							}

						}
					}

					if m.Cuentas != nil {
						for _, c := range *m.Cuentas {
							var concepto_cuentas ConceptoCuentaContable
							concepto_cuentas.Concepto = m.Concepto
							concepto_cuentas.CuentaContable = &c
							if string(c.Codigo[0]) == "9" {
								fmt.Println("CODIGO:", string(c.Codigo[0]))
								concepto_cuentas.CuentaAcreedora = true
							}
							if _, err = o.Insert(&concepto_cuentas); err != nil {
								o.Rollback()
								return err
								//alerta[0] = "error"
								//alerta = append(alerta, "Ocurrio un error al insertar las cuentas al concepto!")
							}
						}
					}

					if m.DelCuentas != nil {
						for _, cdel := range *m.DelCuentas {
							vcdel := ConceptoCuentaContable{Id: cdel.Id}
							// ascertain id exists in the database
							if err = o.Read(&vcdel); err == nil {
								//var num int64
								if _, err = o.Delete(&vcdel); err != nil {
									o.Rollback()
									return
									//fmt.Println("Number of records deleted in database:", num)
								}
							} else {
								o.Rollback()
								return
							}
						}
					}

				}

				o.Commit()
			} else {
				o.Rollback()
			}
			return
		}
	} else {
		o.Rollback()
		return
	}

}
