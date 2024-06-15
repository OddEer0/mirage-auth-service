package pg_repository_test

import (
	postgresRepository "github.com/OddEer0/mirage-auth-service/internal/infrastructure/storage/postgres_repository"
	"github.com/OddEer0/mirage-auth-service/internal/tests/mock"
	testCtx "github.com/OddEer0/mirage-auth-service/internal/tests/test_ctx"
	testLogger "github.com/OddEer0/mirage-auth-service/internal/tests/test_logger"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUserPgRepository(t *testing.T) {
	mockPgData := mock.PostgresData()
	mockUserData := mockPgData.User
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDb := mock.PgxConnMock(ctrl, mockPgData)

	tLog := testLogger.New()
	userRepo := postgresRepository.NewUserRepository(tLog, mockDb)

	t.Run("Testing GetById", func(t *testing.T) {
		t.Run("Should correct work", func(t *testing.T) {
			ctx := testCtx.New()
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

		t.Run("Should not found", func(t *testing.T) {
			ctx := testCtx.New()
			user := mockUserData.NotFoundUser
			userDb, err := userRepo.GetById(ctx, user.Id)
			assert.Nil(t, userDb)
			assert.Equal(t, err, postgresRepository.ErrUserNotFound)
			assert.NotEmpty(t, tLog.Message)
			assert.Equal(t, []any{postgresRepository.TraceUserRepoGetById}, tLog.Stack)
			tLog.Clean()
		})

		t.Run("Should internal", func(t *testing.T) {
			ctx := testCtx.New()
			user := mockUserData.InternalUser
			userDb, err := userRepo.GetById(ctx, user.Id)
			assert.Nil(t, userDb)
			assert.Equal(t, err, postgresRepository.ErrInternal)
			assert.NotEmpty(t, tLog.Message)
			assert.Equal(t, []any{postgresRepository.TraceUserRepoGetById}, tLog.Stack)
			tLog.Clean()
		})
	})

	t.Run("Testing Create", func(t *testing.T) {
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
	})
}
