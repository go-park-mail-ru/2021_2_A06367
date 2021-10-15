package models

import (
	"github.com/google/uuid"
)

type Actors struct {
	Id					uuid.UUID 	`json:"-"`
	Name				string    	`json:"name"`
	Surname				string    	`json:"surname"`
	Height				float32		`json:"height"`
	DateOfBirth			string    	`json:"password"`
	Genres				[]string	`json:"genres"`
}