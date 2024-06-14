package pg_repository_test

import (
	"context"
	"github.com/OddEer0/mirage-auth-service/internal/infrastructure/logger"
	postgresRepository "github.com/OddEer0/mirage-auth-service/internal/infrastructure/storage/postgres_repository"
	"github.com/OddEer0/mirage-auth-service/internal/tests/mock"
	mockgenPostgres "github.com/OddEer0/mirage-auth-service/internal/tests/mockgen/mockgen_postgres"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUserPgRepository(t *testing.T) {
	mockUserData := mock.PostgresData().User
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDb := mockgenPostgres.NewMockQuery(ctrl)
	mockDb.EXPECT().QueryRow(gomock.Any(), postgresRepository.GetUserByIdQuery, mockUserData.CorrectUser1.Id).Return(mock.PgMockRow{
		Data: []any{mockUserData.CorrectUser1.Id, mockUserData.CorrectUser1.Login, mockUserData.CorrectUser1.Email, mockUserData.CorrectUser1.Password, mockUserData.CorrectUser1.Role, mockUserData.CorrectUser1.IsBanned, mockUserData.CorrectUser1.BanReason, mockUserData.CorrectUser1.UpdatedAt, mockUserData.CorrectUser1.CreatedAt}},
	)
	mockDb.EXPECT().QueryRow(gomock.Any(), postgresRepository.GetUserByIdQuery, mockUserData.AdminUser1.Id).Return(mock.PgMockRow{
		Data: []any{mockUserData.AdminUser1.Id, mockUserData.AdminUser1.Login, mockUserData.AdminUser1.Email, mockUserData.AdminUser1.Password, mockUserData.AdminUser1.Role, mockUserData.AdminUser1.IsBanned, mockUserData.AdminUser1.BanReason, mockUserData.AdminUser1.UpdatedAt, mockUserData.AdminUser1.CreatedAt}},
	)
	mockDb.EXPECT().QueryRow(gomock.Any(), postgresRepository.GetUserByIdQuery, mockUserData.BannedUser1.Id).Return(mock.PgMockRow{
		Data: []any{mockUserData.BannedUser1.Id, mockUserData.BannedUser1.Login, mockUserData.BannedUser1.Email, mockUserData.BannedUser1.Password, mockUserData.BannedUser1.Role, mockUserData.BannedUser1.IsBanned, mockUserData.BannedUser1.BanReason, mockUserData.BannedUser1.UpdatedAt, mockUserData.BannedUser1.CreatedAt}},
	)
	mockDb.EXPECT().QueryRow(gomock.Any(), postgresRepository.CreateUserQuery, mockUserData.CreateUser1Res.Id, mockUserData.CreateUser1Res.Login, mockUserData.CreateUser1Res.Email, mockUserData.CreateUser1Res.Password, mockUserData.CreateUser1Res.Role, mockUserData.CreateUser1Res.IsBanned, mockUserData.CreateUser1Res.BanReason, mockUserData.CreateUser1Res.UpdatedAt, mockUserData.CreateUser1Res.CreatedAt).
		Return(mock.PgMockRow{
			Data: []any{mockUserData.CreateUser1Res.Id, mockUserData.CreateUser1Res.Login, mockUserData.CreateUser1Res.Email, mockUserData.CreateUser1Res.Password, mockUserData.CreateUser1Res.Role, mockUserData.CreateUser1Res.IsBanned, mockUserData.CreateUser1Res.BanReason, mockUserData.CreateUser1Res.UpdatedAt, mockUserData.CreateUser1Res.CreatedAt},
		})

	tLog := logger.MustLoad("test", "")
	userRepo := postgresRepository.NewUserRepository(tLog, mockDb)

	t.Run("Should correct GetById", func(t *testing.T) {
		user := mockUserData.CorrectUser1
		userDb, err := userRepo.GetById(ctx, user.Id)
		require.NoError(t, err)
		assert.Equal(t, user, userDb)

		user = mockUserData.BannedUser1
		userDb, err = userRepo.GetById(ctx, user.Id)
		require.NoError(t, err)
		assert.Equal(t, user, userDb)

		user = mockUserData.AdminUser1
		userDb, err = userRepo.GetById(ctx, user.Id)
		require.NoError(t, err)
		assert.Equal(t, user, userDb)
	})

	t.Run("Should correct Create", func(t *testing.T) {
		createUser := mockUserData.CreateUser1Res
		userDb, err := userRepo.Create(ctx, createUser)
		require.NoError(t, err)
		assert.Equal(t, createUser, userDb)
	})

	t.Run("Should correct Delete", func(t *testing.T) {

	})
}
