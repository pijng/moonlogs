-- +goose Up
-- +goose StatementBegin
ALTER TABLE incidents ADD COLUMN schema_name TEXT;
ALTER TABLE incidents ADD COLUMN severity TEXT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE incidents DROP COLUMN schema_name;
ALTER TABLE incidents DROP COLUMN severity;
-- +goose StatementEnd
