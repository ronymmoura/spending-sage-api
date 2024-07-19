package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	auth "github.com/ronymmoura/spending-sage-api/internal/auth/clerk"
)

func (server *Server) signIn(ctx *gin.Context) {
	var event auth.ClerkEvent

	if err := ctx.ShouldBindBodyWithJSON(&event); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if event.Type == "user.created" {
		auth.CreateUser(event)
	}
}
