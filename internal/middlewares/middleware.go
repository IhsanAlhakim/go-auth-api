package middlewares

import (
	"github.com/IhsanAlhakim/go-auth-api/internal/config"
	"github.com/boj/redistore"
)

type Middleware struct {
	store *redistore.RediStore
	cfg   *config.Config
}

func New(store *redistore.RediStore, cfg *config.Config) *Middleware {
	return &Middleware{store: store, cfg: cfg}
}
