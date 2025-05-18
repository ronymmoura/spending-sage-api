package usecases

import (
	"github.com/gin-gonic/gin"
	db "github.com/ronymmoura/spending-sage-api/internal/db/sqlc"
)

func CreateMonthlyExpectedPaymentUseCase(ctx *gin.Context, store db.Store, name string, amount int32, day int16, userID int64) (monthlyExpectedPayment db.MonthlyExpectedPayment, err error) {
	arg := db.CreateMonthlyExpectedPaymentParams{
		Name:   name,
		Amount: amount,
		Day:    day,
		UserID: userID,
	}

	monthlyExpectedPayment, err = store.CreateMonthlyExpectedPayment(ctx, arg)
	return
}
