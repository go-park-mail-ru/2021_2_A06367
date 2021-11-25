package repo

import (
	"github.com/driftprogramming/pgxpoolmock"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/models"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"testing"
	"time"
)

func TestActorsRepo_GetActorById(t *testing.T) {
	uid := uuid.New()
	columns := []string{"id", "name", "surname", "avatar", "height", "date_of_birth", "description", "genres"}
	pgxRows := pgxpoolmock.NewRows(columns).
		AddRow(uid, "tester3",
			"tester3", "img.png", float32(178),
			time.Now(), "test test", []string{"Comedy"}).ToPgxRows()
	pgxRows.Next()
	ctl := gomock.NewController(t)
	mockPool := pgxpoolmock.NewMockPgxPool(ctl)

	repo := NewActorsRepo(mockPool, nil)
	mockPool.EXPECT().QueryRow(gomock.Any(), gomock.Any(), gomock.Any()).Return(pgxRows)

	actor := models.Actors{Id: uid}
	actorCheck, st := repo.GetActorById(actor)

	if st != models.Okey {
		t.Error("wrong status returned")
	}
	if actorCheck.Id != actor.Id {
		t.Error("wrong actor returned")
	}
}

func TestActorsRepo_GetActors(t *testing.T) {
	uid := uuid.New()

	columns := []string{"id", "name", "surname", "avatar", "height", "date_of_birth", "description", "genres"}
	pgxRows := pgxpoolmock.NewRows(columns).
		AddRow(uid, "tester2",
			"tester2", "img.png", float32(178),
			time.Now(), "test test", []string{"Comedy"}).ToPgxRows()

	ctl := gomock.NewController(t)
	mockPool := pgxpoolmock.NewMockPgxPool(ctl)

	repo := NewActorsRepo(mockPool, nil)
	mockPool.EXPECT().Query(gomock.Any(), gomock.Any(), gomock.Any()).Return(pgxRows, nil)

	actor := models.Actors{Id: uid}
	_, st := repo.GetActors([]models.Actors{actor})

	if st != models.Okey {
		t.Error("wrong status returned")
	}
}
