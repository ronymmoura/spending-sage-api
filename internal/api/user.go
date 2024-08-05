package api

import (
	"net/http"

	_ "github.com/clerk/clerk-sdk-go/v2"
	_ "github.com/clerk/clerk-sdk-go/v2/jwt"
	_ "github.com/clerk/clerk-sdk-go/v2/user"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	auth "github.com/ronymmoura/spending-sage-api/internal/auth/clerk"
	db "github.com/ronymmoura/spending-sage-api/internal/db/sqlc"
	"github.com/ronymmoura/spending-sage-api/internal/usecases"
)

func (server *Server) GetUserRoute(ctx *gin.Context) {
	user := GetUser(ctx, server.Store)
	ctx.JSON(http.StatusOK, user)
}

func GetUser(ctx *gin.Context, store db.Store) (user db.User) {
	clerkUser := auth.GetUser(ctx)

	user, err := usecases.GetUserUseCase(ctx, store, clerkUser.EmailAddresses[0].EmailAddress)
	if err != nil {
		if err == pgx.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	return
}
