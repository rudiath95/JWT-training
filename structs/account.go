package structs

import (
	"math/rand"

	"github.com/google/uuid"
)

type Account struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Number    uuid.UUID `json:"number"`
	Balance   int64     `json:"balance"`
}

func NewAccount(firstName string, lastName string) *Account {
	return &Account{
		ID:        rand.Intn(100000),
		FirstName: firstName,
		LastName:  lastName,
		Number:    uuid.New(),
	}
}
