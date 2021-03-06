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

// GiroController operations for Giro
type GiroController struct {
	beego.Controller
}

// URLMapping ...
func (c *GiroController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Giro
// @Param	body		body 	models.Giro	true		"body for Giro content"
// @Success 201 {int} models.Giro
// @Failure 400 the request contains incorrect syntax
// @router / [post]
func (c *GiroController) Post() {
	var v models.Giro
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddGiro(&v); err == nil {
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
// @Description get Giro by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Giro
// @Failure 404 not found resource
// @router /:id [get]
func (c *GiroController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetGiroById(id)
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
// @Description get Giro
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Giro
// @Failure 404 not found resource
// @router / [get]
func (c *GiroController) GetAll() {
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

	l, err := models.GetAllGiro(query, fields, sortby, order, offset, limit)
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
// @Description update the Giro
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Giro	true		"body for Giro content"
// @Success 200 {object} models.Giro
// @Failure 400 the request contains incorrect syntax
// @router /:id [put]
func (c *GiroController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.Giro{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateGiroById(&v); err == nil {
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
// @Description delete the Giro
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 404 not found resource
// @router /:id [delete]
func (c *GiroController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteGiro(id); err == nil {
		c.Data["json"] = map[string]interface{}{"Id": id}
	} else {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("404")
	}
	c.ServeJSON()
}

// RegistrarGiro ...
// @Title RegistrarGiro
// @Description Registrar Giro orden_pago de proveedor, concepto_ordenpago, mivimientos contables
// @Param	body		body 	models.giro	true		"body for giro content"
// @Success 201 {int} models.Giro
// @Failure 403 body is empty
// @router RegistrarGiro [post]
func (c *GiroController) RegistrarGiro() {
	var v interface{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		m := v.(map[string]interface{})
		mensaje := models.RegistrarGiro(m)
		if mensaje.Type != "success" {
			c.Data["json"] = mensaje
		} else {
			c.Ctx.Output.SetStatus(201)
			//alert := models.Alert{Type: mensaje.Type, Code: mensaje.Code, Body: consecutivoOp}
			c.Data["json"] = mensaje
		}
	} else {
		c.Data["json"] = err
	}
	c.ServeJSON()
}

// GetCuentasEspeciales ...
// @Title GetCuentasEspeciales
// @Description obtiene las Cuentas Especiales de cada Orden de Pago
// @Param	query	idordenpago	 int64	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Success 200 {object} models.Giro
// @Failure 403
// @router /GetCuentasEspeciales [get]
func (c *GiroController) GetCuentasEspeciales() {
	var idOrdenPago int64
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("idordenpago"); err == nil {
		idOrdenPago = v
	}
	l, mensaje := models.GetCuentasEspeciales(idOrdenPago)
	if mensaje.Body != nil {
		c.Data["json"] = mensaje.Body
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()

}

// GetSumGiro ...
// @Title GetSumGiro
// @Description obtiene la Suma del Giro
// @Param	query	idordenpago	 int64	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Success 200 {object} models.Giro
// @Failure 403
// @router /GetSumGiro/ [get]
func (c *GiroController) GetSumGiro() {
	var idGiro int64
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("IdGiro"); err == nil {
		idGiro = v
	}
	l, mensaje := models.GetSumGiro(idGiro)
	if mensaje.Body != nil {
		c.Data["json"] = mensaje.Body
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()

}

// GetValueEndoso ...
// @Title GetValueEndoso
// @Description obtiene el valor del Endoso
// @Param	query	IdGiro	 int64	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	query	idordenpago	 int64	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Success 200 {object} models.Giro
// @Failure 403
// @router /GetValueEndoso/ [get]
func (c *GiroController) GetValueEndoso() {
	var idOrdenPago int64
	var idCuentaEspecial int64
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("IdOrdenPago"); err == nil {
		idOrdenPago = v
	}
	if v, err := c.GetInt64("IdCuentaEspecial"); err == nil {
		idCuentaEspecial = v
	}
	l, mensaje := models.GetValueEndoso(idOrdenPago, idCuentaEspecial)
	if mensaje.Body != nil {
		c.Data["json"] = mensaje.Body
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()

}

// RegistrarGiroDescuentos ...
// @Title RegistrarGiroDescuentos
// @Description RegistrarGiroDescuentos orden_pago de proveedor, concepto_ordenpago, movimientos contables
// @Param	body			body 	models.Goru true "body for giro descuentos content"
// @Param	idcuenta		query	string	false	"Limit the size of result set. Must be an integer"
// @Param	idordenpago		query	string	false	"Limit the size of result set. Must be an integer"
// @Param	idgiro			query	string	false	"Limit the size of result set. Must be an integer"
// @Success 201 {int} models.Giro
// @Failure 403 body is empty
// @router /RegistrarGiroDescuentos [post]
func (c *GiroController) RegistrarGiroDescuentos() {
	var v interface{}
	var idCuenta int64
	var idOrdenPago int64
	var idGiro int64
	var resProveedor []interface{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		m := v.(map[string]interface{})
		idCuenta, _ = strconv.ParseInt(m["idCuenta"].(string), 0, 64)
		idOrdenPago, _ = strconv.ParseInt(m["idOrdenPago"].(string), 0, 64)
		idGiro, _ = strconv.ParseInt(m["idGiro"].(string), 0, 64)
		resProveedor = m["resProveedor"].([]interface{})
		mensaje := models.RegistrarGiroDescuentos(resProveedor, idGiro, idCuenta, idOrdenPago)
		if mensaje.Type != "success" {
			c.Data["json"] = mensaje
		} else {
			c.Ctx.Output.SetStatus(201)
			//alert := models.Alert{Type: mensaje.Type, Code: mensaje.Code, Body: consecutivoOp}
			c.Data["json"] = mensaje
		}
	} else {
		c.Data["json"] = err
	}
	c.ServeJSON()
}
