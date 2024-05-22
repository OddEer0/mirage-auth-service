package grpcv1

import (
	"context"
	domainQuery "github.com/OddEer0/mirage-auth-service/internal/domain/repository/domain_query"
	authv1 "github.com/OddEer0/mirage-auth-service/pkg/gen/auth_v1"
)

func (a *AuthServiceServer) GetUsersByQuery(ctx context.Context, query *authv1.PaginationQuery) (*authv1.ManyResponseUser, error) {
	users, pageCount, err := a.userRepo.GetByQuery(ctx, &domainQuery.UserQueryRequest{
		PaginationQuery: domainQuery.PaginationQuery{
			PageCount: uint(query.Count), CurrentPage: uint(query.CurrentPage),
		},
		OrderQuery: domainQuery.OrderQuery{
			OrderBy:        query.OrderBy,
			OrderDirection: query.OrderDirection,
		},
	})
	if err != nil {
		return nil, err
	}

	resUsers := make([]*authv1.ResponseUser, 0, query.Count)

	for _, user := range users {
		resUsers = append(resUsers, &authv1.ResponseUser{
			Id:        user.Id,
			Login:     user.Login,
			Email:     user.Email,
			Role:      user.Role,
			IsBanned:  user.IsBanned,
			BanReason: "",
		})
	}

	return &authv1.ManyResponseUser{
		Users: resUsers, PageCount: uint32(pageCount),
	}, nil
}
