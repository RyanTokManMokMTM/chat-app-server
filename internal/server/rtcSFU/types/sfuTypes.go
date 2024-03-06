package types

type SFUJoinRoomReq struct {
	SessionId string `json:"session_id"`
	Offer     string `json:"offer"`
}

type SFUGetProducerReq struct {
	SessionId string `json:"session_id"`
}

type SFUIceCandindateReq struct {
	SessionId    string `json:"session_id"`
	IsProducer   bool   `json:"is_producer"`
	ToClientId   string `json:"client_id"`
	IceCandidate string `json:"ice_candidate"`
}

type SFUCloseReq struct {
	SessionId string `json:"session_id"`
}

type SFUConsumeReq struct {
	SessionId  string `json:"session_id"`
	ConsumerId string `json:"consumer_id"`
	Offer      string `json:"data"`
}

type SFUConsumeResp struct {
	ProducerId string `json:"producer_id"`
	Data       string `json:"data"`
}

type SFUProducerClosedResp struct {
	ProducerId string `json:"producer_id"`
}

type SFUResponse struct {
	Data string `json:"data"`
}
