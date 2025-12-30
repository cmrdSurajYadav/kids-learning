package auth

import (
	"github.com/go-chi/chi/v5"
	"github.com/cmrdSurajYadav/auth-service/internal/modules/auth/handler"
)

func RegisterRoutes(r *chi.Mux, h *handler.AuthHandler) {
	r.Route("/auth", func(r chi.Router) {
		r.Post("/signup", h.Signup)
		r.Post("/login", h.Login)
	})
}
