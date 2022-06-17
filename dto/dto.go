package dto

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserDTO struct {
	ID       primitive.ObjectID `json:"id"`
	Name     string             `json:"name"`
	Email    string             `json:"email"`
	Token    string             `json:"token"`
}

type WalletDTO struct {
	ID       primitive.ObjectID `json:"id"`
	Name     string             `json:"name"`
	Currency string             `json:"currency"`
	Amount   int             `json:"amount"`
}

type TransactionDTO struct {
	ID       primitive.ObjectID `json:"id"`
	Amount   int                `json:"amount"`
	Category string             `json:"category"`
	Note     string             `json:"note"`
	Time     time.Time          `json:"time"`
}
