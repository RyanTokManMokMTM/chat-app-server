package types

type SFUConnectSessionReq struct {
	SessionId string `json:"session_id"` //RoomId
	SDPType   string `json:"SDPType"`    //Connection WebRTC Offer
}

type SFUGetSessionProducerReq struct {
	SessionId string `json:"session_id"` //RoomId
}

type SFUConsumeProducerReq struct {
	SessionId  string `json:"session_id"`
	ProducerId string `json:"producer_id"` //To produce who
	SDPType    string `json:"SDPType"`
}

// No need to response back to the user
type SFUSendIceCandidateReq struct {
	SessionId        string `json:"session_id"`
	IsProducer       bool   `json:"is_producer"`
	ClientId         string `json:"client_id"`
	IceCandidateType string `json:"ice_candidate_type"`
}

type SFUCloseConnectionReq struct {
	SessionId string `json:"session_id"`
}
