package repo

import (
	"context"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/models"
	"github.com/jackc/pgtype/pgxtype"
	"go.uber.org/zap"
	"math/rand"
	"strings"
)

const (
	GET_ID_BY_SLUG = "SELECT id FROM films where slug=$1;"

	SET_RATING = "INSERT INTO ratings(id, film_id, rating)" +
		"values($1, $2, $3) on conflict(id, film_id) do update set rating = $3;"

	GET_RATING = "SELECT AVG(rating) FROM ratings WHERE film_id=$1;"

	GET_RATING_BY_USER = "SELECT rating FROM ratings WHERE id=$1 AND film_id=$2"

	SElECT_FILM_BY_TOPIC = "SELECT id, genres, country, releaseRus, title, year, director, " +
		"authors, actors, release, duration, language, budget, age, pic, src, description, isSeries, needsPayment, slug " +
		"FROM films " +
		"WHERE $1 = ANY(genres)"

	SElECT_RANDOM_COUNT = "SELECT COUNT(id) FROM films"

	SElECT_RANDOM_FILM = "SELECT id, genres, country, releaseRus, title, year, director, " +
		"authors, actors, release, duration, language, budget, age, pic, src, description, isSeries, needsPayment, slug " +
		"FROM films " +
		"LIMIT 1 OFFSET $1"

	SELECT_FILM_BY_RATING = "SELECT id, genres, country, releaseRus, title, year, director, authors, actors, release, " +
		"duration, language, budget, age, pic, src, description, isSeries, needsPayment, slug " +
		" FROM films JOIN rating ON films.id = rating.film_id ORDER BY rating DESC LIMIT 10"

	SELECT_FILM_BY_DATE = "SELECT id, genres, country, releaseRus, title, year, director, authors, actors, release, " +
		"duration, language, budget, age, pic, src, description, isSeries, needsPayment, slug " +
		"FROM films ORDER BY release DESC LIMIT 10"

	SELECT_FILM_BY_KEYWORD = "SELECT *" +
		"FROM films " +
		"WHERE make_tsvector(title) @@ to_tsquery($1) or LOWER(title) like LOWER($2)  LIMIT 10"

	SELECT_FILM_BY_ID = "SELECT id, genres, country, releaseRus, title, year, director, " +
		"authors, actors, release, duration, language, budget, age, pic, src, description, isSeries, needsPayment, slug " +
		"FROM films " +
		"WHERE id=$1"

	SELECT_FILM_BY_USER = "SELECT f.id, f.genres, f.country, f.releaseRus, f.title, f.year, f.director, " +
		"f.authors, f.actors , f.release, f.duration, f.language, f.budget, f.age, f.pic, f.src, f.description, f.isSeries, f.needsPayment, f.slug " +
		"FROM films f INNER JOIN watchlist w ON f.id=w.film_id " +
		"WHERE w.id=$1"

	SELECT_FILM_BY_ACTOR = "SELECT id, genres, country, releaseRus, title, year, director, " +
		"authors, actors, release, duration, language, budget, age, pic, src, description, isSeries, needsPayment, slug " +
		"FROM films " +
		"WHERE $1=ANY(actors)"

	INSERT_FILM_TO_STARRED = "INSERT INTO starred_films (film_id, user_id) VALUES($1, $2);"

	DELETE_FILM_FROM_STARRED = "DELETE FROM starred_films WHERE film_id=$1 AND user_id=$2;"

	GET_STARRED_FILMS = "SELECT id, genres, country, releaseRus, title, year, director, authors, actors, release, " +
		"duration, language, budget, age, pic, src, description, isSeries, needsPayment, slug " +
		"FROM films f JOIN starred_films sf ON f.id  = sf.film_id " +
		"WHERE sf.user_id=$1"

	GET_IF_STARRED_FILMS = "SELECT id, genres, country, releaseRus, title, year, director, authors, actors, " +
		"release, duration, language, budget, age, pic, src, description, isSeries, needsPayment, slug " +
		"FROM films f JOIN starred_films sf ON f.id  = sf.film_id " +
		"WHERE sf.user_id=$1 AND id=$2"

	GET_IF_WL_FILMS = "SELECT f.id, genres, country, releaseRus, title, year, director, authors, actors, " +
		"release, duration, language, budget, age, pic, src, description, isSeries, needsPayment, slug " +
		"FROM films f JOIN watchlist w ON f.id  = w.film_id " +
		"WHERE w.id=$1 AND w.film_id = $2"

	INSERT_FILM_TO_WATCHLIST = "INSERT INTO watchlist (id, film_id) VALUES($1, $2);"

	DELETE_FILM_FROM_WATCHLIST = "DELETE FROM watchlist WHERE film_id=$1 AND id=$2;"

	GET_WATCHLIST_FILMS = "SELECT f.id, genres, country, releaseRus, title, year, director," +
		" authors, actors, release, duration, language, budget, age, pic, src, description, isSeries, needsPayment, slug " +
		"FROM films f JOIN watchlist w ON f.id  = w.film_id " +
		"WHERE w.id=$1"

	GET_SERIES = "SELECT id, pic, src FROM public.series_seasons WHERE series_id=$1;"
)

type FilmsRepo struct {
	pool   pgxtype.Querier
	logger *zap.SugaredLogger
}

func NewFilmsRepo(pool pgxtype.Querier, logger *zap.SugaredLogger) *FilmsRepo {
	return &FilmsRepo{
		pool:   pool,
		logger: logger,
	}
}

func (r *FilmsRepo) GetFilmsByTopic(topic string) ([]models.Film, models.StatusCode) {

	rows, err := r.pool.Query(context.Background(), SElECT_FILM_BY_TOPIC,
		topic)
	if err != nil {
		return nil, models.InternalError
	}
	defer rows.Close()
	films := make([]models.Film, 0)

	for rows.Next() {
		var film models.Film
		err = rows.Scan(&film.Id, &film.Genres, &film.Country, &film.ReleaseRus, &film.Title,
			&film.Year, &film.Director, &film.Authors, &film.Actors, &film.Release, &film.Duration,
			&film.Language, &film.Budget, &film.Age, &film.Pic, &film.Src, &film.Description, &film.IsSeries, &film.NeedsPayment, &film.Slug)
		if err != nil {
			return nil, models.InternalError
		}

		if film.IsSeries {
			code := r.ReadSeries(&film)
			if code != models.Okey {
				return nil, models.InternalError
			}
		}

		f, code := r.GetRating(film)
		if code == models.NotFound {
			f.Rating = 0
		}
		film.Rating = f.Rating
		films = append(films, film)
	}

	return films, models.Okey
}

func (r *FilmsRepo) GetHottestFilms() ([]models.Film, models.StatusCode) {

	rows, err := r.pool.Query(context.Background(), SELECT_FILM_BY_RATING)
	if err != nil {
		return nil, models.InternalError
	}
	defer rows.Close()

	films := make([]models.Film, 0)

	for rows.Next() {
		var film models.Film
		err = rows.Scan(&film.Id, &film.Genres, &film.Country, &film.ReleaseRus, &film.Title,
			&film.Year, &film.Director, &film.Authors, &film.Actors, &film.Release, &film.Duration,
			&film.Language, &film.Budget, &film.Age, &film.Pic, &film.Src, &film.Description, &film.IsSeries, &film.NeedsPayment, &film.Slug)
		if err != nil {
			return nil, models.InternalError
		}
		if film.IsSeries {
			code := r.ReadSeries(&film)
			if code != models.Okey {
				return nil, models.InternalError
			}
		}
		f, code := r.GetRating(film)
		if code != models.Okey {
			f.Rating = 0
		}
		film.Rating = f.Rating
		films = append(films, film)
	}

	return films, models.Okey
}

func (r *FilmsRepo) GetNewestFilms() ([]models.Film, models.StatusCode) {

	rows, err := r.pool.Query(context.Background(), SELECT_FILM_BY_DATE)
	if err != nil {
		return nil, models.InternalError
	}
	defer rows.Close()

	films := make([]models.Film, 0)

	for rows.Next() {
		var film models.Film
		err = rows.Scan(&film.Id, &film.Genres, &film.Country, &film.ReleaseRus, &film.Title,
			&film.Year, &film.Director, &film.Authors, &film.Actors, &film.Release, &film.Duration,
			&film.Language, &film.Budget, &film.Age, &film.Pic, &film.Src, &film.Description, &film.IsSeries, &film.NeedsPayment, &film.Slug)
		if err != nil {
			return nil, models.InternalError
		}
		if film.IsSeries {
			code := r.ReadSeries(&film)
			if code != models.Okey {
				return nil, models.InternalError
			}
		}
		f, code := r.GetRating(film)
		if code != models.Okey {
			f.Rating = 0
		}
		film.Rating = f.Rating
		films = append(films, film)
	}

	return films, models.Okey
}

func (r *FilmsRepo) GetFilmsByKeyword(keyword string) ([]models.Film, models.StatusCode) {
	rows, err := r.pool.Query(context.Background(), SELECT_FILM_BY_KEYWORD, strings.Replace(keyword, " ", "&", -1), "%"+keyword+"%")

	if err != nil {
		return nil, models.InternalError
	}
	defer rows.Close()

	films := make([]models.Film, 0, 10)

	for rows.Next() {
		var film models.Film
		err = rows.Scan(&film.Id, &film.Genres, &film.Country, &film.ReleaseRus, &film.Title,
			&film.Year, &film.Director, &film.Authors, &film.Actors, &film.Release, &film.Duration,
			&film.Language, &film.Budget, &film.Age, &film.Pic, &film.Src, &film.Description, &film.IsSeries, &film.NeedsPayment, &film.Slug)
		if err != nil {
			return nil, models.InternalError
		}
		if film.IsSeries {
			code := r.ReadSeries(&film)
			if code != models.Okey {
				return nil, models.InternalError
			}
		}
		f, code := r.GetRating(film)
		if code == models.NotFound {
			f.Rating = 0
		}
		film.Rating = f.Rating
		films = append(films, film)
	}

	return films, models.Okey
}

func (r *FilmsRepo) GetFilmsByActor(actor models.Actors) ([]models.Film, models.StatusCode) {
	rows, err := r.pool.Query(context.Background(), SELECT_FILM_BY_ACTOR, actor.Id)
	if err != nil {
		return nil, models.InternalError
	}
	defer rows.Close()

	films := make([]models.Film, 0)

	for rows.Next() {
		var film models.Film
		err = rows.Scan(&film.Id, &film.Genres, &film.Country, &film.ReleaseRus, &film.Title,
			&film.Year, &film.Director, &film.Authors, &film.Actors, &film.Release, &film.Duration,
			&film.Language, &film.Budget, &film.Age, &film.Pic, &film.Src, &film.Description, &film.IsSeries, &film.NeedsPayment, &film.Slug)
		if err != nil {
			return nil, models.InternalError
		}
		if film.IsSeries {
			code := r.ReadSeries(&film)
			if code != models.Okey {
				return nil, models.InternalError
			}
		}
		f, code := r.GetRating(film)
		if code == models.NotFound {
			f.Rating = 0
		}
		film.Rating = f.Rating
		films = append(films, film)
	}
	return films, models.Okey
}

func (r *FilmsRepo) GetFilmById(film models.Film) (models.Film, models.StatusCode) {
	row := r.pool.QueryRow(context.Background(), SELECT_FILM_BY_ID, film.Id)

	err := row.Scan(&film.Id, &film.Genres, &film.Country, &film.ReleaseRus, &film.Title,
		&film.Year, &film.Director, &film.Authors, &film.Actors, &film.Release, &film.Duration,
		&film.Language, &film.Budget, &film.Age, &film.Pic, &film.Src, &film.Description, &film.IsSeries, &film.NeedsPayment, &film.Slug)

	if film.IsSeries {
		code := r.ReadSeries(&film)
		if code != models.Okey {
			return models.Film{}, models.InternalError
		}
	}

	f, code := r.GetRating(film)
	if code != models.Okey {
		f = models.Film{}
	}
	film.Rating = f.Rating
	if err != nil {
		return models.Film{}, models.InternalError
	}
	return film, models.Okey
}

func (r *FilmsRepo) GetFilmsByUser(user models.User) ([]models.Film, models.StatusCode) {
	rows, err := r.pool.Query(context.Background(), SELECT_FILM_BY_USER, user.Id)
	if err != nil {
		return nil, models.InternalError
	}
	defer rows.Close()
	films := make([]models.Film, 0)

	for rows.Next() {
		var film models.Film
		err = rows.Scan(&film.Id, &film.Genres, &film.Country, &film.ReleaseRus, &film.Title,
			&film.Year, &film.Director, &film.Authors, &film.Actors, &film.Release, &film.Duration,
			&film.Language, &film.Budget, &film.Age, &film.Pic, &film.Src, &film.Description, &film.IsSeries, &film.NeedsPayment, &film.Slug)
		if err != nil {
			return nil, models.InternalError
		}
		if film.IsSeries {
			code := r.ReadSeries(&film)
			if code != models.Okey {
				return nil, models.InternalError
			}
		}
		f, code := r.GetRating(film)
		if code == models.NotFound {
			f.Rating = 0
		}
		film.Rating = f.Rating
		films = append(films, film)
	}
	return films, models.Okey
}

func (r *FilmsRepo) InsertStarred(film models.Film, user models.User) models.StatusCode {

	exec, err := r.pool.Exec(context.Background(), INSERT_FILM_TO_STARRED, film.Id, user.Id)
	if err != nil {
		return models.InternalError
	}

	if exec.RowsAffected() != 1 {
		return models.Conflict
	}

	return models.Okey
}

func (r *FilmsRepo) DeleteStarred(film models.Film, user models.User) models.StatusCode {

	exec, err := r.pool.Exec(context.Background(), DELETE_FILM_FROM_STARRED, film.Id, user.Id)
	if err != nil {
		return models.InternalError
	}

	if exec.RowsAffected() != 1 {
		return models.Conflict
	}

	return models.Okey
}

func (r *FilmsRepo) InsertWatchlist(film models.Film, user models.User) models.StatusCode {

	exec, err := r.pool.Exec(context.Background(), INSERT_FILM_TO_WATCHLIST, user.Id, film.Id)
	if err != nil {
		return models.InternalError
	}
	//log.Println(err.Error())

	if exec.RowsAffected() != 1 {
		return models.Conflict
	}

	return models.Okey
}

func (r *FilmsRepo) DeleteWatchlist(film models.Film, user models.User) models.StatusCode {

	exec, err := r.pool.Exec(context.Background(), DELETE_FILM_FROM_WATCHLIST, film.Id, user.Id)
	if err != nil {
		return models.InternalError
	}

	if exec.RowsAffected() != 1 {
		return models.Conflict
	}

	return models.Okey
}

func (r FilmsRepo) GetStarredFilms(user models.User) ([]models.Film, models.StatusCode) {
	rows, err := r.pool.Query(context.Background(), GET_STARRED_FILMS, user.Id)
	if err != nil {
		return nil, models.InternalError
	}
	defer rows.Close()

	films := make([]models.Film, 0)

	for rows.Next() {
		var film models.Film
		err = rows.Scan(&film.Id, &film.Genres, &film.Country, &film.ReleaseRus, &film.Title,
			&film.Year, &film.Director, &film.Authors, &film.Actors, &film.Release, &film.Duration,
			&film.Language, &film.Budget, &film.Age, &film.Pic, &film.Src, &film.Description, &film.IsSeries, &film.NeedsPayment, &film.Slug)
		if err != nil {
			return nil, models.InternalError
		}
		films = append(films, film)
	}

	return films, models.Okey
}

func (r FilmsRepo) IfStarred(film models.Film, user models.User) models.StatusCode {
	rows := r.pool.QueryRow(context.Background(), GET_IF_STARRED_FILMS, user.Id, film.Id)

	err := rows.Scan(&film.Id, &film.Genres, &film.Country, &film.ReleaseRus, &film.Title,
		&film.Year, &film.Director, &film.Authors, &film.Actors, &film.Release, &film.Duration,
		&film.Language, &film.Budget, &film.Age, &film.Pic, &film.Src, &film.Description, &film.IsSeries, &film.NeedsPayment, &film.Slug)
	if err != nil {
		return models.NotFound
	}

	return models.Okey
}

func (r FilmsRepo) IfWatchList(film models.Film, user models.User) models.StatusCode {
	rows := r.pool.QueryRow(context.Background(), GET_IF_WL_FILMS, user.Id, film.Id)

	err := rows.Scan(&film.Id, &film.Genres, &film.Country, &film.ReleaseRus, &film.Title,
		&film.Year, &film.Director, &film.Authors, &film.Actors, &film.Release, &film.Duration,
		&film.Language, &film.Budget, &film.Age, &film.Pic, &film.Src, &film.Description, &film.IsSeries, &film.NeedsPayment, &film.Slug)
	if err != nil {
		return models.NotFound
	}

	return models.Okey
}

func (r FilmsRepo) GetWatchlistFilms(user models.User) ([]models.Film, models.StatusCode) {
	rows, err := r.pool.Query(context.Background(), GET_WATCHLIST_FILMS, user.Id)
	if err != nil {
		return nil, models.InternalError
	}
	defer rows.Close()

	films := make([]models.Film, 0)

	for rows.Next() {
		var film models.Film
		err = rows.Scan(&film.Id, &film.Genres, &film.Country, &film.ReleaseRus, &film.Title,
			&film.Year, &film.Director, &film.Authors, &film.Actors, &film.Release, &film.Duration,
			&film.Language, &film.Budget, &film.Age, &film.Pic, &film.Src, &film.Description, &film.IsSeries, &film.NeedsPayment, &film.Slug)
		if err != nil {
			return nil, models.InternalError
		}
		films = append(films, film)
	}

	return films, models.Okey
}

func (r FilmsRepo) ReadSeries(film *models.Film) models.StatusCode {
	rows, err := r.pool.Query(context.Background(), GET_SERIES,
		film.Id)

	if err != nil {
		return models.InternalError
	}

	defer rows.Close()
	seasons := make([]models.Season, 0)

	for rows.Next() {
		var season models.Season
		err = rows.Scan(&season.Num, &season.Src, &season.Pics)
		if err != nil {
			return models.InternalError
		}
		seasons = append(seasons, season)

	}

	film.Seasons = &seasons
	return models.Okey
}

func (r *FilmsRepo) GetRandom() (models.Film, models.StatusCode) {

	var count int
	row := r.pool.QueryRow(context.Background(), SElECT_RANDOM_COUNT)
	if err := row.Scan(&count); err != nil {
		return models.Film{}, models.InternalError
	}

	id := rand.Intn(count)

	row = r.pool.QueryRow(context.Background(), SElECT_RANDOM_FILM, id)
	var film models.Film
	err := row.Scan(&film.Id, &film.Genres, &film.Country, &film.ReleaseRus, &film.Title,
		&film.Year, &film.Director, &film.Authors, &film.Actors, &film.Release, &film.Duration,
		&film.Language, &film.Budget, &film.Age, &film.Pic, &film.Src, &film.Description, &film.IsSeries, &film.NeedsPayment, &film.Slug)
	if err != nil {
		return models.Film{}, models.InternalError
	}
	if film.IsSeries {
		code := r.ReadSeries(&film)
		if code != models.Okey {
			return models.Film{}, models.InternalError
		}
	}
	f, code := r.GetRating(film)
	if code == models.NotFound {
		f.Rating = 0
	}

	return film, models.Okey
}

func (r *FilmsRepo) SetRating(film models.Film, user models.User, rating float64) models.StatusCode {
	exec, err := r.pool.Exec(context.Background(), SET_RATING, user.Id, film.Id, rating)
	if err != nil {
		return models.InternalError
	}
	if exec.RowsAffected() != 1 {
		return models.Conflict
	}
	return models.Okey
}

func (r *FilmsRepo) GetRating(film models.Film) (models.Film, models.StatusCode) {
	row := r.pool.QueryRow(context.Background(), GET_RATING, film.Id)
	err := row.Scan(&film.Rating)
	if err != nil {
		return film, models.NotFound
	}
	return film, models.Okey
}

func (r *FilmsRepo) GetRatingByUser(film models.Film, user models.User) (models.Film, models.StatusCode) {
	row := r.pool.QueryRow(context.Background(), GET_RATING_BY_USER, user.Id, film.Id)
	err := row.Scan(&film.Rating)
	if err != nil {
		return film, models.NotFound
	}
	return film, models.Okey
}

func (r *FilmsRepo) GetIdBySlug(slug string) (models.Film, models.StatusCode) {

	film := models.Film{}
	row := r.pool.QueryRow(context.Background(), GET_ID_BY_SLUG, slug)
	err := row.Scan(&film.Id)
	if err != nil {
		return models.Film{}, models.NotFound
	}
	return film, models.Okey
}
