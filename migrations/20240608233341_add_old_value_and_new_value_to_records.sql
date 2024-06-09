-- +goose Up
-- +goose StatementBegin
ALTER TABLE records ADD COLUMN old_value TEXT;
ALTER TABLE records ADD COLUMN new_value TEXT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE records DROP COLUMN old_value;
ALTER TABLE records DROP COLUMN new_value;
-- +goose StatementEnd
