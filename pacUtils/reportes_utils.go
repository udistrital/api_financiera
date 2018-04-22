package pacUtils

import (
	"strconv"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/manucorporat/try"
	"github.com/udistrital/api_financiera/models"
	"github.com/udistrital/utils_oas/optimize"
	"github.com/udistrital/utils_oas/request"
)

func AddMovimientoApropiacion(parameter ...interface{}) (err interface{}) {
	try.This(func() {
		var (
			apropiacionMongo []models.MongoApropiacion
			movimientoMongo models.MongoMovimiento
			respuesta  interface{}
		)
		mongoApiURL := beego.AppConfig.String("MongoApi")
		Movimiento := parameter[0].(*models.MovimientoApropiacion)


		request.GetJson("http://"+mongoApiURL+"apropiacion?query=rubro.codigo:"+Movimiento.MovimientoApropiacionDisponibilidadApropiacion[0].CuentaCredito.Rubro.Codigo+",vigencia:"+strconv.Itoa(Movimiento.Vigencia), &apropiacionMongo)
		beego.Info("http://"+mongoApiURL+"apropiacion?query=rubro.codigo:"+Movimiento.MovimientoApropiacionDisponibilidadApropiacion[0].CuentaCredito.Rubro.Codigo+",vigencia:"+strconv.Itoa(Movimiento.Vigencia))
		beego.Info(len(apropiacionMongo))
		beego.Info(apropiacionMongo[0])

		movimientoMongo.Numero = strconv.Itoa(Movimiento.NumeroMovimiento)
		movimientoMongo.Estado_movimiento = "Aprobado"
		movimientoMongo.Fecha_movimiento = Movimiento.FechaMovimiento.Format("2006-01-02")
		movimientoMongo.Numero_oficio = strconv.Itoa(Movimiento.Noficio)
		movimientoMongo.Fecha_oficio = Movimiento.Foficio.Format("2006-01-02")
		movimientoMongo.Descripcion = Movimiento.Descripcion
		movimientoMongo.RubroDestino = Movimiento.MovimientoApropiacionDisponibilidadApropiacion[0].CuentaCredito.Rubro.Codigo
		movimientoMongo.RubroOrigen = ""
		if Movimiento.MovimientoApropiacionDisponibilidadApropiacion[0].CuentaContraCredito != nil {
			movimientoMongo.RubroOrigen = Movimiento.MovimientoApropiacionDisponibilidadApropiacion[0].CuentaContraCredito.Rubro.Codigo
		}
		movimientoMongo.Valor = int64(Movimiento.MovimientoApropiacionDisponibilidadApropiacion[0].Valor)
		movimientoMongo.Tipo_movimiento = Movimiento.MovimientoApropiacionDisponibilidadApropiacion[0].TipoMovimientoApropiacion.Nombre
		movimientoMongo.Apropiacion = apropiacionMongo[0]
		request.SendJson("http://"+mongoApiURL+"movimiento", "POST", &respuesta, movimientoMongo)
		if respuesta.(string) != "" {
			beego.Info("registrado movimiento en mongo")
		}
	}).Catch(func(e try.E) {
		beego.Info(e)
	})
	return nil
}

func AddRegistroPresupuestal(parameter ...interface{}) (err interface{}) {
	return nil
}

func AddOrdenDePago(parameter ...interface{}) (err interface{}) {
	return nil
}

func ReportesInit() {
	optimize.StartDispatcher(1, 200)
	beego.InsertFilter("/v1/movimiento_apropiacion/AprobarMovimietnoApropiacion", beego.AfterExec, saveMovimiento, false)
	beego.InsertFilter("v1/registro_presupuestal", beego.AfterExec, saveRegistroPresupuestal, false)
	beego.InsertFilter("v1/orden_pago_estado_orden_pago/WorkFlowOrdenPago", beego.AfterExec, saveOrdenDePago, false)
	// beego.InsertFilter("/v1/movimiento_apropiacion/AprobarMovimietnoApropiacion", beego.AfterExec, FunctionAfterExecEstadoOrdenP, false)
	// /v1/movimiento_apropiacion/AprobarMovimietnoApropiacion
}

func saveMovimiento(ctx *context.Context) {
	var parameters []interface{}
	try.This(func() {
		if response := ctx.Input.Data()["json"].([]models.Alert)[0].Type; response == "success" {
			parameters = append(parameters, ctx.Input.Data()["json"].([]models.Alert)[0].Body.(map[string]interface{})["Movimiento"])

			work := optimize.WorkRequest{JobParameter: parameters, Job: AddMovimientoApropiacion}
			optimize.WorkQueue <- work
		}

	}).Catch(func(e try.E) { // Aquí se resuelven los errores
		beego.Info(e)
	})
}

func saveRegistroPresupuestal(ctx *context.Context) {
	var parameters []interface{}
	try.This(func() {
		if response := ctx.Input.Data()["json"].([]models.Alert)[0].Type; response == "success" {
			// otros parámetros
			parameters = append(parameters, ctx.Input.Data()["json"].([]models.Alert)[0].Body.(map[string]interface{})["Movimiento"])

			work := optimize.WorkRequest{JobParameter: parameters, Job: AddRegistroPresupuestal}
			optimize.WorkQueue <- work
		}

	}).Catch(func(e try.E) { // Aquí se resuelven los errores
		beego.Info(e)
	})
}

func saveOrdenDePago(ctx *context.Context) {
	var parameters []interface{}
	try.This(func() {
		if response := ctx.Input.Data()["json"].([]models.Alert)[0].Type; response == "success" {
			// otros parámetros
			parameters = append(parameters, ctx.Input.Data()["json"].([]models.Alert)[0].Body.(map[string]interface{})["Movimiento"])

			work := optimize.WorkRequest{JobParameter: parameters, Job: AddOrdenDePago}
			optimize.WorkQueue <- work
		}

	}).Catch(func(e try.E) { // Aquí se resuelven los errores
		beego.Info(e)
	})
}
