package repo

import (
	"context"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	SElECT_FILM_BY_TOPIC = "SELECT id, genres, title, year, director, " +
		"authors, release, duration, language " +
		"FROM films " +
		"WHERE $1 = ANY(genres)"

	SELECT_FILM_BY_RATING = "SELECT id, genres, title, year, director, authors, release, duration, language " +
		" FROM films JOIN rating ON films.id = rating.film_id ORDER BY rating DESC LIMIT 10"
	SELECT_FILM_BY_DATE = "SELECT id, genres, title, year, director, authors, release, duration, language " +
		"FROM films  ORDER BY release DESC LIMIT 10"
)

type FilmsRepo struct {
	pool *pgxpool.Pool
}

func NewFilmsRepo(pool *pgxpool.Pool) *FilmsRepo {
	return &FilmsRepo{pool: pool}
}

func (r *FilmsRepo) GetFilmsByTopic(topic string) ([]models.Film, models.StatusCode) {

	rows, err := r.pool.Query(context.Background(), SElECT_FILM_BY_TOPIC,
		topic)
	if err != nil {
		return nil, models.InternalError
	}

	films := make([]models.Film, 0)

	for rows.Next() {
		var film models.Film
		err = rows.Scan(&film.Id, &film.Genres, &film.Title,
			&film.Year, &film.Director, &film.Authors, &film.Release, &film.Duration,
			&film.Language)
		if err != nil {
			return nil, models.InternalError
		}
		films = append(films, film)
	}

	return films, models.Okey
}

func (r *FilmsRepo) GetHottestFilms() ([]models.Film, models.StatusCode) {

	rows, err := r.pool.Query(context.Background(), SELECT_FILM_BY_RATING)
	if err != nil {
		return nil, models.InternalError
	}

	films := make([]models.Film, 0)

	for rows.Next() {
		var film models.Film
		err = rows.Scan(&film.Id, &film.Genres, &film.Title,
			&film.Year, &film.Director, &film.Authors, &film.Release, &film.Duration,
			&film.Language)
		if err != nil {
			return nil, models.InternalError
		}
		films = append(films, film)
	}

	return films, models.Okey
}

func (r *FilmsRepo) GetNewestFilms() ([]models.Film, models.StatusCode) {

	rows, err := r.pool.Query(context.Background(), SELECT_FILM_BY_DATE)
	if err != nil {
		return nil, models.InternalError
	}

	films := make([]models.Film, 0)

	for rows.Next() {
		var film models.Film
		err = rows.Scan(&film.Id, &film.Genres, &film.Title,
			&film.Year, &film.Director, &film.Authors, &film.Release, &film.Duration,
			&film.Language)
		if err != nil {
			return nil, models.InternalError
		}
		films = append(films, film)
	}

	return films, models.Okey
}
