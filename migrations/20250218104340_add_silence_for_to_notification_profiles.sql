-- +goose Up
-- +goose StatementBegin
ALTER TABLE notification_profiles ADD COLUMN silence_for INTEGER;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE notification_profiles ADD COLUMN silence_for;
-- +goose StatementEnd
