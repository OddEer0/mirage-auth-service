package model

import "time"

type User struct {
	Id        string
	Login     string
	Email     string
	Password  string
	UpdatedAt time.Time
	CreatedAt time.Time
}
