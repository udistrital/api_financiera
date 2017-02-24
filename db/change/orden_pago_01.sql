ALTER TABLE financiera.orden_pago ADD COLUMN iva integer NOT NULL;

-- object: financiera.iva | type: TABLE --
-- DROP TABLE financiera.iva;
CREATE TABLE financiera.iva(
    id serial NOT NULL,
    categoria_iva integer NOT NULL,
    valor numeric(2) NOT NULL,
    estado_activo boolean NOT NULL
);
-- ddl-end --
-- object: financiera.categoria_iva | type: TABLE --
-- DROP TABLE financiera.categoria_iva;
CREATE TABLE financiera.categoria_iva(
    id serial NOT NULL,
    nombre varchar,
    estado_activo boolean NOT NULL
);
-- ddl-end --

-- object: iva_fk | type: CONSTRAINT --
-- ALTER TABLE financiera.orden_pago DROP CONSTRAINT iva_fk;
ALTER TABLE financiera.orden_pago ADD CONSTRAINT iva_fk FOREIGN KEY (iva)
REFERENCES financiera.iva (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: categoria_iva | type: CONSTRAINT --
-- ALTER TABLE financiera.iva DROP CONSTRAINT categoria_iva;
ALTER TABLE financiera.iva ADD CONSTRAINT categoria_iva FOREIGN KEY (categoria_iva)
REFERENCES financiera.categoria_iva (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --