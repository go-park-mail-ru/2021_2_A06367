package repo

import "github.com/jackc/pgx/v4/pgxpool"

type FilmsRepo struct {
	pool pgxpool.Pool
}

func NewFilmsRepo(pool pgxpool.Pool) *FilmsRepo {
	return &FilmsRepo{pool: pool}
}

