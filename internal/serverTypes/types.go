package serverTypes

type SocketClientIf interface {
	//ReadLoop()
	OnEvent(int32, []byte) error
	//WriteLoop()
	Closed()
	SendMessage(socketMessageType int, message []byte)
	ReceiveMessage(message []byte)
}

type SocketServerIf interface {
	Start()
	RegisterClient(SocketClientIf)
	UnRegisterClient(SocketClientIf)
	BroadcastMessage(message []byte)
	MulticastMessage(message []byte)
}
