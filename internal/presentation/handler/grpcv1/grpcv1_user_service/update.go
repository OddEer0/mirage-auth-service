package grpcv1UserService

import (
	"context"
	errorgrpc "github.com/OddEer0/mirage-auth-service/internal/presentation/errors/error_grpc"
	authv1 "github.com/OddEer0/mirage-src/protogen/mirage_auth_service"
)

func (u *UserServiceServer) UpdateUserRole(ctx context.Context, role *authv1.UpdateUserRole) (*authv1.ResponseUser, error) {
	pureUser, err := u.userUseCase.UpdateUserRole(ctx, role.Id, role.Role)
	if err != nil {
		return nil, errorgrpc.Catch(err)
	}
	banReason := ""
	if pureUser.BanReason != nil {
		banReason = *pureUser.BanReason
	}
	return &authv1.ResponseUser{
		Id: pureUser.Id, Login: pureUser.Login, Email: pureUser.Email, Role: pureUser.Role, IsBanned: pureUser.IsBanned, BanReason: banReason,
	}, nil
}
