package models

import "go.mongodb.org/mongo-driver/bson/primitive"


type ToDoList struct {
	Todo string `bson:"todo"`
}

type ToDo struct {
	ID primitive.ObjectID `bson:"_id"`
	Todo string `bson:"todo"`
}

type User struct {
	Username string `bson:"username"`
	Password string `bson:"password"`
}