package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Transaction struct {
	ID       primitive.ObjectID `bson:"_id"`
	WalletID primitive.ObjectID `bson:"wallet_id"`
	Amount   int                `bson:"amount"`
	Category string             `bson:"category"`
	Note     string             `bson:"note"`
	Time     time.Time          `bson:"time"`
}
