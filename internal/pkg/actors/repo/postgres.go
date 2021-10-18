package repo

import (
	"context"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	SElECT_ACTOR_BY_ID = "SELECT * FROM actors WHERE id = $1"
)

type ActorsRepo struct {
	pool *pgxpool.Pool
}

func NewActorsRepo(pool *pgxpool.Pool) *ActorsRepo {
	return &ActorsRepo{pool: pool}
}

func (r *ActorsRepo) GetActorById(actor models.Actors) (models.Actors, models.StatusCode) {
	row := r.pool.QueryRow(context.Background(), SElECT_ACTOR_BY_ID, actor.Id)
	err := row.Scan(&actor.Id, &actor.Name, &actor.Surname, &actor.Avatar, &actor.Height, &actor.DateOfBirth, &actor.Genres)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return models.Actors{}, models.NotFound
		}
		return models.Actors{}, models.InternalError
	}
	return actor, models.Okey
}