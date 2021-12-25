package grpc

import (
	"context"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/models"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/subs"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/subs/delivery/grpc/generated"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/subs/mocks"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"testing"
)

func TestGrpcSubsHandler_GetLicense(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	type fields struct {
		uc subs.SubsUsecase
	}
	type args struct {
		ctx context.Context
		in  *generated.LicenseUUID
	}

	mock := mocks.NewMockSubsUsecase(ctl)
	mock.EXPECT().GetLicense(gomock.Any()).Return(models.License{}, models.Okey)

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *generated.License
		wantErr bool
	}{
		{
			fields: fields{uc: mock},
			args:   args{in: &generated.LicenseUUID{ID: uuid.New().String()}},
		},
		{
			fields: fields{uc: mock},
			args:   args{in: &generated.LicenseUUID{ID: ""}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := GrpcSubsHandler{
				uc: tt.fields.uc,
			}
			_, err := h.GetLicense(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLicense() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestGrpcSubsHandler_SetLicense(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	type fields struct {
		uc subs.SubsUsecase
	}
	type args struct {
		ctx context.Context
		in  *generated.LicenseReq
	}

	mock := mocks.NewMockSubsUsecase(ctl)
	mock.EXPECT().SetLicense(gomock.Any(), gomock.Any()).Return(models.License{}, models.Okey)

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *generated.License
		wantErr bool
	}{
		{
			fields: fields{uc: mock},
			args:   args{in: &generated.LicenseReq{ID: uuid.New().String()}},
		},
		{
			fields: fields{uc: mock},
			args:   args{in: &generated.LicenseReq{ID: ""}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := GrpcSubsHandler{
				uc: tt.fields.uc,
			}
			_, err := h.SetLicense(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("SetLicense() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestNewGrpcSubsHandler(t *testing.T) {
	type args struct {
		uc subs.SubsUsecase
	}
	tests := []struct {
		name string
		args args
		want *GrpcSubsHandler
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGrpcSubsHandler(tt.args.uc); got == nil {
				t.Errorf("NewGrpcSubsHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}
