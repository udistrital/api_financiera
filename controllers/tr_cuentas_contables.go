package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/fatih/structs"
	"github.com/udistrital/api_financiera/models"
	"github.com/udistrital/api_financiera/utilidades"

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
			utilidades.FillStruct(alertdb["Code"], &code)
			alert := models.Alert{Type: "error", Code: "E_" + code, Body: err.Error()}
			c.Data["json"] = alert
		}
	} else {
		fmt.Println(err)
		c.Data["json"] = models.Alert{Type: "error", Code: "E_0458", Body: err.Error()}
	}
	c.ServeJSON()
}
