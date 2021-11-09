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
    src text[] NOT NULL,
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
    description text NOT NULL,
    genres text[] NOT NULL
);

CREATE OR REPLACE FUNCTION make_tsvector(title TEXT)
   RETURNS tsvector AS $$
BEGIN
RETURN (setweight(to_tsvector('english', title),'A') ||
        setweight(to_tsvector('russian', title), 'B'));
END
$$ LANGUAGE 'plpgsql' IMMUTABLE;

-- actors
INSERT INTO public.actors (id, name, surname, avatar, height, date_of_birth, description, genres)
VALUES ('3e06d4e4-3b47-11ec-8d3d-0242ac130003', 'James', 'Bond', '/ewf/xxx/xxx', 1.78,
        '1985-11-01 22:10:57.000000', 'some text', '{Comedy,Thriller}');
INSERT INTO public.actors (id, name, surname, avatar, height, date_of_birth, description, genres)
VALUES ('9743f488-3b47-11ec-8d3d-0242ac130003', 'Nick', 'Ivanov', '/ewf/xxx/xxx', 1.78,
        '1985-11-01 22:10:57.000000', 'some text', '{Comedy,Thriller}');

-- films
INSERT INTO public.films (id, genres, title, year, director, authors, actors, release, duration, language, src)
VALUES ('c7020e69-6a77-4153-97bc-54dc905321a4', '{Comedy}', '007', 2019, '{B. Spiars,
        K. Nolan}', '{Natalio Portman, Sergay Borunov}', '{3e06d4e4-3b47-11ec-8d3d-0242ac130003,
        9743f488-3b47-11ec-8d3d-0242ac130003}', '1970-01-01', 18, 'ru', '{/usr/local/test1.mp4}');
INSERT INTO public.films (id, genres, title, year, director, authors, actors, release, duration, language, src)
VALUES ('f8405178-3b47-11ec-8d3d-0242ac130003', '{Thriller}', 'Friday 13', 2015, '{S. Borunov,
        K. Torantino}', '{Megan Fox, Jarald Lito}', '{3e06d4e4-3b47-11ec-8d3d-0242ac130003,
        9743f488-3b47-11ec-8d3d-0242ac130003}', '1995-01-01', 18, 'ru', '{/usr/local/test2.mp4}');

-- rating
INSERT INTO public.rating (film_id, rating) VALUES ('c7020e69-6a77-4153-97bc-54dc905321a4', 4.7);
INSERT INTO public.rating (film_id, rating) VALUES ('f8405178-3b47-11ec-8d3d-0242ac130003', 4.3);

