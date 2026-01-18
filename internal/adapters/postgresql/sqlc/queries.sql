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
