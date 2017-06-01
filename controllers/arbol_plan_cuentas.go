package controllers

import (
	"encoding/json"
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
	c.Mapping("DeleteBranch", c.DeleteBranch)
	c.Mapping("Post", c.Post)
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

// DeleteBranchPlan ...
// @Title DeleteBranchPlan
// @Description Elimina una rama del plan
// @Param	idCuenta		path 	string	true		"Id de la cuenta"
// @Param	idPlan		path 	string	true		"Id del plan"
// @Success 200 {string} delete success!
// @Failure 403 idCuenta or idPlan is empty
// @router /:idCuenta/:idPlan [delete]
func (c *ArbolPlanCuentasController) DeleteBranch() {
	idCuentaStr := c.Ctx.Input.Param(":idCuenta")
	idPlanStr := c.Ctx.Input.Param(":idPlan")
	idCuenta, _ := strconv.Atoi(idCuentaStr)
	idPlan, _ := strconv.Atoi(idPlanStr)
	if msg, err := models.DeleteBranchPlan(idCuenta, idPlan); err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = msg
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Post ...
// @Title Post
// @Description create Rama arbol plan
// @Param	body		body 	models.CategoriaIva	true		"body for CategoriaIva content"
// @Param	idPlan		path 	string	true		"Id del plan"
// @Success 201 {int} models.CategoriaIva
// @Failure 403 body is empty
// @router /:idPlan [post]
func (c *ArbolPlanCuentasController) Post() {
	idPlanStr := c.Ctx.Input.Param(":idPlan")
	idPlan, _ := strconv.Atoi(idPlanStr)
	var v models.ArbolPlanCuentas
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if msg, err := models.AddBranchPlan(&v, idPlan); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = msg
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
