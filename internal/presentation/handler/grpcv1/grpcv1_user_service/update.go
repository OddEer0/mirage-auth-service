package grpcv1UserService

import (
	"context"
	errorgrpc "github.com/OddEer0/mirage-auth-service/internal/presentation/errors/error_grpc"
	grpcMapper "github.com/OddEer0/mirage-auth-service/internal/presentation/mapper/grpc_mapper"
	authv1 "github.com/OddEer0/mirage-src/protogen/mirage_auth_service"
)

func (u *UserServiceServer) UpdateUserRole(ctx context.Context, role *authv1.UpdateUserRole) (*authv1.ResponseUser, error) {
	pureUser, err := u.userUseCase.UpdateUserRole(ctx, role.Id, role.Role)
	if err != nil {
		return nil, errorgrpc.Catch(err)
	}
	mapper := grpcMapper.UserMapper{}
	return mapper.PureUserToResponseUserV1(pureUser), nil
}
