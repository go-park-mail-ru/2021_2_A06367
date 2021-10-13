package usecase

import (
	"github.com/go-park-mail-ru/2021_2_A06367/internal/models"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/auth"
)

type OnlineUsecase struct {
	onlineRepo auth.OnlineRepo
}

func NewOnlineUsecase(repo auth.OnlineRepo) *OnlineUsecase {
	return &OnlineUsecase{
		onlineRepo: repo,
	}
}

func (ou *OnlineUsecase) Activate(user models.LoginUser) models.StatusCode {
	return ou.onlineRepo.UserOn(user)
}

func (ou *OnlineUsecase) Deactivate(user models.LoginUser) models.StatusCode {
	return ou.onlineRepo.UserOff(user)
}

func (ou *OnlineUsecase) IsAuthed(user models.LoginUser) bool {
	return ou.onlineRepo.IsAuthed(user)
}
