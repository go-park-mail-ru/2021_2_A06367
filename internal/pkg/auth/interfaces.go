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
	GetByKeyword(keyword string) ([]models.Profile, models.StatusCode)
	SetBio(profile models.Profile) models.StatusCode
	SetPass(profile models.User) models.StatusCode
	SetAvatar(profile models.Profile) models.StatusCode
}

type AuthRepo interface {
	CreateUser(user models.User) (models.User, models.StatusCode)
	CheckUser(user models.User) (models.User, models.StatusCode)
	GetProfile(user models.Profile) (models.Profile, models.StatusCode)
	AddFollowing(who, whom uuid.UUID) models.StatusCode
	RemoveFollowing(who, whom uuid.UUID) models.StatusCode
	GetProfileByKeyword(keyword string) ([]models.Profile, models.StatusCode)
	UpdateBio(profile models.Profile) models.StatusCode
	UpdatePass(profile models.User) models.StatusCode
	UpdateAvatar(profile models.Profile) models.StatusCode
}

type TokenGenerator interface {
	GetToken(user models.User) string
}

type Encrypter interface {
	EncryptPswd(pswd string) string
}
