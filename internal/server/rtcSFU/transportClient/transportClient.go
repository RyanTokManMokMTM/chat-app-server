package transportClient

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/pion/rtcp"
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
	"time"
)

type TransportClient struct {
	sync.Mutex
	clientId  string
	sessionId string
	//session                *session.Session
	socketClient           *socketClient.SocketClient
	trackLocalGroup        *trackGroup.TrackGroup
	transportProducer      producer.IProducer
	transportConsumers     map[string]consumer.IConsumer // n-1 consumer.
	producerConnectedVideo chan struct{}
	producerConnectedAudio chan struct{}
}

func NewTransportClient(clientId string, sessionId string, socketClient *socketClient.SocketClient) *TransportClient {
	return &TransportClient{
		clientId:     clientId,
		socketClient: socketClient,
		sessionId:    sessionId,
		//session:            session,
		trackLocalGroup:    trackGroup.NewTrackGroup(),
		transportProducer:  producer.NewProducer(),
		transportConsumers: make(map[string]consumer.IConsumer),
	}
}

//
//func (tc *TransportClient) setUpH256VideoTrack() {
//	if tc.transportProducer.GetPeerConnection() == nil {
//		logx.Error("Connection is null")
//		return
//	}
//	pc := tc.transportProducer.GetPeerConnection()
//
//	videoFileName := "/Users/jackson.tmm/Desktop/Projects/chat-app-server/output.h264"
//	h264FrameDuration := time.Millisecond * 33
//	_, err := os.Stat(videoFileName)
//	if err != nil || os.IsExist(err) {
//		panic(err)
//		return
//	}
//
//	track, err := webrtc.NewTrackLocalStaticSample(webrtc.RTPCodecCapability{
//		MimeType:  webrtc.MimeTypeH264,
//		ClockRate: 90000,
//	}, "video", "demo_pion")
//
//	if err != nil {
//		panic(err)
//		return
//	}
//	sender, err := pc.AddTrack(track)
//	if err != nil {
//		panic(err)
//	}
//
//	go func() {
//		buf := make([]byte, 1500)
//		for {
//			if _, _, err := sender.Read(buf); err != nil {
//				return
//			}
//		}
//	}()
//
//	//Write video track to track local
//	go func() {
//		v, err := os.Open(videoFileName)
//		if err != nil {
//			panic(err)
//		}
//
//		h264, err := h264reader.NewReader(v)
//		if err != nil {
//			panic(err)
//		}
//
//		<-tc.producerConnectedVideo
//		logx.Info("Connection Done........ Starting to send video stream data")
//		ticket := time.NewTicker(h264FrameDuration)
//		for ; true; <-ticket.C {
//			//logx.Info("Sending track./**/..")
//			nal, err := h264.NextNAL()
//			if errors.Is(err, io.EOF) {
//				fmt.Printf("All video frames parsed and sent")
//				return
//			}
//			if err != nil {
//				panic(err)
//			}
//
//			if err = track.WriteSample(media.Sample{
//				Data:     nal.Data,
//				Duration: h264FrameDuration,
//			}); err != nil {
//				panic(err)
//			}
//
//		}
//	}()
//}
//
//func (tc *TransportClient) setUpAudioTrack() {
//	if tc.transportProducer.GetPeerConnection() == nil {
//		logx.Error("Connection is null")
//		return
//	}
//	pc := tc.transportProducer.GetPeerConnection()
//
//	audioFileName := "/Users/jackson.tmm/Desktop/Projects/chat-app-server/output.ogg"
//	oggPageDuration := time.Millisecond * 20
//	_, err := os.Stat(audioFileName)
//	if err != nil || os.IsExist(err) {
//		panic(err)
//		return
//	}
//
//	track, err := webrtc.NewTrackLocalStaticSample(webrtc.RTPCodecCapability{
//		MimeType: webrtc.MimeTypeOpus,
//	}, "audio", "pion_audio")
//
//	if err != nil {
//		panic(err)
//		return
//	}
//	sender, err := pc.a(track)
//	if err != nil {
//		panic(err)
//	}
//
//	go func() {
//		buf := make([]byte, 1500)
//		for {
//			if _, _, err := sender.Read(buf); err != nil {
//				return
//			}
//		}
//	}()
//
//	//Write video track to track local
//	go func() {
//		v, err := os.Open(audioFileName)
//		if err != nil {
//			panic(err)
//		}
//		ogg, _, err := oggreader.NewWith(v)
//		if err != nil {
//			panic(err)
//		}
//		var lastGranule uint64
//		<-tc.producerConnectedAudio
//		logx.Info("Connection Done........ Starting to send audio stream data")
//		ticket := time.NewTicker(oggPageDuration)
//		for ; true; <-ticket.C {
//			//logx.Info("Sending track./**/..")
//			pageData, pageHeader, oggErr := ogg.ParseNextPage()
//			if errors.Is(oggErr, io.EOF) {
//				fmt.Printf("All audio pages parsed and sent")
//				os.Exit(0)
//			}
//
//			if oggErr != nil {
//				panic(oggErr)
//			}
//
//			// The amount of samples is the difference between the last and current timestamp
//			sampleCount := float64(pageHeader.GranulePosition - lastGranule)
//			lastGranule = pageHeader.GranulePosition
//			sampleDuration := time.Duration((sampleCount/48000)*1000) * time.Millisecond
//
//			if oggErr = track.WriteSample(media.Sample{Data: pageData, Duration: sampleDuration}); oggErr != nil {
//				panic(oggErr)
//			}
//
//		}
//	}()
//}

func (tc *TransportClient) NewConnection(iceServer []string, sdpType *types.Signaling,
	onConnectionState func(state webrtc.PeerConnectionState),
	onNewTrackReceived func(clientId string, track *webrtc.TrackLocalStaticRTP)) error {
	if err := tc.transportProducer.NewConnection(iceServer); err != nil {
		return err
	}
	if err := tc.transportProducer.NewDataChannel("WebRTCData", 1); err != nil {
		return err
	}

	////Dummy testing -> sending a mp4
	//tc.setUpH256VideoTrack()
	//tc.setUpAudioTrack()
	ans, err := tc.transportProducer.CreateAnswers(sdpType.SDP)
	if err != nil {
		return err
	}

	conn := tc.transportProducer.GetPeerConnection()
	if conn != nil {
		tc.connectionEventHandler(conn, tc.clientId, true, sdpType, onConnectionState, onNewTrackReceived)
	}
	ansStr := ans.SDP
	//ansStr, err := jsonx.Marshal(ans)
	//if err != nil {
	//	return err
	//}
	//ansStr
	sdpData := &types.Signaling{
		Type: types.ANSWER, //ans
		Call: sdpType.Call,
		SDP:  ansStr,
	}

	sdpResp, err := jsonx.Marshal(sdpData)
	if err != nil {
		return err
	}

	sfuResp := types.SfuConnectSessionResp{
		SessionId: tc.sessionId,
		SDPType:   string(sdpResp),
	}

	resp, err := jsonx.Marshal(sfuResp)
	if err != nil {
		return err
	}

	sfuMsg := &socket_message.Message{ //Send ans to producer
		ToUUID:      tc.clientId, //Back to the user.
		Content:     string(resp),
		ContentType: variable.SFU,
		MessageType: variable.MESSAGE_TYPE_GROUPCHAT,
		EventType:   variable.SFU_EVENT_PRODUCER_SDP,
	}

	msgBytes, err := json.MarshalIndent(sfuMsg, "", "\t")
	if err != nil {
		return err
	}

	tc.socketClient.SendMessage(websocket.BinaryMessage, msgBytes)

	return nil
}

func (tc *TransportClient) Consume(
	clientId string,
	iceServer []string,
	sdpType *types.Signaling,
	producer producer.IProducer,
	onConnectionState func(state webrtc.PeerConnectionState),
	onNewTrackReceived func(clientId string, track *webrtc.TrackLocalStaticRTP)) error {
	//TODO: Create consumer...
	newConsumer := consumer.NewConsumer(
		clientId,
	)

	if err := newConsumer.CreateConnection(iceServer); err != nil {
		return err
	}

	for _, t := range producer.GetLocalTracks() {
		logx.Info("Current producer tracks : Kind: ", t.Kind())
		if err := newConsumer.AddLocalTrack(t); err != nil {
			logx.Error(err)
		}
	}

	ans, err := newConsumer.CreateAnswer(sdpType.SDP)
	conn := newConsumer.GetPeerConnection()

	tc.addConsumer(clientId, newConsumer)
	if conn != nil {
		tc.connectionEventHandler(conn, clientId, false, sdpType, onConnectionState, onNewTrackReceived)
	}

	ansStr := ans.SDP
	sdpData := &types.Signaling{
		Type: types.ANSWER, //ans
		Call: sdpType.Call,
		SDP:  ansStr,
	}

	sdpResp, err := jsonx.Marshal(sdpData)
	if err != nil {
		return err
	}

	sfuResp := types.SFUConsumeProducerResp{
		SessionId:  tc.sessionId,
		ProducerId: clientId,
		SDPType:    string(sdpResp),
	}

	resp, err := jsonx.Marshal(sfuResp)
	if err != nil {
		return err
	}

	sfuMsg := &socket_message.Message{
		ToUUID:      tc.clientId, //Back to the user.
		Content:     string(resp),
		ContentType: variable.SFU,
		MessageType: variable.MESSAGE_TYPE_GROUPCHAT,
		EventType:   variable.SFU_EVENT_CONSUMER_SDP,
	}

	msgBytes, err := json.MarshalIndent(sfuMsg, "", "\t")
	if err != nil {
		return err
	}

	tc.socketClient.SendMessage(websocket.BinaryMessage, msgBytes)

	return nil
}

func (tc *TransportClient) connectionEventHandler(
	conn *webrtc.PeerConnection,
	userId string,
	isProducer bool,
	sdpType *types.Signaling,
	onConnectionStatus func(webrtc.PeerConnectionState),
	onNewTrackReceived func(clientId string, track *webrtc.TrackLocalStaticRTP)) {

	conn.OnDataChannel(func(channel *webrtc.DataChannel) {

		channel.OnOpen(func() {
			logx.Info("data channel opened")
			//for range time.NewTicker(1000 * time.Millisecond).C {
			//	_ = channel.Send([]byte(time.Now().String()))
			//}
		})

		channel.OnError(func(err error) {
			logx.Error("channel err ", err)
		})

		channel.OnMessage(func(msg webrtc.DataChannelMessage) {
			logx.Info("Received an message : ", string(msg.Data))
		})

		channel.OnClose(func() {
			logx.Info("Channel closed")
		})

		channel.OnDial(func() {
			logx.Info("dial")
		})
	})

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
		//Change to the type
		if candidate == nil {
			logx.Error("Candidate is null")
			return
		}

		//logx.Info("Received an candindate and sending to client..！！！！！！！！！！！！.")
		sdpCandidate := candidate.ToJSON().Candidate

		signaling := types.Signaling{
			Type: types.CANDIDATE,
			Call: sdpType.Call,
			SDP:  sdpCandidate,
		}

		candidateData, err := jsonx.Marshal(signaling)
		if err != nil {
			logx.Error("ICECandidate marshal error : ", err)
			return
		}
		resp := types.SFUSendIceCandidateReq{
			SessionId:        tc.sessionId,
			IsProducer:       isProducer,
			ClientId:         userId,
			IceCandidateType: string(candidateData),
		}

		respStr, err := jsonx.Marshal(resp)
		if err != nil {
			logx.Error("resp marshal error : ", err)
			return
		}
		eventType := variable.SFU_EVENT_PRODUCER_ICE
		if !isProducer {
			eventType = variable.SFU_EVENT_CONSUMER_ICE
		}

		msg := &socket_message.Message{
			ToUUID:      tc.clientId,
			Content:     string(respStr),
			ContentType: variable.SFU,
			EventType:   eventType, //join room.
		}

		msgBytes, err := json.MarshalIndent(msg, "", "\t")
		if err != nil {
			logx.Error(err)
			return
		}

		tc.socketClient.SendMessage(websocket.BinaryMessage, msgBytes)
	})

	conn.OnConnectionStateChange(onConnectionStatus)

	conn.OnTrack(func(t *webrtc.TrackRemote, receiver *webrtc.RTPReceiver) {
		if t == nil {
			logx.Info("Track is nil")
			return
		}
		logx.Info(t.StreamID())
		logx.Info(t.Codec().MimeType)
		logx.Info(t.Codec().RTPCodecCapability)
		trackId := fmt.Sprintf("%s_%s_%s", userId, t.Kind(), t.ID())
		trackStreamId := fmt.Sprintf("%s_%s_%s", userId, t.Kind(), t.StreamID())

		trackLocal, err := webrtc.NewTrackLocalStaticRTP(
			t.Codec().RTPCodecCapability,
			trackId,
			trackStreamId)
		if err != nil {
			logx.Error(err)
			return
		}

		p, err := tc.GetProducer()
		if err != nil {
			logx.Error(err)
			return
		}

		p.SetLocalTracks(trackLocal)

		//MARK: Announce any consumer in the session that producer has a new track.
		onNewTrackReceived(userId, trackLocal)

		//TODO: Write track data to track.
		go func() {
			t := time.NewTicker(time.Second * 3)
			for ; true; <-t.C {
				//logx.Info("Sending PictureLossIndication")
				_ = conn.WriteRTCP([]rtcp.Packet{
					&rtcp.PictureLossIndication{
						MediaSSRC: uint32(receiver.Track().SSRC()),
					},
				})
			}
		}()

		defer func() {
			logx.Info("Track is ended")
		}()
		for {
			rtp, _, err := t.ReadRTP()
			if err != nil {
				logx.Error(err)
				return
			}

			if err := trackLocal.WriteRTP(rtp); err != nil {
				logx.Error(err)
				return
			}
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

func (tc *TransportClient) ExchangeIceCandidateForProducer(iceData string) error {
	if tc.transportProducer == nil {
		return errors.New("producer not exist")
	}
	return tc.transportProducer.UpdateIceCandidate([]byte(iceData))
}

func (tc *TransportClient) CloseConsumer(clientId string) error {
	if c, ok := tc.transportConsumers[clientId]; ok {
		_ = c.Close()
		tc.removeConsumer(clientId)
		return nil
	}
	return errors.New("consumer not found while closing the connection")
}

func (tc *TransportClient) ExchangeIceCandidateForConsumers(consumerId, iceData string) error {
	c, err := tc.getConsumer(consumerId)
	if err != nil {
		return err
	}
	return c.UpdateIceCandidate([]byte(iceData))
}

func (tc *TransportClient) GetProducer() (producer.IProducer, error) {
	if tc.transportProducer == nil {
		return nil, errors.New("producer not exist")
	}
	return tc.transportProducer, nil
}

func (tc *TransportClient) GetConsumerById(consumerId string) (consumer.IConsumer, error) {
	c, ok := tc.transportConsumers[consumerId]
	if !ok {
		return nil, errors.New("consumer not exist")
	}

	return c, nil
}
