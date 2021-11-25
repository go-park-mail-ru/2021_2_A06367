package usecase

import (
	"github.com/go-park-mail-ru/2021_2_A06367/internal/models"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/films/mocks"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"testing"
)

func TestFilmsUsecase_GetFilm(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	uid := uuid.New()

	testFilm := models.Film{Id: uid}

	repo := mocks.NewMockFilmsRepository(ctl)
	repo.EXPECT().GetFilmById(testFilm).Times(1).Return(models.Film{}, models.Okey)

	usecase := NewFilmsUsecase(repo, nil)

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

	repo := mocks.NewMockFilmsRepository(ctl)
	repo.EXPECT().GetFilmsByActor(testActor).Times(1).Return([]models.Film{}, models.Okey)

	usecase := NewFilmsUsecase(repo, nil)

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

	repo := mocks.NewMockFilmsRepository(ctl)
	repo.EXPECT().GetFilmsByUser(testUser).Times(1).Return([]models.Film{}, models.Okey)

	usecase := NewFilmsUsecase(repo, nil)

	_, st := usecase.GetCompilationForUser(testUser)
	if st != models.Okey {
		t.Error("Wrong work of usecase")
	}
}

func TestFilmsUsecase_GetCompilation(t *testing.T) {
	topic := "tests"
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mocks.NewMockFilmsRepository(ctl)
	repo.EXPECT().GetFilmsByTopic(topic).Times(1).Return([]models.Film{}, models.Okey)

	usecase := NewFilmsUsecase(repo, nil)

	_, st := usecase.GetCompilation(topic)
	if st != models.Okey {
		t.Error("Wrong work of usecase")
	}
}

func TestFilmsUsecase_GetByKeyword(t *testing.T) {
	keyword := "tests"
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mocks.NewMockFilmsRepository(ctl)
	repo.EXPECT().GetFilmsByKeyword(keyword).Times(1).Return([]models.Film{}, models.Okey)

	usecase := NewFilmsUsecase(repo, nil)

	_, st := usecase.GetByKeyword(keyword)
	if st != models.Okey {
		t.Error("Wrong work of usecase")
	}
}

func TestFilmsUsecase_GetSelection(t *testing.T) {
	selection1 := "hottest"
	selection2 := "not hottest"
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mocks.NewMockFilmsRepository(ctl)
	repo.EXPECT().GetHottestFilms().Times(1).Return([]models.Film{}, models.Okey)
	repo.EXPECT().GetNewestFilms().Times(1).Return([]models.Film{}, models.Okey)

	usecase := NewFilmsUsecase(repo, nil)

	_, st := usecase.GetSelection(selection1)
	if st != models.Okey {
		t.Error("Wrong work of usecase")
	}

	_, st2 := usecase.GetSelection(selection2)
	if st2 != models.Okey {
		t.Error("Wrong work of usecase")
	}
}

func TestFilmsUsecase_GetStartSelections(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mocks.NewMockFilmsRepository(ctl)
	repo.EXPECT().GetHottestFilms().Times(1).Return([]models.Film{}, models.Okey)

	usecase := NewFilmsUsecase(repo, nil)

	_, st := usecase.GetStartSelections(false, models.User{})
	if st != models.Okey {
		t.Error("Wrong work of usecase")
	}

	repo.EXPECT().GetFilmsByUser(models.User{}).Return([]models.Film{}, models.Okey)
	_, st2 := usecase.GetStartSelections(true, models.User{})
	if st2 != models.Okey {
		t.Error("Wrong work of usecase")
	}
}
