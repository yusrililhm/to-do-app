package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yusrililhm/to-do-app/server/api"
)

func ShowAll(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	payload := api.Show()
	json.NewEncoder(w).Encode(payload)
}

func DeleteById(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	params := mux.Vars(r)
	api.DeleteOne(params["ID"])
	json.NewEncoder(w).Encode(params["ID"])
}