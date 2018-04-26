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

// ProductoRubroController operations for ProductoRubro
type ProductoRubroController struct {
	beego.Controller
}

// URLMapping ...
func (c *ProductoRubroController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create ProductoRubro
// @Param	body		body 	models.ProductoRubro	true		"body for ProductoRubro content"
// @Success 201 {int} models.ProductoRubro
// @Failure 403 body is empty
// @router / [post]
func (c *ProductoRubroController) Post() {
	var v models.ProductoRubro
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddProductoRubro(&v); err == nil {
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
// @Description get ProductoRubro by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.ProductoRubro
// @Failure 403 :id is empty
// @router /:id [get]
func (c *ProductoRubroController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetProductoRubroById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get ProductoRubro
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.ProductoRubro
// @Failure 403
// @router / [get]
func (c *ProductoRubroController) GetAll() {
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

	l, err := models.GetAllProductoRubro(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the ProductoRubro
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.ProductoRubro	true		"body for ProductoRubro content"
// @Success 200 {object} models.ProductoRubro
// @Failure 403 :id is not int
// @router /:id [put]
func (c *ProductoRubroController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.ProductoRubro{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateProductoRubroById(&v); err == nil {
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
// @Description delete the ProductoRubro
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *ProductoRubroController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteProductoRubro(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// SetVariacionProducto ...
// @Title SetVariacionProducto
// @Description set new product variation
// @Param	body		body 	models.ProductoRubro	true		"body for ProductoRubro content"
// @Success 201 {int} models.Alert
// @Failure 403 body is empty
// @router /SetVariacionProducto/ [post]
func (c *ProductoRubroController) SetVariacionProducto() {
	var v models.ProductoRubro
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if total, err := models.SetVariacionProducto(&v); err == nil && total <= 1 {
			c.Ctx.Output.SetStatus(201)
			alert := models.Alert{Type: "success", Code: "S_543", Body: v}
			c.Data["json"] = alert
		} else {
			if err != nil {
				alertdb := structs.Map(err)
				var code string
				formatdata.FillStruct(alertdb["Code"], &code)
				alert := models.Alert{Type: "error", Code: "E_" + code, Body: err}
				c.Data["json"] = alert
			} else if total > 1 {
				alert := models.Alert{Type: "error", Code: "E_RB006", Body: v}
				c.Data["json"] = alert
			} else {
				alert := models.Alert{Type: "error", Code: "E_0458", Body: err}
				c.Data["json"] = alert
			}

		}
	} else {
		alert := models.Alert{Type: "error", Code: "E_0458", Body: err}
		c.Data["json"] = alert
	}
	c.ServeJSON()
}

// AddProductoRubrotr ...
// @Title AddProductoRubrotr
// @Description set new product for rubro producto relation
// @Param	body		body 	models.ProductoRubro	true		"body for ProductoRubro content"
// @Success 201 {int} models.Alert
// @Failure 403 body is empty
// @router /AddProductoRubrotr/ [post]
func (c *ProductoRubroController) AddProductoRubrotr() {
	var v models.ProductoRubro
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if total, err := models.AddProductoRubrotr(&v); err == nil && total <= 1 {
			c.Ctx.Output.SetStatus(201)
			alert := models.Alert{Type: "success", Code: "S_543", Body: v}
			c.Data["json"] = alert
		} else {
			if err != nil {
				alertdb := structs.Map(err)
				var code string
				formatdata.FillStruct(alertdb["Code"], &code)
				alert := models.Alert{Type: "error", Code: "E_" + code, Body: err}
				c.Data["json"] = alert
			} else if total > 1 {
				alert := models.Alert{Type: "error", Code: "E_RB006", Body: v}
				c.Data["json"] = alert
			} else {
				alert := models.Alert{Type: "error", Code: "E_0458", Body: err}
				c.Data["json"] = alert
			}

		}
	} else {
		alert := models.Alert{Type: "error", Code: "E_0458", Body: err}
		c.Data["json"] = alert
	}
	c.ServeJSON()
}