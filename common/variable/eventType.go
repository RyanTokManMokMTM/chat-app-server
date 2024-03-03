package variable

// EventType
const (
	HEAT_BEAT_PING int32 = iota + 1
	HEAT_BEAT_PONG
	SYSTEM
	MESSAGE
	WEB_RTC
	MSG_ACK

	//For SFU feature
	SFU_CONNECT
	SFU_ANS
	SFU_ICE
	SFU_GET_RPODUCERS
	SFU_CONSUM
	SFU_CLOSE

	ALL //For multiple communication
)
