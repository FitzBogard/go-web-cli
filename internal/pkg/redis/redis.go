package redis

import (
	"errors"
	"github.com/go-redis/redis/v8"
	"sync"
)

var (
	rds = make(map[string]*redis.Client)
	mu  sync.RWMutex
)

func Set(name string, db *redis.Client) {
	mu.Lock()
	defer mu.Unlock()
	rds[name] = db
}

func Get(name string) (*redis.Client, error) {
	mu.RLock()
	defer mu.RUnlock()

	client, found := rds[name]
	if found {
		return client, nil
	}

	return nil, errors.New("not found redis")
}
