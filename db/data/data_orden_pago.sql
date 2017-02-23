SET search_path TO financiera;

\echo 'Insert Ordenes de Pago\n'

insert into orden_pago
    (vigencia, fecha_creacion, registro_presupuestal, valor_total, persona_elaboro, tipo_orden_pago, unidad_ejecutora, estado_orden_pago)
      values
    (2016,              now(),                     1,     1190000,               1,              2,               1,                 7),
    (2017,              now(),                     1,     1190000,               1,              2,               1,                 7),
    (2017,              now(),                     1,     1190000,               1,              2,               1,                 7);

\echo 'Fin Insert  Ordenes de Pago\n'
