package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Wallets []Wallet

type Wallet struct {
	ID       primitive.ObjectID `bson:"_id"`
	UserID   primitive.ObjectID `bson:"user_id"`
	Name     string             `bson:"name"`
	Currency string             `bson:"currency"`
	Amount   int             `bson:"amount"`
}

func (this *Wallet) Transact(amount int) {
	this.Amount = this.Amount + amount
}