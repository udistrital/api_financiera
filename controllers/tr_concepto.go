package controllers

import (
	"encoding/json"

	"github.com/fatih/structs"
	"github.com/udistrital/api_financiera/models"
	"github.com/udistrital/api_financiera/utilidades"

	"github.com/astaxie/beego"
)

// oprations for TrConcepto
type TrConceptoController struct {
	beego.Controller
}

func (c *TrConceptoController) URLMapping() {
	c.Mapping("Post", c.Post)
}

func (c *TrConceptoController) Post() {

	var v models.TrConcepto
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.AddTransaccionConcepto(&v); err == nil {
			alert := models.Alert{Type: "success", Code: "S_542", Body: v.Concepto.Codigo} //codigo de registro exitoso
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = alert
		} else {
			alertdb := structs.Map(err)
			var code string
			utilidades.FillStruct(alertdb["Code"], &code)
			alert := models.Alert{Type: "error", Code: "E_" + code, Body: err.Error()}
			c.Data["json"] = alert
		}
	} else {
		alert := models.Alert{Type: "error", Code: "E_0458", Body: err.Error()}
		c.Data["json"] = alert
	}
	c.ServeJSON()
}
