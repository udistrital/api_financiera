package controllers

import (
	"github.com/astaxie/beego"
)

// Pasivos_fenecidosController operations for Pasivos_fenecidos
type Pasivos_fenecidosController struct {
	beego.Controller
}

// URLMapping ...
func (c *Pasivos_fenecidosController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Pasivos_fenecidos
// @Param	body		body 	models.Pasivos_fenecidos	true		"body for Pasivos_fenecidos content"
// @Success 201 {object} models.Pasivos_fenecidos
// @Failure 403 body is empty
// @router / [post]
func (c *Pasivos_fenecidosController) Post() {

}

// GetOne ...
// @Title GetOne
// @Description get Pasivos_fenecidos by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Pasivos_fenecidos
// @Failure 403 :id is empty
// @router /:id [get]
func (c *Pasivos_fenecidosController) GetOne() {

}

// GetAll ...
// @Title GetAll
// @Description get Pasivos_fenecidos
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Pasivos_fenecidos
// @Failure 403
// @router / [get]
func (c *Pasivos_fenecidosController) GetAll() {

}

// Put ...
// @Title Put
// @Description update the Pasivos_fenecidos
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Pasivos_fenecidos	true		"body for Pasivos_fenecidos content"
// @Success 200 {object} models.Pasivos_fenecidos
// @Failure 403 :id is not int
// @router /:id [put]
func (c *Pasivos_fenecidosController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Pasivos_fenecidos
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *Pasivos_fenecidosController) Delete() {

}
