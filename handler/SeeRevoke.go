package handler

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/services/utils/common"
)

func (ah *AuthController) SeeRevoke(w http.ResponseWriter, r *http.Request) {
	userIdStr := r.Context().Value("user_id").(uuid.UUID)
	revoked, err := ah.service.SeeRevoke(r.Context(), userIdStr)
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
