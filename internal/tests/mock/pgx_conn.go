package mock

import (
	"database/sql"
	"errors"
	postgresRepository "github.com/OddEer0/mirage-auth-service/internal/infrastructure/storage/postgres_repository"
	mockgenPostgres "github.com/OddEer0/mirage-auth-service/internal/tests/mockgen/mockgen_postgres"
	"github.com/golang/mock/gomock"
)

func PgxConnMock(ctrl *gomock.Controller, mockPgData *Postgres) *mockgenPostgres.MockQuery {
	mockUserData := mockPgData.User
	mockDb := mockgenPostgres.NewMockQuery(ctrl)

	// Create GetById
	mockDb.EXPECT().QueryRow(gomock.Any(), postgresRepository.GetUserByIdQuery, mockUserData.CorrectUser1.Id).Return(PgMockRow{
		Data: []any{mockUserData.CorrectUser1.Id, mockUserData.CorrectUser1.Login, mockUserData.CorrectUser1.Email, mockUserData.CorrectUser1.Password, mockUserData.CorrectUser1.Role, mockUserData.CorrectUser1.IsBanned, mockUserData.CorrectUser1.BanReason, mockUserData.CorrectUser1.UpdatedAt, mockUserData.CorrectUser1.CreatedAt}},
	)
	mockDb.EXPECT().QueryRow(gomock.Any(), postgresRepository.GetUserByIdQuery, mockUserData.AdminUser1.Id).Return(PgMockRow{
		Data: []any{mockUserData.AdminUser1.Id, mockUserData.AdminUser1.Login, mockUserData.AdminUser1.Email, mockUserData.AdminUser1.Password, mockUserData.AdminUser1.Role, mockUserData.AdminUser1.IsBanned, mockUserData.AdminUser1.BanReason, mockUserData.AdminUser1.UpdatedAt, mockUserData.AdminUser1.CreatedAt}},
	)
	mockDb.EXPECT().QueryRow(gomock.Any(), postgresRepository.GetUserByIdQuery, mockUserData.BannedUser1.Id).Return(PgMockRow{
		Data: []any{mockUserData.BannedUser1.Id, mockUserData.BannedUser1.Login, mockUserData.BannedUser1.Email, mockUserData.BannedUser1.Password, mockUserData.BannedUser1.Role, mockUserData.BannedUser1.IsBanned, mockUserData.BannedUser1.BanReason, mockUserData.BannedUser1.UpdatedAt, mockUserData.BannedUser1.CreatedAt}},
	)
	mockDb.EXPECT().QueryRow(gomock.Any(), postgresRepository.GetUserByIdQuery, mockUserData.NotFoundUser.Id).Return(PgMockRowError{err: sql.ErrNoRows})
	mockDb.EXPECT().QueryRow(gomock.Any(), postgresRepository.GetUserByIdQuery, mockUserData.InternalUser.Id).Return(PgMockRowError{err: errors.New("internal")})

	// Create
	mockDb.EXPECT().QueryRow(gomock.Any(), postgresRepository.CreateUserQuery, mockUserData.CreateUser1Res.Id, mockUserData.CreateUser1Res.Login, mockUserData.CreateUser1Res.Email, mockUserData.CreateUser1Res.Password, mockUserData.CreateUser1Res.Role, mockUserData.CreateUser1Res.IsBanned, mockUserData.CreateUser1Res.BanReason, mockUserData.CreateUser1Res.UpdatedAt, mockUserData.CreateUser1Res.CreatedAt).
		Return(PgMockRow{
			Data: []any{mockUserData.CreateUser1Res.Id, mockUserData.CreateUser1Res.Login, mockUserData.CreateUser1Res.Email, mockUserData.CreateUser1Res.Password, mockUserData.CreateUser1Res.Role, mockUserData.CreateUser1Res.IsBanned, mockUserData.CreateUser1Res.BanReason, mockUserData.CreateUser1Res.UpdatedAt, mockUserData.CreateUser1Res.CreatedAt},
		})

	return mockDb
}
