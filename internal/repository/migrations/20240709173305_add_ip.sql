-- +goose Up
-- +goose StatementBegin
ALTER TABLE sessions ADD COLUMN ip text;
UPDATE sessions SET ip = '0.0.0.0';
ALTER TABLE sessions ALTER COLUMN ip SET NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE sessions DROP COLUMN ip;
-- +goose StatementEnd
