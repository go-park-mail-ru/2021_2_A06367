package models

import (
	"github.com/google/uuid"
	"time"
)

type Actors struct {
	Id					uuid.UUID 	`json:"id"`
	Name				string    	`json:"name"`
	Surname				string    	`json:"surname"`
	Avatar				string		`json:"avatar"`
	Height				float32		`json:"height"`
	DateOfBirth			time.Time	`json:"date_of_birth"`
	Description	string		`json:"description"`
	Genres				[]string	`json:"genres"`
}