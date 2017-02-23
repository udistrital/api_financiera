\echo '\n#################### Creacion BD Financiera #################### \n'

--\i financiera_db/financiera_13_02_2016.sql

\echo '\n#################### Fin Creacion BD Financiera #################### \n'


\echo '\n#################### Datos Financiera #################### \n'

\i data/data_estado_orden_pago.sql
\i data/data_tipo_orden_pago.sql
\i data/data_unidad_ejecutora.sql
-- \i data/data_orden_pago.sql

\echo '\n#################### Fin Datos Financiera #################### \n'
