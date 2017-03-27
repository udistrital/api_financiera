package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type TrCuentaContable struct {
	Cuenta      *CuentaContable
	CuentaPadre *CuentaContable
	PlanCuentas *PlanCuentas
}

//funcion para la transaccion de conceptos -Agregar un concepto
func AddTransaccionCuentaContable(m *TrCuentaContable) (alerta []string, err error) {
	o := orm.NewOrm()
	o.Begin()
	alerta = append(alerta, "success")
	if _, err = o.Insert(m.Cuenta); err == nil {
		//m.Concepto.Id = int(id)
		fmt.Println("Cuenta->", m.Cuenta)
		var estructuracuentas = new(EstructuraCuentas)
		estructuracuentas.PlanCuentas = m.PlanCuentas
		fmt.Println("padre:", m.CuentaPadre)
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
			//fmt.Println(err)
			o.Commit()
		}
	} else {
		alerta[0] = "error"
		alerta = append(alerta, "Ocurrio un error al insertar la cuenta!")
		o.Rollback()
	}
	//fmt.Println(err)
	return
}
