package usecase

import (
	"github.com/go-park-mail-ru/2021_2_A06367/internal/models"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/actors/mocks"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"testing"
)

func TestActorsUsecase_GetActorsOfActor(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	uid := uuid.New()

	testActor := models.Actors{Id: uid}

	repo := mocks.NewMockActorsRepository(ctl)
	repo.EXPECT().GetActorById(testActor).Times(1).Return(models.Actors{}, models.Okey)

	usecase := NewActorsUsecase(repo, nil)

	_, st := usecase.GetById(testActor)
	if st != models.Okey {
		t.Error("Wrong work of usecase")
	}
}

func TestActorsUsecase_GetByActors(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	uid := uuid.New()

	testActor := models.Actors{Id: uid}

	actorsArr := []models.Actors{testActor}

	repo := mocks.NewMockActorsRepository(ctl)
	repo.EXPECT().GetActors(actorsArr).Times(1).Return(actorsArr, models.Okey)

	usecase := NewActorsUsecase(repo, nil)

	_, st := usecase.GetByActors(actorsArr)
	if st != models.Okey {
		t.Error("Wrong work of usecase")
	}
}
