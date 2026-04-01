-- +goose Up
-- +goose StatementBegin
CREATE TABLE venue (
  id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  name VARCHAR(255) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS venue;
-- +goose StatementEnd
