package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/fatih/structs"
	"github.com/manucorporat/try"
	"github.com/udistrital/api_financiera/models"
	"github.com/udistrital/utils_oas/formatdata"
	"github.com/udistrital/utils_oas/optimize"
)

// ApropiacionController operations for Apropiacion
type ApropiacionController struct {
	beego.Controller
}

// URLMapping ...
func (c *ApropiacionController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("SaldoApropiacion", c.SaldoApropiacion)
}

// Post ...
// @Title Post
// @Description create Apropiacion
// @Param	body		body 	models.Apropiacion	true		"body for Apropiacion content"
// @Success 201 {int} models.Apropiacion
// @Failure 400 the request contains incorrect syntax
// @router / [post]
func (c *ApropiacionController) Post() {
	var v models.Apropiacion
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddApropiacion(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = models.Alert{Type: "success", Code: "S_543", Body: v}
		} else {
			alertdb := structs.Map(err)
			var code string
			formatdata.FillStruct(alertdb["Code"], &code)
			alert := models.Alert{Type: "error", Code: "E_" + code, Body: err}
			c.Data["json"] = alert
		}
	} else {
		c.Data["json"] = models.Alert{Code: "E_0458", Body: err, Type: "error"}
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Apropiacion by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Apropiacion
// @Failure 404 not found resource
// @router /:id [get]
func (c *ApropiacionController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetApropiacionById(id)
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
// @Description get Apropiacion
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Apropiacion
// @Failure 404 not found resource
// @router / [get]
func (c *ApropiacionController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var exclude = make(map[string]string)
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

	// exclude: k:v,k:v
	if v := c.GetString("exclude"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid exclude key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			exclude[k] = v
		}
	}

	l, err := models.GetAllApropiacion(query, exclude, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"Data": err.Error()}
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Apropiacion
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Apropiacion	true		"body for Apropiacion content"
// @Success 200 {object} models.Apropiacion
// @Failure 400 the request contains incorrect syntax
// @router /:id [put]
func (c *ApropiacionController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.Apropiacion{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateApropiacionById(&v); err == nil {
			c.Data["json"] = models.Alert{Type: "success", Code: "S_542", Body: v}
		} else {
			alertdb := structs.Map(err)
			var code string
			formatdata.FillStruct(alertdb["Code"], &code)
			alert := models.Alert{Type: "error", Code: "E_" + code, Body: err}
			c.Data["json"] = alert
		}
	} else {
		c.Data["json"] = models.Alert{Code: "E_0458", Body: err, Type: "error"}
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Apropiacion
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 404 not found resource
// @router /:id [delete]
func (c *ApropiacionController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteApropiacion(id); err == nil {
		c.Data["json"] = map[string]interface{}{"Id": id}
	} else {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("404")
	}
	c.ServeJSON()
}

// SaldoApropiacion ...
// @Title Get Saldo Apropiacion By Id
// @Description Get Saldo Apropiacion By Id
// @Param	Id	path 	string	true		"Id de la apropiacion"
// @Success 200 {object} float64
// @Failure 403
// @router /SaldoApropiacion/:id [get]
func (c *ApropiacionController) SaldoApropiacion() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)

	valor, err := models.SaldoApropiacion(id)
	//valor, err := models.SaldoRubroPadre(id, 1, 2017)
	if err != nil {
		alertdb := structs.Map(err)
		var code string
		formatdata.FillStruct(alertdb["Code"], &code)
		alert := models.Alert{Type: "error", Code: "E_" + code, Body: err}
		c.Data["json"] = alert
	} else {
		c.Data["json"] = valor
	}

	c.ServeJSON()
}

// SaldoApropiacionPadre ...
// @Title Get Saldo Apropiacion By Id
// @Description Get Saldo Apropiacion By Id
// @Param	Id	path 	string	true		"Id del rubro padre"
// @Param	UnidadEjecutora	query	string	false	"Unidad Ejecutora de los rubros hijo"
// @Param	Vigencia	query	string	false	"Vigencia de las apropiaciones a consultar"
// @Success 200 {object} float64
// @Failure 403
// @router /SaldoApropiacionPadre/:id [get]
func (c *ApropiacionController) SaldoApropiacionPadre() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	ue, err := c.GetInt("UnidadEjecutora")
	vigencia, err1 := c.GetInt("Vigencia")
	if err == nil && err1 == nil {
		valor, err := models.SaldoRubroPadre(id, ue, vigencia)
		if err != nil {
			alertdb := structs.Map(err)
			var code string
			formatdata.FillStruct(alertdb["Code"], &code)
			alert := models.Alert{Type: "error", Code: "E_" + code, Body: err}
			c.Data["json"] = alert
		} else {
			c.Data["json"] = valor
		}
	} else {
		c.Data["json"] = models.Alert{Code: "E_0458", Body: nil, Type: "error"}
	}

	c.ServeJSON()
}

//-----------------------------------------------

// GetApropiacionesHijo ...
// @Title Get Apropiaciones Hijo
// @Description get Apropiaciones Hijo
// @Param	vigencia	path 	string	true		"vigencia filtro de las apropiaciones hijo"
// @Param	tipo	path 	string	true		"tipo del rubro"
// @Success 200 {object} models.Apropiacion
// @Failure 403
// @router /GetApropiacionesHijo/:vigencia [get]
func (c *ApropiacionController) GetApropiacionesHijo() {
	vigStr := c.Ctx.Input.Param(":vigencia")
	tipo := c.GetString("tipo")
	vigencia, err := strconv.Atoi(vigStr)
	if err != nil {
		c.Data["json"] = models.Alert{Code: "E_XXX", Body: err.Error(), Type: "error"}
	} else {
		m, err := models.ListaApropiacionesHijo(vigencia, tipo+"%")
		if err != nil {
			alertdb := structs.Map(err)
			var code string
			formatdata.FillStruct(alertdb["Code"], &code)
			alert := models.Alert{Type: "error", Code: "E_" + code, Body: err}
			c.Data["json"] = alert
		} else {
			c.Data["json"] = m
		}

	}
	c.ServeJSON()
}

// ArbolApropiaciones ...
// @Title ArbolApropiaciones
// @Description genera arbol apropiaciones
// @Success 200 {object} models.Rubro
// @Failure 403 :vigencia is empty
// @router /ArbolApropiaciones/:vigencia [get]
func (c *ApropiacionController) ArbolApropiaciones() {
	vigStr := c.Ctx.Input.Param(":vigencia")
	vig, err := strconv.Atoi(vigStr)
	var parameters []interface{}
	unidadEjecutora := 1
	if err == nil {
		if _, err := os.Stat("apropaciones_" + strconv.Itoa(unidadEjecutora) + "_" + strconv.Itoa(vig) + ".json"); os.IsNotExist(err) {
			parameters = append(parameters, 1)
			parameters = append(parameters, vig)
			work := optimize.WorkRequest{JobParameter: parameters, Job: (models.EncapsuArbolApropiaciones)}
			// Push the work onto the queue.
			select {
			case worker := <-optimize.WorkerQueue:
				optimize.WorkerQueue <- worker
				optimize.WorkQueue <- work
				fmt.Println("Envie el trabajo...")
			default:
				fmt.Println("No mande el trabajo :'(")
			}

			// v, err := models.EncapsuArbolApropiaciones(unidadEjecutora, idpadre)
			fmt.Println("Construyendo archivo (controler)...")
			c.Data["json"] = models.Alert{Type: "success", Code: "procesando archivo....", Body: nil}
		} else {
			// aca devuelve el archivo
			data, _ := ioutil.ReadFile("apropaciones_" + strconv.Itoa(unidadEjecutora) + "_" + strconv.Itoa(vig) + ".json")
			var v interface{}
			err = json.Unmarshal(data, &v)
			fmt.Println("Devolviendo archivo leido....")
			c.Data["json"] = v
		}
	}

	c.ServeJSON()
}

// AprobarPresupuesto ...
// @Title AprobarPresupuesto
// @Description aprueba la asignacion inicial de presupuesto
// @Param	Vigencia		query 	string	true		"vigencia a comprobar"
// @Param	UnidadEjecutora		query 	string	true		"unidad ejecutora de los rubros a comprobar"
// @Success 200 {string} resultado
// @Failure 403
// @router /AprobacionAsignacionInicial/ [get]
func (c *ApropiacionController) AprobarPresupuesto() {
	vigencia, err := c.GetInt("Vigencia")
	if err == nil {
		unidadejecutora, err := c.GetInt("UnidadEjecutora")
		if err == nil {
			err = models.AprobarPresupuesto(unidadejecutora, vigencia)
			if err == nil {
				c.Data["json"] = models.Alert{Code: "S_AP001", Body: nil, Type: "success"}
			} else {
				alertdb := structs.Map(err)
				var code string
				formatdata.FillStruct(alertdb["Code"], &code)
				alert := models.Alert{Type: "error", Code: "E_" + code, Body: err}
				c.Data["json"] = alert
			}
		} else {
			c.Data["json"] = models.Alert{Code: "E_0458", Body: err.Error(), Type: "error"}
		}
	} else {
		c.Data["json"] = models.Alert{Code: "E_0458", Body: err.Error(), Type: "error"}
	}

	c.ServeJSON()
}

// VigenciaApropiaciones ...
// @Title VigenciaApropiaciones
// @Description Obtiene todas las vigencias no repetidas de las apropiaciones
// @Success 200 {string} resultado
// @Failure 403
// @router /VigenciaApropiaciones [get]
func (c *ApropiacionController) VigenciaApropiaciones() {
	m, err := models.VigenciaApropiacion()
	if err != nil {
		c.Data["json"] = models.Alert{Code: "E_458", Body: err.Error(), Type: "error"}
	} else {
		c.Data["json"] = m
	}
	c.ServeJSON()
}

// UpdateApropiacionValue ...
// @Title UpdateApropiacionValue
// @Description Obtiene todas las vigencias no repetidas de las apropiaciones
// @Success 200 {string} resultado
// @Failure 403
// @router /UpdateApropiacionValue/:id/:valor [put]
func (c *ApropiacionController) UpdateApropiacionValue() {
	idStr := c.Ctx.Input.Param(":id")
	valStr := c.Ctx.Input.Param(":valor")
	try.This(func() {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			panic(err.Error())
		}
		val, err := strconv.ParseFloat(valStr, 64)
		if err != nil {
			panic(err.Error())
		}
		models.UpdateApropiacionValue(id, val)
		c.Data["json"] = map[string]interface{}{"Code": "A_S002", "Body": nil, "Type": "success"}

	}).Catch(func(e try.E) {
		fmt.Println("expc ", e)
		c.Data["json"] = map[string]interface{}{"Code": "E_0458", "Body": e, "Type": "error"}
	})
	c.ServeJSON()

}
