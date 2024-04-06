package consumer

import (
	"errors"
	"github.com/pion/webrtc/v3"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
	"log"
)

var _ IConsumer = (*Consumer)(nil)

// Consumer uses to send the media to client, not to receive any media's track from the user it's connected.
type Consumer struct {
	clientId string
	conn     *webrtc.PeerConnection

	trackLocal *webrtc.TrackLocalStaticRTP
}

func NewConsumer(
	clientId string,
) *Consumer {
	return &Consumer{
		clientId: clientId,
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
	for _, typ := range []webrtc.RTPCodecType{webrtc.RTPCodecTypeVideo, webrtc.RTPCodecTypeAudio} {
		if _, err := peerConn.AddTransceiverFromKind(typ, webrtc.RTPTransceiverInit{
			Direction: webrtc.RTPTransceiverDirectionSendonly,
		}); err != nil {
			log.Print(err)
			return err
		}
	}

	c.conn = peerConn
	return nil
}

func (c *Consumer) CreateAnswer(sdp string) (*webrtc.SessionDescription, error) {
	if c.conn == nil {
		return nil, errors.New("peer connection not yet created")
	}

	err := c.conn.SetRemoteDescription(webrtc.SessionDescription{
		Type: webrtc.SDPTypeOffer, //currentSDP is an offer
		SDP:  sdp,
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

func (c *Consumer) UpdateIceCandidate(data []byte) error {
	if c.conn == nil {
		return errors.New("peer connection is nil")
	}
	iceCandidateType := webrtc.ICECandidateInit{}
	//logx.Info("Candidate data : ", string(data))
	if err := jsonx.Unmarshal(data, &iceCandidateType); err != nil {
		return err
	}

	logx.Infof("Adding ice candidate form client %+v", iceCandidateType)

	if err := c.conn.AddICECandidate(iceCandidateType); err != nil {
		return err
	}

	return nil
}

func (c *Consumer) AddLocalTrack(track *webrtc.TrackLocalStaticRTP) error {
	if track == nil {
		return errors.New("track is nil")
	}
	logx.Infof("Add Track")
	sender, err := c.conn.AddTrack(track)
	if err != nil {
		return err
	}
	go func() {
		b := make([]byte, 1600)
		for {
			if _, _, err := sender.Read(b); err != nil {
				logx.Error(err)
				return
			}
		}
	}()
	return nil
}
