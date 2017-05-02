package models

import "github.com/astaxie/beego/orm"

type TrCuentaContable struct {
	Cuenta      *CuentaContable
	CuentaPadre *CuentaContable
	PlanCuentas *PlanCuentas
}

//AddTransaccionCuentaContable funcion para la transaccion de agregar cuentas contables sobre el plan de cuentas
func AddTransaccionCuentaContable(m *TrCuentaContable) (alerta []string, err error) {
	o := orm.NewOrm()
	o.Begin()
	alerta = append(alerta, "success")
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
			alerta[0] = "error"
			alerta = append(alerta, "Ocurrio un error al insertar la cuenta en el plan!")
		} else {
			alerta = append(alerta, "Cuenta "+m.Cuenta.Codigo+"-"+m.Cuenta.Nombre+" agregada exitosamente")
			o.Commit()
		}
	} else {
		alerta[0] = "error"
		alerta = append(alerta, "Ocurrio un error al insertar la cuenta!")
		o.Rollback()
	}
	return
}
