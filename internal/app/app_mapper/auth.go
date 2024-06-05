package appMapper

import (
	appDto "github.com/OddEer0/mirage-auth-service/internal/app/app_dto"
	"github.com/OddEer0/mirage-auth-service/internal/domain/model"
)

type UserMapper struct {
}

func (u *UserMapper) ToPureUser(user *model.User) *appDto.PureUser {
	return &appDto.PureUser{
		Id:        user.Id,
		Login:     user.Login,
		Email:     user.Email,
		IsBanned:  user.IsBanned,
		BanReason: user.BanReason,
		Role:      user.Role,
	}
}

func (u *UserMapper) ToPureUserCopy(user *model.User) appDto.PureUser {
	return appDto.PureUser{
		Id:        user.Id,
		Login:     user.Login,
		Email:     user.Email,
		IsBanned:  user.IsBanned,
		BanReason: user.BanReason,
		Role:      user.Role,
	}
}
