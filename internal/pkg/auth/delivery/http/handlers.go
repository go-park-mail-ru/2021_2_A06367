package http

import (
	"github.com/go-park-mail-ru/2021_2_A06367/internal/models"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/auth"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/middleware"
	"github.com/gorilla/mux"
	"github.com/mailru/easyjson"
	"go.uber.org/zap"
	"net/http"
)

type AuthHandler struct {
	uc     auth.AuthUsecase
	logger *zap.SugaredLogger
	online auth.OnlineRepo
}

func NewAuthHandler(uc auth.AuthUsecase, or auth.OnlineRepo) *AuthHandler {
	return &AuthHandler{
		uc:     uc,
		online: or,
	}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	user := models.LoginUser{}
	err := easyjson.UnmarshalFromReader(r.Body, &user)
	if err != nil || !middleware.LoginUserIsValid(user) {
		middleware.Response(w, models.BadRequest, nil)
		return
	}
	token, status := h.uc.SignIn(user)
	if status == models.Okey {
		status = h.online.UserOn(user)
	}
	middleware.Response(w, status, models.TokenView{Token: token})
}

func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	user := models.LoginUser{}
	accesToken, err := middleware.ExtractTokenMetadata(r, middleware.ExtractToken)
	if err != nil || !middleware.LoginUserIsValid(user) {
		middleware.Response(w, models.BadRequest, nil)
		return
	}

	user.Login = accesToken.Login
	if user.Login == "" {
		middleware.Response(w, models.BadRequest, nil)
		return
	}
	status := h.online.UserOff(user)
	middleware.Response(w, status, nil)
}

func (h *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := easyjson.UnmarshalFromReader(r.Body, &user)
	if err != nil || !middleware.UserIsValid(user) {
		middleware.Response(w, models.BadRequest, nil)
		return
	}
	token, status := h.uc.SignUp(user)
	if status == models.Okey {
		userCopy := models.LoginUser{Login: user.Login, EncryptedPassword: user.EncryptedPassword}
		status = h.online.UserOn(userCopy)
	}

	if token == "" {
		middleware.Response(w, status, nil)
		return
	}
	middleware.Response(w, status, models.TokenView{Token: token})
}

func (h *AuthHandler) AuthStatus(w http.ResponseWriter, r *http.Request) {
	user := models.LoginUser{}
	Var := mux.Vars(r)

	user.Login = Var["user"]
	jwtData, err := middleware.ExtractTokenMetadata(r, middleware.ExtractTokenFromHeader)

	if user.Login == "" {
		middleware.Response(w, models.BadRequest, nil)
		return
	}

	if err != nil || jwtData.Login != user.Login {
		middleware.Response(w, models.Unauthed, nil)
	}

	status := h.online.IsAuthed(user)
	if !status {
		middleware.Response(w, models.Unauthed, nil)
		return
	}
	middleware.Response(w, models.Okey, nil)
}
