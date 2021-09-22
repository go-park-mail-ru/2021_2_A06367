package repo

import (
	"context"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	OnQuery    = `INSERT INTO online_users VALUES($1);`
	OffQuery   = `DELETE FROM online_users WHERE login=$1;`
	CheckQuery = `SELECT * FROM online_users WHERE login=$1;`
)

type OnlineRepo struct {
	conn *pgxpool.Pool
}

func NewOnlineRepo(pool *pgxpool.Pool) *OnlineRepo {
	return &OnlineRepo{
		conn: pool,
	}
}

func (or *OnlineRepo) UserOn(user models.User) models.StatusCode {
	if or.IsOnline(user) {
		return models.Okey
	}

	_, err := or.conn.Exec(context.Background(), OnQuery, user.Login)
	if err != nil {
		return models.InternalError
	}
	return models.Okey
}

func (or *OnlineRepo) UserOff(user models.User) models.StatusCode {
	_, err := or.conn.Exec(context.Background(), OffQuery, user.Login)
	if err != nil {
		return models.InternalError
	}
	return models.Okey
}

func (or *OnlineRepo) IsOnline(user models.User) bool {
	var login string
	err := or.conn.QueryRow(context.Background(), CheckQuery, user.Login).Scan(&login)
	if err != nil || login == "" {
		return false
	}
	return true
}
