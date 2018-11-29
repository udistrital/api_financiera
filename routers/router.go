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

		beego.NSNamespace("/tipo_avance_legalizacion",
			beego.NSInclude(
				&controllers.TipoAvanceLegalizacionController{},
			),
		),

		beego.NSNamespace("/avance_legalizacion",
			beego.NSInclude(
				&controllers.AvanceLegalizacionController{},
			),
		),

		beego.NSNamespace("/avance_legalizacion_cuenta_especial",
			beego.NSInclude(
				&controllers.AvanceLegalizacionCuentaEspecialController{},
			),
		),
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
		beego.NSNamespace("/ingreso_estado_ingreso",
			beego.NSInclude(
				&controllers.IngresoEstadoIngresoController{},
			),
		),
		beego.NSNamespace("/tipo_anulacion_presupuestal",
			beego.NSInclude(
				&controllers.TipoAnulacionPresupuestalController{},
			),
		),

		beego.NSNamespace("/estado_comprobante",
			beego.NSInclude(
				&controllers.EstadoComprobanteController{},
			),
		),

		beego.NSNamespace("/registro_comprobantes",
			beego.NSInclude(
				&controllers.RegistroComprobantesController{},
			),
		),

		beego.NSNamespace("/comprobante",
			beego.NSInclude(
				&controllers.ComprobanteController{},
			),
		),

		beego.NSNamespace("/tipo_comprobante",
			beego.NSInclude(
				&controllers.TipoComprobanteController{},
			),
		),

		beego.NSNamespace("/homologacion_comprobantes",
			beego.NSInclude(
				&controllers.HomologacionComprobantesController{},
			),
		),
		beego.NSNamespace("/documento_generador",
			beego.NSInclude(
				&controllers.DocumentoGeneradorController{},
			),
		),
		beego.NSNamespace("/pasivos_fenecidos",
			beego.NSInclude(
				&controllers.Pasivos_fenecidosController{},
			),
		),
		beego.NSNamespace("/producto",
			beego.NSInclude(
				&controllers.ProductoController{},
			),
		),

		beego.NSNamespace("/producto_rubro",
			beego.NSInclude(
				&controllers.ProductoRubroController{},
			),
		),
		beego.NSNamespace("/estado_inversion",
			beego.NSInclude(
				&controllers.EstadoInversionController{},
			),
		),
		beego.NSNamespace("/acta_inversion",
			beego.NSInclude(
				&controllers.ActaInversionController{},
			),
		),
		beego.NSNamespace("/inversion",
			beego.NSInclude(
				&controllers.InversionController{},
			),
		),
		beego.NSNamespace("/inversiones_acta_inversion",
			beego.NSInclude(
				&controllers.InversionesActaInversionController{},
			),
		),
		beego.NSNamespace("/inversion_estado_inversion",
			beego.NSInclude(
				&controllers.InversionEstadoInversionController{},
			),
		),
		beego.NSNamespace("/titulo_inversion",
			beego.NSInclude(
				&controllers.TituloInversionController{},
			),
		),
		beego.NSNamespace("/inversion_concepto",
			beego.NSInclude(
				&controllers.InversionConceptoController{},
			),
		),
		beego.NSNamespace("/historico_inversion",
			beego.NSInclude(
				&controllers.HistoricoInversionController{},
			),
		),
		beego.NSNamespace("/orden_devolucion",
			beego.NSInclude(
				&controllers.OrdenDevolucionController{},
			),
		),
		beego.NSNamespace("/documento_devolucion",
			beego.NSInclude(
				&controllers.DocumentoDevolucionController{},
			),
		),
		beego.NSNamespace("/cuenta_devolucion",
			beego.NSInclude(
				&controllers.CuentaBancariaEnteController{},
			),
		),

		beego.NSNamespace("/acta_devolucion",
			beego.NSInclude(
				&controllers.ActaDevolucionController{},
			),
		),
		beego.NSNamespace("/solicitud_devolucion",
			beego.NSInclude(
				&controllers.SolicitudDevolucionController{},
			),
		),
		beego.NSNamespace("/devolucion_estado_devolucion",
			beego.NSInclude(
				&controllers.DevolucionEstadoDevolucionController{},
			),
		),
		beego.NSNamespace("/razon_devolucion",
			beego.NSInclude(
				&controllers.RazonDevolucionController{},
			),
		),
		beego.NSNamespace("/tipo_devolucion",
			beego.NSInclude(
				&controllers.TipoDevolucionController{},
			),
		),
		beego.NSNamespace("/estado_devolucion",
			beego.NSInclude(
				&controllers.EstadoDevolucionController{},
			),
		),
		beego.NSNamespace("/solicitud_devolucion_estado_devolucion",
			beego.NSInclude(
				&controllers.SolicitudDevolucionEstadoDevolucionController{},
			),
		),
		beego.NSNamespace("/solicitud_devolucion_concepto",
			beego.NSInclude(
				&controllers.SolicitudDevolucionConceptoController{},
			),
		),
		beego.NSNamespace("/orden_devolucion_solicitud_devolucion",
			beego.NSInclude(
				&controllers.OrdenDevolucionSolicitudDevolucionController{},
			),
		),
		beego.NSNamespace("/orden_devolucion_estado_devolucion",
			beego.NSInclude(
				&controllers.OrdenDevolucionEstadoDevolucionController{},
			),
		),
		beego.NSNamespace("/devolucion_tributaria_movimiento",
			beego.NSInclude(
				&controllers.DevolucionTributariaMovimientoAsociadoController{},
			),
		),

		beego.NSNamespace("/devolucion_tributaria_estado_devolucion",
			beego.NSInclude(
				&controllers.DevolucionTributariaEstadoDevolucionController{},
			),
		),

		beego.NSNamespace("/devolucion_tributaria_concepto",
			beego.NSInclude(
				&controllers.DevolucionTributariaConceptoController{},
			),
		),

		beego.NSNamespace("/devolucion_tributaria",
			beego.NSInclude(
				&controllers.DevolucionTributariaController{},
			),
		),
		beego.NSNamespace("/ingreso_sin_situacion_fondos",
			beego.NSInclude(
				&controllers.IngresoSinSituacionFondosController{},
			),
		),

		beego.NSNamespace("/estado_ingreso_sin_situacion_fondos",
			beego.NSInclude(
				&controllers.EstadoIngresoSinSituacionFondosController{},
			),
		),

		beego.NSNamespace("/ingreso_sin_situacion_fondos_estado",
			beego.NSInclude(
				&controllers.IngresoSinSituacionFondosEstadoController{},
			),
		),
		beego.NSNamespace("/orden_pago_registro_presupuestal",
			beego.NSInclude(
				&controllers.OrdenPagoRegistroPresupuestalController{},
			),
		),

		beego.NSNamespace("/orden_pago_cuenta_especial",
			beego.NSInclude(
				&controllers.OrdenPagoCuentaEspecialController{},
			),
		),

		beego.NSNamespace("/informacion_adicional_banco",
			beego.NSInclude(
				&controllers.InformacionAdicionalBancoController{},
			),
		),
		beego.NSNamespace("/rubro_homologado_rubro",
			beego.NSInclude(
				&controllers.RubroHomologadoRubroController{},
			),
		),
		beego.NSNamespace("/cancelacion_inversion_inversion",
			beego.NSInclude(
				&controllers.CancelacionInversionInversionController{},
			),
		),

		beego.NSNamespace("/estado_cancelacion_inversion",
			beego.NSInclude(
				&controllers.EstadoCancelacionInversionController{},
			),
		),

		beego.NSNamespace("/cancelacion_inversion_estado_cancelacion",
			beego.NSInclude(
				&controllers.CancelacionInversionEstadoCancelacionController{},
			),
		),

		beego.NSNamespace("/cancelacion_inversion",
			beego.NSInclude(
				&controllers.CancelacionInversionController{},
			),
		),
		beego.NSNamespace("/cancelacion_inversion_concepto",
			beego.NSInclude(
				&controllers.CancelacionInversionConceptoController{},
			),
		),
		beego.NSNamespace("/causal_reintegro",
			beego.NSInclude(
				&controllers.CausalReintegroController{},
			),
		),

		beego.NSNamespace("/reintegro",
			beego.NSInclude(
				&controllers.ReintegroController{},
			),
		),

		beego.NSNamespace("/reintegro_avance_legalizacion",
			beego.NSInclude(
				&controllers.ReintegroAvanceLegalizacionController{},
			),
		),

		beego.NSNamespace("/orden_pago_avance_legalizacion",
			beego.NSInclude(
				&controllers.OrdenPagoAvanceLegalizacionController{},
			),
		),

		beego.NSNamespace("/chequera",
			beego.NSInclude(
				&controllers.ChequeraController{},
			),
		),

		beego.NSNamespace("/estado_chequera",
			beego.NSInclude(
				&controllers.EstadoChequeraController{},
			),
		),

		beego.NSNamespace("/chequera_estado_chequera",
			beego.NSInclude(
				&controllers.ChequeraEstadoChequeraController{},
			),
		),
		beego.NSNamespace("/cheque",
			beego.NSInclude(
				&controllers.ChequeController{},
			),
		),

		beego.NSNamespace("/cheque_estado_cheque",
			beego.NSInclude(
				&controllers.ChequeEstadoChequeController{},
			),
		),

		beego.NSNamespace("/estado_cheque",
			beego.NSInclude(
				&controllers.EstadoChequeController{},
			),
		),
		beego.NSNamespace("/avance_legalizacion_tipo",
			beego.NSInclude(
				&controllers.AvanceLegalizacionTipoController{},
			),
		),

		beego.NSNamespace("/concepto_avance_legalizacion_tipo",
			beego.NSInclude(
				&controllers.ConceptoAvanceLegalizacionTipoController{},
			),
		),

		beego.NSNamespace("/estado_legalizacion",
			beego.NSInclude(
				&controllers.EstadoLegalizacionController{},
			),
		),

		beego.NSNamespace("/estado_legalizacion_avance_legalizacion",
			beego.NSInclude(
				&controllers.EstadoLegalizacionAvanceLegalizacionController{},
			),
		),
		beego.NSNamespace("/avance_legalizacion_sub_tipo",
			beego.NSInclude(
				&controllers.AvanceLegalizacionSubTipoController{},
			),
		),
		beego.NSNamespace("/estado_avance_legalizacion_tipo",
			beego.NSInclude(
				&controllers.EstadoAvanceLegalizacionTipoController{},
			),
		),
		beego.NSNamespace("/tipo_transaccion_version",
			beego.NSInclude(
				&controllers.TipoTransaccionVersionController{},
			),
		),

		beego.NSNamespace("/detalle_tipo_transaccion_version",
			beego.NSInclude(
				&controllers.DetalleTipoTransaccionVersionController{},
			),
		),

		beego.NSNamespace("/version_tipo_transaccion",
			beego.NSInclude(
				&controllers.VersionTipoTransaccionController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
