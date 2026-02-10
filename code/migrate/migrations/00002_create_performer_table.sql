-- +goose Up
-- +goose StatementBegin
CREATE TABLE performer (
  id integer primary key generated always as identity,
  name varchar(255) NOT NULL
); 
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS performer;
-- +goose StatementEnd
