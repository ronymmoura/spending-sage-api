package api

import (
	"net/http"

	_ "github.com/clerk/clerk-sdk-go/v2"
	_ "github.com/clerk/clerk-sdk-go/v2/jwt"
	_ "github.com/clerk/clerk-sdk-go/v2/user"
	"github.com/gin-gonic/gin"
)

func (server *Server) getUser(ctx *gin.Context) {
	authPayload := ctx.MustGet(authorizationPayloadKey)
	ctx.JSON(http.StatusOK, authPayload)
}
