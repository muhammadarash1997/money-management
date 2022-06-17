package usecases

import (
	"money-management/domain"
	"money-management/dto"
	"money-management/interfaces/repositories"
)

type WalletInteractor interface {
	CreateWallet(wallet dto.CreateWalletRequest) (dto.CreateWalletResponse, error)
}

type walletInteractor struct {
	walletRepository repositories.WalletRepository
}

func NewWalletInteractor(walletRepository repositories.WalletRepository) WalletInteractor {
	return &walletInteractor{walletRepository}
}

func (this *walletInteractor) CreateWallet(walletRequest dto.CreateWalletRequest) (dto.CreateWalletResponse, error) {
	var wallet domain.Wallet
	var walletDTO dto.WalletDTO
	var walletResponse dto.CreateWalletResponse

	// Mapping CreateWalletRequest to Wallet
	wallet.UserID = walletRequest.UserID
	wallet.Name = walletRequest.Name
	wallet.Currency = walletRequest.Currency
	wallet.Amount = walletRequest.Amount

	// Saving wallet
	wallet, err := this.walletRepository.Save(wallet)
	if err != nil {
		return walletResponse, err
	}

	// Mapping Wallet to WalletDTO
	walletDTO.ID = wallet.ID
	walletDTO.Name = wallet.Name
	walletDTO.Currency = wallet.Currency
	walletDTO.Amount = wallet.Amount

	// Wrapping WalletDTO to CreateWalletResponse
	walletResponse.Wallet = walletDTO
	return walletResponse, nil
}
