-- +goose Up
-- +goose StatementBegin
CREATE TABLE event (
  id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  name VARCHAR(255) NOT NULL,
  venue_id INT REFERENCES venue(id),
  performer_id INT REFERENCES performer(id),
  description VARCHAR(1000) DEFAULT NULL,
  start_time TIMESTAMPTZ DEFAULT NULL,
  category VARCHAR(100) DEFAULT NULL
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS event;
-- +goose StatementEnd
