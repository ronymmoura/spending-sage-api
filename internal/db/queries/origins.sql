-- name: CreateOrigin :one
INSERT INTO origins (
  name,
  type
)
VALUES (
  $1,
  $2
)
RETURNING *;

-- name: GetOrigin :one
SELECT *
FROM origins
WHERE id = $1
LIMIT 1;

-- name: ListOrigins :many
SELECT *
FROM origins;

-- name: DeleteOrigin :exec
DELETE FROM origins
WHERE id = $1;