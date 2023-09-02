package cntechkitgomongodb

import (
	"fmt"

	"github.com/cntech-io/cntechkit-go/utils"
	"github.com/joho/godotenv"
)

type MongoDBEnv struct {
	Username         string
	Password         string
	Database         string
	ConnectionString string
}

func NewMongoDBEnv() *MongoDBEnv {
	if err := godotenv.Load(); err != nil {
		fmt.Println(".env file not found")
	}
	return &MongoDBEnv{
		Username:         utils.GetStringEnv(string(MONGODB_USERNAME), false),
		Password:         utils.GetStringEnv(string(MONGODB_PASSWORD), false),
		Database:         utils.GetStringEnv(string(MONGODB_DATABASE), false),
		ConnectionString: utils.GetStringEnv(string(MONGODB_CONNECTION_STRING), false),
	}
}
