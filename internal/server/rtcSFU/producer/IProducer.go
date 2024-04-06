package producer

import "github.com/pion/webrtc/v3"

type IProducer interface {
	NewConnection(iceServer []string) error
	NewDataChannel(label string, id uint16) error
	CreateAnswers(sdp string) (*webrtc.SessionDescription, error)
	CloseConnection() error
	UpdateIceCandidate([]byte) error

	SetLocalTracks(rtp *webrtc.TrackLocalStaticRTP)
	GetSenderRTPTracks() []*webrtc.TrackLocalStaticRTP
	//WriteBufferToTrack(buf []byte) error //from remote

	GetPeerConnection() *webrtc.PeerConnection
}
