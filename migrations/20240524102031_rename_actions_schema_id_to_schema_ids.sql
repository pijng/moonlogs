-- +goose Up
-- +goose StatementBegin
ALTER TABLE actions DROP COLUMN schema_id;
ALTER TABLE actions DROP COLUMN schema_name;
ALTER TABLE actions ADD COLUMN schema_ids TEXT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE actions ADD COLUMN schema_id;
ALTER TABLE actions ADD COLUMN schema_name TEXT NOT NULL;
ALTER TABLE actions DROP COLUMN schema_ids TEXT;
-- +goose StatementEnd
