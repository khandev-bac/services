package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/services/utils/config"
	"github.com/services/utils/constants"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	log := config.NewLogger(false)
	defer log.Sync()
	r := chi.NewRouter()
	log.Info("Server is connected to port :3000")
	http.ListenAndServe(constants.PORT, r)
}
