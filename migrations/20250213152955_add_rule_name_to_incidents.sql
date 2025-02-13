-- +goose Up
-- +goose StatementBegin
ALTER TABLE incidents ADD COLUMN rule_name TEXT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE incidents DROP COLUMN rule_name;
-- +goose StatementEnd
