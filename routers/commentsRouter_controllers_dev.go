package routers

import (
	"github.com/astaxie/beego"
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

}
