-- +goose Up
-- +goose StatementBegin
ALTER TABLE users
    ADD COLUMN created_at TIMESTAMP NOT NULL DEFAULT now(),
    ADD COLUMN updated_at TIMESTAMP;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users
    DROP COLUMN created_at,
    DROP COLUMN updated_at;
-- +goose StatementEnd
