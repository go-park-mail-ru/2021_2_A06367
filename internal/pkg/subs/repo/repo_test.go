package repo

import (
	"errors"
	"github.com/driftprogramming/pgxpoolmock"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/models"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/jackc/pgtype/pgxtype"
	"reflect"
	"testing"
	"time"
)

func TestNewSubsRepo(t *testing.T) {

	ctl := gomock.NewController(t)
	mockPool := pgxpoolmock.NewMockPgxPool(ctl)

	type args struct {
		pool pgxtype.Querier
	}
	tests := []struct {
		name string
		args args
		want *SubsRepo
	}{
		{name: "",
			args: args{pool: mockPool},
			want: &SubsRepo{pool: mockPool}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSubsRepo(tt.args.pool); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSubsRepo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubsRepo_GetLicense(t *testing.T) {

	date := time.Now()
	columns := []string{"exp_date"}
	pgxRows := pgxpoolmock.NewRows(columns).
		AddRow(date).AddRow(date).RowError(1, errors.New("")).ToPgxRows()
	pgxRows.Next()

	pgxRowsErr := pgxpoolmock.NewRows(columns).
		AddRow(date).AddRow(date).RowError(0, errors.New("")).ToPgxRows()
	pgxRowsErr.Next()

	ctl := gomock.NewController(t)
	mockPool := pgxpoolmock.NewMockPgxPool(ctl)
	mockPool.EXPECT().QueryRow(gomock.Any(), gomock.Any(), gomock.Any()).Return(pgxRows)
	mockPool.EXPECT().QueryRow(gomock.Any(), gomock.Any(), gomock.Any()).Return(pgxRowsErr)

	type fields struct {
		pool pgxtype.Querier
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
			fields: fields{pool: mockPool},
			args: args{id: uuid.New()},
			want: models.License{IsValid: true, ExpDate: date},
			want1: models.Okey,
		},
		{
			name:   "",
			fields: fields{pool: mockPool},
			args: args{id: uuid.New()},
			want: models.License{IsValid: false, ExpDate: time.Time{}},
			want1: models.NotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &SubsRepo{
				pool: tt.fields.pool,
			}
			got, got1 := r.GetLicense(tt.args.id)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLicense() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetLicense() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestSubsRepo_SetLicense(t *testing.T) {
	date := time.Now()
	columns := []string{"user_id"}
	pgxRows := pgxpoolmock.NewRows(columns).
		AddRow(uuid.New()).AddRow(uuid.New()).RowError(1, errors.New("")).ToPgxRows()
	pgxRows.Next()

	pgxRowsErr := pgxpoolmock.NewRows(columns).
		AddRow(uuid.New()).AddRow(uuid.New()).RowError(0, errors.New("")).ToPgxRows()
	pgxRowsErr.Next()

	ctl := gomock.NewController(t)
	mockPool := pgxpoolmock.NewMockPgxPool(ctl)
	mockPool.EXPECT().QueryRow(gomock.Any(), gomock.Any(), gomock.Any()).Return(pgxRows)
	mockPool.EXPECT().QueryRow(gomock.Any(), gomock.Any(), gomock.Any()).Return(pgxRowsErr)

	type fields struct {
		pool pgxtype.Querier
	}
	type args struct {
		id      uuid.UUID
		license models.License
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
			fields: fields{pool: mockPool},
			args: args{id: uuid.New(), license: models.License{ExpDate:date, IsValid: true}},
			want: models.License{IsValid: true, ExpDate: date},
			want1: models.Okey,
		},
		{
			name:   "",
			fields: fields{pool: mockPool},
			args: args{id: uuid.New(), license: models.License{ExpDate:date, IsValid: true}},
			want: models.License{},
			want1: models.BadRequest,
		},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &SubsRepo{
				pool: tt.fields.pool,
			}
			got, got1 := r.SetLicense(tt.args.id, tt.args.license)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetLicense() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("SetLicense() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
