package server

const (
	HEAT_BEAT    string = "HEATBEAT"
	PONG_MESSAGE        = "PONG"
	PING_MESSAGE        = "PING"
)

const (
	MESSAGE_TYPE_USERCHAT
	MESSAGE_TYPE_GROUPCHAT
)

const (
	TEXT = iota + 1
	FILE
	AUDIO
	VIDEO
)

const (
	ReadWait  = 60
	WriteWait = 60
	ReadLimit = 1024
)
