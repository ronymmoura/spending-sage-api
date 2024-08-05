package usecases

import (
	"github.com/gin-gonic/gin"
	db "github.com/ronymmoura/spending-sage-api/internal/db/sqlc"
)

func ListOriginsUseCase(ctx *gin.Context, store db.Store) (origins []db.Origin, err error) {
	origins, err = store.ListOrigins(ctx)

	return
}
