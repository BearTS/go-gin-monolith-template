package redis

import (
	"context"
	"sync"
	"time"

	"github.com/BearTS/go-gin-monolith/config"
	"github.com/redis/go-redis/v9"
)

type Connection struct {
	client *redis.Client
}

var once sync.Once
var ctx = context.Background()

// NewConnection
func (conn *Connection) NewConnection() error {

	once.Do(func() {
		conn.client = redis.NewClient(&redis.Options{
			Addr:            config.Redis.Host + ":" + config.Redis.Port,
			Password:        config.Redis.Password,
			DB:              config.Redis.DB,
			MaxRetries:      config.Redis.MaxRetries,
			MinRetryBackoff: time.Duration(config.Redis.MinRetryBackoffMs) * time.Millisecond,
			MaxRetryBackoff: time.Duration(config.Redis.MaxRetryBackoffMs) * time.Millisecond,
			WriteTimeout:    time.Duration(config.Redis.MaxRetries) * time.Millisecond,
		})
	})

	if _, err := conn.client.Ping(ctx).Result(); err != nil {
		return err
	}

	return nil
}

// GetClient
func (conn *Connection) GetClient() *redis.Client {
	return conn.client
}

// Set method will set single key value in redis
func (conn *Connection) Set(key string, value string) error {
	if err := conn.client.Set(ctx, key, value, 0).Err(); err != nil {
		return err
	}
	return nil
}

// SetWithTimeout method will set single key value in redis with timeout
func (conn *Connection) SetWithTimeout(key string, value string, timeout time.Duration) error {
	if err := conn.client.Set(ctx, key, value, timeout).Err(); err != nil {
		return err
	}
	return nil
}

// Get method will fetch single key value from redis
func (conn *Connection) Get(key string) (string, error) {
	val, err := conn.client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

// Del method will remove single key from redis
func (conn *Connection) Del(key string) error {
	if err := conn.client.Del(ctx, key).Err(); err != nil {
		return err
	}
	return nil
}

// DelMulti method will remove multiple keys from redis
func (conn *Connection) DelMulti(keys []string) error {
	if err := conn.client.Del(ctx, keys...).Err(); err != nil {
		return err
	}
	return nil
}

// Close method closes the redis connection
func (conn *Connection) Close() error {
	if conn.client != nil {
		return conn.client.Close()
	}
	return nil
}
