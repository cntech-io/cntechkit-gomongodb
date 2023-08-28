package cntechkitgomongodb

import (
	"context"
	"fmt"
	"time"

	gokit "github.com/cntech-io/cntechkit-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	context      context.Context
	Client       *mongo.Client
	Collections  map[string]*mongo.Collection
	enableLogger bool
}

type logs struct {
	ID          primitive.ObjectID `json:"_id"`
	AppName     string             `json:"app_name"`
	CreatedAt   time.Time          `json:"created_at"`
	Description string             `json:"description"`
}

var env = NewMongoDBEnv()

func NewMongoDB(enableLogger bool) *MongoDBKit {
	return &MongoDBKit{
		context:      context.Background(),
		enableLogger: enableLogger,
	}
}

func (mdb *MongoDBKit) Connect() *MongoDBKit {

	if env.ConnectionString == "" {
		panic("MongoDB connection string is empty!")
	}

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(env.ConnectionString).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(mdb.context, opts)
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to MongoDB: %v", err))
	}

	if err := client.Database(env.Database).RunCommand(mdb.context, bson.D{{Key: "ping", Value: 1}}).Err(); err != nil {
		panic(fmt.Sprintf("Failed to ping to MongoDB: %v", err))
	}

	gokit.NewLogger(&gokit.LoggerConfig{
		AppName: "cntechkit-gomongodb",
	}).Info("Connected to MongoDB")

	mdb.Client = client

	if mdb.enableLogger {
		newCollectionMap := map[string]*mongo.Collection{
			"logs": mdb.Client.Database(env.Database).Collection("logs"),
		}
		mdb.Collections = newCollectionMap
	}

	return mdb
}

func (mdb *MongoDBKit) AttachCollection(collectionName string) *MongoDBKit {

	collection := mdb.Client.Database(env.Database).Collection(collectionName)
	if mdb.Collections == nil {
		newCollectionMap := map[string]*mongo.Collection{
			collectionName: collection,
		}
		mdb.Collections = newCollectionMap
	} else {
		mdb.Collections[collectionName] = collection
	}
	return mdb
}

func (mdb *MongoDBKit) Disconnect() {
	err := mdb.Client.Disconnect(mdb.context)
	if err != nil {
		panic(fmt.Sprintf("Failed to disconnect from MongoDB: %v", err))
	}
}

func (mdb *MongoDBKit) Do(collectionName string) *mongo.Collection {
	return mdb.Collections[collectionName]
}

func (mdb *MongoDBKit) PushLog(appName string, description string) {
	if mdb.Collections["logs"] == nil {
		gokit.NewLogger(&gokit.LoggerConfig{
			AppName: "cntechkit-gomongodb",
		}).Info("Mongodb logger is not configured!")
		return
	}
	log := &logs{
		ID:          primitive.NewObjectID(),
		AppName:     appName,
		CreatedAt:   time.Now(),
		Description: description,
	}

	mdb.Collections["logs"].InsertOne(mdb.context, log)
}
