// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/udistrital/api_financiera/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/etapa_avance",
			beego.NSInclude(
				&controllers.EtapaAvanceController{},
			),
		),

		beego.NSNamespace("/tipo_movimiento",
			beego.NSInclude(
				&controllers.TipoMovimientoController{},
			),
		),
		beego.NSNamespace("/movimiento_fuente_financiamiento_apropiacion",
			beego.NSInclude(
				&controllers.MovimientoFuenteFinanciamientoApropiacionController{},
			),
		),

		beego.NSNamespace("/fuente_financiamiento",
			beego.NSInclude(
				&controllers.FuenteFinanciamientoController{},
			),
		),

		beego.NSNamespace("/tipo_fuente_financiamiento",
			beego.NSInclude(
				&controllers.TipoFuenteFinanciamientoController{},
			),
		),

		beego.NSNamespace("/fuente_financiamiento_apropiacion",
			beego.NSInclude(
				&controllers.FuenteFinanciamientoApropiacionController{},
			),
		),

		beego.NSNamespace("/requisito_avance",
			beego.NSInclude(
				&controllers.RequisitoAvanceController{},
			),
		),

		beego.NSNamespace("/presupuesto_avance",
			beego.NSInclude(
				&controllers.PresupuestoAvanceController{},
			),
		),

		beego.NSNamespace("/avance_estado_avance",
			beego.NSInclude(
				&controllers.AvanceEstadoAvanceController{},
			),
		),

		beego.NSNamespace("/estado_avance",
			beego.NSInclude(
				&controllers.EstadoAvanceController{},
			),
		),

		beego.NSNamespace("/solicitud_avance",
			beego.NSInclude(
				&controllers.SolicitudAvanceController{},
			),
		),

		beego.NSNamespace("/tipo_avance",
			beego.NSInclude(
				&controllers.TipoAvanceController{},
			),
		),

		beego.NSNamespace("/solicitud_tipo_avance",
			beego.NSInclude(
				&controllers.SolicitudTipoAvanceController{},
			),
		),

		beego.NSNamespace("/requisito_tipo_avance",
			beego.NSInclude(
				&controllers.RequisitoTipoAvanceController{},
			),
		),

		beego.NSNamespace("/solicitud_requisito_tipo_avance",
			beego.NSInclude(
				&controllers.SolicitudRequisitoTipoAvanceController{},
			),
		),

		beego.NSNamespace("/anulacion_registro_presupuestal",
			beego.NSInclude(
				&controllers.AnulacionRegistroPresupuestalController{},
			),
		),

		beego.NSNamespace("/anulacion_registro_presupuestal_disponibilidad_apropiacion",
			beego.NSInclude(
				&controllers.AnulacionRegistroPresupuestalDisponibilidadApropiacionController{},
			),
		),

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

		beego.NSNamespace("/detalle_pac",
			beego.NSInclude(
				&controllers.DetallePacController{},
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

		beego.NSNamespace("/arbol_plan_cuentas",
			beego.NSInclude(
				&controllers.ArbolPlanCuentasController{},
			),
		),

		beego.NSNamespace("/tr_cuentas_contables",
			beego.NSInclude(
				&controllers.TrCuentasContablesController{},
			),
		),

		beego.NSNamespace("/tipo_cuenta_especial",
			beego.NSInclude(
				&controllers.TipoCuentaEspecialController{},
			),
		),

		beego.NSNamespace("/cuenta_especial",
			beego.NSInclude(
				&controllers.CuentaEspecialController{},
			),
		),

		beego.NSNamespace("/homologacion_concepto",
			beego.NSInclude(
				&controllers.HomologacionConceptoController{},
			),
		),

		beego.NSNamespace("/estado_calendario_tributario",
			beego.NSInclude(
				&controllers.EstadoCalendarioTributarioController{},
			),
		),

		beego.NSNamespace("/calendario_tributario",
			beego.NSInclude(
				&controllers.CalendarioTributarioController{},
			),
		),

		beego.NSNamespace("/observacion_calendario",
			beego.NSInclude(
				&controllers.ObservacionCalendarioController{},
			),
		),
		beego.NSNamespace("/cierre_periodo_presupuestal",
			beego.NSInclude(
				&controllers.CierrePeriodoPresupuestalController{},
			),
		),

		beego.NSNamespace("/orden_pago_estado_orden_pago",
			beego.NSInclude(
				&controllers.OrdenPagoEstadoOrdenPagoController{},
			),
		),

		beego.NSNamespace("/giro",
			beego.NSInclude(
				&controllers.GiroController{},
			),
		),

		beego.NSNamespace("/giro_estado_giro",
			beego.NSInclude(
				&controllers.GiroEstadoGiroController{},
			),
		),

		beego.NSNamespace("/estado_giro",
			beego.NSInclude(
				&controllers.EstadoGiroController{},
			),
		),

		beego.NSNamespace("/giro_detalle",
			beego.NSInclude(
				&controllers.GiroDetalleController{},
			),
		),

		beego.NSNamespace("/forma_pago",
			beego.NSInclude(
				&controllers.FormaPagoController{},
			),
		),

		beego.NSNamespace("/estado_movimiento_contable",
			beego.NSInclude(
				&controllers.EstadoMovimientoContableController{},
			),
		),

		beego.NSNamespace("/sub_tipo_orden_pago",
			beego.NSInclude(
				&controllers.SubTipoOrdenPagoController{},
			),
		),

		beego.NSNamespace("/concepto_tesoral_facultad_proyecto",
			beego.NSInclude(
				&controllers.ConceptoTesoralFacultadProyectoController{},
			),
		),

		beego.NSNamespace("/saldo_cuenta_contable",
			beego.NSInclude(
				&controllers.SaldoCuentaContableController{},
			),
		),

		beego.NSNamespace("/homologacion_descuento",
			beego.NSInclude(
				&controllers.HomologacionDescuentoController{},
			),
		),
		beego.NSNamespace("/tipo_disponibilidad",
			beego.NSInclude(
				&controllers.TipoDisponibilidadController{},
			),
		),
		beego.NSNamespace("/disponibilidad_proceso_externo",
			beego.NSInclude(
				&controllers.DisponibilidadProcesoExternoController{},
			),
		),
		beego.NSNamespace("/estado_movimiento_apropiacion",
			beego.NSInclude(
				&controllers.EstadoMovimientoApropiacionController{},
			),
		),

		beego.NSNamespace("/movimiento_apropiacion",
			beego.NSInclude(
				&controllers.MovimientoApropiacionController{},
			),
		),

		beego.NSNamespace("/movimiento_apropiacion_disponibilidad_apropiacion",
			beego.NSInclude(
				&controllers.MovimientoApropiacionDisponibilidadApropiacionController{},
			),
		),

		beego.NSNamespace("/tipo_movimiento_apropiacion",
			beego.NSInclude(
				&controllers.TipoMovimientoApropiacionController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
