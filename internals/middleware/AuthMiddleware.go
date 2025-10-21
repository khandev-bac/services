package middleware

import (
	"context"
	"net/http"

	"github.com/services/utils/common"
)

const (
	AuthHandler        = "x-auth-token"
	UserIdKey   string = "user_id"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenStr := r.Header.Get(AuthHandler)
		if tokenStr == "" {
			http.Error(w, "Missing token", http.StatusUnauthorized)
			return
		}
		payload, err := common.VerifyAccessToken(tokenStr)
		if err != nil {
			http.Error(w, "Invalid token or expired token", http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), UserIdKey, payload.Id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
