package api

import (
	"context"
	"log"

	"github.com/yusrililhm/to-do-app/server/database"
	"github.com/yusrililhm/to-do-app/server/models"
)

func InsertOne()  {
	db, err := database.ConnectMongo()
	if err != nil {
		log.Fatal(err.Error())
	}

	todo := ""

	_, err = db.Collection("todolist").InsertOne(context.Background(), models.ToDoList{Todo: todo})
	if err != nil {
		log.Fatal(err.Error())
	}
}

