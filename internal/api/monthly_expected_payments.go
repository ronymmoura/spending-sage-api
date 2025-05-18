package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/ronymmoura/spending-sage-api/internal/db/sqlc"
	"github.com/ronymmoura/spending-sage-api/internal/usecases"
)

type CreateMonthlyExpectedPaymentRequest struct {
	Name   string  `json:"name"`
	Amount float32 `json:"amount"`
	Day    int16   `json:"day"`
}

type MonthlyExpectedPaymentResponse struct {
	ID     int64   `json:"id"`
	Name   string  `json:"name"`
	Amount float32 `json:"amount"`
	Day    int16   `json:"day"`
}

func monthlyExpectedPaymentMapper(ep db.MonthlyExpectedPayment) (expectedPayment MonthlyExpectedPaymentResponse) {
	return MonthlyExpectedPaymentResponse{
		ID:     ep.ID,
		Name:   ep.Name,
		Day:    ep.Day,
		Amount: float32(ep.Amount) / 100,
	}
}

func monthlyExpectedPaymentsMapper(eps []db.MonthlyExpectedPayment) (expectedPayments []MonthlyExpectedPaymentResponse) {
	list := make([]MonthlyExpectedPaymentResponse, 0)

	for _, ep := range eps {
		list = append(list, monthlyExpectedPaymentMapper(ep))
	}

	return list
}

func (server *Server) CreateMonthlyExpectedPaymentRoute(ctx *gin.Context) {
	user := GetUser(ctx, server.Store)

	var req CreateMonthlyExpectedPaymentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var amount int32 = int32(req.Amount * 100)
	monthlyExpectedPayment, err := usecases.CreateMonthlyExpectedPaymentUseCase(ctx, server.Store, req.Name, amount, req.Day, user.ID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	res := monthlyExpectedPaymentMapper(monthlyExpectedPayment)

	ctx.JSON(http.StatusCreated, res)
}

type ListMonthlyExpectedPaymentResponse struct {
	Items []MonthlyExpectedPaymentResponse `json:"items"`
	Total int64                            `json:"total"`
	Page  int32                            `json:"page"`
	Limit int32                            `json:"limit"`
}

func (server *Server) ListMonthlyExpectedPaymentsRoute(ctx *gin.Context) {
	user := GetUser(ctx, server.Store)

	epsList, err := usecases.ListMonthlyExpectedPaymentsUseCase(ctx, server.Store, user.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	res := ListMonthlyExpectedPaymentResponse{
		Items: monthlyExpectedPaymentsMapper(epsList),
	}

	ctx.JSON(http.StatusOK, res)
}
