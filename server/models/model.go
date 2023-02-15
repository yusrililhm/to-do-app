package models

import "go.mongodb.org/mongo-driver/bson/primitive"


type ToDoList struct {
	Todo string `bson:"todo"`
}

type ToDo struct {
	ID primitive.ObjectID `bson:"_id"`
	ToDoList string `bson:"todo"`
}