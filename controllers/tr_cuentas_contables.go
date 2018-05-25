package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/fatih/structs"
	"github.com/udistrital/api_financiera/models"
	"github.com/udistrital/utils_oas/formatdata"

	"github.com/astaxie/beego"
)

// TrCuentasContablesController operations for plan_cuentas
type TrCuentasContablesController struct {
	beego.Controller
}

// URLMapping ...
func (c *TrCuentasContablesController) URLMapping() {
	c.Mapping("Post", c.Post)
}

// Post ...
// @Title Post
// @Description Post de cuentas contables mediante la transaccion que las crea y ubica en el plan de cuentas
// @Param	body	body 	models.TrCuentaContable	true "body para la transaccion"
// @Success 201 {int} models.TrCuentaContable
// @Failure 403 body is empty
// @router / [post]
func (c *TrCuentasContablesController) Post() {

	var v models.TrCuentaContable
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err = models.AddTransaccionCuentaContable(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = models.Alert{Type: "success", Code: "S_542", Body: v.Cuenta.Codigo}
		} else {
			alertdb := structs.Map(err)
			var code string
			formatdata.FillStruct(alertdb["Code"], &code)
			alert := models.Alert{Type: "error", Code: "E_" + code, Body: err.Error()}
			c.Data["json"] = alert
		}
	} else {
		fmt.Println(err)
		c.Data["json"] = models.Alert{Type: "error", Code: "E_0458", Body: err.Error()}
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the TrCuentaContable
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.TrCuentaContable	true		"body for TrCuentaContable content"
// @Success 200 {object} models.TrCuentaContable
// @Failure 403 :id is not int
// @router /:id [put]
func (c *TrCuentasContablesController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	fmt.Print(id)
	var v models.TrCuentaContable
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err, cod := models.UpdateTrCuentaContable(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = models.Alert{Type: "success", Code: "S_542", Body: cod}
		} else {
			alertdb := structs.Map(err)
			var code string
			formatdata.FillStruct(alertdb["Code"], &code)
			alert := models.Alert{Type: "error", Code: "E_" + code, Body: err.Error()}
			if err.Error() == "04566" {
				alert = models.Alert{Type: "error", Code: "E_" + err.Error(), Body: nil}
			}
			c.Data["json"] = alert
		}
	} else {
		c.Data["json"] = models.Alert{Type: "error", Code: "E_0458", Body: err.Error()}
	}
	c.ServeJSON()
}
