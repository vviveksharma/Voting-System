CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

BEGIN;

CREATE TABLE IF NOT EXISTS user_tbl (
    id UUID PRIMARY KEY,
    username character varying(255) NOT NULL,
    secretkey character varying(500) NOT NULL,
    email character varying(100) NOT NULL,
    first_name character varying(40) NOT NULL,
    last_name character varying(40) NOT NULL,
    is_voted boolean,
    is_logged_in boolean,
    is_validated boolean,
    token UUID,
    voter_id UUID 
);

CREATE TABLE IF NOT EXISTS admin_tbl (
    id UUID PRIMARY KEY,
    role character varying(255) NOT NULL,
    is_super_admin boolean
);


CREATE TABLE IF NOT EXISTS candidate_tbl (
    id UUID PRIMARY KEY,
    name character varying(255) NOT NULL,
    count int8 NULL,
    image character varying(500) NULL
);

CREATE TABLE IF NOT EXISTS employee_tbl (
    id UUID PRIMARY KEY,
    email character varying(100) NOT NULL,
    role character varying(255) NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_candidate_tbl_count ON public.candidate_tbl USING btree (count);

COMMIT;
