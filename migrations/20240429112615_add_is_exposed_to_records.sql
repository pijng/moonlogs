-- +goose Up
-- +goose StatementBegin
ALTER TABLE records ADD COLUMN IS_EXPOSED INTEGER DEFAULT 0;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE records ADD COLUMN IS_EXPOSED;
-- +goose StatementEnd
