package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TrConceptoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TrConceptoController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DisponibilidadController"],
		beego.ControllerComments{
			Method:           "Anular",
			Router:           `Anular`,
			AllowHTTPMethods: []string{"post"},
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DisponibilidadController"],
		beego.ControllerComments{
			Method:           "SaldoCdp",
			Router:           `SaldoCdp`,
			AllowHTTPMethods: []string{"post"},
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ApropiacionController"],
		beego.ControllerComments{
			Method:           "SaldoApropiacion",
			Router:           `SaldoApropiacion/:id`,
			AllowHTTPMethods: []string{"get"},
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenPagoController"],
		beego.ControllerComments{
			Method:           "RegistrarOpProveedor",
			Router:           `RegistrarOpProveedor`,
			AllowHTTPMethods: []string{"post"},
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenPagoController"],
		beego.ControllerComments{
			Method:           "ActualizarOpProveedor",
			Router:           `ActualizarOpProveedor`,
			AllowHTTPMethods: []string{"post"},
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RegistroPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RegistroPresupuestalController"],
		beego.ControllerComments{
			Method:           "ValorTotalRp",
			Router:           `ValorTotalRp/:id`,
			AllowHTTPMethods: []string{"get"},
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenPagoController"],
		beego.ControllerComments{
			Method:           "RegistrarOpPlanta",
			Router:           `RegistrarOpPlanta`,
			AllowHTTPMethods: []string{"post"},
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenPagoController"],
		beego.ControllerComments{
			Method:           "FechaActual",
			Router:           `FechaActual/:formato`,
			AllowHTTPMethods: []string{"get"},
			Params:           nil})
	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RubroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RubroController"],
		beego.ControllerComments{
			Method:           "RubroReporte",
			Router:           `RubroReporte/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

}
