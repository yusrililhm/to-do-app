package routes

import (
	"github.com/gorilla/mux"
	"github.com/yusrililhm/to-do-app/server/middleware"
)

func Routes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/show", middleware.ShowAll).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/insert", middleware.Add).Methods("POST", "OPTIONS")
	return router
}