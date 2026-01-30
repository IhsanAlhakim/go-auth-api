package handlers

import (
	"database/sql"

	"github.com/IhsanAlhakim/go-auth-api/internal/config"
	"github.com/boj/redistore"
)

type Handler struct {
	db    *sql.DB
	store *redistore.RediStore
	cfg   *config.Config
}

func New(db *sql.DB, store *redistore.RediStore, cfg *config.Config) *Handler {
	return &Handler{db: db, store: store, cfg: cfg}
}
