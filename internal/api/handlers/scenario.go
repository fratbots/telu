package handlers

import (
	"encoding/json"
	"net/http"
)

// Scenario represents a game scenario.
type Scenario struct {
	ID             int     `json:"id"`
	Author         string  `json:"author"`
	Title          string  `json:"title"`
	Description    string  `json:"description"`
	Rating         float64 `json:"rating"`
	GamesCompleted int     `json:"games_completed"`
}

// ScenariosGetHandler returns scenarios.
func ScenariosGetHandler(w http.ResponseWriter, r *http.Request) {
	scenarios := []Scenario{
		{
			ID:             1,
			Author:         "author_1",
			Title:          "title_1",
			Description:    "description_1",
			Rating:         3.85,
			GamesCompleted: 120,
		},
		{
			ID:             2,
			Author:         "author_2",
			Title:          "title_2",
			Description:    "description_2",
			Rating:         4.18,
			GamesCompleted: 1092,
		},
		{
			ID:             3,
			Author:         "author_1",
			Title:          "title_3",
			Description:    "description_3",
			Rating:         2.41,
			GamesCompleted: 38,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(w).Encode(scenarios)
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
	}
}
