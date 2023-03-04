package variable

const (
	MESSAGE_TYPE_USERCHAT = iota
	MESSAGE_TYPE_GROUPCHAT
)

const (
	HEAT_BEAT    string = "HEATBEAT"
	PONG_MESSAGE        = "PONG"
	PING_MESSAGE        = "PING"
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
