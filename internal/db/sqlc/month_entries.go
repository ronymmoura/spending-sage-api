package db

import (
	"context"
	"database/sql"
)

const searchMonthEntries = `-- name: ListMonthEntries :many
SELECT id, month_id, name, due_date, pay_date, amount, owner, origin_id, category_id
FROM month_entries
WHERE month_id = $1
  AND (origin_id = $2 OR $2 IS NULL)
  AND (category_id = $3 OR $3 IS NULL)
  AND (owner = $4 OR $4 IS NULL)
ORDER BY due_date ASC
`

type SearchMonthEntriesParams struct {
	MonthID    int64          `json:"month_id"`
	OriginID   sql.NullInt64  `json:"origin_id"`
	CategoryID sql.NullInt64  `json:"category_id"`
	Owner      sql.NullString `json:"owner"`
}

func (q *Queries) SearchMonthEntries(ctx context.Context, arg SearchMonthEntriesParams) ([]MonthEntry, error) {
	rows, err := q.db.Query(ctx, searchMonthEntries,
		arg.MonthID,
		arg.OriginID,
		arg.CategoryID,
		arg.Owner,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []MonthEntry{}
	for rows.Next() {
		var i MonthEntry
		if err := rows.Scan(
			&i.ID,
			&i.MonthID,
			&i.Name,
			&i.DueDate,
			&i.PayDate,
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
