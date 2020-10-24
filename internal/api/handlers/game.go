package handlers

import (
	"encoding/json"
	"net/http"
)

// GameSession represents a game started for a particular player.
type GameSession struct {
	ID       string `json:"id"`
	GameID   string `json:"game_id"`
	PlayerID string `json:"player_id"`
	Token    string `json:"token"`
}

// GameCreateHandler handles game creation request.
func GameCreateHandler(w http.ResponseWriter, r *http.Request) {
	game := &GameSession{
		ID:       "game_session_1",
		GameID:   "game_1",
		PlayerID: "player_1",
		Token:    "token_1",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(w).Encode(game)
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
	}
}
