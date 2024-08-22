package usecases

import (
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/ronymmoura/spending-sage-api/internal/db/sqlc"
	"github.com/ronymmoura/spending-sage-api/internal/util"
)

func EditFixedEntryUseCase(ctx *gin.Context, store db.Store, user db.User, id int64, originID int64, categoryID int64, name string, dueDate time.Time, payDay int16, amount int32, owner string) (fixedEntry db.FixedEntry, err error) {
	foundEntry, err := store.GetFixedEntry(ctx, id)
	if err != nil {
		return
	}

	if foundEntry.UserID != user.ID {
		return foundEntry, util.ErrForbidenEntry
	}

	arg := db.EditFixedEntryParams{
		ID:         id,
		OriginID:   originID,
		CategoryID: categoryID,
		Name:       name,
		DueDate:    dueDate,
		PayDay:     payDay,
		Amount:     amount,
		Owner:      owner,
	}

	fixedEntry, err = store.EditFixedEntry(ctx, arg)
	return
}
