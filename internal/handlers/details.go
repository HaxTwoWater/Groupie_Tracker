package handlers

import (
	"net/http"

	"groupie_tracker/internal/render"
)

func Details(v *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/details" {
			w.WriteHeader(http.StatusNotFound)
			v.Render(w, "error.html", map[string]any{
				"Code":    404,
				"Message": "Page introuvable",
			})
			return
		}
		v.Render(w, "details.html", map[string]any{})
	}
}
