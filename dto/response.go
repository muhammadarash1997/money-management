package dto

type LoginResponse struct {
	User UserDTO
}

type CreateWalletResponse struct {
	Wallet WalletDTO
}

type GetTransactionsByWalletIDResponse struct {
	Transactions []TransactionDTO
}