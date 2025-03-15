-- +goose Up
-- +goose StatementBegin
CREATE TABLE limits (
    id SERIAL PRIMARY KEY,
    user_id int NOT NULL,
    time_limit smallint NOT NULL DEFAULT 7200,
    created_at TIMESTAMP NOT NULL default now(),
    updated_at TIMESTAMP
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users IF EXISTS;
-- +goose StatementEnd
