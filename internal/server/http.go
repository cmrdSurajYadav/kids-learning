package server

import (
	"fmt"
	"net/http"

	"github.com/cmrdSurajYadav/auth-service/internal/config"
	"github.com/cmrdSurajYadav/auth-service/internal/modules/auth"
	"github.com/cmrdSurajYadav/auth-service/internal/modules/auth/handler"
	"github.com/cmrdSurajYadav/auth-service/internal/modules/auth/service"
	"github.com/go-chi/chi/v5"
	"github.com/cmrdSurajYadav/auth-service/internal/database" 
	repo "github.com/cmrdSurajYadav/auth-service/internal/modules/auth/repository"

)

type Server struct {
	Router *chi.Mux
	Port   string
}

func NewServer(cfg *config.Config) *Server {
	r := chi.NewRouter()

	// Database setup
	db := database.ConnectDB(cfg)

	// Auth module
	userRepo := repo.NewUserRepository(db)
	authService := service.NewAuthService(userRepo, cfg.JWTSecret)
	authHandler := handler.NewAuthHandler(authService)
	auth.RegisterRoutes(r, authHandler)

	return &Server{
		Router: r,
		Port:   cfg.AppPort,
	}
}

func (s *Server) Start() error {
	fmt.Println("Server running on port", s.Port)
	return http.ListenAndServe(":"+s.Port, s.Router)
}
