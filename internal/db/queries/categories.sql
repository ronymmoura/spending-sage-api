-- name: CreateCategory :one
INSERT INTO categories (
  name
)
VALUES (
  $1
)
RETURNING *;

-- name: GetCategory :one
SELECT *
FROM categories
WHERE id = $1
LIMIT 1;

-- name: ListCategories :many
SELECT *
FROM categories
ORDER BY name ASC;

-- name: DeleteCategory :exec
DELETE FROM categories
WHERE id = $1;