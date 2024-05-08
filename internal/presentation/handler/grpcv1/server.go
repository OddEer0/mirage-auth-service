package grpcv1

import (
	"context"
	"github.com/OddEer0/mirage-auth-service/internal/domain/repository"
	postgresRepository "github.com/OddEer0/mirage-auth-service/internal/infrastructure/storage/postgres_repository"
	authv1 "github.com/OddEer0/mirage-auth-service/pkg/gen/auth_v1"
	"github.com/jackc/pgx/v5"
)

type AuthServiceServer struct {
	authv1.UnimplementedAuthServiceServer
	userRepo repository.UserRepository
}

func (a *AuthServiceServer) ChangePassword(ctx context.Context, request *authv1.ChangePasswordRequest) (*authv1.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (a *AuthServiceServer) CheckRole(ctx context.Context, request *authv1.CheckRoleRequest) (*authv1.ResponseUser, error) {
	//TODO implement me
	panic("implement me")
}

func (a *AuthServiceServer) GetUserById(ctx context.Context, id *authv1.Id) (*authv1.ResponseUser, error) {
	//TODO implement me
	panic("implement me")
}

func (a *AuthServiceServer) UpdateUserEmail(ctx context.Context, email *authv1.UpdateUserEmail) (*authv1.ResponseUser, error) {
	//TODO implement me
	panic("implement me")
}

func (a *AuthServiceServer) UpdateUserRole(ctx context.Context, role *authv1.UpdateUserRole) (*authv1.ResponseUser, error) {
	//TODO implement me
	panic("implement me")
}

func (a *AuthServiceServer) DeleteUserById(ctx context.Context, id *authv1.Id) (*authv1.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (a *AuthServiceServer) BanUser(ctx context.Context, user *authv1.BanUser) (*authv1.ResponseUser, error) {
	//TODO implement me
	panic("implement me")
}

func (a *AuthServiceServer) UnbanUser(ctx context.Context, id *authv1.Id) (*authv1.ResponseUser, error) {
	//TODO implement me
	panic("implement me")
}

func New(db *pgx.Conn) authv1.AuthServiceServer {
	userRepository := postgresRepository.NewUserRepository(db)
	return &AuthServiceServer{userRepo: userRepository}
}
