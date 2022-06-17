package controllers

import (
	"encoding/json"
	"errors"
	"log"
	"money-management/dto"
	"money-management/usecases"
	"net/http"
)

type WalletController interface {
	CreateWalletHandler(w http.ResponseWriter, r *http.Request)
}

type walletController struct {
	walletInteractor usecases.WalletInteractor
}

func NewWalletController(walletInteractor usecases.WalletInteractor) WalletController {
	return &walletController{walletInteractor}
}

func (this *walletController) CreateWalletHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var wallet dto.CreateWalletRequest

	err := json.NewDecoder(r.Body).Decode(&wallet)
	if err != nil {
		log.Printf("Error %v", err)
		err = errors.New("Invalid payload")
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(dto.Message{
			Code:   http.StatusUnprocessableEntity,
			Status: "Error",
			Data:   err,
		})
		return
	}

	createWalletResponse, err := this.walletInteractor.CreateWallet(wallet)
	if err != nil {
		log.Printf("Error %v", err)
		err = errors.New("Create wallet failed")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(dto.Message{
			Code:   http.StatusInternalServerError,
			Status: "Error",
			Data:   err,
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dto.Message{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   createWalletResponse,
	})
	return
}
