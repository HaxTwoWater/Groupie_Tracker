package server

import (
	"fmt"
	"net/http"
)

func Start() {
	http.Handle("/", http.FileServer(http.Dir("./templates")))

	fmt.Println("Starting server and listening on port 8080 -> http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
