package usecases

import (
	"github.com/gin-gonic/gin"
	db "github.com/ronymmoura/spending-sage-api/internal/db/sqlc"
)

func ListMonthlyExpectedPaymentsUseCase(ctx *gin.Context, store db.Store, userID int64) (monthlyExpectedPayments []db.MonthlyExpectedPayment, err error) {
	monthlyExpectedPayments, err = store.ListMonthlyExpectedPayments(ctx, userID)
	return
}
