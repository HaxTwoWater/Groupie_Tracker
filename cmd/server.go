package server

import (
	"fmt"
	"log"
	"net/http"

	"groupie_tracker/interal/render"
	"groupie_tracker/internal/handlers"
)

func Start() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.HomeHandler)

	fs := http.FileServer(http.Dir("web/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	mux.HandleFunc("/", handlers.Page(v, "home.html"))

	fmt.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
