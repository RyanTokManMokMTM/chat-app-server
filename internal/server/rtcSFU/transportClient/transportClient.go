package transportClient

import (
	"encoding/json"
	"errors"
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

func (tc *TransportClient) NewConnection(iceServer []string, sdp string, onConnectionState func(state webrtc.PeerConnectionState)) error {
	if err := tc.transportProducer.NewConnection(iceServer); err != nil {
		return err
	}
	ans, err := tc.transportProducer.CreateAnswers(sdp)
	if err != nil {
		return err
	}

	conn := tc.transportProducer.GetPeerConnection()
	if conn != nil {
		tc.connectionEventHandler(conn, onConnectionState)
	}

	ansStr, err := jsonx.Marshal(ans)
	if err != nil {
		return err
	}

	sfuResp := types.SFUResponse{
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
		EventType:   variable.SFU_EVENT_SEND_SDP,
	}

	msgBytes, err := json.MarshalIndent(sfuMsg, "", "\t")
	if err != nil {
		return err
	}

	tc.socketClient.SendMessage(websocket.BinaryMessage, msgBytes)

	return nil
}

func (tc *TransportClient) Consume(clientId string, iceServer []string, sdp string, onConnectionState func(state webrtc.PeerConnectionState)) error {
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
		tc.connectionEventHandler(conn, onConnectionState)
	}

	ansStr, err := jsonx.Marshal(ans)
	if err != nil {
		return err
	}

	sfuResp := types.SFUConsumeResp{
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
		EventType:   variable.SFU_EVENT_SEND_SDP,
	}

	msgBytes, err := json.MarshalIndent(sfuMsg, "", "\t")
	if err != nil {
		return err
	}

	tc.socketClient.SendMessage(websocket.BinaryMessage, msgBytes)

	return nil
}

func (tc *TransportClient) connectionEventHandler(conn *webrtc.PeerConnection, onConnectionStatus func(webrtc.PeerConnectionState)) {
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
			EventType:   variable.SFU_EVENT_ICE, //join room.
		}

		msgBytes, err := json.MarshalIndent(msg, "", "\t")
		if err != nil {
			logx.Error(err)
			return
		}

		tc.socketClient.SendMessage(websocket.BinaryMessage, msgBytes)
	})
	conn.OnConnectionStateChange(onConnectionStatus)

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

func (tc *TransportClient) getConsumer(clientId string) (consumer.IConsumer, error) {
	if c, ok := tc.transportConsumers[clientId]; ok {
		return c, nil
	}
	return nil, errors.New("consumer not found")
}

func (tc *TransportClient) Close() error {
	for _, c := range tc.transportConsumers {
		if err := c.Close(); err != nil {
			logx.Error("Close consumer connection err :", err)
		}
		tc.removeConsumer(c.ClientId())
	}
	return tc.transportProducer.CloseConnection()
}

func (tc *TransportClient) GetClientId() string {
	return tc.clientId
}

func (tc *TransportClient) ExchangeIceCandindateForProducer(iceData string) error {
	return tc.transportProducer.UpdateIceCandindate([]byte(iceData))
}

func (tc *TransportClient) CloseConsumer(clientId string) error {
	if c, ok := tc.transportConsumers[clientId]; ok {
		_ = c.Close()
		tc.removeConsumer(clientId)
		return nil
	}
	return errors.New("consumer not found")
}

func (tc *TransportClient) ExchangeIceCandindateForConsumers(consumerId, iceData string) error {
	c, err := tc.getConsumer(consumerId)
	if err != nil {
		return err
	}
	return c.UpdateIceCandindate([]byte(iceData))
}
