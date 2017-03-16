package controllers

import (
	"strconv"

	"github.com/udistrital/api_financiera/models"

	"github.com/astaxie/beego"
)

// ArbolPlanCuentasController operations for estructura_cuentas
type ArbolPlanCuentasController struct {
	beego.Controller
}

// URLMapping ...
func (c *ArbolPlanCuentasController) URLMapping() {
	c.Mapping("MakeTreeCuentas", c.MakeTreeCuentas)
}

// MakeTreeCuentas ...
// @Title MakeTreeCuentas
// @Description get Arbol of Cuentas in Plan_Cuentas
// @Param	body		body 	models.EstructuraCuentas	true		"body for EstructuraCuentas content"
// @Success 201 {int} models.ArbolPlanCuentas
// @Failure 403 body is empty
// @router /:id [get]
func (c *ArbolPlanCuentasController) MakeTreeCuentas() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	l := models.MakeTreePlanCuentas(id)
	//fmt.Println(l)
	c.Data["json"] = l
	//Generera el Json con los datos obtenidos
	c.ServeJSON()
}
