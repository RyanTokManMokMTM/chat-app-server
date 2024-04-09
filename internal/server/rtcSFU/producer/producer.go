package producer

import (
	"errors"
	"github.com/pion/webrtc/v3"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
	"sync"
)

type Producer struct {
	sync.Mutex
	audioTrackLocal *webrtc.TrackLocalStaticRTP
	videoTrackLocal *webrtc.TrackLocalStaticRTP
	conn            *webrtc.PeerConnection
	RTCSenderTracks []webrtc.TrackLocal
	offer           string
	state           string
}

var _ IProducer = (*Producer)(nil)

func NewProducer() *Producer {
	return &Producer{
		RTCSenderTracks: make([]webrtc.TrackLocal, 0),
	}
}

func (p *Producer) NewConnection(
	iceServer []string,
) error {
	peerConn, err := webrtc.NewPeerConnection(webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{
				URLs: iceServer,
			},
		},
		SDPSemantics: webrtc.SDPSemanticsUnifiedPlan,
	})
	if err != nil {
		logx.Error("Create Peer connection err : ", err)
		return err
	}

	if err != nil {
		logx.Error(err)
		return err
	}

	p.conn = peerConn
	//for _, typ := range []webrtc.RTPCodecType{webrtc.RTPCodecTypeVideo, webrtc.RTPCodecTypeAudio} {
	//	if _, err := peerConn.AddTransceiverFromKind(typ, webrtc.RTPTransceiverInit{
	//		Direction: webrtc.RTPTransceiverDirectionSendrecv,
	//	}); err != nil {
	//		log.Print(err)
	//		return err
	//	}
	//}
	return nil
}

func (p *Producer) NewDataChannel(label string, id uint16) error {
	//negotiated := true
	//ordered := false
	channel, err := p.conn.CreateDataChannel(label, &webrtc.DataChannelInit{
		//Ordered:    &ordered,
		//Negotiated: &negotiated,
		//ID:         &id,
	})

	channel.OnOpen(func() {
		logx.Info("data channel opened")
		//
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

	if err != nil {
		return err
	}

	return nil
}

func (p *Producer) CreateAnswers(sdp string) (*webrtc.SessionDescription, error) {
	if p.conn == nil {
		return nil, errors.New("peer connection not yet created")
	}

	err := p.conn.SetRemoteDescription(webrtc.SessionDescription{
		Type: webrtc.SDPTypeOffer, //currentSDP is an offer
		SDP:  sdp,
	})

	if err != nil {
		logx.Error("Set Remote Description err : ", err)
		return nil, err
	}

	ans, err := p.conn.CreateAnswer(&webrtc.AnswerOptions{})
	if err != nil {
		logx.Error("Create Answer err : ", err)
		return nil, err
	}

	err = p.conn.SetLocalDescription(ans)
	if err != nil {
		logx.Error("Set Local Description err : ", err)
		return nil, err
	}
	p.offer = sdp
	//p.peerConnectionListener()
	return &ans, nil
}

func (p *Producer) GetPeerConnection() *webrtc.PeerConnection {
	return p.conn
}

func (p *Producer) SetLocalTracks(kind webrtc.RTPCodecType, rtp *webrtc.TrackLocalStaticRTP) {
	if kind == webrtc.RTPCodecTypeAudio {
		p.audioTrackLocal = rtp
	} else if kind == webrtc.RTPCodecTypeVideo {
		p.videoTrackLocal = rtp
	}
}

func (p *Producer) RemoveLocalTracks(rtp *webrtc.TrackLocalStaticRTP) {
	//TODO: If pc's track has been removed, then update the list
}

func (p *Producer) GetSenderRTPTracks(kind webrtc.RTPCodecType) (*webrtc.TrackLocalStaticRTP, error) {
	if kind == webrtc.RTPCodecTypeAudio {
		return p.audioTrackLocal, nil
	} else if kind == webrtc.RTPCodecTypeVideo {
		return p.videoTrackLocal, nil
	}
	return nil, errors.New("media kind not support")
}
func (p *Producer) CloseConnection() error {

	return p.conn.Close()
}

func (p *Producer) UpdateIceCandidate(data []byte) error {
	if p.conn == nil {
		return errors.New("peer connection is nil")
	}
	iceCandidateType := webrtc.ICECandidateInit{}
	logx.Info("Candidate data : ", string(data))
	if err := jsonx.Unmarshal(data, &iceCandidateType); err != nil {
		return err
	}

	logx.Infof("Adding ice candidate form client %+v", iceCandidateType)

	if err := p.conn.AddICECandidate(iceCandidateType); err != nil {
		return err
	}

	return nil
}
