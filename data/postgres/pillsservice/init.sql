CREATE SCHEMA public AUTHORIZATION pg_database_owner;

COMMENT ON SCHEMA public IS 'standard public schema';

CREATE SEQUENCE public.pill_id_seq
    INCREMENT BY 1
    MINVALUE 1
    MAXVALUE 2147483647
    START 1
    CACHE 1
    NO CYCLE;

CREATE TABLE public."pills" (
    id serial4 NOT NULL,
    "name" varchar(100) NOT NULL,
    dosage varchar(100),
    instructions TEXT,
    createdDate timestamp DEFAULT now() NOT NULL,
    deletedDate timestamp,
    CONSTRAINT pills_pk PRIMARY KEY (id)
);