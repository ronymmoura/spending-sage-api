package usecases

import (
	"github.com/gin-gonic/gin"
	db "github.com/ronymmoura/spending-sage-api/internal/db/sqlc"
)

func ListMonthsUseCase(ctx *gin.Context, store db.Store, user db.User, page int32) (months []db.Month, total int64, limit int32, err error) {
	limit = 1

	total, err = store.CountMonths(ctx, user.ID)
	if err != nil {
		return
	}

	arg := db.ListMonthsParams{
		UserID: user.ID,
		Limit:  limit,
		Offset: (page - 1) * limit,
	}

	months, err = store.ListMonths(ctx, arg)

	return
}
