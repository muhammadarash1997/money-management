package dto

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateWalletRequest struct {
	UserID   primitive.ObjectID `json:"user_id"`
	Name     string             `json:"name"`
	Currency string             `json:"currency"`
	Amount   int                `json:"amount"`
}

type CreateTransactionRequest struct {
	WalletID primitive.ObjectID `json:"wallet_id"`
	Amount   int                `json:"amount"`
	Category string             `json:"category"`
	Note     string             `json:"note"`
	Time     time.Time          `json:"time"`
}

type EditTransactionRequest struct {
	ID       primitive.ObjectID `json:"wallet_id"`
	Amount   int                `json:"amount"`
	Category string             `json:"category"`
	Note     string             `json:"note"`
	Time     time.Time          `json:"time"`
}
