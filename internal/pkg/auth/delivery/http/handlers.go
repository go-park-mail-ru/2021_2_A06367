package http

import (
	"github.com/go-park-mail-ru/2021_2_A06367/internal/models"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/auth"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/middleware"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/utils"
	uuid "github.com/google/uuid"
	"github.com/gorilla/mux"
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

func NewAuthHandler(uc auth.AuthUsecase, ou auth.OnlineUsecase, logger *zap.SugaredLogger) *AuthHandler {
	return &AuthHandler{
		uc:     uc,
		online: ou,
		logger: logger,
	}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	user := models.LoginUser{}
	err := easyjson.UnmarshalFromReader(r.Body, &user)
	if err != nil || !middleware.LoginUserIsValid(user) {
		utils.Response(w, models.Forbidden, nil)
		return
	}
	token, status := h.uc.SignIn(user)
	if status != models.Okey {
		utils.Response(w, models.Unauthed, nil)
		return
	}
	if status == models.Okey {
		status = h.online.Activate(user)
		return
	}
	SSCookie := &http.Cookie{Name: "SSID", Value: token, HttpOnly: true, Secure: true}
	http.SetCookie(w, SSCookie)
	utils.Response(w, status, nil)
}

func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	user := models.LoginUser{}
	accesToken, err := utils.ExtractTokenMetadata(r, utils.ExtractTokenFromCookie)
	if err != nil || accesToken == nil {
		utils.Response(w, models.BadRequest, nil)
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
	utils.Response(w, status, nil)
}

func (h *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := easyjson.UnmarshalFromReader(r.Body, &user)
	if err != nil || !middleware.UserIsValid(user) {
		utils.Response(w, models.BadRequest, nil)
		return
	}
	token, status := h.uc.SignUp(user)
	if token == "" || status != models.Okey {
		utils.Response(w, status, nil)
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
	utils.Response(w, status, nil)
}

func (h *AuthHandler) AuthStatus(w http.ResponseWriter, r *http.Request) {
	user := models.LoginUser{}
	user.Login = r.URL.Query().Get("user")
	jwtData, err := utils.ExtractTokenMetadata(r, utils.ExtractToken)
	if user.Login == "" || jwtData == nil {
		utils.Response(w, models.BadRequest, nil)
		return
	}

	if err != nil || jwtData.Login != user.Login {
		utils.Response(w, models.Unauthed, nil)
		return
	}
	status := h.online.IsAuthed(user)
	if !status {
		utils.Response(w, models.Unauthed, nil)
		return
	}
	utils.Response(w, models.Okey, nil)
}

func (h *AuthHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	profile := models.Profile{}

	vars := mux.Vars(r)
	id, found := vars["id"]
	if !found {
		utils.Response(w, models.BadRequest, nil)
		return
	}
	uid, err := uuid.Parse(id)
	if err != nil {
		utils.Response(w, models.BadRequest, profile)
		return
	}
	profile.Id = uid

	user, status := h.uc.GetProfile(profile)
	utils.Response(w, status, user)
}

func (h *AuthHandler) Follow(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, found := vars["id"]
	if !found {
		utils.Response(w, models.BadRequest, nil)
		return
	}
	uid, err := uuid.Parse(id)
	if err != nil {
		utils.Response(w, models.BadRequest, nil)
		return
	}

	status := h.uc.Follow(uid, uid)
	utils.Response(w, status, nil)
}

func (h *AuthHandler) Unfollow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, found := vars["id"]
	if !found {
		utils.Response(w, models.BadRequest, nil)
		return
	}
	uid, err := uuid.Parse(id)
	if err != nil {
		utils.Response(w, models.BadRequest, nil)
		return
	}

	status := h.uc.Follow(uid, uid)
	utils.Response(w, status, nil)
}
