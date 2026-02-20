-- +goose Up
-- +goose StatementBegin
ALTER TABLE "event" ADD deleted_at TIMESTAMP;
ALTER TABLE venue ADD deleted_at TIMESTAMP;
ALTER TABLE performer ADD deleted_at TIMESTAMP;
ALTER TABLE ticket ADD deleted_at TIMESTAMP;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "event" DROP deleted_at;
ALTER TABLE venue DROP deleted_at;
ALTER TABLE performer DROP deleted_at;
ALTER TABLE ticket DROP deleted_at;
-- +goose StatementEnd
