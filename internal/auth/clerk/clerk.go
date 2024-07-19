package auth

import "fmt"

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

func CreateUser(event ClerkEvent) {
	fmt.Printf("name: %s\n", event.Data.FirstName)
	fmt.Printf("email: %s\n", event.Data.EmailAddresses[0].Email)
}
