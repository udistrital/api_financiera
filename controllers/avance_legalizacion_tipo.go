package controllers

import (
	"encoding/json"
	"errors"
	"github.com/udistrital/api_financiera/models"
	"strconv"
	"strings"
	"github.com/udistrital/utils_oas/formatdata"
	"github.com/fatih/structs"

	"github.com/astaxie/beego"
)

// AvanceLegalizacionTipoController operations for AvanceLegalizacionTipo
type AvanceLegalizacionTipoController struct {
	beego.Controller
}

// URLMapping ...
func (c *AvanceLegalizacionTipoController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create AvanceLegalizacionTipo
// @Param	body		body 	models.AvanceLegalizacionTipo	true		"body for AvanceLegalizacionTipo content"
// @Success 201 {int} models.AvanceLegalizacionTipo
// @Failure 403 body is empty
// @router / [post]
func (c *AvanceLegalizacionTipoController) Post() {
	var v models.AvanceLegalizacionTipo
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddAvanceLegalizacionTipo(&v); err == nil {
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

// GetOne ...
// @Title Get One
// @Description get AvanceLegalizacionTipo by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.AvanceLegalizacionTipo
// @Failure 403 :id is empty
// @router /:id [get]
func (c *AvanceLegalizacionTipoController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetAvanceLegalizacionTipoById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get AvanceLegalizacionTipo
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.AvanceLegalizacionTipo
// @Failure 403
// @router / [get]
func (c *AvanceLegalizacionTipoController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllAvanceLegalizacionTipo(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the AvanceLegalizacionTipo
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.AvanceLegalizacionTipo	true		"body for AvanceLegalizacionTipo content"
// @Success 200 {object} models.AvanceLegalizacionTipo
// @Failure 403 :id is not int
// @router /:id [put]
func (c *AvanceLegalizacionTipoController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.AvanceLegalizacionTipo{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateAvanceLegalizacionTipoById(&v); err == nil {
			c.Data["json"] = "OK"
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the AvanceLegalizacionTipo
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *AvanceLegalizacionTipoController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteAvanceLegalizacionTipo(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// AddAvanceLegalizacionTipo ...
// @Title AddAvanceLegalizacionTipo
// @Description create AvanceLegalizacion
// @Param	body		body 	models.AvanceLegalizacionTipo	true		"body for AvanceLegalizacionTipo content"
// @Success 201 {int} models.AvanceLegalizacionTipo
// @Failure 403 body is empty
// @router /AddEntireAvanceLegalizacionTipo [post]
func (c *AvanceLegalizacionTipoController) AddEntireAvanceLegalizacionTipo () {
	defer c.ServeJSON();
	var v map[string]interface{}
	var alerta interface{}
	var valorLegalizado float64
	var valorAvance float64
	var valorLegalizacion float64
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		valorLegalizado = v["ValorLegalizadoAvance"].(float64)
		valorAvance = v["ValorTotalAvance"].(float64)
		valorLegalizacion = v["Valor"].(float64)
		if valorLegalizado + valorLegalizacion <= valorAvance{
			if alerta, err = models.AddAllAvanceLegalizacionTipo(v); err == nil {
				c.Ctx.Output.SetStatus(201)
				c.Data["json"] = models.Alert{Type: "success", Code: "S_543", Body: alerta}
				} else {
					alertdb := structs.Map(err)
					var code string
					formatdata.FillStruct(alertdb["Code"], &code)
					alert := models.Alert{Type: "error", Code: "E_" + code, Body: err}
					c.Data["json"] = alert
					c.Ctx.Output.SetStatus(500)
				}
		}else{
			alert := models.Alert{Type: "error", Code: "E_LA0001", Body: "Bad Request"}
			c.Data["json"] = alert
			c.Ctx.Output.SetStatus(400)
		}
	} else {
		alert := models.Alert{Type: "error", Code: "E_0458", Body: err}
		c.Data["json"] = alert
		c.Ctx.Output.SetStatus(400)
	}
}

// GetLegalizationValue ...
// @Title Get One
// @Description get sum from all legalization advance payment including refunds
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.AvanceLegalizacion
// @Failure 403 :id is empty
// @router /GetLegalizationValue/:id [get]
func (c *AvanceLegalizacionTipoController) GetLegalizationValue() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetLegalizationValue(id)
	if err != nil {
		alertdb := structs.Map(err)
		var code string
		formatdata.FillStruct(alertdb["Code"], &code)
		alert := models.Alert{Type: "error", Code: "E_" + code, Body: err}
		c.Data["json"] = alert
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

//GetTaxesMovLegalization
// @Title GetTaxesLegalization
// @Description get taxes and movs for legalizationTipo given an id
// @Param	noTipoDoc	query	string 	true		"param for TipoDocumentoAfectante"
// @Param	idLegTipo	query 	string	true	"param for id legalizacion table"
// @Success 200 {object} models.Legalizacion_avance
// @Failure 403 :id is empty
// @router /GetTaxesMovsLegalization [get]
func (c *AvanceLegalizacionTipoController) GetTaxesMovsLegalization() {
	defer c.ServeJSON();
 var legTipo int
 var tipoDoc int
 respuesta:=make(map[string]interface{})
	if v, err := c.GetInt("idLegTipo"); err == nil {
		legTipo = v
	}
	if v, err := c.GetInt("noTipoDoc"); err == nil {
		tipoDoc = v
	}
	v, err := models.GetTaxesLegalization(legTipo,tipoDoc)
	if err != nil {
		alertdb := structs.Map(err)
		var code string
		formatdata.FillStruct(alertdb["Code"], &code)
		alert := models.Alert{Type: "error", Code: "E_" + code, Body: err}
		respuesta["impuestos"]  = alert
	} else {
		respuesta["impuestos"] = v
	}
	v, err = models.GetMovsLegalization(legTipo,tipoDoc)
	if err != nil {
		alertdb := structs.Map(err)
		var code string
		formatdata.FillStruct(alertdb["Code"], &code)
		alert := models.Alert{Type: "error", Code: "E_" + code, Body: err}
		respuesta["movimientos"] = alert
	} else {
		respuesta["movimientos"] = v
	}
	c.Data["json"]=respuesta
}
