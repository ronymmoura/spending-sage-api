package usecases

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/ronymmoura/spending-sage-api/internal/db/sqlc"
)

func CreateMonthUseCase(ctx *gin.Context, store db.Store, date time.Time, user db.User) (month db.Month, err error) {

	fmt.Println(date.Location().String())

	firstDay := time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, date.Location())

	arg := db.CreateMonthParams{
		UserID: user.ID,
		Date:   firstDay,
	}

	month, err = store.CreateMonth(ctx, arg)

	return
}
