package subs

import (
	"github.com/go-park-mail-ru/2021_2_A06367/internal/models"
	"github.com/google/uuid"
)

type SubsUsecase interface {
	GetLicense(id uuid.UUID) (models.License, models.StatusCode)
	SetLicense(id uuid.UUID, license string) (models.License, models.StatusCode)
}

