package http

import (
	"github.com/go-park-mail-ru/2021_2_A06367/internal/models"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/auth"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/middleware"
	"github.com/mailru/easyjson"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type AuthHandler struct {
	uc     auth.AuthUsecase
	logger *zap.SugaredLogger
	online auth.OnlineUsecase
}

func NewAuthHandler(uc auth.AuthUsecase, ou auth.OnlineUsecase) *AuthHandler {
	return &AuthHandler{
		uc:     uc,
		online: ou,
	}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	user := models.LoginUser{}
	err := easyjson.UnmarshalFromReader(r.Body, &user)
	if err != nil || !middleware.LoginUserIsValid(user) {
		Response(w, models.Forbidden, nil)
		return
	}
	token, status := h.uc.SignIn(user)
	if status != models.Okey {
		Response(w, models.Unauthed, nil)
	}
	if status == models.Okey {
		status = h.online.Activate(user)
	}
	SSCookie := &http.Cookie{Name: "SSID", Value: token, HttpOnly: true, Secure: true}
	http.SetCookie(w, SSCookie)
	Response(w, status, models.TokenView{Token: token})
}

func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	user := models.LoginUser{}
	accesToken, err := ExtractTokenMetadata(r, ExtractTokenFromCookie)
	if err != nil || accesToken == nil {
		Response(w, models.BadRequest, nil)
		return
	}
	user.Login = accesToken.Login
	status := h.online.Deactivate(user)
	SSCookie := &http.Cookie{
		Name:     "SSID",
		Value:    "",
		HttpOnly: true,
		Expires:  time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC)}
	http.SetCookie(w, SSCookie)
	Response(w, status, nil)
}

func (h *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := easyjson.UnmarshalFromReader(r.Body, &user)
	if err != nil || !middleware.UserIsValid(user) {
		Response(w, models.BadRequest, nil)
		return
	}
	token, status := h.uc.SignUp(user)
	if token == "" || status != models.Okey {
		Response(w, status, nil)
		return
	}
	if status == models.Okey {
		userCopy := models.LoginUser{Login: user.Login, EncryptedPassword: user.EncryptedPassword}
		status = h.online.Activate(userCopy)
	}
	SSCookie := &http.Cookie{
		Name:     "SSID",
		Value:    token,
		HttpOnly: true,
		Expires:  time.Now().Add(time.Hour * 24)}
	http.SetCookie(w, SSCookie)
	Response(w, status, models.TokenView{Token: token})
}

func (h *AuthHandler) AuthStatus(w http.ResponseWriter, r *http.Request) {
	user := models.LoginUser{}
	user.Login = r.URL.Query().Get("user")
	jwtData, err := ExtractTokenMetadata(r, ExtractToken)
	if user.Login == "" || jwtData == nil {
		Response(w, models.BadRequest, nil)
		return
	}

	if err != nil || jwtData.Login != user.Login {
		Response(w, models.Unauthed, nil)
		return
	}
	status := h.online.IsAuthed(user)
	if !status {
		Response(w, models.Unauthed, nil)
		return
	}
	Response(w, models.Okey, nil)
}
