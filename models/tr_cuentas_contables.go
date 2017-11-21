package models

import (
	"errors"
	"fmt"

	"github.com/astaxie/beego/orm"
)

type TrCuentaContable struct {
	Cuenta      *CuentaContable
	CuentaPadre *CuentaContable
	PlanCuentas *PlanCuentas
}

//AddTransaccionCuentaContable funcion para la transaccion de agregar cuentas contables sobre el plan de cuentas
func AddTransaccionCuentaContable(m *TrCuentaContable) (err error) {
	o := orm.NewOrm()
	o.Begin()
	//alerta = append(alerta, "success")
	if _, err = o.Insert(m.Cuenta); err == nil {
		var estructuracuentas = new(EstructuraCuentas)
		estructuracuentas.PlanCuentas = m.PlanCuentas
		if m.CuentaPadre.Id != 0 {
			estructuracuentas.CuentaPadre = m.CuentaPadre
			estructuracuentas.CuentaHijo = m.Cuenta
		} else {
			estructuracuentas.CuentaPadre = m.Cuenta
			estructuracuentas.CuentaHijo = nil
		}
		if _, err = o.Insert(estructuracuentas); err != nil {
			o.Rollback()
			//alerta[0] = "error"
			//alerta = append(alerta, "Ocurrio un error al insertar la cuenta en el plan!")
		} else {
			//alerta = append(alerta, "Cuenta "+m.Cuenta.Codigo+"-"+m.Cuenta.Nombre+" agregada exitosamente")
			o.Commit()
			return nil
		}
	} else {
		//alerta[0] = "error"
		//alerta = append(alerta, "Ocurrio un error al insertar la cuenta!")
		o.Rollback()
	}
	return
}

// UpdateTrCuentaContable updates CuentaContable by Id and returns error,cod  if the record to be updated doesn't exist
func UpdateTrCuentaContable(m *TrCuentaContable) (err error, cod string) {
	o := orm.NewOrm()
	o.Begin()

	v := CuentaContable{Id: m.Cuenta.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var back int64
		if o.Raw(`select id from financiera.movimiento_contable where cuenta_contable = ?
							union
							select id from financiera.concepto_tesoral_cuenta_contable where cuenta_contable= ?`, v.Id, v.Id).QueryRow(&back); back > 0 {
			err = errors.New("04566")
			o.Rollback()
			return err, ""
		} else {
			if m.CuentaPadre != nil {
				var ec EstructuraCuentas
				if err = o.Raw("select * from financiera.estructura_cuentas where cuenta_hijo=? and plan_cuentas=?", m.Cuenta.Id, m.PlanCuentas.Id).QueryRow(&ec); err == nil {
					ec.CuentaPadre = m.CuentaPadre
					if _, err = o.Update(&ec); err == nil {
						m.Cuenta.Codigo = m.CuentaPadre.Codigo + "-" + m.Cuenta.Codigo[len(m.Cuenta.Codigo)-int(m.Cuenta.NivelClasificacion.Longitud):]
						fmt.Print(m.Cuenta.Codigo)
					} else {
						o.Rollback()
						return err, ""
					}
				} else {
					o.Rollback()
					return err, ""
				}
			} else {
				m.Cuenta.Codigo = v.Codigo[:len(v.Codigo)-int(m.Cuenta.NivelClasificacion.Longitud)] + m.Cuenta.Codigo[len(m.Cuenta.Codigo)-int(m.Cuenta.NivelClasificacion.Longitud):]
				fmt.Print("codigo-->", m.Cuenta.Codigo)
			}
			if _, err = o.Update(m.Cuenta); err == nil {
				o.Commit()
				return nil, m.Cuenta.Codigo
			} else {
				o.Rollback()
				return err, ""
			}
		}
	} else {
		o.Rollback()
		return err, ""
	}
}
