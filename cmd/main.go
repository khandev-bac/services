package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/services/db"
	eventproducer "github.com/services/eventProducer"
	"github.com/services/handler"
	"github.com/services/internals/repository"
	"github.com/services/internals/routes"
	"github.com/services/internals/service"
	"github.com/services/utils/common"
	"github.com/services/utils/config"
	"github.com/services/utils/constants"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	logger := config.NewLogger(false)
	db.Connect_Database()
	db := db.GetQuery()
	kafkaProducer, err := eventproducer.NewKafkaProducer("localhost:9092", "User_created")
	if err != nil {
		log.Println("Kafka-Event error: ", err)
	}
	defer kafkaProducer.Close()
	common.InitKey()
	common.InitPublicKey()
	repo := repository.NewRepository(db)
	service := service.NewAuthService(repo)
	handler := handler.NewAuthHandler(service, kafkaProducer)
	route := routes.AuthRoutes(handler)
	logger.Info("Server is connected to port :3000")
	defer logger.Sync()
	http.ListenAndServe(constants.PORT, route)
}
