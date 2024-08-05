package usecases

import (
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/ronymmoura/spending-sage-api/internal/db/sqlc"
)

func CreateMonthEntryUseCase(ctx *gin.Context, store db.Store, monthID int64, originID int64, categoryID int64, name string, dueDate time.Time, payDate time.Time, amount int32, owner string) (monthEntry db.MonthEntry, err error) {
	arg := db.CreateMonthEntryParams{
		MonthID:    monthID,
		OriginID:   originID,
		CategoryID: categoryID,
		Name:       name,
		DueDate:    dueDate,
		PayDate:    payDate,
		Amount:     amount,
		Owner:      owner,
	}

	monthEntry, err = store.CreateMonthEntry(ctx, arg)

	return
}
