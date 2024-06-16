package pg_repository_test

import (
	"database/sql"
	"errors"
	postgresRepository "github.com/OddEer0/mirage-auth-service/internal/infrastructure/storage/postgres_repository"
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

func TestUserPgRepository(t *testing.T) {
	mockPgData := mock.PostgresData()
	mockUserData := mockPgData.User

	t.Run("Testing GetById", func(t *testing.T) {
		tLog := testLogger.New()
		ctrl := gomock.NewController(t)
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
		mockDb.EXPECT().QueryRow(gomock.Any(), postgresRepository.GetUserByIdQuery, mockUserData.NotFoundUser.Id).Return(mock.PgMockRowError{Err: sql.ErrNoRows})
		mockDb.EXPECT().QueryRow(gomock.Any(), postgresRepository.GetUserByIdQuery, mockUserData.InternalUser.Id).Return(mock.PgMockRowError{Err: errors.New("internal")})

		userRepo := postgresRepository.NewUserRepository(tLog, mockDb)

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

	t.Run("Testing GetByLogin", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		tLog := testLogger.New()
		mockDb := mockgenPostgres.NewMockQuery(ctrl)

		mockDb.EXPECT().QueryRow(gomock.Any(), postgresRepository.GetUserByLoginQuery, mockUserData.CorrectUser1.Login).Return(mock.PgMockRow{
			Data: []any{mockUserData.CorrectUser1.Id, mockUserData.CorrectUser1.Login, mockUserData.CorrectUser1.Email, mockUserData.CorrectUser1.Password, mockUserData.CorrectUser1.Role, mockUserData.CorrectUser1.IsBanned, mockUserData.CorrectUser1.BanReason, mockUserData.CorrectUser1.UpdatedAt, mockUserData.CorrectUser1.CreatedAt}},
		)
		mockDb.EXPECT().QueryRow(gomock.Any(), postgresRepository.GetUserByLoginQuery, mockUserData.AdminUser1.Login).Return(mock.PgMockRow{
			Data: []any{mockUserData.AdminUser1.Id, mockUserData.AdminUser1.Login, mockUserData.AdminUser1.Email, mockUserData.AdminUser1.Password, mockUserData.AdminUser1.Role, mockUserData.AdminUser1.IsBanned, mockUserData.AdminUser1.BanReason, mockUserData.AdminUser1.UpdatedAt, mockUserData.AdminUser1.CreatedAt}},
		)
		mockDb.EXPECT().QueryRow(gomock.Any(), postgresRepository.GetUserByLoginQuery, mockUserData.BannedUser1.Login).Return(mock.PgMockRow{
			Data: []any{mockUserData.BannedUser1.Id, mockUserData.BannedUser1.Login, mockUserData.BannedUser1.Email, mockUserData.BannedUser1.Password, mockUserData.BannedUser1.Role, mockUserData.BannedUser1.IsBanned, mockUserData.BannedUser1.BanReason, mockUserData.BannedUser1.UpdatedAt, mockUserData.BannedUser1.CreatedAt}},
		)
		mockDb.EXPECT().QueryRow(gomock.Any(), postgresRepository.GetUserByLoginQuery, mockUserData.NotFoundUser.Login).Return(mock.PgMockRowError{Err: sql.ErrNoRows})
		mockDb.EXPECT().QueryRow(gomock.Any(), postgresRepository.GetUserByLoginQuery, mockUserData.InternalUser.Login).Return(mock.PgMockRowError{Err: errors.New("internal")})

		userRepo := postgresRepository.NewUserRepository(tLog, mockDb)

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
			assert.Equal(t, err, postgresRepository.ErrUserNotFound)
			assert.NotEmpty(t, tLog.Message)
			assert.Equal(t, []any{postgresRepository.TraceUserRepoGetByLogin}, tLog.Stack)
			tLog.Clean()
		})

		t.Run("Should internal", func(t *testing.T) {
			ctx := testCtx.New()
			user := mockUserData.InternalUser
			userDb, err := userRepo.GetByLogin(ctx, user.Login)
			assert.Nil(t, userDb)
			assert.Equal(t, err, postgresRepository.ErrInternal)
			assert.NotEmpty(t, tLog.Message)
			assert.Equal(t, []any{postgresRepository.TraceUserRepoGetByLogin}, tLog.Stack)
			tLog.Clean()
		})
	})

	t.Run("Testing Create", func(t *testing.T) {
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
	})

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

	t.Run("Testing delete", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		tLog := testLogger.New()
		mockDb := mockgenPostgres.NewMockQuery(ctrl)

		mockDb.EXPECT().QueryRow(gomock.Any(), postgresRepository.HasUserByIdQuery, mockUserData.CorrectUser1.Id).Return(mock.PgMockRow{Data: []any{true}})
		mockDb.EXPECT().Exec(gomock.Any(), postgresRepository.DeleteUserByIdQuery, mockUserData.CorrectUser1.Id).Return(pgconn.CommandTag{}, nil)
		mockDb.EXPECT().QueryRow(gomock.Any(), postgresRepository.HasUserByIdQuery, mockUserData.AdminUser1.Id).Return(mock.PgMockRow{Data: []any{true}})
		mockDb.EXPECT().Exec(gomock.Any(), postgresRepository.DeleteUserByIdQuery, mockUserData.AdminUser1.Id).Return(pgconn.CommandTag{}, nil)
		mockDb.EXPECT().QueryRow(gomock.Any(), postgresRepository.HasUserByIdQuery, mockUserData.BannedUser1.Id).Return(mock.PgMockRow{Data: []any{true}})
		mockDb.EXPECT().Exec(gomock.Any(), postgresRepository.DeleteUserByIdQuery, mockUserData.BannedUser1.Id).Return(pgconn.CommandTag{}, nil)
		mockDb.EXPECT().QueryRow(gomock.Any(), postgresRepository.HasUserByIdQuery, mockUserData.NotFoundUser.Id).Return(mock.PgMockRow{Data: []any{false}})
		mockDb.EXPECT().QueryRow(gomock.Any(), postgresRepository.HasUserByIdQuery, mockUserData.InternalUser.Id).Return(mock.PgMockRowError{Err: errors.New("internal")})
		mockDb.EXPECT().QueryRow(gomock.Any(), postgresRepository.HasUserByIdQuery, mockUserData.InternalUser2.Id).Return(mock.PgMockRow{Data: []any{true}})
		mockDb.EXPECT().Exec(gomock.Any(), postgresRepository.DeleteUserByIdQuery, mockUserData.InternalUser2.Id).Return(pgconn.CommandTag{}, errors.New("internal"))

		userRepo := postgresRepository.NewUserRepository(tLog, mockDb)

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
			assert.Equal(t, err, postgresRepository.ErrUserNotFound)
			assert.Equal(t, []any{postgresRepository.TraceUserRepoDelete}, tLog.Stack)
			assert.NotEmpty(t, tLog.Message)
			tLog.Clean()
		})

		t.Run("Should internal error", func(t *testing.T) {
			ctx := testCtx.New()
			user := mockUserData.InternalUser
			err := userRepo.Delete(ctx, user.Id)
			assert.Equal(t, err, postgresRepository.ErrInternal)
			assert.Equal(t, []any{postgresRepository.TraceUserRepoDelete, postgresRepository.TraceUserRepoHasUserById}, tLog.Stack)
			assert.NotEmpty(t, tLog.Message)
			tLog.Clean()

			user = mockUserData.InternalUser2
			err = userRepo.Delete(ctx, user.Id)
			assert.Equal(t, err, postgresRepository.ErrInternal)
			assert.Equal(t, []any{postgresRepository.TraceUserRepoDelete}, tLog.Stack)
			assert.NotEmpty(t, tLog.Message)
			tLog.Clean()
		})
	})
}
