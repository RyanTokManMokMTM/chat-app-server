package producer

import "github.com/pion/webrtc/v3"

type IProducer interface {
	NewConnection(iceServer []string) error
	CreateAnswers(sdp string) (*webrtc.SessionDescription, error)
	CloseConnection() error

	SetLocalTrack(*webrtc.TrackLocalStaticRTP, webrtc.MediaKind) error
	GetVideoSenderRTPTrack() webrtc.TrackLocal
	GetAudioSenderRTPTrack() webrtc.TrackLocal

	GetPeerConnection() *webrtc.PeerConnection
}
