-- +goose Up
-- +goose StatementBegin
ALTER TABLE "event" ADD created_at TIMESTAMP;
ALTER TABLE "event" ADD updated_at TIMESTAMP;
ALTER TABLE venue ADD created_at TIMESTAMP;
ALTER TABLE venue ADD updated_at TIMESTAMP;
ALTER TABLE performer ADD created_at TIMESTAMP;
ALTER TABLE performer ADD updated_at TIMESTAMP;
ALTER TABLE ticket ADD created_at TIMESTAMP;
ALTER TABLE ticket ADD updated_at TIMESTAMP;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "event" DROP created_at;
ALTER TABLE "event" DROP updated_at;
ALTER TABLE venue DROP created_at;
ALTER TABLE venue DROP updated_at;
ALTER TABLE performer DROP created_at;
ALTER TABLE performer DROP updated_at;
ALTER TABLE ticket DROP created_at;
ALTER TABLE ticket DROP updated_at;
-- +goose StatementEnd
