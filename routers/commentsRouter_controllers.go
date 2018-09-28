package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ActaDevolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ActaDevolucionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ActaDevolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ActaDevolucionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ActaDevolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ActaDevolucionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ActaDevolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ActaDevolucionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ActaDevolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ActaDevolucionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ActaInversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ActaInversionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ActaInversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ActaInversionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ActaInversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ActaInversionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ActaInversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ActaInversionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ActaInversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ActaInversionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AfectacionConceptoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AfectacionConceptoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AfectacionConceptoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AfectacionConceptoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AfectacionConceptoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AfectacionConceptoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AfectacionConceptoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AfectacionConceptoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AfectacionConceptoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AfectacionConceptoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AnulacionDisponibilidadApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AnulacionDisponibilidadApropiacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AnulacionDisponibilidadApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AnulacionDisponibilidadApropiacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AnulacionDisponibilidadApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AnulacionDisponibilidadApropiacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AnulacionDisponibilidadApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AnulacionDisponibilidadApropiacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AnulacionDisponibilidadApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AnulacionDisponibilidadApropiacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AnulacionDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AnulacionDisponibilidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AnulacionDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AnulacionDisponibilidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AnulacionDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AnulacionDisponibilidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AnulacionDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AnulacionDisponibilidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AnulacionDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AnulacionDisponibilidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AnulacionDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AnulacionDisponibilidadController"],
		beego.ControllerComments{
			Method: "TotalAnulacionDisponibilidad",
			Router: `/TotalAnulacionDisponibilidad/:vigencia`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AnulacionRegistroPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AnulacionRegistroPresupuestalController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AnulacionRegistroPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AnulacionRegistroPresupuestalController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AnulacionRegistroPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AnulacionRegistroPresupuestalController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AnulacionRegistroPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AnulacionRegistroPresupuestalController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AnulacionRegistroPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AnulacionRegistroPresupuestalController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AnulacionRegistroPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AnulacionRegistroPresupuestalController"],
		beego.ControllerComments{
			Method: "TotalAnulacionRegistroPresupuestal",
			Router: `/TotalAnulacionRegistroPresupuestal/:vigencia`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AnulacionRegistroPresupuestalDisponibilidadApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AnulacionRegistroPresupuestalDisponibilidadApropiacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AnulacionRegistroPresupuestalDisponibilidadApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AnulacionRegistroPresupuestalDisponibilidadApropiacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AnulacionRegistroPresupuestalDisponibilidadApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AnulacionRegistroPresupuestalDisponibilidadApropiacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AnulacionRegistroPresupuestalDisponibilidadApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AnulacionRegistroPresupuestalDisponibilidadApropiacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AnulacionRegistroPresupuestalDisponibilidadApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AnulacionRegistroPresupuestalDisponibilidadApropiacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AnulacionReservaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AnulacionReservaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AnulacionReservaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AnulacionReservaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AnulacionReservaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AnulacionReservaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AnulacionReservaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AnulacionReservaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AnulacionReservaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AnulacionReservaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ApropiacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ApropiacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ApropiacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ApropiacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ApropiacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ApropiacionController"],
		beego.ControllerComments{
			Method: "AprobarPresupuesto",
			Router: `/AprobacionAsignacionInicial/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ApropiacionController"],
		beego.ControllerComments{
			Method: "ArbolApropiaciones",
			Router: `/ArbolApropiaciones/:vigencia`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ApropiacionController"],
		beego.ControllerComments{
			Method: "GetApropiacionesHijo",
			Router: `/GetApropiacionesHijo/:vigencia`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ApropiacionController"],
		beego.ControllerComments{
			Method: "SaldoApropiacion",
			Router: `/SaldoApropiacion/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ApropiacionController"],
		beego.ControllerComments{
			Method: "SaldoApropiacionPadre",
			Router: `/SaldoApropiacionPadre/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ApropiacionController"],
		beego.ControllerComments{
			Method: "VigenciaApropiaciones",
			Router: `/VigenciaApropiaciones`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ArbolConceptosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ArbolConceptosController"],
		beego.ControllerComments{
			Method: "MakeTree",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ArbolPlanCuentasController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ArbolPlanCuentasController"],
		beego.ControllerComments{
			Method: "MakeTreeCuentas",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ArbolPlanCuentasController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ArbolPlanCuentasController"],
		beego.ControllerComments{
			Method: "DeleteBranch",
			Router: `/:idCuenta/:idPlan`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ArbolPlanCuentasController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ArbolPlanCuentasController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/:idPlan`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AvanceEstadoAvanceController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AvanceEstadoAvanceController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AvanceEstadoAvanceController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AvanceEstadoAvanceController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AvanceEstadoAvanceController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AvanceEstadoAvanceController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AvanceEstadoAvanceController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AvanceEstadoAvanceController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AvanceEstadoAvanceController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AvanceEstadoAvanceController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AvanceLegalizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AvanceLegalizacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AvanceLegalizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AvanceLegalizacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AvanceLegalizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AvanceLegalizacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AvanceLegalizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AvanceLegalizacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AvanceLegalizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AvanceLegalizacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AvanceLegalizacionCuentaEspecialController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AvanceLegalizacionCuentaEspecialController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AvanceLegalizacionCuentaEspecialController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AvanceLegalizacionCuentaEspecialController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AvanceLegalizacionCuentaEspecialController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AvanceLegalizacionCuentaEspecialController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AvanceLegalizacionCuentaEspecialController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AvanceLegalizacionCuentaEspecialController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AvanceLegalizacionCuentaEspecialController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AvanceLegalizacionCuentaEspecialController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AvanceLegalizacionTipoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AvanceLegalizacionTipoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AvanceLegalizacionTipoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AvanceLegalizacionTipoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AvanceLegalizacionTipoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AvanceLegalizacionTipoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AvanceLegalizacionTipoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AvanceLegalizacionTipoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AvanceLegalizacionTipoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AvanceLegalizacionTipoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AvanceLegalizacionTipoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AvanceLegalizacionTipoController"],
		beego.ControllerComments{
			Method: "AddEntireAvanceLegalizacionTipo",
			Router: `/AddEntireAvanceLegalizacionTipo`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AvanceLegalizacionTipoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AvanceLegalizacionTipoController"],
		beego.ControllerComments{
			Method: "GetLegalizationValue",
			Router: `/GetLegalizationValue/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AvanceLegalizacionTipoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:AvanceLegalizacionTipoController"],
		beego.ControllerComments{
			Method: "GetTaxesMovsLegalization",
			Router: `/GetTaxesMovsLegalization`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CalendarioTributarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CalendarioTributarioController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CalendarioTributarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CalendarioTributarioController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CalendarioTributarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CalendarioTributarioController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CalendarioTributarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CalendarioTributarioController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CalendarioTributarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CalendarioTributarioController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CalendarioTributarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CalendarioTributarioController"],
		beego.ControllerComments{
			Method: "GetMovimientos",
			Router: `/movimientos/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CancelacionInversionConceptoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CancelacionInversionConceptoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CancelacionInversionConceptoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CancelacionInversionConceptoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CancelacionInversionConceptoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CancelacionInversionConceptoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CancelacionInversionConceptoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CancelacionInversionConceptoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CancelacionInversionConceptoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CancelacionInversionConceptoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CancelacionInversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CancelacionInversionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CancelacionInversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CancelacionInversionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CancelacionInversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CancelacionInversionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CancelacionInversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CancelacionInversionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CancelacionInversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CancelacionInversionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CancelacionInversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CancelacionInversionController"],
		beego.ControllerComments{
			Method: "CreateCancelacion",
			Router: `CreateCancelacion/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CancelacionInversionEstadoCancelacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CancelacionInversionEstadoCancelacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CancelacionInversionEstadoCancelacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CancelacionInversionEstadoCancelacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CancelacionInversionEstadoCancelacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CancelacionInversionEstadoCancelacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CancelacionInversionEstadoCancelacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CancelacionInversionEstadoCancelacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CancelacionInversionEstadoCancelacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CancelacionInversionEstadoCancelacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CancelacionInversionEstadoCancelacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CancelacionInversionEstadoCancelacionController"],
		beego.ControllerComments{
			Method: "GetActiveCancelations",
			Router: `/GetActiveCancelations`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CancelacionInversionEstadoCancelacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CancelacionInversionEstadoCancelacionController"],
		beego.ControllerComments{
			Method: "AddEstadoCancelacionInversion",
			Router: `AddEstadoCancelacionInversion/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CancelacionInversionEstadoCancelacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CancelacionInversionEstadoCancelacionController"],
		beego.ControllerComments{
			Method: "GetCancelationQuantity",
			Router: `GetCancelationQuantity/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CancelacionInversionInversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CancelacionInversionInversionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CancelacionInversionInversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CancelacionInversionInversionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CancelacionInversionInversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CancelacionInversionInversionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CancelacionInversionInversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CancelacionInversionInversionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CancelacionInversionInversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CancelacionInversionInversionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CategoriaCompromisoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CategoriaCompromisoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CategoriaCompromisoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CategoriaCompromisoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CategoriaCompromisoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CategoriaCompromisoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CategoriaCompromisoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CategoriaCompromisoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CategoriaCompromisoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CategoriaCompromisoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CausalReintegroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CausalReintegroController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CausalReintegroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CausalReintegroController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CausalReintegroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CausalReintegroController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CausalReintegroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CausalReintegroController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CausalReintegroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CausalReintegroController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ChequeController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ChequeController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ChequeController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ChequeController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ChequeController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ChequeController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ChequeController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ChequeController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ChequeController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ChequeController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ChequeController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ChequeController"],
		beego.ControllerComments{
			Method: "CreateChequeEstado",
			Router: `/CreateChequeEstado`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ChequeController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ChequeController"],
		beego.ControllerComments{
			Method: "GetChequeRecordsNumber",
			Router: `/GetChequeRecordsNumber/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ChequeController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ChequeController"],
		beego.ControllerComments{
			Method: "GetChequeSumaOP",
			Router: `/GetChequeSumaOP/:idOP`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ChequeController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ChequeController"],
		beego.ControllerComments{
			Method: "GetNextChequeNumber",
			Router: `/GetNextChequeNumber/:idChequera`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ChequeEstadoChequeController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ChequeEstadoChequeController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ChequeEstadoChequeController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ChequeEstadoChequeController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ChequeEstadoChequeController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ChequeEstadoChequeController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ChequeEstadoChequeController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ChequeEstadoChequeController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ChequeEstadoChequeController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ChequeEstadoChequeController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ChequeEstadoChequeController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ChequeEstadoChequeController"],
		beego.ControllerComments{
			Method: "AddEstadoCheque",
			Router: `/AddEstadoCheque/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ChequeraController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ChequeraController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ChequeraController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ChequeraController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ChequeraController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ChequeraController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ChequeraController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ChequeraController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ChequeraController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ChequeraController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ChequeraController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ChequeraController"],
		beego.ControllerComments{
			Method: "CreateChequeraEstado",
			Router: `/CreateChequeraEstado`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ChequeraController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ChequeraController"],
		beego.ControllerComments{
			Method: "GetChequeraRecordsNumber",
			Router: `/GetChequeraRecordsNumber/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ChequeraEstadoChequeraController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ChequeraEstadoChequeraController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ChequeraEstadoChequeraController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ChequeraEstadoChequeraController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ChequeraEstadoChequeraController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ChequeraEstadoChequeraController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ChequeraEstadoChequeraController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ChequeraEstadoChequeraController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ChequeraEstadoChequeraController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ChequeraEstadoChequeraController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ChequeraEstadoChequeraController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ChequeraEstadoChequeraController"],
		beego.ControllerComments{
			Method: "AddEstadoChequera",
			Router: `/AddEstadoChequera/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CierrePeriodoPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CierrePeriodoPresupuestalController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CierrePeriodoPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CierrePeriodoPresupuestalController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CierrePeriodoPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CierrePeriodoPresupuestalController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CierrePeriodoPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CierrePeriodoPresupuestalController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CierrePeriodoPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CierrePeriodoPresupuestalController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ClaseTransaccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ClaseTransaccionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ClaseTransaccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ClaseTransaccionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ClaseTransaccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ClaseTransaccionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ClaseTransaccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ClaseTransaccionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ClaseTransaccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ClaseTransaccionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ComprobanteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ComprobanteController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ComprobanteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ComprobanteController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ComprobanteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ComprobanteController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ComprobanteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ComprobanteController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ComprobanteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ComprobanteController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CompromisoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CompromisoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CompromisoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CompromisoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CompromisoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CompromisoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CompromisoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CompromisoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CompromisoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CompromisoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ConceptoAvanceLegalizacionTipoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ConceptoAvanceLegalizacionTipoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ConceptoAvanceLegalizacionTipoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ConceptoAvanceLegalizacionTipoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ConceptoAvanceLegalizacionTipoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ConceptoAvanceLegalizacionTipoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ConceptoAvanceLegalizacionTipoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ConceptoAvanceLegalizacionTipoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ConceptoAvanceLegalizacionTipoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ConceptoAvanceLegalizacionTipoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ConceptoConceptoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ConceptoConceptoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ConceptoConceptoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ConceptoConceptoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ConceptoConceptoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ConceptoConceptoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ConceptoConceptoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ConceptoConceptoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ConceptoConceptoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ConceptoConceptoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ConceptoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ConceptoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ConceptoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ConceptoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ConceptoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ConceptoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ConceptoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ConceptoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ConceptoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ConceptoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ConceptoCuentaContableController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ConceptoCuentaContableController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ConceptoCuentaContableController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ConceptoCuentaContableController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ConceptoCuentaContableController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ConceptoCuentaContableController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ConceptoCuentaContableController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ConceptoCuentaContableController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ConceptoCuentaContableController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ConceptoCuentaContableController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ConceptoOrdenPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ConceptoOrdenPagoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ConceptoOrdenPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ConceptoOrdenPagoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ConceptoOrdenPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ConceptoOrdenPagoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ConceptoOrdenPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ConceptoOrdenPagoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ConceptoOrdenPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ConceptoOrdenPagoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ConceptoTesoralFacultadProyectoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ConceptoTesoralFacultadProyectoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ConceptoTesoralFacultadProyectoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ConceptoTesoralFacultadProyectoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ConceptoTesoralFacultadProyectoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ConceptoTesoralFacultadProyectoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ConceptoTesoralFacultadProyectoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ConceptoTesoralFacultadProyectoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ConceptoTesoralFacultadProyectoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ConceptoTesoralFacultadProyectoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CuentaBancariaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CuentaBancariaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CuentaBancariaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CuentaBancariaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CuentaBancariaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CuentaBancariaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CuentaBancariaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CuentaBancariaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CuentaBancariaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CuentaBancariaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CuentaContableController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CuentaContableController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CuentaContableController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CuentaContableController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CuentaContableController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CuentaContableController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CuentaContableController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CuentaContableController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CuentaContableController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CuentaContableController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CuentaDevolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CuentaDevolucionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CuentaDevolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CuentaDevolucionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CuentaDevolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CuentaDevolucionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CuentaDevolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CuentaDevolucionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CuentaDevolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CuentaDevolucionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CuentaEspecialController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CuentaEspecialController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CuentaEspecialController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CuentaEspecialController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CuentaEspecialController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CuentaEspecialController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CuentaEspecialController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CuentaEspecialController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CuentaEspecialController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CuentaEspecialController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DetalleCierrePeriodoPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DetalleCierrePeriodoPresupuestalController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DetalleCierrePeriodoPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DetalleCierrePeriodoPresupuestalController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DetalleCierrePeriodoPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DetalleCierrePeriodoPresupuestalController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DetalleCierrePeriodoPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DetalleCierrePeriodoPresupuestalController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DetalleCierrePeriodoPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DetalleCierrePeriodoPresupuestalController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DetallePacController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DetallePacController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DetallePacController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DetallePacController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DetallePacController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DetallePacController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DetallePacController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DetallePacController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DetallePacController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DetallePacController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DetallePacController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DetallePacController"],
		beego.ControllerComments{
			Method: "InsertarRegistros",
			Router: `/InsertarRegistros`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DevolucionEstadoDevolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DevolucionEstadoDevolucionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DevolucionEstadoDevolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DevolucionEstadoDevolucionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DevolucionEstadoDevolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DevolucionEstadoDevolucionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DevolucionEstadoDevolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DevolucionEstadoDevolucionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DevolucionEstadoDevolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DevolucionEstadoDevolucionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DevolucionTributariaConceptoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DevolucionTributariaConceptoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DevolucionTributariaConceptoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DevolucionTributariaConceptoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DevolucionTributariaConceptoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DevolucionTributariaConceptoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DevolucionTributariaConceptoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DevolucionTributariaConceptoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DevolucionTributariaConceptoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DevolucionTributariaConceptoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DevolucionTributariaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DevolucionTributariaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DevolucionTributariaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DevolucionTributariaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DevolucionTributariaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DevolucionTributariaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DevolucionTributariaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DevolucionTributariaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DevolucionTributariaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DevolucionTributariaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DevolucionTributariaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DevolucionTributariaController"],
		beego.ControllerComments{
			Method: "AddDevolucionTributaria",
			Router: `/AddDevolucionTributaria`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DevolucionTributariaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DevolucionTributariaController"],
		beego.ControllerComments{
			Method: "GetDevolucionRecordsNumber",
			Router: `/GetDevolucionRecordsNumber/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DevolucionTributariaEstadoDevolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DevolucionTributariaEstadoDevolucionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DevolucionTributariaEstadoDevolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DevolucionTributariaEstadoDevolucionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DevolucionTributariaEstadoDevolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DevolucionTributariaEstadoDevolucionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DevolucionTributariaEstadoDevolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DevolucionTributariaEstadoDevolucionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DevolucionTributariaEstadoDevolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DevolucionTributariaEstadoDevolucionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DevolucionTributariaEstadoDevolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DevolucionTributariaEstadoDevolucionController"],
		beego.ControllerComments{
			Method: "AddEstadoDevolTributaria",
			Router: `/AddEstadoDevolTributaria`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DevolucionTributariaMovimientoAsociadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DevolucionTributariaMovimientoAsociadoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DevolucionTributariaMovimientoAsociadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DevolucionTributariaMovimientoAsociadoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DevolucionTributariaMovimientoAsociadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DevolucionTributariaMovimientoAsociadoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DevolucionTributariaMovimientoAsociadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DevolucionTributariaMovimientoAsociadoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DevolucionTributariaMovimientoAsociadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DevolucionTributariaMovimientoAsociadoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DisponibilidadApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DisponibilidadApropiacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DisponibilidadApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DisponibilidadApropiacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DisponibilidadApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DisponibilidadApropiacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DisponibilidadApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DisponibilidadApropiacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DisponibilidadApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DisponibilidadApropiacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DisponibilidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DisponibilidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DisponibilidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DisponibilidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DisponibilidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DisponibilidadController"],
		beego.ControllerComments{
			Method: "AprobarAnulacionDisponibilidad",
			Router: `/AprobarAnulacion`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DisponibilidadController"],
		beego.ControllerComments{
			Method: "TotalDisponibilidades",
			Router: `/TotalDisponibilidades/:vigencia`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DisponibilidadProcesoExternoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DisponibilidadProcesoExternoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DisponibilidadProcesoExternoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DisponibilidadProcesoExternoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DisponibilidadProcesoExternoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DisponibilidadProcesoExternoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DisponibilidadProcesoExternoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DisponibilidadProcesoExternoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DisponibilidadProcesoExternoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DisponibilidadProcesoExternoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DocumentoDevolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DocumentoDevolucionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DocumentoDevolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DocumentoDevolucionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DocumentoDevolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DocumentoDevolucionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DocumentoDevolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DocumentoDevolucionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DocumentoDevolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DocumentoDevolucionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DocumentoGeneradorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DocumentoGeneradorController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DocumentoGeneradorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DocumentoGeneradorController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DocumentoGeneradorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DocumentoGeneradorController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DocumentoGeneradorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DocumentoGeneradorController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DocumentoGeneradorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:DocumentoGeneradorController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EntidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EntidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EntidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EntidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EntidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EntidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EntidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EntidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EntidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EntidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoAnulacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoAnulacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoAnulacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoAnulacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoAnulacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoAnulacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoAnulacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoAnulacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoAnulacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoAnulacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoApropiacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoApropiacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoApropiacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoApropiacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoApropiacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoAvanceController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoAvanceController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoAvanceController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoAvanceController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoAvanceController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoAvanceController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoAvanceController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoAvanceController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoAvanceController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoAvanceController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoCalendarioTributarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoCalendarioTributarioController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoCalendarioTributarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoCalendarioTributarioController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoCalendarioTributarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoCalendarioTributarioController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoCalendarioTributarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoCalendarioTributarioController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoCalendarioTributarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoCalendarioTributarioController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoCancelacionInversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoCancelacionInversionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoCancelacionInversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoCancelacionInversionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoCancelacionInversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoCancelacionInversionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoCancelacionInversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoCancelacionInversionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoCancelacionInversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoCancelacionInversionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoChequeController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoChequeController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoChequeController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoChequeController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoChequeController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoChequeController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoChequeController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoChequeController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoChequeController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoChequeController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoChequeraController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoChequeraController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoChequeraController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoChequeraController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoChequeraController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoChequeraController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoChequeraController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoChequeraController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoChequeraController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoChequeraController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoComprobanteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoComprobanteController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoComprobanteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoComprobanteController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoComprobanteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoComprobanteController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoComprobanteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoComprobanteController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoComprobanteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoComprobanteController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoCompromisoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoCompromisoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoCompromisoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoCompromisoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoCompromisoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoCompromisoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoCompromisoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoCompromisoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoCompromisoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoCompromisoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoDevolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoDevolucionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoDevolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoDevolucionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoDevolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoDevolucionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoDevolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoDevolucionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoDevolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoDevolucionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoDisponibilidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoDisponibilidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoDisponibilidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoDisponibilidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoDisponibilidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoGiroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoGiroController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoGiroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoGiroController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoGiroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoGiroController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoGiroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoGiroController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoGiroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoGiroController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoIngresoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoIngresoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoIngresoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoIngresoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoIngresoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoIngresoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoIngresoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoIngresoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoIngresoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoIngresoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoIngresoSinSituacionFondosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoIngresoSinSituacionFondosController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoIngresoSinSituacionFondosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoIngresoSinSituacionFondosController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoIngresoSinSituacionFondosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoIngresoSinSituacionFondosController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoIngresoSinSituacionFondosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoIngresoSinSituacionFondosController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoIngresoSinSituacionFondosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoIngresoSinSituacionFondosController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoInversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoInversionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoInversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoInversionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoInversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoInversionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoInversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoInversionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoInversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoInversionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoLegalizacionAvanceLegalizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoLegalizacionAvanceLegalizacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoLegalizacionAvanceLegalizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoLegalizacionAvanceLegalizacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoLegalizacionAvanceLegalizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoLegalizacionAvanceLegalizacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoLegalizacionAvanceLegalizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoLegalizacionAvanceLegalizacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoLegalizacionAvanceLegalizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoLegalizacionAvanceLegalizacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoLegalizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoLegalizacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoLegalizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoLegalizacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoLegalizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoLegalizacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoLegalizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoLegalizacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoLegalizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoLegalizacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoMovimientoApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoMovimientoApropiacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoMovimientoApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoMovimientoApropiacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoMovimientoApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoMovimientoApropiacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoMovimientoApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoMovimientoApropiacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoMovimientoApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoMovimientoApropiacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoMovimientoContableController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoMovimientoContableController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoMovimientoContableController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoMovimientoContableController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoMovimientoContableController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoMovimientoContableController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoMovimientoContableController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoMovimientoContableController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoMovimientoContableController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoMovimientoContableController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoOrdenPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoOrdenPagoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoOrdenPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoOrdenPagoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoOrdenPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoOrdenPagoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoOrdenPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoOrdenPagoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoOrdenPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoOrdenPagoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoRegistroPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoRegistroPresupuestalController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoRegistroPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoRegistroPresupuestalController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoRegistroPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoRegistroPresupuestalController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoRegistroPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoRegistroPresupuestalController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoRegistroPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoRegistroPresupuestalController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoReservaPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoReservaPresupuestalController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoReservaPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoReservaPresupuestalController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoReservaPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoReservaPresupuestalController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoReservaPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoReservaPresupuestalController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoReservaPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstadoReservaPresupuestalController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstructuraCuentasController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstructuraCuentasController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstructuraCuentasController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstructuraCuentasController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstructuraCuentasController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstructuraCuentasController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstructuraCuentasController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstructuraCuentasController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstructuraCuentasController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstructuraCuentasController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstructuraNivelesClasificacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstructuraNivelesClasificacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstructuraNivelesClasificacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstructuraNivelesClasificacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstructuraNivelesClasificacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstructuraNivelesClasificacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstructuraNivelesClasificacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstructuraNivelesClasificacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstructuraNivelesClasificacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EstructuraNivelesClasificacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EtapaAvanceController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EtapaAvanceController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EtapaAvanceController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EtapaAvanceController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EtapaAvanceController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EtapaAvanceController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EtapaAvanceController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EtapaAvanceController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EtapaAvanceController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:EtapaAvanceController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:FormaIngresoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:FormaIngresoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:FormaIngresoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:FormaIngresoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:FormaIngresoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:FormaIngresoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:FormaIngresoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:FormaIngresoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:FormaIngresoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:FormaIngresoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:FormaPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:FormaPagoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:FormaPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:FormaPagoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:FormaPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:FormaPagoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:FormaPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:FormaPagoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:FormaPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:FormaPagoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:FuenteFinanciamientoApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:FuenteFinanciamientoApropiacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:FuenteFinanciamientoApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:FuenteFinanciamientoApropiacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:FuenteFinanciamientoApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:FuenteFinanciamientoApropiacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:FuenteFinanciamientoApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:FuenteFinanciamientoApropiacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:FuenteFinanciamientoApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:FuenteFinanciamientoApropiacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:FuenteFinanciamientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:FuenteFinanciamientoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:FuenteFinanciamientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:FuenteFinanciamientoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:FuenteFinanciamientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:FuenteFinanciamientoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:FuenteFinanciamientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:FuenteFinanciamientoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:FuenteFinanciamientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:FuenteFinanciamientoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:GiroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:GiroController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:GiroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:GiroController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:GiroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:GiroController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:GiroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:GiroController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:GiroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:GiroController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:GiroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:GiroController"],
		beego.ControllerComments{
			Method: "RegistrarGiro",
			Router: `RegistrarGiro`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:GiroDetalleController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:GiroDetalleController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:GiroDetalleController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:GiroDetalleController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:GiroDetalleController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:GiroDetalleController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:GiroDetalleController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:GiroDetalleController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:GiroDetalleController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:GiroDetalleController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:GiroEstadoGiroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:GiroEstadoGiroController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:GiroEstadoGiroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:GiroEstadoGiroController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:GiroEstadoGiroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:GiroEstadoGiroController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:GiroEstadoGiroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:GiroEstadoGiroController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:GiroEstadoGiroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:GiroEstadoGiroController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:HistoricoInversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:HistoricoInversionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:HistoricoInversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:HistoricoInversionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:HistoricoInversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:HistoricoInversionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:HistoricoInversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:HistoricoInversionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:HistoricoInversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:HistoricoInversionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:HomologacionComprobantesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:HomologacionComprobantesController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:HomologacionComprobantesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:HomologacionComprobantesController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:HomologacionComprobantesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:HomologacionComprobantesController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:HomologacionComprobantesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:HomologacionComprobantesController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:HomologacionComprobantesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:HomologacionComprobantesController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:HomologacionConceptoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:HomologacionConceptoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:HomologacionConceptoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:HomologacionConceptoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:HomologacionConceptoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:HomologacionConceptoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:HomologacionConceptoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:HomologacionConceptoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:HomologacionConceptoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:HomologacionConceptoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:HomologacionConceptoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:HomologacionConceptoController"],
		beego.ControllerComments{
			Method: "ActualizarHomologacionConcepto",
			Router: `ActualizarHomologacionConcepto`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:HomologacionConceptoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:HomologacionConceptoController"],
		beego.ControllerComments{
			Method: "RegistrarHomologacionConcepto",
			Router: `RegistrarHomologacionConcepto`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:HomologacionDescuentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:HomologacionDescuentoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:HomologacionDescuentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:HomologacionDescuentoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:HomologacionDescuentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:HomologacionDescuentoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:HomologacionDescuentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:HomologacionDescuentoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:HomologacionDescuentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:HomologacionDescuentoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:InformacionAdicionalBancoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:InformacionAdicionalBancoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:InformacionAdicionalBancoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:InformacionAdicionalBancoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:InformacionAdicionalBancoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:InformacionAdicionalBancoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:InformacionAdicionalBancoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:InformacionAdicionalBancoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:InformacionAdicionalBancoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:InformacionAdicionalBancoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:IngresoConceptoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:IngresoConceptoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:IngresoConceptoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:IngresoConceptoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:IngresoConceptoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:IngresoConceptoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:IngresoConceptoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:IngresoConceptoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:IngresoConceptoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:IngresoConceptoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:IngresoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:IngresoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:IngresoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:IngresoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:IngresoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:IngresoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:IngresoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:IngresoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:IngresoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:IngresoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:IngresoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:IngresoController"],
		beego.ControllerComments{
			Method: "AprobacionContableIngreso",
			Router: `/AprobacionContableIngreso`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:IngresoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:IngresoController"],
		beego.ControllerComments{
			Method: "AprobacionPresupuestalIngreso",
			Router: `/AprobacionPresupuestalIngreso`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:IngresoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:IngresoController"],
		beego.ControllerComments{
			Method: "CreateIngresos",
			Router: `/CreateIngresos`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:IngresoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:IngresoController"],
		beego.ControllerComments{
			Method: "RechazoContableIngreso",
			Router: `/RechazoContableIngreso`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:IngresoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:IngresoController"],
		beego.ControllerComments{
			Method: "RechazoPresupuestalIngreso",
			Router: `/RechazoPresupuestalIngreso`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:IngresoEstadoIngresoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:IngresoEstadoIngresoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:IngresoEstadoIngresoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:IngresoEstadoIngresoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:IngresoEstadoIngresoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:IngresoEstadoIngresoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:IngresoEstadoIngresoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:IngresoEstadoIngresoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:IngresoEstadoIngresoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:IngresoEstadoIngresoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:IngresoSinSituacionFondosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:IngresoSinSituacionFondosController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:IngresoSinSituacionFondosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:IngresoSinSituacionFondosController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:IngresoSinSituacionFondosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:IngresoSinSituacionFondosController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:IngresoSinSituacionFondosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:IngresoSinSituacionFondosController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:IngresoSinSituacionFondosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:IngresoSinSituacionFondosController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:IngresoSinSituacionFondosEstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:IngresoSinSituacionFondosEstadoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:IngresoSinSituacionFondosEstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:IngresoSinSituacionFondosEstadoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:IngresoSinSituacionFondosEstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:IngresoSinSituacionFondosEstadoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:IngresoSinSituacionFondosEstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:IngresoSinSituacionFondosEstadoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:IngresoSinSituacionFondosEstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:IngresoSinSituacionFondosEstadoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:IngresoSinSituacionFondosEstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:IngresoSinSituacionFondosEstadoController"],
		beego.ControllerComments{
			Method: "ChangeExistingStates",
			Router: `ChangeExistingStates/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:InversionConceptoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:InversionConceptoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:InversionConceptoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:InversionConceptoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:InversionConceptoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:InversionConceptoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:InversionConceptoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:InversionConceptoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:InversionConceptoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:InversionConceptoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:InversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:InversionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:InversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:InversionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:InversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:InversionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:InversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:InversionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:InversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:InversionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:InversionEstadoInversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:InversionEstadoInversionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:InversionEstadoInversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:InversionEstadoInversionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:InversionEstadoInversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:InversionEstadoInversionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:InversionEstadoInversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:InversionEstadoInversionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:InversionEstadoInversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:InversionEstadoInversionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:InversionEstadoInversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:InversionEstadoInversionController"],
		beego.ControllerComments{
			Method: "AddEstadoInv",
			Router: `/AddEstadoInv`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:InversionesActaInversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:InversionesActaInversionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:InversionesActaInversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:InversionesActaInversionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:InversionesActaInversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:InversionesActaInversionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:InversionesActaInversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:InversionesActaInversionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:InversionesActaInversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:InversionesActaInversionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:InversionesActaInversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:InversionesActaInversionController"],
		beego.ControllerComments{
			Method: "CreateInversion",
			Router: `/CreateInversion`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ModificacionPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ModificacionPresupuestalController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ModificacionPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ModificacionPresupuestalController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ModificacionPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ModificacionPresupuestalController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ModificacionPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ModificacionPresupuestalController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ModificacionPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ModificacionPresupuestalController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ModuloKronosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ModuloKronosController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ModuloKronosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ModuloKronosController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ModuloKronosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ModuloKronosController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ModuloKronosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ModuloKronosController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ModuloKronosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ModuloKronosController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:MovimientoApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:MovimientoApropiacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:MovimientoApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:MovimientoApropiacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:MovimientoApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:MovimientoApropiacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:MovimientoApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:MovimientoApropiacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:MovimientoApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:MovimientoApropiacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:MovimientoApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:MovimientoApropiacionController"],
		beego.ControllerComments{
			Method: "AprobarMovimietnoApropiacion",
			Router: `/AprobarMovimietnoApropiacion`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:MovimientoApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:MovimientoApropiacionController"],
		beego.ControllerComments{
			Method: "RegistroSolicitudMovimientoApropiacion",
			Router: `/RegistroSolicitudMovimientoApropiacion`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:MovimientoApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:MovimientoApropiacionController"],
		beego.ControllerComments{
			Method: "TotalMovimientosApropiacion",
			Router: `/TotalMovimientosApropiacion/:vigencia`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:MovimientoApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:MovimientoApropiacionController"],
		beego.ControllerComments{
			Method: "GetMovimientosApropiacionByApropiacion",
			Router: `GetMovimientosApropiacionByApropiacion/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:MovimientoApropiacionDisponibilidadApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:MovimientoApropiacionDisponibilidadApropiacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:MovimientoApropiacionDisponibilidadApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:MovimientoApropiacionDisponibilidadApropiacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:MovimientoApropiacionDisponibilidadApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:MovimientoApropiacionDisponibilidadApropiacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:MovimientoApropiacionDisponibilidadApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:MovimientoApropiacionDisponibilidadApropiacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:MovimientoApropiacionDisponibilidadApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:MovimientoApropiacionDisponibilidadApropiacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:MovimientoContableController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:MovimientoContableController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:MovimientoContableController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:MovimientoContableController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:MovimientoContableController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:MovimientoContableController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:MovimientoContableController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:MovimientoContableController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:MovimientoContableController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:MovimientoContableController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:MovimientoContableController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:MovimientoContableController"],
		beego.ControllerComments{
			Method: "PostArray",
			Router: `PostArray/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:MovimientoFuenteFinanciamientoApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:MovimientoFuenteFinanciamientoApropiacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:MovimientoFuenteFinanciamientoApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:MovimientoFuenteFinanciamientoApropiacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:MovimientoFuenteFinanciamientoApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:MovimientoFuenteFinanciamientoApropiacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:MovimientoFuenteFinanciamientoApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:MovimientoFuenteFinanciamientoApropiacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:MovimientoFuenteFinanciamientoApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:MovimientoFuenteFinanciamientoApropiacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:NivelClasificacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:NivelClasificacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:NivelClasificacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:NivelClasificacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:NivelClasificacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:NivelClasificacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:NivelClasificacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:NivelClasificacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:NivelClasificacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:NivelClasificacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:NivelClasificacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:NivelClasificacionController"],
		beego.ControllerComments{
			Method: "GetFirstNivel",
			Router: `/primer_nivel`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ObservacionCalendarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ObservacionCalendarioController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ObservacionCalendarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ObservacionCalendarioController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ObservacionCalendarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ObservacionCalendarioController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ObservacionCalendarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ObservacionCalendarioController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ObservacionCalendarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ObservacionCalendarioController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenDevolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenDevolucionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenDevolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenDevolucionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenDevolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenDevolucionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenDevolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenDevolucionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenDevolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenDevolucionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenDevolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenDevolucionController"],
		beego.ControllerComments{
			Method: "AddDevolutionOrder",
			Router: `/AddDevolutionOrder`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenDevolucionEstadoDevolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenDevolucionEstadoDevolucionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenDevolucionEstadoDevolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenDevolucionEstadoDevolucionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenDevolucionEstadoDevolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenDevolucionEstadoDevolucionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenDevolucionEstadoDevolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenDevolucionEstadoDevolucionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenDevolucionEstadoDevolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenDevolucionEstadoDevolucionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenDevolucionEstadoDevolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenDevolucionEstadoDevolucionController"],
		beego.ControllerComments{
			Method: "AddEstadoOrdenDevol",
			Router: `/AddEstadoOrdenDevol`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenDevolucionSolicitudDevolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenDevolucionSolicitudDevolucionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenDevolucionSolicitudDevolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenDevolucionSolicitudDevolucionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenDevolucionSolicitudDevolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenDevolucionSolicitudDevolucionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenDevolucionSolicitudDevolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenDevolucionSolicitudDevolucionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenDevolucionSolicitudDevolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenDevolucionSolicitudDevolucionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenPagoAvanceLegalizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenPagoAvanceLegalizacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenPagoAvanceLegalizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenPagoAvanceLegalizacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenPagoAvanceLegalizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenPagoAvanceLegalizacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenPagoAvanceLegalizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenPagoAvanceLegalizacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenPagoAvanceLegalizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenPagoAvanceLegalizacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenPagoAvanceLegalizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenPagoAvanceLegalizacionController"],
		beego.ControllerComments{
			Method: "AddOPAvance",
			Router: `/AddOPAvance`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenPagoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenPagoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenPagoController"],
		beego.ControllerComments{
			Method: "FechaActual",
			Router: `/:formato`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenPagoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenPagoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenPagoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenPagoController"],
		beego.ControllerComments{
			Method: "GetOrdenPagoByEstado",
			Router: `/GetOrdenPagoByEstado`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenPagoController"],
		beego.ControllerComments{
			Method: "ActualizarOpProveedor",
			Router: `ActualizarOpProveedor`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenPagoController"],
		beego.ControllerComments{
			Method: "RegistrarOpProveedor",
			Router: `RegistrarOpProveedor`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenPagoController"],
		beego.ControllerComments{
			Method: "ValorTotal",
			Router: `ValorTotal/:id`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenPagoCuentaEspecialController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenPagoCuentaEspecialController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenPagoCuentaEspecialController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenPagoCuentaEspecialController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenPagoCuentaEspecialController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenPagoCuentaEspecialController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenPagoCuentaEspecialController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenPagoCuentaEspecialController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenPagoCuentaEspecialController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenPagoCuentaEspecialController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenPagoEstadoOrdenPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenPagoEstadoOrdenPagoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenPagoEstadoOrdenPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenPagoEstadoOrdenPagoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenPagoEstadoOrdenPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenPagoEstadoOrdenPagoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenPagoEstadoOrdenPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenPagoEstadoOrdenPagoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenPagoEstadoOrdenPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenPagoEstadoOrdenPagoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenPagoEstadoOrdenPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenPagoEstadoOrdenPagoController"],
		beego.ControllerComments{
			Method: "WorkFlowOrdenPago",
			Router: `WorkFlowOrdenPago`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenPagoRegistroPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenPagoRegistroPresupuestalController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenPagoRegistroPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenPagoRegistroPresupuestalController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenPagoRegistroPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenPagoRegistroPresupuestalController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenPagoRegistroPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenPagoRegistroPresupuestalController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenPagoRegistroPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:OrdenPagoRegistroPresupuestalController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:PacController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:PacController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:PacController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:PacController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:PacController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:PacController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:PacController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:PacController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:PacController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:PacController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:Pasivos_fenecidosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:Pasivos_fenecidosController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:PeriodoContableController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:PeriodoContableController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:PeriodoContableController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:PeriodoContableController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:PeriodoContableController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:PeriodoContableController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:PeriodoContableController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:PeriodoContableController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:PeriodoContableController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:PeriodoContableController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:PeriodoPlanCuentasController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:PeriodoPlanCuentasController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:PeriodoPlanCuentasController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:PeriodoPlanCuentasController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:PeriodoPlanCuentasController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:PeriodoPlanCuentasController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:PeriodoPlanCuentasController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:PeriodoPlanCuentasController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:PeriodoPlanCuentasController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:PeriodoPlanCuentasController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:PlanCuentasController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:PlanCuentasController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:PlanCuentasController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:PlanCuentasController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:PlanCuentasController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:PlanCuentasController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:PlanCuentasController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:PlanCuentasController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:PlanCuentasController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:PlanCuentasController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:PresupuestoAvanceController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:PresupuestoAvanceController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:PresupuestoAvanceController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:PresupuestoAvanceController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:PresupuestoAvanceController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:PresupuestoAvanceController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:PresupuestoAvanceController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:PresupuestoAvanceController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:PresupuestoAvanceController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:PresupuestoAvanceController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ProductoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ProductoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ProductoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ProductoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ProductoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ProductoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ProductoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ProductoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ProductoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ProductoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ProductoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ProductoController"],
		beego.ControllerComments{
			Method: "TotalProductos",
			Router: `/TotalProductos/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ProductoRubroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ProductoRubroController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ProductoRubroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ProductoRubroController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ProductoRubroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ProductoRubroController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ProductoRubroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ProductoRubroController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ProductoRubroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ProductoRubroController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ProductoRubroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ProductoRubroController"],
		beego.ControllerComments{
			Method: "AddProductoRubrotr",
			Router: `/AddProductoRubrotr/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ProductoRubroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ProductoRubroController"],
		beego.ControllerComments{
			Method: "SetVariacionProducto",
			Router: `/SetVariacionProducto/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RazonDevolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RazonDevolucionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RazonDevolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RazonDevolucionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RazonDevolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RazonDevolucionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RazonDevolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RazonDevolucionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RazonDevolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RazonDevolucionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RegistroComprobantesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RegistroComprobantesController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RegistroComprobantesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RegistroComprobantesController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RegistroComprobantesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RegistroComprobantesController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RegistroComprobantesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RegistroComprobantesController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RegistroComprobantesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RegistroComprobantesController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RegistroPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RegistroPresupuestalController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RegistroPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RegistroPresupuestalController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RegistroPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RegistroPresupuestalController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RegistroPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RegistroPresupuestalController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RegistroPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RegistroPresupuestalController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RegistroPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RegistroPresupuestalController"],
		beego.ControllerComments{
			Method: "AprobarAnulacionRp",
			Router: `/AprobarAnulacion`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RegistroPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RegistroPresupuestalController"],
		beego.ControllerComments{
			Method: "TotalRp",
			Router: `/TotalRp/:vigencia`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RegistroPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RegistroPresupuestalController"],
		beego.ControllerComments{
			Method: "Anular",
			Router: `Anular/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RegistroPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RegistroPresupuestalController"],
		beego.ControllerComments{
			Method: "SaldoRp",
			Router: `SaldoRp/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RegistroPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RegistroPresupuestalController"],
		beego.ControllerComments{
			Method: "ValorActualRp",
			Router: `ValorActualRp/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RegistroPresupuestalDisponibilidadApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RegistroPresupuestalDisponibilidadApropiacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RegistroPresupuestalDisponibilidadApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RegistroPresupuestalDisponibilidadApropiacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RegistroPresupuestalDisponibilidadApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RegistroPresupuestalDisponibilidadApropiacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RegistroPresupuestalDisponibilidadApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RegistroPresupuestalDisponibilidadApropiacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RegistroPresupuestalDisponibilidadApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RegistroPresupuestalDisponibilidadApropiacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ReintegroAvanceLegalizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ReintegroAvanceLegalizacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ReintegroAvanceLegalizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ReintegroAvanceLegalizacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ReintegroAvanceLegalizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ReintegroAvanceLegalizacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ReintegroAvanceLegalizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ReintegroAvanceLegalizacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ReintegroAvanceLegalizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ReintegroAvanceLegalizacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ReintegroAvanceLegalizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ReintegroAvanceLegalizacionController"],
		beego.ControllerComments{
			Method: "AddReintegroAvance",
			Router: `/AddReintegroAvance`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ReintegroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ReintegroController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ReintegroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ReintegroController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ReintegroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ReintegroController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ReintegroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ReintegroController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ReintegroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ReintegroController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ReintegroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ReintegroController"],
		beego.ControllerComments{
			Method: "Create",
			Router: `/Create`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ReintegroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ReintegroController"],
		beego.ControllerComments{
			Method: "GetAllReintegroDisponible",
			Router: `/GetAllReintegroDisponible`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RequisitoAvanceController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RequisitoAvanceController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RequisitoAvanceController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RequisitoAvanceController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RequisitoAvanceController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RequisitoAvanceController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RequisitoAvanceController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RequisitoAvanceController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RequisitoAvanceController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RequisitoAvanceController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RequisitoTipoAvanceController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RequisitoTipoAvanceController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RequisitoTipoAvanceController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RequisitoTipoAvanceController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RequisitoTipoAvanceController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RequisitoTipoAvanceController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RequisitoTipoAvanceController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RequisitoTipoAvanceController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RequisitoTipoAvanceController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RequisitoTipoAvanceController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ReservaPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ReservaPresupuestalController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ReservaPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ReservaPresupuestalController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ReservaPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ReservaPresupuestalController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ReservaPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ReservaPresupuestalController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ReservaPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:ReservaPresupuestalController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RubroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RubroController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RubroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RubroController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RubroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RubroController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RubroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RubroController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RubroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RubroController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RubroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RubroController"],
		beego.ControllerComments{
			Method: "GetIngresoCierre",
			Router: `/GetIngresoCierre`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RubroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RubroController"],
		beego.ControllerComments{
			Method: "GetPacValue",
			Router: `/GetPacValue`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RubroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RubroController"],
		beego.ControllerComments{
			Method: "GetRubroIngreso",
			Router: `/GetRubroIngreso`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RubroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RubroController"],
		beego.ControllerComments{
			Method: "GetRubroOrdenPago",
			Router: `/GetRubroOrdenPago`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RubroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RubroController"],
		beego.ControllerComments{
			Method: "GetRubroPac",
			Router: `/GetRubroPac`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RubroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RubroController"],
		beego.ControllerComments{
			Method: "GetSumbySource",
			Router: `/GetSumbySource`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RubroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RubroController"],
		beego.ControllerComments{
			Method: "ApropiacionReporte",
			Router: `ApropiacionReporte/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RubroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RubroController"],
		beego.ControllerComments{
			Method: "ArbolRubros",
			Router: `ArbolRubros/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RubroHomologadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RubroHomologadoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RubroHomologadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RubroHomologadoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RubroHomologadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RubroHomologadoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RubroHomologadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RubroHomologadoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RubroHomologadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RubroHomologadoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RubroHomologadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RubroHomologadoController"],
		beego.ControllerComments{
			Method: "ArbolRubros",
			Router: `/ArbolRubros`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RubroHomologadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RubroHomologadoController"],
		beego.ControllerComments{
			Method: "GetRecordsNumberByEntity",
			Router: `/GetRecordsNumberByEntity`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RubroHomologadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RubroHomologadoController"],
		beego.ControllerComments{
			Method: "GetRecordsNumberRubroHomologadoById",
			Router: `/GetRecordsNumberRubroHomologadoById/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RubroHomologadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RubroHomologadoController"],
		beego.ControllerComments{
			Method: "GetRecordsNumberRubroHomologadoRubroById",
			Router: `/GetRecordsNumberRubroHomologadoRubroById/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RubroHomologadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RubroHomologadoController"],
		beego.ControllerComments{
			Method: "GetParentHomologation",
			Router: `GetParentHomologation/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RubroHomologadoRubroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RubroHomologadoRubroController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RubroHomologadoRubroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RubroHomologadoRubroController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RubroHomologadoRubroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RubroHomologadoRubroController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RubroHomologadoRubroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RubroHomologadoRubroController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RubroHomologadoRubroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RubroHomologadoRubroController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RubroRubroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RubroRubroController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RubroRubroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RubroRubroController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RubroRubroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RubroRubroController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RubroRubroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RubroRubroController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RubroRubroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:RubroRubroController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:SaldoCuentaContableController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:SaldoCuentaContableController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:SaldoCuentaContableController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:SaldoCuentaContableController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:SaldoCuentaContableController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:SaldoCuentaContableController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:SaldoCuentaContableController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:SaldoCuentaContableController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:SaldoCuentaContableController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:SaldoCuentaContableController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:SolicitudAvanceController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:SolicitudAvanceController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:SolicitudAvanceController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:SolicitudAvanceController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:SolicitudAvanceController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:SolicitudAvanceController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:SolicitudAvanceController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:SolicitudAvanceController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:SolicitudAvanceController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:SolicitudAvanceController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:SolicitudAvanceController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:SolicitudAvanceController"],
		beego.ControllerComments{
			Method: "TrSolicitudAvance",
			Router: `TrSolicitudAvance/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:SolicitudDevolucionConceptoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:SolicitudDevolucionConceptoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:SolicitudDevolucionConceptoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:SolicitudDevolucionConceptoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:SolicitudDevolucionConceptoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:SolicitudDevolucionConceptoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:SolicitudDevolucionConceptoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:SolicitudDevolucionConceptoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:SolicitudDevolucionConceptoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:SolicitudDevolucionConceptoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:SolicitudDevolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:SolicitudDevolucionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:SolicitudDevolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:SolicitudDevolucionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:SolicitudDevolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:SolicitudDevolucionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:SolicitudDevolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:SolicitudDevolucionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:SolicitudDevolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:SolicitudDevolucionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:SolicitudDevolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:SolicitudDevolucionController"],
		beego.ControllerComments{
			Method: "AddDevolution",
			Router: `/AddDevolution`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:SolicitudDevolucionEstadoDevolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:SolicitudDevolucionEstadoDevolucionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:SolicitudDevolucionEstadoDevolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:SolicitudDevolucionEstadoDevolucionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:SolicitudDevolucionEstadoDevolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:SolicitudDevolucionEstadoDevolucionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:SolicitudDevolucionEstadoDevolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:SolicitudDevolucionEstadoDevolucionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:SolicitudDevolucionEstadoDevolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:SolicitudDevolucionEstadoDevolucionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:SolicitudRequisitoTipoAvanceController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:SolicitudRequisitoTipoAvanceController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:SolicitudRequisitoTipoAvanceController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:SolicitudRequisitoTipoAvanceController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:SolicitudRequisitoTipoAvanceController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:SolicitudRequisitoTipoAvanceController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:SolicitudRequisitoTipoAvanceController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:SolicitudRequisitoTipoAvanceController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:SolicitudRequisitoTipoAvanceController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:SolicitudRequisitoTipoAvanceController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:SolicitudRequisitoTipoAvanceController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:SolicitudRequisitoTipoAvanceController"],
		beego.ControllerComments{
			Method: "TrValidarAvance",
			Router: `TrValidarAvance/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:SolicitudTipoAvanceController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:SolicitudTipoAvanceController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:SolicitudTipoAvanceController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:SolicitudTipoAvanceController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:SolicitudTipoAvanceController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:SolicitudTipoAvanceController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:SolicitudTipoAvanceController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:SolicitudTipoAvanceController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:SolicitudTipoAvanceController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:SolicitudTipoAvanceController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:SubTipoOrdenPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:SubTipoOrdenPagoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:SubTipoOrdenPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:SubTipoOrdenPagoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:SubTipoOrdenPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:SubTipoOrdenPagoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:SubTipoOrdenPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:SubTipoOrdenPagoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:SubTipoOrdenPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:SubTipoOrdenPagoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoAfectacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoAfectacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoAfectacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoAfectacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoAfectacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoAfectacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoAfectacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoAfectacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoAfectacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoAfectacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoAnulacionPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoAnulacionPresupuestalController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoAnulacionPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoAnulacionPresupuestalController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoAnulacionPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoAnulacionPresupuestalController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoAnulacionPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoAnulacionPresupuestalController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoAnulacionPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoAnulacionPresupuestalController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoAvanceController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoAvanceController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoAvanceController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoAvanceController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoAvanceController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoAvanceController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoAvanceController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoAvanceController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoAvanceController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoAvanceController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoAvanceLegalizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoAvanceLegalizacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoAvanceLegalizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoAvanceLegalizacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoAvanceLegalizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoAvanceLegalizacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoAvanceLegalizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoAvanceLegalizacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoAvanceLegalizacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoAvanceLegalizacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoComprobanteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoComprobanteController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoComprobanteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoComprobanteController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoComprobanteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoComprobanteController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoComprobanteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoComprobanteController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoComprobanteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoComprobanteController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoCompromisoTesoralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoCompromisoTesoralController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoCompromisoTesoralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoCompromisoTesoralController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoCompromisoTesoralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoCompromisoTesoralController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoCompromisoTesoralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoCompromisoTesoralController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoCompromisoTesoralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoCompromisoTesoralController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoConceptoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoConceptoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoConceptoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoConceptoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoConceptoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoConceptoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoConceptoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoConceptoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoConceptoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoConceptoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoCuentaBancariaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoCuentaBancariaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoCuentaBancariaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoCuentaBancariaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoCuentaBancariaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoCuentaBancariaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoCuentaBancariaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoCuentaBancariaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoCuentaBancariaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoCuentaBancariaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoCuentaEspecialController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoCuentaEspecialController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoCuentaEspecialController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoCuentaEspecialController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoCuentaEspecialController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoCuentaEspecialController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoCuentaEspecialController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoCuentaEspecialController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoCuentaEspecialController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoCuentaEspecialController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoDevolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoDevolucionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoDevolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoDevolucionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoDevolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoDevolucionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoDevolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoDevolucionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoDevolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoDevolucionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoDisponibilidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoDisponibilidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoDisponibilidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoDisponibilidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoDisponibilidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoDocumentoAfectanteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoDocumentoAfectanteController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoDocumentoAfectanteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoDocumentoAfectanteController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoDocumentoAfectanteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoDocumentoAfectanteController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoDocumentoAfectanteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoDocumentoAfectanteController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoDocumentoAfectanteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoDocumentoAfectanteController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoEntidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoEntidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoEntidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoEntidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoEntidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoEntidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoEntidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoEntidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoEntidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoEntidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoFuenteFinanciamientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoFuenteFinanciamientoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoFuenteFinanciamientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoFuenteFinanciamientoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoFuenteFinanciamientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoFuenteFinanciamientoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoFuenteFinanciamientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoFuenteFinanciamientoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoFuenteFinanciamientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoFuenteFinanciamientoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoMovimientoApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoMovimientoApropiacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoMovimientoApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoMovimientoApropiacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoMovimientoApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoMovimientoApropiacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoMovimientoApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoMovimientoApropiacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoMovimientoApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoMovimientoApropiacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoMovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoMovimientoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoMovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoMovimientoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoMovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoMovimientoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoMovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoMovimientoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoMovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoMovimientoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoOrdenPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoOrdenPagoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoOrdenPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoOrdenPagoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoOrdenPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoOrdenPagoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoOrdenPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoOrdenPagoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoOrdenPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoOrdenPagoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoRecursoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoRecursoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoRecursoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoRecursoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoRecursoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoRecursoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoRecursoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoRecursoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoRecursoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoRecursoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoTransaccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoTransaccionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoTransaccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoTransaccionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoTransaccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoTransaccionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoTransaccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoTransaccionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoTransaccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TipoTransaccionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TituloInversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TituloInversionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TituloInversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TituloInversionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TituloInversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TituloInversionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TituloInversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TituloInversionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TituloInversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TituloInversionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TrConceptoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TrConceptoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TrConceptoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TrConceptoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TrCuentasContablesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TrCuentasContablesController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TrCuentasContablesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TrCuentasContablesController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:UnidadEjecutoraController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:UnidadEjecutoraController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:UnidadEjecutoraController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:UnidadEjecutoraController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:UnidadEjecutoraController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:UnidadEjecutoraController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:UnidadEjecutoraController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:UnidadEjecutoraController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:UnidadEjecutoraController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:UnidadEjecutoraController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:VersionTipoTransaccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:VersionTipoTransaccionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:VersionTipoTransaccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:VersionTipoTransaccionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:VersionTipoTransaccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:VersionTipoTransaccionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:VersionTipoTransaccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:VersionTipoTransaccionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:VersionTipoTransaccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:VersionTipoTransaccionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

}
