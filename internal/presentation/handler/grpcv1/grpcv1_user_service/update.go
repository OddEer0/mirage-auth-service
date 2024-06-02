package grpcv1UserService

import (
	"context"
	errorgrpc "github.com/OddEer0/mirage-auth-service/internal/presentation/errors/error_grpc"
	authv1 "github.com/OddEer0/mirage-auth-service/pkg/gen/auth_v1"
)

func (u *UserServiceServer) UpdateUserEmail(ctx context.Context, email *authv1.UpdateUserEmail) (*authv1.ResponseUser, error) {
	pureUser, err := u.userUseCase.UpdateUserEmail(ctx, email.Id, email.Email)
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
