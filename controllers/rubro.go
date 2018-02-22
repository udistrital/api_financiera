package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
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
	c.Mapping("ApropiacionReporte", c.ApropiacionReporte)
}
func genRubrosTreeFile(UnidadEjecutora int) {
	v, err := models.ArbolRubros(UnidadEjecutora, 0)
	rankingsJson, _ := json.Marshal(v)
	if err == nil {
		err = ioutil.WriteFile("RubroTreeUe"+strconv.Itoa(UnidadEjecutora)+".json", rankingsJson, 0644)

	}
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
		if _, err := models.AddRubro(&v); err == nil {
			alert := models.Alert{Type: "success", Code: "S_543", Body: v}
			go genRubrosTreeFile(int(v.UnidadEjecutora))
			c.Data["json"] = alert
		} else {
			alertdb := structs.Map(err)
			var code string
			utilidades.FillStruct(alertdb["Code"], &code)
			alert := models.Alert{Type: "error", Code: "E_" + code, Body: err}
			c.Data["json"] = alert
		}
	} else {
		alert := models.Alert{Type: "error", Code: "E_0458", Body: err}
		c.Data["json"] = alert
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

// ArbolRubros ...
// @Title ArbolRubros
// @Description genera arbol rubros
// @Param	idpadre		path 	string	true		"Id del padre para armar la rama default todos"
// @Param	UnidadEjecutora		path 	string	true		"Id del padre para armar la rama default todos"
// @Success 200 {object} models.Rubro
// @Failure 403 :id is empty
// @router ArbolRubros/ [get]
func (c *RubroController) ArbolRubros() {
	idpadre, _ := c.GetInt("idpadre")
	unidadEjecutora, err := c.GetInt("UnidadEjecutora")
	if err == nil {
		if _, err := os.Stat("RubroTreeUe" + strconv.Itoa(unidadEjecutora) + ".json"); os.IsNotExist(err) {
			// path/to/whatever does not exist
			v, err := models.ArbolRubros(unidadEjecutora, idpadre)
			if err != nil {
				c.Data["json"] = err.Error()
			} else {
				rankingsJson, _ := json.Marshal(v)
				err = ioutil.WriteFile("RubroTreeUe"+strconv.Itoa(unidadEjecutora)+".json", rankingsJson, 0644)
				fmt.Println("err ", err)
				c.Data["json"] = v
			}
		} else {
			data, _ := ioutil.ReadFile("RubroTreeUe" + strconv.Itoa(unidadEjecutora) + ".json")
			var v interface{}
			err = json.Unmarshal(data, &v)
			//fmt.Println("read from file")
			c.Data["json"] = v
		}

	} else {
		e := models.Alert{Type: "error", Code: "E_0458", Body: err.Error()}
		c.Data["json"] = e
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
	v, err1 := models.GetRubroById(id)
	if err := models.DeleteRubro(id); err == nil {
		alert := models.Alert{Type: "success", Code: "S_554", Body: nil}
		if err1 == nil {
			go genRubrosTreeFile(int(v.UnidadEjecutora))
		}
		c.Data["json"] = alert
	} else {
		alertdb := structs.Map(err)
		var code string
		utilidades.FillStruct(alertdb["Code"], &code)
		alert := models.Alert{Type: "error", Code: "E_" + code, Body: err}
		c.Data["json"] = alert
	}
	c.ServeJSON()
}

// ApropiacionReporte ...
// @Title ApropiacionReporte
// @Description Obtener reporte ingresos egresos de los rubros dada una vigencia
// @Param	finicio		path 	string	true		"fecha de inicio para el reporte"
// @Param	ffin		path 	string	true		"fecha final para el reporte"
// @Success 200 {object} interface{}
// @Failure 403 No se encontraron datos
// @router ApropiacionReporte/ [post]
func (c *RubroController) ApropiacionReporte() {
	var v interface{}
	var p interface{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &p); err == nil {
		m := p.(map[string]interface{})
		//fmt.Println("inicio: ", m["inicio"])
		//fmt.Println("inicio: ", m["fin"])
		var inicio time.Time
		err = utilidades.FillStruct(m["inicio"], &inicio)
		var fin time.Time
		err = utilidades.FillStruct(m["fin"], &fin)
		//fmt.Println("format inicio: ", int(inicio.Year()))
		//fmt.Println("fecha mod: ", inicio.AddDate(0, 1, 0))
		reporte := make(map[string]interface{})
		reporte["egresos"], err = models.ApropiacionReporteEgresos(inicio, fin)
		reporte["ingresos"], err = models.ApropiacionReporteIngresos(inicio, fin)
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

// GetRubroOrdenPago ...
// @Title Get Rubro Orden
// @Description get Apropiaciones Hijo
// @Param	apropiacion		path 	int64	true		"apropiacion a consultar"
// @Param	fuente		path 	int64	true		"fuente a consultar"
// @Success 200 {object} models.Rubro
// @Failure 403
// @router /GetRubroOrdenPago [get]
func (c *RubroController) GetRubroOrdenPago() {
	rubro, err := c.GetInt64("rubro")
	if err != nil {
		e := models.Alert{Type: "error", Code: "E_0458", Body: err.Error()}
		c.Data["json"] = e
		c.ServeJSON()
	}
	fuente, err := c.GetInt64("fuente")
	if err != nil {
		e := models.Alert{Type: "error", Code: "E_0458", Body: err.Error()}
		c.Data["json"] = e
		c.ServeJSON()
	}
	res, err := models.RubroOrdenPago(rubro, fuente)
	if err != nil {
		alertdb := structs.Map(err)
		var code string
		utilidades.FillStruct(alertdb["Code"], &code)
		alert := models.Alert{Type: "error", Code: "E_" + code, Body: err}
		c.Data["json"] = alert
		c.ServeJSON()
	}
	c.Data["json"] = res
	c.ServeJSON()
}

// GetRubroIngreso ...
// @Title Get Ingreso Rubro
// @Description get Apropiaciones Hijo
// @Param	apropiacion		path 	int64	true		"apropiacion a consultar"
// @Param	fuente		path 	int64	true		"fuente a consultar"
// @Param	finicio		path 	string	true		"fecha de inicio para el reporte"
// @Param	ffin		path 	string	true		"fecha final para el reporte"
// @Success 200 {object} models.Rubro
// @Failure 403
// @router /GetRubroIngreso [get]
func (c *RubroController) GetRubroIngreso() {
	rubro, err := c.GetInt64("rubro")
	if err != nil {
		e := models.Alert{Type: "error", Code: "E_0458", Body: err.Error()}
		c.Data["json"] = e
		c.ServeJSON()
	}
	fuente, err := c.GetInt64("fuente")
	if err != nil {
		e := models.Alert{Type: "error", Code: "E_0458", Body: err.Error()}
		c.Data["json"] = e
		c.ServeJSON()
	}
	finicioStr := c.GetString("finicio")

	ffinStr := c.GetString("ffin")

	finicio, err := time.ParseInLocation("2006-01-02", finicioStr, time.Local)
	if err != nil {
		e := models.Alert{Type: "error", Code: "E_0458", Body: err.Error()}
		c.Data["json"] = e
		c.ServeJSON()
	}
	ffin, err := time.ParseInLocation("2006-01-02", ffinStr, time.Local)
	if err != nil {
		e := models.Alert{Type: "error", Code: "E_0458", Body: err.Error()}
		c.Data["json"] = e
		c.ServeJSON()
	}
	var fuenteIf interface{}
	fuenteIf = fuente

	res, err := models.RubroIngreso(rubro, fuenteIf, finicio, ffin)
	c.Data["json"] = res
	c.ServeJSON()
}

// GetIngresoCierre ...
// @Title Get Ingreso Cierre
// @Description Obtiene informacion para cierre ingresos mes
// @Param	vigencia	path 	int64	true		"valor vigencia apropiacion"
// @Param	codigo		path 	string	true		"numero inicial de codigo rubro consultar"
// @Param	finicio		path 	string	true		"fecha de inicio para el reporte"
// @Param	ffin		path 	string	true		"fecha final para el reporte"
// @Success 200 {object} models.Rubro
// @Failure 403
// @router /GetIngresoCierre [get]
func (c *RubroController) GetIngresoCierre() {
	vigencia, err := c.GetInt64("vigencia")
	if err != nil {
		e := models.Alert{Type: "error", Code: "E_0458", Body: err.Error()}
		c.Data["json"] = e
		c.ServeJSON()
	}

	codigo := c.GetString("codigo")
	if err != nil {
		e := models.Alert{Type: "error", Code: "E_0458", Body: err.Error()}
		c.Data["json"] = e
		c.ServeJSON()
	}
	finicioStr := c.GetString("finicio")

	ffinStr := c.GetString("ffin")

	finicio, err := time.ParseInLocation("2006-01-02", finicioStr, time.Local)
	if err != nil {
		e := models.Alert{Type: "error", Code: "E_0458", Body: err.Error()}
		c.Data["json"] = e
		c.ServeJSON()
	}
	ffin, err := time.ParseInLocation("2006-01-02", ffinStr, time.Local)
	if err != nil {
		e := models.Alert{Type: "error", Code: "E_0458", Body: err.Error()}
		c.Data["json"] = e
		c.ServeJSON()
	}
	if v := strings.Compare(codigo, "2"); v == 0 {
		res, _ := models.RubroIngresoCierre(finicio, ffin, codigo+"%", vigencia)
		c.Data["json"] = res
		c.ServeJSON()
	} else {
		res, _ := models.RubroIngresoCierre(finicio, ffin, codigo+"%", vigencia)
		c.Data["json"] = res
		c.ServeJSON()
	}

}

// GetPacValue...
// @Title Get Pac value
// @Description Obtiene el valor de ejecucion para meses cesrrados
// @Param	vigencia	path 	int64	true		"valor vigencia apropiacion"
// @Param	mes			path 	int64	true		"valor mes a consultar cierre"
// @Param	rubro		path 	string	true		"id rubro a consultar"
// @Param	fuente		path 	string	true		"valor de la fuente a consultar"
// @Success 200 {object} models.Rubro
// @Failure 403
// @router /GetPacValue [get]
func (c *RubroController) GetPacValue() {
	vigencia, err := c.GetInt64("vigencia")
	if err != nil {
		e := models.Alert{Type: "error", Code: "E_0458", Body: err.Error()}
		c.Data["json"] = e
		c.ServeJSON()
	}

	mes, err := c.GetInt("mes")
	if err != nil {
		e := models.Alert{Type: "error", Code: "E_0458", Body: err.Error()}
		c.Data["json"] = e
		c.ServeJSON()
	}
	rubro := c.GetString("rubro")
	fuente := c.GetString("fuente")

	res, _ := models.ValEjecutadoPac(vigencia, mes, rubro, fuente)
	c.Data["json"] = res
	c.ServeJSON()
}

// GetSumbySource...
// @Title Get Pac value
// @Description Obtiene el valor de ejecucion para meses cesrrados
// @Param	vigencia	path 	int	true		"valor vigencia apropiacion"
// @Param	mes			path 	int	true		"valor mes a consultar cierre"
// @Param	fuente		path 	string	true		"valor de la fuente a consultar"
// @Param	tipo		path 	string	true		"ingresos :2, egresos:3"
// @Success 200 {object} models.Rubro
// @Failure 403
// @router /GetSumbySource [get]
func (c *RubroController) GetSumbySource() {
	vigencia, err := c.GetInt("vigencia")
	mes, err := c.GetInt("mes")
	if err != nil {
		e := models.Alert{Type: "error", Code: "E_0458", Body: err.Error()}
		c.Data["json"] = e
		c.ServeJSON()
	}
	fuente := c.GetString("fuente")
	tipo := c.GetString("tipo") + "%"
	if len(fuente) > 0 {
		res, _ := models.GetSumbySource(vigencia, mes, fuente, tipo)
		c.Data["json"] = res
	} else {
		beego.Error("entra a calcular por tipo")
		res, _ := models.GetSumbyTotal(vigencia, mes, tipo)
		c.Data["json"] = res
	}

	c.ServeJSON()
}

// GetRubroPac ...
// @Title Get Pac Rubro
// @Description get Apropiaciones Hijo
// @Param	vigencia		path 	int	true		"apropiacion a consultar"
// @Param	mes		path 	int	true		"fuente a consultar"
// @Param	idRubro		path 	string	true		"fecha de inicio para el reporte"
// @Param	idFuente		path 	string	true		"fecha final para el reporte"
// @Success 200 interface{}
// @Failure 403
// @router /GetRubroPac [get]
func (c *RubroController) GetRubroPac() {
	vigencia, err := c.GetInt("vigencia")
	if err != nil {
		e := models.Alert{Type: "error", Code: "E_0458", Body: err.Error()}
		c.Data["json"] = e
		c.ServeJSON()
	}
	mes, err := c.GetInt("mes")
	if err != nil {
		e := models.Alert{Type: "error", Code: "E_0458", Body: err.Error()}
		c.Data["json"] = e
		c.ServeJSON()
	}
	rubro := c.GetString("idRubro")
	fuente := c.GetString("idFuente")
	res, err := models.GetRubroPac(vigencia, mes, fuente, rubro)
	c.Data["json"] = res
	c.ServeJSON()
}
