package usecases

import (
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/ronymmoura/spending-sage-api/internal/db/sqlc"
)

func CreateMonthUseCase(ctx *gin.Context, store *db.SQLStore, date time.Time, user db.User) (*db.Month, error) {
	var res db.Month

	err := store.ExecTx(ctx, func(q *db.Queries) error {
		firstDay := time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, date.Location())

		arg := db.CreateMonthParams{
			UserID: user.ID,
			Date:   firstDay,
		}

		month, err := q.CreateMonth(ctx, arg)

		if err != nil {
			return err
		}

		fixedEntries, err := q.SearchFixedEntries(ctx, db.SearchFixedEntriesParams{
			UserID: user.ID,
		})

		if err != nil {
			return err
		}

		for _, fixedEntry := range fixedEntries {
			_, err := q.CreateMonthEntry(ctx, db.CreateMonthEntryParams{
				MonthID:    month.ID,
				OriginID:   fixedEntry.OriginID,
				CategoryID: fixedEntry.CategoryID,
				Name:       fixedEntry.Name,
				DueDate:    fixedEntry.DueDate,
				Amount:     fixedEntry.Amount,
				Owner:      fixedEntry.Owner,
				PayDate:    time.Date(month.Date.Year(), month.Date.Month(), int(fixedEntry.PayDay), 0, 0, 0, 0, month.Date.Location()),
			})

			if err != nil {
				return err
			}
		}

		res = month

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &res, nil
}
