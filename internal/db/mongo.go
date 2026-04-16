package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongo(uri, dbName string) *mongo.Database {
	client, _ := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	return client.Database(dbName)
}