package controllers

import (
	"api_financiera/models"

	"github.com/astaxie/beego"
)

// ConceptoConceptoController operations for ConceptoConcepto
type ArbolConceptosController struct {
	beego.Controller
}

// URLMapping ...
func (c *ArbolConceptosController) URLMapping() {
	c.Mapping("MakeTree", c.MakeTree)
}

// MakeTree ...
// @Title Post
// @Description get ConceptoConcepto Arbol
// @Param	body		body 	models.ConceptoConcepto	true		"body for ConceptoConcepto content"
// @Success 201 {int} models.ConceptoConcepto
// @Failure 403 body is empty
// @router / [get]
func (c *ArbolConceptosController) MakeTree() {

	l := models.MakeTreeConcepto()
	//fmt.Println(l)

	c.Data["json"] = l
	//Generera el Json con los datos obtenidos
	c.ServeJSON()

}
