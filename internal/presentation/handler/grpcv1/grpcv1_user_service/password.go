package grpcv1UserService

import (
	"context"
	authv1 "github.com/OddEer0/mirage-auth-service/pkg/gen/auth_v1"
)

func (u *UserServiceServer) ChangePassword(ctx context.Context, request *authv1.ChangePasswordRequest) (*authv1.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserServiceServer) ConfirmChangePassword(ctx context.Context, link *authv1.Link) (*authv1.Empty, error) {
	//TODO implement me
	panic("implement me")
}
