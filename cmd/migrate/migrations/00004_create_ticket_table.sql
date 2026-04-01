-- +goose Up
-- +goose StatementBegin
CREATE TYPE enum_ticket_status AS ENUM('Booked','Available');
CREATE TABLE ticket (
  id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  seat VARCHAR(20) NOT NULL,
  price FLOAT DEFAULT NULL,
  event_id INT REFERENCES event(id),
  user_id INT DEFAULT NULL,
  status enum_ticket_status DEFAULT 'Available'
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS ticket;
DROP TYPE enum_ticket_status;
-- +goose StatementEnd
