package listener

type EventListener interface {
	AddEventListener(listener SocketEvent)
	RemoveEventListener(listener SocketEvent)
	NotifyListener(name string)
}
