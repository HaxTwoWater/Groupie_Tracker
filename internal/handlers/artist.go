package handlers

import (
	"net/http"
	"strconv"

	"groupie_tracker/internal/api"
	"groupie_tracker/internal/models"
	"groupie_tracker/internal/render"
)

type ArtistPageData struct {
	models.Artist                     // embed => .Name, .Image, .Members, etc.
	Relations     map[string][]string // => .Relations dans le template
}

func Artist(v *render.Render, apiClient *api.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idStr)
		if err != nil || id <= 0 {
			w.WriteHeader(http.StatusBadRequest)
			v.Render(w, "error.html", map[string]any{
				"Code": 400, "Message": "Paramètre id invalide",
			})
			return
		}

		artists, err := apiClient.GetArtists()
		if err != nil {
			w.WriteHeader(http.StatusBadGateway)
			v.Render(w, "error.html", map[string]any{
				"Code": 502, "Message": "Erreur API (artists): " + err.Error(),
			})
			return
		}

		var artist *models.Artist
		for i := range artists {
			if artists[i].Id == id {
				artist = &artists[i]
				break
			}
		}
		if artist == nil {
			w.WriteHeader(http.StatusNotFound)
			v.Render(w, "error.html", map[string]any{
				"Code": 404, "Message": "Artiste introuvable",
			})
			return
		}

		rel, err := apiClient.GetRelations()
		if err != nil {
			w.WriteHeader(http.StatusBadGateway)
			v.Render(w, "error.html", map[string]any{
				"Code": 502, "Message": "Erreur API (relation): " + err.Error(),
			})
			return
		}

		datesLocations := map[string][]string{}
		for _, it := range rel.Index {
			if it.ID == id {
				datesLocations = it.DatesLocations
				break
			}
		}

		data := ArtistPageData{
			Artist:    *artist,
			Relations: datesLocations,
		}

		v.Render(w, "artist.html", data)
	}
}
