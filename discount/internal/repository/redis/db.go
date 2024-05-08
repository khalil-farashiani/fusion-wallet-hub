package redis

import "github.com/khalil-farashiani/fusion-wallet-hub/discount/internal/adapter/redis"

type DB struct {
	adapter redis.Adapter
}

func New(adapter redis.Adapter) DB {
	return DB{adapter: adapter}
}
