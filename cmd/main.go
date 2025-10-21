package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/services/db"
	"github.com/services/handler"
	"github.com/services/internals/repository"
	"github.com/services/internals/routes"
	"github.com/services/internals/service"
	"github.com/services/utils/config"
	"github.com/services/utils/constants"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	log := config.NewLogger(false)
	db.Connect_Database()
	db := db.GetQuery()
	repo := repository.NewRepository(db)
	service := service.NewAuthService(repo)
	handler := handler.NewAuthHandler(service)
	route := routes.AuthRoutes(handler)
	log.Info("Server is connected to port :3000")
	defer log.Sync()
	http.ListenAndServe(constants.PORT, route)
}
