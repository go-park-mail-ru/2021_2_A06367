package models

import "github.com/dgrijalva/jwt-go"

//easyjson:skip
type Token struct {
	Login string
	jwt.StandardClaims
}
