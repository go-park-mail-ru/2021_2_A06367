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
    id UUID NOT NULL,
    genres text[] NOT NULL,
    title text NOT NULL,
    year integer NOT NULL,
    director text[] NOT NULL,
    authors text[] NOT NULL,
    actors UUID[] NOT NULL,
    release date NOT NULL,
    duration integer NOT NULL,
    language text NOT NULL,
    PRIMARY KEY (id),
    Check(year > 0),
	Check(duration > 0)
);

CREATE INDEX films_actors_idx ON films USING gin(actors);

CREATE TABLE watchlist
(
    id UUID REFERENCES users(id) NOT NULL,
    film_id UUID REFERENCES films(id) NOT NULL
);

CREATE TABLE rating
(
    film_id UUID REFERENCES films(id) NOT NULL,
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

-- Заполнение таблички актеров
-- CREATE EXTENSION IF NOT EXISTS “uuid—ossp”;
-- insert into actors (id, name, surname, height, dateofbirth, genres)
-- values (uuid_generate_v4() , 'Sergei', 'Burunov', 1.78, current_timestamp, '{"Comedy"}'),
-- 	   (uuid_generate_v4() , 'Alex', 'Petrov', 1.81, current_timestamp, '{"Tragedy"}'),
-- 	   (uuid_generate_v4() , 'Tom', 'Cruse', 1.68, current_timestamp, '{"Triller"}');

-- Заполнение фильмов
-- insert into films (id, genres, title, "year" , director, authors, actors, "release", duration, "language")
-- values (uuid_generate_v4(),
-- 	'{"Comedy"}',
-- 	'Policeman from Rublevka',
-- 	2017,
-- 	'{"Alex Karamzin"}',
-- 	'{"Alex Karamzin"}',
-- 	'{d878758e-763e-45b2-a6a0-746d39df1b43, 60204c4e-b2be-4f06-95fe-6de72e8d22ab}',
-- 	current_date,
-- 	127,
-- 	'russian'),
-- 	(uuid_generate_v4() ,
-- 	'{"Triller"}',
-- 	'Mission Impossible',
-- 	2017,
-- 	'{"Joe Rover"}',
-- 	'{"Tom Tacker"}',
-- 	'{f68a4740-e24b-4919-ad8c-69184b6633fb}',
-- 	current_date,
-- 	120,
-- 	'english');