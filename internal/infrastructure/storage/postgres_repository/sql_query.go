package postgresRepository

const (
	CreateUserQuery = `INSERT INTO users (id, login, email, password, role, isBanned, banReason, updatedAt, createdAt)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id, login, email, password, role, isBanned, banReason, updatedAt, createdAt;`
	GetUserByIdQuery = `
		SELECT id, login, email, password, role, isBanned, banReason, updatedAt, createdAt FROM users
		WHERE id = $1;
	`
	getUserByLoginQuery = `
		SELECT id, login, email, password, role, isBanned, banReason, updatedAt, createdAt FROM users
		WHERE login = $1;
	`
	hasUserByLoginQuery = `
		SELECT EXISTS(SELECT 1 FROM users WHERE login = $1);
	`
	hasUserByIdQuery = `
		SELECT EXISTS(SELECT 1 FROM users WHERE id = $1);
	`
	hasUserByEmailQuery = `
		SELECT EXISTS(SELECT 1 FROM users WHERE email = $1);
	`
	deleteUserById = `
		DELETE FROM users WHERE id = $1;
	`
	updateUserById = `
		UPDATE users SET login = $2, email = $3, password = $4, role = $5, isBanned = $6, banReason = $7, updatedAt = $8
		WHERE id = $1
		RETURNING id, login, email, password, role, isBanned, banReason, updatedAt, createdAt;
	`
	updateUserRoleById = `
		UPDATE users SET role = $2
		WHERE id = $1
		RETURNING id, login, email, password, role, isBanned, banReason, updatedAt, createdAt;
	`
	updateUserPasswordById = `
		UPDATE users SET password = $2
		WHERE id = $1
		RETURNING id, login, email, password, role, isBanned, banReason, updatedAt, createdAt;
	`
	updateUserBanById = `
		UPDATE users SET isBanned = $2, banReason = $3
		WHERE id = $1
		RETURNING id, login, email, password, role, isBanned, banReason, updatedAt, createdAt;
	`
	checkUserRoleQuery = `
		SELECT EXISTS(SELECT 1 FROM users WHERE id = $1 AND role = $2);
	`
)
