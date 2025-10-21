package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/services/db"
	"github.com/services/handler"
	"github.com/services/internals/repository"
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
	r := chi.NewRouter()
	r.Get("/", handler.Test)
	r.Post("/sign-in", handler.SignupHandler)
	r.Post("/login", handler.Login)
	r.Get("/delete-user", handler.DeleteUserAccount)
	r.Get("/seerevoke", handler.SeeRevoke)
	log.Info("Server is connected to port :3000")
	defer log.Sync()
	http.ListenAndServe(constants.PORT, r)
}
