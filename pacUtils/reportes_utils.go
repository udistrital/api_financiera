package pacUtils

import (
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
			rubroMongo []interface{}
			movimiento = make(map[string]interface{})
			respuesta  interface{}
		)
		mongoApiURL := beego.AppConfig.String("MongoApi")
		Movimiento := parameter[0].(*models.MovimientoApropiacion)

		movimiento["numero"] = Movimiento.NumeroMovimiento
		movimiento["estado_movimiento"] = "Aprobado"
		movimiento["fecha_movimiento"] = Movimiento.FechaMovimiento.Format("2006-01-02")
		movimiento["numero_oficio"] = int(Movimiento.Noficio)
		movimiento["fecha_oficio"] = Movimiento.Foficio.Format("2006-01-02")
		movimiento["descripcion"] = Movimiento.Descripcion
		movimiento["unidad_ejecutora"] = Movimiento.UnidadEjecutora
		movimiento["apropiacion_destino"] = Movimiento.MovimientoApropiacionDisponibilidadApropiacion[0].CuentaCredito.Rubro.Codigo
		movimiento["apropiacion_origen"] = ""

		if Movimiento.MovimientoApropiacionDisponibilidadApropiacion[0].CuentaContraCredito != nil {
			movimiento["apropiacion_origen"] = Movimiento.MovimientoApropiacionDisponibilidadApropiacion[0].CuentaContraCredito.Rubro.Codigo
		}

		movimiento["valor"] = int64(Movimiento.MovimientoApropiacionDisponibilidadApropiacion[0].Valor)
		movimiento["tipo_movimiento"] = Movimiento.MovimientoApropiacionDisponibilidadApropiacion[0].TipoMovimientoApropiacion.Nombre

		request.GetJson("http://"+mongoApiURL+"rubro?query=codigo:"+Movimiento.MovimientoApropiacionDisponibilidadApropiacion[0].CuentaCredito.Rubro.Codigo, &rubroMongo)

		beego.Info(len(rubroMongo))
		beego.Info(rubroMongo[0])

		apropiaciones := rubroMongo[0].(map[string]interface{})["apropiaciones"]
		rubroId := rubroMongo[0].(map[string]interface{})["_id"].(string)

		for i := 0; i < len(apropiaciones.([]interface{})); i++ {

			if int(apropiaciones.([]interface{})[i].(map[string]interface{})["vigencia"].(float64)) == int(Movimiento.FechaMovimiento.Year()) {
				request.SendJson("http://"+mongoApiURL+"movimiento", "POST", &respuesta, movimiento)
				beego.Info(respuesta)
				if respuesta != nil {
					movimiento["_id"] = respuesta
					apropiaciones.([]interface{})[i].(map[string]interface{})["movimientos"] = append(apropiaciones.([]interface{})[i].(map[string]interface{})["movimientos"].([]interface{}), movimiento)
				}
			}
		}
		rubroMongo[0].(map[string]interface{})["apropiaciones"] = apropiaciones
		beego.Info("Movimiento: ", movimiento)
		beego.Info("apropiaciones con movimiento agregada: ", rubroMongo[0])
		request.SendJson("http://"+mongoApiURL+"rubro/"+rubroId, "PUT", &respuesta, rubroMongo[0])
	}).Catch(func(e try.E) {
		beego.Info(e)
	})
	return nil
}

func ReportesInit() {
	optimize.StartDispatcher(1, 200)
	beego.InsertFilter("/v1/movimiento_apropiacion/AprobarMovimietnoApropiacion", beego.AfterExec, saveMovimiento, false)
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
