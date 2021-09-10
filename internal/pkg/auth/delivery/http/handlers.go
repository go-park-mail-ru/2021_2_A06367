package http

import (
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/auth"
	"go.uber.org/zap"
	"net/http"
)

type AuthHandler struct {
	uc auth.AuthUsecase
	logger *zap.SugaredLogger
}

func NewAuthHandler(uc auth.AuthUsecase) *AuthHandler {
	return &AuthHandler{uc: uc}
}

func (h AuthHandler) Login(w http.ResponseWriter, r *http.Request)  {

}

func (h AuthHandler) Logout(w http.ResponseWriter, r *http.Request)  {

}

func (h AuthHandler) SignUp(w http.ResponseWriter, r *http.Request)  {

}
