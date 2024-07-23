-- name: CreateFixedEntry :one
INSERT INTO fixed_entries (
  user_id,
  origin_id,
  category_id,
  name,
  due_date,
  pay_day,
  amount,
  owner
)
VALUES (
  $1,
  $2,
  $3,
  $4,
  $5,
  $6,
  $7,
  $8
)
RETURNING *;

-- name: GetFixedEntry :one
SELECT *
FROM fixed_entries
WHERE id = $1
LIMIT 1;

-- name: EditFixedEntry :one
UPDATE fixed_entries
SET name = $2,
    due_date = $3,
    pay_day = $4,
    amount = $5,
    owner = $6,
    origin_id = $7,
    category_id = $8
WHERE id = $1
RETURNING *;

-- name: DeleteFixedEntry :exec
DELETE FROM fixed_entries
WHERE id = $1;