package repo

import (
	"context"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	SElECT_FILM_BY_TOPIC = "SELECT id, genres, title, year, director, " +
		"authors, release, duration, language " +
		"FROM films "
)

type FilmsRepo struct {
	pool pgxpool.Pool
}

func NewFilmsRepo(pool pgxpool.Pool) *FilmsRepo {
	return &FilmsRepo{pool: pool}
}

func (r *FilmsRepo) GetFilmsByTopic(topic string) ([]models.Film, models.StatusCode)  {

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