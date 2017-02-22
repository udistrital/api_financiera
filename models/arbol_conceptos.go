package models

import (
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
)

type ConceptoArbol struct {
	Id              int           `orm:"column(id);pk;auto"`
	Codigo          string        `orm:"column(codigo)"`
	Nombre          string        `orm:"column(nombre)"`
	FechaCreacion   time.Time     `orm:"column(fecha_creacion);type(date)"`
	Cabeza          bool          `orm:"column(cabeza)"`
	FechaExpiracion time.Time     `orm:"column(fecha_expiracion);type(date);null"`
	Descripcion     string        `orm:"column(descripcion);null"`
	TipoConcepto    *TipoConcepto `orm:"column(tipo_concepto);rel(fk)"`
	Rubro           *Rubro        `orm:"column(rubro);rel(fk);null;"`
	Hijos           *[]ConceptoArbol
}

//MakeTreeConcepto make a tree of a full estructure of concepto and returns error if
// the record to be make fail

func MakeTreeConcepto() (a []ConceptoArbol) {
	o := orm.NewOrm()
	//Arreglo
	var arbol []ConceptoArbol
	_, err := o.Raw("select * from financiera.concepto where id not in (select concepto_hijo from financiera.concepto_concepto) order by id;").QueryRows(&arbol)

	/*for _, concepto := range arbol {
		var l []Concepto
		o.QueryTable(new(Concepto)).Filter("Id", concepto.Id).RelatedSel().All(&l)
		concepto.TipoConcepto = l[0].TipoConcepto
		fmt.Print(l[0].TipoConcepto)
	}*/

	if err == nil {
		//fmt.Println("Menus padre encontrados: ", num)
		//For para que recorra los Ids en busca de hijos
		for i := 0; i < len(arbol); i++ {
			var l Concepto
			o.QueryTable(&l).Filter("Id", arbol[i].Id).RelatedSel().All(&l)
			arbol[i].TipoConcepto = l.TipoConcepto
			arbol[i].Rubro = l.Rubro
			//fmt.Print(l.TipoConcepto)

			//Me verifica que los Id tengan hijos
			MakeBranches(&arbol[i])
		}

	}
	return arbol
}

//Función que construye los Submenús
func MakeBranches(Padre *ConceptoArbol) (a []ConceptoArbol) {
	o := orm.NewOrm()
	//Conversión de entero a string
	idpadre := strconv.Itoa(Padre.Id)

	//Arreglo
	var arbol []ConceptoArbol

	_, err := o.Raw("select a.* from financiera.concepto a left join financiera.concepto_concepto b on a.id =b.concepto_hijo where b.concepto_padre=" + idpadre + " ORDER BY a.id").QueryRows(&arbol)
	//Condicional si el error es nulo
	if err == nil {
		//fmt.Println("Menus Hijos encontrados: ", num)

		//Llena el elemento Opciones en la estructura del menú padre
		Padre.Hijos = &arbol

		//For que recorre el subMenu en busca de hijos (Recursiva)
		for i := 0; i < len(arbol); i++ {
			var l Concepto
			o.QueryTable(&l).Filter("Id", arbol[i].Id).RelatedSel().All(&l)
			arbol[i].TipoConcepto = l.TipoConcepto
			//Me verifica que los Id tengan hijos
			MakeBranches(&arbol[i])
		}
	}
	return arbol
}
