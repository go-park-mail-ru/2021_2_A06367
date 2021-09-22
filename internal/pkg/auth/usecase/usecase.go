package usecase

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/models"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/auth"
	"github.com/google/uuid"
	"os"
	"time"
)

type AuthUsecase struct {
	repo auth.AuthRepo
}

func NewAuthUsecase(repo auth.AuthRepo) *AuthUsecase {
	return &AuthUsecase{repo: repo}
}

func (u *AuthUsecase) SignIn(user models.LoginUser) (string, models.StatusCode) {
	if user.Login == "" || user.EncryptedPassword == "" {
		return "", models.Unauthed
	}
	repoUser := models.User{Login: user.Login, EncryptedPassword: user.EncryptedPassword}

	status := u.repo.CheckUser(repoUser)
	if status == models.Okey {
		return u.GetToken(repoUser), status
	} else {
		return "", status
	}
}

func (u *AuthUsecase) SignUp(user models.User) (string, models.StatusCode) {
	if st := u.repo.CheckUser(user); st == models.Okey {
		return "", models.Conflict
	}
	user.Id = uuid.New()
	status := u.repo.CreateUser(user)

	if status == models.Okey {
		return u.GetToken(user), status
	} else {
		return "", status
	}
}

func (u *AuthUsecase) GetToken(user models.User) string {
	tokenModel := models.Token{
		Login: user.Login,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 5).Unix(),
		},
	}

	SecretKey, err := os.LookupEnv("SECRET")
	if !err {
		panic("where is a secret key!")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenModel)

	jwtCookie, _ := token.SignedString([]byte(SecretKey))
	return jwtCookie
}
