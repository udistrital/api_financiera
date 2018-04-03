package controllers

import (
	"github.com/astaxie/beego"
	"github.com/manucorporat/try"
	"github.com/udistrital/api_financiera/models"
	"github.com/udistrital/utils_oas/optimize"
	"strconv"
	"time"
)

// Pasivos_fenecidosController operations for Pasivos_fenecidos
type Pasivos_fenecidosController struct {
	beego.Controller
}

// URLMapping ...
func (c *Pasivos_fenecidosController) URLMapping() {
	c.Mapping("Post", c.Post)
}

func PFenecidosInit() {
	optimize.StartDispatcher(1, 200)
	//beego.InsertFilter("/v1/ingreso/AprobacionPresupuestalIngreso", beego.AfterExec, FunctionAfterExecIngresoPac, false)
}

// Post ...
// @Title Create
// @Description create Pasivos_fenecidos
// @Param	body		body 	models.Pasivos_fenecidos	true		"body for Pasivos_fenecidos content"
// @Success 201 {object} models.Pasivos_fenecidos
// @Failure 403 body is empty
// @router / [post]
func (c *Pasivos_fenecidosController) Post() {
	/*dataRp, err := models.GetAllRegistroPresupuestal(query, fields, sortby, order, offset, limit)
	if err == nil {

	}*/
	//dataDisp, err := models.GetAllDisponibilidad(query, fields, sortby, order, offset, limit)
	try.This(func() {
		work := optimize.WorkRequest{JobParameter: nil, Job: (pasivosFenecidosProcess)}
		// Push the work onto the queue.

		optimize.WorkQueue <- work
		c.Data["json"] = models.Alert{Code: "S_F001", Body: nil, Type: "success"}
	}).Catch(func(e try.E) {
		// Print crash
		//beego.Info("expc ", e)
		c.Data["json"] = models.Alert{Code: "E_0458", Body: e, Type: "error"}
	})

	c.ServeJSON()
}

func pasivosFenecidosProcess(parameter ...interface{}) (err interface{}) {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = -1
	var offset int64
	query["Estado.Id__in"] = "1|2"
	query["Vigencia"] = strconv.Itoa(time.Now().Local().Year() - 2)
	try.This(func() {
		dataRp, _ := models.GetAllRegistroPresupuestal(query, fields, sortby, order, offset, limit)
		beego.Info("Feneciendo CRP...")
		optimize.ProccDigest(dataRp, fenecerCrp, nil,4)
		dataCdp, _ := models.GetAllDisponibilidad(query, fields, sortby, order, offset, limit)
		beego.Info("Feneciendo CDP...")
		optimize.ProccDigest(dataCdp, fenecerCdp, nil,4)
		beego.Info("Proceso Finalizado... ")
	}).Catch(func(e try.E) {
		// Print crash
		beego.Info("expc ", e)
	})
	return
}

func fenecerCrp(rpintfc interface{}, params ...interface{}) (res interface{}) {
	try.This(func() {
		rp := rpintfc.(models.RegistroPresupuestal)
		//crear la estructura de la info de la anulación
		anulacion := models.AnulacionRegistroPresupuestal{} //Info_rp_a_anular{}
		anulacion.Motivo = "Se Fenece por No utilización de Recursos."
		anulacion.EstadoAnulacion = &models.EstadoAnulacion{Id: 3}
		anulacion.TipoAnulacion = &models.TipoAnulacionPresupuestal{Id: 3}
		anulacion.Expidio = 0
		anulacion.Solicitante = 0
		anulacion.Responsable = 0
		datosAnulacion := models.Info_rp_a_anular{}
		datosAnulacion.Anulacion = anulacion
		datosAnulacion.Rp_apropiacion = rp.RegistroPresupuestalDisponibilidadApropiacion
		datosAnulacion.Valor = 0
		_, err := models.AnulacionTotalRp(&datosAnulacion)
		if err != nil {
			panic(err)
		}
		//beego.Info("Data anulacion: ", datosAnulacion)
	}).Catch(func(e try.E) {
		// Print crash
		beego.Info("Excp ", e)
		return
	})
	return
}

func fenecerCdp(cdpintfc interface{}, params ...interface{}) (res interface{}) {
	try.This(func() {
		cdp := cdpintfc.(models.Disponibilidad)
		//crear la estructura de la info de la anulación
		anulacion := models.AnulacionDisponibilidad{}
		anulacion.Motivo = "Se Fenece por No utilización de Recursos."
		anulacion.EstadoAnulacion = &models.EstadoAnulacion{Id: 3}
		anulacion.TipoAnulacion = &models.TipoAnulacionPresupuestal{Id: 3}
		anulacion.Expidio = 0
		anulacion.Solicitante = 0
		anulacion.Responsable = 0
		datosAnulacion := models.Info_disponibilidad_a_anular{}
		datosAnulacion.Anulacion = anulacion
		datosAnulacion.Disponibilidad_apropiacion = cdp.DisponibilidadApropiacion
		datosAnulacion.Valor = 0
		_, err := models.AnulacionTotal(&datosAnulacion)
		if err != nil {
			beego.Info("err ",err)
		}
		//beego.Info("Data anulacion: ", datosAnulacion)
	}).Catch(func(e try.E) {
		// Print crash
		beego.Info("Excp ", e)
		return
	})
	return
}
