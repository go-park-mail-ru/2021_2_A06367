package middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/models"
	"net/http"
	"os"
	"strings"
	"time"
)

type AccessDetails struct {
	Login string
}

type Extrater func(r *http.Request) string

func ExtractToken(r *http.Request) string {
	token := models.TokenView{}
	err := json.NewDecoder(r.Body).Decode(&token)
	if err != nil {
		return ""
	}

	strArr := strings.Split(token.Token, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return strArr[0]
}

func ExtractTokenFromHeader(r *http.Request) string {
	token := r.Header.Get("Authorization")
	strArr := strings.Split(token, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return strArr[0]
}

func VerifyToken(r *http.Request, extrater Extrater) (*models.Token, error) {
	tokenStr := extrater(r)
	if tokenStr == "" {
		return nil, errors.New("No token")
	}

	token, err := jwt.ParseWithClaims(tokenStr, &models.Token{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	claim := token.Claims
	claims, ok := claim.(*models.Token)
	if ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("no auth data")
}

func ExtractTokenMetadata(r *http.Request, extrater Extrater) (*AccessDetails, error) {
	token, err := VerifyToken(r, extrater)
	if err != nil {
		return nil, err
	}
	exp := token.ExpiresAt
	now := time.Now().Unix()
	if exp < now {
		return nil, errors.New("Token Exprired")
	}
	data := &AccessDetails{Login: token.Login}
	if data.Login == "" {
		return nil, errors.New("Invalid token")
	}

	return data, err
}
