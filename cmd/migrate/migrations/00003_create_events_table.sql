-- +goose Up
-- +goose StatementBegin
CREATE TABLE event (
  id integer primary key generated always as identity,
  name varchar(255) NOT NULL,
  venue_id int REFERENCES venue(id),
  performer_id int REFERENCES performer(id),
  description varchar(1000) DEFAULT NULL,
  start_time timestamptz DEFAULT NULL,
  category varchar(100) DEFAULT NULL
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS event;
-- +goose StatementEnd
