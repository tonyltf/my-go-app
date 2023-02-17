package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Routes() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/healthcheck", HealthCheck)

	return r
}

// Hello godoc
//	@Summary	Health checking
//	@Success	200	{string}	string "ok"
//	@Router		/healthcheck [get]
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}
