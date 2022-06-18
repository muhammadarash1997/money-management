package usecases

import (
	"money-management/domain"
	"money-management/dto"
	"money-management/interfaces/repositories"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TransactionInteractor interface {
	CreateTransaction(dto.CreateTransactionRequest) error
	GetTransactionsByWalletID(int, primitive.ObjectID) ([]dto.TransactionDTO, error)
	EditTransaction(dto.EditTransactionRequest) error
	DeleteTransaction(primitive.ObjectID, primitive.ObjectID) error
}

type transactionInteractor struct {
	transactionRepository repositories.TransactionRepository
	walletRepository repositories.WalletRepository
}

func NewTransactionInteractor(transactionsRepository repositories.TransactionRepository, walletRepository repositories.WalletRepository) TransactionInteractor {
	return &transactionInteractor{transactionsRepository, walletRepository}
}

func (this *transactionInteractor) CreateTransaction(transactionRequest dto.CreateTransactionRequest) error {
	transaction := domain.Transaction{}

	// Mapping CreateTransactionRequest to Transaction
	transaction.WalletID = transactionRequest.WalletID
	transaction.Amount = transactionRequest.Amount
	transaction.Category = transactionRequest.Category
	transaction.Note = transactionRequest.Note
	transaction.Time = transactionRequest.Time

	// Getting Wallet
	wallet, err := this.walletRepository.GetByID(transaction.WalletID)
	if err != nil {
		return err
	}

	// Doing wallet transaction
	wallet.Transact(transaction.Amount)

	// Saving Wallet
	err = this.walletRepository.Update(wallet)
	if err != nil {
		return err
	}

	// Saving Transaction
	err = this.transactionRepository.Save(transaction)
	if err != nil {
		return err
	}

	return nil
}

func (this *transactionInteractor) GetTransactionsByWalletID(limit int, walletID primitive.ObjectID) ([]dto.TransactionDTO, error) {
	var transactionsDTO []dto.TransactionDTO

	// Getting Transactions
	transactions, err := this.transactionRepository.GetByWalletID(limit, walletID)
	if err != nil {
		return nil, err
	}

	// Mapping Transaction to TransactionDTO
	for _, transaction := range transactions {
		transactionDTO := dto.TransactionDTO{
			ID: transaction.ID,
			Amount: transaction.Amount,
			Category: transaction.Category,
			Note: transaction.Note,
			Time: transaction.Time,
		}

		transactionsDTO = append(transactionsDTO, transactionDTO)
	}

	return transactionsDTO, nil
}

func (this *transactionInteractor) EditTransaction(transactionRequest dto.EditTransactionRequest) error {
	transaction := domain.Transaction{}

	// Mapping CreateTransactionRequest to Transaction
	transaction.ID = transactionRequest.ID
	transaction.WalletID = transactionRequest.WalletID
	transaction.Amount = transactionRequest.Amount
	transaction.Category = transactionRequest.Category
	transaction.Note = transactionRequest.Note
	transaction.Time = transactionRequest.Time

	// Getting Transaction
	previousTransaction, err := this.transactionRepository.GetByID(transaction.ID)
	if err != nil {
		return err
	}

	// Getting deviation for updating amount of wallet
	deviation := transaction.Amount - previousTransaction.Amount	// New - Previous

	// Getting Wallet
	wallet, err := this.walletRepository.GetByID(transaction.WalletID)
	if err != nil {
		return err
	}

	// Update amount of wallet
	wallet.Transact(deviation)

	// Updating Wallet
	err = this.walletRepository.Update(wallet)
	if err != nil {
		return err
	}

	// Editing Transaction
	err = this.transactionRepository.Edit(transaction)
	if err != nil {
		return err
	}

	return nil
}

func (this *transactionInteractor) DeleteTransaction(transactionID primitive.ObjectID, walletID primitive.ObjectID) error {
	// Getting Transaction
	transaction, err := this.transactionRepository.GetByID(transactionID)

	// Delete Transaction
	err = this.transactionRepository.Delete(transactionID)
	if err != nil {
		return err
	}

	// Getting Wallet
	wallet, err := this.walletRepository.GetByID(walletID)
	if err != nil {
		return err
	}

	// Update amount of Wallet
	wallet.Transact(-1 * transaction.Amount)

	// Update Wallet
	err = this.walletRepository.Update(wallet)
	if err != nil {
		return err
	}

	return nil
}
