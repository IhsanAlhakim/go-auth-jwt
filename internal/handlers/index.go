package handlers

import (
	"fmt"
	"net/http"

	"github.com/IhsanAlhakim/go-auth-jwt/internal/middlewares"
	"github.com/golang-jwt/jwt/v5"
)

func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
	userInfo := r.Context().Value(middlewares.ContextWithUserInfoKey).(jwt.MapClaims)
	message := fmt.Sprintf("hello %s", userInfo["name"])
	Response(w, P{Message: message}, http.StatusOK)
}
