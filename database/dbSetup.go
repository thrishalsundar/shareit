package database

import (
	"context"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DbSetup() (*mongo.Client, context.Context) {
	var client *mongo.Client
	var ctx context.Context

	if err := godotenv.Load(); err != nil {
		return nil, nil
	}

	uri := os.Getenv("MONGOURI")
	ctx = context.TODO()
	opts := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, opts)

	if err != nil {
		panic(err)
	}

	return client, ctx
}
