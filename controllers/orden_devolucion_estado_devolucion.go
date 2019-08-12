package controllers

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/udistrital/api_financiera/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/fatih/structs"
	"github.com/udistrital/utils_oas/formatdata"
)

// OrdenDevolucionEstadoDevolucionController operations for OrdenDevolucionEstadoDevolucion
type OrdenDevolucionEstadoDevolucionController struct {
	beego.Controller
}

// URLMapping ...
func (c *OrdenDevolucionEstadoDevolucionController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create OrdenDevolucionEstadoDevolucion
// @Param	body		body 	models.OrdenDevolucionEstadoDevolucion	true		"body for OrdenDevolucionEstadoDevolucion content"
// @Success 201 {int} models.OrdenDevolucionEstadoDevolucion
// @Failure 400 the request contains incorrect syntax
// @router / [post]
func (c *OrdenDevolucionEstadoDevolucionController) Post() {
	var v models.OrdenDevolucionEstadoDevolucion
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddOrdenDevolucionEstadoDevolucion(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = v
		} else {
			logs.Error(err)
			//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
			c.Data["system"] = err
			c.Abort("400")
		}
	} else {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("400")
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get OrdenDevolucionEstadoDevolucion by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.OrdenDevolucionEstadoDevolucion
// @Failure 404 not found resource
// @router /:id [get]
func (c *OrdenDevolucionEstadoDevolucionController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetOrdenDevolucionEstadoDevolucionById(id)
	if err != nil {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("404")
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get OrdenDevolucionEstadoDevolucion
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.OrdenDevolucionEstadoDevolucion
// @Failure 404 not found resource
// @router / [get]
func (c *OrdenDevolucionEstadoDevolucionController) GetAll() {
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

	l, err := models.GetAllOrdenDevolucionEstadoDevolucion(query, fields, sortby, order, offset, limit)
	if err != nil {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("404")
	} else {
		if l == nil {
			l = append(l, map[string]interface{}{})
		}
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the OrdenDevolucionEstadoDevolucion
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.OrdenDevolucionEstadoDevolucion	true		"body for OrdenDevolucionEstadoDevolucion content"
// @Success 200 {object} models.OrdenDevolucionEstadoDevolucion
// @Failure 400 the request contains incorrect syntax
// @router /:id [put]
func (c *OrdenDevolucionEstadoDevolucionController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.OrdenDevolucionEstadoDevolucion{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateOrdenDevolucionEstadoDevolucionById(&v); err == nil {
			c.Data["json"] = v
		} else {
			logs.Error(err)
			//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
			c.Data["system"] = err
			c.Abort("400")
		}
	} else {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("400")
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the OrdenDevolucionEstadoDevolucion
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 404 not found resource
// @router /:id [delete]
func (c *OrdenDevolucionEstadoDevolucionController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteOrdenDevolucionEstadoDevolucion(id); err == nil {
		c.Data["json"] = map[string]interface{}{"Id": id}
	} else {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("404")
	}
	c.ServeJSON()
}

// Post ...
// @Title AddEstadoOrdenDevol
// @Description Insert a record to modify status to a devolution
//@Param	body		body 	models.OrdenDevolucionEstadoDevolucion	true		"body for OrdenDevolucionEstadoDevolucion content"
// @Success 200 {object} models.OrdenDevolucionEstadoDevolucion
// @Failure 403 body is empty
// @router /AddEstadoOrdenDevol [post]
func (c *OrdenDevolucionEstadoDevolucionController) AddEstadoOrdenDevol() {
	var request map[string]interface{}
	var code string
	defer c.ServeJSON()
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &request); err == nil {

		if estado, err := models.AddEstadoOrden(request); err == nil {
			alert := models.Alert{Type: "success", Code: "S_543", Body: estado}
			c.Data["json"] = alert
		} else {
			beego.Info(err.Error())
			alertdb := structs.Map(err)
			formatdata.FillStruct(alertdb["Code"], &code)
			alert := models.Alert{Type: "error", Code: "E_" + code, Body: err.Error()}
			c.Data["json"] = alert
		}
	} else {
		beego.Info(err.Error())
		alert := models.Alert{Type: "error", Code: "E_0458" + code, Body: err.Error()}
		c.Data["json"] = alert
	}
}
