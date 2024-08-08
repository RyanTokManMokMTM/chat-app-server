package ws

//
//var Socket serverTypes.ISocketServer
//
//func WebSocketHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
//	if Socket == nil {
//		Socket = socketServer.NewSocketServer(svcCtx.Config.IceServer.Urls)
//		go Socket.Start()
//	}
//
//	return func(w http.ResponseWriter, r *http.Request) {
//		server.ServeWS(svcCtx, w, r, Socket)
//	}
//}
//
//func SendGroupSystemNotification(FromUUID, groupUUID, content string) {
//	logx.Info("Sending notification")
//	msg := &socket_message.Message{
//		MessageID:   uuid.New().String(),
//		FromUUID:    FromUUID,
//		ToUUID:      groupUUID,
//		ContentType: variable.SYS,
//		Content:     content,
//		MessageType: variable.MESSAGE_TYPE_GROUPCHAT,
//		EventType:   variable.MESSAGE,
//	}
//	messageBytes, err := json.MarshalIndent(msg, "", "\t")
//	if err != nil {
//		logx.Error(err)
//		return
//	}
//
//	logx.Info(msg)
//	Socket.MulticastMessage(messageBytes)
//}
