package grpcv1UserService

import (
	"context"
	domainQuery "github.com/OddEer0/mirage-auth-service/internal/domain/repository/domain_query"
	errorgrpc "github.com/OddEer0/mirage-auth-service/internal/presentation/errors/error_grpc"
	grpcMapper "github.com/OddEer0/mirage-auth-service/internal/presentation/mapper/grpc_mapper"
	authv1 "github.com/OddEer0/mirage-src/protogen/mirage_auth_service"
)

func (u *UserServiceServer) GetUserById(ctx context.Context, id *authv1.Id) (*authv1.ResponseUser, error) {
	user, err := u.userUseCase.GetById(ctx, id.Id)
	if err != nil {
		return nil, errorgrpc.Catch(err)
	}

	mapper := grpcMapper.UserMapper{}
	return mapper.PureUserToResponseUserV1(user), nil
}

func (u *UserServiceServer) GetUsersByQuery(ctx context.Context, query *authv1.PaginationQuery) (*authv1.ManyResponseUser, error) {
	users, pageCount, err := u.userRepository.GetByQuery(ctx, &domainQuery.UserQueryRequest{})
	if err != nil {
		return nil, errorgrpc.Catch(err)
	}

	responseUsers := make([]*authv1.ResponseUser, 0, len(users))
	mapper := grpcMapper.UserMapper{}
	for _, user := range users {
		responseUsers = append(responseUsers, mapper.ModelUserToResponseUserV1(user))
	}

	return &authv1.ManyResponseUser{
		Users:     responseUsers,
		PageCount: uint32(pageCount),
	}, err
}
