package routes

import (
	"net/http"

	"github.com/IhsanAlhakim/go-auth-api/internal/handlers"
	"github.com/IhsanAlhakim/go-auth-api/internal/middlewares"
	"github.com/IhsanAlhakim/go-auth-api/internal/mux"
)

func Register(mux *mux.Mux, m *middlewares.Middleware, h *handlers.Handler) {
	mux.Handle("GET /users/{id}", m.Auth(http.HandlerFunc(h.GetUser)))
	mux.HandleFunc("POST /users", h.CreateUser)
	mux.Handle("PUT /users/{id}", m.Auth(http.HandlerFunc(h.UpdateUser)))
	mux.Handle("DELETE /users/{id}", m.Auth(http.HandlerFunc(h.DeleteUser)))
	mux.HandleFunc("POST /sessions", h.SignIn)
	mux.Handle("DELETE /sessions", m.Auth(http.HandlerFunc(h.SignOut)))
}
