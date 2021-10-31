package repo

import (
	"github.com/driftprogramming/pgxpoolmock"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/models"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func prepare(t *testing.T) (*FilmsRepo, uuid.UUID) { // mock on repo and uid used as uuid for everything that has any uid filed
	ctl := gomock.NewController(t)
	mockPool := pgxpoolmock.NewMockPgxPool(ctl)
	uidStr := "40266371-008c-4911-813d-65d222eb4d47"
	uid, _ := uuid.Parse(uidStr)
	filmRepo := NewFilmsRepo(mockPool)
	columns := []string{"id", "genres", "title", "year", "director", "authors", "actors", "release", "duration", "language", "src"}
	pgxRows := pgxpoolmock.NewRows(columns).
		AddRow(uid, []string{"comedy"}, "Policeman from Rublevka",
			2017, []string{"Alex"}, []string{"Alex"},
			[]uuid.UUID{uid}, time.Now(), 120, "russian", []string{"img.png"}).ToPgxRows()
	mockPool.EXPECT().Query(gomock.Any(), gomock.Any(), gomock.Any()).Return(pgxRows, nil)
	return filmRepo, uid
}

func TestFilmsRepo_GetFilmById(t *testing.T) {
	ctl := gomock.NewController(t)

	mockPool := pgxpoolmock.NewMockPgxPool(ctl)
	uidStr := "40266371-008c-4911-813d-65d222eb4d47"
	uid, _ := uuid.Parse(uidStr)
	filmRepo := NewFilmsRepo(mockPool)
	columns := []string{"id", "genres", "title", "year", "director", "authors", "actors", "release", "duration", "language", "src"}
	pgxRows := pgxpoolmock.NewRows(columns).
		AddRow(uid, []string{"comedy"}, "Policeman from Rublevka",
			2017, []string{"Alex"}, []string{"Alex"},
			[]uuid.UUID{uid}, time.Now(), 120, "russian", []string{"img.png"}).ToPgxRows()
	pgxRows.Next()
	mockPool.EXPECT().QueryRow(gomock.Any(), gomock.Any(), gomock.Any()).Return(pgxRows)
	filmRequest := models.Film{
		Id: uid,
	}
	film, status := filmRepo.GetFilmById(filmRequest)
	if status != models.Okey {
		t.Error("Wrong result")
	}
	assert.NotNil(t, film.Id)
}

func TestFilmsRepo_GetFilmsByActor(t *testing.T) {
	filmRepo, uid := prepare(t)
	filmRequest := models.Actors{
		Id: uid,
	}
	film, status := filmRepo.GetFilmsByActor(filmRequest)
	if status != models.Okey {
		t.Error("Wrong result")
	}
	assert.NotNil(t, film[0].Id)
	assert.Equal(t, uid, film[0].Id)
	assert.Equal(t, 2017, film[0].Year)
}

func TestFilmsRepo_GetFilmsByKeyword(t *testing.T) {
	keyword := "SATANA"
	filmRepo, uid := prepare(t)

	film, status := filmRepo.GetFilmsByKeyword(keyword)
	if status != models.Okey {
		t.Error("Wrong result")
	}
	assert.NotNil(t, film[0].Id)
	assert.Equal(t, uid, film[0].Id)
	assert.Equal(t, 2017, film[0].Year)
}

func TestFilmsRepo_GetFilmsByTopic(t *testing.T) {
	topic := "SATANA"
	filmRepo, uid := prepare(t)

	film, status := filmRepo.GetFilmsByTopic(topic)
	if status != models.Okey {
		t.Error("Wrong result")
	}
	assert.NotNil(t, film[0].Id)
	assert.Equal(t, uid, film[0].Id)
	assert.Equal(t, 2017, film[0].Year)
}

func TestFilmsRepo_GetFilmsByUser(t *testing.T) {
	filmRepo, uid := prepare(t)
	user := models.User{Id: uid}
	film, status := filmRepo.GetFilmsByUser(user)
	if status != models.Okey {
		t.Error("Wrong result")
	}
	assert.NotNil(t, film[0].Id)
	assert.Equal(t, uid, film[0].Id)
	assert.Equal(t, 2017, film[0].Year)
}

func TestFilmsRepo_GetHottestFilms(t *testing.T) {
	filmRepo, uid := prepare(t)
	film, status := filmRepo.GetHottestFilms()
	if status != models.Okey {
		t.Error("Wrong result")
	}
	assert.NotNil(t, film[0].Id)
	assert.Equal(t, uid, film[0].Id)
	assert.Equal(t, 2017, film[0].Year)
}

func TestFilmsRepo_GetNewestFilms(t *testing.T) {
	filmRepo, uid := prepare(t)
	film, status := filmRepo.GetNewestFilms()
	if status != models.Okey {
		t.Error("Wrong result")
	}
	assert.NotNil(t, film[0].Id)
	assert.Equal(t, uid, film[0].Id)
	assert.Equal(t, 2017, film[0].Year)
}
