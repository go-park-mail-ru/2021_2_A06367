package usecase

import (
	"github.com/go-park-mail-ru/2021_2_A06367/internal/models"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/auth"
	"github.com/google/uuid"
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
	if status == models.Okey {
		return u.tokenator.GetToken(repoUser), status
	} else {
		return "", status
	}
}

func (u *AuthUsecase) SignUp(user models.User) (string, models.StatusCode) {
	if st := u.repo.CheckUser(user); st == models.Okey {
		return "", models.Conflict
	}
	status := u.repo.CreateUser(user)

	if status == models.Okey {
		return u.tokenator.GetToken(user), status
	} else {
		return "", status
	}
}

func (u *AuthUsecase) GetProfile(user models.Profile) (models.Profile, models.StatusCode) {
	return u.repo.GetProfile(user)
}

func (u *AuthUsecase) Follow(who, whom uuid.UUID) models.StatusCode {
	if who == whom {
		return models.Forbidden
	}
	return u.repo.AddFollowing(who, whom)
}
