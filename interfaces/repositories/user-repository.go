package repositories

import (
	"context"
	"errors"
	"log"
	"money-management/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	Save(domain.User) error
	GetByEmail(email string) (domain.User, error)
}

type userRepository struct {
	db *mongo.Database
}

func NewUserRepository(database *mongo.Database) UserRepository {
	return &userRepository{database}
}

func (this *userRepository) Save(user domain.User) error {
	ctx := context.Background()

	_, err := this.GetByEmail(user.Email)
	if err == nil {
		err = errors.New("Email has been used")
		return err
	}

	_, err = this.db.Collection("users").InsertOne(ctx, &user)
	if err != nil {
		log.Printf("Error %v", err)
		return err
	}

	return nil
}

func (this *userRepository) GetByEmail(email string) (domain.User, error) {
	ctx := context.Background()

	var user domain.User
	query := bson.M{"email": email}
	err := this.db.Collection("users").FindOne(ctx, query).Decode(&user)
	if err != nil {
		log.Printf("Error %v", err)
		err = errors.New("Email not found")
		return user, err
	}

	return user, nil
}
