package http

import (
	"github.com/go-park-mail-ru/2021_2_A06367/internal/models"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/auth"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/middleware"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/utils"
	uuid "github.com/google/uuid"
	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
	"github.com/mailru/easyjson"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type AuthHandler struct {
	uc     auth.AuthUsecase
	logger *zap.SugaredLogger
}

func NewAuthHandler(uc auth.AuthUsecase, logger *zap.SugaredLogger) *AuthHandler {
	return &AuthHandler{
		uc:     uc,
		logger: logger,
	}
}

// Login godoc
// @Summary Get login
// @Description Get login
// @Tags Users
// @Accept json
// @Produce json
// @Param order body models.LoginUser true "Create order"
// @Success 200 {string} 1
// @Header 200 {string} Token "SSID"
// @Failure 400,403,404 {string} 1
// @Router /user/login [post]
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
	SSCookie := &http.Cookie{Name: "SSID", Value: token, HttpOnly: true}
	http.SetCookie(w, SSCookie)
	utils.Response(w, status, nil)
}

func (h *AuthHandler) Token(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-CSRF-Token", csrf.Token(r))
	utils.Response(w, models.Okey, nil)
}

// Logout godoc
// @Summary Get logout
// @Description Get logout
// @Tags Users
// @Accept json
// @Produce json
// @Param order body models.LoginUser true "Create order"
// @Success 200 {string} 1
// @Header 200 {string} Token ""
// @Failure 400 {string} 1
// @Router /user/logout [post]
// @Router /user/logout [options]
func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	user := models.LoginUser{}
	accesToken, err := utils.ExtractTokenMetadata(r, utils.ExtractTokenFromCookie)
	if err != nil {
		utils.Response(w, models.BadRequest, nil)
		return
	}
	if accesToken == nil {
		utils.Response(w, models.Unauthed, nil)
		return
	}
	user.Login = accesToken.Login

	SSCookie := &http.Cookie{
		Name:     "SSID",
		Value:    "",
		HttpOnly: true,
		Expires:  time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC)}
	http.SetCookie(w, SSCookie)
	utils.Response(w, models.Okey, nil)
}

// SignUp godoc
// @Summary Get sign up
// @Description Get sign up
// @Tags Users
// @Accept json
// @Produce json
// @Param order body models.LoginUser true "Create order"
// @Success 200 {string} 1
// @Header 200 {string} Token "SSID"
// @Failure 400 {string} 1
// @Router /user/signup [post]
func (h *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := easyjson.UnmarshalFromReader(r.Body, &user)
	if err != nil || !middleware.UserIsValid(user) {
		utils.Response(w, models.BadRequest, nil)
		return
	}
	token, status := h.uc.SignUp(user)
	if token == "" || token == models.ErrNoSecretKey || status != models.Okey { //TODO в константу
		utils.Response(w, status, nil)
		return
	}

	SSCookie := &http.Cookie{
		Name:     "SSID",
		Value:    token,
		HttpOnly: true,
		Expires:  time.Now().Add(time.Hour * 24)}

	http.SetCookie(w, SSCookie)
	utils.Response(w, status, nil)
}

// AuthStatus godoc
// @Summary Get check auth status
// @Description Get check auth status
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {string} 1
// @Failure 400 {string} 1
// @Router /user/auth [get]
func (h *AuthHandler) AuthStatus(w http.ResponseWriter, r *http.Request) {
	user := models.LoginUser{}
	user.Login = r.URL.Query().Get("user")
	jwtData, err := utils.ExtractTokenMetadata(r, utils.ExtractTokenFromCookie)
	if err != nil && err.Error() != models.ErrNoToken {
		utils.Response(w, models.BadRequest, nil)
		return
	}

	if err.Error() == models.ErrNoToken || jwtData.Login != user.Login || user.Login == "" {
		utils.Response(w, models.Unauthed, nil)
		return
	}
	utils.Response(w, models.Okey, nil)
}

// GetProfile godoc
// @Summary Get details of profile
// @Description Get details of profile
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "768eb570-2e0e-11ec-8d3d-0242ac130004"
// @Success 200 {array} models.Profile
// @Failure 400,404 {string} 1
// @Router /user/profile/{id} [get]
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

// Follow godoc
// @Summary Subscribe
// @Description Subscribe
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "768eb570-2e0e-11ec-8d3d-0242ac130004"
// @Success 200 {string} 1
// @Failure 400,404 {string} 1
// @Router /user/profile/{id}/follow [post]
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

// Unfollow godoc
// @Summary Unsubscribe
// @Description Unsubscribe
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "768eb570-2e0e-11ec-8d3d-0242ac130004"
// @Success 200 {string} 1
// @Failure 400,404 {string} 1
// @Router /user/profile/{id}/unfollow [delete]
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
