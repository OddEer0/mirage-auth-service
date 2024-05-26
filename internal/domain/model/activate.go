package model

import "time"

type UserActivate struct {
	UserId     string
	IsActivate bool
	Link       string
	UpdatedAt  time.Time
	CreatedAt  time.Time
}
