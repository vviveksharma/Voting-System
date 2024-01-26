CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

BEGIN;

CREATE TABLE IF NOT EXISTS user_tbl (
    id UUID PRIMARY KEY,
    username character varying(255) NOT NULL,
    email character varying(100) NOT NULL,
    first_name character varying(40) NOT NULL,
    last_name character varying(40) NOT NULL,
    voter_id UUID 
);

COMMIT;