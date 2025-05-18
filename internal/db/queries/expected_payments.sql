-- name: CreateExpectedPayment :one
INSERT INTO expected_payments (
  name,
  amount,
  date,
  month_id
)
VALUES (
  $1,
  $2,
  $3,
  $4
)
RETURNING *;

-- name: GetExpectedPayment :one
SELECT *
FROM expected_payments
WHERE id = $1
LIMIT $1;

-- name: ListExpectedPayments :many
SELECT *
FROM expected_payments
WHERE month_id = $1
ORDER BY date ASC;

-- name: EditExpectedPayment :one
UPDATE expected_payments
SET name = $2,
    amount = $3,
    date = $4
WHERE id = $1
RETURNING *;

-- name: DeleteExpectedPayment :exec
DELETE FROM expected_payments
WHERE id = $1;