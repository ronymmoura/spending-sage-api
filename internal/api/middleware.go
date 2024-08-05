package api

import (
	"errors"
	"net/http"

	"github.com/clerk/clerk-sdk-go/v2/jwt"
	"github.com/clerk/clerk-sdk-go/v2/user"
	"github.com/gin-gonic/gin"
	"github.com/ronymmoura/spending-sage-api/internal/util"
)

func authMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationCookie, err := ctx.Request.Cookie(util.AuthorizationCookieKey)

		if err != nil {
			err := errors.New("authorization cookie is not provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		claims, err := jwt.Verify(ctx, &jwt.VerifyParams{
			Token: authorizationCookie.Value,
		})

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		usr, err := user.Get(ctx, claims.Subject)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
			return
		}

		ctx.Set(util.AuthorizationPayloadKey, usr)
		ctx.Next()
	}
}
