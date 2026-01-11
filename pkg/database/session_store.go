package database

import (
	"os"

	"github.com/boj/redistore"
)

var store *redistore.RediStore
var SessionID string

func NewSessionStore() (*redistore.RediStore, error) {
	SessionID = os.Getenv("SESSION_ID")
	encryptionKey := os.Getenv("ENCRYPTION_KEY")

	maxIdleConn := 10
	networkType := "tcp"
	redisServerAddress := os.Getenv("REDIS_ADDRESS")
	redisUsername := os.Getenv("REDIS_USERNAME")
	redisPassword := os.Getenv("REDIS_PASSWORD")
	redisDB := os.Getenv("REDIS_DB")

	var err error
	EncryptionKey := []byte(encryptionKey)
	store, err = redistore.NewRediStoreWithDB(maxIdleConn, networkType, redisServerAddress, redisUsername, redisPassword, redisDB, EncryptionKey)
	if err != nil {
		return nil, err
	}
	store.Options.HttpOnly = true
	store.SetMaxLength(4096)
	store.SetKeyPrefix("AUTH_SESSION_")
	store.SetMaxAge(60 * 60) // 1 Hour
	return store, nil
}

func GetSessionStore() *redistore.RediStore {
	return store
}
