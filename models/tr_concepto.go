package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

type TrConcepto struct {
	Concepto      *Concepto
	ConceptoPadre *Concepto
	Afectaciones  *[]AfectacionConcepto
	Cuentas       *[]CuentaContable
}

//funcion para la transaccion de conceptos -Agregar un concepto
func AddTransaccionConcepto(m *TrConcepto) (alerta []string, err error) {
	o := orm.NewOrm()
	alerta = append(alerta, "success")
	m.Concepto.FechaCreacion = time.Now()
	fmt.Println("Concepto!!!!", m.Concepto.FechaCreacion)
	o.Begin()

	if id, err := o.Insert(m.Concepto); err == nil {
		m.Concepto.Id = int(id)

		fmt.Println("Concepto", m.Concepto)
		for _, v := range *m.Afectaciones {
			v.Concepto = m.Concepto
			if _, err = o.Insert(&v); err != nil {
				fmt.Println("Afectacion", &v)
				o.Rollback()
				alerta[0] = "error"
				alerta = append(alerta, "Ocurrio un error al insertar las afectaciones!")
			}
			fmt.Println("Afectacion", &v)
		}
		fmt.Println("padre", m.ConceptoPadre)

		for _, c := range *m.Cuentas {
			var concepto_cuentas ConceptoCuentaContable
			concepto_cuentas.Concepto = m.Concepto
			concepto_cuentas.CuentaContable = &c
			if _, err = o.Insert(&concepto_cuentas); err != nil {
				fmt.Println("error concepto_cuentas", err)
				o.Rollback()
				alerta[0] = "error"
				alerta = append(alerta, "Ocurrio un error al insertar las cuentas al concepto!")
			}
			fmt.Println("exito concepto_cuentas", &concepto_cuentas)
		}

		if m.ConceptoPadre.Id != 0 {
			var conceptoestructura = new(ConceptoConcepto)
			conceptoestructura.ConceptoHijo = m.Concepto
			conceptoestructura.ConceptoPadre = m.ConceptoPadre

			fmt.Println("padre estructura", conceptoestructura.ConceptoPadre)
			if _, err = o.Insert(conceptoestructura); err != nil {
				o.Rollback()
				alerta[0] = "error"
				alerta = append(alerta, "Ocurrio un error al insertar el concepto en la estructura!")
			}
		}
		o.Commit()
		alerta = append(alerta, "El concepto se agrego exitosamente!")
	} else {
		fmt.Println(err)
		o.Rollback()
		alerta[0] = "error"
		alerta = append(alerta, "Ocurrio un error al insertar el concepto!")
	}
	return
}
