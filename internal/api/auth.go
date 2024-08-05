package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	auth "github.com/ronymmoura/spending-sage-api/internal/auth/clerk"
	"github.com/ronymmoura/spending-sage-api/internal/usecases"
)

func (server *Server) SignInRoute(ctx *gin.Context) {
	var event auth.ClerkEvent

	if err := ctx.ShouldBindBodyWithJSON(&event); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if event.Type == "user.created" {
		fullName := event.Data.FirstName + " " + event.Data.LastName
		usecases.CreateUserUseCase(ctx, server.Store, fullName, event.Data.EmailAddresses[0].Email)
	}
}
