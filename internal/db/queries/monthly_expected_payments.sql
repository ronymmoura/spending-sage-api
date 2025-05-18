-- name: CreateMonthlyExpectedPayment :one
INSERT INTO monthly_expected_payments (
  name,
  amount,
  day,
  user_id
)
VALUES (
  $1,
  $2,
  $3,
  $4
)
RETURNING *;

-- name: GetMonthlyExpectedPayment :one
SELECT *
FROM monthly_expected_payments
WHERE id = $1
LIMIT 1;

-- name: ListMonthlyExpectedPayments :many
SELECT *
FROM monthly_expected_payments
WHERE user_id = $1
ORDER BY day ASC;

-- name: EditMonthlyExpectedPayment :one
UPDATE monthly_expected_payments
SET name = $2,
    amount = $3,
    day = $4
WHERE id = $1
RETURNING *;

-- name: DeleteMonthlyExpectedPayment :exec
DELETE FROM monthly_expected_payments
WHERE id = $1;