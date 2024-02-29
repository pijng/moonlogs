-- +goose Up
-- +goose StatementBegin
ALTER TABLE records ADD COLUMN request JSON;
ALTER TABLE records ADD COLUMN response JSON;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE records DROP COLUMN request;
ALTER TABLE records DROP COLUMN response;
-- +goose StatementEnd
