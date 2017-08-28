package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

type TrConcepto struct {
	Concepto      *ConceptoTesoral
	ConceptoPadre *ConceptoTesoral
	Afectaciones  *[]AfectacionConceptoTesoral
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
			v.ConceptoTesoral = m.Concepto
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
			var concepto_cuentas ConceptoTesoralCuentaContable
			concepto_cuentas.ConceptoTesoral = m.Concepto
			concepto_cuentas.CuentaContable = &c
			if string(c.Codigo[0]) == "9" {
				fmt.Println("CODIGO:", string(c.Codigo[0]))
				concepto_cuentas.CuentaAcreedora = true
			}
			if _, err = o.Insert(&concepto_cuentas); err != nil {
				fmt.Println("error concepto_cuentas", err)
				o.Rollback()
				alerta[0] = "error"
				alerta = append(alerta, "Ocurrio un error al insertar las cuentas al concepto!")
			}
			fmt.Println("exito concepto_cuentas", &concepto_cuentas)
		}

		if m.ConceptoPadre.Id != 0 {
			var conceptoestructura = new(EstructuraConceptosTesorales)
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
