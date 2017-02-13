package controllers

import (
	"api_financiera/models"
	"encoding/json"

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

		if _, err = models.AddTransaccionConcepto(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = v

		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
