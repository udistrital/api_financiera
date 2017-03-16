package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/udistrital/api_financiera/models"

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
// @Description Post Cuentas in Plan_Cuentas
// @Param	body		body 	models.TrCuentaContable	true		"body for TrCuentaContable content"
// @Success 201 {int} models.TrCuentaContable
// @Failure 403 body is empty
// @router / [post]
func (c *TrCuentasContablesController) Post() {

	var v models.TrCuentaContable
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {

		if _, err = models.AddTransaccionCuentaContable(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = v

		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		fmt.Println(err)
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
