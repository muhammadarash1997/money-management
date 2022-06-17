package usecases

import (
	"money-management/domain"
	"money-management/dto"
	"money-management/interfaces/repositories"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TransactionInteractor interface {
	CreateTransaction(dto.CreateTransactionRequest) error
	GetTransactionsByWalletIDHandler(int, primitive.ObjectID) ([]domain.Transaction, error)
	EditTransaction(dto.EditTransactionRequest) error
	DeleteTransaction(primitive.ObjectID) error
}

type transactionInteractor struct {
	transactionRepository repositories.TransactionRepository
}

func NewTransactionInteractor(transactionsRepository repositories.TransactionRepository) TransactionInteractor {
	return &transactionInteractor{transactionsRepository}
}

func (this *transactionInteractor) CreateTransaction(transaction dto.CreateTransactionRequest) error

func (this *transactionInteractor) GetTransactionsByWalletIDHandler(limit int, walletID primitive.ObjectID) ([]domain.Transaction, error)

func (this *transactionInteractor) EditTransaction(transaction dto.EditTransactionRequest) error

func (this *transactionInteractor) DeleteTransaction(id primitive.ObjectID) error