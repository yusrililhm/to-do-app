package main

import (
	"fmt"
	"net/http"

	"github.com/yusrililhm/to-do-app/server/routes"
)

func main() {
	routes.Routes()
	fmt.Println("Server Is Running ON PORT 8000")
	http.ListenAndServe(":8000", nil)
}