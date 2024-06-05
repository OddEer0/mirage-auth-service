package grpcv1UserService

import (
	"context"
	domainQuery "github.com/OddEer0/mirage-auth-service/internal/domain/repository/domain_query"
	errorgrpc "github.com/OddEer0/mirage-auth-service/internal/presentation/errors/error_grpc"
	authv1 "github.com/OddEer0/mirage-src/protogen/mirage_auth_service"
)

func (u *UserServiceServer) GetUserById(ctx context.Context, id *authv1.Id) (*authv1.ResponseUser, error) {
	user, err := u.userUseCase.GetById(ctx, id.Id)
	if err != nil {
		return nil, errorgrpc.Catch(err)
	}

	banReason := ""
	if user.BanReason != nil {
		banReason = *user.BanReason
	}
	return &authv1.ResponseUser{
		Id:        user.Id,
		Login:     user.Login,
		Email:     user.Email,
		Role:      user.Role,
		IsBanned:  user.IsBanned,
		BanReason: banReason,
	}, nil
}

func (u *UserServiceServer) GetUsersByQuery(ctx context.Context, query *authv1.PaginationQuery) (*authv1.ManyResponseUser, error) {
	users, pageCount, err := u.userRepository.GetByQuery(ctx, &domainQuery.UserQueryRequest{})
	if err != nil {
		return nil, errorgrpc.Catch(err)
	}

	responseUsers := make([]*authv1.ResponseUser, 0, len(users))
	for _, user := range users {
		banReason := ""
		if user.BanReason != nil {
			banReason = *user.BanReason
		}
		responseUsers = append(responseUsers, &authv1.ResponseUser{
			Id:        user.Id,
			Login:     user.Login,
			Email:     user.Email,
			Role:      user.Role,
			IsBanned:  user.IsBanned,
			BanReason: banReason,
		})
	}

	return &authv1.ManyResponseUser{
		Users:     responseUsers,
		PageCount: uint32(pageCount),
	}, err
}
