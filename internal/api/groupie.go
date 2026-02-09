package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"groupie_tracker/internal/models"
)

type Client struct {
	http *http.Client
}

func NewClient() *Client {
	return &Client{
		http: &http.Client{Timeout: 10 * time.Second},
	}
}

func (c *Client) GetArtists() ([]models.Artist, error) {
	resp, err := c.http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("api error: %s", resp.Status)
	}

	var artists []models.Artist
	if err := json.NewDecoder(resp.Body).Decode(&artists); err != nil {
		return nil, err
	}
	return artists, nil
}

func (c *Client) GetRelations() (models.RelationsResponse, error) {
	resp, err := c.http.Get("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
		return models.RelationsResponse{}, err
	}
	defer resp.Body.Close()

	var rel models.RelationsResponse
	if err := json.NewDecoder(resp.Body).Decode(&rel); err != nil {
		return models.RelationsResponse{}, err
	}

	return rel, nil
}
