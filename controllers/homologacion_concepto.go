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

// HomologacionConceptoController oprations for HomologacionConcepto
type HomologacionConceptoController struct {
	beego.Controller
}

// URLMapping ...
func (c *HomologacionConceptoController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("RegistrarHomologacionConcepto", c.RegistrarHomologacionConcepto)
}

// Post ...
// @Title Post
// @Description create HomologacionConcepto
// @Param	body		body 	models.HomologacionConcepto	true		"body for HomologacionConcepto content"
// @Success 201 {int} models.HomologacionConcepto
// @Failure 400 the request contains incorrect syntax
// @router / [post]
func (c *HomologacionConceptoController) Post() {
	var v models.HomologacionConcepto
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddHomologacionConcepto(&v); err == nil {
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
// @Description get HomologacionConcepto by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.HomologacionConcepto
// @Failure 404 not found resource
// @router /:id [get]
func (c *HomologacionConceptoController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetHomologacionConceptoById(id)
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
// @Description get HomologacionConcepto
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.HomologacionConcepto
// @Failure 404 not found resource
// @router / [get]
func (c *HomologacionConceptoController) GetAll() {
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

	l, err := models.GetAllHomologacionConcepto(query, fields, sortby, order, offset, limit)
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
// @Description update the HomologacionConcepto
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.HomologacionConcepto	true		"body for HomologacionConcepto content"
// @Success 200 {object} models.HomologacionConcepto
// @Failure 400 the request contains incorrect syntax
// @router /:id [put]
func (c *HomologacionConceptoController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.HomologacionConcepto{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateHomologacionConceptoById(&v); err == nil {
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
// @Description delete the HomologacionConcepto
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 404 not found resource
// @router /:id [delete]
func (c *HomologacionConceptoController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteHomologacionConcepto(id); err == nil {
		c.Data["json"] = map[string]interface{}{"Id": id}
	} else {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("404")
	}
	c.ServeJSON()
}

// RegistrarHomologacionConcepto ...
// @Title RegistrarHomologacionConcepto
// @Description Registrar homologacion_concepto y  de concepto_tesoral_facultad_proyecto
// @Param	body		body 	models.homologacion_concepto	true		"body for homologacion_concepto content"
// @Success 201 {int} models.OrdenPago
// @Failure 403 body is empty
// @router RegistrarHomologacionConcepto [post]
func (c *HomologacionConceptoController) RegistrarHomologacionConcepto() {
	var v interface{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		m := v.(map[string]interface{})
		mensaje := models.RegistrarHomologacionConcepto(m)
		if mensaje.Type != "success" {
			c.Data["json"] = mensaje
		} else {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = mensaje
		}
	} else {
		c.Data["json"] = err
	}
	c.ServeJSON()
}

// ActualizarHomologacionConcepto ...
// @Title ActualizarHomologacionConcepto
// @Description Actualiza homologacion_concepto y  de concepto_tesoral_facultad_proyecto
// @Param	body		body 	models.homologacion_concepto	true		"body for homologacion_concepto content"
// @Success 201 {int} models.OrdenPago
// @Failure 403 body is empty
// @router ActualizarHomologacionConcepto [post]
func (c *HomologacionConceptoController) ActualizarHomologacionConcepto() {
	var v interface{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		m := v.(map[string]interface{})
		mensaje := models.ActualizarHomologacionConcepto(m)
		if mensaje.Type != "success" {
			c.Data["json"] = mensaje
		} else {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = mensaje
		}
	} else {
		c.Data["json"] = err
	}
	c.ServeJSON()
}
