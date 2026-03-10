package middlewares

import (
	"context"
	"errors"
	"net/http"

	"github.com/IhsanAlhakim/go-auth-jwt/internal/auth"
)

var ContextWithUserInfoKey = "userInfo"

func (m *Middleware) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			storedToken, err := r.Cookie(m.cfg.TokenCookieName)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			claims, err := auth.VerifyToken(*m.cfg, storedToken.Value)
			if err != nil {
				if errors.Is(err, auth.ErrInvalidSigningMethod) || errors.Is(err, auth.ErrInvalidToken) {
					http.Error(w, "Invalid credentials", http.StatusBadRequest)
					return
				}
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			ctx := context.WithValue(context.Background(), ContextWithUserInfoKey, claims)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
}
