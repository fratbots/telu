package main

import (
	"net/http"

	"github.com/fratbots/telu/internal/api/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/scenarios", handlers.ScenariosGetHandler).Methods("GET")
	r.HandleFunc("/game", handlers.GameCreateHandler).Methods("POST")
	r.HandleFunc("/turn", handlers.MakeTurnHandler).Methods("POST")
	http.ListenAndServe("localhost:8095", r)
}
