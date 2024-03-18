package producer

import (
	"errors"
	"github.com/pion/webrtc/v3"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
	"log"
)

type Producer struct {
	conn           *webrtc.PeerConnection
	videoRTCSender *webrtc.RTPSender
	audioRTCSender *webrtc.RTPSender
	offer          string
	state          string
}

var _ IProducer = (*Producer)(nil)

func NewProducer() *Producer {
	return &Producer{}
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
	})
	if err != nil {
		logx.Error("Create Peer connection err : ", err)
		return err
	}
	p.conn = peerConn
	for _, typ := range []webrtc.RTPCodecType{webrtc.RTPCodecTypeVideo, webrtc.RTPCodecTypeAudio} {
		if _, err := peerConn.AddTransceiverFromKind(typ, webrtc.RTPTransceiverInit{
			Direction: webrtc.RTPTransceiverDirectionRecvonly,
		}); err != nil {
			log.Print(err)
			return err
		}
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

func (p *Producer) SetLocalTrack(track *webrtc.TrackLocalStaticRTP, kind webrtc.MediaKind) error {
	if kind == webrtc.MediaKindAudio {
		audioSender, err := p.conn.AddTrack(track)
		if err != nil {
			logx.Error("Set local track error")
			return err
		}
		p.audioRTCSender = audioSender
	} else if kind == webrtc.MediaKindVideo {
		videoSender, err := p.conn.AddTrack(track)
		if err != nil {
			logx.Error("Set local track error")
			return err
		}
		p.videoRTCSender = videoSender
	}
	return errors.New("kind not supported")
}

func (p *Producer) CloseConnection() error {
	return p.conn.Close()
}

func (p *Producer) GetVideoSenderRTPTrack() webrtc.TrackLocal {
	return p.videoRTCSender.Track()
}

func (p *Producer) GetAudioSenderRTPTrack() webrtc.TrackLocal {
	return p.audioRTCSender.Track()
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
