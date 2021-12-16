package subs

import (
	"github.com/go-park-mail-ru/2021_2_A06367/internal/models"
	"github.com/google/uuid"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks/interfaces_mock.go -package=mocks

type SubsUsecase interface {
	GetLicense(uuid.UUID) (models.License, models.StatusCode)
	SetLicense(uuid.UUID, string) (models.License, models.StatusCode)
}

type SubsRepository interface {
	GetLicense(uuid.UUID) (models.License, models.StatusCode)
	SetLicense(uuid.UUID, models.License) (models.License, models.StatusCode)
}
