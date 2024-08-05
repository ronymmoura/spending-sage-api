package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/ronymmoura/spending-sage-api/internal/db/sqlc"
	"github.com/ronymmoura/spending-sage-api/internal/usecases"
)

type ListsResponse struct {
	Categories []db.Category `json:"categories"`
	Origins    []db.Origin   `json:"origins"`
}

func (server *Server) GetListsRoute(ctx *gin.Context) {
	categories, err := usecases.ListCategoriesUseCase(ctx, server.Store)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	origins, err := usecases.ListOriginsUseCase(ctx, server.Store)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	res := ListsResponse{
		Categories: categories,
		Origins:    origins,
	}

	ctx.JSON(http.StatusOK, res)
}
