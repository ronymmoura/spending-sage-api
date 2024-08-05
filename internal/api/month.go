package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	db "github.com/ronymmoura/spending-sage-api/internal/db/sqlc"
	"github.com/ronymmoura/spending-sage-api/internal/usecases"
)

type CreateMonthRequest struct {
	Date time.Time `json:"date"`
}

func (server *Server) CreateMonthRoute(ctx *gin.Context) {
	user := GetUser(ctx, server.Store)

	var req CreateMonthRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	month, err := usecases.CreateMonthUseCase(ctx, server.Store, req.Date, user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, month)
}

type ListMonthsRequest struct {
	Page int32 `form:"page" binding:"required,min=1"`
}

type ListMonthsResponse struct {
	Items []db.Month `json:"items"`
	Total int64      `json:"total"`
	Page  int32      `json:"page"`
	Limit int32      `json:"limit"`
}

func (server *Server) ListMonthsRoute(ctx *gin.Context) {
	user := GetUser(ctx, server.Store)

	var req ListMonthsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	months, total, limit, err := usecases.ListMonthsUseCase(ctx, server.Store, user, req.Page)
	if err != nil {
		if err == pgx.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	res := ListMonthsResponse{
		Items: months,
		Total: total,
		Page:  req.Page,
		Limit: limit,
	}

	ctx.JSON(http.StatusOK, res)
}
