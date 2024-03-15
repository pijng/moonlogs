-- +goose Up
-- +goose StatementBegin
ALTER TABLE tags ADD COLUMN view_order INTEGER;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE tags DROP COLUMN view_order;
-- +goose StatementEnd
