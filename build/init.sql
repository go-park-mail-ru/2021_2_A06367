-- DROP DATABASE Netflix;
CREATE DATABASE Netflix
    WITH
    OWNER = postgres
    ENCODING = 'UTF8'
    LC_COLLATE = 'Russian_Russia.1251'
    LC_CTYPE = 'Russian_Russia.1251'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1;


CREATE TABLE users(
    id UUID PRIMARY KEY,
    email text NOT NULL,
    login text UNIQUE NOT NULL,
    encrypted_password text NOT NULL,
    created_at TIMESTAMP NOT NULL
);
