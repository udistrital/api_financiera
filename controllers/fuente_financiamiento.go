package controllers

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/udistrital/api_financiera/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

// FuenteFinanciamientoController operations for FuenteFinanciamiento
type FuenteFinanciamientoController struct {
	beego.Controller
}

// URLMapping ...
func (c *FuenteFinanciamientoController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create FuenteFinanciamiento
// @Param	body		body 	models.FuenteFinanciamiento	true		"body for FuenteFinanciamiento content"
// @Success 201 {int} models.FuenteFinanciamiento
// @Failure 400 the request contains incorrect syntax
// @router / [post]
func (c *FuenteFinanciamientoController) Post() {
	var v models.FuenteFinanciamiento
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddFuenteFinanciamiento(&v); err == nil {
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
// @Description get FuenteFinanciamiento by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.FuenteFinanciamiento
// @Failure 404 not found resource
// @router /:id [get]
func (c *FuenteFinanciamientoController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetFuenteFinanciamientoById(id)
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
// @Description get FuenteFinanciamiento
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.FuenteFinanciamiento
// @Failure 404 not found resource
// @router / [get]
func (c *FuenteFinanciamientoController) GetAll() {
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

	l, err := models.GetAllFuenteFinanciamiento(query, fields, sortby, order, offset, limit)
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
// @Description update the FuenteFinanciamiento
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.FuenteFinanciamiento	true		"body for FuenteFinanciamiento content"
// @Success 200 {object} models.FuenteFinanciamiento
// @Failure 400 the request contains incorrect syntax
// @router /:id [put]
func (c *FuenteFinanciamientoController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.FuenteFinanciamiento{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateFuenteFinanciamientoById(&v); err == nil {
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
// @Description delete the FuenteFinanciamiento
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 404 not found resource
// @router /:id [delete]
func (c *FuenteFinanciamientoController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteFuenteFinanciamiento(id); err == nil {
		c.Data["json"] = map[string]interface{}{"Id": id}
	} else {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("404")
	}
	c.ServeJSON()
}

// RegistrarFuenteFinanciamientoTr ...
// @Title RegistrarFuenteFinanciamientoTr
// @Description create FuenteFinanciamiento with Tr
// @Param	body		body 	models.FuenteFinanciamiento	true		"body for FuenteFinanciamiento content"
// @Success 201 {int} models.FuenteFinanciamiento
// @Failure 403 body is empty
// @router /RegistrarFuenteFinanciamientoTr [post]
func (c *FuenteFinanciamientoController) RegistrarFuenteFinanciamientoTr() {
	var v map[string]interface{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if data, err := models.AddFuenteFinanciamientoTr(v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = data
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// MovimientoFuenteFinanciamientoTr ...
// @Title MovimientoFuenteFinanciamientoTr
// @Description create FuenteFinanciamiento with Tr
// @Param	body		body 	models.FuenteFinanciamiento	true		"body for FuenteFinanciamiento content"
// @Success 201 {int} models.FuenteFinanciamiento
// @Failure 403 body is empty
// @router /MovimientoFuenteFinanciamientoTr [post]
func (c *FuenteFinanciamientoController) MovimientoFuenteFinanciamientoTr() {
	var v []map[string]interface{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if data, err := models.AddMovimientoFuenteFinanciamientoTr(v); err == nil {
			c.Ctx.Output.SetStatus(201)
			alert := models.Alert{Type: "success", Code: "S_543", Body: data}
			c.Data["json"] = alert
		} else {
			alert := models.Alert{Type: "error", Code: "E_0458", Body: err}
			c.Data["json"] = alert
		}
	} else {
		alert := models.Alert{Type: "error", Code: "E_0458", Body: err}
		c.Data["json"] = alert
	}
	c.ServeJSON()
}

// DeleteMovimientoFuenteFinanciamientoTr ...
// @Title DeleteMovimientoFuenteFinanciamientoTr
// @Description delete FuenteFinanciamiento with Tr
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 201 {int} models.FuenteFinanciamiento
// @Failure 403 body is empty
// @router /DeleteMovimientoFuenteFinanciamientoTr/:id [delete]
func (c *FuenteFinanciamientoController) DeleteMovimientoFuenteFinanciamientoTr() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteMovimientoFuenteFinanciamientoTr(id); err == nil {
		c.Ctx.Output.SetStatus(201)
		alert := models.Alert{Type: "success", Code: "S_543", Body: err}
		c.Data["json"] = alert
	} else {
		alert := models.Alert{Type: "error", Code: "E_0458", Body: err}
		c.Data["json"] = alert
	}

	c.ServeJSON()
}

// DeleteModificacionFuenteFinanciamientoTr ...
// @Title DeleteModificacionFuenteFinanciamientoTr
// @Description Delete ModificacionFuenteFinanciamiento with Tr
// @Param	body		body 	models.FuenteFinanciamientoApropiacion	true		"body for FuenteFinanciamiento content"
// @Success 201 {int} models.FuenteFinanciamiento
// @Failure 403 body is empty
// @router /DeleteModificacionFuenteFinanciamientoTr [post]
func (c *FuenteFinanciamientoController) DeleteModificacionFuenteFinanciamientoTr() {
	var v []map[string]interface{}

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.DeleteModificacionFuenteFinanciamiento(v); err == nil {
			c.Ctx.Output.SetStatus(201)
			alert := models.Alert{Type: "success", Code: "S_543", Body: nil}
			c.Data["json"] = alert
		} else {
			alert := models.Alert{Type: "error", Code: "E_0458", Body: err}
			c.Data["json"] = alert
		}
	} else {
		alert := models.Alert{Type: "error", Code: "E_0458", Body: err}
		c.Data["json"] = alert
	}
	c.ServeJSON()
}
