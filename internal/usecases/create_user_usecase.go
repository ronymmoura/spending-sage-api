package usecases

import (
	"github.com/gin-gonic/gin"
	db "github.com/ronymmoura/spending-sage-api/internal/db/sqlc"
)

func CreateUserUseCase(ctx *gin.Context, store db.Store, fullName string, email string) (user db.User, err error) {
	arg := db.CreateUserParams{
		Email:    email,
		FullName: fullName,
	}

	user, err = store.CreateUser(ctx, arg)
	if err != nil {
		return
	}

	return
}
