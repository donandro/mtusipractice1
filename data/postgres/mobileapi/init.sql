CREATE SCHEMA public AUTHORIZATION pg_database_owner;

COMMENT ON SCHEMA public IS 'standard public schema';

CREATE SEQUENCE public.user_id_seq
    INCREMENT BY 1
    MINVALUE 1
    MAXVALUE 2147483647
    START 1
    CACHE 1
    NO CYCLE;

CREATE TABLE public."users" (
    id serial4 NOT NULL,
    "name" varchar(100) NOT NULL,
    login varchar(100) NOT NULL,
    password varchar(100) NOT NULL,
    email varchar(100) NOT NULL,
    created timestamp DEFAULT now() NOT NULL,
    deleted timestamp NULL,
    CONSTRAINT user_pk PRIMARY KEY (id)
);