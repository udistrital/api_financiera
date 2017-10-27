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

//AddTransaccionConcepto funcion para la transaccion de conceptos -Agregar un concepto
func AddTransaccionConcepto(m *TrConcepto) (err error) {
	o := orm.NewOrm()
	//alerta = append(alerta, "success")
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
				//alerta := models.Alert{Type: "error", Code: "E_3542", Body: nil}
				return err
				//alerta[0] = "error"
				//alerta = append(alerta, "Ocurrio un error al insertar las afectaciones!")
			}
			fmt.Println("Afectacion", &v)
		}
		fmt.Println("padre", m.ConceptoPadre)

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

		if m.ConceptoPadre.Id != 0 {
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
	}
	fmt.Println(err)
	o.Rollback()
	return err
	//alerta[0] = "error"
	//alerta = append(alerta, "Ocurrio un error al insertar el concepto!")

}
