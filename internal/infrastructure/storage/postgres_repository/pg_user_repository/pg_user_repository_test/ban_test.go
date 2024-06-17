package pg_user_repository_test

import (
	"database/sql"
	"errors"
	pgUserRepository "github.com/OddEer0/mirage-auth-service/internal/infrastructure/storage/postgres_repository/pg_user_repository"
	"github.com/OddEer0/mirage-auth-service/internal/tests/mock"
	mockgenPostgres "github.com/OddEer0/mirage-auth-service/internal/tests/mockgen/mockgen_postgres"
	testCtx "github.com/OddEer0/mirage-auth-service/internal/tests/test_ctx"
	testLogger "github.com/OddEer0/mirage-auth-service/internal/tests/test_logger"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBan(t *testing.T) {
	mockPgData := mock.PostgresData()
	mockUserData := mockPgData.User

	t.Run("Testing BanUserById", func(t *testing.T) {
		tLog := testLogger.New()
		ctrl := gomock.NewController(t)
		mockDb := mockgenPostgres.NewMockQuery(ctrl)

		banReason := "toxic"
		updatedUser := mockUserData.CorrectUser1
		mockDb.EXPECT().QueryRow(gomock.Any(), pgUserRepository.UpdateUserBanById, mockUserData.CorrectUser1.Id, true, banReason).Return(mock.PgMockRow{
			Data: []any{
				updatedUser.Id,
				updatedUser.Login,
				updatedUser.Email,
				updatedUser.Password,
				updatedUser.Role,
				true,
				&banReason,
				updatedUser.UpdatedAt,
				updatedUser.CreatedAt,
			},
		})
		mockDb.EXPECT().QueryRow(gomock.Any(), pgUserRepository.UpdateUserBanById, mockUserData.NotFoundUser.Id, true, banReason).Return(mock.PgMockRowError{
			Err: sql.ErrNoRows,
		})
		mockDb.EXPECT().QueryRow(gomock.Any(), pgUserRepository.UpdateUserBanById, mockUserData.InternalUser.Id, true, banReason).Return(mock.PgMockRowError{
			Err: errors.New("internal"),
		})

		userRepo := pgUserRepository.New(tLog, mockDb)

		t.Run("Should correct work", func(t *testing.T) {
			ctx := testCtx.New()
			user := *mockUserData.CorrectUser1
			userDb, err := userRepo.BanUserById(ctx, user.Id, banReason)
			user.IsBanned = true
			user.BanReason = &banReason
			require.NoError(t, err)
			assert.Equal(t, &user, userDb)
		})

		t.Run("Should not found", func(t *testing.T) {
			ctx := testCtx.New()
			user := mockUserData.NotFoundUser
			userDb, err := userRepo.BanUserById(ctx, user.Id, banReason)
			assert.Nil(t, userDb)
			assert.Equal(t, pgUserRepository.ErrUserNotFound, err)
			assert.NotEmpty(t, tLog.Message)
			assert.NotEmpty(t, []any{pgUserRepository.TraceBanUserById}, tLog.Stack)
			tLog.Clean()
		})

		t.Run("Should internal", func(t *testing.T) {
			ctx := testCtx.New()
			user := mockUserData.InternalUser
			userDb, err := userRepo.BanUserById(ctx, user.Id, banReason)
			assert.Nil(t, userDb)
			assert.Equal(t, pgUserRepository.ErrInternal, err)
			assert.NotEmpty(t, tLog.Message)
			assert.NotEmpty(t, []any{pgUserRepository.TraceBanUserById}, tLog.Stack)
			tLog.Clean()
		})
	})

	t.Run("Testing UnbanUserById", func(t *testing.T) {
		tLog := testLogger.New()
		ctrl := gomock.NewController(t)
		mockDb := mockgenPostgres.NewMockQuery(ctrl)

		updatedUser := mockUserData.BannedUser1
		var nullableStr *string
		mockDb.EXPECT().QueryRow(gomock.Any(), pgUserRepository.UpdateUserBanById, mockUserData.BannedUser1.Id, false, nil).Return(mock.PgMockRow{
			Data: []any{
				updatedUser.Id,
				updatedUser.Login,
				updatedUser.Email,
				updatedUser.Password,
				updatedUser.Role,
				false,
				nullableStr,
				updatedUser.UpdatedAt,
				updatedUser.CreatedAt,
			},
		})
		mockDb.EXPECT().QueryRow(gomock.Any(), pgUserRepository.UpdateUserBanById, mockUserData.NotFoundUser.Id, false, nil).Return(mock.PgMockRowError{
			Err: sql.ErrNoRows,
		})
		mockDb.EXPECT().QueryRow(gomock.Any(), pgUserRepository.UpdateUserBanById, mockUserData.InternalUser.Id, false, nil).Return(mock.PgMockRowError{
			Err: errors.New("internal"),
		})

		userRepo := pgUserRepository.New(tLog, mockDb)

		t.Run("Should correct work", func(t *testing.T) {
			ctx := testCtx.New()
			user := *mockUserData.BannedUser1
			userDb, err := userRepo.UnbanUserById(ctx, user.Id)
			t.Log(userDb)
			require.NoError(t, err)
			user.IsBanned = false
			user.BanReason = nil
			assert.Equal(t, &user, userDb)
		})

		t.Run("Should not found", func(t *testing.T) {
			ctx := testCtx.New()
			user := *mockUserData.NotFoundUser
			userDb, err := userRepo.UnbanUserById(ctx, user.Id)
			assert.Nil(t, userDb)
			assert.Equal(t, pgUserRepository.ErrUserNotFound, err)
			assert.NotEmpty(t, tLog.Message)
			assert.Equal(t, []any{pgUserRepository.TraceUnbanUserById}, tLog.Stack)
			tLog.Clean()
		})

		t.Run("Should internal", func(t *testing.T) {
			ctx := testCtx.New()
			user := *mockUserData.InternalUser
			userDb, err := userRepo.UnbanUserById(ctx, user.Id)
			assert.Nil(t, userDb)
			assert.Equal(t, pgUserRepository.ErrInternal, err)
			assert.NotEmpty(t, tLog.Message)
			assert.Equal(t, []any{pgUserRepository.TraceUnbanUserById}, tLog.Stack)
			tLog.Clean()
		})
	})
}
