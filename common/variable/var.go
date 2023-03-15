package variable

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
	FILE
	Image
	AUDIO
	VIDEO
)

const (
	ReadWait  = 60
	WriteWait = 60
	ReadLimit = 1024
)
