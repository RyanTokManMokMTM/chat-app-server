package types

type Signaling struct {
	Type string `json:"type"`
	Call string `json:"call"`
	SDP  string `json:"sdp"`
}

type IceCandidateType struct {
	Type          string `json:"type"`
	SDPMLineIndex string `json:"sdpMLineIndex"`
	SDPMid        string `json:"sdpMid"`
	Candidate     string `json:"candidate"`
}

const (
	OFFER     = "offer"
	ANSWER    = "answer"
	CANDIDATE = "candidate"
)

const (
	SIGNALING_OFFER     = "offer"
	SIGNALING_ANSWER    = "answer"
	SIGNALING_CANDIDATE = "candidate"
	SIGNALING_BYE       = "bye"
)

const (
	RTC_CALLING__VOICE = iota + 1
	RTC_CALLING__VIDEO
)
