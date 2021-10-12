package usecase

import (
	"github.com/go-park-mail-ru/2021_2_A06367/internal/models"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/auth"
)

type AuthUsecase struct {
	repo      auth.AuthRepo
	tokenator auth.TokenGenerator
	encrypter auth.Encrypter
}

func NewAuthUsecase(repo auth.AuthRepo, tokenator auth.TokenGenerator, encrypter auth.Encrypter) *AuthUsecase {
	AuthUC := &AuthUsecase{repo: repo, tokenator: tokenator, encrypter: encrypter}
	return AuthUC
}

func (u *AuthUsecase) SignIn(user models.LoginUser) (string, models.StatusCode) {
	if user.Login == "" || user.EncryptedPassword == "" {
		return "", models.BadRequest
	}

	user.EncryptedPassword = u.encrypter.EncryptPswd(user.EncryptedPassword)
	DBUser, status := u.repo.CheckUser(models.User{Login: user.Login, EncryptedPassword: user.EncryptedPassword})

	token := u.tokenator.GetToken(DBUser)
	if status == models.Okey && token != "" {
		return token, status
	}
	return "", status
}

func (u *AuthUsecase) SignUp(user models.User) (string, models.StatusCode) {
	user.EncryptedPassword = u.encrypter.EncryptPswd(user.EncryptedPassword)
	_, st := u.repo.CheckUser(user)
	if st == models.Okey {
		return "", models.Conflict
	}

	NewUser, status := u.repo.CreateUser(user)
	token := u.tokenator.GetToken(NewUser)
	if status == models.Okey && token != "" {
		return token, status
	}
	return "", status
}
