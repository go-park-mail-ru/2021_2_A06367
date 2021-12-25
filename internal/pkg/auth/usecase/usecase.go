package usecase

import (
	"github.com/go-park-mail-ru/2021_2_A06367/internal/models"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/auth"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type AuthUsecase struct {
	repo      auth.AuthRepo
	tokenator auth.TokenGenerator
	encrypter auth.Encrypter
	logger    *zap.SugaredLogger
}

func NewAuthUsecase(repo auth.AuthRepo, tokenator auth.TokenGenerator, encrypter auth.Encrypter, logger *zap.SugaredLogger) *AuthUsecase {
	AuthUC := &AuthUsecase{
		repo:      repo,
		tokenator: tokenator,
		encrypter: encrypter,
		logger:    logger,
	}
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
	if st == models.Okey || st == models.Unauthed {
		return "", models.Conflict
	}

	NewUser, status := u.repo.CreateUser(user)
	token := u.tokenator.GetToken(NewUser)
	if status == models.Okey && token != "" {
		return token, status
	}
	return "", status
}

func (u *AuthUsecase) CheckUser(user models.User) (models.User, models.StatusCode) {
	return u.repo.CheckUser(user)
}
func (u *AuthUsecase) CheckUserLogin(user models.User) (models.User, models.StatusCode) {
	return u.repo.CheckUserLogin(user)
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

func (u *AuthUsecase) Unfollow(who, whom uuid.UUID) models.StatusCode {
	if who == whom {
		return models.Forbidden
	}
	return u.repo.RemoveFollowing(who, whom)
}

func (u *AuthUsecase) GetSubscriptions() ([]models.Profile, models.StatusCode) {
	return nil, models.Okey
}

func (u *AuthUsecase) GetSubscribers() ([]models.Profile, models.StatusCode) {
	return nil, models.Okey
}

func (u *AuthUsecase) GetByKeyword(keyword string) ([]models.Profile, models.StatusCode) {
	return u.repo.GetProfileByKeyword(keyword)
}

func (u *AuthUsecase) SetBio(profile models.Profile) models.StatusCode {
	return u.repo.UpdateBio(profile)
}

func (u *AuthUsecase) SetPass(profile models.User) models.StatusCode {
	profile.EncryptedPassword = u.encrypter.EncryptPswd(profile.EncryptedPassword)
	return u.repo.UpdatePass(profile)
}

func (u *AuthUsecase) SetAvatar(profile models.Profile) models.StatusCode {
	return u.repo.UpdateAvatar(profile)
}
