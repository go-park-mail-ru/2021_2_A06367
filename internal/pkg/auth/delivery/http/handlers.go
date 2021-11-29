package http

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/models"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/auth/delivery/grpc/generated"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/middleware"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/utils"
	"github.com/gorilla/csrf"
	"github.com/mailru/easyjson"
	"go.uber.org/zap"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type AuthHandler struct {
	logger *zap.SugaredLogger
	client generated.AuthServiceClient
}

func NewAuthHandler(cl generated.AuthServiceClient, logger *zap.SugaredLogger) *AuthHandler {
	return &AuthHandler{
		client: cl,
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

	login, err := h.client.Login(context.Background(), &generated.LoginUser{
		Login:             user.Login,
		EncryptedPassword: user.EncryptedPassword,
	})
	if err != nil {
		utils.Response(w, models.Unauthed, nil)
	}
	if models.StatusCode(login.Status) != models.Okey {
		utils.Response(w, models.Unauthed, nil)
		return
	}
	SSCookie := &http.Cookie{
		Name:   "SSID",
		Value:  login.Cookie,
		Path:   "/",
		Domain: "a06367.ru",
		//SameSite: http.SameSiteNoneMode,
		HttpOnly: true,
		Expires:  time.Now().Add(time.Hour * 24),
	}
	http.SetCookie(w, SSCookie)
	utils.Response(w, models.StatusCode(login.Status), nil)
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
	//user := models.LoginUser{}
	//accesToken, err := utils.ExtractTokenMetadata(r, utils.ExtractTokenFromCookie)
	//if err != nil {
	//	utils.Response(w, models.BadRequest, nil)
	//	return
	//}
	//if accesToken == nil {
	//	utils.Response(w, models.Unauthed, nil)
	//	return
	//}
	//user.Login = accesToken.Login

	SSCookie := &http.Cookie{
		Name:     "SSID",
		Value:    "",
		Path:     "/",
		Domain:   "a06367.ru",
		HttpOnly: true,
		Expires:  time.Unix(0, 0),
	}
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

	us, err := h.client.SignUp(context.Background(), &generated.User{
		Login:             user.Login,
		EncryptedPassword: user.EncryptedPassword,
	})
	if err != nil {
		utils.Response(w, models.Unauthed, nil)
		return
	}

	if models.StatusCode(us.Status) != models.Okey || us.Cookie == "" || us.Cookie == "no secret key" { //TODO в константу
		utils.Response(w, models.StatusCode(us.Status), nil)
		return
	}

	SSCookie := &http.Cookie{
		Name:   "SSID",
		Value:  us.Cookie,
		Path:   "/",
		Domain: "a06367.ru",
		//SameSite: http.SameSiteNoneMode,
		HttpOnly: true,
		Expires:  time.Now().Add(time.Hour * 24),
	}

	http.SetCookie(w, SSCookie)
	w.WriteHeader(http.StatusOK)
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
	if err != nil && err.Error() != "no token" { //TODO в константу
		utils.Response(w, models.BadRequest, nil)
		return
	}

	if (err != nil && err.Error() == "no token") || jwtData.Login == "" {
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

	accesToken, err := utils.ExtractTokenMetadata(r, utils.ExtractTokenFromCookie)
	if err != nil {
		utils.Response(w, models.BadRequest, nil)
		return
	}
	if accesToken == nil {
		utils.Response(w, models.Unauthed, nil)
		return
	}
	uid := accesToken.Id

	//vars := mux.Vars(r)
	//id, found := vars["id"]
	//if !found {
	//	utils.Response(w, models.BadRequest, nil)
	//	return
	//}

	profile.Id = uid

	us, err := h.client.GetProfile(context.Background(), &generated.UserUUID{
		ID: profile.Id.String(),
	})
	if err != nil {
		utils.Response(w, models.InternalError, nil)
	}
	utils.Response(w, models.StatusCode(us.Status), models.Profile{
		Id:            uid,
		Login:         us.Login,
		About:         us.About,
		Avatar:        us.Avatar,
		Subscriptions: uint(us.Subscriptions),
		Subscribers:   uint(us.Subscribers),
	})
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
	/*
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

		//status := h.uc.Follow(uid, uid)
		utils.Response(w, status, nil)

	*/
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
	/*vars := mux.Vars(r)
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

	*/
}

func (h *AuthHandler) UpdateProfilePic(w http.ResponseWriter, r *http.Request) {
	jwtData, err := utils.ExtractTokenMetadata(r, utils.ExtractTokenFromCookie)
	if err != nil && err.Error() != "no token" {
		utils.Response(w, models.Unauthed, nil)
		return
	}

	err = r.ParseMultipartForm(5 * 1024 * 1025)
	if err != nil {
		utils.Response(w, models.BadRequest, nil)
		return
	}

	file, _, err := r.FormFile("pic")
	if err != nil {
		utils.Response(w, models.BadRequest, nil)
		return
	}

	all, err := ioutil.ReadAll(file)
	if err != nil {
		utils.Response(w, models.InternalError, nil)
		log.Fatal("readall", err)
		return
	}
	hash := md5.New()
	_, err = hash.Write(all)
	if err != nil {
		return
	}
	name := hash.Sum(nil)

	err = os.WriteFile("/image/"+hex.EncodeToString(name[:])+".png", all, 0644)
	if err != nil {
		utils.Response(w, models.InternalError, nil)

		log.Fatal("write", err)
		return
	}

	user := models.Profile{
		Id:     jwtData.Id,
		Login:  jwtData.Login,
		Avatar: hex.EncodeToString(name[:]) + ".png",
	}

	us, err := h.client.UpdateProfilePic(context.Background(), &generated.UserUpdatePic{
		Login:  user.Login,
		Avatar: user.Avatar,
		ID:     user.Id.String(),
	})
	if err != nil {
		utils.Response(w, models.Unauthed, nil)
	}
	utils.Response(w, models.StatusCode(us.Status), nil)
}

func (h *AuthHandler) UpdateProfilePass(w http.ResponseWriter, r *http.Request) {

	var pass models.PassUpdate
	err := easyjson.UnmarshalFromReader(r.Body, &pass)
	if err != nil {
		utils.Response(w, models.BadRequest, nil)
		return
	}

	jwtData, err := utils.ExtractTokenMetadata(r, utils.ExtractTokenFromCookie)
	if err != nil && err.Error() != "no token" {
		utils.Response(w, models.Unauthed, nil)
		return
	}

	user := models.User{
		EncryptedPassword: pass.Password,
		Id:                jwtData.Id,
		Login:             jwtData.Login,
	}

	us, err := h.client.UpdateProfilePass(context.Background(), &generated.UserUpdatePass{
		Login:    user.Login,
		Password: user.EncryptedPassword,
		ID:       user.Id.String(),
	})
	if err != nil {
		utils.Response(w, models.Unauthed, nil)
	}
	utils.Response(w, models.StatusCode(us.Status), nil)
}

func (h *AuthHandler) UpdateProfileBio(w http.ResponseWriter, r *http.Request) {

	var bio models.BioUpdate
	err := easyjson.UnmarshalFromReader(r.Body, &bio)
	if err != nil {
		utils.Response(w, models.BadRequest, nil)
		return
	}

	jwtData, err := utils.ExtractTokenMetadata(r, utils.ExtractTokenFromCookie)
	if err != nil && err.Error() != "no token" {
		utils.Response(w, models.Unauthed, nil)
		return
	}

	user := models.Profile{
		About: bio.About,
		Id:    jwtData.Id,
		Login: jwtData.Login,
	}

	us, err := h.client.UpdateProfileBio(context.Background(), &generated.UserUpdateBio{
		Login: user.Login,
		About: user.About,
		ID:    user.Id.String(),
	})
	if err != nil {
		utils.Response(w, models.Unauthed, nil)
	}
	utils.Response(w, models.StatusCode(us.Status), nil)
}
