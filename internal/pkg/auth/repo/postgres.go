package repo

import "github.com/jackc/pgx/v4/pgxpool"

type AuthRepo struct {
	pool pgxpool.Pool
}

func NewAuthRepo(pool pgxpool.Pool) *AuthRepo {
	return &AuthRepo{pool: pool}
}

