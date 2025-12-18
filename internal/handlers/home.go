package handlers

import (
	"html/template"
	"net/http"
)

var homeTpl = template.Must(template.ParseFiles("web/templates/index.html"))

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	err := homeTpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
		return
	}
}
