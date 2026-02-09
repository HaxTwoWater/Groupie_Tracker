package handlers

import (
	"net/http"

	"groupie_tracker/internal/api"
	"groupie_tracker/internal/render"
)

func Artists(v *render.Render, apiClient *api.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/artists" {
			w.WriteHeader(http.StatusNotFound)
			v.Render(w, "error.html", map[string]any{
				"Code":    404,
				"Message": "Page introuvable",
			})
			return
		}

		artists, err := apiClient.GetArtists()
		if err != nil {
			w.WriteHeader(http.StatusBadGateway)
			v.Render(w, "error.html", map[string]any{
				"Code":    502,
				"Message": "Erreur API (artists) : " + err.Error(),
			})
			return
		}

		v.Render(w, "artists.html", artists)
	}
}
