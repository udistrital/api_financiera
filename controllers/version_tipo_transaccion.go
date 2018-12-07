package controllers

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/fatih/structs"
	"github.com/udistrital/api_financiera/models"
	"github.com/udistrital/utils_oas/formatdata"

	"github.com/astaxie/beego"
)

// VersionTipoTransaccionController operations for VersionTipoTransaccion
type VersionTipoTransaccionController struct {
	beego.Controller
}

// URLMapping ...
func (c *VersionTipoTransaccionController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create VersionTipoTransaccion
// @Param	body		body 	models.VersionTipoTransaccion	true		"body for VersionTipoTransaccion content"
// @Success 201 {int} models.VersionTipoTransaccion
// @Failure 403 body is empty
// @router / [post]
func (c *VersionTipoTransaccionController) Post() {
	var v models.VersionTipoTransaccion
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddVersionTipoTransaccion(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = models.Alert{Type: "success", Code: "S_543", Body: v}
		} else {
			alertdb := structs.Map(err)
			var code string
			formatdata.FillStruct(alertdb["Code"], &code)
			alert := models.Alert{Type: "error", Code: "E_" + code, Body: err}
			c.Data["json"] = alert
		}
	} else {
		alert := models.Alert{Type: "error", Code: "E_0458", Body: err}
		c.Data["json"] = alert
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get VersionTipoTransaccion by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.VersionTipoTransaccion
// @Failure 403 :id is empty
// @router /:id [get]
func (c *VersionTipoTransaccionController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetVersionTipoTransaccionById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get VersionTipoTransaccion
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.VersionTipoTransaccion
// @Failure 403
// @router / [get]
func (c *VersionTipoTransaccionController) GetAll() {
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

	l, err := models.GetAllVersionTipoTransaccion(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the VersionTipoTransaccion
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.VersionTipoTransaccion	true		"body for VersionTipoTransaccion content"
// @Success 200 {object} models.VersionTipoTransaccion
// @Failure 403 :id is not int
// @router /:id [put]
func (c *VersionTipoTransaccionController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.VersionTipoTransaccion{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateVersionTipoTransaccionById(&v); err == nil {
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
// @Description delete the VersionTipoTransaccion
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *VersionTipoTransaccionController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	beego.Error("parametro version ", idStr)
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteVersionTipoTransaccion(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// VersionTipoTransaccionController ...
// @Title VersionTipoTransaccionController
// @Description get Version Tipo Transaccion
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.VersionTipoTransaccion
// @Failure 403
// @router /GetVersionesTipoNumber/ [get]
func (c *VersionTipoTransaccionController) GetVersionesTipoNumber() {
	var query = make(map[string]string)
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
	l, err := models.GetRecordsNumberVersionTipo(query)
	if err != nil {
		alertdb := structs.Map(err)
		c.Data["json"] = &models.Alert{Code: "E_" + alertdb["Code"].(string), Type: "error", Body: err.Error()}
	} else {
		c.Data["json"] = &models.Alert{Code: "E_S545", Type: "success", Body: l}
	}
	c.ServeJSON()
}

// GetVersionToType ...
// @Title GetVersionToType
// @Description get last version for TipoTransaccion
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.VersionTipoTransaccion
// @Failure 403 :id is empty
// @router /GetVersionToType/:idTipo [get]
func (c *VersionTipoTransaccionController) GetVersionToType() {
	idStr := c.Ctx.Input.Param(":idTipo")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetVersionToTipoTransaccion(id)
	if err != nil {
		alertdb := structs.Map(err)
		var code string
		formatdata.FillStruct(alertdb["Code"], &code)
		alert := models.Alert{Type: "error", Code: "E_" + code, Body: err}
		c.Data["json"] = alert
	} else {
		c.Data["json"] = models.Alert{Type: "success", Code: "S_543", Body: v}
	}
	c.ServeJSON()
}

// GetVersionInEspecifiedDate ...
// @Title GetVersionInEspecifiedDate
// @Description get all versions in a especified set of dates
// @Param	fechaInicio	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fechaFin	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	tipo	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Success 200 {object} models.VersionTipoTransaccion
// @Failure 403 :id is empty
// @router /GetVersionToType/:idTipo [get]
func (c *VersionTipoTransaccionController) GetVersionInEspecifiedDate() {
	var fechaInicio string
	var fechaFin string
	var tipo int

	if v := c.GetString("fechaInicio"); v != "" {
		fechaInicio = v
	}

	if v := c.GetString("fechaFin"); v != "" {
		fechaFin = v
	}

	if v, err := c.GetInt("tipo"); err == nil {
		tipo = v
	}

	v, err := models.GetVersionInEspecifiedDate(fechaInicio, fechaFin, tipo)
	if err != nil {
		alertdb := structs.Map(err)
		var code string
		formatdata.FillStruct(alertdb["Code"], &code)
		alert := models.Alert{Type: "error", Code: "E_" + code, Body: err}
		c.Data["json"] = alert
	} else {
		c.Data["json"] = models.Alert{Type: "success", Code: "S_543", Body: v}
	}
	c.ServeJSON()
}
