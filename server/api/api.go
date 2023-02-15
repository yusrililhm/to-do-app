package api

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/yusrililhm/to-do-app/server/database"
	"github.com/yusrililhm/to-do-app/server/models"
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
	fmt.Println(todo)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

