package api

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/ronymmoura/spending-sage-api/internal/db/sqlc"
	"github.com/ronymmoura/spending-sage-api/internal/usecases"
)

type CreateMonthEntryRequest struct {
	Name       string    `json:"name"`
	DueDate    time.Time `json:"due_date"`
	PayDate    time.Time `json:"pay_date"`
	Amount     float32   `json:"amount"`
	Owner      string    `json:"owner"`
	CategoryID int64     `json:"category_id"`
	OriginID   int64     `json:"origin_id"`
	MonthID    int64     `uri:"month_id"`
}

func (server *Server) CreateMonthEntryRoute(ctx *gin.Context) {
	var req CreateMonthEntryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var amount int32 = int32(req.Amount * 100)

	entry, err := usecases.CreateMonthEntryUseCase(ctx, server.Store, req.MonthID, req.OriginID, req.CategoryID, req.Name, req.DueDate, req.PayDate, amount, req.Owner)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, entry)
}

type SearchMonthEntriesRequest struct {
	MonthID    int64          `uri:"month_id"`
	Page       int32          `form:"page" binding:"required,min=1"`
	CategoryID sql.NullInt64  `form:"category_id"`
	OriginID   sql.NullInt64  `form:"origin_id"`
	Owner      sql.NullString `form:"owner"`
}

type SearchMonthEntriesResponse struct {
	Items []MonthEntryResponse `json:"items"`
	Total int64                `json:"total"`
	Page  int32                `json:"page"`
	Limit int32                `json:"limit"`
}

type MonthEntryResponse struct {
	ID       int64       `json:"id"`
	Name     string      `json:"name"`
	DueDate  time.Time   `json:"due_date"`
	PayDate  time.Time   `json:"pay_date"`
	PaidDate *time.Time  `json:"paid_date"`
	Amount   float32     `json:"amount"`
	Owner    string      `json:"owner"`
	Origin   db.Origin   `json:"origin"`
	Category db.Category `json:"category"`
}

func monthEntryMapper(ctx *gin.Context, server *Server, entries []db.MonthEntry) (monthEntries []MonthEntryResponse, err error) {
	for _, entry := range entries {
		entryCategory, err := server.Cache.GetCategory(ctx, entry.CategoryID)
		if err != nil {
			return nil, err
		}

		entryOrigin, err := server.Cache.GetOrigin(ctx, entry.OriginID)
		if err != nil {
			return nil, err
		}

		monthEntries = append(monthEntries, MonthEntryResponse{
			ID:       entry.ID,
			Name:     entry.Name,
			DueDate:  entry.DueDate,
			PayDate:  entry.PayDate,
			PaidDate: entry.PaidDate,
			Amount:   float32(entry.Amount) / 100,
			Owner:    entry.Owner,
			Origin:   entryOrigin,
			Category: entryCategory,
		})
	}

	return
}

func (server *Server) SearchMonthEntriesRoute(ctx *gin.Context) {
	var req SearchMonthEntriesRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	monthEntries, err := usecases.SearchMonthEntriesUseCase(ctx, server.Store, req.Page, req.MonthID, req.OriginID, req.CategoryID, req.Owner)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	items, err := monthEntryMapper(ctx, server, monthEntries)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	res := SearchMonthEntriesResponse{
		Items: items,
		Total: 0,
		Page:  1,
		Limit: 1,
	}

	ctx.JSON(http.StatusOK, res)
}
