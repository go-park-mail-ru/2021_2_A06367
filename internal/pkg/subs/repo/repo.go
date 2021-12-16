package repo

import (
	"context"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/models"
	"github.com/google/uuid"
	"github.com/jackc/pgtype/pgxtype"
)

const (
	SelectLicense = "SELECT exp_date FROM subs WHERE user_id = $1;"
	UpsertLicense = "INSERT INTO subs(user_id, exp_date) VALUES ($1, $2) " +
		"on conflict(user_id) do update set exp_date = $2 " +
		"RETURNING user_id;"
)

type SubsRepo struct {
	pool pgxtype.Querier
}

func NewSubsRepo(pool pgxtype.Querier) *SubsRepo {
	return &SubsRepo{pool: pool}
}

func (r *SubsRepo) GetLicense(id uuid.UUID) (models.License, models.StatusCode) {
	var l models.License
	rows := r.pool.QueryRow(context.Background(), SelectLicense, id)
	err := rows.Scan(&l.ExpDate)
	if err != nil {
		return models.License{}, models.NotFound
	}
	l.IsValid = true
	return l, models.Okey
}

func (r *SubsRepo) SetLicense(id uuid.UUID, license models.License) (models.License, models.StatusCode) {
	row := r.pool.QueryRow(context.Background(), UpsertLicense, id, license.ExpDate)
	err := row.Scan(&id)
	if err != nil {
		return models.License{}, models.BadRequest
	}
	return license, models.Okey
}
