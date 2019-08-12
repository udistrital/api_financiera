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

// oprations for TrConcepto
type TrConceptoController struct {
	beego.Controller
}

func (c *TrConceptoController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("Put", c.Put)
}

// Post ...
// @Title Post
// @Description create Concepto
// @Param	body		body 	models.TrConcepto	true		"body for Concepto content"
// @Success 201 {int} models.Alert
// @Failure 400 the request contains incorrect syntax
// @router / [post]
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
			formatdata.FillStruct(alertdb["Code"], &code)

			alert := models.Alert{Type: "error", Code: "E_" + code, Body: err.Error()}
			if err.Error() == "C92011" {
				alert = models.Alert{Type: "error", Code: "E_" + err.Error(), Body: nil}
			}

			c.Data["json"] = alert
		}
	} else {
		alert := models.Alert{Type: "error", Code: "E_0458", Body: err.Error()}
		c.Data["json"] = alert
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the TrConcepto
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.TrConcepto	true		"body for Concepto content"
// @Success 200 {object} models.Alert
// @Failure 400 the request contains incorrect syntax
// @router /:id [put]
func (c *TrConceptoController) Put() {
	fmt.Printf("actualizando...............................................................")
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	var v models.TrConcepto
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if int(id) == v.Concepto.Id {
			if err := models.UpdateConceptoTr(&v); err == nil {
				alert := models.Alert{Type: "success", Code: "S_542", Body: v.Concepto.Codigo} //codigo de registro exitoso
				c.Ctx.Output.SetStatus(201)
				c.Data["json"] = alert
			} else {
				alertdb := structs.Map(err)
				var code string
				formatdata.FillStruct(alertdb["Code"], &code)
				alert := models.Alert{Type: "error", Code: "E_" + code, Body: err.Error()}
				if err.Error() == "04566" {
					alert = models.Alert{Type: "error", Code: "E_" + err.Error(), Body: nil}
				}
				if err.Error() == "C92011" {
					alert = models.Alert{Type: "error", Code: "E_" + err.Error(), Body: nil}
				}
				c.Data["json"] = alert
			}
		} else {
			alert := models.Alert{Type: "error", Code: "E_C15973", Body: err.Error()}
			c.Data["json"] = alert
		}
	} else {
		alert := models.Alert{Type: "error", Code: "E_0458", Body: err.Error()}
		c.Data["json"] = alert
	}
	c.ServeJSON()
}
