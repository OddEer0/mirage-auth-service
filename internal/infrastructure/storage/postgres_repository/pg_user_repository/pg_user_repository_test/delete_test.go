package pg_user_repository_test

import (
	"errors"
	pgUserRepository "github.com/OddEer0/mirage-auth-service/internal/infrastructure/storage/postgres_repository/pg_user_repository"
	"github.com/OddEer0/mirage-auth-service/internal/tests/mock"
	mockgenPostgres "github.com/OddEer0/mirage-auth-service/internal/tests/mockgen/mockgen_postgres"
	testCtx "github.com/OddEer0/mirage-auth-service/internal/tests/test_ctx"
	testLogger "github.com/OddEer0/mirage-auth-service/internal/tests/test_logger"
	"github.com/golang/mock/gomock"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDelete(t *testing.T) {
	mockPgData := mock.PostgresData()
	mockUserData := mockPgData.User

	t.Run("Testing Delete", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		tLog := testLogger.New()
		mockDb := mockgenPostgres.NewMockQuery(ctrl)

		mockDb.EXPECT().QueryRow(gomock.Any(), pgUserRepository.HasUserByIdQuery, mockUserData.CorrectUser1.Id).Return(mock.PgMockRow{Data: []any{true}})
		mockDb.EXPECT().Exec(gomock.Any(), pgUserRepository.DeleteUserByIdQuery, mockUserData.CorrectUser1.Id).Return(pgconn.CommandTag{}, nil)
		mockDb.EXPECT().QueryRow(gomock.Any(), pgUserRepository.HasUserByIdQuery, mockUserData.AdminUser1.Id).Return(mock.PgMockRow{Data: []any{true}})
		mockDb.EXPECT().Exec(gomock.Any(), pgUserRepository.DeleteUserByIdQuery, mockUserData.AdminUser1.Id).Return(pgconn.CommandTag{}, nil)
		mockDb.EXPECT().QueryRow(gomock.Any(), pgUserRepository.HasUserByIdQuery, mockUserData.BannedUser1.Id).Return(mock.PgMockRow{Data: []any{true}})
		mockDb.EXPECT().Exec(gomock.Any(), pgUserRepository.DeleteUserByIdQuery, mockUserData.BannedUser1.Id).Return(pgconn.CommandTag{}, nil)
		mockDb.EXPECT().QueryRow(gomock.Any(), pgUserRepository.HasUserByIdQuery, mockUserData.NotFoundUser.Id).Return(mock.PgMockRow{Data: []any{false}})
		mockDb.EXPECT().QueryRow(gomock.Any(), pgUserRepository.HasUserByIdQuery, mockUserData.InternalUser.Id).Return(mock.PgMockRowError{Err: errors.New("internal")})
		mockDb.EXPECT().QueryRow(gomock.Any(), pgUserRepository.HasUserByIdQuery, mockUserData.InternalUser2.Id).Return(mock.PgMockRow{Data: []any{true}})
		mockDb.EXPECT().Exec(gomock.Any(), pgUserRepository.DeleteUserByIdQuery, mockUserData.InternalUser2.Id).Return(pgconn.CommandTag{}, errors.New("internal"))

		userRepo := pgUserRepository.New(tLog, mockDb)

		t.Run("Should correct delete", func(t *testing.T) {
			ctx := testCtx.New()
			user := mockUserData.CorrectUser1
			err := userRepo.Delete(ctx, user.Id)
			require.NoError(t, err)

			user = mockUserData.AdminUser1
			err = userRepo.Delete(ctx, user.Id)
			require.NoError(t, err)

			user = mockUserData.BannedUser1
			err = userRepo.Delete(ctx, user.Id)
			require.NoError(t, err)
		})

		t.Run("Should not found deleted user", func(t *testing.T) {
			ctx := testCtx.New()
			user := mockUserData.NotFoundUser
			err := userRepo.Delete(ctx, user.Id)
			assert.Equal(t, err, pgUserRepository.ErrUserNotFound)
			assert.Equal(t, []any{pgUserRepository.TraceDelete}, tLog.Stack)
			assert.NotEmpty(t, tLog.Message)
			tLog.Clean()
		})

		t.Run("Should internal error", func(t *testing.T) {
			ctx := testCtx.New()
			user := mockUserData.InternalUser
			err := userRepo.Delete(ctx, user.Id)
			assert.Equal(t, err, pgUserRepository.ErrInternal)
			assert.Equal(t, []any{pgUserRepository.TraceDelete, pgUserRepository.TraceHasUserById}, tLog.Stack)
			assert.NotEmpty(t, tLog.Message)
			tLog.Clean()

			user = mockUserData.InternalUser2
			err = userRepo.Delete(ctx, user.Id)
			assert.Equal(t, err, pgUserRepository.ErrInternal)
			assert.Equal(t, []any{pgUserRepository.TraceDelete}, tLog.Stack)
			assert.NotEmpty(t, tLog.Message)
			tLog.Clean()
		})
	})

}
