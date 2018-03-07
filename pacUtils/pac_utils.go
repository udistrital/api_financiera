package pacUtils

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/udistrital/api_financiera/models"
	"github.com/udistrital/utils_oas/formatdata"
	"github.com/udistrital/utils_oas/optimize"
)

func FunctionAfterExecIngresoPac(ctx *context.Context) {
	//beego.Info("Llamada...")
	var u map[string]interface{}
	ingreso := models.Ingreso{}
	var paramsIngreso []interface{}
	var tipo string
	if err := formatdata.FillStruct(ctx.Input.Data()["json"], &u); err == nil {
		if err = formatdata.FillStruct(u["Body"], &ingreso); err == nil && ingreso.Id != 0 {
			if err = formatdata.FillStruct(u["Type"], &tipo); err == nil && tipo == "success" {
				paramsIngreso = append(paramsIngreso, ingreso)
				work := optimize.WorkRequest{JobParameter: paramsIngreso, Job: (models.AddIngresoPac)}
				// Push the work onto the queue.
				optimize.WorkQueue <- work
			}

		}

	}
	//work := WorkRequest{JobParameter: ingreso, Job: FunctionJobExample}

	beego.Info("Work request queued")
}

func FunctionAfterExecEstadoOrdenP(ctx *context.Context) {
	var u map[string]interface{}
	var u2 map[string]interface{}
	var nuevoEstado map[string]interface{}
	var idEstado int
	var tipo string
	var parameters []interface{}

	egresoArr := make([]models.OrdenPago, 0)
	egreso := models.OrdenPago{}

	if err := json.Unmarshal(ctx.Input.RequestBody, &u); err == nil {

		if err = formatdata.FillStruct(u["NuevoEstado"], &nuevoEstado); err != nil {
			beego.Error(err.Error())
		} else {
			if err = formatdata.FillStruct(nuevoEstado["Id"], &idEstado); err != nil {
				beego.Error(err.Error())
			}
		}

		if err = formatdata.FillStruct(u["OrdenPago"], &egresoArr); err == nil {
			egreso = egresoArr[0]
		} else {
			beego.Error(err.Error())
		}
	} else {
		beego.Error(err.Error())
	}

	parameters = append(parameters, egreso)
	parameters = append(parameters, idEstado)
	beego.Error("valor u  ", u)
	if idEstado == 4 && egreso.Id != 0 {
		if err := formatdata.FillStruct(ctx.Input.Data()["json"], &u2); err == nil {
			if err := formatdata.FillStruct(u2["Type"], &tipo); err == nil && tipo == "success" {
				work := optimize.WorkRequest{JobParameter: parameters, Job: (models.AddEgresoPac)}
				// Push the work onto the queue.
				optimize.WorkQueue <- work

			} else {
				beego.Error("Error", err.Error())
			}
		}
		beego.Error("tipo ", tipo)
	}
}

func FunctionJobExample(parameter ...interface{}) (res interface{}) {
	beego.Info("Job's Parameter: ", parameter[0].(models.Ingreso).Id)
	return
}

func Init() {
	optimize.StartDispatcher(1, 200)
	beego.InsertFilter("/v1/ingreso/AprobacionPresupuestalIngreso", beego.AfterExec, FunctionAfterExecIngresoPac, false)
	beego.InsertFilter("/v1/orden_pago_estado_orden_pago/WorkFlowOrdenPago", beego.AfterExec, FunctionAfterExecEstadoOrdenP, false)
}
