package auth

import "github.com/go-park-mail-ru/2021_2_A06367/internal/models"

//go:generate mockgen -source=interfaces.go -destination=interfaces_mock.go -package=auth

type AuthUsecase interface {
	SignIn(user models.LoginUser) (string, models.StatusCode)
	SignUp(user models.User) (string, models.StatusCode)
}

type AuthRepo interface {
	CreateUser(user models.User) models.StatusCode
	CheckUser(user models.User) models.StatusCode
}

type OnlineRepo interface {
	UserOn(user models.User) models.StatusCode
	UserOff(user models.User) models.StatusCode
	IsOnline(user models.User) bool
}
