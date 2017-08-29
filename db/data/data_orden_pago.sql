SET search_path TO financiera;

-- Estados
INSERT INTO estado_orden_pago (nombre, activo, descripcion, codigo_abreviacion, numero_orden) VALUES ('Elaborado', true, 'Primer estado de la orden de pago', 'EOP_01', 1.00);
INSERT INTO estado_orden_pago (nombre, activo, descripcion, codigo_abreviacion, numero_orden) VALUES ('Aprobacion Contable', true, 'Funcionario de Contabilidad Aproeba la Elaboracion de la OP', 'EOP_02', 2.00);
INSERT INTO estado_orden_pago (nombre, activo, descripcion, codigo_abreviacion, numero_orden) VALUES ('Desaprobacion Contable', true, 'Funcionario de Contabilidad Desaprueba la Elaboracion de la OP', 'EOP_03', 3.00);
INSERT INTO estado_orden_pago (nombre, activo, descripcion, codigo_abreviacion, numero_orden) VALUES ('Aprobacion Presupuestal', true, 'Funcionario de Presupuesto Aproeba la Elaboracion de la OP', 'EOP_04', 4.00);
INSERT INTO estado_orden_pago (nombre, activo, descripcion, codigo_abreviacion, numero_orden) VALUES ('Desaprobacion Presupuestal', true, 'Funcionario de Presupuesto Desaprueba la Elaboracion de la OP', 'EOP_05', 5.00);
INSERT INTO estado_orden_pago (nombre, activo, descripcion, codigo_abreviacion, numero_orden) VALUES ('Enviada', true, 'Funcionario de Presupuesto Envia la OP a Tesoreria', 'EOP_06', 6.00);
INSERT INTO estado_orden_pago (nombre, activo, descripcion, codigo_abreviacion, numero_orden) VALUES ('Radicaca', true, 'Funcionario de Tesoreria Radica la OP', 'EOP_07', 7.00);

-- Tipo // falata estanar
INSERT INTO tipo_orden_pago (nombre, estado_activo, descripcion) VALUES ('Factura', true, 'Documento de tipo OP');
INSERT INTO tipo_orden_pago (nombre, estado_activo, descripcion) VALUES ('Cuenta de Cobro', true, 'Documento de tipo OP');
INSERT INTO tipo_orden_pago (nombre, estado_activo, descripcion) VALUES ('Resolución', true, 'Documento de tipo OP');
INSERT INTO tipo_orden_pago (nombre, estado_activo, descripcion) VALUES ('Sentencia Juridica', true, 'Documento de tipo OP');
INSERT INTO tipo_orden_pago (nombre, estado_activo, descripcion) VALUES ('Pago Invitación', true, 'Documento de tipo OP');
INSERT INTO tipo_orden_pago (nombre, estado_activo, descripcion) VALUES ('Poliza', true, 'Documento de tipo OP');

-- Tipo documento afectante  // falta estandar
INSERT INTO tipo_documento_afectante (nombre, descripcion) VALUES ('Orden de Pago', NULL);

