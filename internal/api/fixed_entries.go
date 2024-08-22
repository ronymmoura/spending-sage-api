package api

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/ronymmoura/spending-sage-api/internal/db/sqlc"
	"github.com/ronymmoura/spending-sage-api/internal/usecases"
	"github.com/ronymmoura/spending-sage-api/internal/util"
)

type SearchFixedEntriesRequest struct {
	Page       int32          `form:"page" binding:"required,min=1"`
	CategoryID sql.NullInt64  `form:"category_id"`
	OriginID   sql.NullInt64  `form:"origin_id"`
	Owner      sql.NullString `form:"owner"`
}

type SearchFixedEntriesResponse struct {
	Items []FixedEntryResponse `json:"items"`
	Total int64                `json:"total"`
	Page  int32                `json:"page"`
	Limit int32                `json:"limit"`
}

type FixedEntryResponse struct {
	ID       int64       `json:"id"`
	Name     string      `json:"name"`
	DueDate  time.Time   `json:"due_date"`
	PaidDate *time.Time  `json:"paid_date"`
	PayDay   int16       `json:"pay_day"`
	Amount   float32     `json:"amount"`
	Owner    string      `json:"owner"`
	Origin   db.Origin   `json:"origin"`
	Category db.Category `json:"category"`
}

func fixedEntryMapper(ctx *gin.Context, server *Server, entries []db.FixedEntry) (fixedEntries []FixedEntryResponse, err error) {
	for _, entry := range entries {
		entryCategory, err := server.Cache.GetCategory(ctx, entry.CategoryID)
		if err != nil {
			return nil, err
		}

		entryOrigin, err := server.Cache.GetOrigin(ctx, entry.OriginID)
		if err != nil {
			return nil, err
		}

		fixedEntries = append(fixedEntries, FixedEntryResponse{
			ID:       entry.ID,
			Name:     entry.Name,
			DueDate:  entry.DueDate,
			PayDay:   entry.PayDay,
			Amount:   float32(entry.Amount) / 100,
			Owner:    entry.Owner,
			Origin:   entryOrigin,
			Category: entryCategory,
		})
	}

	return
}

func (server *Server) SearchFixedEntriesRoute(ctx *gin.Context) {
	user := GetUser(ctx, server.Store)

	var req SearchFixedEntriesRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	fixedEntries, err := usecases.SearchFixedEntriesUseCase(ctx, server.Store, user.ID, req.OriginID, req.CategoryID, req.Owner)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	items, err := fixedEntryMapper(ctx, server, fixedEntries)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	res := SearchFixedEntriesResponse{
		Items: items,
		Total: 1,
		Page:  1,
		Limit: 1,
	}

	ctx.JSON(http.StatusOK, res)
}

type CreateFixedEntryRequest struct {
	Name       string    `json:"name"`
	DueDate    time.Time `json:"due_date"`
	PayDay     int16     `json:"pay_day"`
	Amount     float32   `json:"amount"`
	Owner      string    `json:"owner"`
	CategoryID int64     `json:"category_id"`
	OriginID   int64     `json:"origin_id"`
}

type CreateFixedEntryResponse struct {
	ID         int64     `json:"id"`
	OriginID   int64     `json:"origin_id"`
	CategoryID int64     `json:"category_id"`
	Name       string    `json:"name"`
	DueDate    time.Time `json:"due_date"`
	PayDay     int16     `json:"pay_day"`
	Amount     int32     `json:"amount"`
	Owner      string    `json:"owner"`
}

func (server *Server) CreateFixedEntryRoute(ctx *gin.Context) {
	user := GetUser(ctx, server.Store)

	var req CreateFixedEntryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var amount int32 = int32(req.Amount * 100)

	fixedEntry, err := usecases.CreateFixedEntryUseCase(ctx, server.Store, user.ID, req.OriginID, req.CategoryID, req.Name, req.DueDate, req.PayDay, amount, req.Owner)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	res := CreateFixedEntryResponse{
		ID:         fixedEntry.ID,
		OriginID:   fixedEntry.OriginID,
		CategoryID: fixedEntry.CategoryID,
		Name:       fixedEntry.Name,
		DueDate:    fixedEntry.DueDate,
		PayDay:     fixedEntry.PayDay,
		Amount:     fixedEntry.Amount,
		Owner:      fixedEntry.Owner,
	}

	ctx.JSON(http.StatusCreated, res)
}

type UpdateFixedEntryRequest struct {
	ID         int64     `uri:"fixed_entry_id"`
	Name       string    `json:"name"`
	DueDate    time.Time `json:"due_date"`
	PayDay     int16     `json:"pay_day"`
	Amount     float32   `json:"amount"`
	Owner      string    `json:"owner"`
	CategoryID int64     `json:"category_id"`
	OriginID   int64     `json:"origin_id"`
}

type UpdateFixedEntryResponse struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	DueDate    time.Time `json:"due_date"`
	PayDay     int16     `json:"pay_day"`
	Amount     float32   `json:"amount"`
	Owner      string    `json:"owner"`
	CategoryID int64     `json:"category_id"`
	OriginID   int64     `json:"origin_id"`
}

func (server *Server) EditFixedEntryRoute(ctx *gin.Context) {
	user := GetUser(ctx, server.Store)

	var req UpdateFixedEntryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var amount int32 = int32(req.Amount * 100)

	fixedEntry, err := usecases.EditFixedEntryUseCase(ctx, server.Store, user, req.ID, req.OriginID, req.CategoryID, req.Name, req.DueDate, req.PayDay, amount, req.Owner)
	if err != nil {
		var status int

		switch err {
		case sql.ErrNoRows:
			status = http.StatusNotFound
		case util.ErrForbidenEntry:
			status = http.StatusForbidden
		default:
			status = http.StatusBadRequest
		}

		ctx.JSON(status, errorResponse(err))
		return
	}

	res := CreateFixedEntryResponse{
		ID:         fixedEntry.ID,
		OriginID:   fixedEntry.OriginID,
		CategoryID: fixedEntry.CategoryID,
		Name:       fixedEntry.Name,
		DueDate:    fixedEntry.DueDate,
		PayDay:     fixedEntry.PayDay,
		Amount:     fixedEntry.Amount,
		Owner:      fixedEntry.Owner,
	}

	ctx.JSON(http.StatusCreated, res)
}
