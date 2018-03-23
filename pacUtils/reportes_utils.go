package pacUtils

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/manucorporat/try"
	"github.com/udistrital/api_financiera/models"
	"github.com/udistrital/utils_oas/optimize"
)

func AddMovimientoApropiacion(parameter ...interface{}) (err interface{}) {
	try.This(func() {
		Movimiento := parameter[0].(*models.MovimientoApropiacion)
		beego.Info(Movimiento)
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
			// beego.Info("save movimiento: ",ctx.Input.Data()["json"].([]models.Alert)[0].Body.(map[string]interface{})["Movimiento"].(*models.MovimientoApropiacion).MovimientoApropiacionDisponibilidadApropiacion[0].Valor)
			parameters = append(parameters, ctx.Input.Data()["json"].([]models.Alert)[0].Body.(map[string]interface{})["Movimiento"])

			work := optimize.WorkRequest{JobParameter: parameters, Job: AddMovimientoApropiacion}
			optimize.WorkQueue <- work
		}

		// 	switch response := ctx.Input.Data()["json"].([]models.Alert)[0].Type; response {
		// 	case "darwin":
		// 		fmt.Println("OS X.")
		// 	case "linux":
		// 		fmt.Println("Linux.")
		// 	default:
		// }
	}).Catch(func(e try.E) {
		beego.Info(e)
	})

}
