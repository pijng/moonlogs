-- +goose Up
-- +goose StatementBegin
ALTER TABLE records ADD COLUMN changes JSON;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE records DROP COLUMN changes;
-- +goose StatementEnd
