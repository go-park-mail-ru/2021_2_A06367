package usecase

import (
	"github.com/go-park-mail-ru/2021_2_A06367/internal/models"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/auth"
)

type AuthUsecase struct {
	repo      auth.AuthRepo
	tokenator auth.TokenGenerator
}

func NewAuthUsecase(repo auth.AuthRepo, tokenator auth.TokenGenerator) *AuthUsecase {
	return &AuthUsecase{repo: repo, tokenator: tokenator}
}

func (u *AuthUsecase) SignIn(user models.LoginUser) (string, models.StatusCode) {
	if user.Login == "" || user.EncryptedPassword == "" {
		return "", models.BadRequest
	}
	repoUser := models.User{Login: user.Login, EncryptedPassword: user.EncryptedPassword}

	status := u.repo.CheckUser(repoUser)
	token := u.tokenator.GetToken(repoUser)
	if status == models.Okey && token != "" {
		return token, status
	}
	return "", status
}

func (u *AuthUsecase) SignUp(user models.User) (string, models.StatusCode) {
	if st := u.repo.CheckUser(user); st == models.Okey {
		return "", models.Conflict
	}
	status := u.repo.CreateUser(user)
	token := u.tokenator.GetToken(user)
	if status == models.Okey && token != "" {
		return token, status
	}
	return "", status
}
