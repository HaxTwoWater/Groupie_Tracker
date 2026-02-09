package handlers

import (
	"net/http"
	"sort"
	"strings"

	"groupie_tracker/internal/api"
	"groupie_tracker/internal/render"
)

func Live(v *render.Render, apiClient *api.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/live" {
			w.WriteHeader(http.StatusNotFound)
			v.Render(w, "error.html", map[string]any{
				"Code": 404, "Message": "Page introuvable",
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

		seen := make(map[string]bool)
		locations := make([]string, 0)

		for _, item := range rel.Index {
			for loc := range item.DatesLocations {
				loc = strings.TrimSpace(loc)
				if loc == "" {
					continue
				}
				if !seen[loc] {
					seen[loc] = true
					locations = append(locations, loc)
				}
			}
		}

		sort.Strings(locations)

		v.Render(w, "live.html", locations)
	}
}
