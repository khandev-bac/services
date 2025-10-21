package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/services/internals/middleware"
	"github.com/services/utils/common"
)

func (ah *AuthController) SeeRevoke(w http.ResponseWriter, r *http.Request) {
	userIdStr, ok := r.Context().Value(middleware.UserIdKey).(string)
	if !ok {
		WriteJSONError(w, &common.AppError{
			Message: "Unauthorized",
			Code:    http.StatusUnauthorized,
			Err:     errors.New("user ID missing from context"),
		})
		return
	}
	userID, err := uuid.Parse(userIdStr)
	if err != nil {
		WriteJSONError(w, &common.AppError{
			Message: "Invalid user ID",
			Code:    http.StatusBadRequest,
			Err:     err,
		})
		return
	}
	revoked, err := ah.service.SeeRevoke(r.Context(), userID)
	if err != nil {
		WriteJSONError(w, &common.AppError{
			Message: "Could not check revoke status",
			Code:    http.StatusInternalServerError,
			Err:     err,
		})
		return
	}
	res := map[string]any{
		"revoked": revoked,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
