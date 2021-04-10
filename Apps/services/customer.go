package services

import (
	"go.mongodb.org/mongo-driver/mongo"
	"fmt"
	"log"
	"context"
	database "github.com/customers/database"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/customers/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetCustomers() (customers []entities.Customer) {
	cur, err := database.GetCollection("customers").Find(context.TODO(), bson.D{})


	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		customer := entities.Customer{}
		err := cur.Decode(&customer)

		if err != nil {
			log.Fatal(err)
		}

		customers = append(customers,customer)
	}

	return
}

func CreateCustomer(customer entities.Customer) *mongo.InsertOneResult {
	customer.ID = primitive.NewObjectID()
	result, _ := database.GetCollection("customers").InsertOne(context.TODO(),customer)
	return result
}