package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/yusrililhm/to-do-app/server/api"
	"github.com/yusrililhm/to-do-app/server/models"
)

func ShowAll(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	payload := api.Show()
	json.NewEncoder(w).Encode(payload)
}

func Add(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Origin", "POST")
	w.Header().Set("Access-Control-Allow-Origin", "Content-Type")
	
	var todo models.ToDoList
	_ = json.NewDecoder(r.Body).Decode(&todo)
	api.InsertOne(todo)
	json.NewEncoder(w).Encode(&todo)
}