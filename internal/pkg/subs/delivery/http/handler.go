package http

import (
	"context"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/models"
	subs "github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/subs/delivery/grpc/generated"
	util "github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/utils"
	"github.com/mailru/easyjson"
	"io/ioutil"
	"log"
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
	var check models.UMoney

	all, _ := ioutil.ReadAll(r.Body)
	log.Println(string(all))

	_ = easyjson.UnmarshalFromReader(r.Body, &check)
	util.Response(w, models.Okey, nil)

	log.Println(string(all))
	log.Println(check)
	log.Println(check.Label)
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
