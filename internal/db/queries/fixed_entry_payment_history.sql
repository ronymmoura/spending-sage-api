-- name: CreateFixedEntryPaymentHistory :one
INSERT INTO fixed_entry_payment_history (
  entry_id,
  amount,
  date
)
VALUES (
  $1,
  $2,
  $3
)
RETURNING *;

-- name: GetFixedEntryPaymentHistory :one
SELECT *
FROM fixed_entry_payment_history
WHERE id = $1
LIMIT 1;

-- name: ListFixedEntryPaymentHistory :many
SELECT *
FROM fixed_entry_payment_history
WHERE entry_id = $1;

-- name: DeleteFixedEntryPaymentHistory :exec
DELETE FROM fixed_entry_payment_history
WHERE id = $1;