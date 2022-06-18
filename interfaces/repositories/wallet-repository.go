package repositories

import (
	"context"
	"errors"
	"log"
	"money-management/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type WalletRepository interface {
	Save(domain.Wallet) (domain.Wallet, error)
	GetByID(primitive.ObjectID) (domain.Wallet, error)
	Update(domain.Wallet) error
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

func (this *walletRepository) GetByID(id primitive.ObjectID) (domain.Wallet, error) {
	ctx := context.Background()

	var wallet domain.Wallet
	query := bson.M{"_id": id}
	err := this.db.Collection("wallets").FindOne(ctx, query).Decode(&wallet)
	if err != nil {
		log.Printf("Error %v", err)
		err = errors.New("Failed to get wallet")
		return wallet, err
	}

	return wallet, nil
}

func (this *walletRepository) Update(wallet domain.Wallet) error {
	ctx := context.Background()

	filter := bson.M{"wallet_id": wallet.ID}
	_, err := this.db.Collection("wallets").ReplaceOne(ctx, filter, wallet)
	if err != nil {
		log.Printf("Error %v", err)
		err = errors.New("Wallets not found")
		return err
	}

	return nil
}