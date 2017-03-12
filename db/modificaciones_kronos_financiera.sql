
-- object: financiera.anulacion_registro_presupuestal | type: TABLE --
-- DROP TABLE IF EXISTS financiera.anulacion_registro_presupuestal CASCADE;
CREATE TABLE financiera.anulacion_registro_presupuestal(
	id serial NOT NULL,
	motivo character varying NOT NULL,
	fecha_registro date NOT NULL,
	tipo_anulacion character varying NOT NULL,
	CONSTRAINT pk_anulacion_rp PRIMARY KEY (id),
	CONSTRAINT ck_tipo_anulacion CHECK (tipo_anulacion::text = ANY (ARRAY['T'::character varying::text, 'P'::character varying::text]))

);
-- ddl-end --
COMMENT ON COLUMN financiera.anulacion_registro_presupuestal.id IS 'identificador del registro de la anulacion de rp';
-- ddl-end --
COMMENT ON COLUMN financiera.anulacion_registro_presupuestal.motivo IS 'motivo por el cual se realiza la anulacion';
-- ddl-end --
COMMENT ON COLUMN financiera.anulacion_registro_presupuestal.fecha_registro IS 'fecha en la que se realizo la anulacion';
-- ddl-end --
COMMENT ON COLUMN financiera.anulacion_registro_presupuestal.tipo_anulacion IS 'tipo de anulacion que se realizo (total o parcial)';
-- ddl-end --
COMMENT ON CONSTRAINT pk_anulacion_rp ON financiera.anulacion_registro_presupuestal  IS 'primary key de la tabla';
-- ddl-end --
ALTER TABLE financiera.anulacion_registro_presupuestal OWNER TO postgres;
-- ddl-end --

-- object: financiera.anulacion_registro_presupuestal_disponibilidad_apropiacion | type: TABLE --
-- DROP TABLE IF EXISTS financiera.anulacion_registro_presupuestal_disponibilidad_apropiacion CASCADE;
CREATE TABLE financiera.anulacion_registro_presupuestal_disponibilidad_apropiacion(
	id serial NOT NULL,
	anulacion_registro_presupuestal integer NOT NULL,
	registro_presupuestal_disponibilidad_apropiacion integer NOT NULL,
	valor numeric(30,4) NOT NULL,
	CONSTRAINT pk_anulacion_registro_presupuestal_apropiacion PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON COLUMN financiera.anulacion_registro_presupuestal_disponibilidad_apropiacion.id IS 'identificador unico de la anulacion a la apropiacion del CDP';
-- ddl-end --
COMMENT ON COLUMN financiera.anulacion_registro_presupuestal_disponibilidad_apropiacion.anulacion_registro_presupuestal IS 'identificador de la info general de la anulacion';
-- ddl-end --
COMMENT ON COLUMN financiera.anulacion_registro_presupuestal_disponibilidad_apropiacion.registro_presupuestal_disponibilidad_apropiacion IS 'identificador de la apropiacion del CDP a donde se dirige el RP';
-- ddl-end --
COMMENT ON COLUMN financiera.anulacion_registro_presupuestal_disponibilidad_apropiacion.valor IS 'valor anulado del RP';
-- ddl-end --
COMMENT ON CONSTRAINT pk_anulacion_registro_presupuestal_apropiacion ON financiera.anulacion_registro_presupuestal_disponibilidad_apropiacion  IS 'primary key del detalle de la anulacion';
-- ddl-end --
ALTER TABLE financiera.anulacion_registro_presupuestal_disponibilidad_apropiacion OWNER TO postgres;
-- ddl-end --

-- object: fk_anulacion_registro_presupuestal | type: CONSTRAINT --
-- ALTER TABLE financiera.anulacion_registro_presupuestal_disponibilidad_apropiacion DROP CONSTRAINT IF EXISTS fk_anulacion_registro_presupuestal CASCADE;
ALTER TABLE financiera.anulacion_registro_presupuestal_disponibilidad_apropiacion ADD CONSTRAINT fk_anulacion_registro_presupuestal FOREIGN KEY (anulacion_registro_presupuestal)
REFERENCES financiera.anulacion_registro_presupuestal (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: fk_registro_presupuestal_disponibilidad_anulacion | type: CONSTRAINT --
-- ALTER TABLE financiera.anulacion_registro_presupuestal_disponibilidad_apropiacion DROP CONSTRAINT IF EXISTS fk_registro_presupuestal_disponibilidad_anulacion CASCADE;
ALTER TABLE financiera.anulacion_registro_presupuestal_disponibilidad_apropiacion ADD CONSTRAINT fk_registro_presupuestal_disponibilidad_anulacion FOREIGN KEY (registro_presupuestal_disponibilidad_apropiacion)
REFERENCES financiera.registro_presupuestal_disponibilidad_apropiacion (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --
--modificacion RP para referenciar el compromiso
ALTER TABLE financiera.registro_presupuestal
  ADD COLUMN compromiso integer;
ALTER TABLE financiera.registro_presupuestal
  ADD CONSTRAINT fk_compromiso FOREIGN KEY (compromiso) REFERENCES financiera.compromiso (id) ON UPDATE NO ACTION ON DELETE NO ACTION;
COMMENT ON COLUMN financiera.registro_presupuestal.compromiso IS 'refrencia al compromiso del RP';
