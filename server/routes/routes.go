package routes

import (
	"github.com/gorilla/mux"
	"github.com/yusrililhm/to-do-app/server/api"
	"github.com/yusrililhm/to-do-app/server/middleware"
)

func Routes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/show", middleware.ShowAll).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/insert", api.InsertOne).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/delete/{ID}", middleware.DeleteById).Methods("DELETE", "OPTIONS")
	return router
}