package auth

import (
	"github.com/go-park-mail-ru/2021_2_A06367/internal/models"
	"github.com/google/uuid"
)

//go:generate mockgen -source=interfaces.go -destination=interfaces_mock.go -package=auth

type AuthUsecase interface {
	SignIn(user models.LoginUser) (string, models.StatusCode)
	SignUp(user models.User) (string, models.StatusCode)
	GetProfile(user models.Profile) (models.Profile, models.StatusCode)
	Follow(who, whom uuid.UUID) models.StatusCode
}

type AuthRepo interface {
	CreateUser(user models.User) models.StatusCode
	CheckUser(user models.User) models.StatusCode
	GetProfile(user models.Profile) (models.Profile, models.StatusCode)
	AddFollowing(who, whom uuid.UUID) models.StatusCode
}

type OnlineRepo interface {
	UserOn(user models.LoginUser) models.StatusCode
	UserOff(user models.LoginUser) models.StatusCode
	IsAuthed(user models.LoginUser) bool
}

type TokenGenerator interface {
	GetToken(user models.User) string
}
