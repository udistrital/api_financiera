-- Database generated with pgModeler (PostgreSQL Database Modeler).
-- pgModeler  version: 0.7.2
-- PostgreSQL version: 9.4
-- Project Site: pgmodeler.com.br
-- Model Author: ---

SET check_function_bodies = false;
-- ddl-end --


-- Database creation must be done outside an multicommand file.
-- These commands were put in this file only for convenience.
-- -- object: union_financiera | type: DATABASE --
-- -- DROP DATABASE union_financiera;
-- CREATE DATABASE union_financiera
-- 	ENCODING = 'UTF8'
-- 	LC_COLLATE = 'es_CO.UTF8'
-- 	LC_CTYPE = 'es_CO.UTF8'
-- 	TABLESPACE = pg_default
-- 	OWNER = postgres
-- ;
-- -- ddl-end --
-- 

-- object: financiera | type: SCHEMA --
-- DROP SCHEMA financiera;
CREATE SCHEMA financiera;
ALTER SCHEMA financiera OWNER TO postgres;
COMMENT ON SCHEMA financiera IS 'esquema para el módulo presupuestal del sistema financiero';
-- ddl-end --

SET search_path TO pg_catalog,public,financiera;
-- ddl-end --

-- object: financiera.anulacion_cdp_id_seq | type: SEQUENCE --
-- DROP SEQUENCE financiera.anulacion_cdp_id_seq;
CREATE SEQUENCE financiera.anulacion_cdp_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
ALTER SEQUENCE financiera.anulacion_cdp_id_seq OWNER TO postgres;
-- ddl-end --

-- object: financiera.anulacion_disponibilidad | type: TABLE --
-- DROP TABLE financiera.anulacion_disponibilidad;
CREATE  TABLE financiera.anulacion_disponibilidad(
	id integer NOT NULL DEFAULT nextval('anulacion_cdp_id_seq'::regclass),
	motivo character varying(600) NOT NULL,
	fecha_registro date NOT NULL,
	tipo_anulacion character varying NOT NULL,
	CONSTRAINT ck_tipo_anulacion CHECK (((tipo_anulacion)::text = ANY (ARRAY[('T'::character varying)::text, ('P'::character varying)::text]))),
	CONSTRAINT pk_anulacion_cdp PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE financiera.anulacion_disponibilidad IS 'registro de detalle de las anulaciones de CDP';
COMMENT ON CONSTRAINT ck_tipo_anulacion ON financiera.anulacion_disponibilidad IS 'si la anulacion es total o parcial';
ALTER TABLE financiera.anulacion_disponibilidad OWNER TO postgres;
-- ddl-end --

-- object: financiera.anulacion_disponibilidad_apropiacion_id_seq | type: SEQUENCE --
-- DROP SEQUENCE financiera.anulacion_disponibilidad_apropiacion_id_seq;
CREATE SEQUENCE financiera.anulacion_disponibilidad_apropiacion_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
ALTER SEQUENCE financiera.anulacion_disponibilidad_apropiacion_id_seq OWNER TO postgres;
-- ddl-end --

-- object: financiera.anulacion_disponibilidad_apropiacion | type: TABLE --
-- DROP TABLE financiera.anulacion_disponibilidad_apropiacion;
CREATE  TABLE financiera.anulacion_disponibilidad_apropiacion(
	id integer NOT NULL DEFAULT nextval('anulacion_disponibilidad_apropiacion_id_seq'::regclass),
	disponibilidad_apropiacion bigint NOT NULL,
	valor numeric(18,5) NOT NULL,
	anulacion bigint NOT NULL,
	CONSTRAINT pk_anulacion_disponibilidad_apropiacion PRIMARY KEY (id)

);
-- ddl-end --
ALTER TABLE financiera.anulacion_disponibilidad_apropiacion OWNER TO postgres;
-- ddl-end --

-- object: financiera.anulacion_reserva_id_seq | type: SEQUENCE --
-- DROP SEQUENCE financiera.anulacion_reserva_id_seq;
CREATE SEQUENCE financiera.anulacion_reserva_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
ALTER SEQUENCE financiera.anulacion_reserva_id_seq OWNER TO postgres;
-- ddl-end --

-- object: financiera.anulacion_reserva | type: TABLE --
-- DROP TABLE financiera.anulacion_reserva;
CREATE  TABLE financiera.anulacion_reserva(
	id integer NOT NULL DEFAULT nextval('anulacion_reserva_id_seq'::regclass),
	fecha_anulacion date NOT NULL,
	reserva_presupuestal integer NOT NULL,
	valor numeric(16,2) NOT NULL,
	CONSTRAINT pk_anulacion PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE financiera.anulacion_reserva IS 'anulaciones de reservas presupuestales';
COMMENT ON COLUMN financiera.anulacion_reserva.id IS 'identificador de la reserva';
COMMENT ON COLUMN financiera.anulacion_reserva.fecha_anulacion IS 'fecha de la anulación';
COMMENT ON COLUMN financiera.anulacion_reserva.reserva_presupuestal IS 'identificador de la reserva asociada';
COMMENT ON COLUMN financiera.anulacion_reserva.valor IS 'valor de la anulación';
ALTER TABLE financiera.anulacion_reserva OWNER TO postgres;
-- ddl-end --

-- object: financiera.apropiacion_id_seq | type: SEQUENCE --
-- DROP SEQUENCE financiera.apropiacion_id_seq;
CREATE SEQUENCE financiera.apropiacion_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
ALTER SEQUENCE financiera.apropiacion_id_seq OWNER TO postgres;
-- ddl-end --

-- object: financiera.apropiacion | type: TABLE --
-- DROP TABLE financiera.apropiacion;
CREATE  TABLE financiera.apropiacion(
	id integer NOT NULL DEFAULT nextval('apropiacion_id_seq'::regclass),
	vigencia numeric(4,0),
	rubro integer,
	unidad_ejecutora integer,
	valor_rezago numeric(16,2),
	valor numeric(16,2),
	tipo_documento character varying(30),
	documento_numero character varying(10),
	documento_fecha date,
	estado integer NOT NULL DEFAULT 1,
	CONSTRAINT pk_apropiacion PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE financiera.apropiacion IS 'apropiaciones iniciales';
COMMENT ON COLUMN financiera.apropiacion.id IS 'identificador de la apropiación';
COMMENT ON COLUMN financiera.apropiacion.vigencia IS 'año de la vigencia';
COMMENT ON COLUMN financiera.apropiacion.rubro IS 'rubro asociado';
COMMENT ON COLUMN financiera.apropiacion.unidad_ejecutora IS 'unidad ejecutora';
COMMENT ON COLUMN financiera.apropiacion.valor_rezago IS 'valor rezago de la apropiación';
COMMENT ON COLUMN financiera.apropiacion.valor IS 'valor de la apropiación';
COMMENT ON COLUMN financiera.apropiacion.tipo_documento IS 'tipo de documento';
COMMENT ON COLUMN financiera.apropiacion.documento_numero IS 'número del documento';
COMMENT ON COLUMN financiera.apropiacion.documento_fecha IS 'fecha del documento';
ALTER TABLE financiera.apropiacion OWNER TO postgres;
-- ddl-end --

-- object: financiera.concepto_id_seq | type: SEQUENCE --
-- DROP SEQUENCE financiera.concepto_id_seq;
CREATE SEQUENCE financiera.concepto_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
ALTER SEQUENCE financiera.concepto_id_seq OWNER TO postgres;
-- ddl-end --

-- object: financiera.desagregacion_ingreso_id_seq | type: SEQUENCE --
-- DROP SEQUENCE financiera.desagregacion_ingreso_id_seq;
CREATE SEQUENCE financiera.desagregacion_ingreso_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
ALTER SEQUENCE financiera.desagregacion_ingreso_id_seq OWNER TO postgres;
-- ddl-end --

-- object: financiera.destino_disponibilidad_id_seq | type: SEQUENCE --
-- DROP SEQUENCE financiera.destino_disponibilidad_id_seq;
CREATE SEQUENCE financiera.destino_disponibilidad_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
ALTER SEQUENCE financiera.destino_disponibilidad_id_seq OWNER TO postgres;
-- ddl-end --

-- object: financiera.disponibilidad_id_seq | type: SEQUENCE --
-- DROP SEQUENCE financiera.disponibilidad_id_seq;
CREATE SEQUENCE financiera.disponibilidad_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
ALTER SEQUENCE financiera.disponibilidad_id_seq OWNER TO postgres;
-- ddl-end --

-- object: financiera.disponibilidad_apropiacion_id_seq | type: SEQUENCE --
-- DROP SEQUENCE financiera.disponibilidad_apropiacion_id_seq;
CREATE SEQUENCE financiera.disponibilidad_apropiacion_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
ALTER SEQUENCE financiera.disponibilidad_apropiacion_id_seq OWNER TO postgres;
-- ddl-end --

-- object: financiera.disponibilidad_apropiacion | type: TABLE --
-- DROP TABLE financiera.disponibilidad_apropiacion;
CREATE  TABLE financiera.disponibilidad_apropiacion(
	id integer NOT NULL DEFAULT nextval('disponibilidad_apropiacion_id_seq'::regclass),
	disponibilidad bigint,
	apropiacion bigint,
	valor numeric(25,7),
	CONSTRAINT pk_disponibilidad_apropiacion PRIMARY KEY (id),
	CONSTRAINT uq_apropiacion_disponibilidad UNIQUE (disponibilidad,apropiacion)

);
-- ddl-end --
COMMENT ON TABLE financiera.disponibilidad_apropiacion IS 'tabla para registrar la desagrgación presupuestal de la apropiacion en CDP';
COMMENT ON COLUMN financiera.disponibilidad_apropiacion.disponibilidad IS 'referencia al CDP que solicita el monto de la disponibilidad';
COMMENT ON COLUMN financiera.disponibilidad_apropiacion.apropiacion IS 'referencia a la apropiacion objetivo del CDP';
COMMENT ON COLUMN financiera.disponibilidad_apropiacion.valor IS 'valor del CDP';
ALTER TABLE financiera.disponibilidad_apropiacion OWNER TO postgres;
-- ddl-end --

-- object: financiera.disponibilidad | type: TABLE --
-- DROP TABLE financiera.disponibilidad;
CREATE  TABLE financiera.disponibilidad(
	id integer NOT NULL DEFAULT nextval('disponibilidad_id_seq'::regclass),
	unidad_ejecutora integer NOT NULL,
	vigencia numeric(4,0) NOT NULL,
	numero_disponibilidad numeric(6,0),
	responsable bigint,
	solicitante bigint,
	fecha_registro date,
	modalidad_giro smallint,
	estado smallint,
	numero_oficio character varying(10),
	objeto character varying(400),
	vigencia_futura numeric,
	destino bigint,
	solicitud bigint NOT NULL,
	CONSTRAINT pk_id_entidad PRIMARY KEY (id),
	CONSTRAINT uq_disponibilidad_vigencia_numero UNIQUE (vigencia,numero_disponibilidad)

);
-- ddl-end --
COMMENT ON COLUMN financiera.disponibilidad.id IS 'identificador de la disponibilidad';
COMMENT ON COLUMN financiera.disponibilidad.unidad_ejecutora IS 'codigo de la unidad ejecutora';
COMMENT ON COLUMN financiera.disponibilidad.vigencia IS 'Vigencia de la disponibilidad';
COMMENT ON COLUMN financiera.disponibilidad.numero_disponibilidad IS 'número de la disponibilidad';
COMMENT ON COLUMN financiera.disponibilidad.responsable IS 'crear tabla de usuarios para este rol';
COMMENT ON COLUMN financiera.disponibilidad.solicitante IS 'sacar de la tabla de usuarios el id';
COMMENT ON COLUMN financiera.disponibilidad.fecha_registro IS 'fecha de registro de la disponibilidad';
COMMENT ON COLUMN financiera.disponibilidad.modalidad_giro IS 'campo paramétrico';
COMMENT ON COLUMN financiera.disponibilidad.estado IS 'paramétrica estado disponibilidad';
COMMENT ON COLUMN financiera.disponibilidad.numero_oficio IS 'numero del oficio asociado';
COMMENT ON COLUMN financiera.disponibilidad.objeto IS 'objeto de la disponibilidad';
COMMENT ON COLUMN financiera.disponibilidad.vigencia_futura IS 'indicador de si es para vigencias futuras';
COMMENT ON COLUMN financiera.disponibilidad.destino IS 'paramétrica destino disponibilidad por ejemplo modificación o anulación';
ALTER TABLE financiera.disponibilidad OWNER TO postgres;
-- ddl-end --

-- object: financiera.disponibilidad_rubro_id_seq | type: SEQUENCE --
-- DROP SEQUENCE financiera.disponibilidad_rubro_id_seq;
CREATE SEQUENCE financiera.disponibilidad_rubro_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
ALTER SEQUENCE financiera.disponibilidad_rubro_id_seq OWNER TO postgres;
-- ddl-end --

-- object: financiera.estado_apropiacion | type: TABLE --
-- DROP TABLE financiera.estado_apropiacion;
CREATE  TABLE financiera.estado_apropiacion(
	id integer NOT NULL,
	nombre character varying(20),
	descripcion character varying(300),
	CONSTRAINT pk_id_estado_apropiacion PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE financiera.estado_apropiacion IS 'paramétrica estado de apropiaciones';
COMMENT ON COLUMN financiera.estado_apropiacion.id IS 'identificador único del registro';
COMMENT ON COLUMN financiera.estado_apropiacion.nombre IS 'nombre del estado, será el mostrado';
COMMENT ON COLUMN financiera.estado_apropiacion.descripcion IS 'descripción del estado';
ALTER TABLE financiera.estado_apropiacion OWNER TO postgres;
-- ddl-end --

-- object: financiera.estado_disponibilidad_id_seq | type: SEQUENCE --
-- DROP SEQUENCE financiera.estado_disponibilidad_id_seq;
CREATE SEQUENCE financiera.estado_disponibilidad_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
ALTER SEQUENCE financiera.estado_disponibilidad_id_seq OWNER TO postgres;
-- ddl-end --

-- object: financiera.estado_disponibilidad | type: TABLE --
-- DROP TABLE financiera.estado_disponibilidad;
CREATE  TABLE financiera.estado_disponibilidad(
	id integer NOT NULL DEFAULT nextval('estado_disponibilidad_id_seq'::regclass),
	nombre character varying(25) NOT NULL,
	descripcion character varying(150) NOT NULL,
	CONSTRAINT pk_id_estado_disponibilidad PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE financiera.estado_disponibilidad IS 'paramétrica estado de disponibilidades';
COMMENT ON COLUMN financiera.estado_disponibilidad.descripcion IS 'descripción del estado';
ALTER TABLE financiera.estado_disponibilidad OWNER TO postgres;
-- ddl-end --

-- object: financiera.estado_registro_presupuestal_id_seq | type: SEQUENCE --
-- DROP SEQUENCE financiera.estado_registro_presupuestal_id_seq;
CREATE SEQUENCE financiera.estado_registro_presupuestal_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
ALTER SEQUENCE financiera.estado_registro_presupuestal_id_seq OWNER TO postgres;
-- ddl-end --

-- object: financiera.estado_registro_presupuestal | type: TABLE --
-- DROP TABLE financiera.estado_registro_presupuestal;
CREATE  TABLE financiera.estado_registro_presupuestal(
	id integer NOT NULL DEFAULT nextval('estado_registro_presupuestal_id_seq'::regclass),
	nombre character varying(25) NOT NULL,
	descripcion character varying(150) NOT NULL,
	CONSTRAINT pk_id_estado_registro PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE financiera.estado_registro_presupuestal IS 'paramétrica estado de registros';
COMMENT ON COLUMN financiera.estado_registro_presupuestal.nombre IS 'nombre del estado';
COMMENT ON COLUMN financiera.estado_registro_presupuestal.descripcion IS 'descripción del estado';
ALTER TABLE financiera.estado_registro_presupuestal OWNER TO postgres;
-- ddl-end --

-- object: financiera.estado_reserva_presupuestal_id_seq | type: SEQUENCE --
-- DROP SEQUENCE financiera.estado_reserva_presupuestal_id_seq;
CREATE SEQUENCE financiera.estado_reserva_presupuestal_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
ALTER SEQUENCE financiera.estado_reserva_presupuestal_id_seq OWNER TO postgres;
-- ddl-end --

-- object: financiera.estado_reserva_presupuestal | type: TABLE --
-- DROP TABLE financiera.estado_reserva_presupuestal;
CREATE  TABLE financiera.estado_reserva_presupuestal(
	id integer NOT NULL DEFAULT nextval('estado_reserva_presupuestal_id_seq'::regclass),
	nombre character varying(15) NOT NULL,
	descripcion character varying(100) NOT NULL,
	CONSTRAINT pk_estado_reserva PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE financiera.estado_reserva_presupuestal IS 'estados de las reservas presupuestales';
COMMENT ON COLUMN financiera.estado_reserva_presupuestal.id IS 'identificador del estado de reserva';
COMMENT ON COLUMN financiera.estado_reserva_presupuestal.nombre IS 'nombre del estado';
COMMENT ON COLUMN financiera.estado_reserva_presupuestal.descripcion IS 'descripción del estado';
ALTER TABLE financiera.estado_reserva_presupuestal OWNER TO postgres;
-- ddl-end --

-- object: financiera.fuente_financiacion_id_seq | type: SEQUENCE --
-- DROP SEQUENCE financiera.fuente_financiacion_id_seq;
CREATE SEQUENCE financiera.fuente_financiacion_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
ALTER SEQUENCE financiera.fuente_financiacion_id_seq OWNER TO postgres;
-- ddl-end --

-- object: financiera.fuente_financiacion_entidad_id_seq | type: SEQUENCE --
-- DROP SEQUENCE financiera.fuente_financiacion_entidad_id_seq;
CREATE SEQUENCE financiera.fuente_financiacion_entidad_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
ALTER SEQUENCE financiera.fuente_financiacion_entidad_id_seq OWNER TO postgres;
-- ddl-end --

-- object: financiera.ingreso_id_seq | type: SEQUENCE --
-- DROP SEQUENCE financiera.ingreso_id_seq;
CREATE SEQUENCE financiera.ingreso_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
ALTER SEQUENCE financiera.ingreso_id_seq OWNER TO postgres;
-- ddl-end --

-- object: financiera.modificacion_presupuestal_id_seq | type: SEQUENCE --
-- DROP SEQUENCE financiera.modificacion_presupuestal_id_seq;
CREATE SEQUENCE financiera.modificacion_presupuestal_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
ALTER SEQUENCE financiera.modificacion_presupuestal_id_seq OWNER TO postgres;
-- ddl-end --

-- object: financiera.modificacion_presupuestal | type: TABLE --
-- DROP TABLE financiera.modificacion_presupuestal;
CREATE  TABLE financiera.modificacion_presupuestal(
	id integer NOT NULL DEFAULT nextval('modificacion_presupuestal_id_seq'::regclass),
	apropiacion smallint,
	tipo_documento character varying(30),
	documento_numero character varying(10),
	tipo_movimiento character varying(30),
	numero_disponibilidad numeric(6,0),
	fecha_registro date,
	valor_credito numeric(16,2),
	valor_contra_credito numeric(16,2),
	CONSTRAINT pk_modificacion_presupuestal PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON COLUMN financiera.modificacion_presupuestal.id IS 'identificador de la modificación';
COMMENT ON COLUMN financiera.modificacion_presupuestal.apropiacion IS 'apropiación asociada';
COMMENT ON COLUMN financiera.modificacion_presupuestal.tipo_documento IS 'tipo de documento';
COMMENT ON COLUMN financiera.modificacion_presupuestal.tipo_movimiento IS 'tipo de movimiento';
COMMENT ON COLUMN financiera.modificacion_presupuestal.numero_disponibilidad IS 'número de la disponibilidad que modifica la apropiación';
COMMENT ON COLUMN financiera.modificacion_presupuestal.fecha_registro IS 'fecha del registro';
COMMENT ON COLUMN financiera.modificacion_presupuestal.valor_credito IS 'valor de credito';
COMMENT ON COLUMN financiera.modificacion_presupuestal.valor_contra_credito IS 'valor contra credito';
ALTER TABLE financiera.modificacion_presupuestal OWNER TO postgres;
-- ddl-end --

-- object: financiera.registo_presupuestal_id_seq | type: SEQUENCE --
-- DROP SEQUENCE financiera.registo_presupuestal_id_seq;
CREATE SEQUENCE financiera.registo_presupuestal_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
ALTER SEQUENCE financiera.registo_presupuestal_id_seq OWNER TO postgres;
-- ddl-end --

-- object: financiera.registro_presupuestal | type: TABLE --
-- DROP TABLE financiera.registro_presupuestal;
CREATE  TABLE financiera.registro_presupuestal(
	id integer NOT NULL DEFAULT nextval('registo_presupuestal_id_seq'::regclass),
	unidad_ejecutora smallint NOT NULL,
	vigencia numeric(4,0) NOT NULL,
	fecha_movimiento date,
	responsable integer,
	estado smallint NOT NULL,
	numero_registro_presupuestal integer NOT NULL,
	beneficiario integer,
	CONSTRAINT pk_registro PRIMARY KEY (id),
	CONSTRAINT uq_registro_presupuestal_vigencia_numero UNIQUE (vigencia,numero_registro_presupuestal)

);
-- ddl-end --
COMMENT ON COLUMN financiera.registro_presupuestal.id IS 'identificador del registro';
COMMENT ON COLUMN financiera.registro_presupuestal.unidad_ejecutora IS 'unidad ejecutora';
COMMENT ON COLUMN financiera.registro_presupuestal.vigencia IS 'vigencia del registro presupuestal';
COMMENT ON COLUMN financiera.registro_presupuestal.fecha_movimiento IS 'fecha del movimiento';
COMMENT ON COLUMN financiera.registro_presupuestal.responsable IS 'crear tabla de usuarios para asignar la responsabilidad';
COMMENT ON COLUMN financiera.registro_presupuestal.estado IS 'estado del registros';
ALTER TABLE financiera.registro_presupuestal OWNER TO postgres;
-- ddl-end --

-- object: financiera.registro_presupuestal_disponibilidad_id_seq | type: SEQUENCE --
-- DROP SEQUENCE financiera.registro_presupuestal_disponibilidad_id_seq;
CREATE SEQUENCE financiera.registro_presupuestal_disponibilidad_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
ALTER SEQUENCE financiera.registro_presupuestal_disponibilidad_id_seq OWNER TO postgres;
-- ddl-end --

-- object: financiera.registro_presupuestal_disponibilidad_apropiacion_id_seq | type: SEQUENCE --
-- DROP SEQUENCE financiera.registro_presupuestal_disponibilidad_apropiacion_id_seq;
CREATE SEQUENCE financiera.registro_presupuestal_disponibilidad_apropiacion_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
ALTER SEQUENCE financiera.registro_presupuestal_disponibilidad_apropiacion_id_seq OWNER TO postgres;
-- ddl-end --

-- object: financiera.registro_presupuestal_disponibilidad_apropiacion | type: TABLE --
-- DROP TABLE financiera.registro_presupuestal_disponibilidad_apropiacion;
CREATE  TABLE financiera.registro_presupuestal_disponibilidad_apropiacion(
	id integer NOT NULL DEFAULT nextval('registro_presupuestal_disponibilidad_apropiacion_id_seq'::regclass),
	registro_presupuestal bigint,
	disponibilidad_apropiacion bigint,
	valor numeric(20,5),
	CONSTRAINT pk_registro_presupuestal_disponibilidad_apropiacion PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON COLUMN financiera.registro_presupuestal_disponibilidad_apropiacion.id IS 'serial de identificacion ';
COMMENT ON COLUMN financiera.registro_presupuestal_disponibilidad_apropiacion.registro_presupuestal IS 'referencia al rp expedido';
COMMENT ON COLUMN financiera.registro_presupuestal_disponibilidad_apropiacion.disponibilidad_apropiacion IS 'referencia al dinero reservado por el cdp para un rubro';
COMMENT ON COLUMN financiera.registro_presupuestal_disponibilidad_apropiacion.valor IS 'valor a tomar de la apropiacion del CDP objetivo';
ALTER TABLE financiera.registro_presupuestal_disponibilidad_apropiacion OWNER TO postgres;
-- ddl-end --

-- object: financiera.reserva_presupuestal_id_seq | type: SEQUENCE --
-- DROP SEQUENCE financiera.reserva_presupuestal_id_seq;
CREATE SEQUENCE financiera.reserva_presupuestal_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
ALTER SEQUENCE financiera.reserva_presupuestal_id_seq OWNER TO postgres;
-- ddl-end --

-- object: financiera.reserva_presupuestal | type: TABLE --
-- DROP TABLE financiera.reserva_presupuestal;
CREATE  TABLE financiera.reserva_presupuestal(
	id integer NOT NULL DEFAULT nextval('reserva_presupuestal_id_seq'::regclass),
	unidad_ejecutora smallint NOT NULL,
	vigencia numeric(4,0),
	fecha_movimiento date,
	responsable smallint,
	estado smallint NOT NULL,
	rubro integer NOT NULL,
	valor numeric(16,2) NOT NULL,
	CONSTRAINT pk_reserva_presupuestal PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE financiera.reserva_presupuestal IS 'reservas presupuestales';
COMMENT ON COLUMN financiera.reserva_presupuestal.id IS 'identificador reserva';
COMMENT ON COLUMN financiera.reserva_presupuestal.unidad_ejecutora IS 'unidad ejecutora';
COMMENT ON COLUMN financiera.reserva_presupuestal.fecha_movimiento IS 'fecha del movimiento';
COMMENT ON COLUMN financiera.reserva_presupuestal.responsable IS 'crear tabla de usuarios para asignar la responsabilidad';
COMMENT ON COLUMN financiera.reserva_presupuestal.estado IS 'crear tabla de estado_reserva';
COMMENT ON COLUMN financiera.reserva_presupuestal.rubro IS 'identificador del rubro asociado a la reserva';
COMMENT ON COLUMN financiera.reserva_presupuestal.valor IS 'valor de la reserva';
ALTER TABLE financiera.reserva_presupuestal OWNER TO postgres;
-- ddl-end --

-- object: financiera.rubro_id_seq | type: SEQUENCE --
-- DROP SEQUENCE financiera.rubro_id_seq;
CREATE SEQUENCE financiera.rubro_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
ALTER SEQUENCE financiera.rubro_id_seq OWNER TO postgres;
-- ddl-end --

-- object: financiera.rubro_homologado_id_seq | type: SEQUENCE --
-- DROP SEQUENCE financiera.rubro_homologado_id_seq;
CREATE SEQUENCE financiera.rubro_homologado_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
ALTER SEQUENCE financiera.rubro_homologado_id_seq OWNER TO postgres;
-- ddl-end --

-- object: financiera.rubro_homologado | type: TABLE --
-- DROP TABLE financiera.rubro_homologado;
CREATE  TABLE financiera.rubro_homologado(
	id integer NOT NULL DEFAULT nextval('rubro_homologado_id_seq'::regclass),
	rubro smallint,
	codigo_homologado character varying(15) NOT NULL,
	nombre_homologado character varying(60) NOT NULL,
	entidad_homologado integer NOT NULL,
	vigencia numeric(4,0) NOT NULL,
	CONSTRAINT pk_rubro_homologado PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON COLUMN financiera.rubro_homologado.id IS 'identificador del rubro homologado';
COMMENT ON COLUMN financiera.rubro_homologado.rubro IS 'rubro a homologar';
COMMENT ON COLUMN financiera.rubro_homologado.codigo_homologado IS 'codigo de homologación';
COMMENT ON COLUMN financiera.rubro_homologado.nombre_homologado IS 'nombre del codigo homologado';
COMMENT ON COLUMN financiera.rubro_homologado.entidad_homologado IS 'identificador de la entidad de homologación';
COMMENT ON COLUMN financiera.rubro_homologado.vigencia IS 'vigencia de la homologación';
ALTER TABLE financiera.rubro_homologado OWNER TO postgres;
-- ddl-end --

-- object: financiera.rubro | type: TABLE --
-- DROP TABLE financiera.rubro;
CREATE  TABLE financiera.rubro(
	id integer NOT NULL DEFAULT nextval('rubro_id_seq'::regclass),
	entidad integer NOT NULL,
	codigo character varying(100) NOT NULL,
	vigencia numeric(4,0) NOT NULL,
	descripcion text,
	tipo_plan smallint,
	administracion character varying(30),
	estado smallint,
	CONSTRAINT "CK_rubro_codigo" CHECK (((codigo)::text ~ '^([0-9]+-){0,7}[0-9]+$'::text)),
	CONSTRAINT pk_rubro PRIMARY KEY (id),
	CONSTRAINT uq_rubro_codigo_vigencia UNIQUE (codigo,vigencia)

);
-- ddl-end --
COMMENT ON COLUMN financiera.rubro.id IS 'identificador del rubro';
COMMENT ON COLUMN financiera.rubro.entidad IS 'entidad a la que pertenece el rubro';
COMMENT ON COLUMN financiera.rubro.codigo IS 'codigo interno';
COMMENT ON COLUMN financiera.rubro.vigencia IS 'vigencia del rubro';
COMMENT ON COLUMN financiera.rubro.descripcion IS 'descripción del rubro';
COMMENT ON COLUMN financiera.rubro.tipo_plan IS 'tipo plan de cuentas';
COMMENT ON COLUMN financiera.rubro.administracion IS 'Nombre De La Administracion';
COMMENT ON COLUMN financiera.rubro.estado IS 'parámetrica de estados del rubro';
ALTER TABLE financiera.rubro OWNER TO postgres;
-- ddl-end --

-- object: financiera.rubro_rubro_id_seq | type: SEQUENCE --
-- DROP SEQUENCE financiera.rubro_rubro_id_seq;
CREATE SEQUENCE financiera.rubro_rubro_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
ALTER SEQUENCE financiera.rubro_rubro_id_seq OWNER TO postgres;
-- ddl-end --

-- object: financiera.rubro_rubro | type: TABLE --
-- DROP TABLE financiera.rubro_rubro;
CREATE  TABLE financiera.rubro_rubro(
	id integer NOT NULL DEFAULT nextval('rubro_rubro_id_seq'::regclass),
	rubro_padre integer,
	rubro_hijo integer,
	CONSTRAINT pk_rubro_rubro PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE financiera.rubro_rubro IS 'Tabla para normalizar la relacion reflexiva entre rubros';
ALTER TABLE financiera.rubro_rubro OWNER TO postgres;
-- ddl-end --

-- object: financiera.cuenta_contable_id_seq | type: SEQUENCE --
-- DROP SEQUENCE financiera.cuenta_contable_id_seq;
CREATE SEQUENCE financiera.cuenta_contable_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
ALTER SEQUENCE financiera.cuenta_contable_id_seq OWNER TO postgres;
-- ddl-end --

-- object: financiera.cuenta_contable | type: TABLE --
-- DROP TABLE financiera.cuenta_contable;
CREATE  TABLE financiera.cuenta_contable(
	id integer NOT NULL DEFAULT nextval('cuenta_contable_id_seq'::regclass),
	saldo numeric(12,0) NOT NULL,
	nombre character varying(250) NOT NULL,
	naturaleza character varying(7) NOT NULL,
	descripcion character varying,
	codigo character varying NOT NULL,
	nivel_clasificacion integer NOT NULL,
	cuenta_bancaria integer,
	CONSTRAINT naturaleza_ck CHECK (((naturaleza)::text = ANY ((ARRAY['debito'::character varying, 'credito'::character varying])::text[]))),
	CONSTRAINT cuenta_contable_pk PRIMARY KEY (id),
	CONSTRAINT codigo_uq UNIQUE (codigo)

);
-- ddl-end --
COMMENT ON TABLE financiera.cuenta_contable IS 'tabla en la que se registran las cuentas contables ';
COMMENT ON COLUMN financiera.cuenta_contable.id IS 'identificador de la tabla cuenta_contable';
COMMENT ON COLUMN financiera.cuenta_contable.saldo IS 'campo en el que se registra el saldo de la cuenta contable y es modificado segun sus movimientos';
COMMENT ON COLUMN financiera.cuenta_contable.nombre IS 'campo que indica el nombre de la cuenta contable';
COMMENT ON COLUMN financiera.cuenta_contable.naturaleza IS 'campo en el que se indica la naturaleza con la cuenta la cuenta contable registrada (debito o credito)';
COMMENT ON COLUMN financiera.cuenta_contable.descripcion IS 'campo en el que se puede indicar una descripcion sobre la cuenta contable registrada';
COMMENT ON COLUMN financiera.cuenta_contable.codigo IS 'campo que indica el codigo consecutivo unico para diferenciar una cuenta contable';
COMMENT ON COLUMN financiera.cuenta_contable.nivel_clasificacion IS 'identificador de la tabla nivel_clasificacion referenciado para diferenciar a que nivel pertenece la cuenta contable registrada';
COMMENT ON COLUMN financiera.cuenta_contable.cuenta_bancaria IS 'identificador de la tabla cuenta_bancaria referemciado para confirmar si la cuenta contable se asocia a una bancaria';
ALTER TABLE financiera.cuenta_contable OWNER TO postgres;
-- ddl-end --

-- object: financiera.movimiento_contable_id_seq | type: SEQUENCE --
-- DROP SEQUENCE financiera.movimiento_contable_id_seq;
CREATE SEQUENCE financiera.movimiento_contable_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
ALTER SEQUENCE financiera.movimiento_contable_id_seq OWNER TO postgres;
-- ddl-end --

-- object: financiera.movimiento_contable | type: TABLE --
-- DROP TABLE financiera.movimiento_contable;
CREATE  TABLE financiera.movimiento_contable(
	id integer NOT NULL DEFAULT nextval('movimiento_contable_id_seq'::regclass),
	debito numeric NOT NULL,
	credito numeric NOT NULL,
	fecha timestamp NOT NULL,
	cuenta_contable integer NOT NULL,
	tipo_documento_afectante integer NOT NULL,
	codigo_documento_afectante integer NOT NULL,
	CONSTRAINT movimiento_contable_pk PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE financiera.movimiento_contable IS 'tabla en la que se registran los movimientos contables que son realizados sobre una cuenta contable';
COMMENT ON COLUMN financiera.movimiento_contable.id IS 'identificador de la tabla movimiento_contable';
COMMENT ON COLUMN financiera.movimiento_contable.debito IS 'campo en el que se debe indical el valor que se registro por el lado debito en un movimiento';
COMMENT ON COLUMN financiera.movimiento_contable.credito IS 'campo en el que se debe indical el valor que se registro por el lado credito en un movimiento';
COMMENT ON COLUMN financiera.movimiento_contable.fecha IS 'campo en el que se indica la fecha y hora del movimiento realizado';
COMMENT ON COLUMN financiera.movimiento_contable.cuenta_contable IS 'identificador de la tabla cuenta_contable que se referencia para diferenciar en que cuenta se realizo el movimiento';
COMMENT ON COLUMN financiera.movimiento_contable.tipo_documento_afectante IS 'identificador de la tabla tipo_documento_afectante referenciado para diferenciar sobre que tipo de documento se realizo el movimiento y de esta manera poder buscarlo en el modelo por su codigo.';
COMMENT ON COLUMN financiera.movimiento_contable.codigo_documento_afectante IS 'campo en el que se registra el codigo del documento en el que se realizo el movimiento, este biene dado por el tipo de documento afectante';
ALTER TABLE financiera.movimiento_contable OWNER TO postgres;
-- ddl-end --

-- object: financiera.tipo_documento_afectante_id_seq | type: SEQUENCE --
-- DROP SEQUENCE financiera.tipo_documento_afectante_id_seq;
CREATE SEQUENCE financiera.tipo_documento_afectante_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
ALTER SEQUENCE financiera.tipo_documento_afectante_id_seq OWNER TO postgres;
-- ddl-end --

-- object: financiera.tipo_documento_afectante | type: TABLE --
-- DROP TABLE financiera.tipo_documento_afectante;
CREATE  TABLE financiera.tipo_documento_afectante(
	id integer NOT NULL DEFAULT nextval('tipo_documento_afectante_id_seq'::regclass),
	nombre character varying(150) NOT NULL,
	descripcion character varying,
	CONSTRAINT tipo_documento_afectante_pk PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE financiera.tipo_documento_afectante IS 'tabla en la que se registran los documentos que pueden afectar cuentas contables, es decir en los que se realizan movimientos contables';
COMMENT ON COLUMN financiera.tipo_documento_afectante.id IS 'identificador de la tabla tipo_documento_afectante';
COMMENT ON COLUMN financiera.tipo_documento_afectante.nombre IS 'campo en el que se indica el nombre del tipo de documento afectante';
COMMENT ON COLUMN financiera.tipo_documento_afectante.descripcion IS 'campo en el que se puede registrar una descripcion hacerca del tipo de documento que afecta a la cuenta  contable';
ALTER TABLE financiera.tipo_documento_afectante OWNER TO postgres;
-- ddl-end --

-- object: financiera.estructura_cuentas_id_seq | type: SEQUENCE --
-- DROP SEQUENCE financiera.estructura_cuentas_id_seq;
CREATE SEQUENCE financiera.estructura_cuentas_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
ALTER SEQUENCE financiera.estructura_cuentas_id_seq OWNER TO postgres;
-- ddl-end --

-- object: financiera.estructura_cuentas | type: TABLE --
-- DROP TABLE financiera.estructura_cuentas;
CREATE  TABLE financiera.estructura_cuentas(
	id integer NOT NULL DEFAULT nextval('estructura_cuentas_id_seq'::regclass),
	cuenta_padre integer NOT NULL,
	cuenta_hijo integer NOT NULL,
	plan_cuentas integer NOT NULL,
	CONSTRAINT estructura_cuentas_pk PRIMARY KEY (id),
	CONSTRAINT plan_cuentas_cuenta_hijo_uq UNIQUE (cuenta_hijo,plan_cuentas)

);
-- ddl-end --
COMMENT ON TABLE financiera.estructura_cuentas IS 'tabla en la que se construyen las estructuras ligadas al plan de cuentas';
COMMENT ON COLUMN financiera.estructura_cuentas.id IS 'identificador de la tabla estructura_cuentas';
COMMENT ON COLUMN financiera.estructura_cuentas.cuenta_padre IS 'identificador de la tabla cuenta_contable referenciada para indicar que se trata de la cuenta padre';
COMMENT ON COLUMN financiera.estructura_cuentas.cuenta_hijo IS 'identificador de la tabla cuenta_contable referenciada para indicar que se trata de la cuenta hijo';
COMMENT ON COLUMN financiera.estructura_cuentas.plan_cuentas IS 'campo que indica a que plan de cuentas pertenece la estructura';
ALTER TABLE financiera.estructura_cuentas OWNER TO postgres;
-- ddl-end --

-- object: financiera.nivel_clasificacion_id_seq | type: SEQUENCE --
-- DROP SEQUENCE financiera.nivel_clasificacion_id_seq;
CREATE SEQUENCE financiera.nivel_clasificacion_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
ALTER SEQUENCE financiera.nivel_clasificacion_id_seq OWNER TO postgres;
-- ddl-end --

-- object: financiera.nivel_clasificacion | type: TABLE --
-- DROP TABLE financiera.nivel_clasificacion;
CREATE  TABLE financiera.nivel_clasificacion(
	id integer NOT NULL DEFAULT nextval('nivel_clasificacion_id_seq'::regclass),
	nombre character varying(80) NOT NULL,
	longitud numeric(2,0) NOT NULL,
	descripcion character varying,
	CONSTRAINT nivel_clasificacion_pk PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE financiera.nivel_clasificacion IS 'tabla en la que se registran los niveles para la clasificacion de las cuentas contables en la estructura del plan de cuentas';
COMMENT ON COLUMN financiera.nivel_clasificacion.id IS 'identificador de la tabla nivel_clasificacion';
COMMENT ON COLUMN financiera.nivel_clasificacion.nombre IS 'campo en el que se indica el nombre que se le da al nivel de clasificacion para las cuentas contables';
COMMENT ON COLUMN financiera.nivel_clasificacion.longitud IS 'campo en el que se indica el numero de digitos que se pueden tener en un nivel';
COMMENT ON COLUMN financiera.nivel_clasificacion.descripcion IS 'campo en el que se puede ingresar una descripcion acerca del nivel de clasificacion registrado';
ALTER TABLE financiera.nivel_clasificacion OWNER TO postgres;
-- ddl-end --

-- object: financiera.plan_cuentas_id_seq | type: SEQUENCE --
-- DROP SEQUENCE financiera.plan_cuentas_id_seq;
CREATE SEQUENCE financiera.plan_cuentas_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
ALTER SEQUENCE financiera.plan_cuentas_id_seq OWNER TO postgres;
-- ddl-end --

-- object: financiera.plan_cuentas | type: TABLE --
-- DROP TABLE financiera.plan_cuentas;
CREATE  TABLE financiera.plan_cuentas(
	id integer NOT NULL DEFAULT nextval('plan_cuentas_id_seq'::regclass),
	nombre character varying(150) NOT NULL,
	descripcion character varying,
	unidad_ejecutora integer NOT NULL,
	CONSTRAINT plan_cuentas_pk PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE financiera.plan_cuentas IS 'tabla en la que se registran el plan de cuentas maestro y los alternos';
COMMENT ON COLUMN financiera.plan_cuentas.id IS 'identificador de la tabla plan_cuentas';
COMMENT ON COLUMN financiera.plan_cuentas.nombre IS 'campo en el que se registra el nombre dado al plan de cuentas';
COMMENT ON COLUMN financiera.plan_cuentas.descripcion IS 'campo en el que se puede ingresar una descripcion al plan de cuentas';
COMMENT ON COLUMN financiera.plan_cuentas.unidad_ejecutora IS 'identificador de la tabla unidad ejecutora referenciado para indicar a que entidad y unidad ejecutora pertenece el plan de cuentas';
ALTER TABLE financiera.plan_cuentas OWNER TO postgres;
-- ddl-end --

-- object: financiera.periodo_contable_id_seq | type: SEQUENCE --
-- DROP SEQUENCE financiera.periodo_contable_id_seq;
CREATE SEQUENCE financiera.periodo_contable_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
ALTER SEQUENCE financiera.periodo_contable_id_seq OWNER TO postgres;
-- ddl-end --

-- object: financiera.periodo_contable | type: TABLE --
-- DROP TABLE financiera.periodo_contable;
CREATE  TABLE financiera.periodo_contable(
	id integer NOT NULL DEFAULT nextval('periodo_contable_id_seq'::regclass),
	estado_activo boolean NOT NULL,
	fecha_modificacion date NOT NULL,
	fecha_inicio date NOT NULL,
	fecha_fin date NOT NULL,
	descripcion character varying,
	mes numeric(2,0) NOT NULL,
	ano numeric(4,0) NOT NULL,
	CONSTRAINT periodo_contable_pk PRIMARY KEY (id),
	CONSTRAINT ano_mes_uq UNIQUE (ano,mes)

);
-- ddl-end --
COMMENT ON TABLE financiera.periodo_contable IS 'tabla encargada de registrar los periodos contables';
COMMENT ON COLUMN financiera.periodo_contable.id IS 'identificador de la tabla periodo_contable';
COMMENT ON COLUMN financiera.periodo_contable.estado_activo IS 'campo en que se identifica si un periodo se encuentra activo';
COMMENT ON COLUMN financiera.periodo_contable.fecha_modificacion IS 'campo que indica la fecha del registro o ultima modificacion del periodo';
COMMENT ON COLUMN financiera.periodo_contable.fecha_inicio IS 'campo en el que se especifica la fecha de inicio del periodo';
COMMENT ON COLUMN financiera.periodo_contable.fecha_fin IS 'campo en el que se especifica la fecha de finalizacion del periodo';
COMMENT ON COLUMN financiera.periodo_contable.descripcion IS 'campo en el que se puede registrar una descripcion del periodo contable';
COMMENT ON COLUMN financiera.periodo_contable.mes IS 'campo en el que se registra el mes como el periodo';
COMMENT ON COLUMN financiera.periodo_contable.ano IS 'campo que indica el año para el que se define el periodo';
COMMENT ON CONSTRAINT ano_mes_uq ON financiera.periodo_contable IS 'restriccion para controlar un unico registro de un periodo por año';
ALTER TABLE financiera.periodo_contable OWNER TO postgres;
-- ddl-end --

-- object: financiera.periodo_plan_cuentas_id_seq | type: SEQUENCE --
-- DROP SEQUENCE financiera.periodo_plan_cuentas_id_seq;
CREATE SEQUENCE financiera.periodo_plan_cuentas_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
ALTER SEQUENCE financiera.periodo_plan_cuentas_id_seq OWNER TO postgres;
-- ddl-end --

-- object: financiera.periodo_plan_cuentas | type: TABLE --
-- DROP TABLE financiera.periodo_plan_cuentas;
CREATE  TABLE financiera.periodo_plan_cuentas(
	id integer NOT NULL DEFAULT nextval('periodo_plan_cuentas_id_seq'::regclass),
	definitivo boolean NOT NULL,
	plan_cuentas integer NOT NULL,
	periodo_contable integer NOT NULL,
	CONSTRAINT periodo_plan_cuentas_pk PRIMARY KEY (id),
	CONSTRAINT periodo_plan_cuentas_uq UNIQUE (plan_cuentas,periodo_contable)

);
-- ddl-end --
COMMENT ON TABLE financiera.periodo_plan_cuentas IS 'tabla en la que se estipulan los periodos de un plan de cuentas';
COMMENT ON COLUMN financiera.periodo_plan_cuentas.id IS 'identificador de la tabla periodo_plan_cuentas';
COMMENT ON COLUMN financiera.periodo_plan_cuentas.definitivo IS 'campo que indica si un plan de cuentas es definitivo en el periodo referenciado';
COMMENT ON COLUMN financiera.periodo_plan_cuentas.plan_cuentas IS 'identificador de la tabla plan_cuentas referenciado para indicar a que periodo se asigna el plan de cuentas';
COMMENT ON COLUMN financiera.periodo_plan_cuentas.periodo_contable IS 'identificador de la tabla periodo_contable referenciado para aclarar el periodo del plan de cuentas';
ALTER TABLE financiera.periodo_plan_cuentas OWNER TO postgres;
-- ddl-end --

-- object: financiera.concepto_cuenta_contable_id_seq | type: SEQUENCE --
-- DROP SEQUENCE financiera.concepto_cuenta_contable_id_seq;
CREATE SEQUENCE financiera.concepto_cuenta_contable_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
ALTER SEQUENCE financiera.concepto_cuenta_contable_id_seq OWNER TO postgres;
-- ddl-end --

-- object: financiera.concepto_cuenta_contable | type: TABLE --
-- DROP TABLE financiera.concepto_cuenta_contable;
CREATE  TABLE financiera.concepto_cuenta_contable(
	id integer NOT NULL DEFAULT nextval('concepto_cuenta_contable_id_seq'::regclass),
	cuenta_contable integer NOT NULL,
	concepto integer NOT NULL,
	CONSTRAINT concepto_cuenta_contable_pk PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON COLUMN financiera.concepto_cuenta_contable.id IS 'identificador de la tabla concepto_cuenta_contable';
COMMENT ON COLUMN financiera.concepto_cuenta_contable.cuenta_contable IS 'identificador de la tabla cuenta_contable referenciada para la relacion con el concepto';
COMMENT ON COLUMN financiera.concepto_cuenta_contable.concepto IS 'identificador de la tabla concepto referenciado para la relacion con cuenta contable';
ALTER TABLE financiera.concepto_cuenta_contable OWNER TO postgres;
-- ddl-end --

-- object: financiera.cuenta_bancaria_id_seq | type: SEQUENCE --
-- DROP SEQUENCE financiera.cuenta_bancaria_id_seq;
CREATE SEQUENCE financiera.cuenta_bancaria_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
ALTER SEQUENCE financiera.cuenta_bancaria_id_seq OWNER TO postgres;
-- ddl-end --

-- object: financiera.cuenta_bancaria | type: TABLE --
-- DROP TABLE financiera.cuenta_bancaria;
CREATE  TABLE financiera.cuenta_bancaria(
	id integer NOT NULL DEFAULT nextval('cuenta_bancaria_id_seq'::regclass),
	nombre character varying NOT NULL,
	numero_cuenta character varying(25) NOT NULL,
	estado_activo boolean NOT NULL DEFAULT true,
	saldo numeric(15,3) NOT NULL,
	tipo_cuenta integer NOT NULL,
	tipo_recurso integer NOT NULL,
	sucursal integer NOT NULL,
	unidad_ejecutora integer NOT NULL,
	CONSTRAINT cuenta_bancaria_pk PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE financiera.cuenta_bancaria IS 'tabla en la que se registran las diferentes cuentas bancarias asociadas a la universidad';
COMMENT ON COLUMN financiera.cuenta_bancaria.id IS 'identificador de la tabla cuenta_bancaria';
COMMENT ON COLUMN financiera.cuenta_bancaria.nombre IS 'campo en el que se indica el nombre de la cuenta bancaria';
COMMENT ON COLUMN financiera.cuenta_bancaria.numero_cuenta IS 'campo en el que se registra el numero de la cuenta';
COMMENT ON COLUMN financiera.cuenta_bancaria.estado_activo IS 'campo que indica si la cuenta se encuentra activa o no';
COMMENT ON COLUMN financiera.cuenta_bancaria.saldo IS 'campo en el que se registra el saldo de una cuenta';
COMMENT ON COLUMN financiera.cuenta_bancaria.tipo_cuenta IS 'identificador de la tabla tipo_cuenta referenciado para distinguir el tipo de cuenta';
COMMENT ON COLUMN financiera.cuenta_bancaria.tipo_recurso IS 'identificador de la tabla recurso referenciado para diferenciar los recursos de una cuenta bancaria';
COMMENT ON COLUMN financiera.cuenta_bancaria.sucursal IS 'identificador de la tabla sucursal referenciado para indicar a que sucursal pertenece la cuenta';
COMMENT ON COLUMN financiera.cuenta_bancaria.unidad_ejecutora IS 'identificador de la tabla unidad_ejecutora referenciado para indicar que esta es la que maneja la cuenta a traves de la entidad.';
ALTER TABLE financiera.cuenta_bancaria OWNER TO postgres;
-- ddl-end --

-- object: financiera.tipo_cuenta_bancaria_id_seq | type: SEQUENCE --
-- DROP SEQUENCE financiera.tipo_cuenta_bancaria_id_seq;
CREATE SEQUENCE financiera.tipo_cuenta_bancaria_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
ALTER SEQUENCE financiera.tipo_cuenta_bancaria_id_seq OWNER TO postgres;
-- ddl-end --

-- object: financiera.tipo_cuenta_bancaria | type: TABLE --
-- DROP TABLE financiera.tipo_cuenta_bancaria;
CREATE  TABLE financiera.tipo_cuenta_bancaria(
	id integer NOT NULL DEFAULT nextval('tipo_cuenta_bancaria_id_seq'::regclass),
	nombre character varying(80) NOT NULL,
	descripcion character varying,
	CONSTRAINT tipo_cuenta_pk PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE financiera.tipo_cuenta_bancaria IS 'tabla en la que se indican los tipos de cuentas bancarias existentes';
COMMENT ON COLUMN financiera.tipo_cuenta_bancaria.id IS 'identificador de la tabla tipo_cuenta';
COMMENT ON COLUMN financiera.tipo_cuenta_bancaria.nombre IS 'campo en el que se registra el nombre del tipo de cuenta';
COMMENT ON COLUMN financiera.tipo_cuenta_bancaria.descripcion IS 'campo en el que se puede registrar una descripcion del tipo de cuenta';
ALTER TABLE financiera.tipo_cuenta_bancaria OWNER TO postgres;
-- ddl-end --

-- object: financiera.tipo_recurso_id_seq | type: SEQUENCE --
-- DROP SEQUENCE financiera.tipo_recurso_id_seq;
CREATE SEQUENCE financiera.tipo_recurso_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
ALTER SEQUENCE financiera.tipo_recurso_id_seq OWNER TO postgres;
-- ddl-end --

-- object: financiera.tipo_recurso | type: TABLE --
-- DROP TABLE financiera.tipo_recurso;
CREATE  TABLE financiera.tipo_recurso(
	id integer NOT NULL DEFAULT nextval('tipo_recurso_id_seq'::regclass),
	nombre character varying(80) NOT NULL,
	descripcion character varying,
	CONSTRAINT recurso_pk PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE financiera.tipo_recurso IS 'tabla en la que se registran los recursos de cuentas bancarias';
COMMENT ON COLUMN financiera.tipo_recurso.id IS 'identificador de la tabla tipo_recurso';
COMMENT ON COLUMN financiera.tipo_recurso.nombre IS 'campo en el que se registra el nombre de un recurso';
COMMENT ON COLUMN financiera.tipo_recurso.descripcion IS 'campo en el que se puede registrar una descripcion asociada a un recurso';
ALTER TABLE financiera.tipo_recurso OWNER TO postgres;
-- ddl-end --

-- object: financiera.entidad_id_seq | type: SEQUENCE --
-- DROP SEQUENCE financiera.entidad_id_seq;
CREATE SEQUENCE financiera.entidad_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
ALTER SEQUENCE financiera.entidad_id_seq OWNER TO postgres;
-- ddl-end --

-- object: financiera.entidad | type: TABLE --
-- DROP TABLE financiera.entidad;
CREATE  TABLE financiera.entidad(
	id integer NOT NULL DEFAULT nextval('entidad_id_seq'::regclass),
	nombre character varying(60) NOT NULL,
	codigo_entidad character varying(6) NOT NULL,
	tipo_entidad integer NOT NULL,
	CONSTRAINT entidad_pk PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE financiera.entidad IS 'tabla en la que se registran las entidades referentes al esquema';
COMMENT ON COLUMN financiera.entidad.id IS 'identificador de la tabla entidad';
COMMENT ON COLUMN financiera.entidad.nombre IS 'campo en el que se indica el nombre de la entidad';
COMMENT ON COLUMN financiera.entidad.codigo_entidad IS 'campo en el que se asigna un codigo a la entidad';
COMMENT ON COLUMN financiera.entidad.tipo_entidad IS 'identificador de la tabla tipo_entidad relacionada para diferenciar el tipo de entidad que se registra';
ALTER TABLE financiera.entidad OWNER TO postgres;
-- ddl-end --

-- object: financiera.unidad_ejecutora_id_seq | type: SEQUENCE --
-- DROP SEQUENCE financiera.unidad_ejecutora_id_seq;
CREATE SEQUENCE financiera.unidad_ejecutora_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
ALTER SEQUENCE financiera.unidad_ejecutora_id_seq OWNER TO postgres;
-- ddl-end --

-- object: financiera.unidad_ejecutora | type: TABLE --
-- DROP TABLE financiera.unidad_ejecutora;
CREATE  TABLE financiera.unidad_ejecutora(
	id integer NOT NULL DEFAULT nextval('unidad_ejecutora_id_seq'::regclass),
	nombre character varying(300) NOT NULL,
	descripcion character varying(500),
	entidad integer NOT NULL,
	CONSTRAINT unidad_ejecutora_pk PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON COLUMN financiera.unidad_ejecutora.entidad IS 'identificador de la tabla entidad que se referncia para saber a que entidad corresponde la unidad ejecutora';
ALTER TABLE financiera.unidad_ejecutora OWNER TO postgres;
-- ddl-end --

-- object: financiera.tipo_entidad_id_seq | type: SEQUENCE --
-- DROP SEQUENCE financiera.tipo_entidad_id_seq;
CREATE SEQUENCE financiera.tipo_entidad_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
ALTER SEQUENCE financiera.tipo_entidad_id_seq OWNER TO postgres;
-- ddl-end --

-- object: financiera.tipo_entidad | type: TABLE --
-- DROP TABLE financiera.tipo_entidad;
CREATE  TABLE financiera.tipo_entidad(
	id integer NOT NULL DEFAULT nextval('tipo_entidad_id_seq'::regclass),
	nombre character varying(80) NOT NULL,
	descripcion character varying,
	CONSTRAINT tipo_entidad_pk PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE financiera.tipo_entidad IS 'tabla en la que se registran los tipos de entidades en el esquema kronos';
COMMENT ON COLUMN financiera.tipo_entidad.id IS 'identificador de la tabla tipo_entidad';
COMMENT ON COLUMN financiera.tipo_entidad.nombre IS 'campo en el que se indica el nombre de el tipo entidad';
COMMENT ON COLUMN financiera.tipo_entidad.descripcion IS 'campo en el que se puede registrar una descripcion al tipo entidad';
ALTER TABLE financiera.tipo_entidad OWNER TO postgres;
-- ddl-end --

-- object: financiera.tipo_compromiso_tesoral_id_seq | type: SEQUENCE --
-- DROP SEQUENCE financiera.tipo_compromiso_tesoral_id_seq;
CREATE SEQUENCE financiera.tipo_compromiso_tesoral_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
ALTER SEQUENCE financiera.tipo_compromiso_tesoral_id_seq OWNER TO postgres;
-- ddl-end --

-- object: financiera.tipo_compromiso_tesoral | type: TABLE --
-- DROP TABLE financiera.tipo_compromiso_tesoral;
CREATE  TABLE financiera.tipo_compromiso_tesoral(
	id integer NOT NULL DEFAULT nextval('tipo_compromiso_tesoral_id_seq'::regclass),
	nombre character varying(150) NOT NULL,
	estado_activo boolean NOT NULL,
	categoria_compromiso integer NOT NULL,
	descripcion character varying,
	CONSTRAINT tipo_compromiso_tesoral_pk PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE financiera.tipo_compromiso_tesoral IS 'tabla en la que se registran los tipos de compromisos tesorales ';
COMMENT ON COLUMN financiera.tipo_compromiso_tesoral.id IS 'identificador de la tabla tipo_compromiso_tesoral';
COMMENT ON COLUMN financiera.tipo_compromiso_tesoral.nombre IS 'campo que indica el nombre dado al tipo de comrpomiso';
COMMENT ON COLUMN financiera.tipo_compromiso_tesoral.estado_activo IS 'campo en el que se indica si el tipo de compromiso se encuentra activo';
COMMENT ON COLUMN financiera.tipo_compromiso_tesoral.categoria_compromiso IS 'identificador de la tabla categoria referenciado para diferenciar los tipos de compromisos';
COMMENT ON COLUMN financiera.tipo_compromiso_tesoral.descripcion IS 'campo en el que se puede registrar una descripcion del tipo compromiso si se requiere';
ALTER TABLE financiera.tipo_compromiso_tesoral OWNER TO postgres;
-- ddl-end --

-- object: financiera.compromiso_id_seq | type: SEQUENCE --
-- DROP SEQUENCE financiera.compromiso_id_seq;
CREATE SEQUENCE financiera.compromiso_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
ALTER SEQUENCE financiera.compromiso_id_seq OWNER TO postgres;
-- ddl-end --

-- object: financiera.compromiso | type: TABLE --
-- DROP TABLE financiera.compromiso;
CREATE  TABLE financiera.compromiso(
	id integer NOT NULL DEFAULT nextval('compromiso_id_seq'::regclass),
	objeto character varying NOT NULL,
	vigencia numeric(4,0) NOT NULL,
	fecha_inicio date NOT NULL,
	fecha_fin date NOT NULL,
	fecha_modificacion date NOT NULL,
	estado_compromiso integer NOT NULL,
	tipo_compromiso_tesoral integer NOT NULL,
	unidad_ejecutora integer NOT NULL,
	CONSTRAINT compromiso_pk PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE financiera.compromiso IS 'tabla en la que se registran los compromisos financieros para un periodo determinado';
COMMENT ON COLUMN financiera.compromiso.id IS 'identificador de la tabla compromiso';
COMMENT ON COLUMN financiera.compromiso.objeto IS 'campo en el que se registra el objeto determinado en el compromiso';
COMMENT ON COLUMN financiera.compromiso.vigencia IS 'campo en el que se indica el año de vigencia de un compromiso';
COMMENT ON COLUMN financiera.compromiso.fecha_inicio IS 'campo que indica el momento en el que iniciara el compromiso.';
COMMENT ON COLUMN financiera.compromiso.fecha_fin IS 'campo que indica la fecha en la que el compromiso finaliza';
COMMENT ON COLUMN financiera.compromiso.fecha_modificacion IS 'campo que indica la fecha en la que se creo el compromiso y si se actualiza la ultima fecha de  modificacion';
COMMENT ON COLUMN financiera.compromiso.estado_compromiso IS 'campo que indica en que estado se encuentra el compromiso a la fecha';
COMMENT ON COLUMN financiera.compromiso.tipo_compromiso_tesoral IS 'identificador de la tabla tipo_compromiso_tesoral referenciado para situar un compromiso en un tipo';
COMMENT ON COLUMN financiera.compromiso.unidad_ejecutora IS 'identificador de la tabla unidad_ejecutora referenciado para indicar el lugar en el que se aplica el compromiso';
ALTER TABLE financiera.compromiso OWNER TO postgres;
-- ddl-end --

-- object: financiera.estado_compromiso_id_seq | type: SEQUENCE --
-- DROP SEQUENCE financiera.estado_compromiso_id_seq;
CREATE SEQUENCE financiera.estado_compromiso_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
ALTER SEQUENCE financiera.estado_compromiso_id_seq OWNER TO postgres;
-- ddl-end --

-- object: financiera.estado_compromiso | type: TABLE --
-- DROP TABLE financiera.estado_compromiso;
CREATE  TABLE financiera.estado_compromiso(
	id integer NOT NULL DEFAULT nextval('estado_compromiso_id_seq'::regclass),
	nombre character varying(80) NOT NULL,
	descripcion character varying,
	CONSTRAINT estado_compromiso_pk PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE financiera.estado_compromiso IS 'tabla en la que se registran los diferentes estados en los que se encuentra un compromiso';
COMMENT ON COLUMN financiera.estado_compromiso.id IS 'identificador de la tabla estado_compromiso';
COMMENT ON COLUMN financiera.estado_compromiso.nombre IS 'campo que indica el nombre del estado para un compromiso';
COMMENT ON COLUMN financiera.estado_compromiso.descripcion IS 'campo en el que se puede registrar una descripcion al estado del compromiso';
ALTER TABLE financiera.estado_compromiso OWNER TO postgres;
-- ddl-end --

-- object: financiera.categoria_compromiso_id_seq | type: SEQUENCE --
-- DROP SEQUENCE financiera.categoria_compromiso_id_seq;
CREATE SEQUENCE financiera.categoria_compromiso_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
ALTER SEQUENCE financiera.categoria_compromiso_id_seq OWNER TO postgres;
-- ddl-end --

-- object: financiera.categoria_compromiso | type: TABLE --
-- DROP TABLE financiera.categoria_compromiso;
CREATE  TABLE financiera.categoria_compromiso(
	id integer NOT NULL DEFAULT nextval('categoria_compromiso_id_seq'::regclass),
	nombre character varying(80) NOT NULL,
	descripcion character varying,
	CONSTRAINT categoria_compromiso_pk PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE financiera.categoria_compromiso IS 'tabla encargada de diferenciar las categorias en las que se encuentran los diferentes tipos de compromisos';
COMMENT ON COLUMN financiera.categoria_compromiso.id IS 'identificador de la tabla categoria_compromiso';
COMMENT ON COLUMN financiera.categoria_compromiso.nombre IS 'campo en el que se indica el nombre de la categoria';
COMMENT ON COLUMN financiera.categoria_compromiso.descripcion IS 'campo en el que se puede agregar una descripcion a la categoria';
ALTER TABLE financiera.categoria_compromiso OWNER TO postgres;
-- ddl-end --

-- object: financiera.version_tipo_transaccion_id_seq | type: SEQUENCE --
-- DROP SEQUENCE financiera.version_tipo_transaccion_id_seq;
CREATE SEQUENCE financiera.version_tipo_transaccion_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
ALTER SEQUENCE financiera.version_tipo_transaccion_id_seq OWNER TO postgres;
-- ddl-end --

-- object: financiera.version_tipo_transaccion | type: TABLE --
-- DROP TABLE financiera.version_tipo_transaccion;
CREATE  TABLE financiera.version_tipo_transaccion(
	id integer NOT NULL DEFAULT nextval('version_tipo_transaccion_id_seq'::regclass),
	nombre character varying(80) NOT NULL,
	definitiva boolean NOT NULL,
	fecha_modificacion date NOT NULL,
	fecha_inicio date NOT NULL,
	fecha_fin date,
	tipo_transaccion integer NOT NULL,
	CONSTRAINT version_tipo_transaccion_pk PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE financiera.version_tipo_transaccion IS 'tabla en la que se registran las versiones para el control de transacciones de un tipo de transaccion';
COMMENT ON COLUMN financiera.version_tipo_transaccion.id IS 'identificador de la tabla version_tipo_transaccion';
COMMENT ON COLUMN financiera.version_tipo_transaccion.nombre IS 'campo en el que se indica el nombre a una version del tipo transaccion';
COMMENT ON COLUMN financiera.version_tipo_transaccion.definitiva IS 'campo que indica si la version registrada es definitiva y por lo tanto es la que aplica para el manejo de transacciones';
COMMENT ON COLUMN financiera.version_tipo_transaccion.fecha_modificacion IS 'campo que registra la fecha de creacion o de la ultima modificacion que tenga la version registrada';
COMMENT ON COLUMN financiera.version_tipo_transaccion.fecha_inicio IS 'campo en el que se indica la fecha en que ese tipo de transaccion rige con esta version';
COMMENT ON COLUMN financiera.version_tipo_transaccion.fecha_fin IS 'campo en el que se puede registrar la fecha de finalizacion de la version para el tipo de transaccion referenciado';
COMMENT ON COLUMN financiera.version_tipo_transaccion.tipo_transaccion IS 'identificador de la tabla tipo_transaccion referenciada para comprobar el tipo al que pertenece la version registrada';
ALTER TABLE financiera.version_tipo_transaccion OWNER TO postgres;
-- ddl-end --

-- object: financiera.tipo_transaccion_id_seq | type: SEQUENCE --
-- DROP SEQUENCE financiera.tipo_transaccion_id_seq;
CREATE SEQUENCE financiera.tipo_transaccion_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
ALTER SEQUENCE financiera.tipo_transaccion_id_seq OWNER TO postgres;
-- ddl-end --

-- object: financiera.tipo_transaccion | type: TABLE --
-- DROP TABLE financiera.tipo_transaccion;
CREATE  TABLE financiera.tipo_transaccion(
	id integer NOT NULL DEFAULT nextval('tipo_transaccion_id_seq'::regclass),
	nombre character varying(150) NOT NULL,
	descripcion character varying,
	clase_transaccion integer NOT NULL,
	CONSTRAINT tipo_transaccion_pk PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE financiera.tipo_transaccion IS 'tabla en la que se registran los tipos de transacciones ubicados en una clase de transaccion';
COMMENT ON COLUMN financiera.tipo_transaccion.id IS 'identificador de la tabla tipo_transaccion';
COMMENT ON COLUMN financiera.tipo_transaccion.nombre IS 'campo en el que se indica el nombre del tipo de transaccion';
COMMENT ON COLUMN financiera.tipo_transaccion.descripcion IS 'campo en el que se puede indicar una descripcion para el tipo de transaccion';
COMMENT ON COLUMN financiera.tipo_transaccion.clase_transaccion IS 'identificador de la clase transaccion referenciada para indicar a que clase pertenece el tipo de transaccion registrado';
ALTER TABLE financiera.tipo_transaccion OWNER TO postgres;
-- ddl-end --

-- object: financiera.clase_transaccion_id_seq | type: SEQUENCE --
-- DROP SEQUENCE financiera.clase_transaccion_id_seq;
CREATE SEQUENCE financiera.clase_transaccion_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
ALTER SEQUENCE financiera.clase_transaccion_id_seq OWNER TO postgres;
-- ddl-end --

-- object: financiera.clase_transaccion | type: TABLE --
-- DROP TABLE financiera.clase_transaccion;
CREATE  TABLE financiera.clase_transaccion(
	id integer NOT NULL DEFAULT nextval('clase_transaccion_id_seq'::regclass),
	nombre character varying(80) NOT NULL,
	descripcion character varying,
	modulo_kronos integer NOT NULL,
	CONSTRAINT clase_transaccion_pk PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE financiera.clase_transaccion IS 'tabla en la que se registran las clases de transacciones disponibles para cada modulo';
COMMENT ON COLUMN financiera.clase_transaccion.id IS 'identificador de la tabla clase_transaccion';
COMMENT ON COLUMN financiera.clase_transaccion.nombre IS 'campo en el que se indica el nombre de la clase de transaccion';
COMMENT ON COLUMN financiera.clase_transaccion.descripcion IS 'campo en el que se puede registrar una descriopcion concerniete a la clase de transaccion';
COMMENT ON COLUMN financiera.clase_transaccion.modulo_kronos IS 'identificador de la tabla modulo-aplicacion referenciada para indicar a que modulo pertenece la clase de transaccion registrada';
ALTER TABLE financiera.clase_transaccion OWNER TO postgres;
-- ddl-end --

-- object: financiera.modulo_kronos_id_seq | type: SEQUENCE --
-- DROP SEQUENCE financiera.modulo_kronos_id_seq;
CREATE SEQUENCE financiera.modulo_kronos_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
ALTER SEQUENCE financiera.modulo_kronos_id_seq OWNER TO postgres;
-- ddl-end --

-- object: financiera.modulo_kronos | type: TABLE --
-- DROP TABLE financiera.modulo_kronos;
CREATE  TABLE financiera.modulo_kronos(
	id integer NOT NULL DEFAULT nextval('modulo_kronos_id_seq'::regclass),
	nombre character varying(80) NOT NULL,
	descripcion character varying NOT NULL,
	CONSTRAINT modulo_aplicacion_pk PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE financiera.modulo_kronos IS 'tabla en la que son registrados los modulos de la aplicacion de kronos con el fin de gestionar las transacciones que se pueden realizar en estos';
COMMENT ON COLUMN financiera.modulo_kronos.id IS 'identificador de la tabla modulo_kronos';
COMMENT ON COLUMN financiera.modulo_kronos.nombre IS 'campo en el que se registra el nombre del modulo de la aplicacion';
COMMENT ON COLUMN financiera.modulo_kronos.descripcion IS 'campo en el que se puede indicar una descripcion al modulo registrado';
ALTER TABLE financiera.modulo_kronos OWNER TO postgres;
-- ddl-end --

-- object: financiera.afectacion_concepto_id_seq | type: SEQUENCE --
-- DROP SEQUENCE financiera.afectacion_concepto_id_seq;
CREATE SEQUENCE financiera.afectacion_concepto_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
ALTER SEQUENCE financiera.afectacion_concepto_id_seq OWNER TO postgres;
-- ddl-end --

-- object: financiera.afectacion_concepto | type: TABLE --
-- DROP TABLE financiera.afectacion_concepto;
CREATE  TABLE financiera.afectacion_concepto(
	id integer NOT NULL DEFAULT nextval('afectacion_concepto_id_seq'::regclass),
	afectacion_ingreso boolean NOT NULL,
	afectacion_egreso boolean NOT NULL,
	concepto integer NOT NULL,
	tipo_afectacion integer NOT NULL,
	CONSTRAINT afectacion_concepto_pk PRIMARY KEY (id),
	CONSTRAINT tipo_afectacion_concepto_uk UNIQUE (concepto,tipo_afectacion)

);
-- ddl-end --
COMMENT ON TABLE financiera.afectacion_concepto IS 'tabla en la que se relacionan los conceptos y el tipo de afectacion, en esta se indica si afecta en egresos o ingresos';
COMMENT ON COLUMN financiera.afectacion_concepto.id IS 'identificador de la tabla afectacion concepto';
COMMENT ON COLUMN financiera.afectacion_concepto.afectacion_ingreso IS 'campo que indica si un concepto afecta el ingreso de un tipo de afectacion';
COMMENT ON COLUMN financiera.afectacion_concepto.afectacion_egreso IS 'campo que indica si un concepto afecta el egreso de un tipo de afectacion';
COMMENT ON COLUMN financiera.afectacion_concepto.concepto IS 'identificador de la tabla concepto que se relaciona spara la afectacion sobre un tipo de afectacion';
COMMENT ON COLUMN financiera.afectacion_concepto.tipo_afectacion IS 'identificador de la tabla tipo_afectacion que se relaciona para indicar a que tipo de afectacion afecta un concepto ';
COMMENT ON CONSTRAINT afectacion_concepto_pk ON financiera.afectacion_concepto IS 'constraint primary key tabla afectacion_concepto';
ALTER TABLE financiera.afectacion_concepto OWNER TO postgres;
-- ddl-end --

-- object: financiera.tipo_afectacion_id_seq | type: SEQUENCE --
-- DROP SEQUENCE financiera.tipo_afectacion_id_seq;
CREATE SEQUENCE financiera.tipo_afectacion_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
ALTER SEQUENCE financiera.tipo_afectacion_id_seq OWNER TO postgres;
-- ddl-end --

-- object: financiera.tipo_afectacion | type: TABLE --
-- DROP TABLE financiera.tipo_afectacion;
CREATE  TABLE financiera.tipo_afectacion(
	id integer NOT NULL DEFAULT nextval('tipo_afectacion_id_seq'::regclass),
	nombre character varying(80) NOT NULL,
	descripcion character varying,
	CONSTRAINT tipo_afectacion_pk PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE financiera.tipo_afectacion IS 'tabla que indica los tipo de afectacion sobre un concepto';
COMMENT ON COLUMN financiera.tipo_afectacion.id IS 'identificador de la tabla tipo_afectacion';
COMMENT ON COLUMN financiera.tipo_afectacion.nombre IS 'campo que identifica el nombre del tipo de afectacion';
COMMENT ON COLUMN financiera.tipo_afectacion.descripcion IS 'campo que indica la descripcion de un tipo de afectacion';
COMMENT ON CONSTRAINT tipo_afectacion_pk ON financiera.tipo_afectacion IS 'constraint primary key tabla tipo_afectacion';
ALTER TABLE financiera.tipo_afectacion OWNER TO postgres;
-- ddl-end --

-- object: financiera.tipo_concepto_id_seq | type: SEQUENCE --
-- DROP SEQUENCE financiera.tipo_concepto_id_seq;
CREATE SEQUENCE financiera.tipo_concepto_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
ALTER SEQUENCE financiera.tipo_concepto_id_seq OWNER TO postgres;
-- ddl-end --

-- object: financiera.tipo_concepto | type: TABLE --
-- DROP TABLE financiera.tipo_concepto;
CREATE  TABLE financiera.tipo_concepto(
	id integer NOT NULL DEFAULT nextval('tipo_concepto_id_seq'::regclass),
	nombre character varying(80) NOT NULL,
	descripcion character varying,
	CONSTRAINT tipo_concepto_pk PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE financiera.tipo_concepto IS 'tabla en la que se registran los tipos de concepto que existen en el sistema';
COMMENT ON COLUMN financiera.tipo_concepto.id IS 'identificador de la tabla tipo_concepto';
COMMENT ON COLUMN financiera.tipo_concepto.nombre IS 'campo que indica el nombre del tipo de concepto';
COMMENT ON COLUMN financiera.tipo_concepto.descripcion IS 'campo en el que se puede registrar una descripcion relacionada al tipo de concepto';
ALTER TABLE financiera.tipo_concepto OWNER TO postgres;
-- ddl-end --

-- object: financiera.concepto_concepto_id_seq | type: SEQUENCE --
-- DROP SEQUENCE financiera.concepto_concepto_id_seq;
CREATE SEQUENCE financiera.concepto_concepto_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
ALTER SEQUENCE financiera.concepto_concepto_id_seq OWNER TO postgres;
-- ddl-end --

-- object: financiera.concepto_concepto | type: TABLE --
-- DROP TABLE financiera.concepto_concepto;
CREATE  TABLE financiera.concepto_concepto(
	id integer NOT NULL DEFAULT nextval('concepto_concepto_id_seq'::regclass),
	concepto_padre integer NOT NULL,
	concepto_hijo integer NOT NULL,
	CONSTRAINT concepto_concepto_pk PRIMARY KEY (id),
	CONSTRAINT concepto_hijo_uk UNIQUE (concepto_hijo)

);
-- ddl-end --
COMMENT ON TABLE financiera.concepto_concepto IS 'tabla en la que se realiza la relacion de un concepto hijo y un concepto padre';
COMMENT ON COLUMN financiera.concepto_concepto.id IS 'identificador de la tabla concepto_concepto';
COMMENT ON COLUMN financiera.concepto_concepto.concepto_padre IS 'identificador de la tabla concepto que relaciona al concepto padre';
COMMENT ON COLUMN financiera.concepto_concepto.concepto_hijo IS 'identificador de la tabla concepto que relaciona al concepto hijo';
ALTER TABLE financiera.concepto_concepto OWNER TO postgres;
-- ddl-end --

-- object: financiera.concepto_id_seq1 | type: SEQUENCE --
-- DROP SEQUENCE financiera.concepto_id_seq1;
CREATE SEQUENCE financiera.concepto_id_seq1
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
ALTER SEQUENCE financiera.concepto_id_seq1 OWNER TO postgres;
-- ddl-end --

-- object: financiera.concepto | type: TABLE --
-- DROP TABLE financiera.concepto;
CREATE  TABLE financiera.concepto(
	id integer NOT NULL DEFAULT nextval('concepto_id_seq1'::regclass),
	codigo character varying(25) NOT NULL,
	nombre character varying NOT NULL,
	fecha_creacion date NOT NULL,
	cabeza boolean NOT NULL DEFAULT false,
	fecha_expiracion date,
	descripcion character varying,
	tipo_concepto integer NOT NULL,
	rubro integer,
	CONSTRAINT concepto_pk PRIMARY KEY (id),
	CONSTRAINT codigo_uk UNIQUE (codigo)

);
-- ddl-end --
COMMENT ON TABLE financiera.concepto IS 'tabla que soporta el manejo de conceptos tanto de ingresos como de gastos';
COMMENT ON COLUMN financiera.concepto.id IS 'identificador de la tabla concepto';
COMMENT ON COLUMN financiera.concepto.codigo IS 'campo en el que se registra el codigo unico para el concepto';
COMMENT ON COLUMN financiera.concepto.nombre IS 'campo que indica el nombre referido a un concepto';
COMMENT ON COLUMN financiera.concepto.fecha_creacion IS 'campo que indica la fecha en que el concepto es creado';
COMMENT ON COLUMN financiera.concepto.cabeza IS 'campo que indica si un concepto es cabeza, en caso que no lo sea sera hoja.';
COMMENT ON COLUMN financiera.concepto.fecha_expiracion IS 'campo que indica la fecha de expiracion de un concepto.';
COMMENT ON COLUMN financiera.concepto.descripcion IS 'campo que provee la descripcion de un concepto';
COMMENT ON COLUMN financiera.concepto.tipo_concepto IS 'campo que indica a que tipo pertenece el concepto, ingreso, egreso o gasto';
COMMENT ON COLUMN financiera.concepto.rubro IS 'identificador de la tabla rubro que se relaciona para los conceptos de tipo gasto';
COMMENT ON CONSTRAINT concepto_pk ON financiera.concepto IS 'constraint primary key tabla concepto';
COMMENT ON CONSTRAINT codigo_uk ON financiera.concepto IS 'restriccion que permite que el codigo de cada concepto sea unico';
ALTER TABLE financiera.concepto OWNER TO postgres;
-- ddl-end --

-- object: financiera.ingreso_id_seq1 | type: SEQUENCE --
-- DROP SEQUENCE financiera.ingreso_id_seq1;
CREATE SEQUENCE financiera.ingreso_id_seq1
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
ALTER SEQUENCE financiera.ingreso_id_seq1 OWNER TO postgres;
-- ddl-end --

-- object: financiera.ingreso | type: TABLE --
-- DROP TABLE financiera.ingreso;
CREATE  TABLE financiera.ingreso(
	id integer NOT NULL DEFAULT nextval('ingreso_id_seq1'::regclass),
	consecutivo numeric NOT NULL,
	vigencia numeric(4,0) NOT NULL,
	fecha_ingreso date NOT NULL,
	fecha_consignacion date NOT NULL,
	valor numeric NOT NULL,
	observaciones character varying,
	origen_ingreso character varying,
	forma_ingreso integer NOT NULL,
	estado_ingreso integer NOT NULL,
	unidad_ejecutora integer NOT NULL,
	aportante integer,
	reviso integer,
	elaboro integer NOT NULL,
	CONSTRAINT ingreso_pk PRIMARY KEY (id),
	CONSTRAINT consecutivo_vigencia_uq UNIQUE (consecutivo,vigencia)

);
-- ddl-end --
ALTER TABLE financiera.ingreso OWNER TO postgres;
-- ddl-end --

-- object: financiera.forma_ingreso_id_seq | type: SEQUENCE --
-- DROP SEQUENCE financiera.forma_ingreso_id_seq;
CREATE SEQUENCE financiera.forma_ingreso_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
ALTER SEQUENCE financiera.forma_ingreso_id_seq OWNER TO postgres;
-- ddl-end --

-- object: financiera.forma_ingreso | type: TABLE --
-- DROP TABLE financiera.forma_ingreso;
CREATE  TABLE financiera.forma_ingreso(
	id integer NOT NULL DEFAULT nextval('forma_ingreso_id_seq'::regclass),
	nombre character varying(80) NOT NULL,
	descripcion character varying,
	CONSTRAINT forma_ingreso_pk PRIMARY KEY (id)

);
-- ddl-end --
ALTER TABLE financiera.forma_ingreso OWNER TO postgres;
-- ddl-end --

-- object: financiera.estado_ingreso_id_seq | type: SEQUENCE --
-- DROP SEQUENCE financiera.estado_ingreso_id_seq;
CREATE SEQUENCE financiera.estado_ingreso_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
ALTER SEQUENCE financiera.estado_ingreso_id_seq OWNER TO postgres;
-- ddl-end --

-- object: financiera.estado_ingreso | type: TABLE --
-- DROP TABLE financiera.estado_ingreso;
CREATE  TABLE financiera.estado_ingreso(
	id integer NOT NULL DEFAULT nextval('estado_ingreso_id_seq'::regclass),
	nombre character varying(80) NOT NULL,
	descripcion character varying,
	CONSTRAINT estado_ingreso_pk PRIMARY KEY (id)

);
-- ddl-end --
ALTER TABLE financiera.estado_ingreso OWNER TO postgres;
-- ddl-end --

-- object: financiera.ingreso_concepto_id_seq | type: SEQUENCE --
-- DROP SEQUENCE financiera.ingreso_concepto_id_seq;
CREATE SEQUENCE financiera.ingreso_concepto_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
ALTER SEQUENCE financiera.ingreso_concepto_id_seq OWNER TO postgres;
-- ddl-end --

-- object: financiera.ingreso_concepto | type: TABLE --
-- DROP TABLE financiera.ingreso_concepto;
CREATE  TABLE financiera.ingreso_concepto(
	id integer NOT NULL DEFAULT nextval('ingreso_concepto_id_seq'::regclass),
	valor_agregado numeric NOT NULL,
	ingreso integer NOT NULL,
	concepto integer NOT NULL,
	CONSTRAINT ingreso_concepto_pk PRIMARY KEY (id),
	CONSTRAINT concepto_ingreso_uq UNIQUE (concepto,ingreso)

);
-- ddl-end --
COMMENT ON COLUMN financiera.ingreso_concepto.concepto IS 'identificador de la tabla concepto';
ALTER TABLE financiera.ingreso_concepto OWNER TO postgres;
-- ddl-end --

-- object: financiera.orden_pago_id_seq | type: SEQUENCE --
-- DROP SEQUENCE financiera.orden_pago_id_seq;
CREATE SEQUENCE financiera.orden_pago_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
ALTER SEQUENCE financiera.orden_pago_id_seq OWNER TO postgres;
-- ddl-end --

-- object: financiera.orden_pago | type: TABLE --
-- DROP TABLE financiera.orden_pago;
CREATE  TABLE financiera.orden_pago(
	id integer NOT NULL DEFAULT nextval('orden_pago_id_seq'::regclass),
	vigencia numeric(4,0) NOT NULL,
	fecha_creacion date NOT NULL,
	registro_presupuestal integer NOT NULL,
	valor_total numeric NOT NULL,
	persona_elaboro integer NOT NULL,
	convenio integer,
	tipo_orden_pago integer NOT NULL,
	unidad_ejecutora integer,
	estado_orden_pago integer NOT NULL,
	CONSTRAINT orden_pago_pk PRIMARY KEY (id)

);
-- ddl-end --
ALTER TABLE financiera.orden_pago OWNER TO postgres;
-- ddl-end --

-- object: financiera.concepto_orden_pago_id_seq | type: SEQUENCE --
-- DROP SEQUENCE financiera.concepto_orden_pago_id_seq;
CREATE SEQUENCE financiera.concepto_orden_pago_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
ALTER SEQUENCE financiera.concepto_orden_pago_id_seq OWNER TO postgres;
-- ddl-end --

-- object: financiera.concepto_orden_pago | type: TABLE --
-- DROP TABLE financiera.concepto_orden_pago;
CREATE  TABLE financiera.concepto_orden_pago(
	id integer NOT NULL DEFAULT nextval('concepto_orden_pago_id_seq'::regclass),
	valor numeric NOT NULL,
	concepto integer NOT NULL,
	orden_de_pago integer NOT NULL,
	CONSTRAINT detalle_orden_pago_pk PRIMARY KEY (id)

);
-- ddl-end --
ALTER TABLE financiera.concepto_orden_pago OWNER TO postgres;
-- ddl-end --

-- object: financiera.tipo_orden_pago_id_seq | type: SEQUENCE --
-- DROP SEQUENCE financiera.tipo_orden_pago_id_seq;
CREATE SEQUENCE financiera.tipo_orden_pago_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
ALTER SEQUENCE financiera.tipo_orden_pago_id_seq OWNER TO postgres;
-- ddl-end --

-- object: financiera.tipo_orden_pago | type: TABLE --
-- DROP TABLE financiera.tipo_orden_pago;
CREATE  TABLE financiera.tipo_orden_pago(
	id integer NOT NULL DEFAULT nextval('tipo_orden_pago_id_seq'::regclass),
	nombre character varying(150) NOT NULL,
	estado_activo boolean NOT NULL,
	descripcion character varying,
	CONSTRAINT tipo_orden_pago_pk PRIMARY KEY (id)

);
-- ddl-end --
ALTER TABLE financiera.tipo_orden_pago OWNER TO postgres;
-- ddl-end --

-- object: financiera.estado_orden_pago_id_seq | type: SEQUENCE --
-- DROP SEQUENCE financiera.estado_orden_pago_id_seq;
CREATE SEQUENCE financiera.estado_orden_pago_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
ALTER SEQUENCE financiera.estado_orden_pago_id_seq OWNER TO postgres;
-- ddl-end --

-- object: financiera.estado_orden_pago | type: TABLE --
-- DROP TABLE financiera.estado_orden_pago;
CREATE  TABLE financiera.estado_orden_pago(
	id integer NOT NULL DEFAULT nextval('estado_orden_pago_id_seq'::regclass),
	nombre character varying(150) NOT NULL,
	estado_activo boolean NOT NULL,
	descripcion character varying,
	CONSTRAINT estado_orden_pago_pk PRIMARY KEY (id)

);
-- ddl-end --
ALTER TABLE financiera.estado_orden_pago OWNER TO postgres;
-- ddl-end --

-- object: fk_anulacion | type: CONSTRAINT --
-- ALTER TABLE financiera.anulacion_disponibilidad_apropiacion DROP CONSTRAINT fk_anulacion;
ALTER TABLE financiera.anulacion_disponibilidad_apropiacion ADD CONSTRAINT fk_anulacion FOREIGN KEY (anulacion)
REFERENCES financiera.anulacion_disponibilidad (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --


-- object: fk_disponibilidad_apropiacion | type: CONSTRAINT --
-- ALTER TABLE financiera.anulacion_disponibilidad_apropiacion DROP CONSTRAINT fk_disponibilidad_apropiacion;
ALTER TABLE financiera.anulacion_disponibilidad_apropiacion ADD CONSTRAINT fk_disponibilidad_apropiacion FOREIGN KEY (disponibilidad_apropiacion)
REFERENCES financiera.disponibilidad_apropiacion (id) MATCH SIMPLE
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --


-- object: fk_anulacion_reserva | type: CONSTRAINT --
-- ALTER TABLE financiera.anulacion_reserva DROP CONSTRAINT fk_anulacion_reserva;
ALTER TABLE financiera.anulacion_reserva ADD CONSTRAINT fk_anulacion_reserva FOREIGN KEY (reserva_presupuestal)
REFERENCES financiera.reserva_presupuestal (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --


-- object: fk_apropiacion_estado_apropiacion | type: CONSTRAINT --
-- ALTER TABLE financiera.apropiacion DROP CONSTRAINT fk_apropiacion_estado_apropiacion;
ALTER TABLE financiera.apropiacion ADD CONSTRAINT fk_apropiacion_estado_apropiacion FOREIGN KEY (estado)
REFERENCES financiera.estado_apropiacion (id) MATCH SIMPLE
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --


-- object: fk_apropiacion_rubro | type: CONSTRAINT --
-- ALTER TABLE financiera.apropiacion DROP CONSTRAINT fk_apropiacion_rubro;
ALTER TABLE financiera.apropiacion ADD CONSTRAINT fk_apropiacion_rubro FOREIGN KEY (rubro)
REFERENCES financiera.rubro (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --


-- object: fk_disponibilidad_apropiacion_apropiacion | type: CONSTRAINT --
-- ALTER TABLE financiera.disponibilidad_apropiacion DROP CONSTRAINT fk_disponibilidad_apropiacion_apropiacion;
ALTER TABLE financiera.disponibilidad_apropiacion ADD CONSTRAINT fk_disponibilidad_apropiacion_apropiacion FOREIGN KEY (apropiacion)
REFERENCES financiera.apropiacion (id) MATCH SIMPLE
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --


-- object: fk_disponibilidad_apropiacion_disponibilidad | type: CONSTRAINT --
-- ALTER TABLE financiera.disponibilidad_apropiacion DROP CONSTRAINT fk_disponibilidad_apropiacion_disponibilidad;
ALTER TABLE financiera.disponibilidad_apropiacion ADD CONSTRAINT fk_disponibilidad_apropiacion_disponibilidad FOREIGN KEY (disponibilidad)
REFERENCES financiera.disponibilidad (id) MATCH SIMPLE
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --


-- object: fk_disponibilidad_estado_disponibilidad | type: CONSTRAINT --
-- ALTER TABLE financiera.disponibilidad DROP CONSTRAINT fk_disponibilidad_estado_disponibilidad;
ALTER TABLE financiera.disponibilidad ADD CONSTRAINT fk_disponibilidad_estado_disponibilidad FOREIGN KEY (estado)
REFERENCES financiera.estado_disponibilidad (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --


-- object: fk_unidad_ejecutora | type: CONSTRAINT --
-- ALTER TABLE financiera.disponibilidad DROP CONSTRAINT fk_unidad_ejecutora;
ALTER TABLE financiera.disponibilidad ADD CONSTRAINT fk_unidad_ejecutora FOREIGN KEY (unidad_ejecutora)
REFERENCES financiera.unidad_ejecutora (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --


-- object: fk_modificacion_presupuestal_apropiacion | type: CONSTRAINT --
-- ALTER TABLE financiera.modificacion_presupuestal DROP CONSTRAINT fk_modificacion_presupuestal_apropiacion;
ALTER TABLE financiera.modificacion_presupuestal ADD CONSTRAINT fk_modificacion_presupuestal_apropiacion FOREIGN KEY (apropiacion)
REFERENCES financiera.apropiacion (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --


-- object: fk_registo_presupuestal_estado | type: CONSTRAINT --
-- ALTER TABLE financiera.registro_presupuestal DROP CONSTRAINT fk_registo_presupuestal_estado;
ALTER TABLE financiera.registro_presupuestal ADD CONSTRAINT fk_registo_presupuestal_estado FOREIGN KEY (estado)
REFERENCES financiera.estado_registro_presupuestal (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --


-- object: fk_disponibilidad_apropiacion | type: CONSTRAINT --
-- ALTER TABLE financiera.registro_presupuestal_disponibilidad_apropiacion DROP CONSTRAINT fk_disponibilidad_apropiacion;
ALTER TABLE financiera.registro_presupuestal_disponibilidad_apropiacion ADD CONSTRAINT fk_disponibilidad_apropiacion FOREIGN KEY (disponibilidad_apropiacion)
REFERENCES financiera.disponibilidad_apropiacion (id) MATCH SIMPLE
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --


-- object: fk_registro_presupuestal | type: CONSTRAINT --
-- ALTER TABLE financiera.registro_presupuestal_disponibilidad_apropiacion DROP CONSTRAINT fk_registro_presupuestal;
ALTER TABLE financiera.registro_presupuestal_disponibilidad_apropiacion ADD CONSTRAINT fk_registro_presupuestal FOREIGN KEY (registro_presupuestal)
REFERENCES financiera.registro_presupuestal (id) MATCH SIMPLE
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --


-- object: fk_reserva_estado_reserva | type: CONSTRAINT --
-- ALTER TABLE financiera.reserva_presupuestal DROP CONSTRAINT fk_reserva_estado_reserva;
ALTER TABLE financiera.reserva_presupuestal ADD CONSTRAINT fk_reserva_estado_reserva FOREIGN KEY (estado)
REFERENCES financiera.estado_reserva_presupuestal (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --


-- object: fk_reserva_presupuestal_rubro | type: CONSTRAINT --
-- ALTER TABLE financiera.reserva_presupuestal DROP CONSTRAINT fk_reserva_presupuestal_rubro;
ALTER TABLE financiera.reserva_presupuestal ADD CONSTRAINT fk_reserva_presupuestal_rubro FOREIGN KEY (rubro)
REFERENCES financiera.rubro (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --


-- object: fk_rubro_homologado_rubro | type: CONSTRAINT --
-- ALTER TABLE financiera.rubro_homologado DROP CONSTRAINT fk_rubro_homologado_rubro;
ALTER TABLE financiera.rubro_homologado ADD CONSTRAINT fk_rubro_homologado_rubro FOREIGN KEY (rubro)
REFERENCES financiera.rubro (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --


-- object: fk_entidad_homologado | type: CONSTRAINT --
-- ALTER TABLE financiera.rubro_homologado DROP CONSTRAINT fk_entidad_homologado;
ALTER TABLE financiera.rubro_homologado ADD CONSTRAINT fk_entidad_homologado FOREIGN KEY (entidad_homologado)
REFERENCES financiera.entidad (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --


-- object: fk_entidad | type: CONSTRAINT --
-- ALTER TABLE financiera.rubro DROP CONSTRAINT fk_entidad;
ALTER TABLE financiera.rubro ADD CONSTRAINT fk_entidad FOREIGN KEY (entidad)
REFERENCES financiera.entidad (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --


-- object: fk_rubro_hijo | type: CONSTRAINT --
-- ALTER TABLE financiera.rubro_rubro DROP CONSTRAINT fk_rubro_hijo;
ALTER TABLE financiera.rubro_rubro ADD CONSTRAINT fk_rubro_hijo FOREIGN KEY (rubro_hijo)
REFERENCES financiera.rubro (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --


-- object: fk_rubro_padre | type: CONSTRAINT --
-- ALTER TABLE financiera.rubro_rubro DROP CONSTRAINT fk_rubro_padre;
ALTER TABLE financiera.rubro_rubro ADD CONSTRAINT fk_rubro_padre FOREIGN KEY (rubro_padre)
REFERENCES financiera.rubro (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --


-- object: nivel_clasificacion_fk | type: CONSTRAINT --
-- ALTER TABLE financiera.cuenta_contable DROP CONSTRAINT nivel_clasificacion_fk;
ALTER TABLE financiera.cuenta_contable ADD CONSTRAINT nivel_clasificacion_fk FOREIGN KEY (nivel_clasificacion)
REFERENCES financiera.nivel_clasificacion (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --


-- object: cuenta_bancaria_fk | type: CONSTRAINT --
-- ALTER TABLE financiera.cuenta_contable DROP CONSTRAINT cuenta_bancaria_fk;
ALTER TABLE financiera.cuenta_contable ADD CONSTRAINT cuenta_bancaria_fk FOREIGN KEY (cuenta_bancaria)
REFERENCES financiera.cuenta_bancaria (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --


-- object: tipo_documento_afectante_fk | type: CONSTRAINT --
-- ALTER TABLE financiera.movimiento_contable DROP CONSTRAINT tipo_documento_afectante_fk;
ALTER TABLE financiera.movimiento_contable ADD CONSTRAINT tipo_documento_afectante_fk FOREIGN KEY (tipo_documento_afectante)
REFERENCES financiera.tipo_documento_afectante (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --


-- object: cuenta_contable_fk | type: CONSTRAINT --
-- ALTER TABLE financiera.movimiento_contable DROP CONSTRAINT cuenta_contable_fk;
ALTER TABLE financiera.movimiento_contable ADD CONSTRAINT cuenta_contable_fk FOREIGN KEY (cuenta_contable)
REFERENCES financiera.cuenta_contable (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --


-- object: cuenta_padre_fk | type: CONSTRAINT --
-- ALTER TABLE financiera.estructura_cuentas DROP CONSTRAINT cuenta_padre_fk;
ALTER TABLE financiera.estructura_cuentas ADD CONSTRAINT cuenta_padre_fk FOREIGN KEY (cuenta_padre)
REFERENCES financiera.cuenta_contable (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --


-- object: cuenta_hijo_fk | type: CONSTRAINT --
-- ALTER TABLE financiera.estructura_cuentas DROP CONSTRAINT cuenta_hijo_fk;
ALTER TABLE financiera.estructura_cuentas ADD CONSTRAINT cuenta_hijo_fk FOREIGN KEY (cuenta_hijo)
REFERENCES financiera.cuenta_contable (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --


-- object: plan_cuentas_fk | type: CONSTRAINT --
-- ALTER TABLE financiera.estructura_cuentas DROP CONSTRAINT plan_cuentas_fk;
ALTER TABLE financiera.estructura_cuentas ADD CONSTRAINT plan_cuentas_fk FOREIGN KEY (plan_cuentas)
REFERENCES financiera.plan_cuentas (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --


-- object: unidad_ejecutora_fk | type: CONSTRAINT --
-- ALTER TABLE financiera.plan_cuentas DROP CONSTRAINT unidad_ejecutora_fk;
ALTER TABLE financiera.plan_cuentas ADD CONSTRAINT unidad_ejecutora_fk FOREIGN KEY (unidad_ejecutora)
REFERENCES financiera.unidad_ejecutora (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --


-- object: periodo_contable_fk | type: CONSTRAINT --
-- ALTER TABLE financiera.periodo_plan_cuentas DROP CONSTRAINT periodo_contable_fk;
ALTER TABLE financiera.periodo_plan_cuentas ADD CONSTRAINT periodo_contable_fk FOREIGN KEY (periodo_contable)
REFERENCES financiera.periodo_contable (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --


-- object: plan_cuentas_fk | type: CONSTRAINT --
-- ALTER TABLE financiera.periodo_plan_cuentas DROP CONSTRAINT plan_cuentas_fk;
ALTER TABLE financiera.periodo_plan_cuentas ADD CONSTRAINT plan_cuentas_fk FOREIGN KEY (plan_cuentas)
REFERENCES financiera.plan_cuentas (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --


-- object: cuenta_contable_fk | type: CONSTRAINT --
-- ALTER TABLE financiera.concepto_cuenta_contable DROP CONSTRAINT cuenta_contable_fk;
ALTER TABLE financiera.concepto_cuenta_contable ADD CONSTRAINT cuenta_contable_fk FOREIGN KEY (cuenta_contable)
REFERENCES financiera.cuenta_contable (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --


-- object: concepto_fk | type: CONSTRAINT --
-- ALTER TABLE financiera.concepto_cuenta_contable DROP CONSTRAINT concepto_fk;
ALTER TABLE financiera.concepto_cuenta_contable ADD CONSTRAINT concepto_fk FOREIGN KEY (concepto)
REFERENCES financiera.concepto (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --


-- object: tipo_cuenta_fk | type: CONSTRAINT --
-- ALTER TABLE financiera.cuenta_bancaria DROP CONSTRAINT tipo_cuenta_fk;
ALTER TABLE financiera.cuenta_bancaria ADD CONSTRAINT tipo_cuenta_fk FOREIGN KEY (tipo_cuenta)
REFERENCES financiera.tipo_cuenta_bancaria (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --


-- object: tipo_recurso_fk | type: CONSTRAINT --
-- ALTER TABLE financiera.cuenta_bancaria DROP CONSTRAINT tipo_recurso_fk;
ALTER TABLE financiera.cuenta_bancaria ADD CONSTRAINT tipo_recurso_fk FOREIGN KEY (tipo_recurso)
REFERENCES financiera.tipo_recurso (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --


-- object: unidad_ejecutora_fk | type: CONSTRAINT --
-- ALTER TABLE financiera.cuenta_bancaria DROP CONSTRAINT unidad_ejecutora_fk;
ALTER TABLE financiera.cuenta_bancaria ADD CONSTRAINT unidad_ejecutora_fk FOREIGN KEY (unidad_ejecutora)
REFERENCES financiera.unidad_ejecutora (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --


-- object: tipo_entidad_fk | type: CONSTRAINT --
-- ALTER TABLE financiera.entidad DROP CONSTRAINT tipo_entidad_fk;
ALTER TABLE financiera.entidad ADD CONSTRAINT tipo_entidad_fk FOREIGN KEY (tipo_entidad)
REFERENCES financiera.tipo_entidad (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --


-- object: entidad_fk | type: CONSTRAINT --
-- ALTER TABLE financiera.unidad_ejecutora DROP CONSTRAINT entidad_fk;
ALTER TABLE financiera.unidad_ejecutora ADD CONSTRAINT entidad_fk FOREIGN KEY (entidad)
REFERENCES financiera.entidad (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --


-- object: categoria_compromiso_fk | type: CONSTRAINT --
-- ALTER TABLE financiera.tipo_compromiso_tesoral DROP CONSTRAINT categoria_compromiso_fk;
ALTER TABLE financiera.tipo_compromiso_tesoral ADD CONSTRAINT categoria_compromiso_fk FOREIGN KEY (categoria_compromiso)
REFERENCES financiera.categoria_compromiso (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --


-- object: tipo_compromiso_tesoral_fk | type: CONSTRAINT --
-- ALTER TABLE financiera.compromiso DROP CONSTRAINT tipo_compromiso_tesoral_fk;
ALTER TABLE financiera.compromiso ADD CONSTRAINT tipo_compromiso_tesoral_fk FOREIGN KEY (tipo_compromiso_tesoral)
REFERENCES financiera.tipo_compromiso_tesoral (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --


-- object: estado_compromiso_fk | type: CONSTRAINT --
-- ALTER TABLE financiera.compromiso DROP CONSTRAINT estado_compromiso_fk;
ALTER TABLE financiera.compromiso ADD CONSTRAINT estado_compromiso_fk FOREIGN KEY (estado_compromiso)
REFERENCES financiera.estado_compromiso (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --


-- object: unidad_ejecutora_fk | type: CONSTRAINT --
-- ALTER TABLE financiera.compromiso DROP CONSTRAINT unidad_ejecutora_fk;
ALTER TABLE financiera.compromiso ADD CONSTRAINT unidad_ejecutora_fk FOREIGN KEY (unidad_ejecutora)
REFERENCES financiera.unidad_ejecutora (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --


-- object: tipo_transaccion_fk | type: CONSTRAINT --
-- ALTER TABLE financiera.version_tipo_transaccion DROP CONSTRAINT tipo_transaccion_fk;
ALTER TABLE financiera.version_tipo_transaccion ADD CONSTRAINT tipo_transaccion_fk FOREIGN KEY (tipo_transaccion)
REFERENCES financiera.tipo_transaccion (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --


-- object: clase_transaccion_fk | type: CONSTRAINT --
-- ALTER TABLE financiera.tipo_transaccion DROP CONSTRAINT clase_transaccion_fk;
ALTER TABLE financiera.tipo_transaccion ADD CONSTRAINT clase_transaccion_fk FOREIGN KEY (clase_transaccion)
REFERENCES financiera.clase_transaccion (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --


-- object: modulo_kronos_fk | type: CONSTRAINT --
-- ALTER TABLE financiera.clase_transaccion DROP CONSTRAINT modulo_kronos_fk;
ALTER TABLE financiera.clase_transaccion ADD CONSTRAINT modulo_kronos_fk FOREIGN KEY (modulo_kronos)
REFERENCES financiera.modulo_kronos (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --


-- object: concepto_fk | type: CONSTRAINT --
-- ALTER TABLE financiera.afectacion_concepto DROP CONSTRAINT concepto_fk;
ALTER TABLE financiera.afectacion_concepto ADD CONSTRAINT concepto_fk FOREIGN KEY (concepto)
REFERENCES financiera.concepto (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --


-- object: tipo_afectacion_fk | type: CONSTRAINT --
-- ALTER TABLE financiera.afectacion_concepto DROP CONSTRAINT tipo_afectacion_fk;
ALTER TABLE financiera.afectacion_concepto ADD CONSTRAINT tipo_afectacion_fk FOREIGN KEY (tipo_afectacion)
REFERENCES financiera.tipo_afectacion (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --


-- object: concepto_padre_fk | type: CONSTRAINT --
-- ALTER TABLE financiera.concepto_concepto DROP CONSTRAINT concepto_padre_fk;
ALTER TABLE financiera.concepto_concepto ADD CONSTRAINT concepto_padre_fk FOREIGN KEY (concepto_padre)
REFERENCES financiera.concepto (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --


-- object: concepto_hijo_fk | type: CONSTRAINT --
-- ALTER TABLE financiera.concepto_concepto DROP CONSTRAINT concepto_hijo_fk;
ALTER TABLE financiera.concepto_concepto ADD CONSTRAINT concepto_hijo_fk FOREIGN KEY (concepto_hijo)
REFERENCES financiera.concepto (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --


-- object: tipo_concepto_fk | type: CONSTRAINT --
-- ALTER TABLE financiera.concepto DROP CONSTRAINT tipo_concepto_fk;
ALTER TABLE financiera.concepto ADD CONSTRAINT tipo_concepto_fk FOREIGN KEY (tipo_concepto)
REFERENCES financiera.tipo_concepto (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --


-- object: rubro_fk | type: CONSTRAINT --
-- ALTER TABLE financiera.concepto DROP CONSTRAINT rubro_fk;
ALTER TABLE financiera.concepto ADD CONSTRAINT rubro_fk FOREIGN KEY (rubro)
REFERENCES financiera.rubro (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --


-- object: forma_ingreso_fk | type: CONSTRAINT --
-- ALTER TABLE financiera.ingreso DROP CONSTRAINT forma_ingreso_fk;
ALTER TABLE financiera.ingreso ADD CONSTRAINT forma_ingreso_fk FOREIGN KEY (forma_ingreso)
REFERENCES financiera.forma_ingreso (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --


-- object: estado_ingreso_fk | type: CONSTRAINT --
-- ALTER TABLE financiera.ingreso DROP CONSTRAINT estado_ingreso_fk;
ALTER TABLE financiera.ingreso ADD CONSTRAINT estado_ingreso_fk FOREIGN KEY (estado_ingreso)
REFERENCES financiera.estado_ingreso (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --


-- object: unidad_ejecutora_fk | type: CONSTRAINT --
-- ALTER TABLE financiera.ingreso DROP CONSTRAINT unidad_ejecutora_fk;
ALTER TABLE financiera.ingreso ADD CONSTRAINT unidad_ejecutora_fk FOREIGN KEY (unidad_ejecutora)
REFERENCES financiera.unidad_ejecutora (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --


-- object: ingreso_fk | type: CONSTRAINT --
-- ALTER TABLE financiera.ingreso_concepto DROP CONSTRAINT ingreso_fk;
ALTER TABLE financiera.ingreso_concepto ADD CONSTRAINT ingreso_fk FOREIGN KEY (ingreso)
REFERENCES financiera.ingreso (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --


-- object: concepto_fk | type: CONSTRAINT --
-- ALTER TABLE financiera.ingreso_concepto DROP CONSTRAINT concepto_fk;
ALTER TABLE financiera.ingreso_concepto ADD CONSTRAINT concepto_fk FOREIGN KEY (concepto)
REFERENCES financiera.concepto (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --


-- object: unidad_ejecutora_fk | type: CONSTRAINT --
-- ALTER TABLE financiera.orden_pago DROP CONSTRAINT unidad_ejecutora_fk;
ALTER TABLE financiera.orden_pago ADD CONSTRAINT unidad_ejecutora_fk FOREIGN KEY (unidad_ejecutora)
REFERENCES financiera.unidad_ejecutora (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --


-- object: estado_orden_pago_fk | type: CONSTRAINT --
-- ALTER TABLE financiera.orden_pago DROP CONSTRAINT estado_orden_pago_fk;
ALTER TABLE financiera.orden_pago ADD CONSTRAINT estado_orden_pago_fk FOREIGN KEY (estado_orden_pago)
REFERENCES financiera.estado_orden_pago (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --


-- object: tipo_orden_pago_fk | type: CONSTRAINT --
-- ALTER TABLE financiera.orden_pago DROP CONSTRAINT tipo_orden_pago_fk;
ALTER TABLE financiera.orden_pago ADD CONSTRAINT tipo_orden_pago_fk FOREIGN KEY (tipo_orden_pago)
REFERENCES financiera.tipo_orden_pago (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --


-- object: fk_registro_presupuestal | type: CONSTRAINT --
-- ALTER TABLE financiera.orden_pago DROP CONSTRAINT fk_registro_presupuestal;
ALTER TABLE financiera.orden_pago ADD CONSTRAINT fk_registro_presupuestal FOREIGN KEY (registro_presupuestal)
REFERENCES financiera.registro_presupuestal (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --


-- object: orden_de_pago_fk | type: CONSTRAINT --
-- ALTER TABLE financiera.concepto_orden_pago DROP CONSTRAINT orden_de_pago_fk;
ALTER TABLE financiera.concepto_orden_pago ADD CONSTRAINT orden_de_pago_fk FOREIGN KEY (orden_de_pago)
REFERENCES financiera.orden_pago (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --


-- object: concepto_fk | type: CONSTRAINT --
-- ALTER TABLE financiera.concepto_orden_pago DROP CONSTRAINT concepto_fk;
ALTER TABLE financiera.concepto_orden_pago ADD CONSTRAINT concepto_fk FOREIGN KEY (concepto)
REFERENCES financiera.concepto (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --



