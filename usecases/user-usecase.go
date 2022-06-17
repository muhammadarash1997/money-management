package usecases

import (
	"errors"
	"log"
	"money-management/domain"
	"money-management/dto"
	"money-management/helper"
	"money-management/interfaces/repositories"

	"golang.org/x/crypto/bcrypt"
)

type UserInteractor interface {
	SaveUser(dto.RegisterRequest) error
	Login(dto.LoginRequest) (dto.LoginResponse, error)
}

type userInteractor struct {
	userRepository   repositories.UserRepository
}

func NewUserInteractor(userRepository repositories.UserRepository) UserInteractor {
	return &userInteractor{userRepository}
}

func (this *userInteractor) SaveUser(userRegister dto.RegisterRequest) error {
	var user domain.User

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(userRegister.Password), bcrypt.MinCost)
	if err != nil {
		log.Printf("Error %v", err)
		return err
	}

	user.Name = userRegister.Name
	user.Email = userRegister.Email
	user.PasswordHash = string(passwordHash)

	err = this.userRepository.Save(user)
	if err != nil {
		log.Printf("Error %v", err)
		return err
	}

	return nil
}

func (this *userInteractor) Login(userLogin dto.LoginRequest) (dto.LoginResponse, error) {
	var loginResponse dto.LoginResponse
	var userDTO dto.UserDTO

	// Checking email and getting user
	email := userLogin.Email
	user, err := this.userRepository.GetByEmail(email)
	if err != nil {
		return loginResponse, err
	}

	// Checking password
	password := userLogin.Password
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		log.Printf("Error %v", err)
		err = errors.New("Wrong password")
		return loginResponse, err
	}

	// Generating token
	token, err := helper.GenerateToken(user.ID.Hex())
	if err != nil {
		return loginResponse, err
	}

	// Mapping User to UserDTO
	userDTO.ID = user.ID
	userDTO.Name = user.Name
	userDTO.Email = user.Email
	userDTO.Token = token

	// Mapping UserDTO to LoginResponse
	loginResponse.User = userDTO

	return loginResponse, err
}
