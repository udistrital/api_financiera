SET search_path TO financiera;

-- Estados
INSERT INTO estado_orden_pago (nombre, activo, descripcion, codigo_abreviacion, numero_orden) VALUES ('Elaborado', true, 'Primer estado de la orden de pago', 'EOP_01', 1.00);
INSERT INTO estado_orden_pago (nombre, activo, descripcion, codigo_abreviacion, numero_orden) VALUES ('Aprobacion Contable', true, 'Funcionario de Contabilidad Aproeba la Elaboracion de la OP', 'EOP_02', 2.00);
INSERT INTO estado_orden_pago (nombre, activo, descripcion, codigo_abreviacion, numero_orden) VALUES ('Desaprobacion Contable', true, 'Funcionario de Contabilidad Desaprueba la Elaboracion de la OP', 'EOP_03', 3.00);
INSERT INTO estado_orden_pago (nombre, activo, descripcion, codigo_abreviacion, numero_orden) VALUES ('Aprobacion Presupuestal', true, 'Funcionario de Presupuesto Aproeba la Elaboracion de la OP', 'EOP_04', 4.00);
INSERT INTO estado_orden_pago (nombre, activo, descripcion, codigo_abreviacion, numero_orden) VALUES ('Desaprobacion Presupuestal', true, 'Funcionario de Presupuesto Desaprueba la Elaboracion de la OP', 'EOP_05', 5.00);
INSERT INTO estado_orden_pago (nombre, activo, descripcion, codigo_abreviacion, numero_orden) VALUES ('Enviada', true, 'Funcionario de Presupuesto Envia la OP a Tesoreria', 'EOP_06', 6.00);
INSERT INTO estado_orden_pago (nombre, activo, descripcion, codigo_abreviacion, numero_orden) VALUES ('Radicaca', true, 'Funcionario de Tesoreria Radica la OP', 'EOP_07', 7.00);

-- Tipo
alter sequence financiera.tipo_orden_pago_id_seq restart 1;
INSERT INTO financiera.tipo_orden_pago (nombre, activo, descripcion, codigo_abreviacion, numero_orden) VALUES ('Proveedor', true, 'Orden de pago para proveedores o contratistas', 'OP-PROV', 1);
INSERT INTO financiera.tipo_orden_pago (nombre, activo, descripcion, codigo_abreviacion, numero_orden) VALUES ('Planta', true, 'Orden de pago para nominas de planta', 'OP-PLAN', 2);

INSERT INTO financiera.sub_tipo_orden_pago (nombre, activo, descripcion, codigo_abreviacion, numero_orden, tipo_orden_pago) VALUES ('Cuenta de Cobro', true, 'Orden de pago para Cuenta de Cobro', 'OP-PROV-CC', 1.2, 1);
INSERT INTO financiera.sub_tipo_orden_pago (nombre, activo, descripcion, codigo_abreviacion, numero_orden, tipo_orden_pago) VALUES ('Factura', true, 'Orden de pago tipo Factura', 'OP-PROV-FACT', 1.1, 1);
INSERT INTO financiera.sub_tipo_orden_pago (nombre, activo, descripcion, codigo_abreviacion, numero_orden, tipo_orden_pago) VALUES ('Sentencia Juridica', true, 'Orden de pago para Sentencia Juridica', 'OP-PROV-SJ', 1.3, 1);
INSERT INTO financiera.sub_tipo_orden_pago (nombre, activo, descripcion, codigo_abreviacion, numero_orden, tipo_orden_pago) VALUES ('Pago Invitación', true, 'Orden de pago para Pago Invitación', 'OP-PROV-PI', 1.4, 1);
INSERT INTO financiera.sub_tipo_orden_pago (nombre, activo, descripcion, codigo_abreviacion, numero_orden, tipo_orden_pago) VALUES ('Poliza', true, 'Orden de pago para Poliza', 'OP-PROV-POLI', 1.5, 1);
INSERT INTO financiera.sub_tipo_orden_pago (nombre, activo, descripcion, codigo_abreviacion, numero_orden, tipo_orden_pago) VALUES ('Seguridad Social', true, 'Orden de pago para Seguridad Social', 'OP-PROV-SS', 1.6, 1);
INSERT INTO financiera.sub_tipo_orden_pago (nombre, activo, descripcion, codigo_abreviacion, numero_orden, tipo_orden_pago) VALUES ('Administrativa', true, 'Orden de pago para nominas de planta administrativa', 'OP-PLAN-ADMI', 2.1, 2);
INSERT INTO financiera.sub_tipo_orden_pago (nombre, activo, descripcion, codigo_abreviacion, numero_orden, tipo_orden_pago) VALUES ('Docente', true, 'Orden de pago para nominas de nomina de Docentes', 'OP-PLAN-ADMI', 2.2, 2);

-- Tipo documento afectante  // falta estandar
INSERT INTO tipo_documento_afectante (nombre, descripcion) VALUES ('Orden de Pago', NULL);

-- Homologacion
INSERT INTO homologacion_concepto (vigencia, fecha_creacion, concepto_kronos, concepto_titan) VALUES (2017, '2017-08-29', 73, 239);
INSERT INTO homologacion_concepto (vigencia, fecha_creacion, concepto_kronos, concepto_titan) VALUES (2017, '2017-08-29', 74, 11);
INSERT INTO homologacion_concepto (vigencia, fecha_creacion, concepto_kronos, concepto_titan) VALUES (2017, '2017-08-29', 75, 212);
INSERT INTO homologacion_concepto (vigencia, fecha_creacion, concepto_kronos, concepto_titan) VALUES (2017, '2017-08-29', 75, 231);
INSERT INTO homologacion_concepto (vigencia, fecha_creacion, concepto_kronos, concepto_titan) VALUES (2017, '2017-08-29', 76, 291);
INSERT INTO homologacion_concepto (vigencia, fecha_creacion, concepto_kronos, concepto_titan) VALUES (2017, '2017-08-29', 77, 1213);
INSERT INTO homologacion_concepto (vigencia, fecha_creacion, concepto_kronos, concepto_titan) VALUES (2017, '2017-08-29', 76, 269);
