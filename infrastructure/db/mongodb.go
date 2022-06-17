package db

import (
	"context"
	"os"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDatabase() *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoPoolMin, err := strconv.Atoi(os.Getenv("MONGO_POOL_MIN"))
	if err != nil {
		panic(err)
	}

	mongoPoolMax, err := strconv.Atoi(os.Getenv("MONGO_POOL_MAX"))
	if err != nil {
		panic(err)
	}

	mongoMaxIdleTime, err := strconv.Atoi(os.Getenv("MONGO_MAX_IDLE_TIME_SECOND"))
	if err != nil {
		panic(err)
	}

	option := options.Client().
		ApplyURI(os.Getenv("MONGO_URI")).
		SetMinPoolSize(uint64(mongoPoolMin)).
		SetMaxPoolSize(uint64(mongoPoolMax)).
		SetMaxConnIdleTime(time.Duration(mongoMaxIdleTime) * time.Second)

	client, err := mongo.NewClient(option)
	if err != nil {
		panic(err)
	}

	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}

	database := client.Database(os.Getenv("MONGO_DATABASE"))
	return database
}
