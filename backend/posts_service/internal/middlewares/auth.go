package middlewares

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
)

type ContextKey string

const (
	UserIDKey ContextKey = "user_id"
	TokenKey  ContextKey = "token"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Исключить эндпоинты /health и /ready
		if r.URL.Path == "/health" || r.URL.Path == "/ready" {
			next.ServeHTTP(w, r)
			return
		}

		tokenString := r.Header.Get("Authorization")

		if tokenString == "" {
			http.Error(w, "Authorization token missing", http.StatusUnauthorized)
			return
		}

		if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
			tokenString = tokenString[7:]
		}

		claims := jwt.MapClaims{}
		secret := os.Getenv("JWT_SECRET")
		if secret == "" {
			log.Println("JWT_SECRET not found in environment")
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})
		if err != nil || !token.Valid {
			log.Println("AuthMiddleware: Invalid token")
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		log.Printf("Claims: %v", claims)

		userIDFloat, ok := claims["user_id"].(float64)
		if !ok {
			log.Println("AuthMiddleware: user_id not found in claims")
			http.Error(w, "Invalid user ID in token", http.StatusUnauthorized)
			return
		}
		userID := int(userIDFloat)

		log.Printf("Authorized user: %d", userID)

		ctx := context.WithValue(r.Context(), UserIDKey, userID)
		ctx = context.WithValue(ctx, TokenKey, tokenString)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
