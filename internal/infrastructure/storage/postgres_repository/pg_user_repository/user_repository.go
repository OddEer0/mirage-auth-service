package pgUserRepository

import (
	"github.com/OddEer0/mirage-auth-service/internal/domain"
	"github.com/OddEer0/mirage-auth-service/internal/domain/repository"
	"github.com/OddEer0/mirage-auth-service/internal/infrastructure/storage/postgres"
	stacktrace "github.com/OddEer0/stack-trace/stack_trace"
)

var (
	ErrUserNotFound = domain.NewErr(domain.ErrNotFoundCode, "user not found")
	ErrInternal     = domain.NewErr(domain.ErrInternalCode, "internal error")

	TraceGetById            = &stacktrace.Method{Package: "pgUserRepository", Type: "userRepository", Method: "GetById"}
	TraceGetByLogin         = &stacktrace.Method{Package: "pgUserRepository", Type: "userRepository", Method: "GetByLogin"}
	TraceCreate             = &stacktrace.Method{Package: "pgUserRepository", Type: "userRepository", Method: "Create"}
	TraceCheckUserRole      = &stacktrace.Method{Package: "pgUserRepository", Type: "userRepository", Method: "CheckUserRole"}
	TraceGetByQuery         = &stacktrace.Method{Package: "pgUserRepository", Type: "userRepository", Method: "GetByQuery"}
	TraceDelete             = &stacktrace.Method{Package: "pgUserRepository", Type: "userRepository", Method: "Delete"}
	TraceUpdateById         = &stacktrace.Method{Package: "pgUserRepository", Type: "userRepository", Method: "UpdateById"}
	TraceUpdateRoleById     = &stacktrace.Method{Package: "pgUserRepository", Type: "userRepository", Method: "UpdateRoleById"}
	TraceUpdatePasswordById = &stacktrace.Method{Package: "pgUserRepository", Type: "userRepository", Method: "UpdatePasswordById"}
	TraceBanUserById        = &stacktrace.Method{Package: "pgUserRepository", Type: "userRepository", Method: "BanUserById"}
	TraceUnbanUserById      = &stacktrace.Method{Package: "pgUserRepository", Type: "userRepository", Method: "UnbanUserById"}
	TraceHasUserById        = &stacktrace.Method{Package: "pgUserRepository", Type: "userRepository", Method: "HasUserById"}
	TraceHasUserByLogin     = &stacktrace.Method{Package: "pgUserRepository", Type: "userRepository", Method: "HasUserByLogin"}
	TraceHasUserByEmail     = &stacktrace.Method{Package: "pgUserRepository", Type: "userRepository", Method: "HasUserByEmail"}
)

type userRepository struct {
	log domain.Logger
	db  postgres.Query
}

func New(logger domain.Logger, db postgres.Query) repository.UserRepository {
	return &userRepository{db: db, log: logger}
}
