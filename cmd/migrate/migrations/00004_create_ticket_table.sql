-- +goose Up
-- +goose StatementBegin
CREATE TYPE enum_ticket_status as enum('booked','available');
CREATE TABLE ticket (
  id integer primary key generated always as identity,
  seat varchar(20) NOT NULL,
  price float DEFAULT NULL,
  event_id int REFERENCES event(id),
  user_id int DEFAULT NULL,
  status enum_ticket_status DEFAULT 'available'
) 
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS ticket;
DROP TYPE enum_ticket_status;
-- +goose StatementEnd
