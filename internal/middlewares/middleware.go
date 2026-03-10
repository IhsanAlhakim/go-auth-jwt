package middlewares

import (
	"github.com/IhsanAlhakim/go-auth-jwt/internal/config"
)

type Middleware struct {
	cfg *config.Config
}

func New(cfg *config.Config) *Middleware {
	return &Middleware{cfg: cfg}
}
