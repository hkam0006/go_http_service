-- +goose Up
-- +goose StatementBegin
ALTER TABLE users ADD COLUMN password VARCHAR(255);
ALTER TABLE users ADD COLUMN email VARCHAR(255);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users DROP COLUMN password;
ALTER TABLE users DROP COLUMN email;
-- +goose StatementEnd
