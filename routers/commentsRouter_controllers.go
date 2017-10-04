package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

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

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CategoriaIvaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CategoriaIvaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CategoriaIvaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CategoriaIvaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CategoriaIvaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CategoriaIvaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CategoriaIvaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CategoriaIvaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CategoriaIvaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:CategoriaIvaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
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
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
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
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:IngresoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:IngresoController"],
		beego.ControllerComments{
			Method: "AprobarIngreso",
			Router: `/AprobarIngreso`,
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
			Method: "RechazarIngreso",
			Router: `/RechazarIngreso`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:IvaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:IvaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:IvaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:IvaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:IvaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:IvaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:IvaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:IvaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:IvaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:IvaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
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
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
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
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
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
			Method: "RegistrarOpNomina",
			Router: `RegistrarOpNomina`,
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
			Method: "RegistrarOpSeguridadSocial",
			Router: `RegistrarOpSeguridadSocial`,
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

	beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TrCuentasContablesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/api_financiera/controllers:TrCuentasContablesController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
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
