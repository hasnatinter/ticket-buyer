-- +goose Up
ALTER TABLE "ticket"
ADD COLUMN booking_id INT REFERENCES booking(id);

-- +goose Down
ALTER TABLE "ticket" DROP booking_id;
