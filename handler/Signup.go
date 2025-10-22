package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	eventproducer "github.com/services/eventProducer"
	"github.com/services/internals/service"
	"github.com/services/utils/common"
)

type AuthController struct {
	service       *service.AuthService
	kafkaProducer *eventproducer.KafkaProducer
}

func WriteJSONError(w http.ResponseWriter, appErr *common.AppError) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(appErr.Code)
	json.NewEncoder(w).Encode(appErr)
}
func WriteJSONResponse(w http.ResponseWriter, users *common.SuccessResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(users.Code)
	json.NewEncoder(w).Encode(users)
}
func NewAuthHandler(service *service.AuthService, producer *eventproducer.KafkaProducer) *AuthController {
	return &AuthController{
		service:       service,
		kafkaProducer: producer,
	}
}
func (ah *AuthController) Test(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ok fine"))
}
func (ah *AuthController) SignupHandler(w http.ResponseWriter, r *http.Request) {
	var req common.UserBody
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println("error in handler signup: ", err)
		WriteJSONError(w, &common.AppError{
			Message: "Invalid request",
			Code:    http.StatusBadRequest,
			Err:     err,
		})
		return
	}
	user, err := ah.service.SignUp(r.Context(), req.Username, req.Email, req.Password)
	if err != nil {
		log.Println("error in user signup: ", err.Error())
		WriteJSONError(w, &common.AppError{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
			Err:     err,
		})
		return
	}
	go func() {
		err := ah.kafkaProducer.SendEvent(r.Context(), "user_created", common.KafkaSendValues{
			UserId:   user.ID,
			Email:    user.Email,
			Username: user.Username,
			Picture:  user.Picture,
		})
		if err != nil {
			log.Println("❌ Event sending failed:", err)
			return
		}
		fmt.Println("✅ Event successfully sent")
	}()
	WriteJSONResponse(w, &common.SuccessResponse{
		Message: "Successfully created user",
		Code:    http.StatusOK,
		Data: common.UserResponse{
			ID:           user.ID,
			Email:        user.Email,
			AccessToken:  user.AccessToken,
			RefreshToken: user.RefreshToken,
			Username:     user.Username,
			Picture:      user.Picture,
		},
	})
}
