package consumer

import (
	"errors"
	"github.com/pion/webrtc/v3"
	"github.com/zeromicro/go-zero/core/jsonx"
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

func (c *Consumer) Close() error {
	return c.conn.Close()
}

func (c *Consumer) ClientId() string {
	return c.clientId
}

func (c *Consumer) UpdateIceCandindate(data []byte) error {
	if c.conn == nil {
		return errors.New("peer connection is nil")
	}

	var iceCandindate webrtc.ICECandidateInit
	if err := jsonx.Unmarshal(data, &iceCandindate); err != nil {
		return err
	}

	if err := c.conn.AddICECandidate(iceCandindate); err != nil {
		return err
	}

	return nil
}
