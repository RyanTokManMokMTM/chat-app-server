package types

type SFUJoinRoomReq struct {
	SessionId string `json:"session_id"`
	Offer     string `json:"offer"`
}

type SFUConsumeResp struct {
	Type       string `json:"type"`
	ProducerId string `json:"producer_id"`
	Data       string `json:"data"`
}

type SFUResponse struct {
	Type string `json:"type"`
	Data string `json:"data"`
}
