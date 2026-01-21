
-- +goose Up
-- +goose StatementBegin
ALTER TABLE order_items RENAME COLUMN price_in_cents TO price_per_product_in_cents;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE order_items RENAME COLUMN price_per_product_in_cents TO price_in_cents;
-- +goose StatementEnd
