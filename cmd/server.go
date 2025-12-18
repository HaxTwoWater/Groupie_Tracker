package server

import (
	"fmt"
	"groupie_tracker/internal/handlers"
	"log"
	"net/http"
)

func Start() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.HomeHandler)

	fs := http.FileServer(http.Dir("web/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
