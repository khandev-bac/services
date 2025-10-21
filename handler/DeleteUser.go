package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/services/internals/middleware"
	"github.com/services/utils/common"
)

func (ah *AuthController) DeleteUserAccount(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value(middleware.UserIdKey).(string)
	if !ok {
		WriteJSONError(w, &common.AppError{
			Message: "Unauthorized",
			Code:    http.StatusUnauthorized,
			Err:     errors.New("Unauthorized"),
		})
		return
	}
	userIdToUUID, _ := uuid.Parse(userId)
	s, err := ah.service.DeleteUserAccount(r.Context(), userIdToUUID)
	if err != nil {
		WriteJSONError(w, &common.AppError{
			Message: "Something went wrong",
			Code:    http.StatusInternalServerError,
			Err:     err,
		})
		return
	}
	res := map[string]any{
		"message": s,
		"code":    http.StatusOK,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
