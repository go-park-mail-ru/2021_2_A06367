package usecase

import (
	"github.com/go-park-mail-ru/2021_2_A06367/internal/models"
	"github.com/golang-jwt/jwt"
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
		Id:    user.Id.String(),
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
