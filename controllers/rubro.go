package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/structs"
	"github.com/udistrital/api_financiera/models"
	"github.com/udistrital/api_financiera/utilidades"

	"github.com/astaxie/beego"
)

// RubroController operations for Rubro
type RubroController struct {
	beego.Controller
}

// URLMapping ...
func (c *RubroController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("RubroReporte", c.RubroReporte)
}

// Post ...
// @Title Post
// @Description create Rubro
// @Param	body		body 	models.Rubro	true		"body for Rubro content"
// @Success 201 {int} models.Rubro
// @Failure 403 body is empty
// @router / [post]
func (c *RubroController) Post() {
	var v models.Rubro
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if id, err := models.AddRubro(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = id
		} else {
			//fmt.Println("error: ", err )
			c.Data["json"] = err
		}
	} else {
		//fmt.Println("error: ", err )
		c.Data["json"] = err
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Rubro by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Rubro
// @Failure 403 :id is empty
// @router /:id [get]
func (c *RubroController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetRubroById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Rubro
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	group	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Rubro
// @Failure 403
// @router / [get]
func (c *RubroController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var group []string
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
	// related: value__related
	if v := c.GetString("group"); v != "" {
		grp := strings.Split(v, ",")
		for _, val := range grp {
			group = append(group, val)
		}
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

	l, err := models.GetAllRubro(query, group, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Rubro
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Rubro	true		"body for Rubro content"
// @Success 200 {object} models.Rubro
// @Failure 403 :id is not int
// @router /:id [put]
func (c *RubroController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.Rubro{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateRubroById(&v); err == nil {
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
// @Description delete the Rubro
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *RubroController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteRubro(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// RubroReporte ...
// @Title RubroReporte
// @Description Obtener reporte ingresos egresos de los rubros dada una vigencia
// @Param	finicio		path 	string	true		"fecha de inicio para el reporte"
// @Param	ffin		path 	string	true		"fecha final para el reporte"
// @Success 200 {object} interface{}
// @Failure 403 No se en contraron datos
// @router RubroReporte/ [post]

func (c *RubroController) RubroReporte() {
	var v interface{}
	var p interface{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &p); err == nil {
		m := p.(map[string]interface{})
		fmt.Println("inicio: ", m["inicio"])
		fmt.Println("inicio: ", m["fin"])
		var inicio time.Time
		err = utilidades.FillStruct(m["inicio"], &inicio)
		var fin time.Time
		err = utilidades.FillStruct(m["fin"], &fin)
		fmt.Println("format inicio: ", int(inicio.Year()))
		fmt.Println("fecha mod: ", inicio.AddDate(0, 1, 0))
		reporte := make(map[string]interface{})
		reporte["egresos"], err = models.RubroReporteEgresos(inicio, fin)
		reporte["ingresos"], err = models.RubroReporteIngresos(inicio, fin)
		//v, err = models.ListaApropiacionesHijo(2017)
		v = reporte
		if err != nil {
			alertdb := structs.Map(err)
			var code string
			utilidades.FillStruct(alertdb["Code"], &code)
			alert := models.Alert{Type: "error", Code: "E_" + code, Body: err}
			c.Data["json"] = alert
		} else {
			c.Data["json"] = v
		}
	} else {
		e := models.Alert{Type: "error", Code: "E_0458", Body: p}
		c.Data["json"] = e
	}

	c.ServeJSON()
}
