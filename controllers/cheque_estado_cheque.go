package controllers

import (
	"encoding/json"
	"errors"
	"github.com/udistrital/api_financiera/models"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/fatih/structs"
	"github.com/udistrital/utils_oas/formatdata"
)

// ChequeEstadoChequeController operations for ChequeEstadoCheque
type ChequeEstadoChequeController struct {
	beego.Controller
}

// URLMapping ...
func (c *ChequeEstadoChequeController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create ChequeEstadoCheque
// @Param	body		body 	models.ChequeEstadoCheque	true		"body for ChequeEstadoCheque content"
// @Success 201 {int} models.ChequeEstadoCheque
// @Failure 403 body is empty
// @router / [post]
func (c *ChequeEstadoChequeController) Post() {
	var v models.ChequeEstadoCheque
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddChequeEstadoCheque(&v); err == nil {
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
// @Description get ChequeEstadoCheque by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.ChequeEstadoCheque
// @Failure 403 :id is empty
// @router /:id [get]
func (c *ChequeEstadoChequeController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetChequeEstadoChequeById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get ChequeEstadoCheque
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.ChequeEstadoCheque
// @Failure 403
// @router / [get]
func (c *ChequeEstadoChequeController) GetAll() {
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

	l, err := models.GetAllChequeEstadoCheque(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the ChequeEstadoCheque
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.ChequeEstadoCheque	true		"body for ChequeEstadoCheque content"
// @Success 200 {object} models.ChequeEstadoCheque
// @Failure 403 :id is not int
// @router /:id [put]
func (c *ChequeEstadoChequeController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.ChequeEstadoCheque{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateChequeEstadoChequeById(&v); err == nil {
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
// @Description delete the ChequeEstadoCheque
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *ChequeEstadoChequeController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteChequeEstadoCheque(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}


// Post ...
// @Title AddEstadoCheque
// @Description create estate for cheque
// @Param	body		body 	models.ChequeraEstadoChequera	true		"body for ChequeraEstadoChequera content"
// @Success 201 {int} models.ChequeraEstadoChequera
// @Failure 403 body is empty
// @router /AddEstadoCheque/ [post]
func (c *ChequeEstadoChequeController) AddEstadoCheque() {
	defer c.ServeJSON()
	var v map[string]interface{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if estadoCheque, err := models.AddNewEstadoCheque(v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = models.Alert{Type:"success",Code:"S_543",Body:estadoCheque}
		} else {
			var code string
			alertdb:=structs.Map(err)
			formatdata.FillStruct(alertdb["Code"],&code)
			c.Data["json"] = models.Alert{Type:"error",Code:"E_"+code,Body:err}
		}
	} else {
		c.Data["json"] = models.Alert{Type: "error", Code: "E_0458", Body: err}
	}

}
