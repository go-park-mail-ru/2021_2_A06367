-- DROP DATABASE netflix;
CREATE DATABASE netflix
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
    about text,
    avatar text,
    subscriptions int,
    subscribers int,
    created_at TIMESTAMP NOT NULL
);

CREATE TABLE subscriptions (
    id serial PRIMARY KEY,
    user_id UUID REFERENCES users(id) NOT NULL,
    subscribed_at UUID REFERENCES users(id) NOT NULL,
    UNIQUE (user_id, subscribed_at)
);

CREATE TABLE online_users(
    login text UNIQUE NOT NULL
);

CREATE UNIQUE INDEX online_idx ON online_users (login);

CREATE TABLE films
(
    id bigserial NOT NULL,
    genres text[] NOT NULL,
    title text NOT NULL,
    year integer NOT NULL,
    director text[] NOT NULL,
    authors text[] NOT NULL,
    release date NOT NULL,
    duration integer NOT NULL,
    language text NOT NULL,
    PRIMARY KEY (id),
    Check(year > 0),
	Check(duration > 0)
);


CREATE TABLE rating
(
    film_id integer REFERENCES films(id) NOT NULL,
    rating double precision NOT NULL,
    CONSTRAINT rating_pkey PRIMARY KEY (film_id)
);

CREATE TABLE actors
(
    id UUID PRIMARY KEY,
    name text NOT NULL,
    surname text NOT NULL,
    avatar text NOT NULL,
    height float NOT NULL,
    date_of_birth TIMESTAMP NOT NULL,
    genres text[] NOT NULL
);

CREATE OR REPLACE FUNCTION make_tsvector(title TEXT)
   RETURNS tsvector AS $$
BEGIN
RETURN (setweight(to_tsvector('english', title),'A') ||
        setweight(to_tsvector('russian', title), 'B'));
END
$$ LANGUAGE 'plpgsql' IMMUTABLE;
