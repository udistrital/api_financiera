SET search_path TO financiera;

\echo 'Insert tipo_entidad\n'

insert into tipo_entidad (nombre, descripcion)
    values
    ('Reportante',   'descripcion'),
    ('Homologada',   'descripcion');

\echo 'Fin Insert tipo_entidad\n'


\echo 'Insert Entidad\n'

insert into entidad (nombre, codigo_entidad, tipo_entidad)
    values ('Universidad Distrital Francisco José de Caldas', '230', 1);

\echo 'Fin Insert Tipo Ordenes de Pago\n'


\echo 'Insert Unidad Ejecutora\n'

insert into unidad_ejecutora (nombre, descripcion, entidad)
    values
    ('Rector',                    'Unidad ejecutora sede administrativa, abora las dependencias de juridica y compras  que realizan la ejecucion de contratos', 1),
    ('Convenios',                 'Unidad ejecutora idexud, la cual se encarga de la gestion de contratacion y compras a traves y con convenios', 1);

\echo 'Fin Insert Unidad Ejecutora\n'
