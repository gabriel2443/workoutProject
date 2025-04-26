-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS users(
 id BIGSERIAL PRIMARY KEY,
 username VARCHAR(50) UNIQUE NOT NULL,
 email VARCHAR(255) UNIQUE NOT NULL,
 passwrod_hash VARCHAR(255) NOT NULL,
 bio TEXT,
 created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
 updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,


)

-- +goose StatementEnd
-- +goose Down
-- +Goose StatementBegin
DROP TABLE users;
-- +goose statementEnd