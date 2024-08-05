package usecases

import (
	"github.com/gin-gonic/gin"
	db "github.com/ronymmoura/spending-sage-api/internal/db/sqlc"
)

func ListCategoriesUseCase(ctx *gin.Context, store db.Store) (categories []db.Category, err error) {
	categories, err = store.ListCategories(ctx)

	return
}
