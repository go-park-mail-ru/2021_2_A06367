package repo

import (
	"context"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"time"
)

const (
	SElECT_USER = "SELECT login, about, avatar, subscriptions, subscribers FROM public.users WHERE id=$1;"
	CHECK_USER  = "SELECT encrypted_password FROM public.users WHERE login=$1;"
	CREATE_USER = "INSERT INTO public.users(id, email, login, encrypted_password, created_at) VALUES($1, $2, $3, $4, $5) RETURNING id;"
	FOLLOW      = "INSERT INTO public.subscriptions (user_id, subscribed_at) VALUES($1, $2) RETURNING id;"
	UNFOLLOW    = "DELETE FROM public.subscriptions WHERE user_id=$1 AND subscribed_at=$2;"
	SELECT_FOLLOWING = "SELECT users.id, email, login, encrypted_password, about, avatar, subscriptions, " +
		"subscribers, created_at FROM users JOIN subscriptions ON users.id = subscriptions.user_id;"
	SELECT_FOLLOWERS = "SELECT users.id, email, login, encrypted_password, about, avatar, subscriptions, " +
		"subscribers, created_at FROM users JOIN subscriptions ON users.id = subscriptions.subscribed_at;"
)

type AuthRepo struct {
	pool *pgxpool.Pool
}

func NewAuthRepo(pool *pgxpool.Pool) *AuthRepo {
	return &AuthRepo{pool: pool}
}

func (r *AuthRepo) CreateUser(user models.User) models.StatusCode {

	var id uuid.UUID
	user.Id = uuid.New()
	row := r.pool.QueryRow(context.Background(), CREATE_USER,
		user.Id, user.Email, user.Login, user.EncryptedPassword, time.Now())

	err := row.Scan(&id)
	if err != nil {
		return models.InternalError
	}
	return models.Okey
}

func (r *AuthRepo) CheckUser(user models.User) models.StatusCode {
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

func (r *AuthRepo) GetProfile(user models.Profile) (models.Profile, models.StatusCode) {

	row := r.pool.QueryRow(context.Background(), SElECT_USER,
		user.Id)

	err := row.Scan(&user.Login, &user.About, &user.Avatar, &user.Subscriptions, &user.Subscribers)
	if err != nil {
		return models.Profile{}, models.InternalError
	}
	return user, models.Okey
}

func (r *AuthRepo) AddFollowing(who, whom uuid.UUID) models.StatusCode {

	var id int
	row := r.pool.QueryRow(context.Background(), FOLLOW,
		who, whom)

	err := row.Scan(&id)
	if err != nil {
		//TODO: добавить проверку ошибок
		return models.InternalError
	}
	return models.Okey
}

func (r *AuthRepo) RemoveFollowing(who, whom uuid.UUID) models.StatusCode {

	exec, err := r.pool.Exec(context.Background(), UNFOLLOW,
		who, whom)
	if err != nil {
		return models.InternalError
	}

	if exec.RowsAffected() == 0 {
		return models.NotFound
	}
	return models.Okey
}
