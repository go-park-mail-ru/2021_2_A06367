package usecase

import (
	"github.com/go-park-mail-ru/2021_2_A06367/internal/models"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/actors"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"testing"
)


func TestActorsUsecase_GetActorsOfActor(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	uid := uuid.New().String()

	repo := actors.NewMockActorsRepository(ctl)
	repo.EXPECT().GetActorById(uid).Times(1).Return(models.Actors{}, models.Okey)

	usecase := NewActorsUsecase(repo)

	_, st := usecase.GetById(uid)
	if st != models.Okey {
		t.Error("Wrong work of usecase")
	}
}