-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY,
    login VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(50) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(30) NOT NULL,
    isBanned BOOLEAN NOT NULL,
    banReason VARCHAR(255),
    updatedAt DATE NOT NULL,
    createdAt DATE NOT NULL
);

CREATE TABLE IF NOT EXISTS tokens (
    id UUID PRIMARY KEY REFERENCES users(id),
    value VARCHAR(255) NOT NULL,
    updatedAt DATE NOT NULL,
    createdAt DATE NOT NULL
);

CREATE TABLE IF NOT EXISTS user_activate (
    id UUID PRIMARY KEY REFERENCES users(id),
    isActivate BOOLEAN NOT NULL,
    link VARCHAR(255) NOT NULL,
    updatedAt DATE NOT NULL,
    createdAt DATE NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_activate;
DROP TABLE IF EXISTS tokens;
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
