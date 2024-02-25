package sfu

import (
	rtc "github.com/pion/webrtc/v3"
)

type SFUPeer struct {
	clientId string
	peerConn *rtc.PeerConnection //Client's peerConnection
}

func NewSFUPeer(clientId string, peerConn *rtc.PeerConnection) (sfuPeer *SFUPeer) {
	sfuPeer = &SFUPeer{
		clientId: clientId,
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

func (sp *SFUPeer) GetClientID() string {
	return sp.clientId
}
