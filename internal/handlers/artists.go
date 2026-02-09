package handlers

import (
	"net/http"

	"groupie_tracker/internal/render"
)

func Artists(v *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/artists" {
			w.WriteHeader(http.StatusNotFound)
			v.Render(w, "error.html", map[string]any{
				"Code":    404,
				"Message": "Page introuvable",
			})
			return
		}
		v.Render(w, "artists.html", map[string]any{})
	}
}
