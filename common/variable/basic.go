package variable

import (
	"github.com/redis/go-redis/v9"
)

const (
	PONG_MESSAGE = "PONG"
	PING_MESSAGE = "PING"
)

const (
	ReadWait  = 60
	WriteWait = 60
	ReadLimit = 1024 * 1024 * 1024
)

var RedisConnection *redis.Client
