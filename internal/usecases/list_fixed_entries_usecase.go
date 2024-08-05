package usecases

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	db "github.com/ronymmoura/spending-sage-api/internal/db/sqlc"
)

func SearchFixedEntriesUseCase(ctx *gin.Context, store db.Store, userID int64, originID sql.NullInt64, categoryID sql.NullInt64, owner sql.NullString) (fixedEntries []db.FixedEntry, err error) {
	arg := db.SearchFixedEntriesParams{
		UserID:     userID,
		OriginID:   originID,
		CategoryID: categoryID,
		Owner:      owner,
	}

	fixedEntries, err = store.SearchFixedEntries(ctx, arg)
	return
}
