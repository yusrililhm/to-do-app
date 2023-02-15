package routes

import (
	"net/http"

	"github.com/yusrililhm/to-do-app/server/api"
)

func Routes()  {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "http://localhost:3000", http.StatusSeeOther)
	})

	http.HandleFunc("/Add", api.InsertOne)
	http.HandleFunc("/Delete", api.DeleteOne)
	http.HandleFunc("/Find", api.Show)
}