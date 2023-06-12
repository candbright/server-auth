package redis

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"piano-server/config"
	"time"
)

type DB struct {
	timeout time.Duration
	*redis.Client
}

func NewDB() (*DB, error) {
	var (
		ip       = config.Get("db.redis.ip")
		port     = config.GetInt("db.redis.port")
		password = config.Get("db.redis.password")
		dbName   = config.GetInt("db.mysql.db")
		poolSize = config.GetInt("db.mysql.pool-size")
	)
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", ip, port),
		Password: password,
		DB:       dbName,
		PoolSize: poolSize,
	})
	instance := &DB{
		timeout: 12 * time.Second,
		Client:  client,
	}
	return instance, nil
}
