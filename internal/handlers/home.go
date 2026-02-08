package handlers

import (
	"net/http"

	"groupie_tracker/internal/render"
)

func Home(v *render.View) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			w.WriteHeader(http.StatusNotFound)
			v.Render(w, "error.html", map[string]any{
				"Code":    404,
				"Message": "Page introuvable",
			})
			return
		}
		v.Render(w, "home.html", map[string]any{})
	}
}
