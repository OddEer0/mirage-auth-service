package grpcv1UserService

import (
	"context"
	authv1 "github.com/OddEer0/mirage-auth-service/pkg/gen/auth_v1"
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
