package types

type SFUConnectSessionReq struct {
	SessionId string `json:"session_id"` //RoomId
	Offer     string `json:"offer"`      //Connection WebRTC Offer
}

type SFUGetSessionProducerReq struct {
	SessionId string `json:"session_id"` //RoomId
}

type SFUConsumeProducerReq struct {
	SessionId  string `json:"session_id"`
	ProducerId string `json:"producer_id"` //To produce who
	Offer      string `json:"offer"`
}

// No need to response back to the user
type SFUSendIceCandindateReq struct {
	SessionId    string `json:"session_id"`
	IsProducer   bool   `json:"is_producer"`
	ClientId     string `json:"client_id"`
	IceCandidate string `json:"ice_candidate"`
}

type SFUCloseConnectionReq struct {
	SessionId string `json:"session_id"`
}
