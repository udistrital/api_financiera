package pacUtils

import (
	"fmt"

	"reflect"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/manucorporat/try"
	"github.com/udistrital/api_financiera/models"
	"github.com/udistrital/utils_oas/optimize"
	"github.com/udistrital/utils_oas/request"
)

func AddMovimientoApropiacion(parameter ...interface{}) (err interface{}) {
	try.This(func() {

		var rubroMongo []interface{}
		mongoApiURL := beego.AppConfig.String("MongoApi")
		// var respuesta interface{}
		Movimiento := parameter[0].(*models.MovimientoApropiacion)
		beego.Info("numero: ", Movimiento.NumeroMovimiento)
		beego.Info("estado_movimiento: apropbado")
		beego.Info("fecha_movmieinto: ", Movimiento.FechaMovimiento)
		beego.Info("numero_oficiio: ", Movimiento.Noficio)
		beego.Info("fecha Oficio: ", Movimiento.Foficio)
		beego.Info("descripcion: ", Movimiento.Descripcion)
		beego.Info("unidad_ejecutora: ", Movimiento.UnidadEjecutora)
		beego.Info("apropiacion_destino: ", Movimiento.MovimientoApropiacionDisponibilidadApropiacion[0].CuentaCredito.Rubro.Codigo)
		if Movimiento.MovimientoApropiacionDisponibilidadApropiacion[0].CuentaContraCredito.Rubro == nil {
			beego.Info("apropiacion_destino: ''")
		} else {
			beego.Info("apropiacion_origen: ", Movimiento.MovimientoApropiacionDisponibilidadApropiacion[0].CuentaContraCredito.Rubro)
		}
		beego.Info("valor: ", Movimiento.MovimientoApropiacionDisponibilidadApropiacion[0].Valor)
		beego.Info("Tipo movimiento: ", Movimiento.MovimientoApropiacionDisponibilidadApropiacion[0].TipoMovimientoApropiacion.Nombre)
		// beego.Info(Movimiento.MovimientoApropiacionDisponibilidadApropiacion[0].CuentaCredito.Rubro.Codigo)
		// beego.Info(Movimiento.MovimientoApropiacionDisponibilidadApropiacion[0].CuentaCredito.Rubro.Nombre)
		// beego.Info(Movimiento.MovimientoApropiacionDisponibilidadApropiacion[0].CuentaCredito.Valor) // Valor de la apropiacion inicial
		fmt.Printf("%g\n", Movimiento.MovimientoApropiacionDisponibilidadApropiacion[0].CuentaCredito.Valor)
		beego.Info(int(Movimiento.FechaMovimiento.Year()))
		request.GetJson("http://"+mongoApiURL+"rubro?query=codigo:"+Movimiento.MovimientoApropiacionDisponibilidadApropiacion[0].CuentaCredito.Rubro.Codigo, &rubroMongo)
		// request.GetJson("http://"+mongoApiURL+"")
		// beego.Info("rubroMongo: ", rubroMongo)
		// for key, value := range rubroMongo[0].(map[string]interface{}) {
		// 	fmt.Println("key: ", key)
		// 	fmt.Println("value: ", value)
		// }
		apropiaciones := rubroMongo[0].(map[string]interface{})["apropiaciones"]
		beego.Info(apropiaciones)
		beego.Info(reflect.TypeOf(apropiaciones))
		for i := 0; i < len(apropiaciones.([]interface{})); i++ {
			beego.Info(reflect.TypeOf(apropiaciones.([]interface{})[i].(map[string]interface{})["vigencia"]))
			if int(apropiaciones.([]interface{})[i].(map[string]interface{})["vigencia"].(float64)) == int(Movimiento.FechaMovimiento.Year()) {
				beego.Info(reflect.TypeOf(apropiaciones.([]interface{})[i].(map[string]interface{})["movimientos"]))
			}
		}

		// request.SendJson(beego.AppConfig.String("MongoApi")+"/movimiento", "POST",)
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
	}).Catch(func(e try.E) { // Aquí se resuelven los errores
		beego.Info(e)
	})

}
