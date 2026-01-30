package database

import (
	"github.com/IhsanAlhakim/go-auth-api/internal/config"
	"github.com/boj/redistore"
)

func NewSessionStore(cfg *config.Config) (*redistore.RediStore, error) {

	maxIdleConn := 10
	networkType := "tcp"

	var err error
	EncryptionKey := []byte(cfg.EcryptKey)
	store, err := redistore.NewRediStoreWithDB(maxIdleConn, networkType, cfg.RedisAddr, cfg.RedisUsername, cfg.RedisPassword, cfg.RedisDB, EncryptionKey)
	if err != nil {
		return nil, err
	}
	store.Options.HttpOnly = true
	store.SetMaxLength(4096)
	store.SetKeyPrefix("AUTH_SESSION_")
	store.SetMaxAge(60 * 60) // 1 Hour
	return store, nil
}
