package auth

import (
	"context"
	"links/configs"
	"links/pkg/jwt"
	"net/http"
	"strings"
)

// JWTMiddleware - middleware для проверки токена
func IsAuthorized(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		t := strings.TrimPrefix(authHeader, "Bearer ")
		token := strings.TrimSpace(t)

		isValue, data := jwt.NewJWT(configs.Config.Secret).Parse(token)
		req := r
		if isValue {
			ctx := context.WithValue(r.Context(), "user_phone", data.Phone)
			req = req.WithContext(ctx)
		}
		next.ServeHTTP(writer, req)
	})
}
