package serverTypes

type ISocketClient interface {
	//ReadLoop()
	OnEvent(string, []byte) error
	//WriteLoop()
	Closed()
	SendMessage(socketMessageType int, message []byte)
	ReceiveMessage(message []byte)
}

type ISocketServer interface {
	Start()
	RegisterClient(ISocketClient)
	UnRegisterClient(ISocketClient)
	BroadcastMessage(message []byte)
	MulticastMessage(message []byte)
}
