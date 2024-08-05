package usecases

import (
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/ronymmoura/spending-sage-api/internal/db/sqlc"
)

func CreateFixedEntryUseCase(ctx *gin.Context, store db.Store, userID int64, originID int64, categoryID int64, name string, dueDate time.Time, payDay int16, amount int32, owner string) (fixedEntry db.FixedEntry, err error) {
	arg := db.CreateFixedEntryParams{
		UserID:     userID,
		OriginID:   originID,
		CategoryID: categoryID,
		Name:       name,
		DueDate:    dueDate,
		PayDay:     payDay,
		Amount:     amount,
		Owner:      owner,
	}

	fixedEntry, err = store.CreateFixedEntry(ctx, arg)
	return
}
