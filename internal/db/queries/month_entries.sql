-- name: CreateMonthEntry :one
INSERT INTO month_entries (
  month_id,
  origin_id,
  category_id,
  name,
  due_date,
  pay_date,
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

-- name: GetMonthEntry :one
SELECT *
FROM month_entries
WHERE id = $1
LIMIT 1;

-- name: EditMonthEntry :one
UPDATE month_entries
SET name = $2,
    due_date = $3,
    pay_date = $4,
    amount = $5,
    owner = $6,
    origin_id = $7,
    category_id = $8
WHERE id = $1
RETURNING *;

-- name: PayEntry :one
UPDATE month_entries
SET paid_date = $2
WHERE id = $1
RETURNING *;

-- name: DeleteMonthEntry :exec
DELETE FROM month_entries
WHERE id = $1;