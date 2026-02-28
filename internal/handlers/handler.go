package handlers

import (
	"database/sql"

	"github.com/IhsanAlhakim/go-auth-api/internal/config"
)

type Handler struct {
	db  *sql.DB
	cfg *config.Config
}

func New(db *sql.DB, cfg *config.Config) *Handler {
	return &Handler{db: db, cfg: cfg}
}
