package pg_user_repository_test

import (
	"database/sql"
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

func TestCheck(t *testing.T) {
	mockPgData := mock.PostgresData()
	mockUserData := mockPgData.User

	t.Run("Testing CheckUserRole", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		tLog := testLogger.New()
		mockDb := mockgenPostgres.NewMockQuery(ctrl)

		mockDb.EXPECT().QueryRow(gomock.Any(), postgresRepository.CheckUserRoleQuery, mockUserData.AdminUser1.Id, mockUserData.AdminUser1.Role).Return(mock.PgMockRow{
			Data: []any{true},
		})
		mockDb.EXPECT().QueryRow(gomock.Any(), postgresRepository.CheckUserRoleQuery, mockUserData.CorrectUser1.Id, mockUserData.CorrectUser1.Role).Return(mock.PgMockRow{
			Data: []any{true},
		})
		mockDb.EXPECT().QueryRow(gomock.Any(), postgresRepository.CheckUserRoleQuery, mockUserData.AdminUser1.Id, mockUserData.CorrectUser1.Role).Return(mock.PgMockRow{
			Data: []any{false},
		})
		mockDb.EXPECT().QueryRow(gomock.Any(), postgresRepository.CheckUserRoleQuery, mockUserData.CorrectUser1.Id, mockUserData.AdminUser1.Role).Return(mock.PgMockRow{
			Data: []any{false},
		})
		mockDb.EXPECT().QueryRow(gomock.Any(), postgresRepository.CheckUserRoleQuery, mockUserData.NotFoundUser.Id, mockUserData.NotFoundUser.Role).Return(mock.PgMockRowError{Err: sql.ErrNoRows})
		mockDb.EXPECT().QueryRow(gomock.Any(), postgresRepository.CheckUserRoleQuery, mockUserData.InternalUser.Id, mockUserData.InternalUser.Role).Return(mock.PgMockRowError{Err: errors.New("internal")})

		userRepo := postgresRepository.NewUserRepository(tLog, mockDb)

		t.Run("Should correct check", func(t *testing.T) {
			ctx := testCtx.New()
			user := mockUserData.CorrectUser1
			adminUser := mockUserData.AdminUser1
			check, err := userRepo.CheckUserRole(ctx, user.Id, user.Role)
			require.NoError(t, err)
			assert.True(t, check)
			check, err = userRepo.CheckUserRole(ctx, adminUser.Id, adminUser.Role)
			require.NoError(t, err)
			assert.True(t, check)
			check, err = userRepo.CheckUserRole(ctx, user.Id, adminUser.Role)
			require.NoError(t, err)
			assert.False(t, check)
			check, err = userRepo.CheckUserRole(ctx, adminUser.Id, user.Role)
			require.NoError(t, err)
			assert.False(t, check)
		})

		t.Run("Should not found", func(t *testing.T) {
			ctx := testCtx.New()
			user := mockUserData.NotFoundUser
			check, err := userRepo.CheckUserRole(ctx, user.Id, user.Role)
			assert.False(t, check)
			assert.Equal(t, postgresRepository.ErrUserNotFound, err)
			assert.NotEmpty(t, tLog.Message)
			assert.Equal(t, []any{postgresRepository.TraceUserRepoCheckUserRole}, tLog.Stack)
			tLog.Clean()
		})

		t.Run("Should internal", func(t *testing.T) {
			ctx := testCtx.New()
			user := mockUserData.InternalUser
			check, err := userRepo.CheckUserRole(ctx, user.Id, user.Role)
			assert.False(t, check)
			assert.Equal(t, postgresRepository.ErrInternal, err)
			assert.NotEmpty(t, tLog.Message)
			assert.Equal(t, []any{postgresRepository.TraceUserRepoCheckUserRole}, tLog.Stack)
			tLog.Clean()
		})
	})

	t.Run("Testing HasUserById", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		tLog := testLogger.New()
		mockDb := mockgenPostgres.NewMockQuery(ctrl)

		mockDb.EXPECT().QueryRow(gomock.Any(), postgresRepository.HasUserByIdQuery, mockUserData.CorrectUser1.Id).Return(mock.PgMockRow{Data: []any{true}})
		mockDb.EXPECT().QueryRow(gomock.Any(), postgresRepository.HasUserByIdQuery, mockUserData.AdminUser1.Id).Return(mock.PgMockRow{Data: []any{true}})
		mockDb.EXPECT().QueryRow(gomock.Any(), postgresRepository.HasUserByIdQuery, mockUserData.BannedUser1.Id).Return(mock.PgMockRow{Data: []any{true}})
		mockDb.EXPECT().QueryRow(gomock.Any(), postgresRepository.HasUserByIdQuery, mockUserData.NotFoundUser.Id).Return(mock.PgMockRow{Data: []any{false}})
		mockDb.EXPECT().QueryRow(gomock.Any(), postgresRepository.HasUserByIdQuery, mockUserData.InternalUser.Id).Return(mock.PgMockRowError{Err: errors.New("internal")})

		userRepo := postgresRepository.NewUserRepository(tLog, mockDb)

		t.Run("Should correct work", func(t *testing.T) {
			ctx := testCtx.New()
			user := mockUserData.CorrectUser1
			has, err := userRepo.HasUserById(ctx, user.Id)
			require.NoError(t, err)
			assert.True(t, has)

			user = mockUserData.AdminUser1
			has, err = userRepo.HasUserById(ctx, user.Id)
			require.NoError(t, err)
			assert.True(t, has)

			user = mockUserData.BannedUser1
			has, err = userRepo.HasUserById(ctx, user.Id)
			require.NoError(t, err)
			assert.True(t, has)

			user = mockUserData.NotFoundUser
			has, err = userRepo.HasUserById(ctx, user.Id)
			require.NoError(t, err)
			assert.False(t, has)
		})

		t.Run("Should internal", func(t *testing.T) {
			ctx := testCtx.New()
			user := mockUserData.InternalUser
			has, err := userRepo.HasUserById(ctx, user.Id)
			assert.False(t, has)
			assert.Equal(t, err, postgresRepository.ErrInternal)
			assert.NotEmpty(t, tLog.Message)
			assert.Equal(t, []any{postgresRepository.TraceUserRepoHasUserById}, tLog.Stack)
			tLog.Clean()
		})
	})

	t.Run("Testing HasUserByLogin", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		tLog := testLogger.New()
		mockDb := mockgenPostgres.NewMockQuery(ctrl)

		mockDb.EXPECT().QueryRow(gomock.Any(), postgresRepository.HasUserByLoginQuery, mockUserData.CorrectUser1.Login).Return(mock.PgMockRow{Data: []any{true}})
		mockDb.EXPECT().QueryRow(gomock.Any(), postgresRepository.HasUserByLoginQuery, mockUserData.AdminUser1.Login).Return(mock.PgMockRow{Data: []any{true}})
		mockDb.EXPECT().QueryRow(gomock.Any(), postgresRepository.HasUserByLoginQuery, mockUserData.BannedUser1.Login).Return(mock.PgMockRow{Data: []any{true}})
		mockDb.EXPECT().QueryRow(gomock.Any(), postgresRepository.HasUserByLoginQuery, mockUserData.NotFoundUser.Login).Return(mock.PgMockRow{Data: []any{false}})
		mockDb.EXPECT().QueryRow(gomock.Any(), postgresRepository.HasUserByLoginQuery, mockUserData.InternalUser.Login).Return(mock.PgMockRowError{Err: errors.New("internal")})

		userRepo := postgresRepository.NewUserRepository(tLog, mockDb)

		t.Run("Should correct work", func(t *testing.T) {
			ctx := testCtx.New()
			user := mockUserData.CorrectUser1
			has, err := userRepo.HasUserByLogin(ctx, user.Login)
			require.NoError(t, err)
			assert.True(t, has)

			user = mockUserData.AdminUser1
			has, err = userRepo.HasUserByLogin(ctx, user.Login)
			require.NoError(t, err)
			assert.True(t, has)

			user = mockUserData.BannedUser1
			has, err = userRepo.HasUserByLogin(ctx, user.Login)
			require.NoError(t, err)
			assert.True(t, has)

			user = mockUserData.NotFoundUser
			has, err = userRepo.HasUserByLogin(ctx, user.Login)
			require.NoError(t, err)
			assert.False(t, has)
		})

		t.Run("Should internal", func(t *testing.T) {
			ctx := testCtx.New()
			user := mockUserData.InternalUser
			has, err := userRepo.HasUserByLogin(ctx, user.Login)
			assert.False(t, has)
			assert.Equal(t, err, postgresRepository.ErrInternal)
			assert.NotEmpty(t, tLog.Message)
			assert.Equal(t, []any{postgresRepository.TraceUserRepoHasUserByLogin}, tLog.Stack)
			tLog.Clean()
		})
	})

	t.Run("Testing HasUserByEmail", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		tLog := testLogger.New()
		mockDb := mockgenPostgres.NewMockQuery(ctrl)

		mockDb.EXPECT().QueryRow(gomock.Any(), postgresRepository.HasUserByEmailQuery, mockUserData.CorrectUser1.Email).Return(mock.PgMockRow{Data: []any{true}})
		mockDb.EXPECT().QueryRow(gomock.Any(), postgresRepository.HasUserByEmailQuery, mockUserData.AdminUser1.Email).Return(mock.PgMockRow{Data: []any{true}})
		mockDb.EXPECT().QueryRow(gomock.Any(), postgresRepository.HasUserByEmailQuery, mockUserData.BannedUser1.Email).Return(mock.PgMockRow{Data: []any{true}})
		mockDb.EXPECT().QueryRow(gomock.Any(), postgresRepository.HasUserByEmailQuery, mockUserData.NotFoundUser.Email).Return(mock.PgMockRow{Data: []any{false}})
		mockDb.EXPECT().QueryRow(gomock.Any(), postgresRepository.HasUserByEmailQuery, mockUserData.InternalUser.Email).Return(mock.PgMockRowError{Err: errors.New("internal")})

		userRepo := postgresRepository.NewUserRepository(tLog, mockDb)

		t.Run("Should correct work", func(t *testing.T) {
			ctx := testCtx.New()
			user := mockUserData.CorrectUser1
			has, err := userRepo.HasUserByEmail(ctx, user.Email)
			require.NoError(t, err)
			assert.True(t, has)

			user = mockUserData.AdminUser1
			has, err = userRepo.HasUserByEmail(ctx, user.Email)
			require.NoError(t, err)
			assert.True(t, has)

			user = mockUserData.BannedUser1
			has, err = userRepo.HasUserByEmail(ctx, user.Email)
			require.NoError(t, err)
			assert.True(t, has)

			user = mockUserData.NotFoundUser
			has, err = userRepo.HasUserByEmail(ctx, user.Email)
			require.NoError(t, err)
			assert.False(t, has)
		})

		t.Run("Should internal", func(t *testing.T) {
			ctx := testCtx.New()
			user := mockUserData.InternalUser
			has, err := userRepo.HasUserByEmail(ctx, user.Email)
			assert.False(t, has)
			assert.Equal(t, err, postgresRepository.ErrInternal)
			assert.NotEmpty(t, tLog.Message)
			assert.Equal(t, []any{postgresRepository.TraceUserRepoHasUserByEmail}, tLog.Stack)
			tLog.Clean()
		})
	})
}
