package handlers

import (
	"net/http"
	"sort"
	"strings"

	"groupie_tracker/internal/api"
	"groupie_tracker/internal/render"
)

type LiveMarker struct {
	Location string
	Artist   string
	Dates    []string
}

func Live(v *render.Render, apiClient *api.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/live" {
			w.WriteHeader(http.StatusNotFound)
			v.Render(w, "error.html", map[string]any{
				"Code": 404, "Message": "Page introuvable",
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

		artistName := make(map[int]string)
		for _, a := range artists {
			artistName[a.ID] = a.Name
		}

		rel, err := apiClient.GetRelations()
		if err != nil {
			w.WriteHeader(http.StatusBadGateway)
			v.Render(w, "error.html", map[string]any{
				"Code": 502, "Message": "Erreur API (relation): " + err.Error(),
			})
			return
		}

		markers := make([]LiveMarker, 0)

		for _, item := range rel.Index {
			name := artistName[item.ID]
			for loc, dates := range item.DatesLocations {
				loc = strings.TrimSpace(loc)
				if loc == "" {
					continue
				}

				markers = append(markers, LiveMarker{
					Location: loc,
					Artist:   name,
					Dates:    dates,
				})
			}
		}

		sort.Slice(markers, func(i, j int) bool {
			if markers[i].Location == markers[j].Location {
				return markers[i].Artist < markers[j].Artist
			}
			return markers[i].Location < markers[j].Location
		})

		v.Render(w, "live.html", markers)
	}
}
