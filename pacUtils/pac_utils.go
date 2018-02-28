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
	var tipo string
	if err := formatdata.FillStruct(ctx.Input.Data()["json"], &u); err == nil {
		if err = formatdata.FillStruct(u["Body"], &ingreso); err == nil && ingreso.Id != 0 {
			if err = formatdata.FillStruct(u["Type"], &tipo); err == nil && tipo == "success" {
				work := optimize.WorkRequest{JobParameter: ingreso, Job: (models.AddIngresoPac)}
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
	var u2 []map[string]interface{}
	var u3 map[string]interface{}

	beego.Error(ctx.Input.RequestBody)

	if err := json.Unmarshal(ctx.Input.RequestBody, &u3); err == nil {
		beego.Error(u3)
	} else {
		beego.Info(err.Error())
	}

	if err := formatdata.FillStruct(ctx.Input.Data()["json"], &u); err == nil {
		if err = formatdata.FillStruct(u["Body"], &u2); err == nil {
			beego.Error(u2)
		} else {
			beego.Info(err.Error())
		}

	} else {
		beego.Info(err.Error())
	}

	beego.Info("Work request queued")
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
