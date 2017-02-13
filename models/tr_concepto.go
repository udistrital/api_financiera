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
}

//funcion para la transaccion de conceptos -Agregar un concepto
func AddTransaccionConcepto(m *TrConcepto) (id int64, err error) {
	o := orm.NewOrm()
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
				err = o.Rollback()
			}
			fmt.Println("Afectacion", &v)
		}
		fmt.Println("padre", m.ConceptoPadre)

		if m.ConceptoPadre.Id != 0 {
			var conceptoestructura = new(ConceptoConcepto)
			conceptoestructura.ConceptoHijo = m.Concepto
			conceptoestructura.ConceptoPadre = m.ConceptoPadre

			fmt.Println("padre estructura", conceptoestructura.ConceptoPadre)
			if _, err = o.Insert(conceptoestructura); err != nil {
				err = o.Rollback()
			}
		}
		err = o.Commit()
	} else {
		err = o.Rollback()
	}
	return
}
