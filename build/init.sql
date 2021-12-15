-- DROP DATABASE netflix;
CREATE DATABASE netflix;


CREATE TABLE users(
    id UUID PRIMARY KEY,
    login text UNIQUE NOT NULL,
    encrypted_password text NOT NULL,
    about text DEFAULT '',
    avatar text DEFAULT 'userpic.png',
    subscriptions int DEFAULT 0,
    subscribers int DEFAULT 0,
    created_at TIMESTAMP NOT NULL
);

CREATE TABLE subscriptions (
    id serial PRIMARY KEY,
    user_id UUID REFERENCES users(id) NOT NULL,
    subscribed_at UUID REFERENCES users(id) NOT NULL,
    UNIQUE (user_id, subscribed_at)
);


CREATE TABLE films
(
    id UUID NOT NULL,
    genres text[] NOT NULL,
    country text NOT NULL,
    releaserus TIMESTAMP WITH TIME ZONE NOT NULL,
    title text NOT NULL,
    year integer NOT NULL,
    director text[] NOT NULL,
    authors text[] NOT NULL,
    actors UUID[] NOT NULL,
    release date NOT NULL,
    duration integer NOT NULL,
    language text NOT NULL,
    budget text NOT NULL,
    age integer NOT NULL,
    pic text[] NOT NULL,
    src text[] NOT NULL,
    description text NOT NULL,
    isSeries bool DEFAULT false,
    needsPayment bool DEFAULT false,
    slug text NOT NULL,
    PRIMARY KEY (id),
    Check(year > 0),
	Check(duration > 0),
	UNIQUE (slug)
);

CREATE TABLE series_seasons
(
    series_id UUID REFERENCES films(id) NOT NULL,
    id integer,
    pic text[] NOT NULL,
    src text[] NOT NULL
);

ALTER TABLE series_seasons
    ADD CONSTRAINT series_seasons_uniq UNIQUE(series_id, id);

CREATE TABLE starred_films
(
    film_id UUID NOT NULL,
    user_id UUID NOT NULL
);

ALTER TABLE starred_films
    ADD CONSTRAINT uniq_list UNIQUE(film_id, user_id);

CREATE INDEX films_actors_idx ON films USING gin(actors);

CREATE TABLE watchlist
(
    id UUID NOT NULL,
    film_id UUID NOT NULL
);

ALTER TABLE watchlist
    ADD CONSTRAINT watchlist_uniq_list UNIQUE(id,film_id);


CREATE TABLE ratings
(
    id UUID NOT NULL,
    film_id UUID NOT NULL,
    rating float NOT NULL
);

ALTER TABLE ratings
    ADD CONSTRAINT ratings_uniq_list UNIQUE(id,film_id);


CREATE TABLE rating
(
    film_id UUID REFERENCES films(id) NOT NULL,
    rating double precision NOT NULL
--     CONSTRAINT rating_pkey PRIMARY KEY (film_id)
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
VALUES ('3e06d4e4-3b47-11ec-8d3d-0242ac130003', 'Джеймс', 'Бонд', 'jb.png', 1.78,
        '1985-11-01 22:10:57.000000', 'Красивый', '{Комедия ,Триллер}');
INSERT INTO public.actors (id, name, surname, avatar, height, date_of_birth, description, genres)
VALUES ('9743f488-3b47-11ec-8d3d-0242ac130003', 'Павел', 'Табаков', 'pt.png', 1.78,
        '1985-11-01 22:10:57.000000', 'Умный', '{Комедия, Триллер}');

-- films
INSERT INTO public.films (id, genres, country, releaseRus, description, title, year, director, authors, actors, release, duration, language, budget, age, pic, src, slug)
VALUES ('c7020e69-6a77-4153-97bc-54dc905321a4', '{Комедия}', 'США', '2001-01-01', 'История о том, как обычный паренёк из провинции превратился в одного из самых состоятельных людей США. Джордж Джанга (Джонни Депп) с детства мечтал о роскошной жизни: шикарном доме, дорогих автомобилях и лучших девушках США. Находчивый паренёк быстро отыскал оригинальный способ осуществить свою мечту: он решил распространять кокаин в промышленных масштабах. Как найти рынки сбыта? Очень просто — необходимо подсадить на белый порошок главных знаменитостей Америки!',
        '007', 2015, '{S. Borunov,
        K. Тарантино}', '{Роберт Уэйд}', '{3e06d4e4-3b47-11ec-8d3d-0242ac130003,
        9743f488-3b47-11ec-8d3d-0242ac130003}', '1970-01-01', 120, 'Русский', '120 млн', 12, '{matrix.png}', '{video.mp4}', 'matrix');

INSERT INTO public.films (id, genres, country, releaseRus, description, title, year, director, authors, actors, release, duration, language, budget, age, pic, src, slug)
VALUES ('f8405178-3b47-11ec-8d3d-0242ac130003', '{Триллер}', 'США', '2001-01-01', 'История о том, как обычный паренёк из провинции превратился в одного из самых состоятельных людей США. Джордж Джанга (Джонни Депп) с детства мечтал о роскошной жизни: шикарном доме, дорогих автомобилях и лучших девушках США. Находчивый паренёк быстро отыскал оригинальный способ осуществить свою мечту: он решил распространять кокаин в промышленных масштабах. Как найти рынки сбыта? Очень просто — необходимо подсадить на белый порошок главных знаменитостей Америки!',
        'Птичий короб', 2015, '{S. Borunov,
        K. Тарантино}', '{Роберт Уэйд}', '{3e06d4e4-3b47-11ec-8d3d-0242ac130003,
        9743f488-3b47-11ec-8d3d-0242ac130003}', '1995-01-01', 120, 'Русский', '120 млн', 12, '{bird.png}', '{video.mp4}', 'birdbox');

INSERT INTO public.films (id, genres, country, releaseRus, description, title, year, director, authors, actors, release, duration, language, budget, age, pic, src, slug)
VALUES ('649be5f4-46e6-11ec-81d3-0242ac130003', '{Триллер}', 'США', '2001-01-01', 'История о том, как обычный паренёк из провинции превратился в одного из самых состоятельных людей США. Джордж Джанга (Джонни Депп) с детства мечтал о роскошной жизни: шикарном доме, дорогих автомобилях и лучших девушках США. Находчивый паренёк быстро отыскал оригинальный способ осуществить свою мечту: он решил распространять кокаин в промышленных масштабах. Как найти рынки сбыта? Очень просто — необходимо подсадить на белый порошок главных знаменитостей Америки!',
        'Кровавый алмаз', 2015, '{S. Borunov,
        K. Тарантино}', '{Роберт Уэйд}', '{3e06d4e4-3b47-11ec-8d3d-0242ac130003,
        9743f488-3b47-11ec-8d3d-0242ac130003}', '1995-01-01', 120, 'Русский', '120 млн', 12, '{blood.png}', '{video.mp4}', 'blooddiamond');

INSERT INTO public.films (id, genres, country, releaseRus, description, title, year, director, authors, actors, release, duration, language, budget, age, pic, src, slug)
VALUES ('68896eac-46e6-11ec-81d3-0242ac130003', '{Боевик}', 'США', '2001-01-01', 'История о том, как обычный паренёк из провинции превратился в одного из самых состоятельных людей США. Джордж Джанга (Джонни Депп) с детства мечтал о роскошной жизни: шикарном доме, дорогих автомобилях и лучших девушках США. Находчивый паренёк быстро отыскал оригинальный способ осуществить свою мечту: он решил распространять кокаин в промышленных масштабах. Как найти рынки сбыта? Очень просто — необходимо подсадить на белый порошок главных знаменитостей Америки!',
        'Отцовство', 2015, '{S. Borunov,
        K. Тарантино}', '{Роберт Уэйд}', '{3e06d4e4-3b47-11ec-8d3d-0242ac130003,
        9743f488-3b47-11ec-8d3d-0242ac130003}', '1995-01-01', 120, 'Русский', '120 млн', 12, '{father.png}', '{video.mp4}', 'dad');

INSERT INTO public.films (id, genres, country, releaseRus, description, title, year, director, authors, actors, release, duration, language, budget, age, pic, src, slug)
VALUES ('7f1ff974-46e6-11ec-81d3-0242ac130003', '{Фантастика}', 'США', '2001-01-01', 'История о том, как обычный паренёк из провинции превратился в одного из самых состоятельных людей США. Джордж Джанга (Джонни Депп) с детства мечтал о роскошной жизни: шикарном доме, дорогих автомобилях и лучших девушках США. Находчивый паренёк быстро отыскал оригинальный способ осуществить свою мечту: он решил распространять кокаин в промышленных масштабах. Как найти рынки сбыта? Очень просто — необходимо подсадить на белый порошок главных знаменитостей Америки!',
        'Изгой', 2015, '{S. Borunov,
        K. Тарантино}', '{Роберт Уэйд}', '{3e06d4e4-3b47-11ec-8d3d-0242ac130003,
        9743f488-3b47-11ec-8d3d-0242ac130003}', '1995-01-01', 120, 'Русский', '120 млн', 12, '{izgoy.png}', '{video.mp4}', 'alone');

INSERT INTO public.films (id, genres, country, releaseRus, description, title, year, director, authors, actors, release, duration, language, budget, age, pic, src, slug)
VALUES ('9dd8cef4-46e6-11ec-81d3-0242ac130003', '{Комедия}', 'Англия', '2001-01-01', 'История о том, как обычный паренёк из провинции превратился в одного из самых состоятельных людей США. Джордж Джанга (Джонни Депп) с детства мечтал о роскошной жизни: шикарном доме, дорогих автомобилях и лучших девушках США. Находчивый паренёк быстро отыскал оригинальный способ осуществить свою мечту: он решил распространять кокаин в промышленных масштабах. Как найти рынки сбыта? Очень просто — необходимо подсадить на белый порошок главных знаменитостей Америки!',
        'Убийство', 2015, '{S. Borunov,
        K. Тарантино}', '{Роберт Уэйд}', '{3e06d4e4-3b47-11ec-8d3d-0242ac130003,
        9743f488-3b47-11ec-8d3d-0242ac130003}', '1995-01-01', 120, 'Русский', '120 млн', 12, '{kill.png}', '{video.mp4}', 'murder');

INSERT INTO public.films (id, genres, country, releaseRus, description, title, year, director, authors, actors, release, duration, language, budget, age, pic, src, slug)
VALUES ('a8c67370-46e6-11ec-81d3-0242ac130003', '{Фентези}', 'Южная Корея', '2001-01-01', 'История о том, как обычный паренёк из провинции превратился в одного из самых состоятельных людей США. Джордж Джанга (Джонни Депп) с детства мечтал о роскошной жизни: шикарном доме, дорогих автомобилях и лучших девушках США. Находчивый паренёк быстро отыскал оригинальный способ осуществить свою мечту: он решил распространять кокаин в промышленных масштабах. Как найти рынки сбыта? Очень просто — необходимо подсадить на белый порошок главных знаменитостей Америки!',
        'Кокаин', 2015, '{S. Borunov,
        K. Тарантино}', '{Роберт Уэйд}', '{3e06d4e4-3b47-11ec-8d3d-0242ac130003,
        9743f488-3b47-11ec-8d3d-0242ac130003}', '1995-01-01', 120, 'Русский', '120 млн', 12, '{img.png}', '{video.mp4}', 'cocain');

INSERT INTO public.films (id, genres, country, releaseRus, description, title, year, director, authors, actors, release, duration, language, budget, age, pic, src, slug)
VALUES ('ac615b08-46e6-11ec-81d3-0242ac130003', '{Триллер}', 'США', '2001-01-01', 'История о том, как обычный паренёк из провинции превратился в одного из самых состоятельных людей США. Джордж Джанга (Джонни Депп) с детства мечтал о роскошной жизни: шикарном доме, дорогих автомобилях и лучших девушках США. Находчивый паренёк быстро отыскал оригинальный способ осуществить свою мечту: он решил распространять кокаин в промышленных масштабах. Как найти рынки сбыта? Очень просто — необходимо подсадить на белый порошок главных знаменитостей Америки!',
        'Красное уведомление', 2015, '{S. Borunov,
        K. Тарантино}', '{Роберт Уэйд}', '{3e06d4e4-3b47-11ec-8d3d-0242ac130003,
        9743f488-3b47-11ec-8d3d-0242ac130003}','1995-01-01', 120, 'Русский', '120 млн', 12, '{red.png}', '{video.mp4}', 'rednotification');

INSERT INTO public.films (id, genres, country, releaseRus, description, title, year, director, authors, actors, release, duration, language, budget, age, pic, src, slug)
VALUES ('b0840514-46e6-11ec-81d3-0242ac130003', '{Триллер}', 'США','2001-01-01',
        'История о том, как обычный паренёк из провинции превратился в одного из самых состоятельных людей США. Джордж Джанга (Джонни Депп) с детства мечтал о роскошной жизни: шикарном доме, дорогих автомобилях и лучших девушках США. Находчивый паренёк быстро отыскал оригинальный способ осуществить свою мечту: он решил распространять кокаин в промышленных масштабах. Как найти рынки сбыта? Очень просто — необходимо подсадить на белый порошок главных знаменитостей Америки!',
        'Варкрафт', 2015, '{S. Borunov,
        K. Тарантино}', '{Роберт Уэйд}', '{3e06d4e4-3b47-11ec-8d3d-0242ac130003,
        9743f488-3b47-11ec-8d3d-0242ac130003}','1995-01-01', 120, 'Русский', '120 млн', 12, '{warcraft.png}', '{video.mp4}', 'warcraft');

INSERT INTO public.films (id, genres, country, releaseRus, description, title, year, director, authors, actors, release, duration, language, budget, age, pic, src, isSeries, needsPayment ,slug)
VALUES ('9743f488-3b47-11ec-8d3d-0242ac130003', '{Фентези}', 'США', '2001-01-01',
        'История о том, как обычный паренёк из провинции превратился в одного из самых состоятельных людей США. Джордж Джанга (Джонни Депп) с детства мечтал о роскошной жизни: шикарном доме, дорогих автомобилях и лучших девушках США. Находчивый паренёк быстро отыскал оригинальный способ осуществить свою мечту: он решил распространять кокаин в промышленных масштабах. Как найти рынки сбыта? Очень просто — необходимо подсадить на белый порошок главных знаменитостей Америки!',
        'Ирландец', 2015, '{S. Borunov, K. Тарантино}', '{Роберт Уэйд}', '{3e06d4e4-3b47-11ec-8d3d-0242ac130003,
        9743f488-3b47-11ec-8d3d-0242ac130003}', '2001-01-02', 120, 'Русский', '120 млн', 12, '{irish.png}', '{video.mp4}', true, true,'irishman');

INSERT INTO public.series_seasons (series_id, id, pic, src)
VALUES ('9743f488-3b47-11ec-8d3d-0242ac130003', 1, '{warcraft_s1_ser1.png, warcraft_s1_ser2.png, warcraft_s1_ser3.png, warcraft_s1_ser4.png, warcraft_s1_ser5.png}',
        '{video.mp4, video.mp4, video.mp4, video.mp4, video.mp4}');
INSERT INTO public.series_seasons (series_id, id, pic, src)
VALUES ('9743f488-3b47-11ec-8d3d-0242ac130003', 2, '{warcraft_s2_ser1.png, warcraft_s2_ser2.png, warcraft_s2_ser3.png, warcraft_s2_ser4.png, warcraft_s2_ser5.png}',
    '{video.mp4, video.mp4, video.mp4, video.mp4, video.mp4}');

-- rating
INSERT INTO public.rating (film_id, rating) VALUES ('9743f488-3b47-11ec-8d3d-0242ac130003', 4.7);
INSERT INTO public.rating (film_id, rating) VALUES ('f8405178-3b47-11ec-8d3d-0242ac130003', 4.3);
INSERT INTO public.rating (film_id, rating) VALUES ('b0840514-46e6-11ec-81d3-0242ac130003', 4.7);
INSERT INTO public.rating (film_id, rating) VALUES ('ac615b08-46e6-11ec-81d3-0242ac130003', 4.2);
INSERT INTO public.rating (film_id, rating) VALUES ('9dd8cef4-46e6-11ec-81d3-0242ac130003', 4.1);

INSERT INTO public.rating (film_id, rating) VALUES ('68896eac-46e6-11ec-81d3-0242ac130003', 3.3);
INSERT INTO public.rating (film_id, rating) VALUES ('649be5f4-46e6-11ec-81d3-0242ac130003', 2.7);
INSERT INTO public.rating (film_id, rating) VALUES ('6d0191bc-46e6-11ec-81d3-0242ac130003', 1.7);
INSERT INTO public.rating (film_id, rating) VALUES ('7f1ff974-46e6-11ec-81d3-0242ac130003', 6.3);
INSERT INTO public.rating (film_id, rating) VALUES ('a8c67370-46e6-11ec-81d3-0242ac130003', 8.3);
