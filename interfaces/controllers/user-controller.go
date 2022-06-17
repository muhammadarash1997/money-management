package controllers

import (
	"encoding/json"
	"errors"
	"log"
	"money-management/dto"
	"money-management/usecases"
	"net/http"
)

type UserController interface {
	RegisterHandler(w http.ResponseWriter, r *http.Request)
	LoginHandler(w http.ResponseWriter, r *http.Request)
}

type userController struct {
	userInteractor usecases.UserInteractor
}

func NewUserController(userInteractor usecases.UserInteractor) UserController {
	return &userController{userInteractor}
}

func (this *userController) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user dto.RegisterRequest

	err := json.NewDecoder(r.Body).Decode(&user)
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

	err = this.userInteractor.SaveUser(user)
	if err != nil {
		log.Printf("Error %v", err)
		err = errors.New("Register account failed")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(dto.Message{
			Code: http.StatusInternalServerError,
			Status: "Error",
			Data: err,
		})
		return
	}
	
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dto.Message{
		Code: http.StatusOK,
		Status: "Ok",
		Data: nil,
	})
	return
}

func (this *userController) LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user dto.LoginRequest

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Printf("Error %v", err)
		err = errors.New("Invalid payload")
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(err)
		return
	}

	loginResponse, err := this.userInteractor.Login(user)
	if err != nil {
		log.Printf("Error %v", err)
		err = errors.New("Login failed")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(dto.Message{
			Code: http.StatusInternalServerError,
			Status: "Error",
			Data: err,
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dto.Message{
		Code: http.StatusOK,
		Status: "Ok",
		Data: loginResponse,
	})
	return
}
