package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/yusrililhm/to-do-app/server/api"
)

func ShowAll(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	payload := api.Show()
	json.NewEncoder(w).Encode(payload)
}