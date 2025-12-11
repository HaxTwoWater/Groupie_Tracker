package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
		
}

func main() {
	fs := http.FileServer(http.Dir("./"))
	http.Handle("/", http.StripPrefix("/static/", fs))

	fmt.Println("Listening on port 8080 http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
