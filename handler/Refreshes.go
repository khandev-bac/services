package handler

import (
	"encoding/json"
	"net/http"

	"github.com/services/utils/common"
)

func (ah *AuthController) Refreshes(w http.ResponseWriter, r *http.Request) {
	var req common.Refreshes
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		WriteJSONError(w, &common.AppError{
			Message: "Invalid request",
			Code:    http.StatusBadRequest,
			Err:     err,
		})
		return
	}
	payload, err := common.VerifyRefreshToken(req.RefreshToken)
	if err != nil {
		WriteJSONError(w, &common.AppError{
			Message: "Failed verify user",
			Code:    http.StatusUnauthorized,
			Err:     err,
		})
		return
	}
	newTokens := common.GenerateToken(common.Payloads{
		Id:       payload.Id,
		Email:    payload.Email,
		Username: payload.Username,
	})
	WriteJSONResponse(w, &common.SuccessResponse{
		Message: "verifed",
		Code:    http.StatusOK,
		Data: common.UserResponse{
			ID:           payload.Id,
			Email:        payload.Email,
			AccessToken:  newTokens.AccessToken,
			RefreshToken: newTokens.RefreshToken,
		},
	})
}
