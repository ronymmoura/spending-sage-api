// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: fixed_entries.sql

package db

import (
	"context"
	"time"
)

const createFixedEntry = `-- name: CreateFixedEntry :one
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
RETURNING id, user_id, name, due_date, pay_day, amount, owner, origin_id, category_id
`

type CreateFixedEntryParams struct {
	UserID     int64     `json:"user_id"`
	OriginID   int64     `json:"origin_id"`
	CategoryID int64     `json:"category_id"`
	Name       string    `json:"name"`
	DueDate    time.Time `json:"due_date"`
	PayDay     time.Time `json:"pay_day"`
	Amount     int32     `json:"amount"`
	Owner      string    `json:"owner"`
}

func (q *Queries) CreateFixedEntry(ctx context.Context, arg CreateFixedEntryParams) (FixedEntry, error) {
	row := q.db.QueryRow(ctx, createFixedEntry,
		arg.UserID,
		arg.OriginID,
		arg.CategoryID,
		arg.Name,
		arg.DueDate,
		arg.PayDay,
		arg.Amount,
		arg.Owner,
	)
	var i FixedEntry
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Name,
		&i.DueDate,
		&i.PayDay,
		&i.Amount,
		&i.Owner,
		&i.OriginID,
		&i.CategoryID,
	)
	return i, err
}

const deleteFixedEntry = `-- name: DeleteFixedEntry :exec
DELETE FROM fixed_entries
WHERE id = $1
`

func (q *Queries) DeleteFixedEntry(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteFixedEntry, id)
	return err
}

const editFixedEntry = `-- name: EditFixedEntry :one
UPDATE fixed_entries
SET name = $2,
    due_date = $3,
    pay_day = $4,
    amount = $5,
    owner = $6,
    origin_id = $7,
    category_id = $8
WHERE id = $1
RETURNING id, user_id, name, due_date, pay_day, amount, owner, origin_id, category_id
`

type EditFixedEntryParams struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	DueDate    time.Time `json:"due_date"`
	PayDay     time.Time `json:"pay_day"`
	Amount     int32     `json:"amount"`
	Owner      string    `json:"owner"`
	OriginID   int64     `json:"origin_id"`
	CategoryID int64     `json:"category_id"`
}

func (q *Queries) EditFixedEntry(ctx context.Context, arg EditFixedEntryParams) (FixedEntry, error) {
	row := q.db.QueryRow(ctx, editFixedEntry,
		arg.ID,
		arg.Name,
		arg.DueDate,
		arg.PayDay,
		arg.Amount,
		arg.Owner,
		arg.OriginID,
		arg.CategoryID,
	)
	var i FixedEntry
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Name,
		&i.DueDate,
		&i.PayDay,
		&i.Amount,
		&i.Owner,
		&i.OriginID,
		&i.CategoryID,
	)
	return i, err
}

const getFixedEntry = `-- name: GetFixedEntry :one
SELECT id, user_id, name, due_date, pay_day, amount, owner, origin_id, category_id
FROM fixed_entries
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetFixedEntry(ctx context.Context, id int64) (FixedEntry, error) {
	row := q.db.QueryRow(ctx, getFixedEntry, id)
	var i FixedEntry
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Name,
		&i.DueDate,
		&i.PayDay,
		&i.Amount,
		&i.Owner,
		&i.OriginID,
		&i.CategoryID,
	)
	return i, err
}
