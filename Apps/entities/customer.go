package entities

import(
	"go.mongodb.org/mongo-driver/bson/primitive"
)


type Customer struct {
	ID primitive.ObjectID `bson:"_id"`
	Name string `bson:"name"`
	Email string `bson:"email"`
	Password string `bson:"password"`
}