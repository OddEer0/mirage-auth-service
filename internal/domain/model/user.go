package model

import "time"

type (
	CreateUser struct {
		Id       string
		Login    string
		Email    string
		Password string
		Role     string
	}

	User struct {
		Id        string
		Login     string
		Email     string
		Password  string
		Role      string
		IsBanned  bool
		BanReason *string
		CreatedAt time.Time
		UpdatedAt time.Time
	}

	UserAggregate struct {
		User
		*JwtToken
		Activate *UserActivate
	}
)
