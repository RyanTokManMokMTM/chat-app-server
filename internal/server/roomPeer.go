package server

import (
	rtc "github.com/pion/webrtc/v3"
)

type SFUPeer struct {
	client   *SocketClient       //Client Websocket information.
	peerConn *rtc.PeerConnection //Client's peerConnection
}

func NewSFUPeer(client *SocketClient, peerConn *rtc.PeerConnection) (sfuPeer *SFUPeer) {
	sfuPeer = &SFUPeer{
		client:   client,
		peerConn: peerConn,
	}
	return
}

func (sp *SFUPeer) SetLocalDescription(sdp string) {}

func (sp *SFUPeer) SetRemoteDescDescription(sdp string) {}

func (sp *SFUPeer) AddTrack(track *rtc.TrackRemote) {}

func (sp *SFUPeer) AddIceCandidate(candidate string) error {
	return sp.peerConn.AddICECandidate(rtc.ICECandidateInit{Candidate: candidate})
}
