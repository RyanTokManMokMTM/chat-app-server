package listener

type Event interface {
	Notify()
}

type SocketEvent struct {
	eventName string
}

func NewSocketEvent(eventName string) *SocketEvent {
	return &SocketEvent{
		eventName: eventName,
	}
}
