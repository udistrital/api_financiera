package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/fatih/structs"
	"github.com/udistrital/api_financiera/models"
	"github.com/udistrital/utils_oas/formatdata"

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
	id, err := strconv.Atoi(idStr)
    if err == nil {
    	if _, err := os.Stat("PlanCuentasTreeId" + idStr + ".json"); os.IsNotExist(err) {
			l, err := models.MakeTreePlanCuentas(id)
			if err != nil {
				alertdb := structs.Map(err)
				var code string
				formatdata.FillStruct(alertdb["Code"], &code)
				alert := models.Alert{Type: "error", Code: "E_" + code, Body: err.Error()}
				c.Data["json"] = alert
			} else {
				rankingsJson, _ := json.Marshal(l)
				err = ioutil.WriteFile("PlanCuentasTreeId" + idStr + ".json", rankingsJson, 0644)
				fmt.Println("err ", err)	
				c.Data["json"] = l
			}
		} else {
			data, _ := ioutil.ReadFile("PlanCuentasTreeId" + idStr + ".json")
			var l interface {}
			err = json.Unmarshal(data, &l)
			// fmt.Println("read from file PlanCuentas")
			c.Data["json"] = l
		}		
	} else {
		e := models.Alert{Type: "error", Code: "E_0458", Body: err.Error()}
		c.Data["json"] = e

	}	
		
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
	if err := models.DeleteBranchPlan(idCuenta, idPlan); err == nil {
		alert := models.Alert{Type: "success", Code: "S_554", Body: nil}
		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = alert
	} else {
		alertdb := structs.Map(err)
		var code string
		formatdata.FillStruct(alertdb["Code"], &code)
		alert := models.Alert{Type: "error", Code: "E_" + code, Body: err.Error()}
		c.Data["json"] = alert
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
		if err = models.AddBranchPlan(&v, idPlan); err == nil {
			alert := models.Alert{Type: "success", Code: "S_543", Body: nil}
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = alert
		} else {
			alertdb := structs.Map(err)
			var code string
			formatdata.FillStruct(alertdb["Code"], &code)
			alert := models.Alert{Type: "error", Code: "E_" + code, Body: err.Error()}
			c.Data["json"] = alert
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
