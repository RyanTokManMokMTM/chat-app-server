package consumer

import (
	"errors"
	"github.com/pion/webrtc/v3"
	"github.com/zeromicro/go-zero/core/logx"
)

var _ IConsumer = (*Consumer)(nil)

// Consumer uses to send the media to client, not to receive any media's track from the user it's connected.
type Consumer struct {
	clientId string
	conn     *webrtc.PeerConnection
	state    string

	remoteAudioTrack webrtc.TrackLocal
	remoteVideoTrack webrtc.TrackLocal

	videoRTCSender *webrtc.RTPSender
	audioRTCSender *webrtc.RTPSender
}

func NewConsumer(
	clientId string,
	remoteAudioTrack webrtc.TrackLocal,
	remoteVideoTrack webrtc.TrackLocal,
) *Consumer {
	return &Consumer{
		clientId:         clientId,
		remoteVideoTrack: remoteVideoTrack,
		remoteAudioTrack: remoteAudioTrack,
	}
}

func (c *Consumer) CreateConnection(iceServer []string) error {
	peerConn, err := webrtc.NewPeerConnection(webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{
				URLs: iceServer,
			},
		},
	})
	if err != nil {
		logx.Error("Create Peer connection err : ", err)
		return err
	}
	c.conn = peerConn
	return nil
}

func (c *Consumer) CreateAnswer(sdp string) (*webrtc.SessionDescription, error) {
	if c.conn == nil {
		return nil, errors.New("peer connection not yet created")
	}

	err := c.conn.SetRemoteDescription(webrtc.SessionDescription{
		SDP: sdp,
	})
	if err != nil {
		logx.Error("Set Remote Description err : ", err)
		return nil, err
	}

	ans, err := c.conn.CreateAnswer(&webrtc.AnswerOptions{})
	if err != nil {
		logx.Error("Create Answer err : ", err)
		return nil, err
	}

	err = c.conn.SetLocalDescription(ans)
	if err != nil {
		logx.Error("Set Local Description err : ", err)
		return nil, err
	}
	//c.peerConnectionListener()
	return &ans, nil
}

func (c *Consumer) GetPeerConnection() *webrtc.PeerConnection {
	return c.conn
}

//
//func (c *Consumer) peerConnectionListener() {
//	c.conn.OnICEConnectionStateChange(func(state webrtc.ICEConnectionState) {
//		logx.Info("Ice connection State change : ", state)
//	})
//
//	c.conn.OnICEGatheringStateChange(func(state webrtc.ICEGathererState) {
//		logx.Info("Ice Gathering State change : ", state)
//	})
//
//	c.conn.OnNegotiationNeeded(func() {
//		logx.Info("Negotiation needed")
//	})
//
//	c.conn.OnSignalingStateChange(func(state webrtc.SignalingState) {
//		logx.Info("Signaling State Change ", state)
//	})
//
//	c.conn.OnConnectionStateChange(func(state webrtc.PeerConnectionState) {
//		switch state {
//		case webrtc.PeerConnectionStateNew:
//			logx.Info("Connection State Change : New Connection")
//			break
//		case webrtc.PeerConnectionStateConnecting:
//			logx.Info("Connection State Change : Connecting")
//			break
//		case webrtc.PeerConnectionStateConnected:
//			logx.Info("Connection State Change : Connected")
//			//TODO: added the track into the connection.
//			videoTrackSender, err := c.conn.AddTrack(c.remoteVideoTrack)
//			if err != nil {
//				logx.Error("Consumer add remote video track error: ", err)
//				return
//			}
//			audioTrackSender, err := c.conn.AddTrack(c.remoteAudioTrack)
//			if err != nil {
//				logx.Error("Consumer add remote audio track error: ", err)
//				return
//			}
//
//			c.audioRTCSender = audioTrackSender
//			c.videoRTCSender = videoTrackSender
//			break
//		case webrtc.PeerConnectionStateDisconnected:
//			logx.Info("Connection State Change : Disconnected")
//			break
//		case webrtc.PeerConnectionStateFailed:
//			logx.Info("Connection State Change : Failed")
//			break
//		case webrtc.PeerConnectionStateClosed:
//			logx.Info("Connection State Change : Closed")
//			break
//		}
//	})
//}

// If the connected is connected, we can start to send to track

func (c *Consumer) Close() error {
	return c.conn.Close()
}

func (c *Consumer) ClientId() string {
	return c.clientId
}
