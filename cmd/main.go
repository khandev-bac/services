package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/services/db"
	"github.com/services/utils/config"
	"github.com/services/utils/constants"
)

func main() {
	log := config.NewLogger(false)
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db.Connect_Database()
	r := chi.NewRouter()
	log.Info("Server is connected to port :3000")
	defer log.Sync()
	http.ListenAndServe(constants.PORT, r)
}
