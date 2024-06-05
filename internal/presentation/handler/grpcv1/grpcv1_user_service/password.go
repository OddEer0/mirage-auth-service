package grpcv1UserService

import (
	"context"
	authv1 "github.com/OddEer0/mirage-src/protogen/mirage_auth_service"
)

func (u *UserServiceServer) ChangePassword(ctx context.Context, request *authv1.ChangePasswordRequest) (*authv1.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserServiceServer) ConfirmChangePassword(ctx context.Context, link *authv1.Link) (*authv1.Empty, error) {
	//TODO implement me
	panic("implement me")
}
