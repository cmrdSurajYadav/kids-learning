package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const (
	UserIDKey contextKey = "userID"
	RoleKey   contextKey = "role"
)

func JWTAuth(secret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			// 1️⃣ Read Authorization header
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "missing authorization header", http.StatusUnauthorized)
				return
			}

			// 2️⃣ Expect: Bearer <token>
			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				http.Error(w, "invalid authorization format", http.StatusUnauthorized)
				return
			}

			tokenStr := parts[1]

			// 3️⃣ Parse & validate token
			token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
				// 🔐 Enforce HMAC signing method
				if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, jwt.ErrSignatureInvalid
				}
				return []byte(secret), nil
			})

			if err != nil || !token.Valid {
				http.Error(w, "invalid or expired token", http.StatusUnauthorized)
				return
			}

			// 4️⃣ Extract claims safely
			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				http.Error(w, "invalid token claims", http.StatusUnauthorized)
				return
			}

			// 5️⃣ Extract user_id
			userIDFloat, ok := claims["user_id"].(float64)
			if !ok {
				http.Error(w, "invalid user id", http.StatusUnauthorized)
				return
			}

			role, _ := claims["role"].(string)

			// 6️⃣ Store values in context
			ctx := context.WithValue(r.Context(), UserIDKey, uint(userIDFloat))
			ctx = context.WithValue(ctx, RoleKey, role)

			// 7️⃣ Continue request
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
