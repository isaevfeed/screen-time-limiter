-- +goose Up
-- +goose StatementBegin
CREATE TABLE limit_histories (
    id SERIAL PRIMARY KEY,
    limit_id int,
    time_amount smallint,
    sent_at timestamp,
    limit_date date
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE limit_histories;
-- +goose StatementEnd
