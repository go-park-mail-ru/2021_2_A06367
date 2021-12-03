package repo

import (
	"context"
	"fmt"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/models"
	"github.com/google/uuid"
	"github.com/jackc/pgtype/pgxtype"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

const (
	SElECT_ACTOR_BY_ID       = "SELECT id, name, surname, avatar, height, date_of_birth, description, genres FROM actors WHERE id = $1"
	SELECT_ACTORS_BY_ID      = "SELECT id, name, surname, avatar, height,date_of_birth,genres FROM actors WHERE id IN ($1)"
	SELECT_ACTORS_BY_KEYWORD = "SELECT * FROM actors WHERE make_tsvector(name) @@ to_tsquery($1) or LOWER(name) like LOWER($2)   LIMIT 10"
)

type ActorsRepo struct {
	pool   pgxtype.Querier
	logger *zap.SugaredLogger
}

func NewActorsRepo(pool pgxtype.Querier, logger *zap.SugaredLogger) *ActorsRepo {
	return &ActorsRepo{
		pool:   pool,
		logger: logger,
	}
}

func (r *ActorsRepo) GetActorById(actor models.Actors) (models.Actors, models.StatusCode) {
	row := r.pool.QueryRow(context.Background(), SElECT_ACTOR_BY_ID, actor.Id)
	err := row.Scan(&actor.Id, &actor.Name, &actor.Surname, &actor.Avatar, &actor.Height, &actor.DateOfBirth, &actor.Description, &actor.Genres)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return models.Actors{}, models.NotFound
		}
		return models.Actors{}, models.InternalError
	}
	if r.logger != nil {
		r.logger.Info(zap.String("Status:", string(rune(http.StatusOK))))
	}

	return actor, models.Okey
}

func (r *ActorsRepo) GetActors(actors []models.Actors) ([]models.Actors, models.StatusCode) {
	uids := []uuid.UUID{}
	for _, actor := range actors {
		uids = append(uids, actor.Id)
	}

	args := []string{}
	for _, el := range uids {
		args = append(args, fmt.Sprintf("'%v'", el))
	}
	arg := strings.Join(args, ",")

	rows, err := r.pool.Query(context.Background(), fmt.Sprintf("SELECT id, name, surname, avatar, height, date_of_birth, description, genres FROM actors WHERE id IN (%s)", arg))

	if err != nil {
		return nil, models.InternalError
	}

	i := 0
	for rows.Next() {
		actor := models.Actors{}
		err = rows.Scan(&actor.Id, &actor.Name, &actor.Surname, &actor.Avatar,
			&actor.Height, &actor.DateOfBirth, &actor.Description, &actor.Genres)
		if err != nil {
			return nil, models.InternalError
		}
		actors[i] = actor
		i++
	}

	if i == 0 {
		return nil, models.NotFound
	}

	if r.logger != nil {
		r.logger.Info(zap.String("Status:", string(rune(http.StatusOK))))
	}

	return actors, models.Okey
}

func (r *ActorsRepo) GetActorsByKeyword(keyword string) ([]models.Actors, models.StatusCode) {
	rows, err := r.pool.Query(context.Background(), SELECT_ACTORS_BY_KEYWORD, strings.Replace(keyword, " ", "&", -1), "%"+keyword+"%")
	if err != nil {
		return []models.Actors{}, models.InternalError
	}
	actors := make([]models.Actors, 0, 10)

	for rows.Next() {
		var actor models.Actors
		err = rows.Scan(&actor.Id, &actor.Name, &actor.Surname, &actor.Avatar, &actor.Height, &actor.DateOfBirth, &actor.Description, &actor.Genres)
		if err != nil {
			return nil, models.InternalError
		}
		actors = append(actors, actor)
	}
	return actors, models.Okey
}
