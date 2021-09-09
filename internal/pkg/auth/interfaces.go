package auth

import "github.com/go-park-mail-ru/2021_2_A06367/internal/models"

//go:generate mockgen -source=interfaces.go -destination=interfaces_mock.go -package=auth

type AuthUsecase interface {
	SignIn(user models.User) models.StatusCode
	SignUp(user models.User) (int, models.StatusCode)
}

type AuthRepo interface {
	CreateUser(user models.User)  (int,models.StatusCode)
	CheckUser(user models.User) int
}
