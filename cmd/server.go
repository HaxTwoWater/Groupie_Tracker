package server

import (
	"fmt"
	"log"
	"net/http"

	"groupie_tracker/interal/render"
	"groupie_tracker/internal/handlers"
)

func Start() {
	v, err := render.New("web/templates/*.html")
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("web/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	mux.HandleFunc("/", handlers.Home(v))
	mux.HandleFunc("/artists", handlers.Artists(v))
	mux.HandleFunc("/details", handlers.Details(v))
	mux.HandleFunc("/live", handlers.Live(v))

	fmt.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
