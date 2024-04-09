package socketServer

import (
	"context"
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/pion/webrtc/v3"
	"github.com/ryantokmanmokmtm/chat-app-server/common/errx"
	"github.com/ryantokmanmokmtm/chat-app-server/common/variable"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/server/rtcSFU/sessionManager"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/server/rtcSFU/transportClient"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/server/rtcSFU/types"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/server/socketClient"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/serverTypes"
	svc "github.com/ryantokmanmokmtm/chat-app-server/internal/svc"
	socket_message "github.com/ryantokmanmokmtm/chat-app-server/socket-proto"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/protobuf/encoding/protojson"
	"log"
	"net/http"
	"sync"
)

var Upgrader = websocket.Upgrader{
	//ReadBufferSize:  1024,
	//WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
var _ serverTypes.ISocketServer = (*SocketServer)(nil)

type SocketServer struct {
	sync.Mutex
	Clients    map[string]*socketClient.SocketClient
	register   chan *socketClient.SocketClient
	unRegister chan *socketClient.SocketClient
	multicast  chan []byte
	broadcast  chan []byte

	sessionManager *sessionManager.SessionManager
	testConn       *webrtc.PeerConnection
	//eventListener *listener.SocketEvent
}

func NewSocketServer(iceServerURLs []string) *SocketServer {
	return &SocketServer{
		Clients:        make(map[string]*socketClient.SocketClient),
		register:       make(chan *socketClient.SocketClient),
		unRegister:     make(chan *socketClient.SocketClient),
		multicast:      make(chan []byte),
		broadcast:      make(chan []byte),
		sessionManager: sessionManager.NewSessionManager(),
	}
}

func (s *SocketServer) Start() {
	logx.Info("Starting ws server")
	for {
		select {
		case client := <-s.register:
			logx.Infof("New User is connecting : uuid: %v and name: %v ", client.UUID, client.Name)
			old, ok := s.add(client.UUID, client)

			//TODO: Send a welcome message?
			if ok {
				old.SendMessage(websocket.CloseMessage, []byte("Closed connection")) //TODO: send a close message to client
				old.Closed()                                                         //TODO: close the sending channel of old client
			}

		case client := <-s.unRegister:
			logx.Infof("User %v is leaving.", client.UUID)
			s.remove(client)
		case message := <-s.multicast:
			err := s.multicastMessageHandler(message)
			if err != nil {
				logx.Error("multicast error ", err)
				//send back to the sender?
			}

		case message := <-s.broadcast: //received protoBuffer message -> it need to be decoded
			err := s.broadcastMessageHandler(message)
			if err != nil {
				logx.Error(err)
			}
		}
	}
}

func (s *SocketServer) RegisterClient(client serverTypes.ISocketClient) {
	s.register <- (client).(*socketClient.SocketClient)
}

func (s *SocketServer) UnRegisterClient(client serverTypes.ISocketClient) {
	s.unRegister <- (client).(*socketClient.SocketClient)
}

func (s *SocketServer) BroadcastMessage(message []byte) {
	s.broadcast <- message
}

func (s *SocketServer) MulticastMessage(message []byte) {
	s.multicast <- message
}

// Internal function
func (s *SocketServer) multicastMessageHandler(message []byte) error {
	var socketMessage socket_message.Message
	err := protojson.Unmarshal(message, &socketMessage)
	if err != nil {
		logx.Error(err)
		return err
	}

	if socketMessage.ToUUID != "" && socketMessage.ToUUID != "SFU" {
		//TODO: Send it to someone with a specific Uuid
		logx.Infof("Sending message to %s and type %s", socketMessage.ToUUID, socketMessage.ContentType)
		switch socketMessage.ContentType {
		case variable.TEXT:
			fallthrough
		case variable.IMAGE:
			fallthrough
		case variable.FILE:
			fallthrough
		case variable.AUDIO:
			fallthrough
		case variable.VIDEO:
			fallthrough
		case variable.STORY:
			fallthrough
		case variable.SYS:
			fallthrough
		case variable.REPLY:
			fallthrough
		case variable.STICKER:
			fallthrough
		case variable.SHARED:
			//TODO: save message
			conn, ok := s.Clients[socketMessage.FromUUID]
			if ok && socketMessage.ContentType != variable.SYS {
				//TODO: No need to save system message
				s.saveMessage(conn.SvcCtx, &socketMessage)
			}

			bytes, err := protojson.Marshal(&socketMessage)
			if err != nil {
				logx.Error(err)
				return err
			}

			if socketMessage.MessageType == variable.MESSAGE_TYPE_USERCHAT {
				client, ok := s.Clients[socketMessage.ToUUID]
				if ok {
					client.SendMessage(websocket.BinaryMessage, bytes)
				} else {
					logx.Info("User not connect.")
					ctx := context.Background()
					_, err := variable.RedisConnection.RPush(ctx, socketMessage.ToUUID, string(bytes)).Result()
					if err != nil {
						logx.Error("offline message to redis err %s", err.Error())
						return err
					}
				}

			} else if socketMessage.MessageType == variable.MESSAGE_TYPE_GROUPCHAT {
				logx.Info("Sending Group Message")
				s.sendGroupMessage(&socketMessage, s, conn.SvcCtx)
			}

			if socketMessage.ContentType != variable.SYS {
				//MARK: system message no need to ack??
				logx.Infof("Sending ack with seqId %s to userId :%s", socketMessage.MessageID, socketMessage.FromUUID)
				s.sendAcknowledgement(socketMessage.MessageID, socketMessage.FromUUID)
			}

			break
		default:
			logx.Error("Content type no supported")
			break

		}
	} else if socketMessage.ToUUID == "SFU" {

		switch socketMessage.EventType {
		case variable.SFU_EVENT_PRODUCER_SDP:
			var joinRoomData types.SFUConnectSessionReq
			jsonString := socketMessage.Content //Can be a json string?
			userId := socketMessage.FromUUID
			if err := jsonx.Unmarshal([]byte(jsonString), &joinRoomData); err != nil {
				logx.Error("json unmarshal error", err)
				break
			}
			sdpType := &types.Signaling{}
			if err := jsonx.Unmarshal([]byte(joinRoomData.SDPType), &sdpType); err != nil {
				logx.Error("json unmarshal error(sdp type)", err)
				break
			}

			c, err := s.GetOneClient(userId)
			if err != nil {
				logx.Error("SocketClient not found")
				break
			}
			logx.Info("Joining to session : ", joinRoomData.SessionId)

			//Find a session
			session, err := s.sessionManager.GetOneSession(joinRoomData.SessionId)
			if err != nil {
				logx.Error("Session not found")
				session = s.sessionManager.CreateOneSession(joinRoomData.SessionId)
			}

			logx.Info("Current session Id : ", session.SessionId)

			//Create transport client
			tc := transportClient.NewTransportClient(userId, joinRoomData.SessionId, c)
			logx.Info("Created transport client for ", userId)

			session.AddNewSessionClient(userId, tc)
			//logx.Info("Offer : ", sdpType.SDP)
			//Create the SFU connection
			err = tc.NewConnection(c.SvcCtx.Config.IceServer.Urls, sdpType, func(state webrtc.PeerConnectionState) {
				logx.Info("Connection State changed : ", state)
				switch state {
				case webrtc.PeerConnectionStateConnected:
					logx.Info("Connection State Change : Connected")
					//tc.SignalProducerConnected()
					//TODO: send a signal to all client in the session
					//Send current client info to client that is in the group

					clients := session.GetSessionClients()
					sessionProducersList := make([]types.SFUProducerUserInfo, 0)

					//Send a message to all connected producer.
					ctx := context.Background()
					currentUser, err := c.SvcCtx.DAO.FindOneUserByUUID(ctx, userId)
					if err != nil {
						logx.Error("Get User Info err : ", err)
						break
					}

					currentUserInfo := types.SFUProducerUserInfo{
						ProducerUserId:     currentUser.Uuid,
						ProducerUserName:   currentUser.NickName,
						ProducerUserAvatar: currentUser.Avatar,
					}
					for _, clientId := range clients {
						if clientId != userId {
							//sessionProducersList = append(sessionProducersList, clientId)
							//logx.Infof("Current user %s is Producer", clientId)
							curClient, err := s.GetOneClient(clientId)
							if err != nil {
								logx.Error("Get client err : ", err)
								continue
							}
							logx.Info("Getting user info")
							//TODO: Get Current client info
							ctx := context.Background()
							producerInfo, err := curClient.SvcCtx.DAO.FindOneUserByUUID(ctx, clientId)
							if err != nil {
								logx.Error("Get User Info err : ", err)
								continue
							}

							producerUserInfo := types.SFUProducerUserInfo{
								ProducerUserId:     producerInfo.Uuid,
								ProducerUserName:   producerInfo.NickName,
								ProducerUserAvatar: producerInfo.Avatar,
							}

							//add to session producers list
							sessionProducersList = append(sessionProducersList, producerUserInfo)

							resp := types.SfuNewProducerResp{
								SessionId:    session.SessionId,
								ProducerId:   currentUserInfo.ProducerUserId,
								ProducerInfo: currentUserInfo,
							}

							respStr, err := jsonx.Marshal(resp)
							if err != nil {
								logx.Error("resp marshal error : ", err)
								break
							}

							msg := &socket_message.Message{
								FromUUID:    userId,
								ToUUID:      clientId,
								Content:     string(respStr),
								ContentType: variable.SFU,
								EventType:   variable.SFU_EVENT_SEND_NEW_PRODUCER, //join room.
							}

							msgBytes, err := json.MarshalIndent(msg, "", "\t")
							if err != nil {
								logx.Error(err)
								break
							}
							logx.Info("Sending new producer in to client : ", clientId)
							curClient.SendMessage(websocket.BinaryMessage, msgBytes)
						}
					}

					//Response to producer.
					resp := types.SFUConnectSessionResp{
						SessionId:        session.SessionId,
						ProducerId:       socketMessage.FromUUID,
						SessionProducers: sessionProducersList,
					}

					respStr, err := jsonx.Marshal(resp)
					if err != nil {
						logx.Error("resp marshal error : ", err)
						break
					}

					msg := &socket_message.Message{
						ToUUID:      socketMessage.FromUUID,
						Content:     string(respStr),
						ContentType: variable.SFU,
						EventType:   variable.SFU_EVENT_PRODUCER_CONNECTED, //join room.
					}

					msgBytes, err := json.MarshalIndent(msg, "", "\t")
					if err != nil {
						logx.Error(err)
						break
					}
					c.SendMessage(websocket.BinaryMessage, msgBytes)

					break
				case webrtc.PeerConnectionStateDisconnected:
				case webrtc.PeerConnectionStateClosed:
					logx.Info("Connection State Change : Disconnected")
					//TODO: send a signal to all client in the session
					//clients := session.GetSessionClients()
					//
					//for _, c := range clients {
					//	if c != userId {
					//		receiver, err := s.GetOneClient(c)
					//		if err != nil {
					//			logx.Error(err)
					//			continue
					//		}
					//
					//		resp := types.SFUCloseConnectionResp{
					//			SessionId:  session.SessionId,
					//			ProducerId: userId,
					//		}
					//
					//		respStr, err := jsonx.Marshal(resp)
					//		if err != nil {
					//			logx.Error("resp marshal error : ", err)
					//			break
					//		}
					//
					//		msg := &socket_message.Message{
					//			ToUUID:      socketMessage.FromUUID,
					//			Content:     string(respStr),
					//			ContentType: variable.SFU,
					//			EventType:   variable.SFU_EVENT_CONSUMER_CLOSE, //producer is leave.
					//		}
					//
					//		msgBytes, err := json.MarshalIndent(msg, "", "\t")
					//		if err != nil {
					//			logx.Error(err)
					//			break
					//		}
					//
					//		receiver.SendMessage(websocket.BinaryMessage, msgBytes)
					//
					//	}
					//}
					break
				case webrtc.PeerConnectionStateFailed:
					logx.Info("Connection State Change : Failed")
					//TODO: Close the connection when failed
					if err := tc.Close(); err != nil {
						log.Print(err)
					}
					break
				default:
					break
				}

			})
			if err != nil {
				logx.Error("Create ans error ", err)
				break
			}

			break
		case variable.SFU_EVENT_CONSUMER_SDP:
			//MARK: Same as Create?
			consumeReq := types.SFUConsumeProducerReq{}
			jsonString := socketMessage.Content //Can be a json string?
			userId := socketMessage.FromUUID

			c, err := s.GetOneClient(userId)
			if err != nil {
				logx.Error("Get SocketClient error : ", err)
				break
			}

			if err := jsonx.Unmarshal([]byte(jsonString), &consumeReq); err != nil {
				logx.Error("Unmarshal error")
				break
			}
			logx.Info("Receive offer for consumer : Producer Id : ", consumeReq.ProducerId)

			sdpType := &types.Signaling{}
			if err := jsonx.Unmarshal([]byte(consumeReq.SDPType), &sdpType); err != nil {
				logx.Error("json unmarshal error(sdp type)", err)
				break
			}

			session, err := s.sessionManager.GetOneSession(consumeReq.SessionId)
			if err != nil {
				logx.Error("Get session error : ", err)
				break
			}

			transC, err := session.GetTransportClient(userId)
			if err != nil {
				logx.Error("Get transportClient error : ", err)
				break
			}

			producerClient, err := session.GetTransportClient(consumeReq.ProducerId)
			if err != nil {
				logx.Error(err)
				break
			}

			producer, err := producerClient.GetProducer()
			if err != nil {
				logx.Error(err)
				break
			}
			logx.Info("Consuming....")
			if err := transC.Consume(
				consumeReq.ProducerId,
				c.SvcCtx.Config.IceServer.Urls,
				sdpType,
				producer,
				func(state webrtc.PeerConnectionState) {
					logx.Info("(Consumer)Connection State changed : ", state)
					switch state {
					case webrtc.PeerConnectionStateConnected:
						logx.Info("(Consumer)Connection State Change : Connected")
						break
					case webrtc.PeerConnectionStateDisconnected:
					case webrtc.PeerConnectionStateClosed:
						logx.Info("(Consumer)Connection State Change : Disconnected")
						if err := transC.CloseConsumer(consumeReq.ProducerId); err != nil {
							logx.Error(err)
						}
						break
					case webrtc.PeerConnectionStateFailed:
						logx.Info("(Consumer)Connection State Change : Failed")
						//TODO: Close the connection when failed
						if err := transC.CloseConsumer(consumeReq.ProducerId); err != nil {
							log.Print(err)
						}
						break
					default:
						break
					}
				}); err != nil {
				logx.Errorf("Consume %s error %s", consumeReq.ProducerId, err)
				break
			}
			break
		case variable.SFU_EVENT_CONSUMER_ICE: //same logic as SFU_EVENT_PRODUCER_ICE but is Prodcuer is false
			fallthrough
		case variable.SFU_EVENT_PRODUCER_ICE:
			//Add to ice candindate info into the peer connection that data is provided
			//MARK: Get All producer -> return a list of producerUserId
			iceCandidateReq := types.SFUSendIceCandidateReq{}
			jsonString := socketMessage.Content //Can be a json string?
			iceCandidateType := types.IceCandidateType{}
			userId := socketMessage.FromUUID
			//logx.Info("Received ice candidate from client")
			_, err := s.GetOneClient(userId)
			if err != nil {
				logx.Error("SocketClient not found")
				break
			}

			if err := jsonx.Unmarshal([]byte(jsonString), &iceCandidateReq); err != nil {
				logx.Error("Unmarshal get producers request error : ", err)
				break
			}
			if err := jsonx.Unmarshal([]byte(iceCandidateReq.IceCandidateType), &iceCandidateType); err != nil {
				logx.Error("Unmarshal IceCandidateType error : ", err)
				break
			}

			session, err := s.sessionManager.GetOneSession(iceCandidateReq.SessionId)
			if err != nil {
				logx.Error(err)
				break
			}
			transC, err := session.GetTransportClient(userId) //get current user - transport client obj
			if err != nil {
				logx.Error("Get Transport client error,", err)
				break
			}
			if iceCandidateReq.IsProducer {
				if err := transC.ExchangeIceCandidateForProducer(iceCandidateReq.IceCandidateType); err != nil {
					logx.Error("Exchange ice candzidate for producer error,", err)
					break
				}
			} else {
				if err := transC.ExchangeIceCandidateForConsumers(iceCandidateReq.ClientId, iceCandidateReq.IceCandidateType); err != nil {
					logx.Error("Exchange ice candidate for consumer error,", err)
					break
				}
			}
			break
		case variable.SFU_EVENT_PRODUCER_CLOSE:
			closeConnReq := types.SFUCloseConnectionReq{}
			jsonString := socketMessage.Content //Can be a json string?
			userId := socketMessage.FromUUID

			_, err := s.GetOneClient(socketMessage.FromUUID)
			if err != nil {
				logx.Error("SocketClient not found")
				break
			}

			if err := jsonx.Unmarshal([]byte(jsonString), &closeConnReq); err != nil {
				logx.Error("Unmarshal get producers request error : ", err)
				break
			}

			session, err := s.sessionManager.GetOneSession(closeConnReq.SessionId)
			if err != nil {
				logx.Error("Get one session error , ", err)
				break
			}

			//Send a close message to all session client.
			//Disconnect consumer with userId
			for _, clientId := range session.GetSessionClients() {
				if clientId != userId {
					sessionClient, err := s.GetOneClient(clientId)
					if err != nil {
						logx.Error("Socket client error : ", err)
						continue
					}

					closeResp := types.SFUCloseConnectionResp{
						ProducerId: userId,
					}

					respStr, err := jsonx.Marshal(closeResp)
					if err != nil {
						logx.Error(err)
						continue
					}

					msg := &socket_message.Message{
						ToUUID:      clientId,
						Content:     string(respStr),
						ContentType: variable.SFU,
						EventType:   variable.SFU_EVENT_CONSUMER_CLOSE, // a producer is left
					}

					msgBytes, err := json.MarshalIndent(msg, "", "\t")
					if err != nil {
						logx.Error(err)
						break
					}

					sessionClient.SendMessage(websocket.BinaryMessage, msgBytes)

					//Disconnect consumer
					transClient, err := session.GetTransportClient(clientId)
					if err != nil {
						logx.Error(err)
						break
					}

					if err := transClient.CloseConsumer(userId); err != nil {
						logx.Error(err)
						break
					}
				}
			}

			transC, err := session.GetTransportClient(socketMessage.FromUUID)
			if err != nil {
				logx.Error("Get transport client , ", err)
				break
			}

			//Close and disconnect all consumer connection.
			if err := transC.Close(); err != nil {
				logx.Error("Close connection error ,", err)
				break
			}

			//Remove Current Client from Session
			session.RemoveSessionClient(socketMessage.FromUUID)

			if session.IsEmpty() {
				logx.Info("Session is empty --- removing.....")
				s.sessionManager.RemoveOneSession(session.SessionId)
			}

			break

			//Send from server when producer connected?
		//case variable.SFU_EVENT_GET_PRODUCERS:
		//	//Send all info to client. User Name, User Avatar etc...
		//	//MARK: Get All producer -> return a list of producerUserId
		//	getProducersReq := types.SFUGetSessionProducerReq{}
		//	jsonString := socketMessage.Content //Can be a json string?
		//
		//	c, err := s.GetOneClient(socketMessage.FromUUID)
		//	if err != nil {
		//		logx.Error("SocketClient not found")
		//		break
		//	}
		//
		//	if err := jsonx.Unmarshal([]byte(jsonString), &getProducersReq); err != nil {
		//		logx.Error("Unmarshal get producers request error : ", err)
		//		break
		//	}
		//
		//	session, err := s.sessionManager.GetOneSession(getProducersReq.SessionId)
		//	if err != nil {
		//		logx.Error(err)
		//		break
		//	}
		//
		//	producersList := session.GetSessionClients()
		//
		//	resp := types.SfuGetSessionProducerResp{
		//		SessionId:    session.SessionId,
		//		ProducerList: producersList,
		//	}
		//
		//	respStr, err := jsonx.Marshal(resp)
		//	if err != nil {
		//		logx.Error("resp marshal error : ", err)
		//		break
		//	}
		//
		//	msg := &socket_message.Message{
		//		ToUUID:      socketMessage.FromUUID,
		//		Content:     string(respStr),
		//		ContentType: variable.SFU,
		//		EventType:   variable.SFU_EVENT_GET_PRODUCERS, //join room.
		//	}
		//
		//	msgBytes, err := json.MarshalIndent(msg, "", "\t")
		//	if err != nil {
		//		logx.Error(err)
		//		break
		//	}
		//
		//	c.SendMessage(websocket.BinaryMessage, msgBytes)
		//
		//	break
		//case variable.WEB_RTC:
		//	//logx.Info("Testing RTC testing.")
		//	content := socketMessage.Content
		//	sdp := types.Signaling{}
		//	err := jsonx.Unmarshal([]byte(content), &sdp)
		//	if err != nil {
		//		logx.Error(err)
		//		break
		//	}
		//
		//	client, err := s.GetOneClient(socketMessage.FromUUID)
		//	if err != nil {
		//		logx.Error(err)
		//		break
		//	}
		//	switch sdp.Type {
		//	case "candidate":
		//		candidate := webrtc.ICECandidateInit{}
		//
		//		err = jsonx.Unmarshal([]byte(content), &candidate)
		//		if err != nil {
		//			logx.Error(err)
		//			break
		//		}
		//
		//		if err := s.testConn.AddICECandidate(candidate); err != nil {
		//			logx.Error(err)
		//		}
		//		break
		//	case "offer":
		//		conn, err := webrtc.NewPeerConnection(webrtc.Configuration{
		//			ICEServers: []webrtc.ICEServer{
		//				{
		//					URLs: []string{
		//						"stun:stun.l.google.com:19302",
		//						"stun:stun1.l.google.com:19302",
		//						"stun:stun2.l.google.com:19302",
		//						"stun:stun3.l.google.com:19302",
		//						"stun:stun4.l.google.com:19302",
		//					},
		//				},
		//			},
		//		})
		//
		//		if err != nil {
		//			logx.Error(err)
		//			break
		//		}
		//		s.testConn = conn
		//		for _, typ := range []webrtc.RTPCodecType{webrtc.RTPCodecTypeVideo, webrtc.RTPCodecTypeAudio} {
		//			if _, err := conn.AddTransceiverFromKind(typ, webrtc.RTPTransceiverInit{
		//				Direction: webrtc.RTPTransceiverDirectionRecvonly,
		//			}); err != nil {
		//				logx.Error(err)
		//				break
		//			}
		//		}
		//
		//		err = conn.SetRemoteDescription(webrtc.SessionDescription{
		//			Type: webrtc.SDPTypeOffer, //currentSDP is an offer
		//			SDP:  sdp.SDP,
		//		})
		//
		//		if err != nil {
		//			logx.Error(err)
		//			return err
		//		}
		//
		//		conn.OnICECandidate(func(candidate *webrtc.ICECandidate) {
		//			//logx.Info("Receieving an candindate")
		//			if candidate == nil {
		//				return
		//			}
		//			sdpStr := candidate.ToJSON().Candidate
		//
		//			resp := types.Signaling{
		//				Type: types.CANDIDATE,
		//				Call: sdp.Call,
		//				SDP:  sdpStr,
		//			}
		//
		//			data, err := jsonx.Marshal(resp)
		//
		//			sfuMsg := &socket_message.Message{
		//				ToUUID:      socketMessage.FromUUID, //Back to the user.
		//				Content:     string(data),
		//				ContentType: variable.SYS,
		//				MessageType: variable.MESSAGE_TYPE_USERCHAT,
		//				EventType:   variable.WEB_RTC,
		//			}
		//
		//			msgBytes, err := json.MarshalIndent(sfuMsg, "", "\t")
		//			if err != nil {
		//				logx.Error(err)
		//				return
		//			}
		//
		//			client.SendMessage(websocket.BinaryMessage, msgBytes)
		//		})
		//
		//		conn.OnConnectionStateChange(func(state webrtc.PeerConnectionState) {
		//			logx.Info("Connection State changed : ", state)
		//		})
		//
		//		conn.OnICEConnectionStateChange(func(state webrtc.ICEConnectionState) {
		//			logx.Info("Ice Candidate State changed : ", state)
		//		})
		//
		//		conn.OnTrack(func(remote *webrtc.TrackRemote, receiver *webrtc.RTPReceiver) {
		//			logx.Info("Received tracks from Client ......... waiting for handing a new track.")
		//		})
		//
		//		conn.OnICEGatheringStateChange(func(state webrtc.ICEGathererState) {
		//			logx.Info("Ice Gathering State changed : ", state)
		//		})
		//
		//		conn.OnNegotiationNeeded(func() {
		//			logx.Info("Negotiation Needed State changed")
		//		})
		//
		//		ansDesc, err := conn.CreateAnswer(&webrtc.AnswerOptions{})
		//		if err != nil {
		//			logx.Error(err)
		//			break
		//		}
		//
		//		if err := conn.SetLocalDescription(ansDesc); err != nil {
		//			logx.Error(err)
		//			break
		//		}
		//
		//		ans := ansDesc.SDP
		//
		//		data := types.Signaling{
		//			Type: types.ANSWER,
		//			Call: sdp.Call,
		//			SDP:  ans,
		//		}
		//
		//		resp, err := json.Marshal(data)
		//		if err != nil {
		//			logx.Error(err)
		//			return err
		//		}
		//		sfuMsg := &socket_message.Message{
		//			ToUUID:      socketMessage.FromUUID, //Back to the user.
		//			Content:     string(resp),
		//			ContentType: variable.SYS,
		//			MessageType: variable.MESSAGE_TYPE_USERCHAT,
		//			EventType:   variable.WEB_RTC,
		//		}
		//
		//		msgBytes, err := json.MarshalIndent(sfuMsg, "", "\t")
		//		if err != nil {
		//			logx.Error(err)
		//			return err
		//		}
		//
		//		client.SendMessage(websocket.BinaryMessage, msgBytes)
		//		break
		//	case "answer":
		//		logx.Info("Answer")
		//		break
		//	case "bye":
		//		logx.Info("bye")
		//		break
		//	default:
		//		break
		//
		//	}

		default:
			logx.Infof("Event Type not support")
		}

	}
	return nil
}

func (s *SocketServer) broadcastMessageHandler(message []byte) error {
	var socketMessage socket_message.Message
	err := protojson.Unmarshal(message, &socketMessage)
	if err != nil {
		logx.Error(err)
		return err
	}
	//MARK: to all user....
	//MARK: just like system message..
	for _, client := range s.Clients {
		client.SendMessage(websocket.BinaryMessage, message)
	}
	return nil
}

func (s *SocketServer) add(uuid string, client *socketClient.SocketClient) (*socketClient.SocketClient, bool) {
	s.Lock()
	defer s.Unlock()
	old, ok := s.Clients[uuid]
	s.Clients[client.UUID] = client //add to our map
	if ok {
		return old, ok //TODO: Close the older connection
	}
	return nil, false
}

func (s *SocketServer) remove(client *socketClient.SocketClient) {
	s.Lock()
	defer s.Unlock()

	if _, ok := s.Clients[client.UUID]; ok {
		delete(s.Clients, client.UUID)
	}
}

func (s *SocketServer) GetOneClient(clientUUId string) (*socketClient.SocketClient, error) {
	for _, client := range s.Clients {
		if client.UUID == clientUUId {
			return client, nil
		}
	}
	return nil, errx.NewCustomErrCode(errx.CLIENT_NOT_FOUND)
}

func (s *SocketServer) sendGroupMessage(message *socket_message.Message, server *SocketServer, svcCtx *svc.ServiceContext) {
	//TODO: GET ALL GROUP MEMBER
	//TODO: Check if group is exist
	ctx := context.Background()
	group, err := svcCtx.DAO.FindOneGroupByUUID(ctx, message.ToUUID)
	if err != nil {
		logx.Error(err.Error())
		return
	}

	//TODO: Get All Group Members
	members, err := svcCtx.DAO.FindOneGroupMembers(ctx, group.Id)
	if err != nil {
		logx.Error(err.Error())
		return

	}
	for _, mem := range members {
		if mem.MemberInfo.Uuid == message.FromUUID && message.ContentType != variable.SYS {
			continue
		}

		conn, ok := server.Clients[mem.MemberInfo.Uuid]

		socketMessage := &socket_message.Message{
			MessageID:      message.MessageID,
			ReplyMessageID: message.ReplyMessageID,
			Avatar:         message.Avatar,
			FromUserName:   message.FromUserName,
			FromUUID:       message.ToUUID,   //From Group UUID
			ToUUID:         message.FromUUID, //To Member UUID
			Content:        message.Content,
			ContentType:    message.ContentType,
			MessageType:    message.MessageType,
			EventType:      message.EventType,
			UrlPath:        message.UrlPath,
			GroupName:      group.GroupName,
			GroupAvatar:    group.GroupAvatar,
			FileName:       message.FileName,
			FileSize:       message.FileSize,
		}

		messageBytes, err := json.MarshalIndent(socketMessage, "", "\t")
		if err != nil {
			logx.Error(err)
			continue
		}

		if !ok {
			logx.Infof("Group %v 's member %v is offline", message.ToUUID, mem.MemberInfo.Uuid)
			ctx := context.Background()
			_, err := variable.RedisConnection.RPush(ctx, mem.MemberInfo.Uuid, messageBytes).Result()
			if err != nil {
				logx.Error("offline message to redis err %s", err.Error())
			}
			continue

		}
		conn.SendMessage(websocket.BinaryMessage, messageBytes)
	}
}

// saveMessage, TEXT:Save directly and other types need to be store to FS
func (s *SocketServer) saveMessage(svcCtx *svc.ServiceContext, message *socket_message.Message) {
	//TODO : Save Message into db
	svcCtx.DAO.InsertOneMessage(context.Background(), message)
}

func (s *SocketServer) sendAcknowledgement(seqID string, toUUID string) {
	logx.Infof("Sending ack message with SeqID : %s", seqID)
	acknowledgement := &socket_message.Message{
		MessageID: seqID,
		ToUUID:    toUUID,
		EventType: variable.MSG_ACK,
	}
	ackMessage, err := json.MarshalIndent(acknowledgement, "", "\t")
	if err != nil {
		logx.Error(err)
		return
	}

	client, ok := s.Clients[toUUID]
	logx.Error(toUUID)
	if ok {
		client.SendMessage(websocket.BinaryMessage, ackMessage)
	} else {
		//TODO: Offline
		logx.Info("user is offline -> ack failed")
	}
}
