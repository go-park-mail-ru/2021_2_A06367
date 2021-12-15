package grpc

import (
	"context"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/models/grpc"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/subs"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/subs/delivery/grpc/generated"
	"github.com/google/uuid"
)

type GrpcSubsHandler struct {
	uc subs.SubsUsecase
	generated.SubsServiceServer
}

func NewGrpcSubsHandler(uc subs.SubsUsecase) *GrpcSubsHandler {
	return &GrpcSubsHandler{uc: uc}
}

func (h GrpcSubsHandler) GetLicense(ctx context.Context, in *generated.LicenseUUID) (*generated.License, error) {
	id, err := uuid.Parse(in.ID)
	if err != nil {
		return &generated.License{Status: grpc.StatusCode_BadRequest}, nil
	}
	l, status := h.uc.GetLicense(id)
	return &generated.License{
		IsValid:     l.IsValid,
		ExpiresDate: l.ExpDate.String(),
		Status:      grpc.StatusCode(status),
	}, nil
}

func (h GrpcSubsHandler) SetLicense(ctx context.Context, in *generated.LicenseReq) (*generated.License, error) {
	id, err := uuid.Parse(in.ID)
	if err != nil {
		return &generated.License{Status: grpc.StatusCode_BadRequest}, nil
	}
	l, status := h.uc.SetLicense(id, in.Type)
	return &generated.License{
		IsValid:     l.IsValid,
		ExpiresDate: l.ExpDate.Format("2006-01-02 15:04:05"),
		Status:      grpc.StatusCode(status),
	}, nil
}
