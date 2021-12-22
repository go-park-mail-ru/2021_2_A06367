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

CREATE TABLE subs
(
    user_id uuid NOT NULL,
    exp_date date NOT NULL,
    CONSTRAINT subs_pkey PRIMARY KEY (user_id)
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



INSERT INTO actors VALUES ('42903802-d80f-4f9f-aa6e-4f3dd08f9821', 'Хилари', 'Дафф', '31505.jpg', 161, '0001-01-01 00:00:00', '', '{мелодрама,приключения,комедия,музыка,семейный}');
INSERT INTO actors VALUES ('8c8e9d7e-f553-4aba-9da2-8ab137ec7ca4', 'Адам', 'Лэмберг', '31506.jpg', 177, '0001-01-01 00:00:00', '', '{мелодрама,приключения,комедия,музыка,семейный}');
INSERT INTO actors VALUES ('ffb54b75-adb4-4cf2-adb0-8ad1ae223d45', 'Киану', 'Ривз', '7836.jpg', 179, '0001-01-01 00:00:00', '', '{фантастика,боевик}');
INSERT INTO actors VALUES ('e5717e55-d822-461a-bca1-a19764d59a06', 'Лоренс', 'Фишбёрн', '9838.jpg', 151, '0001-01-01 00:00:00', '', '{фантастика,боевик}');
INSERT INTO actors VALUES ('0d4b2616-f636-4457-9e21-4ffe7b828149', 'Сигурни', 'Уивер', '6915.jpg', 175, '0001-01-01 00:00:00', '', '{драма,детектив,приключения,комедия,семейный}');
INSERT INTO actors VALUES ('313598fb-6eae-4cb7-9629-ccf10aea89f4', 'Джон', 'Войт', '515.jpg', 170, '0001-01-01 00:00:00', '', '{драма,детектив,приключения,комедия,семейный}');
INSERT INTO actors VALUES ('a6909e2d-8fc8-43f0-ad7f-e0ae11c75964', 'Адам', 'Сэндлер', '7627.jpg', 150, '0001-01-01 00:00:00', '', '{мелодрама,комедия}');
INSERT INTO actors VALUES ('942a7956-e666-4ab9-a6b9-1616c8092fe1', 'Джек', 'Николсон', '30056.jpg', 164, '0001-01-01 00:00:00', '', '{мелодрама,комедия}');
INSERT INTO actors VALUES ('3b502535-e074-4cd3-87fb-70314f2e5988', 'Эдвард', 'Бёрнс', '719.jpg', 162, '0001-01-01 00:00:00', '', '{триллер,криминал}');
INSERT INTO actors VALUES ('dfec6bf7-d5f0-4ad1-ae43-893ae9793c99', 'Рэйчел', 'Вайс', '12586.jpg', 179, '0001-01-01 00:00:00', '', '{триллер,криминал}');
INSERT INTO actors VALUES ('2704d03b-bfa9-42e2-9895-2144809d39de', 'Колин', 'Фаррелл', '373.jpg', 164, '0001-01-01 00:00:00', '', '{триллер,криминал}');
INSERT INTO actors VALUES ('72dd2b7d-f90b-492b-bfa9-42dd63cc86f9', 'Кифер', 'Сазерленд', '7131.jpg', 171, '0001-01-01 00:00:00', '', '{триллер,криминал}');
INSERT INTO actors VALUES ('7d4e974a-2af6-4c41-b760-23225e8fef4c', 'Аль', 'Пачино', '26240.jpg', 167, '0001-01-01 00:00:00', '', '{триллер,боевик}');
INSERT INTO actors VALUES ('7efcb4cc-77a0-4b99-8847-b30ac9c14fcb', 'Колин', 'Фаррелл', '373.jpg', 176, '0001-01-01 00:00:00', '', '{триллер,боевик}');
INSERT INTO actors VALUES ('b310248d-f398-4199-a63c-b2dc79263ea7', 'Джейми', 'Кеннеди', '1576.jpg', 156, '0001-01-01 00:00:00', '', '{криминал,комедия}');
INSERT INTO actors VALUES ('015f5813-5e82-4747-b599-bc75a9757223', 'Тэй', 'Диггз', '7090.jpg', 158, '0001-01-01 00:00:00', '', '{криминал,комедия}');
INSERT INTO actors VALUES ('94f31064-c88d-497a-b03a-f818768d408f', 'Эрик', 'Кристиан', '26880.jpg', 177, '0001-01-01 00:00:00', '', '{комедия}');
INSERT INTO actors VALUES ('02f3dfcb-235a-4628-a4a1-2a3577109305', 'Дерек', 'Ричардсон', '30160.jpg', 157, '0001-01-01 00:00:00', '', '{комедия}');
INSERT INTO actors VALUES ('328db934-874e-4283-ab1d-1da4c12ea012', 'Кристиан', 'Бэйл', '21495.jpg', 178, '0001-01-01 00:00:00', '', '{триллер,драма,фантастика,боевик}');
INSERT INTO actors VALUES ('1f1df240-f513-40ed-acd1-fc8dbd3e8a4e', 'Тэй', 'Диггз', '7090.jpg', 150, '0001-01-01 00:00:00', '', '{триллер,драма,фантастика,боевик}');

INSERT INTO actors VALUES ('d8f2012f-3801-48cf-a824-6ed716c7165b', 'Эдди', 'Мёрфи', '25674.jpg', 151, '0001-01-01 00:00:00', '', '{комедия,семейный}');
INSERT INTO actors VALUES ('cd6f9ff4-6573-4314-b14c-329cd7be270d', 'Стив', 'Зан', '10083.jpg', 168, '0001-01-01 00:00:00', '', '{комедия,семейный}');
INSERT INTO actors VALUES ('7ae1466a-e35f-4309-8fa2-978998ee48fe', 'Чоу', 'Юнь-Фат', '10919.jpg', 161, '0001-01-01 00:00:00', '', '{мелодрама,боевик,фэнтези,комедия}');
INSERT INTO actors VALUES ('d14802a0-fd9d-4d4f-9fc6-ab6f8f8845ca', 'Шонн', 'Уильям', '297.jpg', 159, '0001-01-01 00:00:00', '', '{мелодрама,боевик,фэнтези,комедия}');
INSERT INTO actors VALUES ('e45797cf-4427-4903-99a7-595e4e6d2ad1', 'Элайджа', 'Вуд', '20287.jpg', 177, '0001-01-01 00:00:00', '', '{драма,приключения,фэнтези}');
INSERT INTO actors VALUES ('b7b1ec06-9e60-472c-828a-c1e7c9090db5', 'Иэн', 'Маккеллен', '8215.jpg', 151, '0001-01-01 00:00:00', '', '{драма,приключения,фэнтези}');
INSERT INTO actors VALUES ('0e66ce4b-cf81-4f85-81a5-245081d26b41', 'Парминдер', 'К.', '37287.jpg', 156, '0001-01-01 00:00:00', '', '{драма,мелодрама,комедия,спорт}');
INSERT INTO actors VALUES ('48d5a192-fbae-4423-9a48-441a37c5bca8', 'Кира', 'Найтли', '24302.jpg', 163, '0001-01-01 00:00:00', '', '{драма,мелодрама,комедия,спорт}');
INSERT INTO actors VALUES ('3806b3e1-3e16-4464-9c0c-a2c15387c449', 'Роб', 'Шнайдер', '7642.jpg', 154, '0001-01-01 00:00:00', '', '{фэнтези,комедия}');
INSERT INTO actors VALUES ('6c14e6ac-6372-43f2-bf90-385e18cb2c69', 'Анна', 'Фэрис', '22471.jpg', 153, '0001-01-01 00:00:00', '', '{фэнтези,комедия}');
INSERT INTO actors VALUES ('fe6ba91c-7f3a-4ab6-93e7-fa031a626bc5', 'Киану', 'Ривз', '7836.jpg', 177, '0001-01-01 00:00:00', '', '{фантастика,боевик}');
INSERT INTO actors VALUES ('9ce13ada-31ba-48f3-9ef8-535c1849d6b9', 'Хьюго', 'Уивинг', '1491.jpg', 178, '0001-01-01 00:00:00', '', '{фантастика,боевик}');
INSERT INTO actors VALUES ('d4adbae3-7487-41f4-9acd-4cd21a216cfb', 'Сандра', 'Буллок', '6542.jpg', 169, '0001-01-01 00:00:00', '', '{мелодрама,комедия}');
INSERT INTO actors VALUES ('42c7e022-7590-4da5-858a-e913e20a89fa', 'Хью', 'Грант', '8090.jpg', 153, '0001-01-01 00:00:00', '', '{мелодрама,комедия}');
INSERT INTO actors VALUES ('ef274fb8-4d57-4a2f-a6b8-d16c6937b8cd', 'Шарлиз', 'Терон', '730.jpg', 151, '0001-01-01 00:00:00', '', '{триллер,криминал,боевик}');
INSERT INTO actors VALUES ('e8c39d1f-4dc8-4cc4-8429-ee551cf3e9de', 'Марк', 'Уолберг', '21038.jpg', 159, '0001-01-01 00:00:00', '', '{триллер,криминал,боевик}');
INSERT INTO actors VALUES ('078e10be-09eb-4046-9d9c-17bb72a08bf3', 'Арнольд', 'Шварценеггер', '6264.jpg', 160, '0001-01-01 00:00:00', '', '{фантастика,боевик}');
INSERT INTO actors VALUES ('46d71ef1-7903-4b3c-91b0-2848ef0bb750', 'Ник', 'Стал', '13297.jpg', 175, '0001-01-01 00:00:00', '', '{фантастика,боевик}');


INSERT INTO films VALUES ('e16fc41a-dce8-44fa-bba6-d9af1eecfebd', '{мелодрама,приключения,комедия,музыка,семейный}', 'США', '2003-01-01 03:00:00+03', 'Лиззи Магуайр', 2003, '{"Джим Фолл"}', '{"Сьюзэн Эстель Джэнсен","Эд Дектер","Джон Дж. Штраусс","Терри Мински"}', '{42903802-d80f-4f9f-aa6e-4f3dd08f9821,8c8e9d7e-f553-4aba-9da2-8ab137ec7ca4}', '2003-01-01', 94, 'Ru', '???', 12, '{lizzi-maguair.png}', '{lizzi-maguair.mp4}', 'Тринадцатилетняя школьница Лиззи Магуайер и ее приятели Гордо, Кейт и Эсан собираются оттянуться по полной программе во время их поездки с классом в Италию.

Но там случается весьма неожиданное происшествие: девочку ошибочно принимают за итальянскую поп-звезду Изабеллу, да к тому же девушка влюбляется в бывшего дружка Изабеллы Паоло. Когда родители Лизи обо всем узнают, они вместе с ее братом Мэттом срочно вылетают в Италию.

Но Лиззи уже не та закомплексованная девочка-подросток, кем была раньше, она до такой степени вжилась в роль певицы, что и на самом деле стала самой настоящей звездой.', false, true, 'lizzi-maguair');
INSERT INTO films VALUES ('9d574931-e719-44fc-bad2-69a1de36d00e', '{фантастика,боевик}', 'США', '1999-01-01 03:00:00+03', 'Матрица', 1999, '{"Лана Вачовски","Лилли Вачовски"}', '{"Лилли Вачовски","Лана Вачовски"}', '{ffb54b75-adb4-4cf2-adb0-8ad1ae223d45,e5717e55-d822-461a-bca1-a19764d59a06}', '1999-01-01', 136, 'Ru', '???', 16, '{matritsa.png}', '{matritsa.mp4}', 'Жизнь Томаса Андерсона разделена на две части: днём он — самый обычный офисный работник, получающий нагоняи от начальства, а ночью превращается в хакера по имени Нео, и нет места в сети, куда он бы не смог проникнуть. Но однажды всё меняется. Томас узнаёт ужасающую правду о реальности.', false, false, 'matritsa');
INSERT INTO films VALUES ('f0285cb7-d7c7-48d7-b94a-1bdf62469553', '{драма,детектив,приключения,комедия,семейный}', 'США', '2003-01-01 03:00:00+03', 'Клад', 2003, '{"Эндрю Дэвис"}', '{"Луис Сачар"}', '{0d4b2616-f636-4457-9e21-4ffe7b828149,313598fb-6eae-4cb7-9629-ccf10aea89f4}', '2003-01-01', 117, 'Ru', '???', 12, '{klad.png}', '{klad.mp4}', 'Стэнли арестован по ложному обвинению в краже кроссовок и отправлен в трудовой лагерь, расположенный в техасской пустыне. Воспитатели «закаляют характер» подростков странным наказанием. Ребята копают ямы в иссушенной земле, но не знают, что их на самом деле используют для раскопок таинственного клада. Однако Стэнли удается раскрыть загадочную связь между сокровищами и проклятием, тяготеющим долгие годы над его семьей…', false, false, 'klad');
INSERT INTO films VALUES ('c8ae08e9-ee3d-4b59-9bdf-af456de7f97a', '{мелодрама,комедия}', 'США', '2003-01-01 03:00:00+03', 'Управление гневом', 2003, '{"Питер Сигал"}', '{"Дэвид Дорфман"}', '{a6909e2d-8fc8-43f0-ad7f-e0ae11c75964,942a7956-e666-4ab9-a6b9-1616c8092fe1}', '2003-01-01', 106, 'Ru', '???', 12, '{upravlenie-gnevom.png}', '{upravlenie-gnevom.mp4}', 'Скромному клерку отчаянно не везет. Парня по обвинению в нападении на бортпроводницу приговаривают к лечению у психиатра. Но верно говорят, что большинство психиатров сами немного безумны. Или сильно не в себе...', false, true, 'upravlenie-gnevom');
INSERT INTO films VALUES ('cda6387e-a256-4f85-96cc-3cb27343fcc4', '{триллер,криминал}', 'США', '2003-01-01 03:00:00+03', 'Афера', 2003, '{"Джеймс Фоули"}', '{"Даг Джанг"}', '{3b502535-e074-4cd3-87fb-70314f2e5988,dfec6bf7-d5f0-4ad1-ae43-893ae9793c99}', '2003-01-01', 97, 'Ru', '???', 16, '{afera.png}', '{afera.mp4}', 'Джейк Вига – хитроумный и обаятельный мошенник. Последняя афера Джейка привела к тому, что его дорожки пересеклись с мафией – при помощи своей команды он лишил нескольких тысяч долларов Лайонела Долби, – счетовода эксцентричного мафиозного босса Уинстона Кинга по прозвищу «Король».

Мафия шутить не любит, а любит выбивать долги из тех, кто пытается ее надуть. Чтобы сохранить жизнь и расквитаться с долгами, Джейку приходится устроить новую, еще более изощренную аферу – сложнейшую схему, в которой требуется «творческий подход к бухгалтерии».

Неожиданные помехи появляются одна за другой: палки в колеса Джейку вставляют его старый враг, агент ФБР Гюнтер Бутан, Трэвис, правая рука «Короля» и хитроумная карманщица Лили на которую Джейк успел положить глаз...', false, false, 'afera');
INSERT INTO films VALUES ('3100072a-e1e5-44c6-9a10-77be314dd492', '{триллер,криминал}', 'США', '2002-01-01 03:00:00+03', 'Телефонная будка', 2002, '{"Джоэл Шумахер"}', '{"Ларри Коэн"}', '{2704d03b-bfa9-42e2-9895-2144809d39de,72dd2b7d-f90b-492b-bfa9-42dd63cc86f9}', '2002-01-01', 81, 'Ru', '???', 16, '{telefonnaia-budka.png}', '{telefonnaia-budka.mp4}', 'Один телефонный звонок может изменить всю жизнь человека или даже оборвать ее. Герой фильма Стью Шеферд становится пленником телефонной будки.

Что вы сделаете, если услышите, как в телефонной будке зазвонил телефон? Скорее всего, инстинктивно поднимете трубку, хотя прекрасно знаете, что кто-то просто ошибся номером. Вот и Стью кажется, что на звонок надо обязательно ответить, а в результате он оказывается втянутым в чудовищную игру. «Только положи трубку, и ты – труп», – говорит ему невидимый собеседник.', false, false, 'telefonnaia-budka');
INSERT INTO films VALUES ('335b9fdd-64e9-4403-ba24-7269ca8a5a52', '{триллер,боевик}', 'США', '2003-01-01 03:00:00+03', 'Рекрут', 2003, '{"Роджер Дональдсон"}', '{"Роджер Таун","Курт Уиммер","Митч Глейзер"}', '{7d4e974a-2af6-4c41-b760-23225e8fef4c,7efcb4cc-77a0-4b99-8847-b30ac9c14fcb}', '2003-01-01', 115, 'Ru', '???', 16, '{rekrut.png}', '{rekrut.mp4}', 'Джеймс Клэйтон - студент и опытный хакер. Он привлекает внимание спецслужб, и его вербуют в ЦРУ, упоминая таинственное исчезновение его отца в 90-х. Джеймс обучается у Уолтера Бёрка. Он отлично сдает все тесты, кроме последнего. Так он становится агентом без прикрытия и получает задание: найти «крота», который похищает опасный вирус из Лэнгли.', false, true, 'rekrut');
INSERT INTO films VALUES ('2a149c42-d48e-4ad6-a537-e95c19b0a4fd', '{криминал,комедия}', 'США', '2003-01-01 03:00:00+03', 'Разыскиваются в Малибу', 2003, '{"Джон Уайтселл"}', '{"Факс Бар","Адам Смолл","Джейми Кеннеди","Ник Свардсон"}', '{b310248d-f398-4199-a63c-b2dc79263ea7,015f5813-5e82-4747-b599-bc75a9757223}', '2003-01-01', 86, 'Ru', '???', 12, '{razyskivaiutsia-v-malibu.png}', '{razyskivaiutsia-v-malibu.mp4}', 'Не стоит ненавидеть его за то, кем он является. Или кем он не является. Все, чем хочет заниматься Брэд - или Би-рэд - разъезжать со своими дружками по Малибу и вести себя, как самый крутой черный рэппер в округе.

Но все вокруг знают, что паренек, лихо читающий рэп, родом из респектабельного квартала Малибу. И отец Брэда, Билл Глакман, кандидат в губернаторы, серьезно боится, что увлечение Би-рэда «черной» культурой может разнести в пух и прах его предвыборную кампанию.', false, false, 'razyskivaiutsia-v-malibu');
INSERT INTO films VALUES ('ddd73f7f-3ae7-467f-a79b-f3465485a8bf', '{комедия}', 'США', '2003-01-01 03:00:00+03', 'Тупой и еще тупее тупого: Когда Гарри встретил Ллойда', 2003, '{"Трой Миллер"}', '{"Питер Фаррелли","Беннетт Йеллин","Дэнни Байерс","Бобби Фаррелли","Роберт Бренер","Трой Миллер"}', '{94f31064-c88d-497a-b03a-f818768d408f,02f3dfcb-235a-4628-a4a1-2a3577109305}', '2003-01-01', 85, 'Ru', '???', 16, '{tupoi-i-eshche-tupee-tupogo:-kogda-garri-vstretil-lloida.png}', '{tupoi-i-eshche-tupee-tupogo:-kogda-garri-vstretil-lloida.mp4}', 'Как же встретились два героя-недоумка Гарри и Ллойд, известные по фильму «Тупой и еще тупее»? Оба несколько лет не ходили в школу, а учились на дому. Пришло время отправляться в школу, и прямо на улице Гарри и Ллойд столкнулись лбами...

А в это время школьный директор Коллинз и мисс Хеллер, буфетчица-официантка школьной столовой, задумали провернуть аферу - получить благотворительную премию в сто тысяч долларов за организацию класса для умственно отсталых. Злоумышленники решили создать липовый класс. Два кандидата в спецкласс нашлись сразу - Гарри и Ллойд, а они разыскали остальных...', false, false, 'tupoi-i-eshche-tupee-tupogo:-kogda-garri-vstretil-lloida');
INSERT INTO films VALUES ('fbf17d95-f8bb-4ed7-b385-5111c321cbdc', '{триллер,драма,фантастика,боевик}', 'США', '2002-01-01 03:00:00+03', 'Эквилибриум', 2002, '{"Курт Уиммер"}', '{"Курт Уиммер"}', '{328db934-874e-4283-ab1d-1da4c12ea012,1f1df240-f513-40ed-acd1-fc8dbd3e8a4e}', '2002-01-01', 107, 'Ru', '???', 16, '{ekvilibrium.png}', '{ekvilibrium.mp4}', 'В будущем люди лишены возможности выражать эмоции. Это цена, которую человечество платит за устранение из своей жизни войны. Теперь книги, искусство и музыка находятся вне закона, а любое чувство — преступление, наказуемое смертью.

Для приведения в жизнь существующего правила используется принудительное применение лекарства прозиум. Правительственный агент Джон Престон борется с теми, кто нарушает правила. В один прекрасный момент он забывает принять очередную дозу лекарства, и с ним происходит духовное преображение, что приводит его к конфликту не только с режимом, но и с самим собой.', false, true, 'ekvilibrium');
INSERT INTO films VALUES ('36fe5b42-da42-40a7-a554-985c6351f310', '{комедия,семейный}', 'США', '2003-01-01 03:00:00+03', 'Дежурный папа', 2003, '{"Стив Карр"}', '{"Джефф Родкей"}', '{d8f2012f-3801-48cf-a824-6ed716c7165b,cd6f9ff4-6573-4314-b14c-329cd7be270d}', '2003-01-01', 89, 'Ru', '???', 6, '{dezhurnyi-papa.png}', '{dezhurnyi-papa.mp4}', 'Чарли и Фила увольняют с работы в крупной корпорации. Теперь им самим придётся сидеть со своими сыновьями, так как оплачивать счета дорогостоящего детского центра уже не на что. Промучившись пару недель со своими отпрысками, папаши так увлекаются этим делом, что решают поставить дело на деловые рельсы и открывают новый центр дневного пребывания для детей.

Чарли и Филл находят всё новые и новые нетрадиционные и забавные способы воспитательного воздействия на малышей, и центр «Дежурный папа» становится всё более популярным. Почувствовав жёсткое соперничество со стороны усатых няней, директриса дорогого детского центра решает выжить конкурентов-новичков из бизнеса.', false, false, 'dezhurnyi-papa');
INSERT INTO films VALUES ('55e547e7-c344-4990-9952-8bcf94661f0b', '{мелодрама,боевик,фэнтези,комедия}', 'США', '2003-01-01 03:00:00+03', 'Пуленепробиваемый', 2003, '{"Пол Хантер"}', '{"Этан Райфф","Сайрус Ворис"}', '{7ae1466a-e35f-4309-8fa2-978998ee48fe,d14802a0-fd9d-4d4f-9fc6-ab6f8f8845ca}', '2003-01-01', 104, 'Ru', '???', 12, '{puleneprobivaemyi.png}', '{puleneprobivaemyi.mp4}', 'Монах - мастер боевых искусств, который охраняет могущественный древний свиток - таинственный артефакт, содержащий ключ к безграничной власти. Монаху нужно найти следующего хранителя свитка, и поиски приводят его в Америку. Согласно древнему пророчеству и к изумлению Монаха его преемником оказывается обаятельный хулиган по имени Кар.

Пока Монах обучает Кара, как выполнять свою работу, этому невероятному дуэту приходится защищать свиток от одержимого жаждой власти человека, который гоняется за ним 60 лет.

В суматохе невероятной акробатики, схваток с применением боевых искусств и остроумных шуток эта комическая странная парочка должна сделать всё, чтобы свиток не попал в руки злодея.', false, false, 'puleneprobivaemyi');
INSERT INTO films VALUES ('ba7c44b4-1fac-48f2-832f-200fd18dfe8c', '{драма,приключения,фэнтези}', 'США', '2002-01-01 03:00:00+03', 'Властелин колец: Две крепости', 2002, '{"Питер Джексон"}', '{"Дж.Р.Р. Толкин","Фрэн Уолш","Филиппа Бойенс","Стивен Синклер","Питер Джексон"}', '{e45797cf-4427-4903-99a7-595e4e6d2ad1,b7b1ec06-9e60-472c-828a-c1e7c9090db5}', '2002-01-01', 179, 'Ru', '???', 12, '{vlastelin-kolets:-dve-kreposti.png}', '{vlastelin-kolets:-dve-kreposti.mp4}', 'Братство распалось, но Кольцо Всевластья должно быть уничтожено. Фродо и Сэм вынуждены доверится Голлуму, который взялся провести их к вратам Мордора. Громадная армия Сарумана приближается: члены братства и их союзники готовы принять бой. Битва за Средиземье продолжается.', false, true, 'vlastelin-kolets:-dve-kreposti');
INSERT INTO films VALUES ('a377a105-0b4a-40b7-a6e1-f31bf1b5c8e4', '{драма,мелодрама,комедия,спорт}', 'США', '2002-01-01 03:00:00+03', 'Играй, как Бекхэм', 2002, '{"Гуриндер Чадха"}', '{"Гуриндер Чадха","Гюльджит Биндра","Пол Маеда Берджес"}', '{0e66ce4b-cf81-4f85-81a5-245081d26b41,48d5a192-fbae-4423-9a48-441a37c5bca8}', '2002-01-01', 112, 'Ru', '???', 16, '{"igrai,-kak-bekkhem.png"}', '{"igrai,-kak-bekkhem.mp4"}', 'Джесс только 18, но она твердо знает, что сделает ее счастливой: футбольная карьера, такая же, как у знаменитого игрока «Манчестер Юнайтед» Дэвида Бекхема. Пока она гоняет мяч в лондонском парке, доказывая соседским мальчишкам, что девушки играют в футбол не хуже, ее родители и многочисленные родственники, как и положено традиционной индийской семье, подыскивают для нее достойного мужа и строят планы о ее будущей карьере юриста.

Однажды, во время очередной разминки, Джесс знакомится с Джулз и та приглашает ее на тренировку в женскую футбольную секцию. С замиранием сердца она следует за своей новой приятельницей, и мечта становится реальностью: Джесс принимают в настоящую футбольную команду. В довершение ко всему она влюбляется в своего тренера Джо.

Понимая, что семья никогда не примирится с ее выбором, Джесс наслаждается своим кратковременным счастьем и готовится отстаивать свои права до конца...', false, false, 'igrai,-kak-bekkhem');
INSERT INTO films VALUES ('82879482-9ceb-492a-8333-80faab5edd64', '{фэнтези,комедия}', 'США', '2002-01-01 03:00:00+03', 'Цыпочка', 2002, '{"Том Брэди"}', '{"Том Брэди","Роб Шнайдер"}', '{3806b3e1-3e16-4464-9c0c-a2c15387c449,6c14e6ac-6372-43f2-bf90-385e18cb2c69}', '2002-01-01', 104, 'Ru', '???', 16, '{tsypochka.png}', '{tsypochka.mp4}', 'Популярная, но неприятная в общении старшеклассница Джессика однажды утром просыпается в теле 30-летнего мужчины с не самой привлекательной внешностью. Девушка отправляется на поиски своего тела, и это приключение помогает ей увидеть себя со стороны и понять, насколько поверхностной и недалёкой она была.', false, false, 'tsypochka');
INSERT INTO films VALUES ('8d7db471-6810-4045-a841-fa47a72f5242', '{фантастика,боевик}', 'США', '2003-01-01 03:00:00+03', 'Матрица: Революция', 2003, '{"Лана Вачовски","Лилли Вачовски"}', '{"Лилли Вачовски","Лана Вачовски"}', '{fe6ba91c-7f3a-4ab6-93e7-fa031a626bc5,9ce13ada-31ba-48f3-9ef8-535c1849d6b9}', '2003-01-01', 129, 'Ru', '???', 16, '{matritsa:-revoliutsiia.png}', '{matritsa:-revoliutsiia.mp4}', 'Пока армия Машин пытается уничтожить Зион, его жители из последних сил держат оборону. Но удастся ли им предотвратить полное вторжение в город кишащей орды беспощадных машин до того, как Нео соберет все свои силы и положит конец войне?', false, false, 'matritsa:-revoliutsiia');
INSERT INTO films VALUES ('6c6abded-9206-44b8-b149-92f08569ca8d', '{мелодрама,комедия}', 'США', '2002-01-01 03:00:00+03', 'Любовь с уведомлением', 2002, '{"Марк Лоуренс"}', '{"Марк Лоуренс"}', '{d4adbae3-7487-41f4-9acd-4cd21a216cfb,42c7e022-7590-4da5-858a-e913e20a89fa}', '2002-01-01', 101, 'Ru', '???', 16, '{liubov-s-uvedomleniem.png}', '{liubov-s-uvedomleniem.mp4}', 'Джордж Уэйд и шага не может сделать без Люси Келсон, работающей главным консультантом в его корпорации. Однако обращается он с ней скорее как с няней, а не как с блестящим юристом, окончившим Гарвард. По прошествии года все это надоедает Люси, и она решает уволиться. Джордж соглашается ее отпустить, но с одним условием — она должна найти себе достойную замену...', false, false, 'liubov-s-uvedomleniem');
INSERT INTO films VALUES ('c2e011a5-42cc-40f5-97a3-4dc430be3c49', '{триллер,криминал,боевик}', 'США', '2003-01-01 03:00:00+03', 'Ограбление по-итальянски', 2003, '{"Ф. Гэри Грей"}', '{"Трой Кеннеди-Мартин","Донна Пауэрс","Уэйн Пауэрс"}', '{ef274fb8-4d57-4a2f-a6b8-d16c6937b8cd,e8c39d1f-4dc8-4cc4-8429-ee551cf3e9de}', '2003-01-01', 111, 'Ru', '???', 12, '{ograblenie-po-italianski.png}', '{ograblenie-po-italianski.mp4}', 'Джон Бриджер всегда умел спланировать идеальное ограбление. Вместе со своей командой опытных бандитов он провернул не одно дело, но теперь решил уйти на покой. Впереди у Бриджера последнее задание: кража золотых слитков, в которой принимают участие инсайдер Стив, водитель Роб, взрыватель Левое ухо, технарь Лайл и Чарли - верный друг Бриджера и второй «планировщик» в их команде. Ограбление, изящное и быстрое, было разыграно как по нотам, но после его завершения веселье преступников было омрачено предательством...', false, true, 'ograblenie-po-italianski');
INSERT INTO films VALUES ('1dfc3231-c9c9-4962-8954-01da59c5dcdd', '{фантастика,боевик}', 'США', '2003-01-01 03:00:00+03', 'Терминатор 3: Восстание машин', 2003, '{"Джонатан Мостоу"}', '{"Джеймс Кэмерон","Джон Бренкето","Майкл Феррис","Гейл Энн Хёрд","","Теди Сарафьян"}', '{078e10be-09eb-4046-9d9c-17bb72a08bf3,46d71ef1-7903-4b3c-91b0-2848ef0bb750}', '2003-01-01', 109, 'Ru', '???', 16, '{terminator-3:-vosstanie-mashin.png}', '{terminator-3:-vosstanie-mashin.mp4}', 'Прошло десять лет с тех пор, как Джон Коннор помог предотвратить Судный День и спасти человечество от массового уничтожения. Теперь ему 25, Коннор не живет «как все» - у него нет дома, нет кредитных карт, нет сотового телефона и никакой работы.

Его существование нигде не зарегистрировано. Он не может быть прослежен системой Skynet - высокоразвитой сетью машин, которые когда-то попробовали убить его и развязать войну против человечества. Пока из теней будущего не появляется T-X - Терминатрикс, самый сложный киборг-убийца Skynet.

Посланная назад сквозь время, чтобы завершить работу, начатую её предшественником, T-1000, эта машина так же упорна, как прекрасен её человеческий облик. Теперь единственная надежда Коннору выжить - Терминатор, его таинственный прежний убийца. Вместе они должны одержать победу над новыми технологиями T-X и снова предотвратить Судный День...', false, false, 'terminator-3:-vosstanie-mashin');


INSERT INTO rating VALUES ('e16fc41a-dce8-44fa-bba6-d9af1eecfebd', 4);
INSERT INTO rating VALUES ('9d574931-e719-44fc-bad2-69a1de36d00e', 1);
INSERT INTO rating VALUES ('f0285cb7-d7c7-48d7-b94a-1bdf62469553', 5);
INSERT INTO rating VALUES ('c8ae08e9-ee3d-4b59-9bdf-af456de7f97a', 1);
INSERT INTO rating VALUES ('cda6387e-a256-4f85-96cc-3cb27343fcc4', 2);
INSERT INTO rating VALUES ('3100072a-e1e5-44c6-9a10-77be314dd492', 4);
INSERT INTO rating VALUES ('335b9fdd-64e9-4403-ba24-7269ca8a5a52', 9);
INSERT INTO rating VALUES ('2a149c42-d48e-4ad6-a537-e95c19b0a4fd', 5);
INSERT INTO rating VALUES ('ddd73f7f-3ae7-467f-a79b-f3465485a8bf', 8);
INSERT INTO rating VALUES ('fbf17d95-f8bb-4ed7-b385-5111c321cbdc', 1);
INSERT INTO rating VALUES ('36fe5b42-da42-40a7-a554-985c6351f310', 8);
INSERT INTO rating VALUES ('55e547e7-c344-4990-9952-8bcf94661f0b', 5);
INSERT INTO rating VALUES ('ba7c44b4-1fac-48f2-832f-200fd18dfe8c', 8);
INSERT INTO rating VALUES ('a377a105-0b4a-40b7-a6e1-f31bf1b5c8e4', 9);
INSERT INTO rating VALUES ('82879482-9ceb-492a-8333-80faab5edd64', 3);
INSERT INTO rating VALUES ('8d7db471-6810-4045-a841-fa47a72f5242', 2);
INSERT INTO rating VALUES ('6c6abded-9206-44b8-b149-92f08569ca8d', 5);
INSERT INTO rating VALUES ('c2e011a5-42cc-40f5-97a3-4dc430be3c49', 9);
INSERT INTO rating VALUES ('1dfc3231-c9c9-4962-8954-01da59c5dcdd', 8);

