package usecase

import (
	"github.com/go-park-mail-ru/2021_2_A06367/internal/models"
	"github.com/google/uuid"
	"time"
)

type SubsUsecase struct {
}

func NewSubsUsecase() *SubsUsecase {
	return &SubsUsecase{}
}

func (u SubsUsecase) GetLicense(id uuid.UUID) (models.License, models.StatusCode) {
	return models.License{ExpDate: time.Now().AddDate(0, 1, 0), IsValid: true}, models.Okey

}

func (u SubsUsecase) SetLicense(id uuid.UUID, license string) (models.License, models.StatusCode) {
	return models.License{ExpDate: time.Now().AddDate(0, 1, 0), IsValid: true}, models.Okey
}
