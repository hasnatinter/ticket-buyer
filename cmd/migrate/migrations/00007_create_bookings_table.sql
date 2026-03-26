-- +goose Up
CREATE TYPE enum_booking_status as enum('booked','reserved');

CREATE TABLE booking (
  id integer primary key generated always as identity,
  user_name VARCHAR(255) NOT NULL,
  user_address VARCHAR(255) NOT NULL,
  ticket_ids integer[],
  status enum_booking_status DEFAULT 'booked',
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP 
);

CREATE INDEX idx_status ON ticket (status);

-- +goose Down
DROP TABLE IF EXISTS booking;

DROP TYPE IF EXISTS enum_booking_status;