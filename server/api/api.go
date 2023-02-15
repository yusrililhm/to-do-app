package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/yusrililhm/to-do-app/server/database"
	"github.com/yusrililhm/to-do-app/server/models"
	"go.mongodb.org/mongo-driver/bson"
)

func InsertOne(w http.ResponseWriter, r *http.Request)  {
	db, err := database.ConnectMongo()
	if err != nil {
		log.Fatal(err.Error())
	}

	todo := r.FormValue("todo")

	_, err = db.Collection("todolist").InsertOne(context.Background(), models.ToDoList{Todo: todo})
	if err != nil {
		log.Fatal(err.Error())
	}
	
	http.Redirect(w, r, "/", http.StatusSeeOther)
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

func Show(w http.ResponseWriter, r *http.Request)  {
	db, err := database.ConnectMongo()
	if err != nil {
		log.Fatal(err.Error())
	}

	doc, err := db.Collection("todolist").Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err.Error())
	}

	var results []models.ToDo
	if err = doc.All(context.Background(), &results); err != nil {
		log.Fatal(err.Error())
	}

	for _, result := range results {
		doc.Decode(&result)
		output, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Fprintf(w, "%s\n", output)
	}
}