CREATE SCHEMA manager;
ALTER SCHEMA manager OWNER TO postgres;
CREATE SEQUENCE manager."CTHEA_ENDPOINT_id_endpoint_seq"
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER TABLE manager."CTHEA_ENDPOINT_id_endpoint_seq" OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

CREATE TABLE manager."CTHEA_ENDPOINT" (
    id_endpoint numeric(8,0) DEFAULT nextval('manager."CTHEA_ENDPOINT_id_endpoint_seq"'::regclass) NOT NULL,
    name character varying(70),
    description character varying(250),
    creation_date timestamp without time zone DEFAULT now() NOT NULL,
    edit_date timestamp without time zone,
    status boolean DEFAULT true
);


ALTER TABLE manager."CTHEA_ENDPOINT" OWNER TO postgres;

COMMENT ON COLUMN manager."CTHEA_ENDPOINT".id_endpoint IS 'LLAVE PRIMARIA ROLE';

COMMENT ON COLUMN manager."CTHEA_ENDPOINT".name IS 'NOMBRE DE LA PAGINA';

COMMENT ON COLUMN manager."CTHEA_ENDPOINT".description IS 'DESCRIPCIÓN ';

COMMENT ON COLUMN manager."CTHEA_ENDPOINT".creation_date IS 'FECHA DE CREACIÓN';

COMMENT ON COLUMN manager."CTHEA_ENDPOINT".edit_date IS 'FECHA DE EDICIÓN';

COMMENT ON COLUMN manager."CTHEA_ENDPOINT".status IS 'STATUS ACTIVO/INACTIVO';

CREATE SEQUENCE manager."CTHEA_ROLE_id_role_seq"
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE manager."CTHEA_ROLE_id_role_seq" OWNER TO postgres;

CREATE TABLE manager."CTHEA_ROLE" (
    id_role numeric(8,0) DEFAULT nextval('manager."CTHEA_ROLE_id_role_seq"'::regclass) NOT NULL,
    name character varying(50),
    code character varying(50),
    creation_date timestamp without time zone DEFAULT now() NOT NULL,
    edit_date timestamp without time zone,
    status boolean DEFAULT true
);


ALTER TABLE manager."CTHEA_ROLE" OWNER TO postgres;

COMMENT ON COLUMN manager."CTHEA_ROLE".id_role IS 'LLAVE PRIMARIA ROLE';

COMMENT ON COLUMN manager."CTHEA_ROLE".name IS 'NOMBRE DEL ROL';

COMMENT ON COLUMN manager."CTHEA_ROLE".code IS 'CODIGO';

COMMENT ON COLUMN manager."CTHEA_ROLE".creation_date IS 'FECHA DE CREACIÓN';

COMMENT ON COLUMN manager."CTHEA_ROLE".edit_date IS 'FECHA DE EDICIÓN';

COMMENT ON COLUMN manager."CTHEA_ROLE".status IS 'STATUS ACTIVO/INACTIVO';

CREATE SEQUENCE manager."THEA_DETAILS_id_details_endpoint_and_hostin_seq"
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE manager."THEA_DETAILS_id_details_endpoint_and_hostin_seq" OWNER TO postgres;

CREATE TABLE manager."THEA_DETAILS_ENDPOINT_AND_HOSTIN" (
    id_details_endpoint_and_hostin numeric(8,0) DEFAULT nextval('manager."THEA_DETAILS_id_details_endpoint_and_hostin_seq"'::regclass) NOT NULL,
    id_details_hosting numeric(8,0) NOT NULL,
    id_endpoint numeric(8,0) NOT NULL,
    creation_date timestamp without time zone DEFAULT now() NOT NULL,
    edit_date timestamp without time zone,
    status boolean DEFAULT true
);


ALTER TABLE manager."THEA_DETAILS_ENDPOINT_AND_HOSTIN" OWNER TO postgres;

COMMENT ON COLUMN manager."THEA_DETAILS_ENDPOINT_AND_HOSTIN".id_details_endpoint_and_hostin IS 'LLAVE PRIMARIA ROLE';

COMMENT ON COLUMN manager."THEA_DETAILS_ENDPOINT_AND_HOSTIN".id_details_hosting IS 'LLAVE FORANEA DETAILS HOSTING';

COMMENT ON COLUMN manager."THEA_DETAILS_ENDPOINT_AND_HOSTIN".id_endpoint IS 'LLAVE FORANEA ENDPOINT';

COMMENT ON COLUMN manager."THEA_DETAILS_ENDPOINT_AND_HOSTIN".creation_date IS 'FECHA DE CREACIÓN';

COMMENT ON COLUMN manager."THEA_DETAILS_ENDPOINT_AND_HOSTIN".edit_date IS 'FECHA DE EDICIÓN';

COMMENT ON COLUMN manager."THEA_DETAILS_ENDPOINT_AND_HOSTIN".status IS 'STATUS ACTIVO/INACTIVO';

CREATE SEQUENCE manager."THEA_DETAILS_HOSTING_id_details_hosting_seq"
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE manager."THEA_DETAILS_HOSTING_id_details_hosting_seq" OWNER TO postgres;

CREATE TABLE manager."THEA_DETAILS_HOSTING" (
    id_details_hosting numeric(8,0) DEFAULT nextval('manager."THEA_DETAILS_HOSTING_id_details_hosting_seq"'::regclass) NOT NULL,
    id_user numeric(8,0) NOT NULL,
    host character varying(50) NOT NULL,
    session_start timestamp without time zone NOT NULL,
    session_closure timestamp without time zone NOT NULL,
    type_log_out character varying(50) NOT NULL,
    creation_date timestamp without time zone DEFAULT now() NOT NULL,
    edit_date timestamp without time zone,
    status boolean DEFAULT true
);


ALTER TABLE manager."THEA_DETAILS_HOSTING" OWNER TO postgres;

COMMENT ON COLUMN manager."THEA_DETAILS_HOSTING".id_details_hosting IS 'LLAVE PRIMARIA DETAILS HOSTING';

COMMENT ON COLUMN manager."THEA_DETAILS_HOSTING".id_user IS 'LLAVE FORANEA USER';

COMMENT ON COLUMN manager."THEA_DETAILS_HOSTING".host IS 'HOSTING DE INICIO DE SESION';

COMMENT ON COLUMN manager."THEA_DETAILS_HOSTING".session_start IS 'INICIO DE SESION';

COMMENT ON COLUMN manager."THEA_DETAILS_HOSTING".session_closure IS 'CIERRE DE SESION';

COMMENT ON COLUMN manager."THEA_DETAILS_HOSTING".type_log_out IS 'TIPO DE CIERRE DE SESION';

COMMENT ON COLUMN manager."THEA_DETAILS_HOSTING".creation_date IS 'FECHA DE CREACIÓN';

COMMENT ON COLUMN manager."THEA_DETAILS_HOSTING".edit_date IS 'FECHA DE EDICIÓN';

COMMENT ON COLUMN manager."THEA_DETAILS_HOSTING".status IS 'STATUS ACTIVO/INACTIVO';

CREATE SEQUENCE manager."THEA_USER_id_user_seq"
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE manager."THEA_USER_id_user_seq" OWNER TO postgres;

CREATE TABLE manager."THEA_USER" (
    id_user numeric(8,0) DEFAULT nextval('manager."THEA_USER_id_user_seq"'::regclass) NOT NULL,
    id_role numeric(8,0) DEFAULT 1 NOT NULL,
    name character varying(150),
    paternal_surname character varying(150),
    maternal_surname character varying(150),
    email character varying(150),
    pasword character varying(150),
    creation_date timestamp without time zone DEFAULT now() NOT NULL,
    edit_date timestamp without time zone,
    status boolean DEFAULT true
);


ALTER TABLE manager."THEA_USER" OWNER TO postgres;

COMMENT ON COLUMN manager."THEA_USER".id_user IS 'LLAVE PRIMARIA USER';

COMMENT ON COLUMN manager."THEA_USER".id_role IS 'LLAVE FORANEA ROLE';

COMMENT ON COLUMN manager."THEA_USER".name IS 'NOMBRE DE USUARIO';

COMMENT ON COLUMN manager."THEA_USER".paternal_surname IS 'APELLIDO PATERNO';

COMMENT ON COLUMN manager."THEA_USER".maternal_surname IS 'APELLIDO MATERNO';

COMMENT ON COLUMN manager."THEA_USER".email IS 'EMAIL';

COMMENT ON COLUMN manager."THEA_USER".pasword IS 'PASSWORD';

COMMENT ON COLUMN manager."THEA_USER".creation_date IS 'FECHA DE CREACIÓN';

COMMENT ON COLUMN manager."THEA_USER".edit_date IS 'FECHA DE EDICIÓN';

COMMENT ON COLUMN manager."THEA_USER".status IS 'STATUS ACTIVO/INACTIVO';

SELECT pg_catalog.setval('manager."CTHEA_ENDPOINT_id_endpoint_seq"', 36, true);

SELECT pg_catalog.setval('manager."CTHEA_ROLE_id_role_seq"', 1, true);

SELECT pg_catalog.setval('manager."THEA_DETAILS_HOSTING_id_details_hosting_seq"', 1, false);

SELECT pg_catalog.setval('manager."THEA_DETAILS_id_details_endpoint_and_hostin_seq"', 1, false);

SELECT pg_catalog.setval('manager."THEA_USER_id_user_seq"', 1, false);

ALTER TABLE ONLY manager."CTHEA_ENDPOINT"
    ADD CONSTRAINT "PK_CTHEA_ENDPOINT" PRIMARY KEY (id_endpoint);

ALTER TABLE ONLY manager."THEA_DETAILS_ENDPOINT_AND_HOSTIN"
    ADD CONSTRAINT "PK_THEA_DETAILS_ENDPOINT_AND_HOSTIN" PRIMARY KEY (id_details_endpoint_and_hostin);

ALTER TABLE ONLY manager."THEA_DETAILS_HOSTING"
    ADD CONSTRAINT "PK_THEA_DETAILS_HOSTING" PRIMARY KEY (id_details_hosting);

ALTER TABLE ONLY manager."CTHEA_ROLE"
    ADD CONSTRAINT "PK_THEA_ROLE" PRIMARY KEY (id_role);

ALTER TABLE ONLY manager."THEA_USER"
    ADD CONSTRAINT "PK_THEA_USER" PRIMARY KEY (id_user);

ALTER TABLE ONLY manager."THEA_DETAILS_ENDPOINT_AND_HOSTIN"
    ADD CONSTRAINT "FK_THEA_DETAILS_ENDPOINT_AND_HOSTIN_DETAILS_HOSTING" FOREIGN KEY (id_details_hosting) REFERENCES manager."THEA_DETAILS_HOSTING"(id_details_hosting);


ALTER TABLE ONLY manager."THEA_DETAILS_ENDPOINT_AND_HOSTIN"
    ADD CONSTRAINT "FK_THEA_DETAILS_ENDPOINT_AND_HOSTIN_ENDPOINT" FOREIGN KEY (id_endpoint) REFERENCES manager."CTHEA_ENDPOINT"(id_endpoint);


ALTER TABLE ONLY manager."THEA_DETAILS_HOSTING"
    ADD CONSTRAINT "FK_THEA_DETAILS_HOSTING_USER" FOREIGN KEY (id_user) REFERENCES manager."THEA_USER"(id_user);


ALTER TABLE ONLY manager."THEA_USER"
    ADD CONSTRAINT "FK_THEA_USER_ROLE" FOREIGN KEY (id_role) REFERENCES manager."CTHEA_ROLE"(id_role);
