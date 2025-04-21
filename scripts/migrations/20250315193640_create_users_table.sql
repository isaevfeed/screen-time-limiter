-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR,
    last_name VARCHAR
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
