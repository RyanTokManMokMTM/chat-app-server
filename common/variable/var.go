package variable

import (
	"github.com/redis/go-redis/v9"
)

const (
	MESSAGE_TYPE_USERCHAT = iota + 1
	MESSAGE_TYPE_GROUPCHAT
)

const (
	PONG_MESSAGE = "PONG"
	PING_MESSAGE = "PING"
)

// EventType
const (
	HEAT_BEAT_PING int32 = iota + 1
	HEAT_BEAT_PONG
	SYSTEM
	MESSAGE
	WEB_RTC
	MSG_ACK

	//For SFU feature
	SFU_JOIN
	SFU_ANS
	SFU_ICE
	SFU_CONSUM
	SFU_CLOSE

	ALL //For multiple communication
)

const (
	TEXT = iota + 1
	IMAGE
	FILE
	AUDIO
	VIDEO

	//Story Reply?
	STORY //with Url path and content and available time
	SYS
	REPLY
	STICKER
	SHARED
	SFU
)

const (
	ReadWait  = 60
	WriteWait = 60
	ReadLimit = 1024 * 1024 * 1024
)

var RedisConnection *redis.Client
