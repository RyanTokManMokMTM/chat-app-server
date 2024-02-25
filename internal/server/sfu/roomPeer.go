package sfu

import (
	rtc "github.com/pion/webrtc/v3"
	"github.com/zeromicro/go-zero/core/logx"
)

type SFUPeer struct {
	//socketClient   *socketClient.SocketClient
	ClientId       string
	peerConn       *rtc.PeerConnection //Client's peerConnection
	videoRtpSender *rtc.RTPSender
	audioRtpSender *rtc.RTPSender
}

func NewSFUPeer(clientId string) (sfuPeer *SFUPeer) {
	sfuPeer = &SFUPeer{
		ClientId: clientId,
	}
	return
}

func (sp *SFUPeer) CreatePeerConnection(
	iceServer []string,
	remoteSDP string,
	clientTrackGroup *PeerTrackGroup,
) (*rtc.SessionDescription, error) {
	peerConn, err := rtc.NewPeerConnection(rtc.Configuration{
		ICEServers: []rtc.ICEServer{
			{
				URLs: iceServer,
			},
		},
	})

	if err != nil {
		logx.Error("Create peer connection error : ", err)
		return nil, err
	}

	remoteDecs := rtc.SessionDescription{
		SDP: remoteSDP,
	}

	if err := peerConn.SetRemoteDescription(remoteDecs); err != nil {
		logx.Error("Set remoteDesc err ", err)
		return nil, err
	}

	ans, err := peerConn.CreateAnswer(&rtc.AnswerOptions{})

	if err != nil {
		logx.Error("Create ans : ", err)
		return nil, err
	}

	err = peerConn.SetLocalDescription(ans)

	if err != nil {
		logx.Error("Create ans : ", err)
		return nil, err
	}

	//peerConn.OnTrack(func(remote *rtc.TrackRemote, receiver *rtc.RTPReceiver) {
	//	if receiver != nil && receiver.Tracks()[0] != nil {
	//		//peerConn.AddTrack(webrtc.NewTrackLocalStaticSample())
	//		logx.Infof("Received a track from a new connection with information \n : %+v , %s,%s",
	//			remote.Codec().RTPCodecCapability,
	//			remote.ID(),
	//			remote.StreamID())
	//
	//		sp.videoRtpSender, err = peerConn.AddTrack(sp.socketClient.TrackGroup.VideoTrack)
	//		if err != nil {
	//			return
	//		}
	//
	//		sp.audioRtpSender, err = peerConn.AddTrack(sp.socketClient.TrackGroup.VideoTrack)
	//		if err != nil {
	//			return
	//		}
	//	}
	//})

	peerConn.OnICEConnectionStateChange(func(state rtc.ICEConnectionState) {

	})

	peerConn.OnICEGatheringStateChange(func(state rtc.ICEGathererState) {

	})

	peerConn.OnNegotiationNeeded(func() {

	})

	peerConn.OnSignalingStateChange(func(state rtc.SignalingState) {

	})

	peerConn.OnConnectionStateChange(func(state rtc.PeerConnectionState) {
		//MARK: To handle connection state.
	})

	sp.videoRtpSender, err = peerConn.AddTrack(clientTrackGroup.VideoTrack)
	if err != nil {
		return nil, err
	}

	sp.audioRtpSender, err = peerConn.AddTrack(clientTrackGroup.VideoTrack)
	if err != nil {
		return nil, err
	}

	sp.SetPeerConnection(peerConn)
	return &ans, nil
}

func (sp *SFUPeer) SetPeerConnection(peerConn *rtc.PeerConnection) {
	sp.peerConn = peerConn
}

func (sp *SFUPeer) SetAudioRTPSender(sender *rtc.RTPSender) {
	sp.audioRtpSender = sender
}

func (sp *SFUPeer) SetVideoRTPSender(sender *rtc.RTPSender) {
	sp.videoRtpSender = sender
}

func (sp *SFUPeer) SetLocalDescription(sdp string) {}

func (sp *SFUPeer) SetRemoteDescDescription(sdp string) {}

func (sp *SFUPeer) AddIceCandidate(candidate string) error {
	return sp.peerConn.AddICECandidate(rtc.ICECandidateInit{Candidate: candidate})
}

func (sp *SFUPeer) GetClientID() string {
	return sp.ClientId
}
