package handlers

import (
	"net/http"
	"strings"

	"groupie_tracker/internal/api"
	"groupie_tracker/internal/render"
	"groupie_tracker/internal/models"
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

		q := strings.TrimSpace(r.URL.Query().Get("q"))
		if q != "" {
			artists = filterArtists(artists, q)
		}

		v.Render(w, "artists.html", artists)
	}
}

func filterArtists(artists []models.Artist, q string) []models.Artist {
	q = strings.ToLower(q)
	out := make([]models.Artist, 0)

	for _, a := range artists {
		if strings.Contains(strings.ToLower(a.Name), q) {
			out = append(out, a)
			continue
		}
		for _, m := range a.Members {
			if strings.Contains(strings.ToLower(m), q) {
				out = append(out, a)
				break
			}
		}
	}
	return out
}
