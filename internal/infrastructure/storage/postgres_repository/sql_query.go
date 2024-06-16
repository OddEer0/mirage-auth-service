package postgresRepository

const (
	CreateUserQuery = `INSERT INTO users (id, login, email, password, role, isBanned, banReason, updatedAt, createdAt)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id, login, email, password, role, isBanned, banReason, updatedAt, createdAt;`
	GetUserByIdQuery = `
		SELECT id, login, email, password, role, isBanned, banReason, updatedAt, createdAt FROM users
		WHERE id = $1;
	`
	GetUserByLoginQuery = `
		SELECT id, login, email, password, role, isBanned, banReason, updatedAt, createdAt FROM users
		WHERE login = $1;
	`
	HasUserByLoginQuery = `
		SELECT EXISTS(SELECT 1 FROM users WHERE login = $1);
	`
	HasUserByIdQuery = `
		SELECT EXISTS(SELECT 1 FROM users WHERE id = $1);
	`
	HasUserByEmailQuery = `
		SELECT EXISTS(SELECT 1 FROM users WHERE email = $1);
	`
	DeleteUserByIdQuery = `
		DELETE FROM users WHERE id = $1;
	`
	UpdateUserById = `
		UPDATE users SET login = $2, email = $3, password = $4, role = $5, isBanned = $6, banReason = $7, updatedAt = $8
		WHERE id = $1
		RETURNING id, login, email, password, role, isBanned, banReason, updatedAt, createdAt;
	`
	UpdateUserRoleById = `
		UPDATE users SET role = $2
		WHERE id = $1
		RETURNING id, login, email, password, role, isBanned, banReason, updatedAt, createdAt;
	`
	UpdateUserPasswordById = `
		UPDATE users SET password = $2
		WHERE id = $1
		RETURNING id, login, email, password, role, isBanned, banReason, updatedAt, createdAt;
	`
	UpdateUserBanById = `
		UPDATE users SET isBanned = $2, banReason = $3
		WHERE id = $1
		RETURNING id, login, email, password, role, isBanned, banReason, updatedAt, createdAt;
	`
	CheckUserRoleQuery = `
		SELECT EXISTS(SELECT 1 FROM users WHERE id = $1 AND role = $2);
	`
)
