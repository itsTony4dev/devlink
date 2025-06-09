package middleware

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"devlink/internal/utils"

	"github.com/golang-jwt/jwt/v5"
)

// userCtxKeyType is a custom type to avoid context key collisions
type userCtxKeyType string

const userCtxKey userCtxKeyType = "user"

// JWTAuthMiddleware validates JWT and attaches user info to the request context
func JWTAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		if header == "" || !strings.HasPrefix(header, "Bearer ") {
			http.Error(w, "Missing or invalid Authorization header", http.StatusUnauthorized)
			return
		}
		tokenString := strings.TrimPrefix(header, "Bearer ")

		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return utils.GetJWTSecret(), nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		// Attach claims to context for use in handlers
		ctx := context.WithValue(r.Context(), userCtxKey, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GetUserClaims extracts JWT claims from context
func GetUserClaims(r *http.Request) (jwt.MapClaims, bool) {
	claims, ok := r.Context().Value(userCtxKey).(jwt.MapClaims)
	return claims, ok
}

// IsUserSelf checks if the user in the JWT matches the user ID in the route param
func IsUserSelf(r *http.Request, paramID string) bool {
	claims, ok := GetUserClaims(r)
	if !ok {
		return false
	}
	jwtUserID, ok := claims["user_id"].(float64)
	if !ok {
		return false
	}
	id, err := strconv.Atoi(paramID)
	if err != nil {
		return false
	}
	return uint(id) == uint(jwtUserID)
}
