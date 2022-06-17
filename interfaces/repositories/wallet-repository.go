package repositories

import (
	"context"
	"errors"
	"log"
	"money-management/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type WalletRepository interface {
	Save(domain.Wallet) (domain.Wallet, error)
}

type walletRepository struct {
	db *mongo.Database
}

func NewWalletRepository(db *mongo.Database) WalletRepository {
	return &walletRepository{db}
}

func (this *walletRepository) Save(wallet domain.Wallet) (domain.Wallet, error) {
	ctx := context.Background()

	result, err := this.db.Collection("wallets").InsertOne(ctx, &wallet)
	if err != nil {
		log.Printf("Error %v", err)
		err = errors.New("Failed to save wallet")
		return wallet, err
	}

	wallet.ID = result.InsertedID.(primitive.ObjectID)

	return wallet, nil
}