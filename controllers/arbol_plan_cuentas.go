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
// @Description construye y muestra la estructura de cuentas en un plan de cuentas
// @Param	id		path 	string	true		"Id del plan de cuentas a construirse"
// @Success 201 {int} models.ArbolPlanCuentas
// @Failure 403 body is empty
// @router /:id [get]
func (c *ArbolPlanCuentasController) MakeTreeCuentas() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	l := models.MakeTreePlanCuentas(id)
	c.Data["json"] = l
	//Generera el Json con los datos obtenidos
	c.ServeJSON()
}
