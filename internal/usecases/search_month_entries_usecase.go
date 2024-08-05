package usecases

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	db "github.com/ronymmoura/spending-sage-api/internal/db/sqlc"
)

func SearchMonthEntriesUseCase(ctx *gin.Context, store db.Store, page int32, monthID int64, originID sql.NullInt64, categoryID sql.NullInt64, owner sql.NullString) (monthEntries []db.MonthEntry, err error) {
	arg := db.SearchMonthEntriesParams{
		MonthID:    monthID,
		OriginID:   originID,
		CategoryID: categoryID,
		Owner:      owner,
	}

	monthEntries, err = store.SearchMonthEntries(ctx, arg)
	return
}
