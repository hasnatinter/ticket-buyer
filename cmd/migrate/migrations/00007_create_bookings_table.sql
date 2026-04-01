-- +goose Up
CREATE TYPE enum_booking_status AS ENUM('Booked','Reserved');

CREATE TABLE booking (
  id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  user_name VARCHAR(255) NOT NULL,
  user_address VARCHAR(255) NOT NULL,
  ticket_ids INTEGER[],
  status enum_booking_status DEFAULT 'Booked',
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP
);

CREATE INDEX idx_status ON ticket (status);

-- +goose Down
DROP TABLE IF EXISTS booking;

DROP TYPE IF EXISTS enum_booking_status;
