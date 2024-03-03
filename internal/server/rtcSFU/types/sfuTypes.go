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
	Type       string `json:"type"`
	SessionId  string `json:"session_id"`
	ConsumerId string `json:"consumer_id"`
	Offer      string `json:"data"`
}

type SFUConsumeResp struct {
	Type       string `json:"type"`
	ProducerId string `json:"producer_id"`
	Data       string `json:"data"`
}

type SFUProducerClosedResp struct {
	Type       string `json:"type"`
	ProducerId string `json:"producer_id"`
}

type SFUResponse struct {
	Type string `json:"type"`
	Data string `json:"data"`
}
