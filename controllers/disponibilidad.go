package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/udistrital/api_financiera/models"
)

// DisponibilidadController operations for Disponibilidad
type DisponibilidadController struct {
	beego.Controller
}

// URLMapping ...
func (c *DisponibilidadController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("SaldoCdp", c.SaldoCdp)
	c.Mapping("Anular", c.Anular)
}

// Post ...
// @Title Post
// @Description create Disponibilidad
// @Param	body		body 	models.Disponibilidad	true		"body for Disponibilidad content"
// @Success 201 {int} models.Disponibilidad
// @Failure 403 body is empty
// @router / [post]
func (c *DisponibilidadController) Post() {
	var v models.Disponibilidad
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddDisponibilidad(&v); err == nil {
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
// @Description get Disponibilidad by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Disponibilidad
// @Failure 403 :id is empty
// @router /:id [get]
func (c *DisponibilidadController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetDisponibilidadById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Disponibilidad
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Disponibilidad
// @Failure 403
// @router / [get]
func (c *DisponibilidadController) GetAll() {
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

	l, err := models.GetAllDisponibilidad(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Disponibilidad
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Disponibilidad	true		"body for Disponibilidad content"
// @Success 200 {object} models.Disponibilidad
// @Failure 403 :id is not int
// @router /:id [put]
func (c *DisponibilidadController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.Disponibilidad{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateDisponibilidadById(&v); err == nil {
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
// @Description delete the Disponibilidad
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *DisponibilidadController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteDisponibilidad(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

func (c *DisponibilidadController) Anular() {
	var v models.Info_disponibilidad_a_anular
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if v.Anulacion.TipoAnulacion == "T" {
			alertas, err := models.AnulacionTotal(&v)
			if err != nil {
				c.Data["json"] = err
			} else {
				c.Data["json"] = alertas
			}
		} else if v.Anulacion.TipoAnulacion == "P" {
			alertas, err := models.AnulacionParcial(&v)
			if err != nil {
				c.Data["json"] = err
			} else {
				c.Data["json"] = alertas
			}
		} else {
			c.Data["json"] = "No se pudo cargar el tipo de la anulacion"
		}

	} else {
		c.Data["json"] = err
	}
	c.ServeJSON()
}
func (c *DisponibilidadController) SaldoCdp() {
	var v models.DatosRubroRegistroPresupuestal
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		var ff int
		if v.FuenteFinanciacion != nil {
			ff = v.FuenteFinanciacion.Id
		} else {
			ff = 0
		}
		saldo, comprometido, anulado, err := models.SaldoCdp(v.Disponibilidad.Id, v.Apropiacion.Id, ff)
		if err != nil {
			c.Data["json"] = err
		} else {
			var m map[string]float64
			m = make(map[string]float64)
			m["saldo"] = saldo
			m["comprometido"] = comprometido
			m["anulado"] = anulado
			c.Data["json"] = m
		}
	} else {
		c.Data["json"] = err
		fmt.Println("error: ", err)
	}

	c.ServeJSON()
}

// AprobarAnulacionDisponibilidad ...
// @Title AprobarAnulacionDisponibilidad
// @Description aprueba la anulacion de un cdp ya sea total o parcial
// @Param	body		body 	models.AnulacionDisponibilidad	true		"body for AnulacionDisponibilidad content"
// @Success 201 {int} models.AnulacionDisponibilidad
// @Failure 403 body is empty
// @router /AprobarAnulacion [post]
func (c *DisponibilidadController) AprobarAnulacionDisponibilidad() {
	var v models.AnulacionDisponibilidad
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if alert, err := models.AprobacionAnulacion(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = alert
		} else {
			c.Data["json"] = alert
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

//-------------------
