package auth

import (
	"github.com/clerk/clerk-sdk-go/v2"
	"github.com/gin-gonic/gin"
	"github.com/ronymmoura/spending-sage-api/internal/util"
)

type ClerkEvent struct {
	Data ClerkUserData `json:"data"`
	Type string        `json:"type"`
}

type ClerkUserData struct {
	ID             string                        `json:"id"`
	Banned         bool                          `json:"banned"`
	FirstName      string                        `json:"first_name"`
	LastName       string                        `json:"last_name"`
	EmailAddresses []ClerkUserDataEmailAddresses `json:"email_addresses"`
}

type ClerkUserDataEmailAddresses struct {
	Email string `json:"email_address"`
}

func GetUser(ctx *gin.Context) *clerk.User {
	authPayload := ctx.MustGet(util.AuthorizationPayloadKey).(*clerk.User)
	return authPayload
}
