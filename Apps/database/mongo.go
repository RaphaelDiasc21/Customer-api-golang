package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetCollection(collection string) *mongo.Collection {
	client, _ := mongo.Connect(context.TODO(),options.Client().ApplyURI("mongodb://172.17.0.2:27017"))
	return client.Database("customers").Collection(collection)
}