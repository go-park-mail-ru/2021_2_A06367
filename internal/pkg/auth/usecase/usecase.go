package usecase

import (
	"github.com/go-park-mail-ru/2021_2_A06367/internal/models"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/auth"
	"github.com/google/uuid"
)

type AuthUsecase struct {
	repo auth.AuthRepo
}

func NewAuthUsecase(repo auth.AuthRepo) *AuthUsecase {
	return &AuthUsecase{repo: repo}
}

func (u *AuthUsecase) SignIn(user models.User) models.StatusCode {
	return u.repo.CheckUser(user)
}

func (u *AuthUsecase) SignUp(user models.User) models.StatusCode  {
	user.Id = uuid.New()
	return u.repo.CreateUser(user)
}
