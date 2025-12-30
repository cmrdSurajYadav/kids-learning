package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/cmrdSurajYadav/auth-service/internal/modules/health"
)

func RegisterRoutes(r *chi.Mux) {
	// Register health routes
	health.RegisterRoutes(r)
}
