package handlers

import (
	"encoding/json"
	"net/http"
)

// Game represents a game session started for a particular player.
type Game struct {
	ID         string `json:"id"`
	ScenarioID string `json:"scenario_id"`
	PlayerID   string `json:"player_id"`
	Token      string `json:"token"`
}

// GameCreateHandler handles game creation request.
func GameCreateHandler(w http.ResponseWriter, r *http.Request) {
	game := &Game{
		ID:         "game_1",
		ScenarioID: "scenario_1",
		PlayerID:   "player_1",
		Token:      "token_1",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(w).Encode(game)
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
	}
}
