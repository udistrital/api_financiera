package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/udistrital/api_financiera/models"

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

		if alerta, err := models.AddTransaccionConcepto(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = alerta

		} else {
			c.Data["json"] = alerta
		}
	} else {
		fmt.Println(err)
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
