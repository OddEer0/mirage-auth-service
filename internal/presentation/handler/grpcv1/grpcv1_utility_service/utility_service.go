package grpcv1UtilityService

import (
	"context"
	userActivateUsecase "github.com/OddEer0/mirage-auth-service/internal/app/usecases/user_activate_usecase"
	errorgrpc "github.com/OddEer0/mirage-auth-service/internal/presentation/errors/error_grpc"
	authv1 "github.com/OddEer0/mirage-auth-service/pkg/gen/auth_v1"
	"log/slog"
)

type (
	Dependencies struct {
		UserActivateUseCase userActivateUsecase.UseCase
		Log                 *slog.Logger
	}

	UtilityServiceServer struct {
		authv1.UnimplementedUtilityServiceServer
		userActivateUseCase userActivateUsecase.UseCase
		log                 *slog.Logger
	}
)

func (u *UtilityServiceServer) LinkActivate(ctx context.Context, link *authv1.Link) (*authv1.Empty, error) {
	err := u.userActivateUseCase.ActivateLink(ctx, link.Link)
	return &authv1.Empty{}, errorgrpc.Catch(err)
}

func (u *UtilityServiceServer) IsUserActivate(ctx context.Context, id *authv1.Id) (*authv1.Bool, error) {
	isActivate, err := u.userActivateUseCase.IsUserActivate(ctx, id.Id)
	return &authv1.Bool{Value: isActivate}, err
}

func (u *UtilityServiceServer) BanUser(ctx context.Context, user *authv1.BanUser) (*authv1.ResponseUser, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UtilityServiceServer) UnbanUser(ctx context.Context, id *authv1.Id) (*authv1.ResponseUser, error) {
	//TODO implement me
	panic("implement me")
}

func New(dependencies *Dependencies) authv1.UtilityServiceServer {
	return &UtilityServiceServer{log: dependencies.Log, userActivateUseCase: dependencies.UserActivateUseCase}
}
