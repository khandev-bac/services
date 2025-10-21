package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/services/handler"
	"github.com/services/internals/middleware"
)

func AuthRoutes(authHandler *handler.AuthController) http.Handler {
	r := chi.NewRouter()
	r.Route("/v1", func(v1 chi.Router) {
		v1.Route("/api", func(api chi.Router) {
			api.Route("/auth", func(auth chi.Router) {
				auth.Get("/", authHandler.Test)
				auth.Post("/sign-up", authHandler.SignupHandler)
				auth.Post("/sign-in", authHandler.Login)
				auth.Post("/refresh", authHandler.Refreshes)
				auth.With(middleware.AuthMiddleware).Get("/delete-user", authHandler.DeleteUserAccount)
				auth.With(middleware.AuthMiddleware).Get("/seerevoke", authHandler.SeeRevoke)
			})
		})
	})
	return r
}
