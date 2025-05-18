package usecases

import (
	"github.com/gin-gonic/gin"
	db "github.com/ronymmoura/spending-sage-api/internal/db/sqlc"
	"github.com/ronymmoura/spending-sage-api/internal/util"
)

func EditMonthlyExpectedPaymentUseCase(ctx *gin.Context, store db.Store, id int64, name string, amount int32, day int16, userID int64) (monthlyExpectedPayment db.MonthlyExpectedPayment, err error) {
	foundEP, err := store.GetMonthlyExpectedPayment(ctx, id)
	if err != nil {
		return
	}

	if foundEP.UserID != userID {
		return foundEP, util.ErrForbidenEntry
	}

	arg := db.EditMonthlyExpectedPaymentParams{
		ID:     id,
		Name:   name,
		Amount: amount,
		Day:    day,
	}

	monthlyExpectedPayment, err = store.EditMonthlyExpectedPayment(ctx, arg)
	return
}
