package repo

import (
	"errors"
	"github.com/driftprogramming/pgxpoolmock"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/models"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"log"
	"strings"
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

	lg := CreateLogger()
	repo := NewActorsRepo(mockPool, lg)
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

func TestActorsRepo_GetActorByIdNotFound(t *testing.T) {
	uid := uuid.New()
	columns := []string{"id", "name", "surname", "avatar", "height", "date_of_birth", "description", "genres"}
	pgxRows := pgxpoolmock.NewRows(columns).
		AddRow(uid, "tester3",
			"tester3", "img.png", float32(178),
			time.Now(), "test test", []string{"Comedy"}).RowError(0, errors.New("no rows in result set")).ToPgxRows()
	pgxRows.Next()
	ctl := gomock.NewController(t)
	mockPool := pgxpoolmock.NewMockPgxPool(ctl)

	lg := CreateLogger()
	repo := NewActorsRepo(mockPool, lg)
	mockPool.EXPECT().QueryRow(gomock.Any(), gomock.Any(), gomock.Any()).Return(pgxRows)

	actor := models.Actors{Id: uid}
	_, st := repo.GetActorById(actor)

	if st != models.NotFound {
		t.Error("wrong status returned")
	}
}

func TestActorsRepo_GetActorByIdInternal(t *testing.T) {
	uid := uuid.New()
	columns := []string{"id", "name", "surname", "avatar", "height", "date_of_birth", "description", "genres"}
	pgxRows := pgxpoolmock.NewRows(columns).
		AddRow(uid, "tester3",
			"tester3", "img.png", float32(178),
			time.Now(), "test test", []string{"Comedy"}).RowError(0, errors.New("surprise error")).ToPgxRows()
	pgxRows.Next()
	ctl := gomock.NewController(t)
	mockPool := pgxpoolmock.NewMockPgxPool(ctl)

	lg := CreateLogger()
	repo := NewActorsRepo(mockPool, lg)
	mockPool.EXPECT().QueryRow(gomock.Any(), gomock.Any(), gomock.Any()).Return(pgxRows)

	actor := models.Actors{Id: uid}
	_, st := repo.GetActorById(actor)

	if st != models.InternalError {
		t.Error("wrong status returned")
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

	lg := CreateLogger()
	repo := NewActorsRepo(mockPool, lg)
	mockPool.EXPECT().Query(gomock.Any(), gomock.Any(), gomock.Any()).Return(pgxRows, nil)

	actor := models.Actors{Id: uid}
	_, st := repo.GetActors([]models.Actors{actor})

	if st != models.Okey {
		t.Error("wrong status returned")
	}
}

func TestActorsRepo_GetActorsError(t *testing.T) {
	uid := uuid.New()

	columns := []string{"id", "name", "surname", "avatar", "height", "date_of_birth", "description", "genres"}
	pgxRows := pgxpoolmock.NewRows(columns).
		AddRow(uid, "tester2",
			"tester2", "img.png", float32(178),
			time.Now(), "test test", []string{"Comedy"}).RowError(0, errors.New("")).ToPgxRows()

	ctl := gomock.NewController(t)
	mockPool := pgxpoolmock.NewMockPgxPool(ctl)

	lg := CreateLogger()
	repo := NewActorsRepo(mockPool, lg)
	mockPool.EXPECT().Query(gomock.Any(), gomock.Any(), gomock.Any()).Return(pgxRows, errors.New("wow"))

	actor := models.Actors{Id: uid}
	_, st := repo.GetActors([]models.Actors{actor})

	if st != models.InternalError {
		t.Error("wrong status returned")
	}
}

func TestActorsRepo_GetActorsScanError(t *testing.T) {
	uid := uuid.New()

	columns := []string{"id", "name", "surname", "avatar", "height", "date_of_birth", "description", "genres"}
	pgxRows := pgxpoolmock.NewRows(columns).
		AddRow(uid, "tester2",
			"tester2", "img.png", float32(178),
			time.Now(), "test test", []string{"Comedy"}).RowError(0, errors.New("")).ToPgxRows()

	ctl := gomock.NewController(t)
	mockPool := pgxpoolmock.NewMockPgxPool(ctl)

	lg := CreateLogger()
	repo := NewActorsRepo(mockPool, lg)
	mockPool.EXPECT().Query(gomock.Any(), gomock.Any(), gomock.Any()).Return(pgxRows, nil)

	actor := models.Actors{Id: uid}
	_, st := repo.GetActors([]models.Actors{actor})

	if st != models.InternalError {
		t.Error("wrong status returned")
	}
}

func TestActorsRepo_GetActorsNotFoundError(t *testing.T) {
	uid := uuid.New()

	columns := []string{"id", "name", "surname", "avatar", "height", "date_of_birth", "description", "genres"}
	pgxRows := pgxpoolmock.NewRows(columns).ToPgxRows()

	ctl := gomock.NewController(t)
	mockPool := pgxpoolmock.NewMockPgxPool(ctl)

	lg := CreateLogger()
	repo := NewActorsRepo(mockPool, lg)
	mockPool.EXPECT().Query(gomock.Any(), gomock.Any(), gomock.Any()).Return(pgxRows, nil)

	actor := models.Actors{Id: uid}
	_, st := repo.GetActors([]models.Actors{actor})

	if st != models.NotFound {
		t.Error("wrong status returned")
	}
}

func TestActorsRepo_GetKeyWord(t *testing.T) {
	uid := uuid.New()

	columns := []string{"id", "name", "surname", "avatar", "height", "date_of_birth", "description", "genres"}
	pgxRows := pgxpoolmock.NewRows(columns).
		AddRow(uid, "tester2",
			"tester2", "img.png", float32(178),
			time.Now(), "test test", []string{"Comedy"}).ToPgxRows()

	ctl := gomock.NewController(t)
	mockPool := pgxpoolmock.NewMockPgxPool(ctl)

	lg := CreateLogger()
	repo := NewActorsRepo(mockPool, lg)
	mockPool.EXPECT().Query(gomock.Any(), SELECT_ACTORS_BY_KEYWORD, strings.Replace("abc", " ", "&", -1), "%"+"abc"+"%").Return(pgxRows, nil)

	_, st := repo.GetActorsByKeyword("abc")

	if st != models.Okey {
		t.Error("wrong status returned")
	}
}

func TestActorsRepo_GetKeyWordError(t *testing.T) {
	uid := uuid.New()

	columns := []string{"id", "name", "surname", "avatar", "height", "date_of_birth", "description", "genres"}
	pgxRows := pgxpoolmock.NewRows(columns).
		AddRow(uid, "tester2",
			"tester2", "img.png", float32(178),
			time.Now(), "test test", []string{"Comedy"}).RowError(0, errors.New("")).ToPgxRows()

	ctl := gomock.NewController(t)
	mockPool := pgxpoolmock.NewMockPgxPool(ctl)

	lg := CreateLogger()
	repo := NewActorsRepo(mockPool, lg)
	mockPool.EXPECT().Query(gomock.Any(), SELECT_ACTORS_BY_KEYWORD, strings.Replace("abc", " ", "&", -1), "%"+"abc"+"%").Return(pgxRows, errors.New(""))

	_, st := repo.GetActorsByKeyword("abc")
	if st != models.InternalError {
		t.Error("wrong status returned")
	}
}

func TestActorsRepo_GetKeyWordScanError(t *testing.T) {
	uid := uuid.New()

	columns := []string{"id", "name", "surname", "avatar", "height", "date_of_birth", "description", "genres"}
	pgxRows := pgxpoolmock.NewRows(columns).
		AddRow(uid, "tester2",
			"tester2", "img.png", float32(178),
			time.Now(), "test test", []string{"Comedy"}).RowError(0, errors.New("")).ToPgxRows()

	ctl := gomock.NewController(t)
	mockPool := pgxpoolmock.NewMockPgxPool(ctl)

	lg := CreateLogger()
	repo := NewActorsRepo(mockPool, lg)
	mockPool.EXPECT().Query(gomock.Any(), SELECT_ACTORS_BY_KEYWORD, strings.Replace("abc", " ", "&", -1), "%"+"abc"+"%").Return(pgxRows, nil)

	_, st := repo.GetActorsByKeyword("abc")

	if st != models.InternalError {
		t.Error("wrong status returned")
	}
}

func TestActorsRepo_GetKeyWordNotFoundError(t *testing.T) {

	columns := []string{"id", "name", "surname", "avatar", "height", "date_of_birth", "description", "genres"}
	pgxRows := pgxpoolmock.NewRows(columns).ToPgxRows()

	ctl := gomock.NewController(t)
	mockPool := pgxpoolmock.NewMockPgxPool(ctl)

	lg := CreateLogger()
	repo := NewActorsRepo(mockPool, lg)
	mockPool.EXPECT().Query(gomock.Any(), SELECT_ACTORS_BY_KEYWORD, strings.Replace("abc", " ", "&", -1), "%"+"abc"+"%").Return(pgxRows, nil)

	_, st := repo.GetActorsByKeyword("abc")

	if st != models.NotFound {
		t.Error("wrong status returned")
	}
}

func CreateLogger() *zap.SugaredLogger {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Print(err)
	}
	defer func(logger *zap.Logger) {
		err = logger.Sync()
		if err != nil {
			log.Print(err)
		}
	}(logger)
	zapSugar := logger.Sugar()
	return zapSugar
}
