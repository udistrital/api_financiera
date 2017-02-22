// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"api_financiera/controllers"

	"github.com/astaxie/beego"

	"github.com/astaxie/beego/plugins/cors"
)

func init() {
	beego.Debug("Filters init...")
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))

	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/anulacion_disponibilidad",
			beego.NSInclude(
				&controllers.AnulacionDisponibilidadController{},
			),
		),

		beego.NSNamespace("/anulacion_disponibilidad_apropiacion",
			beego.NSInclude(
				&controllers.AnulacionDisponibilidadApropiacionController{},
			),
		),

		beego.NSNamespace("/anulacion_reserva",
			beego.NSInclude(
				&controllers.AnulacionReservaController{},
			),
		),

		beego.NSNamespace("/estado_apropiacion",
			beego.NSInclude(
				&controllers.EstadoApropiacionController{},
			),
		),

		beego.NSNamespace("/estado_disponibilidad",
			beego.NSInclude(
				&controllers.EstadoDisponibilidadController{},
			),
		),

		beego.NSNamespace("/disponibilidad",
			beego.NSInclude(
				&controllers.DisponibilidadController{},
			),
		),

		beego.NSNamespace("/apropiacion",
			beego.NSInclude(
				&controllers.ApropiacionController{},
			),
		),

		beego.NSNamespace("/modificacion_presupuestal",
			beego.NSInclude(
				&controllers.ModificacionPresupuestalController{},
			),
		),

		beego.NSNamespace("/estado_registro_presupuestal",
			beego.NSInclude(
				&controllers.EstadoRegistroPresupuestalController{},
			),
		),

		beego.NSNamespace("/disponibilidad_apropiacion",
			beego.NSInclude(
				&controllers.DisponibilidadApropiacionController{},
			),
		),

		beego.NSNamespace("/registro_presupuestal_disponibilidad_apropiacion",
			beego.NSInclude(
				&controllers.RegistroPresupuestalDisponibilidadApropiacionController{},
			),
		),

		beego.NSNamespace("/estado_reserva_presupuestal",
			beego.NSInclude(
				&controllers.EstadoReservaPresupuestalController{},
			),
		),

		beego.NSNamespace("/reserva_presupuestal",
			beego.NSInclude(
				&controllers.ReservaPresupuestalController{},
			),
		),

		beego.NSNamespace("/rubro_homologado",
			beego.NSInclude(
				&controllers.RubroHomologadoController{},
			),
		),

		beego.NSNamespace("/rubro_rubro",
			beego.NSInclude(
				&controllers.RubroRubroController{},
			),
		),

		beego.NSNamespace("/nivel_clasificacion",
			beego.NSInclude(
				&controllers.NivelClasificacionController{},
			),
		),

		beego.NSNamespace("/tipo_documento_afectante",
			beego.NSInclude(
				&controllers.TipoDocumentoAfectanteController{},
			),
		),

		beego.NSNamespace("/movimiento_contable",
			beego.NSInclude(
				&controllers.MovimientoContableController{},
			),
		),

		beego.NSNamespace("/estructura_cuentas",
			beego.NSInclude(
				&controllers.EstructuraCuentasController{},
			),
		),

		beego.NSNamespace("/periodo_contable",
			beego.NSInclude(
				&controllers.PeriodoContableController{},
			),
		),

		beego.NSNamespace("/plan_cuentas",
			beego.NSInclude(
				&controllers.PlanCuentasController{},
			),
		),

		beego.NSNamespace("/periodo_plan_cuentas",
			beego.NSInclude(
				&controllers.PeriodoPlanCuentasController{},
			),
		),

		beego.NSNamespace("/cuenta_contable",
			beego.NSInclude(
				&controllers.CuentaContableController{},
			),
		),

		beego.NSNamespace("/concepto_cuenta_contable",
			beego.NSInclude(
				&controllers.ConceptoCuentaContableController{},
			),
		),

		beego.NSNamespace("/tipo_cuenta_bancaria",
			beego.NSInclude(
				&controllers.TipoCuentaBancariaController{},
			),
		),

		beego.NSNamespace("/tipo_recurso",
			beego.NSInclude(
				&controllers.TipoRecursoController{},
			),
		),

		beego.NSNamespace("/cuenta_bancaria",
			beego.NSInclude(
				&controllers.CuentaBancariaController{},
			),
		),

		beego.NSNamespace("/tipo_entidad",
			beego.NSInclude(
				&controllers.TipoEntidadController{},
			),
		),

		beego.NSNamespace("/entidad",
			beego.NSInclude(
				&controllers.EntidadController{},
			),
		),

		beego.NSNamespace("/categoria_compromiso",
			beego.NSInclude(
				&controllers.CategoriaCompromisoController{},
			),
		),

		beego.NSNamespace("/tipo_compromiso_tesoral",
			beego.NSInclude(
				&controllers.TipoCompromisoTesoralController{},
			),
		),

		beego.NSNamespace("/estado_compromiso",
			beego.NSInclude(
				&controllers.EstadoCompromisoController{},
			),
		),

		beego.NSNamespace("/compromiso",
			beego.NSInclude(
				&controllers.CompromisoController{},
			),
		),

		beego.NSNamespace("/version_tipo_transaccion",
			beego.NSInclude(
				&controllers.VersionTipoTransaccionController{},
			),
		),

		beego.NSNamespace("/tipo_transaccion",
			beego.NSInclude(
				&controllers.TipoTransaccionController{},
			),
		),

		beego.NSNamespace("/modulo_kronos",
			beego.NSInclude(
				&controllers.ModuloKronosController{},
			),
		),

		beego.NSNamespace("/clase_transaccion",
			beego.NSInclude(
				&controllers.ClaseTransaccionController{},
			),
		),

		beego.NSNamespace("/tipo_afectacion",
			beego.NSInclude(
				&controllers.TipoAfectacionController{},
			),
		),

		beego.NSNamespace("/afectacion_concepto",
			beego.NSInclude(
				&controllers.AfectacionConceptoController{},
			),
		),

		beego.NSNamespace("/concepto_concepto",
			beego.NSInclude(
				&controllers.ConceptoConceptoController{},
			),
		),

		beego.NSNamespace("/tipo_concepto",
			beego.NSInclude(
				&controllers.TipoConceptoController{},
			),
		),

		beego.NSNamespace("/rubro",
			beego.NSInclude(
				&controllers.RubroController{},
			),
		),

		beego.NSNamespace("/forma_ingreso",
			beego.NSInclude(
				&controllers.FormaIngresoController{},
			),
		),

		beego.NSNamespace("/estado_ingreso",
			beego.NSInclude(
				&controllers.EstadoIngresoController{},
			),
		),

		beego.NSNamespace("/ingreso",
			beego.NSInclude(
				&controllers.IngresoController{},
			),
		),

		beego.NSNamespace("/ingreso_concepto",
			beego.NSInclude(
				&controllers.IngresoConceptoController{},
			),
		),

		beego.NSNamespace("/unidad_ejecutora",
			beego.NSInclude(
				&controllers.UnidadEjecutoraController{},
			),
		),

		beego.NSNamespace("/estado_orden_pago",
			beego.NSInclude(
				&controllers.EstadoOrdenPagoController{},
			),
		),

		beego.NSNamespace("/tipo_orden_pago",
			beego.NSInclude(
				&controllers.TipoOrdenPagoController{},
			),
		),

		beego.NSNamespace("/registro_presupuestal",
			beego.NSInclude(
				&controllers.RegistroPresupuestalController{},
			),
		),

		beego.NSNamespace("/orden_pago",
			beego.NSInclude(
				&controllers.OrdenPagoController{},
			),
		),

		beego.NSNamespace("/concepto",
			beego.NSInclude(
				&controllers.ConceptoController{},
			),
		),

		beego.NSNamespace("/concepto_orden_pago",
			beego.NSInclude(
				&controllers.ConceptoOrdenPagoController{},
			),
		),

		beego.NSNamespace("/estructura_niveles_clasificacion",
			beego.NSInclude(
				&controllers.EstructuraNivelesClasificacionController{},
			),
		),

		beego.NSNamespace("/arbol_conceptos",
			beego.NSInclude(
				&controllers.ArbolConceptosController{},
			),
		),

		beego.NSNamespace("/tr_concepto",
			beego.NSInclude(
				&controllers.TrConceptoController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
