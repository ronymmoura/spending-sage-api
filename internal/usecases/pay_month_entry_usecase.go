package usecases

import (
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/ronymmoura/spending-sage-api/internal/db/sqlc"
	"github.com/ronymmoura/spending-sage-api/internal/util"
)

func PayMonthEntryUseCase(ctx *gin.Context, store db.Store, user db.User, id int64, monthId int64, date *time.Time) (monthEntry db.MonthEntry, err error) {
	foundMonth, err := store.GetMonth(ctx, monthId)
	if err != nil {
		return
	}

	if foundMonth.UserID != user.ID {
		return db.MonthEntry{}, util.ErrForbidenEntry
	}

	foundEntry, err := store.GetMonthEntry(ctx, id)
	if err != nil {
		return
	}

	if foundEntry.MonthID != monthId {
		return db.MonthEntry{}, util.ErrForbidenEntry
	}

	arg := db.PayEntryParams{
		ID:       id,
		PaidDate: date,
	}

	monthEntry, err = store.PayEntry(ctx, arg)
	return
}
