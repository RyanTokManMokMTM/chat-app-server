package sfu

import (
	"fmt"
	"github.com/pion/randutil"
	"github.com/pion/webrtc/v3"
	"github.com/zeromicro/go-zero/core/logx"
)

type PeerTrackGroup struct {
	VideoTrack *webrtc.TrackLocalStaticRTP
	AudioTrack *webrtc.TrackLocalStaticRTP
}

func NewPeerTrackGroup() *PeerTrackGroup {
	//Get RTP
	videoTrack, err := webrtc.NewTrackLocalStaticRTP(
		webrtc.RTPCodecCapability{
			MimeType: webrtc.MimeTypeVP8,
		},
		fmt.Sprintf("video_%d", randutil.NewMathRandomGenerator().Uint32()), //generate a video id
		fmt.Sprintf("video_%d", randutil.NewMathRandomGenerator().Uint32()), // generate a stream id
	)
	if err != nil {
		logx.Error(err)
		return nil
	}
	audioTrack, err := webrtc.NewTrackLocalStaticRTP(
		webrtc.RTPCodecCapability{
			MimeType: webrtc.MimeTypeOpus,
		},
		fmt.Sprintf("audio_%d", randutil.NewMathRandomGenerator().Uint32()), //generate a video id
		fmt.Sprintf("audio_%d", randutil.NewMathRandomGenerator().Uint32()),
	)
	if err != nil {
		logx.Error(err)
		return nil
	}
	return &PeerTrackGroup{
		VideoTrack: videoTrack,
		AudioTrack: audioTrack,
	}
}