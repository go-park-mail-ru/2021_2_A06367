package models

import "time"

type License struct {
	IsValid bool
	ExpDate time.Time
}
