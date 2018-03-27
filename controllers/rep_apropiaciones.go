package controllers

import (
	"github.com/astaxie/beego"
	"github.com/udistrital/api_financiera/models"
	"strconv"
)

// RepApropiacionesController operations for RepApropiaciones
type RepApropiacionesController struct {
	beego.Controller
}

// URLMapping ...
func (c *RepApropiacionesController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create RepApropiaciones
// @Param	body		body 	models.mongoDels.RepApropiaciones	true		"body for RepApropiaciones content"
// @Success 201 {object} models.mongoModels.RepApropiaciones
// @Failure 403 body is empty
// @router / [post]
func (c *RepApropiacionesController) Post() {

}

// GetOne ...
// @Title GetOne
// @Description get RepApropiaciones by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.RepApropiaciones
// @Failure 403 :anio is empty
// @router /:anio [get]
func (c *RepApropiacionesController) GetOne() {
	anioStr := c.Ctx.Input.Param(":anio")
	beego.Info("anio: ", anioStr)
	anio, _ := strconv.Atoi(anioStr)
	v, err := models.GetRepApropiacionesById(anio, 4, 1)
	if err != nil {
		beego.Info("error")
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title GetAll
// @Description get RepApropiaciones
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.mongoModels.RepApropiaciones
// @Failure 403
// @router / [get]
func (c *RepApropiacionesController) GetAll() {

}

// Put ...
// @Title Put
// @Description update the RepApropiaciones
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.mongoModels.RepApropiaciones	true		"body for RepApropiaciones content"
// @Success 200 {object} models.mongoModels.RepApropiaciones
// @Failure 403 :id is not int
// @router /:id [put]
func (c *RepApropiacionesController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the RepApropiaciones
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *RepApropiacionesController) Delete() {

}
