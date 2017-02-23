SET search_path TO financiera;

\echo 'Insert Estados de las Ordenes de Pago\n'

insert into estado_orden_pago (nombre, estado_activo, descripcion)
    values
    ('Elaborado',                  TRUE, 'Primer estado de la orden de pago'),
    ('Aprobacion Contable',        TRUE, 'Funcionario de Contabilidad Aproeba la Elaboracion de la OP'),
    ('Desaprobacion Contable',     TRUE, 'Funcionario de Contabilidad Desaprueba la Elaboracion de la OP'),
    ('Aprobacion Presupuestal',    TRUE, 'Funcionario de Presupuesto Aproeba la Elaboracion de la OP'),
    ('Desaprobacion Presupuestal', TRUE, 'Funcionario de Presupuesto Desaprueba la Elaboracion de la OP'),
    ('Enviada',                    TRUE, 'Funcionario de Presupuesto Envia la OP a Tesoreria'),
    ('Radicaca',                   TRUE, 'Funcionario de Tesoreria Radica la OP');

\echo 'Fin Insert Estados de las Ordenes de Pago\n'
