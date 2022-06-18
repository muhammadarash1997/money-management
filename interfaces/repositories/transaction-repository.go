package repositories

import (
	"context"
	"errors"
	"log"
	"money-management/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TransactionRepository interface {
	Save(domain.Transaction) error
	GetByID(primitive.ObjectID) (domain.Transaction, error)
	GetByWalletID(int, primitive.ObjectID) ([]domain.Transaction, error)
	Edit(domain.Transaction) error
	Delete(primitive.ObjectID) error
}

type transactionRepository struct {
	db *mongo.Database
}

func NewTransactionRepository(database *mongo.Database) TransactionRepository {
	return &transactionRepository{database}
}

func (this *transactionRepository) Save(transaction domain.Transaction) error {
	ctx := context.Background()

	_, err := this.db.Collection("transactions").InsertOne(ctx, &transaction)
	if err != nil {
		log.Printf("Error %v", err)
		return err
	}

	return nil
}

func (this *transactionRepository) GetByID(id primitive.ObjectID) (domain.Transaction, error) {
	ctx := context.Background()

	var transaction domain.Transaction
	filter := bson.M{"_id": id}
	err := this.db.Collection("transactions").FindOne(ctx, filter).Decode(&transaction)
	if err != nil {
		log.Printf("Error %v", err)
		err = errors.New("Transaction not found")
		return transaction, err
	}

	return transaction, nil
}

func (this *transactionRepository) GetByWalletID(limit int, walletID primitive.ObjectID) ([]domain.Transaction, error) {
	ctx := context.Background()

	options := options.Find()
	options.SetLimit(int64(limit))
	options.SetSort(bson.D{{"time", -1}})

	filter := bson.M{"wallet_id": walletID}
	cursor, err := this.db.Collection("transactions").Find(ctx, filter)
	if err != nil {
		log.Printf("Error %v", err)
		err = errors.New("Transactions not found")
		return nil, err
	}

	var transactions []domain.Transaction
	err = cursor.All(ctx, &transactions)
	if err != nil {
		log.Printf("Error %v", err)
		err = errors.New("Transactions not found")
		return nil, err
	}

	return transactions, nil
}

func (this *transactionRepository) Edit(transaction domain.Transaction) error {
	ctx := context.Background()

	filter := bson.M{"transaction_id": transaction.ID}
	_, err := this.db.Collection("transactions").ReplaceOne(ctx, filter, transaction)
	if err != nil {
		log.Printf("Error %v", err)
		err = errors.New("Transactions not found")
		return err
	}

	return nil
}

func (this *transactionRepository) Delete(id primitive.ObjectID) error {
	ctx := context.Background()

	filter := bson.M{"_id": id}
	_, err := this.db.Collection("transactions").DeleteOne(ctx, filter)
	if err != nil {
		log.Printf("Error %v", err)
		err = errors.New("Transactions not found")
		return err
	}

	return nil
}