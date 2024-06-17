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

func TestGetBy(t *testing.T) {
	mockPgData := mock.PostgresData()
	mockUserData := mockPgData.User

	t.Run("Testing GetById", func(t *testing.T) {
		tLog := testLogger.New()
		ctrl := gomock.NewController(t)
		mockDb := mockgenPostgres.NewMockQuery(ctrl)

		mockDb.EXPECT().QueryRow(gomock.Any(), pgUserRepository.GetUserByIdQuery, mockUserData.CorrectUser1.Id).Return(mock.PgMockRow{
			Data: []any{mockUserData.CorrectUser1.Id, mockUserData.CorrectUser1.Login, mockUserData.CorrectUser1.Email, mockUserData.CorrectUser1.Password, mockUserData.CorrectUser1.Role, mockUserData.CorrectUser1.IsBanned, mockUserData.CorrectUser1.BanReason, mockUserData.CorrectUser1.UpdatedAt, mockUserData.CorrectUser1.CreatedAt}},
		)
		mockDb.EXPECT().QueryRow(gomock.Any(), pgUserRepository.GetUserByIdQuery, mockUserData.AdminUser1.Id).Return(mock.PgMockRow{
			Data: []any{mockUserData.AdminUser1.Id, mockUserData.AdminUser1.Login, mockUserData.AdminUser1.Email, mockUserData.AdminUser1.Password, mockUserData.AdminUser1.Role, mockUserData.AdminUser1.IsBanned, mockUserData.AdminUser1.BanReason, mockUserData.AdminUser1.UpdatedAt, mockUserData.AdminUser1.CreatedAt}},
		)
		mockDb.EXPECT().QueryRow(gomock.Any(), pgUserRepository.GetUserByIdQuery, mockUserData.BannedUser1.Id).Return(mock.PgMockRow{
			Data: []any{mockUserData.BannedUser1.Id, mockUserData.BannedUser1.Login, mockUserData.BannedUser1.Email, mockUserData.BannedUser1.Password, mockUserData.BannedUser1.Role, mockUserData.BannedUser1.IsBanned, mockUserData.BannedUser1.BanReason, mockUserData.BannedUser1.UpdatedAt, mockUserData.BannedUser1.CreatedAt}},
		)
		mockDb.EXPECT().QueryRow(gomock.Any(), pgUserRepository.GetUserByIdQuery, mockUserData.NotFoundUser.Id).Return(mock.PgMockRowError{Err: sql.ErrNoRows})
		mockDb.EXPECT().QueryRow(gomock.Any(), pgUserRepository.GetUserByIdQuery, mockUserData.InternalUser.Id).Return(mock.PgMockRowError{Err: errors.New("internal")})

		userRepo := pgUserRepository.New(tLog, mockDb)

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
			assert.Equal(t, err, pgUserRepository.ErrUserNotFound)
			assert.NotEmpty(t, tLog.Message)
			assert.Equal(t, []any{pgUserRepository.TraceGetById}, tLog.Stack)
			tLog.Clean()
		})

		t.Run("Should internal", func(t *testing.T) {
			ctx := testCtx.New()
			user := mockUserData.InternalUser
			userDb, err := userRepo.GetById(ctx, user.Id)
			assert.Nil(t, userDb)
			assert.Equal(t, err, pgUserRepository.ErrInternal)
			assert.NotEmpty(t, tLog.Message)
			assert.Equal(t, []any{pgUserRepository.TraceGetById}, tLog.Stack)
			tLog.Clean()
		})
	})

	t.Run("Testing GetByLogin", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		tLog := testLogger.New()
		mockDb := mockgenPostgres.NewMockQuery(ctrl)

		mockDb.EXPECT().QueryRow(gomock.Any(), pgUserRepository.GetUserByLoginQuery, mockUserData.CorrectUser1.Login).Return(mock.PgMockRow{
			Data: []any{mockUserData.CorrectUser1.Id, mockUserData.CorrectUser1.Login, mockUserData.CorrectUser1.Email, mockUserData.CorrectUser1.Password, mockUserData.CorrectUser1.Role, mockUserData.CorrectUser1.IsBanned, mockUserData.CorrectUser1.BanReason, mockUserData.CorrectUser1.UpdatedAt, mockUserData.CorrectUser1.CreatedAt}},
		)
		mockDb.EXPECT().QueryRow(gomock.Any(), pgUserRepository.GetUserByLoginQuery, mockUserData.AdminUser1.Login).Return(mock.PgMockRow{
			Data: []any{mockUserData.AdminUser1.Id, mockUserData.AdminUser1.Login, mockUserData.AdminUser1.Email, mockUserData.AdminUser1.Password, mockUserData.AdminUser1.Role, mockUserData.AdminUser1.IsBanned, mockUserData.AdminUser1.BanReason, mockUserData.AdminUser1.UpdatedAt, mockUserData.AdminUser1.CreatedAt}},
		)
		mockDb.EXPECT().QueryRow(gomock.Any(), pgUserRepository.GetUserByLoginQuery, mockUserData.BannedUser1.Login).Return(mock.PgMockRow{
			Data: []any{mockUserData.BannedUser1.Id, mockUserData.BannedUser1.Login, mockUserData.BannedUser1.Email, mockUserData.BannedUser1.Password, mockUserData.BannedUser1.Role, mockUserData.BannedUser1.IsBanned, mockUserData.BannedUser1.BanReason, mockUserData.BannedUser1.UpdatedAt, mockUserData.BannedUser1.CreatedAt}},
		)
		mockDb.EXPECT().QueryRow(gomock.Any(), pgUserRepository.GetUserByLoginQuery, mockUserData.NotFoundUser.Login).Return(mock.PgMockRowError{Err: sql.ErrNoRows})
		mockDb.EXPECT().QueryRow(gomock.Any(), pgUserRepository.GetUserByLoginQuery, mockUserData.InternalUser.Login).Return(mock.PgMockRowError{Err: errors.New("internal")})

		userRepo := pgUserRepository.New(tLog, mockDb)

		t.Run("Should correct work", func(t *testing.T) {
			ctx := testCtx.New()
			user := mockUserData.CorrectUser1
			userDb, err := userRepo.GetByLogin(ctx, user.Login)
			require.NoError(t, err)
			assert.Equal(t, user, userDb)

			user = mockUserData.BannedUser1
			userDb, err = userRepo.GetByLogin(ctx, user.Login)
			require.NoError(t, err)
			assert.Equal(t, user, userDb)

			user = mockUserData.AdminUser1
			userDb, err = userRepo.GetByLogin(ctx, user.Login)
			require.NoError(t, err)
			assert.Equal(t, user, userDb)
		})

		t.Run("Should not found", func(t *testing.T) {
			ctx := testCtx.New()
			user := mockUserData.NotFoundUser
			userDb, err := userRepo.GetByLogin(ctx, user.Login)
			assert.Nil(t, userDb)
			assert.Equal(t, err, pgUserRepository.ErrUserNotFound)
			assert.NotEmpty(t, tLog.Message)
			assert.Equal(t, []any{pgUserRepository.TraceGetByLogin}, tLog.Stack)
			tLog.Clean()
		})

		t.Run("Should internal", func(t *testing.T) {
			ctx := testCtx.New()
			user := mockUserData.InternalUser
			userDb, err := userRepo.GetByLogin(ctx, user.Login)
			assert.Nil(t, userDb)
			assert.Equal(t, err, pgUserRepository.ErrInternal)
			assert.NotEmpty(t, tLog.Message)
			assert.Equal(t, []any{pgUserRepository.TraceGetByLogin}, tLog.Stack)
			tLog.Clean()
		})
	})
}
