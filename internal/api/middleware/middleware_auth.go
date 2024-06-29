package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/naijasonezeiru/go-phish-backend/internal/api/helper"
	"github.com/naijasonezeiru/go-phish-backend/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jwtToken, err := helper.GetBearerToken(r.Header)
		if err != nil {
			helper.RespondWithError(w, 401, fmt.Sprintf("Authorization err: %v", err))
			return
		}
		user_id, err := helper.DecodeJWTToken(jwtToken)
		if err != nil {
			helper.RespondWithError(w, http.StatusUnauthorized, fmt.Sprintf("Authorization err: %v", err))
			return
		}
		ctx := context.WithValue(r.Context(), "userId", user_id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
