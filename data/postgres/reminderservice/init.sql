CREATE SCHEMA public AUTHORIZATION pg_database_owner;

COMMENT ON SCHEMA public IS 'standard public schema';

CREATE SEQUENCE public.plan_id_seq
    INCREMENT BY 1
    MINVALUE 1
    MAXVALUE 2147483647
    START 1
    CACHE 1
    NO CYCLE;

CREATE SEQUENCE public.intake_log_id_seq
    INCREMENT BY 1
    MINVALUE 1
    MAXVALUE 2147483647
    START 1
    CACHE 1
    NO CYCLE;

CREATE SEQUENCE public.notification_id_seq
    INCREMENT BY 1
    MINVALUE 1
    MAXVALUE 2147483647
    START 1
    CACHE 1
    NO CYCLE;

CREATE TABLE public."plans" (
    id serial4 NOT NULL,
    "name" varchar(100) NOT NULL,
    pillId int NOT NULL,
    reminderTime timestamp NOT NULL,
    reminderFrequency  varchar(100) NOT NULL,
    daysOfWeek varchar(100) NOT NULL,
    startDate timestamp NOT NULL,
    endDate timestamp NULL,
    createdDate timestamp DEFAULT now() NOT NULL,
    CONSTRAINT plans_pk PRIMARY KEY (id)
);

CREATE TABLE public."notifications" (
    id serial4 NOT NULL,
    planId int NOT NULL,
    userId int NOT NULL,
    reminderTime timestamp NOT NULL,
    sent boolean DEFAULT false,
    CONSTRAINT notifications_pk PRIMARY KEY (id)
);

CREATE TABLE public."intake_log" (
     id serial4 NOT NULL,
     planId int NOT NULL,
     intakeTime timestamp NOT NULL,
     status varchar(10) NOT NULL,
     CONSTRAINT intake_log_pk PRIMARY KEY (id)
);
