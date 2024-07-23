// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: fixed_entry_payment_history.sql

package db

import (
	"context"
	"time"
)

const createFixedEntryPaymentHistory = `-- name: CreateFixedEntryPaymentHistory :one
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
RETURNING id, entry_id, amount, date
`

type CreateFixedEntryPaymentHistoryParams struct {
	EntryID int64     `json:"entry_id"`
	Amount  int32     `json:"amount"`
	Date    time.Time `json:"date"`
}

func (q *Queries) CreateFixedEntryPaymentHistory(ctx context.Context, arg CreateFixedEntryPaymentHistoryParams) (FixedEntryPaymentHistory, error) {
	row := q.db.QueryRow(ctx, createFixedEntryPaymentHistory, arg.EntryID, arg.Amount, arg.Date)
	var i FixedEntryPaymentHistory
	err := row.Scan(
		&i.ID,
		&i.EntryID,
		&i.Amount,
		&i.Date,
	)
	return i, err
}

const deleteFixedEntryPaymentHistory = `-- name: DeleteFixedEntryPaymentHistory :exec
DELETE FROM fixed_entry_payment_history
WHERE id = $1
`

func (q *Queries) DeleteFixedEntryPaymentHistory(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteFixedEntryPaymentHistory, id)
	return err
}

const getFixedEntryPaymentHistory = `-- name: GetFixedEntryPaymentHistory :one
SELECT id, entry_id, amount, date
FROM fixed_entry_payment_history
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetFixedEntryPaymentHistory(ctx context.Context, id int64) (FixedEntryPaymentHistory, error) {
	row := q.db.QueryRow(ctx, getFixedEntryPaymentHistory, id)
	var i FixedEntryPaymentHistory
	err := row.Scan(
		&i.ID,
		&i.EntryID,
		&i.Amount,
		&i.Date,
	)
	return i, err
}

const listFixedEntryPaymentHistory = `-- name: ListFixedEntryPaymentHistory :many
SELECT id, entry_id, amount, date
FROM fixed_entry_payment_history
WHERE entry_id = $1
`

func (q *Queries) ListFixedEntryPaymentHistory(ctx context.Context, entryID int64) ([]FixedEntryPaymentHistory, error) {
	rows, err := q.db.Query(ctx, listFixedEntryPaymentHistory, entryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []FixedEntryPaymentHistory{}
	for rows.Next() {
		var i FixedEntryPaymentHistory
		if err := rows.Scan(
			&i.ID,
			&i.EntryID,
			&i.Amount,
			&i.Date,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}