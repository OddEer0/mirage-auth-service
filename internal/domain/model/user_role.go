package model

import "time"

type UserRole struct {
	Id        string
	UserId    string
	RoleId    string
	UpdatedAt time.Time
	CreatedAt time.Time
}
