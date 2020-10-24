package handlers

import (
	"encoding/json"
	"net/http"
)

const (
	headerContentType = "Content-Type"
)

// Request is a user request.
type Request struct {
	Action   Action `json:"action"`
	PlayerID string `json:"player_id"`
	GameID   string `json:"game_id"`
	Token    string `json:"token"`
}

// Action is a possible action player may perform.
type Action struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
}

// Situation is a current state of things from the player's perspective.
type Situation struct {
	Plot    string `json:"plot"`
	Actions []Action
}

// MakeTurnHandler handles request to make next turn.
func MakeTurnHandler(w http.ResponseWriter, r *http.Request) {
	request := new(Request)

	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if request.GameID != "game_1" {
		http.Error(w, "Unknown game ID", http.StatusBadRequest)
		return
	}
	if request.PlayerID != "player_1" {
		http.Error(w, "Unknown player ID", http.StatusBadRequest)
		return
	}
	if request.Token != "token_1" {
		http.Error(w, "Wrong token", http.StatusBadRequest)
		return
	}

	situation := &Situation{
		Plot: "You see the light in the end of tonnel.",
		Actions: []Action{
			{
				ID:          1,
				Description: "Switch off the light",
			},
			{
				ID:          2,
				Description: "Turn back and run",
			},
			{
				ID:          3,
				Description: "Cry for help",
			},
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(situation)
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
	}
}
