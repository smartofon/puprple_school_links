package auth

import (
	"context"
	"links/configs"
	"links/internal/user"
	"links/pkg/api"
	"links/pkg/jwt"
	"net/http"
	"strings"
)

// JWTMiddleware - middleware для проверки токена
func IsAuthorized(repo *user.UserRepository, next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		t := strings.TrimPrefix(authHeader, "Bearer ")
		token := strings.TrimSpace(t)

		req := r
		ctx := context.WithValue(r.Context(), "authorized", false)
		ctx = context.WithValue(ctx, "user_phone", "")
		ctx = context.WithValue(ctx, "user_id", "")

		isValue, data := jwt.NewJWT(configs.Config.Secret).Parse(token)

		if !isValue {
			api.Json(writer, struct{}{}, http.StatusUnauthorized)
			return
		}

		ctx = context.WithValue(ctx, "user_phone", data.Phone)
		user, err := repo.Find(data.Phone)
		if err != nil {
			api.Json(writer, struct{}{}, http.StatusUnauthorized)
			return
		}
		ctx = context.WithValue(ctx, "authorized", true)
		ctx = context.WithValue(ctx, "user_id", int(user.ID))

		req = req.WithContext(ctx)
		next.ServeHTTP(writer, req)
	})
}
