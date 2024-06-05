package grpcv1UserService

import (
	"context"
	authv1 "github.com/OddEer0/mirage-src/protogen/mirage_auth_service"
)

func (u *UserServiceServer) CheckRole(ctx context.Context, request *authv1.CheckRoleRequest) (*authv1.Bool, error) {
	hasReqRole, err := u.userUseCase.CheckUserRole(ctx, request.UserId, request.Role)
	if err != nil {
		return nil, err
	}
	return &authv1.Bool{
		Value: hasReqRole,
	}, nil
}
