package consumer

import "github.com/pion/webrtc/v3"

type IConsumer interface {
	CreateConnection(iceServer []string) error
	CreateAnswer(sdp string) (*webrtc.SessionDescription, error)
	UpdateIceCandidate(data []byte) error
	Close() error
	ClientId() string

	AddLocalTrack(track *webrtc.TrackLocalStaticRTP) error
	RemoveLocal(track *webrtc.TrackLocalStaticRTP) error

	GetPeerConnection() *webrtc.PeerConnection
}
