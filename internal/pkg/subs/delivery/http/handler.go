package http

import (
	"context"
	"fmt"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/models"
	subs "github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/subs/delivery/grpc/generated"
	util "github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/utils"
	"net/http"
	"time"
)

type SubsHandler struct {
	subsClient subs.SubsServiceClient
}

func NewSubsHandler(subsClient subs.SubsServiceClient) *SubsHandler {
	return &SubsHandler{subsClient: subsClient}
}

func (h SubsHandler) GetLicense(w http.ResponseWriter, r *http.Request) {
	access, err := util.ExtractTokenMetadata(r, util.ExtractTokenFromCookie)
	if err != nil || access == nil {
		util.Response(w, models.BadRequest, nil)
	}

	l, err := h.subsClient.GetLicense(context.Background(), &subs.LicenseUUID{ID: access.Id.String()})
	//если микросервис отвалился
	if err != nil {
		util.Response(w, models.InternalError, nil)
	}
	parsed, _ := time.Parse(time.RFC3339, l.ExpiresDate)
	license := models.License{IsValid: true, ExpDate: parsed}
	util.Response(w, models.Okey, license)
}

func (h SubsHandler) SetLicense(w http.ResponseWriter, r *http.Request) {

	label := r.URL.Query().Get("label")
	fmt.Println(label)
	fmt.Println(r.URL.Query())
	util.Response(w, models.Okey, nil)
	return

	/*
		access, err := util.ExtractTokenMetadata(r, util.ExtractTokenFromCookie)
		if err != nil || access == nil {
			util.Response(w, models.BadRequest, nil)
		}

		l, err := h.subsClient.SetLicense(context.Background(), &subs.LicenseReq{
			ID:   access.Id.String(),
			Type: "MONTH"})
		//если микросервис отвалился
		if err != nil {
			util.Response(w, models.InternalError, nil)
		}
		parsed, _ := time.Parse(time.RFC3339, l.ExpiresDate)
		license := models.License{IsValid: true, ExpDate: parsed}
		util.Response(w, models.Okey, license)
	*/
}
