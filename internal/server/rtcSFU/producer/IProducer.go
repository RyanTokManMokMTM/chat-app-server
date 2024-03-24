package producer

import "github.com/pion/webrtc/v3"

type IProducer interface {
	NewConnection(iceServer []string) error
	CreateAnswers(sdp string) (*webrtc.SessionDescription, error)
	CloseConnection() error
	UpdateIceCandidate([]byte) error

	SetLocalTrack(rtp *webrtc.TrackLocalStaticRTP) error
	GetSenderRTPTrack() webrtc.TrackLocal
	WriteBufferToTrack(buf []byte) error //from remote

	GetPeerConnection() *webrtc.PeerConnection
}
