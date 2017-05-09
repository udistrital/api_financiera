package models

import (
	"strconv"

	"github.com/astaxie/beego/orm"
)

//ArbolPlanCuentas estructura que se retorna al consultar un plan de cuentas
type ArbolPlanCuentas struct {
	Id                 int                 `orm:"column(id);pk;auto"`
	Saldo              float64             `orm:"column(saldo)"`
	Nombre             string              `orm:"column(nombre)"`
	Naturaleza         string              `orm:"column(naturaleza)"`
	Descripcion        string              `orm:"column(descripcion);null"`
	Codigo             string              `orm:"column(codigo)"`
	NivelClasificacion *NivelClasificacion `orm:"column(nivel_clasificacion);rel(fk)"`
	CuentaBancaria     *CuentaBancaria     `orm:"column(cuenta_bancaria);rel(fk);null"`
	Hijos              *[]ArbolPlanCuentas
}

//MakeTreePlanCuentas construye el arbol de la estructura de un plan de cuentas
func MakeTreePlanCuentas(plan int) (a []ArbolPlanCuentas) {
	o := orm.NewOrm()
	//Arreglo
	var arbol []ArbolPlanCuentas
	idplan := strconv.Itoa(plan)

	_, err := o.Raw("select * from financiera.cuenta_contable where id not in (select cuenta_hijo from financiera.estructura_cuentas where cuenta_hijo is not null) and id in (select cuenta_padre from financiera.estructura_cuentas where plan_cuentas=" + idplan + ") order by id;").QueryRows(&arbol)

	if err == nil {
		//For para que recorra los Ids en busca de hijos
		for i := 0; i < len(arbol); i++ {
			var l CuentaContable
			o.QueryTable(&l).Filter("Id", arbol[i].Id).RelatedSel().All(&l)
			arbol[i].NivelClasificacion = l.NivelClasificacion
			//verifica que los Id tengan hijos
			MakeBranchesPlan(&arbol[i], plan)
		}

	}
	return arbol
}

//MakeBranchesPlan Función que construye los hijos del arbol
func MakeBranchesPlan(Padre *ArbolPlanCuentas, plan int) (a []ArbolPlanCuentas) {
	o := orm.NewOrm()
	//Conversión de entero a string
	idpadre := strconv.Itoa(Padre.Id)
	idplan := strconv.Itoa(plan)
	//Arreglo
	var arbol []ArbolPlanCuentas
	_, err := o.Raw("select a.* from financiera.cuenta_contable a left join financiera.estructura_cuentas b on a.id =b.cuenta_hijo where b.cuenta_padre=" + idpadre + "and b.plan_cuentas = " + idplan + " ORDER BY a.id").QueryRows(&arbol)
	//Condicional si el error es nulo
	if err == nil {
		//Llena el elemento Opciones en la estructura del padre
		Padre.Hijos = &arbol
		//For que recorre el subMenu en busca de hijos (Recursiva)
		for i := 0; i < len(arbol); i++ {
			var l CuentaContable
			o.QueryTable(&l).Filter("Id", arbol[i].Id).RelatedSel().All(&l)
			arbol[i].NivelClasificacion = l.NivelClasificacion
			//verifica que los Id tengan hijos
			MakeBranchesPlan(&arbol[i], plan)
		}
	}
	return arbol
}
