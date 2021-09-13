package repo

import (
	"context"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"time"
)

const (
	SElECT_USER = "SELECT id, email, login, encrypted_password, created_at FROM users;"
	CHECK_USER = "SELECT encrypted_password FROM users WHERE login=$1;"
	CREATE_USER = "INSERT INTO users(id, email, login, encrypted_password, created_at) VALUES($1, $2, $3, $4, $5) RETURNING id;"
)

type AuthRepo struct {
	pool pgxpool.Pool
}

func NewAuthRepo(pool pgxpool.Pool) *AuthRepo {
	return &AuthRepo{pool: pool}
}

func (r *AuthRepo) CreateUser(user models.User)  models.StatusCode {

	var id uuid.UUID
	row := r.pool.QueryRow(context.Background(), CREATE_USER,
		user.Id, user.Email, user.Login, user.EncryptedPassword, time.Now())

	err := row.Scan(&id)
	if err != nil {
		return models.InternalError
	}
	return models.Okey
}

func (r *AuthRepo)  CheckUser(user models.User) models.StatusCode {
	var pwd string
	row := r.pool.QueryRow(context.Background(), CHECK_USER, user.Login)

	if err := row.Scan(&pwd); err != nil {
		return models.InternalError
	}

	if pwd != user.EncryptedPassword {
		return models.Unauthed
	}

	return models.Okey
}
