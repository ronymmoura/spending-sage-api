package usecases

import (
	"github.com/gin-gonic/gin"
	db "github.com/ronymmoura/spending-sage-api/internal/db/sqlc"
)

func GetUserUseCase(ctx *gin.Context, store db.Store, email string) (user db.User, err error) {
	user, err = store.GetUserByEmail(ctx, email)

	return
}
