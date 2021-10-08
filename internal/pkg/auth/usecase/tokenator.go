package usecase

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/models"
	"os"
	"time"
)

type Tokenator struct {
}

func NewTokenator() *Tokenator {
	return &Tokenator{}
}

func (t *Tokenator) GetToken(user models.User) string {
	tokenModel := models.Token{
		Login: user.Login,
		Id:    user.Id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 5).Unix(),
		},
	}

	SecretKey, err := os.LookupEnv("SECRET")
	if !err {
		return "no secret key"
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenModel)

	jwtCookie, _ := token.SignedString([]byte(SecretKey))
	return jwtCookie
}
