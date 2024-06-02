package grpcv1UserService

import (
	"context"
	errorgrpc "github.com/OddEer0/mirage-auth-service/internal/presentation/errors/error_grpc"
	authv1 "github.com/OddEer0/mirage-auth-service/pkg/gen/auth_v1"
)

func (u *UserServiceServer) DeleteUserById(ctx context.Context, id *authv1.Id) (*authv1.Empty, error) {
	err := u.userRepository.Delete(ctx, id.Id)
	return nil, errorgrpc.Catch(err)
}
