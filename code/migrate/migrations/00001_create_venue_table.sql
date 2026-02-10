-- +goose Up
-- +goose StatementBegin
CREATE TABLE venue (
  id integer primary key generated always as identity,
  name VARCHAR(255) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS venue;
-- +goose StatementEnd
