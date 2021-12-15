package usecase

import (
	"fmt"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/models"
	"github.com/google/uuid"
	"time"
)

type SubsUsecase struct {
	data map[uuid.UUID]models.License
}

func NewSubsUsecase() *SubsUsecase {
	return &SubsUsecase{
		data: map[uuid.UUID]models.License{},
	}
}

func (u *SubsUsecase) GetLicense(id uuid.UUID) (models.License, models.StatusCode) {
	fmt.Print(u.data)
	if l, flag := u.data[id];flag {
		return l, models.Okey
	}
	return models.License{}, models.NotFound

}

func (u *SubsUsecase) SetLicense(id uuid.UUID, license string) (models.License, models.StatusCode) {
	l := models.License{ExpDate: time.Now().AddDate(0, 1, 0), IsValid: true}
	u.data[id] = l
	fmt.Print(u.data)
	return l, models.Okey
}
