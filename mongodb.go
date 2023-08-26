package cntechkitgomongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBEnvName string

const (
	MONGODB_USERNAME          MongoDBEnvName = "MONGODB_USERNAME"
	MONGODB_PASSWORD          MongoDBEnvName = "MONGODB_PASSWORD"
	MONGODB_DATABASE          MongoDBEnvName = "MONGODB_DATABASE"
	MONGODB_CONNECTION_STRING MongoDBEnvName = "MONGODB_CONNECTION_STRING"
)

type MongoDBKit struct {
	Context     context.Context
	Client      *mongo.Client
	Collections map[string]*mongo.Collection
}

func NewMongoDB() *MongoDBKit {
	return &MongoDBKit{
		Context: context.Background(),
	}
}

func (mdb *MongoDBKit) Connect() *MongoDBKit {

	env := NewMongoDBEnv()

	if env.ConnectionString == "" {
		panic("MongoDB connection string is empty!")
	}

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(env.ConnectionString).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(mdb.Context, opts)
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to MongoDB: %v", err))
	}

	if err := client.Database(env.Database).RunCommand(mdb.Context, bson.D{{Key: "ping", Value: 1}}).Err(); err != nil {
		panic(fmt.Sprintf("Failed to ping to MongoDB: %v", err))
	}

	mdb.Client = client
	return mdb
}
