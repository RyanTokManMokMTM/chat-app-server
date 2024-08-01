package producer

import "github.com/pion/webrtc/v3"

type IProducer interface {
	NewConnection(iceServer []string) error
	NewDataChannel(label string, id uint16) error

	CreateAnswers(sdp string) (*webrtc.SessionDescription, error)
	CloseConnection() error
	UpdateIceCandidate([]byte) error

	//SetLocalTracks(kind webrtc.RTPCodecType, rtp *webrtc.TrackLocalStaticRTP)
	SetLocalTracks(rtp *webrtc.TrackLocalStaticRTP)
	GetLocalTracks() []*webrtc.TrackLocalStaticRTP

	RemoveLocalTracks(rtp *webrtc.TrackLocalStaticRTP)

	//GetSenderRTPTracks(kind webrtc.RTPCodecType) (*webrtc.TrackLocalStaticRTP, error)
	//WriteBufferToTrack(buf []byte) error //from remote

	GetPeerConnection() *webrtc.PeerConnection
}
