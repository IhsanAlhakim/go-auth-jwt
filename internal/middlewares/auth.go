package middlewares

import (
	"net/http"

	"github.com/boj/redistore"
)

var store *redistore.RediStore

func (m *Middleware) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r)
		})
}
