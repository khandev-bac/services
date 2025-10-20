package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/services/utils/constants"
)

func main() {
	r := chi.NewRouter()
	http.ListenAndServe(constants.PORT, r)
}
