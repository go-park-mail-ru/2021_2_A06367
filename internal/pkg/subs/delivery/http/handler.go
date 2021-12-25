package http

import (
	"context"
	"fmt"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/models"
	generated2 "github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/auth/delivery/grpc/generated"
	subs "github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/subs/delivery/grpc/generated"
	util "github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/utils"
	"net/http"
	"time"
)

type SubsHandler struct {
	subsClient subs.SubsServiceClient
	cl         generated2.AuthServiceClient
}

func NewSubsHandler(subsClient subs.SubsServiceClient, cl generated2.AuthServiceClient) *SubsHandler {
	return &SubsHandler{subsClient: subsClient, cl:cl }
}

func (h SubsHandler) GetLicense(w http.ResponseWriter, r *http.Request) {
	access, err := util.ExtractTokenMetadata(r, util.ExtractTokenFromCookie)
	if err != nil || access == nil {
		util.Response(w, models.BadRequest, nil)
	}

	l, err := h.subsClient.GetLicense(context.Background(), &subs.LicenseUUID{ID: access.Id.String()})
	if err != nil {
		util.Response(w, models.InternalError, nil)
	}
	parsed, err := time.Parse("2006-01-02", l.ExpiresDate[:10])
	if err != nil {
		util.Response(w, models.InternalError, nil)
	}
	license := models.License{IsValid: true, ExpDate: parsed}
	util.Response(w, models.Okey, license)
}

func (h SubsHandler) SetLicense(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		util.Response(w, models.BadRequest, nil)
		return
	}

	data := r.PostFormValue("label")
	if data == "" {
		util.Response(w, models.BadRequest, nil)
		return
	}
	profile, err := h.cl.CheckByLogin(context.Background(), &generated2.LoginUser{Login: data})
	if err != nil {
		util.Response(w, models.InternalError, nil)
		return
	}

	fmt.Print(profile)
	_, err = h.subsClient.SetLicense(context.Background(), &subs.LicenseReq{
		ID:   profile.ID,
		Type: "MONTH"})
	//если микросервис отвалился
	if err != nil {
		util.Response(w, models.InternalError, nil)
	}
	util.Response(w, models.Okey, nil)
}
