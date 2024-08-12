package redisx

import (
	"errors"
	"github.com/redis/go-redis/v9"
)

func ConnectToClient(addr, password string) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0,
	})

	if client == nil {
		return nil, errors.New("failed to connect to redisClient")
	}
	return client, nil
}
