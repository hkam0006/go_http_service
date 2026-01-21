-- name: ListProducts :many
SELECT
    *
FROM
    products;

-- name: FindProductsByID :one
SELECT
    *
FROM
    products
WHERE
    id = $1;

-- name: CreateProduct :one
INSERT INTO products (
    name,
    price_in_cents,
    quantity
) VALUES (
    $1,
    $2,
    $3
) RETURNING *;

-- name: AddUser :one
INSERT INTO users (
    first_name,
    last_name,
    email,
    password
) VALUES (
    $1,
    $2,
    $3,
    $4
) RETURNING *;

-- name: DeleteUser :one
DELETE FROM users WHERE id = $1 RETURNING id;

-- name: DeleteProduct :one
DELETE FROM products WHERE id=$1 RETURNING id;

-- name: PlaceOrder :one
INSERT INTO orders (user_id) VALUES ($1) RETURNING *;

-- name: CreateOrderItems :many
INSERT INTO order_items (order_id, product_id, quantity, price_per_product_in_cents)
SELECT
  $1::uuid,
  (x->>'product_id')::uuid,
  (x->>'quantity')::int,
  (x->>'price_per_product_in_cents')::int
FROM jsonb_array_elements($2::jsonb) AS x
RETURNING *;

-- name: GetProductsByIds :many
SELECT * FROM products WHERE id = ANY($1::uuid[]);

-- name: GetOrderById :one
SELECT
    o.*,
    COALESCE(
      jsonb_agg(
        jsonb_build_object(
          'product_id', oi.product_id,
          'quantity', oi.quantity,
          'product', p.name,
          'original_price', p.price_in_cents
        )
      ) FILTER (WHERE oi.id IS NOT NULL),
      '[]'::jsonb
    ) AS order_items,
    COALESCE(
        SUM(oi.quantity * oi.price_per_product_in_cents),
        0
    )::int AS total_price_in_cents
FROM
    orders o
LEFT JOIN
    order_items oi ON oi.order_id = o.id
LEFT JOIN
    products p ON oi.product_id = p.id
WHERE
    o.id=$1
GROUP BY
    o.id;
