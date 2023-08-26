package cntechkitgomongodb

import (
	"github.com/cntech-io/cntechkit-go/utils"
)

type MongoDBEnv struct {
	Username         string
	Password         string
	Database         string
	ConnectionString string
}

func NewMongoDBEnv() *MongoDBEnv {
	return &MongoDBEnv{
		Username:         utils.GetStringEnv(string(MONGODB_USERNAME), false),
		Password:         utils.GetStringEnv(string(MONGODB_PASSWORD), false),
		Database:         utils.GetStringEnv(string(MONGODB_DATABASE), false),
		ConnectionString: utils.GetStringEnv(string(MONGODB_CONNECTION_STRING), false),
	}
}
