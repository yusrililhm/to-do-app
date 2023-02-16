package api

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/yusrililhm/to-do-app/server/database"
	"github.com/yusrililhm/to-do-app/server/models"
	"go.mongodb.org/mongo-driver/bson"
)

func InsertOne(todo models.ToDoList)  {
	db, err := database.ConnectMongo()
	if err != nil {
		log.Fatal(err.Error())
	}

	insert, err := db.Collection("todolist").InsertOne(context.Background(), todo)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Insert Successfull", insert.InsertedID)	
}

func DeleteOne(w http.ResponseWriter, r *http.Request)  {
	db, err := database.ConnectMongo()
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = db.Collection("todolist").DeleteOne(context.Background(), models.ToDoList{Todo: ""})
	if err != nil {
		log.Fatal(err.Error())
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func Update(w http.ResponseWriter, r *http.Request)  {
	db, err := database.ConnectMongo()
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = db.Collection("todolist").UpdateOne(context.Background(), bson.D{}, bson.D{})
	if err != nil {
		log.Fatal(err.Error())
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func Show() []models.ToDo {
	db, err := database.ConnectMongo()
	if err != nil {
		log.Fatal(err.Error())
	}

	var results []models.ToDo

	doc, err := db.Collection("todolist").Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err.Error())
	}

	for doc.Next(context.Background()) {
		var result models.ToDo
		err := doc.Decode(&result)
		if err != nil {
			log.Fatal(err.Error())
		}
		results = append(results, result)
	}

	if err := doc.Err() ;err != nil {
		log.Fatal(err.Error())
	}

	return results
}