package grpcMapper

import (
	appDto "github.com/OddEer0/mirage-auth-service/internal/app/app_dto"
	"github.com/OddEer0/mirage-auth-service/internal/domain/model"
	authv1 "github.com/OddEer0/mirage-src/protogen/mirage_auth_service"
)

type UserMapper struct {
}

func (u *UserMapper) PureUserToResponseUserV1(pureUser *appDto.PureUser) *authv1.ResponseUser {
	banReason := ""
	if pureUser.BanReason != nil {
		banReason = *pureUser.BanReason
	}
	return &authv1.ResponseUser{
		Id: pureUser.Id, Login: pureUser.Login, Email: pureUser.Email, Role: pureUser.Role, IsBanned: pureUser.IsBanned, BanReason: banReason,
	}
}

func (u *UserMapper) ModelUserToResponseUserV1(user *model.User) *authv1.ResponseUser {
	banReason := ""
	if user.BanReason != nil {
		banReason = *user.BanReason
	}
	return &authv1.ResponseUser{
		Id: user.Id, Login: user.Login, Email: user.Email, Role: user.Role, IsBanned: user.IsBanned, BanReason: banReason,
	}
}
