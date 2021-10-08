package http

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

type Extracter func(r *http.Request) string

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

func ExtractTokenFromCookie(r *http.Request) string {
	tokenCookie, _ := r.Cookie("SSID")
	token := tokenCookie.Value
	strArr := strings.Split(token, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return strArr[0]
}

func VerifyToken(r *http.Request, extracter Extracter) (*models.Token, error) {
	tokenStr := extracter(r)
	if tokenStr == "" {
		return nil, errors.New("no token")
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

func ExtractTokenMetadata(r *http.Request, extracter Extracter) (*models.AccessDetails, error) {
	token, err := VerifyToken(r, extracter)
	if err != nil {
		return nil, err
	}
	exp := token.ExpiresAt
	now := time.Now().Unix()
	if exp < now {
		return nil, errors.New("token expired")
	}
	data := &models.AccessDetails{Login: token.Login, Id: token.Id}
	if data.Login == "" || data.Id.String() == "" {
		return nil, errors.New("invalid token")
	}

	return data, err
}
