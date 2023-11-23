-- +goose Up
-- +goose StatementBegin
CREATE TABLE businesses (
    datetime TIMESTAMP,
    priority INTEGER,
    deadline TIMESTAMP,
    description TEXT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE businesses;
-- +goose StatementEnd
