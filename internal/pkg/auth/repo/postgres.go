package repo

import (
	"context"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"time"
)

const (
	SElECT_USER = "SELECT id, email, login, encrypted_password, created_at FROM public.users;"
	CHECK_USER  = "SELECT id, encrypted_password FROM public.users WHERE login=$1;"
	CREATE_USER = "INSERT INTO public.users(id, email, login, encrypted_password, created_at) VALUES($1, $2, $3, $4, $5) RETURNING id;"
)

type AuthRepo struct {
	pool *pgxpool.Pool
}

func NewAuthRepo(pool *pgxpool.Pool) *AuthRepo {
	return &AuthRepo{pool: pool}
}

func (r *AuthRepo) CreateUser(user models.User) (models.User, models.StatusCode) {
	var id uuid.UUID
	user.Id = uuid.New()
	row := r.pool.QueryRow(context.Background(), CREATE_USER,
		user.Id, user.Email, user.Login, user.EncryptedPassword, time.Now())

	err := row.Scan(&id)
	if err != nil && id == user.Id {
		return models.User{}, models.InternalError
	}
	userOut := models.User{
		Id:                id,
		Login:             user.Login,
		EncryptedPassword: user.EncryptedPassword,
		Email:             user.Email,
	}
	return userOut, models.Okey
}

func (r *AuthRepo) CheckUser(user models.User) (models.User, models.StatusCode) {
	var (
		pwd string
		id  uuid.UUID
	)
	row := r.pool.QueryRow(context.Background(), CHECK_USER, user.Login)

	if err := row.Scan(&id, &pwd); err != nil {
		return models.User{}, models.InternalError
	}
	if pwd != user.EncryptedPassword {
		return models.User{}, models.Unauthed
	}

	userOut := models.User{
		Id:                id,
		Login:             user.Login,
		EncryptedPassword: user.EncryptedPassword,
		Email:             user.Email,
	}
	return userOut, models.Okey
}
