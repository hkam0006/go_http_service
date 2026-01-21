-- +goose Up
-- +goose StatementBegin
ALTER TABLE order_items ADD COLUMN price_in_cents INTEGER NOT NULL CHECK(price_in_cents >= 0);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE order_items DROP COLUMN price_in_cents;
-- +goose StatementEnd
