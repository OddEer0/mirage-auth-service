package model

import "time"

type Role struct {
	Id          string
	Value       string
	Description string
	UpdatedAt   time.Time
	CreatedAt   time.Time
}
