package pg_user_repository_test

import (
	"errors"
	postgresRepository "github.com/OddEer0/mirage-auth-service/internal/infrastructure/storage/postgres_repository"
	"github.com/OddEer0/mirage-auth-service/internal/tests/mock"
	mockgenPostgres "github.com/OddEer0/mirage-auth-service/internal/tests/mockgen/mockgen_postgres"
	testCtx "github.com/OddEer0/mirage-auth-service/internal/tests/test_ctx"
	testLogger "github.com/OddEer0/mirage-auth-service/internal/tests/test_logger"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreate(t *testing.T) {
	mockPgData := mock.PostgresData()
	mockUserData := mockPgData.User

	ctrl := gomock.NewController(t)
	tLog := testLogger.New()
	mockDb := mockgenPostgres.NewMockQuery(ctrl)

	mockDb.EXPECT().QueryRow(gomock.Any(), postgresRepository.CreateUserQuery, mockUserData.CreateUser1Res.Id, mockUserData.CreateUser1Res.Login, mockUserData.CreateUser1Res.Email, mockUserData.CreateUser1Res.Password, mockUserData.CreateUser1Res.Role, mockUserData.CreateUser1Res.IsBanned, mockUserData.CreateUser1Res.BanReason, mockUserData.CreateUser1Res.UpdatedAt, mockUserData.CreateUser1Res.CreatedAt).
		Return(mock.PgMockRow{
			Data: []any{mockUserData.CreateUser1Res.Id, mockUserData.CreateUser1Res.Login, mockUserData.CreateUser1Res.Email, mockUserData.CreateUser1Res.Password, mockUserData.CreateUser1Res.Role, mockUserData.CreateUser1Res.IsBanned, mockUserData.CreateUser1Res.BanReason, mockUserData.CreateUser1Res.UpdatedAt, mockUserData.CreateUser1Res.CreatedAt},
		})
	mockDb.EXPECT().QueryRow(gomock.Any(), postgresRepository.CreateUserQuery, mockUserData.InternalUser.Id, mockUserData.InternalUser.Login, mockUserData.InternalUser.Email, mockUserData.InternalUser.Password, mockUserData.InternalUser.Role, mockUserData.InternalUser.IsBanned, mockUserData.InternalUser.BanReason, mockUserData.InternalUser.UpdatedAt, mockUserData.InternalUser.CreatedAt).
		Return(mock.PgMockRowError{Err: errors.New("internal")})

	userRepo := postgresRepository.NewUserRepository(tLog, mockDb)

	t.Run("Should correct create user", func(t *testing.T) {
		ctx := testCtx.New()
		createUser := mockUserData.CreateUser1Res
		userDb, err := userRepo.Create(ctx, createUser)
		require.NoError(t, err)
		assert.Equal(t, createUser, userDb)
	})

	t.Run("Should internal user", func(t *testing.T) {
		ctx := testCtx.New()
		createUser := mockUserData.InternalUser
		userDb, err := userRepo.Create(ctx, createUser)
		assert.Nil(t, userDb)
		assert.Equal(t, postgresRepository.ErrInternal, err)
		assert.Equal(t,
			[]any{postgresRepository.TraceUserRepoCreate}, tLog.Stack)
		tLog.Clean()
	})
}
