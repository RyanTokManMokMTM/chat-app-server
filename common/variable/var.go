package variable

import "github.com/redis/go-redis/v9"

const (
	MESSAGE_TYPE_USERCHAT = iota + 1
	MESSAGE_TYPE_GROUPCHAT
)

const (
	PONG_MESSAGE = "PONG"
	PING_MESSAGE = "PING"
)

const (
	HEAT_BEAT_PING int32 = iota + 1
	HEAT_BEAT_PONG
	SYSTEM
	MESSAGE
	WEB_RTC
)

const (
	TEXT = iota + 1
	IMAGE
	FILE
	AUDIO
	VIDEO
)

const (
	ReadWait  = 60
	WriteWait = 60
	ReadLimit = 1024
)

var RedisConnection *redis.Client
