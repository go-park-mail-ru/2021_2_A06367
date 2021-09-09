package http

import "github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/auth"

type AuthHandler struct {
	uc auth.AuthUsecase
}

func NewAuthHandler(uc auth.AuthUsecase) *AuthHandler {
	return &AuthHandler{uc: uc}
}
