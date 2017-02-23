SET search_path TO financiera;

\echo 'Insert Tipo Ordenes de Pago\n'

insert into tipo_orden_pago (nombre, estado_activo, descripcion)
    values
    ('Factura',                     TRUE, 'Documento de tipo OP'),
    ('Cuenta de Cobro',             TRUE, 'Documento de tipo OP'),
    ('Resolución',                  TRUE, 'Documento de tipo OP'),
    ('Sentencia Juridica',          TRUE, 'Documento de tipo OP'),
    ('Pago Invitación',             TRUE, 'Documento de tipo OP'),
    ('Poliza',                      TRUE, 'Documento de tipo OP');

\echo 'Fin Insert Tipo Ordenes de Pago\n'
