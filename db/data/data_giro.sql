-- giro_detalle
delete from financiera.giro_detalle;
alter sequence financiera.giro_detalle_id_seq restart 1;
-- giro_estado_giro
delete from financiera.giro_estado_giro;
alter sequence financiera.giro_estado_giro_id_seq restart 1;
-- giro
delete from financiera.giro;
alter sequence financiera.giro_id_seq restart 1;
-- estado_giro
delete from financiera.estado_giro;
alter sequence financiera.estado_giro_id_seq restart 1;
INSERT INTO financiera.estado_giro (nombre, activo, descripcion, codigo_abreviacion, numero_orden) VALUES ('Elaborado', true, 'Primer estado de la orden de pago', 'EGI_01', 1.00);
INSERT INTO financiera.estado_giro (nombre, activo, descripcion, codigo_abreviacion, numero_orden) VALUES ('Aprobacion Contable', true, 'Funcionario de Contabilidad Aproeba la Elaboracion de la OP', 'EGI_02', 2.00);
