package model

import "time"

type JwtToken struct {
	UserId    string
	Value     string
	UpdatedAt time.Time
	CreatedAt time.Time
}
