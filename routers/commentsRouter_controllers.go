package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["api_financiera/controllers:AfectacionConceptoController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:AfectacionConceptoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:AfectacionConceptoController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:AfectacionConceptoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:AfectacionConceptoController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:AfectacionConceptoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:AfectacionConceptoController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:AfectacionConceptoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:AfectacionConceptoController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:AfectacionConceptoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:AnulacionDisponibilidadApropiacionController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:AnulacionDisponibilidadApropiacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:AnulacionDisponibilidadApropiacionController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:AnulacionDisponibilidadApropiacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:AnulacionDisponibilidadApropiacionController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:AnulacionDisponibilidadApropiacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:AnulacionDisponibilidadApropiacionController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:AnulacionDisponibilidadApropiacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:AnulacionDisponibilidadApropiacionController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:AnulacionDisponibilidadApropiacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:AnulacionDisponibilidadController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:AnulacionDisponibilidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:AnulacionDisponibilidadController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:AnulacionDisponibilidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:AnulacionDisponibilidadController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:AnulacionDisponibilidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:AnulacionDisponibilidadController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:AnulacionDisponibilidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:AnulacionDisponibilidadController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:AnulacionDisponibilidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:AnulacionReservaController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:AnulacionReservaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:AnulacionReservaController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:AnulacionReservaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:AnulacionReservaController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:AnulacionReservaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:AnulacionReservaController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:AnulacionReservaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:AnulacionReservaController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:AnulacionReservaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:ApropiacionController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:ApropiacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:ApropiacionController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:ApropiacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:ApropiacionController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:ApropiacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:ApropiacionController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:ApropiacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:ApropiacionController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:ApropiacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:ArbolConceptosController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:ArbolConceptosController"],
		beego.ControllerComments{
			Method: "MakeTree",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:CategoriaCompromisoController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:CategoriaCompromisoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:CategoriaCompromisoController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:CategoriaCompromisoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:CategoriaCompromisoController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:CategoriaCompromisoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:CategoriaCompromisoController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:CategoriaCompromisoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:CategoriaCompromisoController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:CategoriaCompromisoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:ClaseTransaccionController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:ClaseTransaccionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:ClaseTransaccionController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:ClaseTransaccionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:ClaseTransaccionController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:ClaseTransaccionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:ClaseTransaccionController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:ClaseTransaccionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:ClaseTransaccionController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:ClaseTransaccionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:CompromisoController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:CompromisoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:CompromisoController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:CompromisoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:CompromisoController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:CompromisoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:CompromisoController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:CompromisoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:CompromisoController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:CompromisoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:ConceptoConceptoController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:ConceptoConceptoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:ConceptoConceptoController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:ConceptoConceptoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:ConceptoConceptoController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:ConceptoConceptoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:ConceptoConceptoController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:ConceptoConceptoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:ConceptoConceptoController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:ConceptoConceptoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:ConceptoController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:ConceptoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:ConceptoController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:ConceptoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:ConceptoController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:ConceptoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:ConceptoController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:ConceptoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:ConceptoController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:ConceptoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:ConceptoCuentaContableController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:ConceptoCuentaContableController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:ConceptoCuentaContableController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:ConceptoCuentaContableController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:ConceptoCuentaContableController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:ConceptoCuentaContableController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:ConceptoCuentaContableController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:ConceptoCuentaContableController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:ConceptoCuentaContableController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:ConceptoCuentaContableController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:ConceptoOrdenPagoController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:ConceptoOrdenPagoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:ConceptoOrdenPagoController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:ConceptoOrdenPagoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:ConceptoOrdenPagoController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:ConceptoOrdenPagoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:ConceptoOrdenPagoController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:ConceptoOrdenPagoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:ConceptoOrdenPagoController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:ConceptoOrdenPagoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:CuentaBancariaController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:CuentaBancariaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:CuentaBancariaController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:CuentaBancariaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:CuentaBancariaController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:CuentaBancariaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:CuentaBancariaController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:CuentaBancariaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:CuentaBancariaController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:CuentaBancariaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:CuentaContableController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:CuentaContableController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:CuentaContableController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:CuentaContableController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:CuentaContableController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:CuentaContableController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:CuentaContableController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:CuentaContableController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:CuentaContableController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:CuentaContableController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:DisponibilidadApropiacionController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:DisponibilidadApropiacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:DisponibilidadApropiacionController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:DisponibilidadApropiacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:DisponibilidadApropiacionController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:DisponibilidadApropiacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:DisponibilidadApropiacionController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:DisponibilidadApropiacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:DisponibilidadApropiacionController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:DisponibilidadApropiacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:DisponibilidadController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:DisponibilidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:DisponibilidadController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:DisponibilidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:DisponibilidadController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:DisponibilidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:DisponibilidadController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:DisponibilidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:DisponibilidadController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:DisponibilidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:EntidadController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:EntidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:EntidadController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:EntidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:EntidadController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:EntidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:EntidadController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:EntidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:EntidadController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:EntidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:EstadoApropiacionController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:EstadoApropiacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:EstadoApropiacionController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:EstadoApropiacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:EstadoApropiacionController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:EstadoApropiacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:EstadoApropiacionController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:EstadoApropiacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:EstadoApropiacionController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:EstadoApropiacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:EstadoCompromisoController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:EstadoCompromisoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:EstadoCompromisoController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:EstadoCompromisoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:EstadoCompromisoController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:EstadoCompromisoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:EstadoCompromisoController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:EstadoCompromisoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:EstadoCompromisoController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:EstadoCompromisoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:EstadoDisponibilidadController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:EstadoDisponibilidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:EstadoDisponibilidadController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:EstadoDisponibilidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:EstadoDisponibilidadController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:EstadoDisponibilidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:EstadoDisponibilidadController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:EstadoDisponibilidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:EstadoDisponibilidadController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:EstadoDisponibilidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:EstadoIngresoController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:EstadoIngresoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:EstadoIngresoController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:EstadoIngresoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:EstadoIngresoController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:EstadoIngresoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:EstadoIngresoController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:EstadoIngresoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:EstadoIngresoController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:EstadoIngresoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:EstadoOrdenPagoController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:EstadoOrdenPagoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:EstadoOrdenPagoController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:EstadoOrdenPagoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:EstadoOrdenPagoController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:EstadoOrdenPagoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:EstadoOrdenPagoController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:EstadoOrdenPagoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:EstadoOrdenPagoController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:EstadoOrdenPagoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:EstadoRegistroPresupuestalController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:EstadoRegistroPresupuestalController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:EstadoRegistroPresupuestalController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:EstadoRegistroPresupuestalController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:EstadoRegistroPresupuestalController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:EstadoRegistroPresupuestalController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:EstadoRegistroPresupuestalController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:EstadoRegistroPresupuestalController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:EstadoRegistroPresupuestalController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:EstadoRegistroPresupuestalController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:EstadoReservaPresupuestalController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:EstadoReservaPresupuestalController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:EstadoReservaPresupuestalController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:EstadoReservaPresupuestalController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:EstadoReservaPresupuestalController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:EstadoReservaPresupuestalController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:EstadoReservaPresupuestalController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:EstadoReservaPresupuestalController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:EstadoReservaPresupuestalController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:EstadoReservaPresupuestalController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:EstructuraCuentasController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:EstructuraCuentasController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:EstructuraCuentasController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:EstructuraCuentasController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:EstructuraCuentasController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:EstructuraCuentasController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:EstructuraCuentasController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:EstructuraCuentasController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:EstructuraCuentasController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:EstructuraCuentasController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:EstructuraNivelesClasificacionController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:EstructuraNivelesClasificacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:EstructuraNivelesClasificacionController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:EstructuraNivelesClasificacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:EstructuraNivelesClasificacionController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:EstructuraNivelesClasificacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:EstructuraNivelesClasificacionController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:EstructuraNivelesClasificacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:EstructuraNivelesClasificacionController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:EstructuraNivelesClasificacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:FormaIngresoController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:FormaIngresoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:FormaIngresoController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:FormaIngresoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:FormaIngresoController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:FormaIngresoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:FormaIngresoController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:FormaIngresoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:FormaIngresoController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:FormaIngresoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:IngresoConceptoController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:IngresoConceptoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:IngresoConceptoController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:IngresoConceptoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:IngresoConceptoController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:IngresoConceptoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:IngresoConceptoController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:IngresoConceptoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:IngresoConceptoController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:IngresoConceptoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:IngresoController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:IngresoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:IngresoController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:IngresoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:IngresoController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:IngresoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:IngresoController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:IngresoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:IngresoController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:IngresoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:ModificacionPresupuestalController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:ModificacionPresupuestalController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:ModificacionPresupuestalController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:ModificacionPresupuestalController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:ModificacionPresupuestalController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:ModificacionPresupuestalController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:ModificacionPresupuestalController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:ModificacionPresupuestalController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:ModificacionPresupuestalController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:ModificacionPresupuestalController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:ModuloKronosController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:ModuloKronosController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:ModuloKronosController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:ModuloKronosController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:ModuloKronosController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:ModuloKronosController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:ModuloKronosController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:ModuloKronosController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:ModuloKronosController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:ModuloKronosController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:MovimientoContableController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:MovimientoContableController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:MovimientoContableController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:MovimientoContableController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:MovimientoContableController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:MovimientoContableController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:MovimientoContableController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:MovimientoContableController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:MovimientoContableController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:MovimientoContableController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:NivelClasificacionController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:NivelClasificacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:NivelClasificacionController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:NivelClasificacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:NivelClasificacionController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:NivelClasificacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:NivelClasificacionController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:NivelClasificacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:NivelClasificacionController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:NivelClasificacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:OrdenPagoController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:OrdenPagoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:OrdenPagoController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:OrdenPagoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:OrdenPagoController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:OrdenPagoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:OrdenPagoController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:OrdenPagoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:OrdenPagoController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:OrdenPagoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:PeriodoContableController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:PeriodoContableController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:PeriodoContableController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:PeriodoContableController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:PeriodoContableController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:PeriodoContableController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:PeriodoContableController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:PeriodoContableController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:PeriodoContableController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:PeriodoContableController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:PeriodoPlanCuentasController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:PeriodoPlanCuentasController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:PeriodoPlanCuentasController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:PeriodoPlanCuentasController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:PeriodoPlanCuentasController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:PeriodoPlanCuentasController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:PeriodoPlanCuentasController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:PeriodoPlanCuentasController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:PeriodoPlanCuentasController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:PeriodoPlanCuentasController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:PlanCuentasController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:PlanCuentasController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:PlanCuentasController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:PlanCuentasController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:PlanCuentasController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:PlanCuentasController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:PlanCuentasController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:PlanCuentasController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:PlanCuentasController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:PlanCuentasController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:RegistroPresupuestalController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:RegistroPresupuestalController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:RegistroPresupuestalController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:RegistroPresupuestalController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:RegistroPresupuestalController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:RegistroPresupuestalController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:RegistroPresupuestalController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:RegistroPresupuestalController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:RegistroPresupuestalController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:RegistroPresupuestalController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:RegistroPresupuestalDisponibilidadApropiacionController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:RegistroPresupuestalDisponibilidadApropiacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:RegistroPresupuestalDisponibilidadApropiacionController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:RegistroPresupuestalDisponibilidadApropiacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:RegistroPresupuestalDisponibilidadApropiacionController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:RegistroPresupuestalDisponibilidadApropiacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:RegistroPresupuestalDisponibilidadApropiacionController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:RegistroPresupuestalDisponibilidadApropiacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:RegistroPresupuestalDisponibilidadApropiacionController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:RegistroPresupuestalDisponibilidadApropiacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:ReservaPresupuestalController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:ReservaPresupuestalController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:ReservaPresupuestalController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:ReservaPresupuestalController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:ReservaPresupuestalController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:ReservaPresupuestalController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:ReservaPresupuestalController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:ReservaPresupuestalController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:ReservaPresupuestalController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:ReservaPresupuestalController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:RubroController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:RubroController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:RubroController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:RubroController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:RubroController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:RubroController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:RubroController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:RubroController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:RubroController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:RubroController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:RubroHomologadoController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:RubroHomologadoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:RubroHomologadoController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:RubroHomologadoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:RubroHomologadoController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:RubroHomologadoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:RubroHomologadoController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:RubroHomologadoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:RubroHomologadoController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:RubroHomologadoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:RubroRubroController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:RubroRubroController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:RubroRubroController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:RubroRubroController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:RubroRubroController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:RubroRubroController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:RubroRubroController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:RubroRubroController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:RubroRubroController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:RubroRubroController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:TipoAfectacionController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:TipoAfectacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:TipoAfectacionController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:TipoAfectacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:TipoAfectacionController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:TipoAfectacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:TipoAfectacionController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:TipoAfectacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:TipoAfectacionController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:TipoAfectacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:TipoCompromisoTesoralController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:TipoCompromisoTesoralController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:TipoCompromisoTesoralController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:TipoCompromisoTesoralController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:TipoCompromisoTesoralController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:TipoCompromisoTesoralController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:TipoCompromisoTesoralController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:TipoCompromisoTesoralController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:TipoCompromisoTesoralController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:TipoCompromisoTesoralController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:TipoConceptoController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:TipoConceptoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:TipoConceptoController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:TipoConceptoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:TipoConceptoController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:TipoConceptoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:TipoConceptoController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:TipoConceptoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:TipoConceptoController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:TipoConceptoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:TipoCuentaBancariaController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:TipoCuentaBancariaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:TipoCuentaBancariaController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:TipoCuentaBancariaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:TipoCuentaBancariaController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:TipoCuentaBancariaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:TipoCuentaBancariaController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:TipoCuentaBancariaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:TipoCuentaBancariaController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:TipoCuentaBancariaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:TipoDocumentoAfectanteController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:TipoDocumentoAfectanteController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:TipoDocumentoAfectanteController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:TipoDocumentoAfectanteController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:TipoDocumentoAfectanteController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:TipoDocumentoAfectanteController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:TipoDocumentoAfectanteController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:TipoDocumentoAfectanteController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:TipoDocumentoAfectanteController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:TipoDocumentoAfectanteController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:TipoEntidadController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:TipoEntidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:TipoEntidadController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:TipoEntidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:TipoEntidadController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:TipoEntidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:TipoEntidadController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:TipoEntidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:TipoEntidadController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:TipoEntidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:TipoOrdenPagoController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:TipoOrdenPagoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:TipoOrdenPagoController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:TipoOrdenPagoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:TipoOrdenPagoController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:TipoOrdenPagoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:TipoOrdenPagoController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:TipoOrdenPagoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:TipoOrdenPagoController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:TipoOrdenPagoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:TipoRecursoController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:TipoRecursoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:TipoRecursoController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:TipoRecursoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:TipoRecursoController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:TipoRecursoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:TipoRecursoController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:TipoRecursoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:TipoRecursoController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:TipoRecursoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:TipoTransaccionController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:TipoTransaccionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:TipoTransaccionController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:TipoTransaccionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:TipoTransaccionController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:TipoTransaccionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:TipoTransaccionController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:TipoTransaccionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:TipoTransaccionController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:TipoTransaccionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:UnidadEjecutoraController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:UnidadEjecutoraController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:UnidadEjecutoraController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:UnidadEjecutoraController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:UnidadEjecutoraController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:UnidadEjecutoraController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:UnidadEjecutoraController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:UnidadEjecutoraController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:UnidadEjecutoraController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:UnidadEjecutoraController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:VersionTipoTransaccionController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:VersionTipoTransaccionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:VersionTipoTransaccionController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:VersionTipoTransaccionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:VersionTipoTransaccionController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:VersionTipoTransaccionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:VersionTipoTransaccionController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:VersionTipoTransaccionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["api_financiera/controllers:VersionTipoTransaccionController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:VersionTipoTransaccionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

		beego.GlobalControllerRouter["api_financiera/controllers:CategoriaIvaController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:CategoriaIvaController"],
		  beego.ControllerComments{
		    Method: "Post",
		    Router: `/`,
		    AllowHTTPMethods: []string{"post"},
		    Params: nil})

		beego.GlobalControllerRouter["api_financiera/controllers:CategoriaIvaController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:CategoriaIvaController"],
		  beego.ControllerComments{
		    Method: "GetOne",
		    Router: `/:id`,
		    AllowHTTPMethods: []string{"get"},
		    Params: nil})

		beego.GlobalControllerRouter["api_financiera/controllers:CategoriaIvaController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:CategoriaIvaController"],
		  beego.ControllerComments{
		    Method: "GetAll",
		    Router: `/`,
		    AllowHTTPMethods: []string{"get"},
		    Params: nil})

		beego.GlobalControllerRouter["api_financiera/controllers:CategoriaIvaController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:CategoriaIvaController"],
		  beego.ControllerComments{
		    Method: "Put",
		    Router: `/:id`,
		    AllowHTTPMethods: []string{"put"},
		    Params: nil})

		beego.GlobalControllerRouter["api_financiera/controllers:CategoriaIvaController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:CategoriaIvaController"],
		  beego.ControllerComments{
		    Method: "Delete",
		    Router: `/:id`,
		    AllowHTTPMethods: []string{"delete"},
		    Params: nil})

			beego.GlobalControllerRouter["api_financiera/controllers:IvaController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:IvaController"],
			  beego.ControllerComments{
			    Method: "Post",
			    Router: `/`,
			    AllowHTTPMethods: []string{"post"},
			    Params: nil})

			beego.GlobalControllerRouter["api_financiera/controllers:IvaController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:IvaController"],
			  beego.ControllerComments{
			    Method: "GetOne",
			    Router: `/:id`,
			    AllowHTTPMethods: []string{"get"},
			    Params: nil})

			beego.GlobalControllerRouter["api_financiera/controllers:IvaController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:IvaController"],
			  beego.ControllerComments{
			    Method: "GetAll",
			    Router: `/`,
			    AllowHTTPMethods: []string{"get"},
			    Params: nil})

			beego.GlobalControllerRouter["api_financiera/controllers:IvaController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:IvaController"],
			  beego.ControllerComments{
			    Method: "Put",
			    Router: `/:id`,
			    AllowHTTPMethods: []string{"put"},
			    Params: nil})

			beego.GlobalControllerRouter["api_financiera/controllers:IvaController"] = append(beego.GlobalControllerRouter["api_financiera/controllers:IvaController"],
			  beego.ControllerComments{
			    Method: "Delete",
			    Router: `/:id`,
			    AllowHTTPMethods: []string{"delete"},
			    Params: nil})

}
