package usecase

import (
	"github.com/go-park-mail-ru/2021_2_A06367/internal/models"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/subs"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/subs/mocks"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"reflect"
	"testing"
)

func TestNewSubsUsecase(t *testing.T) {
	ctl := gomock.NewController(t)
	r := mocks.NewMockSubsRepository(ctl)

	type args struct {
		r subs.SubsRepository
	}
	tests := []struct {
		name string
		args args
		want *SubsUsecase
	}{
		{name: "",
			args: args{r},
			want: &SubsUsecase{r: r}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSubsUsecase(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSubsUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubsUsecase_GetLicense(t *testing.T) {
	ctl := gomock.NewController(t)
	r := mocks.NewMockSubsRepository(ctl)

	r.EXPECT().GetLicense(gomock.Any()).Return(models.License{}, models.Okey)

	type fields struct {
		r subs.SubsRepository
	}
	type args struct {
		id uuid.UUID
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   models.License
		want1  models.StatusCode
	}{
		{
			name:   "",
			fields: fields{r: r},
			args:   args{id: uuid.New()},
			want:   models.License{},
			want1:  models.Okey,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &SubsUsecase{
				r: tt.fields.r,
			}
			got, got1 := u.GetLicense(tt.args.id)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLicense() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetLicense() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestSubsUsecase_SetLicense(t *testing.T) {
	ctl := gomock.NewController(t)
	r := mocks.NewMockSubsRepository(ctl)

	r.EXPECT().SetLicense(gomock.Any(), gomock.Any()).Return(models.License{}, models.Okey)
	type fields struct {
		r    subs.SubsRepository
	}
	type args struct {
		id      uuid.UUID
		license string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   models.License
		want1  models.StatusCode
	}{
		{name: "",
			fields: fields{r: r},
			args:   args{id: uuid.New(), license: ""},
			want1:  models.Okey,
			want:   models.License{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &SubsUsecase{
				r: tt.fields.r,
			}
			got, got1 := u.SetLicense(tt.args.id, tt.args.license)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetLicense() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("SetLicense() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
