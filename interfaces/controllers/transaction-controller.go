package controllers

import (
	"encoding/json"
	"errors"
	"log"
	"money-management/dto"
	"money-management/usecases"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TransactionController interface {
	CreateTransactionHandler(w http.ResponseWriter, r *http.Request)
	GetTransactionsByWalletIDHandler(w http.ResponseWriter, r *http.Request)
	EditTransactionHandler(w http.ResponseWriter, r *http.Request)
	DeleteTransactionHandler(w http.ResponseWriter, r *http.Request)
}

type transactionController struct {
	transactionInteractor usecases.TransactionInteractor
}

func NewTransactionController(transactionInteractor usecases.TransactionInteractor) TransactionController {
	return &transactionController{transactionInteractor}
}

func (this *transactionController) CreateTransactionHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var transaction dto.CreateTransactionRequest

	err := json.NewDecoder(r.Body).Decode(&transaction)
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

	err = this.transactionInteractor.CreateTransaction(transaction)
	if err != nil {
		log.Printf("Error %v", err)
		err = errors.New("Create transaction failed")
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
		Data:   nil,
	})
	return
}

func (this *transactionController) GetTransactionsByWalletIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	limit, err:= strconv.Atoi(params["limit"])
	if err != nil {
		log.Printf("Error %v", err)
		err = errors.New("Invalid payload")
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(dto.Message{
			Code: http.StatusUnprocessableEntity,
			Status: "Error",
			Data: err,
		})
		return
	}

	walletID, err := primitive.ObjectIDFromHex(params["wallet_id"])
	if err != nil {
		log.Printf("Error %v", err)
		err = errors.New("Invalid payload")
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(dto.Message{
			Code: http.StatusUnprocessableEntity,
			Status: "Error",
			Data: err,
		})
		return
	}

	transactions, err := this.transactionInteractor.GetTransactionsByWalletID(limit, walletID)
	if err != nil {
		log.Printf("Error %v", err)
		err = errors.New("Login failed")
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
		Data:   transactions,
	})
	return
}

func (this *transactionController) EditTransactionHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var transaction dto.EditTransactionRequest

	err := json.NewDecoder(r.Body).Decode(&transaction)
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

	err = this.transactionInteractor.EditTransaction(transaction)
	if err != nil {
		log.Printf("Error %v", err)
		err = errors.New("Edit transaction failed")
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
		Data:   nil,
	})
	return
}

func (this *transactionController) DeleteTransactionHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	transactionID, err := primitive.ObjectIDFromHex(params["transaction_id"])
	if err != nil {
		log.Printf("Error %v", err)
		err = errors.New("Invalid payload")
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(dto.Message{
			Code: http.StatusUnprocessableEntity,
			Status: "Error",
			Data: err,
		})
		return
	}
	walletID, err := primitive.ObjectIDFromHex(params["wallet_id"])
	if err != nil {
		log.Printf("Error %v", err)
		err = errors.New("Invalid payload")
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(dto.Message{
			Code: http.StatusUnprocessableEntity,
			Status: "Error",
			Data: err,
		})
		return
	}

	err = this.transactionInteractor.DeleteTransaction(transactionID, walletID)
	if err != nil {
		log.Printf("Error %v", err)
		err = errors.New("Delete transaction failed")
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
		Data:   nil,
	})
	return
}