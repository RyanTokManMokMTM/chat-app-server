package types

type SFUConnectSessionReq struct {
	SessionID string `json:"session_id"` //RoomID
	SDPType   string `json:"SDPType"`    //Connection WebRTC Offer
	CallType  string `json:"callType"`
}

type SFUGetSessionProducerReq struct {
	SessionID string `json:"session_id"` //RoomID
}

type SFUConsumeProducerReq struct {
	SessionID  string `json:"session_id"`
	ProducerID string `json:"producer_id"` //To produce who
	SDPType    string `json:"SDPType"`
}

// No need to response back to the user
type SFUSendIceCandidateReq struct {
	SessionID        string `json:"session_id"`
	IsProducer       bool   `json:"is_producer"`
	ClientID         string `json:"client_id"`
	IceCandidateType string `json:"ice_candidate_type"`
}

type SFUCloseConnectionReq struct {
	SessionID string `json:"session_id"`
}

type SFUProducerMediaStatusReq struct {
	SessionID string `json:"session_id"`
	ClientID  string `json:"client_id"`
	MediaType string `json:"media_type"`
	IsOn      bool   `json:"is_on"`
}
