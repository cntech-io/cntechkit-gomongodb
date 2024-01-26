package mongodb

import (
	"fmt"

	e "github.com/cntech-io/cntechkit-go/v2/env"
	"github.com/joho/godotenv"
)

const (
	MONGODB_USERNAME          MongoDBEnvName = "MONGODB_USERNAME"
	MONGODB_PASSWORD          MongoDBEnvName = "MONGODB_PASSWORD"
	MONGODB_DATABASE          MongoDBEnvName = "MONGODB_DATABASE"
	MONGODB_CONNECTION_STRING MongoDBEnvName = "MONGODB_CONNECTION_STRING"
)

type MongoDBEnv struct {
	Username         string
	Password         string
	Database         string
	ConnectionString string
}

type MongoDBEnvName string

func NewMongoDBEnv() *MongoDBEnv {
	if err := godotenv.Load(); err != nil {
		fmt.Println(".env file not found")
	}
	return &MongoDBEnv{
		Username:         e.GetString(string(MONGODB_USERNAME), false),
		Password:         e.GetString(string(MONGODB_PASSWORD), false),
		Database:         e.GetString(string(MONGODB_DATABASE), false),
		ConnectionString: e.GetString(string(MONGODB_CONNECTION_STRING), false),
	}
}
