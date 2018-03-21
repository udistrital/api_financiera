package pacUtils

import (
	//"encoding/json"
	
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/udistrital/api_financiera/models"
	"github.com/udistrital/utils_oas/formatdata"
	"github.com/udistrital/utils_oas/optimize"
)


func  FunctionAfterExecInsertarOP(ctx *context.Context) {
	//beego.Info("Llamada...")
	var u map[string]interface{}
	ordenPago := models.OrdenPago{}
	//var paramsIngreso []interface{}
	//var tipo string
	if(ctx.Input.Is("POST")){
		if err := formatdata.FillStruct(ctx.Input.Data()["json"], &u); err == nil {
			if err2 := formatdata.FillStruct(u, &ordenPago); err2 == nil && ordenPago.Id != 0 {
				models.CrearComprobante(ordenPago)
				//work := optimize.WorkRequest{JobParameter: paramsIngreso, Job: (models.AddIngresoPac)}
				//optimize.WorkQueue <- work
			}else{

			}

		}
	}


}


func InitComprobante() {
	optimize.StartDispatcher(1, 200)
	beego.InsertFilter("/v1/orden_pago", beego.AfterExec, FunctionAfterExecInsertarOP, false) //crud

}
