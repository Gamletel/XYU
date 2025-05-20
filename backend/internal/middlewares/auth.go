package middlewares

import (
	"backend/internal/db"
	"backend/internal/models"
	"backend/internal/repositories"
	"backend/pkg/utils"
	"context"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
)

type UserCtxKey struct{}

// Auth — middleware, который проверяет JWT токен
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "missing authorization header", http.StatusUnauthorized)
			return
		}

		if strings.HasPrefix(tokenString, "Bearer ") {
			tokenString = tokenString[7:]
		}

		token, err := jwt.ParseWithClaims(tokenString, &models.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(utils.JWTSecret), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(*models.UserClaims)
		if !ok {
			http.Error(w, "invalid token claims", http.StatusUnauthorized)
			return
		}

		userRepo := repositories.NewUserRepository(db.DB)
		user, err := userRepo.GetUserByEmail(claims.Email)
		if err != nil {
			http.Error(w, "user not found", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), UserCtxKey{}, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
