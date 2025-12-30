package health

import "github.com/go-chi/chi/v5"

func RegisterRoutes(r *chi.Mux) {
	r.Get("/health", HealthCheck)
}
