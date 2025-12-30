package service

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/cmrdSurajYadav/auth-service/internal/modules/auth/helper"
	"github.com/cmrdSurajYadav/auth-service/internal/modules/auth/model"
	"github.com/cmrdSurajYadav/auth-service/internal/modules/auth/repository"
	"github.com/cmrdSurajYadav/auth-service/internal/modules/dtos"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepo  *repository.UserRepository
	JWTSecret string
}

func NewAuthService(repo *repository.UserRepository, secret string) *AuthService {
	return &AuthService{UserRepo: repo, JWTSecret: secret}
}

func (s *AuthService) Signup(
	ctx context.Context,
	email, password string,
) (*dtos.SignupResponse, error) {

	existing, _ := s.UserRepo.GetByEmail(ctx, email)
	if existing != nil {
		return nil, errors.New("user already exists")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		Email:    email,
		Password: string(hashed),
		Role:     "user",
	}

	if err := s.UserRepo.CreateUser(ctx, user); err != nil {
		return nil, err
	}

	token, err := helper.GenerateToken(
		user.ID,
		user.Email,
		user.Role,
		os.Getenv("JWT_SECRET"),
	)
	if err != nil {
		return nil, err
	}

	return &dtos.SignupResponse{
		ID:    user.ID,
		Email: user.Email,
		Role:  user.Role,
		Token: token,
	}, nil
}

func (s *AuthService) Login(ctx context.Context, email, password string) (string, error) {
	user, err := s.UserRepo.GetByEmail(ctx, email)
	if err != nil || user == nil {
		return "", errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})

	signed, err := token.SignedString([]byte(s.JWTSecret))
	if err != nil {
		return "", err
	}

	return signed, nil
}
