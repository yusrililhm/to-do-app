package main

import (
	"net/http"

	"github.com/yusrililhm/to-do-app/server/routes"
)

func main() {
	routes.Routes()
	http.ListenAndServe(":8000", nil)
}