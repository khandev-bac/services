package handler

import (
	"encoding/json"
	"net/http"

	"github.com/services/utils/common"
)

func (ah *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	var req common.LoginBody
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		WriteJSONError(w, &common.AppError{
			Message: "Invalid request",
			Code:    http.StatusBadRequest,
			Err:     err,
		})
		return
	}
	user, err := ah.service.Login(r.Context(), req.Email, req.Password)
	if err != nil {
		WriteJSONError(w, &common.AppError{
			Message: err.Error(),
			Code:    http.StatusInsufficientStorage,
			Err:     err,
		})
		return
	}
	WriteJSONResponse(w,
		&common.SuccessResponse{
			Message: "Successfully logged in",
			Code:    http.StatusOK,
			Data: common.UserResponse{
				ID:           user.ID,
				Email:        user.Email,
				AccessToken:  user.AccessToken,
				RefreshToken: user.RefreshToken,
			},
		})
}
