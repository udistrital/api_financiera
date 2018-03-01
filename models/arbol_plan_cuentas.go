package models

import (
	"fmt"
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
	NivelClasificacion *NivelClasificacion `orm:"column(nivel_clasificacion_cuenta_contable);rel(fk)"`
	CuentaBancaria     *CuentaBancaria     `orm:"column(cuenta_bancaria);rel(fk);null"`
	Hijos              *[]ArbolPlanCuentas
}

//MakeTreePlanCuentas construye el arbol de la estructura de un plan de cuentas
func MakeTreePlanCuentas(plan int) (a []ArbolPlanCuentas, err error) {
	o := orm.NewOrm()
	//Arreglo
	var arbol []ArbolPlanCuentas
	idplan := strconv.Itoa(plan)

	_, err = o.Raw("select * from financiera.cuenta_contable where id not in (select cuenta_hijo from financiera.estructura_cuentas where cuenta_hijo is not null) and id in (select cuenta_padre from financiera.estructura_cuentas where plan_cuentas=" + idplan + ") order by id;").QueryRows(&arbol)

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
	return arbol, nil
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

//DeleteBranchPlan Funcion para eliminar una rama del arbol
func DeleteBranchPlan(IdCuenta int, IdPlan int) (err error) {
	var hijos []EstructuraCuentas
	o := orm.NewOrm()
	o.Begin()
	if _, err = o.Raw("select * from financiera.estructura_cuentas where plan_cuentas = ? and cuenta_padre = ?", IdPlan, IdCuenta).QueryRows(&hijos); err == nil {
		for _, hijo := range hijos {
			if err = DeleteBranchPlan(hijo.CuentaHijo.Id, IdPlan); err != nil {
				o.Rollback()
				return
			}
		}
	}
	_, err = o.Raw("delete from financiera.estructura_cuentas where plan_cuentas = ? and cuenta_hijo=? ", IdPlan, IdCuenta).Exec()
	if err != nil {
		o.Rollback()
		return
	}
	o.Commit()
	return
}

//AddBranchPlan Funcion que agrega una rama al plan como venga construida
func AddBranchPlan(Rama *ArbolPlanCuentas, IdPlan int) (err error) {
	o := orm.NewOrm()
	o.Begin()
	var v EstructuraCuentas
	v.PlanCuentas = &PlanCuentas{Id: IdPlan}
	v.CuentaPadre = &CuentaContable{Id: Rama.Id}
	if Rama.Hijos != nil {
		for _, cuenta := range *Rama.Hijos {
			p := v
			p.CuentaHijo = &CuentaContable{Id: cuenta.Id}
			var f EstructuraCuentas
			fmt.Println(p.CuentaHijo.Id)
			if err = o.Raw("Select * from financiera.estructura_cuentas where cuenta_hijo=? and plan_cuentas=?", p.CuentaHijo.Id, p.PlanCuentas.Id).QueryRow(&f); err != nil {
				if _, err = o.Insert(&p); err != nil {
					o.Rollback()
					return
				}
			}
			if err = AddBranchPlan(&cuenta, IdPlan); err != nil {
				return
			}
		}
	}
	o.Commit()
	return
}

/*
// Generar arbol de rubros.
func MakeTreePlanCuentas(plan int) (padres []map[string]interface{}, err error) {
	o := orm.NewOrm()
	var m []orm.Params
	idplan := strconv.Itoa(plan)
	//funcion para conseguir los rubros padre. OR (id not in (select DISTINCT rubro_padre from financiera.rubro_rubro))
	_, err = o.Raw(`select cuenta_contable.id as "Id", cuenta_contable.nombre as "Nombre", cuenta_contable.naturaleza as "Naturaleza",
		cuenta_contable.descripcion as "Descripcion", cuenta_contable.codigo as "Codigo", cuenta_contable.cuenta_bancaria as "CuentaBancaria",
		cuenta_contable.nivel_clasificacion_cuenta_contable as "NivelClasificacion" from financiera.cuenta_contable where id not in
		(select cuenta_hijo from financiera.estructura_cuentas where cuenta_hijo is not null)
		 and id in (select cuenta_padre from financiera.estructura_cuentas where plan_cuentas= ? ) order by id;`, idplan).Values(&m)
	if err == nil {
		var res []interface{}
		err = formatdata.FillStruct(m, &res)
		done := make(chan interface{})
		defer close(done)
		resch := utilidades.GenChanInterface(res...)
		var params []interface{}
		params = append(params, idplan)
		charbolcuentas := utilidades.Digest(done, MakeBranchesPlan, resch, params)
		for padre := range charbolcuentas {
			p := padre.(map[string]interface{})
			kl, _ := strconv.Atoi(p["Id"].(string))
			p["Id"] = kl
			kln, _ := strconv.Atoi(p["NivelClasificacion"].(string))
			v := NivelClasificacion{Id: kln}
			if err = o.Read(&v); err == nil {
				p["NivelClasificacion"] = v
			}

			padres = append(padres, p) //tomar valores del canal y agregarlos al array de hijos.
		}
	}
	return
}

func MakeBranchesPlan(forkin interface{}, params ...interface{}) (forkout interface{}) {
	fork := forkin.(map[string]interface{})
	o := orm.NewOrm()
	var m []orm.Params
	var res []interface{}
	//funcion para conseguir los hijos de los rubros padre.
	_, err := o.Raw(`select a.id as "Id", a.nombre as "Nombre", a.naturaleza as "Naturaleza",
		a.descripcion as "Descripcion", a.codigo as "Codigo", a.cuenta_bancaria as "CuentaBancaria",
		a.nivel_clasificacion_cuenta_contable as "NivelClasificacion" from financiera.cuenta_contable a
		left join financiera.estructura_cuentas b on a.id =b.cuenta_hijo
		where b.cuenta_padre= ? and b.plan_cuentas = ? ORDER BY a.id`, fork["Id"], params).Values(&m)
	if err == nil {
		err = formatdata.FillStruct(m, &res)
		var hijos []map[string]interface{}
		done := make(chan interface{})
		defer close(done)
		resch := utilidades.GenChanInterface(res...)
		charbolcuentas := utilidades.Digest(done, MakeBranchesPlan, resch, params)
		for hijo := range charbolcuentas {
			if hijo != nil {
				h := hijo.(map[string]interface{})
				kl, _ := strconv.Atoi(h["Id"].(string))
				h["Id"] = kl
				kln, _ := strconv.Atoi(h["NivelClasificacion"].(string))
				v := NivelClasificacion{Id: kln}
				if err = o.Read(&v); err == nil {
					h["NivelClasificacion"] = v
				}

				hijos = append(hijos, h) //tomar valores del canal y agregarlos al array de hijos.
			}
		}
		fork["Hijos"] = hijos
		return fork
	}
	return
}*/
