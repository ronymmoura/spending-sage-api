-- name: CreateMonth :one
INSERT INTO months (
  user_id,
  date
)
VALUES (
  $1,
  $2
)
RETURNING *;

-- name: GetMonth :one
SELECT *
FROM months
WHERE id = $1
LIMIT 1;

-- name: ListMonths :many
SELECT *
FROM months
WHERE user_id = $1
ORDER BY date DESC
LIMIT $2
OFFSET $3;

-- name: CountMonths :one
SELECT COUNT(*)
FROM months
WHERE user_id = $1;

-- name: DeleteMonth :exec
DELETE FROM months
WHERE id = $1;