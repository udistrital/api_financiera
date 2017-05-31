package controllers

import (
	"github.com/udistrital/api_financiera/models"
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"fmt"

	"github.com/astaxie/beego"
)

// OrdenPagoController operations for OrdenPago
type OrdenPagoController struct {
	beego.Controller
}

// URLMapping ...
func (c *OrdenPagoController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create OrdenPago
// @Param	body		body 	models.OrdenPago	true		"body for OrdenPago content"
// @Success 201 {int} models.OrdenPago
// @Failure 403 body is empty
// @router / [post]
func (c *OrdenPagoController) Post() {
	var v models.OrdenPago
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddOrdenPago(&v); err == nil {
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
// @Description get OrdenPago by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.OrdenPago
// @Failure 403 :id is empty
// @router /:id [get]
func (c *OrdenPagoController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetOrdenPagoById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get OrdenPago
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.OrdenPago
// @Failure 403
// @router / [get]
func (c *OrdenPagoController) GetAll() {
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

	l, err := models.GetAllOrdenPago(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the OrdenPago
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.OrdenPago	true		"body for OrdenPago content"
// @Success 200 {object} models.OrdenPago
// @Failure 403 :id is not int
// @router /:id [put]
func (c *OrdenPagoController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.OrdenPago{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateOrdenPagoById(&v); err == nil {
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
// @Description delete the OrdenPago
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *OrdenPagoController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteOrdenPago(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// personalizado Registrar orden_pago, concepto_ordenpago y transacciones
func (c *OrdenPagoController) RegistrarOpProveedor() {
    var v models.Data_OrdenPago_Concepto
    if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
        mensaje, err, id_orden := models.RegistrarOpProveedor(&v)
        if err != nil {
          fmt.Println(mensaje)
          c.Data["json"] = err
        } else {
          c.Data["json"] = id_orden
        }
    } else {
        c.Data["json"] = err
    }
    c.ServeJSON()
}

// personalizado Actualizar orden_pago, concepto_ordenpago y Movimientos Contables
func (c *OrdenPagoController) ActualizarOpProveedor() {
    var v models.Data_OrdenPago_Concepto
    if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
        mensaje, err, id_orden := models.ActualizarOpProveedor(&v)
        if err != nil {
          fmt.Println(mensaje)
          c.Data["json"] = err
        } else {
          c.Data["json"] = id_orden
        }
    } else {
      c.Data["json"] = err
    }
    c.ServeJSON()
}

// personalizado Registrar orden_pago nomina planta, homologa conceptos titan-kronos, concepto_ordenpago y transacciones
func (c *OrdenPagoController) RegistrarOpPlanta() {
		fmt.Println("1-------")
    var v models.OrdenPago
    if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
        mensaje, err, id_orden := models.RegistrarOpPlanta(&v)
        if err != nil {
          fmt.Println("2-------")
          c.Data["json"] = err
        } else {
					fmt.Println("3-------")
          c.Data["json"] = id_orden
					fmt.Println(mensaje)
        }
    } else {
				fmt.Println("4-------")
        c.Data["json"] = err
				fmt.Println(err)
    }
    c.ServeJSON()
}


// personalizado Retrona la fecha actual del servidor
func (c *OrdenPagoController) FechaActual() {
		formatoInput := c.Ctx.Input.Param(":formato")
		fechaActual, err := models.FechaActual(formatoInput)
		if err == nil{
			c.Data["json"] = fechaActual
		} else {
			c.Data["json"] = err.Error()
		}
		c.ServeJSON()
}
