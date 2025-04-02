package middleware

import (
	"context"
	customStatus "god/internal/common/error"
	"god/pkg/jwt"
	"god/pkg/resp"
	"net/http"
	"strings"
)

func JwtMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if len(authHeader) == 0 {
			resp.Return(w, http.StatusUnauthorized, customStatus.UNAUTHORIZED, "Authorization header is required")
			return
		}

		splitToken := strings.Split(authHeader, "Bearer ")
		if len(splitToken) != 2 {
			resp.Return(w, http.StatusUnauthorized, customStatus.UNAUTHORIZED, "Authorization header is invalid")
			return
		}

		payload, err := jwt.Valid(splitToken[1])
		if err != nil {
			resp.Return(w, http.StatusUnauthorized, customStatus.UNAUTHORIZED, err.Error())
			return
		}

		ctx := context.WithValue(r.Context(), "user_id", payload.UserID)
		ctx = context.WithValue(ctx, "role", payload.Role)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
