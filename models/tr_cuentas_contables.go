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
func AddTransaccionCuentaContable(m *TrCuentaContable) (id int64, err error) {
	o := orm.NewOrm()
	o.Begin()

	if _, err = o.Insert(m.Cuenta); err == nil {
		//m.Concepto.Id = int(id)

		fmt.Println("Cuenta->", m.Cuenta)
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
		} else {
			fmt.Println(err)
			o.Commit()
		}
	} else {
		//fmt.Println(err)
		o.Rollback()
	}
	fmt.Println(err)
	return
}
