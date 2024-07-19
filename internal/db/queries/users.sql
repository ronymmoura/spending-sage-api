-- name: CreateUser :one
INSERT INTO users (
  email,
  full_name
)
VALUES (
  $1,
  $2
)
RETURNING *;

-- name: GetUser :one
SELECT *
FROM users
WHERE id = $1
LIMIT 1;

-- name: GetUserByEmail :one
SELECT *
FROM users
WHERE email = $1
LIMIT 1;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;
