package usecase

import (
	"github.com/go-park-mail-ru/2021_2_A06367/internal/models"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/films"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"testing"
)

func TestFilmsUsecase_GetFilm(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	uid := uuid.New()

	testFilm := models.Film{Id: uid}

	repo := films.NewMockFilmsRepository(ctl)
	repo.EXPECT().GetFilmById(testFilm).Times(1).Return(models.Film{}, models.Okey)

	usecase := NewFilmsUsecase(repo)

	_, st := usecase.GetFilm(testFilm)
	if st != models.Okey {
		t.Error("Wrong work of usecase")
	}
}

func TestFilmsUsecase_GetFilmsOfActor(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	uid := uuid.New()

	testActor := models.Actors{Id: uid}

	repo := films.NewMockFilmsRepository(ctl)
	repo.EXPECT().GetFilmsByActor(testActor).Times(1).Return([]models.Film{}, models.Okey)

	usecase := NewFilmsUsecase(repo)

	_, st := usecase.GetFilmsOfActor(testActor)
	if st != models.Okey {
		t.Error("Wrong work of usecase")
	}
}

func TestFilmsUsecase_GetCompilationForUser(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	uid := uuid.New()

	testUser := models.User{Id: uid}

	repo := films.NewMockFilmsRepository(ctl)
	repo.EXPECT().GetFilmsByUser(testUser).Times(1).Return([]models.Film{}, models.Okey)

	usecase := NewFilmsUsecase(repo)

	_, st := usecase.GetCompilationForUser(testUser)
	if st != models.Okey {
		t.Error("Wrong work of usecase")
	}
}
