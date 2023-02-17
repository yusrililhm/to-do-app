package api

import (
	"context"
	"log"
	"net/http"

	"github.com/yusrililhm/to-do-app/server/database"
	"github.com/yusrililhm/to-do-app/server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertOne(w http.ResponseWriter, r *http.Request)  {
	db, err := database.ConnectMongo()
	if err != nil {
		log.Fatal(err.Error())
	}

	Todo := r.FormValue("Todo")

	_, err = db.Collection("todolist").InsertOne(context.Background(), models.ToDoList{Todo: Todo})
	if err != nil {
		log.Fatal(err.Error())
	}
	http.Redirect(w, r, "http://localhost:3000", http.StatusSeeOther)
}

func DeleteOne(Todo string)  {
	db, err := database.ConnectMongo()
	if err != nil {
		log.Fatal(err.Error())
	}
	id, _ := primitive.ObjectIDFromHex(Todo)

	_, err = db.Collection("todolist").DeleteMany(context.Background(), bson.M{"_id": id})
	if err != nil {
		log.Fatal(err.Error())
	}
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

	if err := doc.Err(); err != nil {
		log.Fatal(err.Error())
	}

	return results
}