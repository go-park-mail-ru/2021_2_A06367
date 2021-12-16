package usecase

import (
	"github.com/go-park-mail-ru/2021_2_A06367/internal/models"
	"github.com/go-park-mail-ru/2021_2_A06367/internal/pkg/subs"
	"github.com/google/uuid"
	"time"
)

type SubsUsecase struct {
	r subs.SubsRepository
}

func NewSubsUsecase(r subs.SubsRepository) *SubsUsecase {
	return &SubsUsecase{r: r}
}

func (u *SubsUsecase) GetLicense(id uuid.UUID) (models.License, models.StatusCode) {
	return u.r.GetLicense(id)

}

func (u *SubsUsecase) SetLicense(id uuid.UUID, license string) (models.License, models.StatusCode) {

	l := models.License{ExpDate: time.Now().AddDate(0, 1, 0), IsValid: true}
	return u.r.SetLicense(id, l)
}
