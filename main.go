package main

import (
	"money-management/infrastructure/db"
	"money-management/infrastructure/router"
	"money-management/interfaces/controllers"
	"money-management/interfaces/repositories"
	"money-management/usecases"
)

func main() {
	// Setup MongoDB
	database := db.NewMongoDatabase()

	userRepository := repositories.NewUserRepository(database)
	userInteractor := usecases.NewUserInteractor(userRepository)
	userController := controllers.NewUserController(userInteractor)

	walletRepository := repositories.NewWalletRepository(database)
	walletInteractor := usecases.NewWalletInteractor(walletRepository)
	walletController := controllers.NewWalletController(walletInteractor)

	transactionRepository := repositories.NewTransactionRepository(database)
	transactionInteractor := usecases.NewTransactionInteractor(transactionRepository, walletRepository)
	transactionController := controllers.NewTransactionController(transactionInteractor)

	// Setup Gin
	router := router.NewMuxRouter()
	router.POST("/api/user/register", userController.RegisterHandler)
	router.POST("/api/user/login", userController.LoginHandler)
	router.POST("/api/wallet", walletController.CreateWalletHandler)
	router.GET("/api/transactions/:limit/:wallet_id", transactionController.GetTransactionsByWalletIDHandler)    // Get Wallet's Transactions
	router.POST("/api/transaction", transactionController.CreateTransactionHandler)                              // Add Transaction
	router.PUT("/api/transaction", transactionController.EditTransactionHandler)                                 // Update Transaction
	router.DELETE("/api/transaction/:transaction_id/:wallet_id", transactionController.DeleteTransactionHandler) // Delete Transaction

	router.SERVE("8080")
}
