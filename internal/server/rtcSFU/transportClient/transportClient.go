package transportClient

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/pion/webrtc/v3"
	"github.com/ryantokmanmokmtm/chat-app-server/common/variable"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/server/rtcSFU/consumer"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/server/rtcSFU/producer"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/server/rtcSFU/trackGroup"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/server/rtcSFU/types"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/server/socketClient"
	socket_message "github.com/ryantokmanmokmtm/chat-app-server/socket-proto"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
	"sync"
)

const (
	SDP = "SDP"
	ICE = "ICE"
)

type TransportClient struct {
	sync.Mutex
	clientId           string
	socketClient       *socketClient.SocketClient
	trackLocalGroup    *trackGroup.TrackGroup
	transportProducer  producer.IProducer
	transportConsumers map[string]consumer.IConsumer // n-1 consumer.
}

func NewTransportClient(clientId string, socketClient *socketClient.SocketClient) *TransportClient {
	return &TransportClient{
		clientId:           clientId,
		socketClient:       socketClient,
		trackLocalGroup:    trackGroup.NewTrackGroup(),
		transportProducer:  producer.NewProducer(),
		transportConsumers: make(map[string]consumer.IConsumer),
	}
}

func (tc *TransportClient) NewConnection(iceServer []string, sdp string) error {
	if err := tc.transportProducer.NewConnection(iceServer); err != nil {
		return err
	}
	ans, err := tc.transportProducer.CreateAnswers(sdp)
	if err != nil {
		return err
	}

	conn := tc.transportProducer.GetPeerConnection()
	if conn != nil {
		tc.connectionEventHandler(conn)
	}

	ansStr, err := jsonx.Marshal(ans)
	if err != nil {
		return err
	}

	sfuResp := types.SFUResponse{
		Type: SDP,
		Data: string(ansStr),
	}

	resp, err := jsonx.Marshal(sfuResp)
	if err != nil {
		return err
	}

	sfuMsg := socket_message.Message{
		ToUUID:      tc.clientId, //Back to the user.
		Content:     string(resp),
		ContentType: variable.SFU,
		MessageType: variable.MESSAGE_TYPE_GROUPCHAT,
		EventType:   variable.SFU_ANS,
	}

	msgBytes, err := json.MarshalIndent(sfuMsg, "", "\t")
	if err != nil {
		return err
	}

	tc.socketClient.SendMessage(websocket.BinaryMessage, msgBytes)

	return nil
}

func (tc *TransportClient) Consume(clientId string, iceServer []string, sdp string) error {
	//TODO: Create consumer...
	newConsumer := consumer.NewConsumer(
		clientId,
		tc.transportProducer.GetAudioSenderRTPTrack(),
		tc.transportProducer.GetVideoSenderRTPTrack(),
	)

	if err := newConsumer.CreateConnection(iceServer); err != nil {
		return err
	}
	ans, err := newConsumer.CreateAnswer(sdp)
	conn := newConsumer.GetPeerConnection()

	tc.addConsumer(clientId, newConsumer)
	if conn != nil {
		tc.connectionEventHandler(conn)
	}

	ansStr, err := jsonx.Marshal(ans)
	if err != nil {
		return err
	}

	sfuResp := types.SFUConsumeResp{
		Type:       SDP,
		ProducerId: clientId,
		Data:       string(ansStr),
	}

	resp, err := jsonx.Marshal(sfuResp)
	if err != nil {
		return err
	}

	sfuMsg := socket_message.Message{
		ToUUID:      tc.clientId, //Back to the user.
		Content:     string(resp),
		ContentType: variable.SFU,
		MessageType: variable.MESSAGE_TYPE_GROUPCHAT,
		EventType:   variable.SFU_ANS,
	}

	msgBytes, err := json.MarshalIndent(sfuMsg, "", "\t")
	if err != nil {
		return err
	}

	tc.socketClient.SendMessage(websocket.BinaryMessage, msgBytes)

	return nil
}

func (tc *TransportClient) connectionEventHandler(conn *webrtc.PeerConnection) {
	conn.OnICEConnectionStateChange(func(state webrtc.ICEConnectionState) {
		logx.Info("Ice connection State change : ", state)
	})

	conn.OnICEGatheringStateChange(func(state webrtc.ICEGathererState) {
		logx.Info("Ice Gathering State change : ", state)
	})

	conn.OnNegotiationNeeded(func() {
		logx.Info("Negotiation needed")
	})

	conn.OnSignalingStateChange(func(state webrtc.SignalingState) {
		logx.Info("Signaling State Change ", state)
	})

	conn.OnICECandidate(func(candidate *webrtc.ICECandidate) {
		//Send the candidate to client.
		iceStr, err := jsonx.Marshal(candidate)
		if err != nil {
			logx.Error("ICECandidate marshal error : ", err)
			return
		}

		resp := types.SFUResponse{
			Type: ICE,
			Data: string(iceStr),
		}

		respStr, err := jsonx.Marshal(resp)
		if err != nil {
			logx.Error("resp marshal error : ", err)
			return
		}
		msg := socket_message.Message{
			ToUUID:      tc.clientId,
			Content:     string(respStr),
			ContentType: variable.SFU,
			EventType:   variable.SFU_ICE, //join room.
		}

		msgBytes, err := json.MarshalIndent(msg, "", "\t")
		if err != nil {
			logx.Error(err)
			return
		}

		tc.socketClient.SendMessage(websocket.BinaryMessage, msgBytes)
	})

	conn.OnConnectionStateChange(func(state webrtc.PeerConnectionState) {
		switch state {
		case webrtc.PeerConnectionStateNew:
			logx.Info("Connection State Change : New Connection")
			break
		case webrtc.PeerConnectionStateConnecting:
			logx.Info("Connection State Change : Connecting")
			break
		case webrtc.PeerConnectionStateConnected:
			logx.Info("Connection State Change : Connected")
			break
		case webrtc.PeerConnectionStateDisconnected:
			logx.Info("Connection State Change : Disconnected")
			break
		case webrtc.PeerConnectionStateFailed:
			logx.Info("Connection State Change : Failed")
			break
		case webrtc.PeerConnectionStateClosed:
			logx.Info("Connection State Change : Closed")
			break
		}
	})
}

func (tc *TransportClient) addConsumer(clientId string, ic consumer.IConsumer) {
	tc.Lock()
	defer tc.Unlock()
	if c, ok := tc.transportConsumers[clientId]; ok {
		_ = c.Close()
	}
	tc.transportConsumers[clientId] = ic
}

func (tc *TransportClient) removeConsumer(clientId string) {
	tc.Lock()
	defer tc.Unlock()
	if c, ok := tc.transportConsumers[clientId]; ok {
		_ = c.Close()
		delete(tc.transportConsumers, clientId)
	}
}

func (tc *TransportClient) Close() error {
	return tc.transportProducer.CloseConnection()
}

func (tc *TransportClient) GetClientId() string {
	return tc.clientId
}
