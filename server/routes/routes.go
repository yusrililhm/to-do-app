package routes

import (
	"net/http"

	"github.com/yusrililhm/to-do-app/server/api"
	"github.com/yusrililhm/to-do-app/server/auth"
)

func Routes()  {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "http://localhost:3000", http.StatusSeeOther)
	})

	// api
	http.HandleFunc("/add", api.InsertOne)
	http.HandleFunc("/delete", api.DeleteOne)
	http.HandleFunc("/find", api.Show)

	// auth
	http.HandleFunc("/signup", auth.AddUser)
	http.HandleFunc("/show", auth.ShowUser)
	http.HandleFunc("/signin", auth.SignIn)
}