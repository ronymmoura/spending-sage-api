package db

import (
	"context"
	"database/sql"
)

const searchFixedEntries = `-- name: ListFixedEntries :many
SELECT id, user_id, name, due_date, pay_day, amount, owner, origin_id, category_id
FROM fixed_entries
WHERE user_id = $1
  AND (origin_id = $2 OR $2 IS NULL)
  AND (category_id = $3 OR $3 IS NULL)
  AND (owner = $4 OR $4 IS NULL)
ORDER BY due_date ASC
`

type SearchFixedEntriesParams struct {
	UserID     int64          `json:"user_id"`
	OriginID   sql.NullInt64  `json:"origin_id"`
	CategoryID sql.NullInt64  `json:"category_id"`
	Owner      sql.NullString `json:"owner"`
}

func (q *Queries) SearchFixedEntries(ctx context.Context, arg SearchFixedEntriesParams) ([]FixedEntry, error) {
	rows, err := q.db.Query(ctx, searchFixedEntries,
		arg.UserID,
		arg.OriginID,
		arg.CategoryID,
		arg.Owner,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []FixedEntry{}
	for rows.Next() {
		var i FixedEntry
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Name,
			&i.DueDate,
			&i.PayDay,
			&i.Amount,
			&i.Owner,
			&i.OriginID,
			&i.CategoryID,
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
