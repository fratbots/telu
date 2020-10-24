package main

import (
	"log"
	"net/http"
	"os"

	"github.com/fratbots/telu/internal/telegram_bot/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	if token == "" {
		log.Fatalln("No telegram bot token provided")
		os.Exit(1)
	}

	r := mux.NewRouter()
	r.HandleFunc("/", handlers.WebhookHandler).Methods("POST")
	http.ListenAndServe("localhost:8090", r)
}
