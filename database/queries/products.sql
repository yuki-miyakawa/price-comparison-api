-- name: GetAllProductsByUserID :many
SELECT * FROM products WHERE user_id = $1;

-- name: GetAllProductsByCategory :many
SELECT * FROM products WHERE user_id = $1 AND category = $2;

-- name: GetAllProductsByCategoryAndShops :many
SELECT * FROM products WHERE user_id = $1 AND category = $2 AND shops = $3;

-- name: GetAllProducts :many
SELECT * FROM products;

-- name: CreateProduct :one
INSERT INTO products (user_id, name, category, shops, price) VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: UpdateProduct :one
UPDATE products SET user_id = $2, name = $3, category = $4,shops = $5, price = $6 WHERE id = $1 RETURNING *;

-- name: DeleteProduct :one
DELETE FROM products WHERE id = $1 RETURNING *;