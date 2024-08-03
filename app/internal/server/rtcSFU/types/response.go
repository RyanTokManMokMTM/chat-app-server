package types

type SfuNewProducerResp struct {
	SessionId string `json:"session_id"` //RoomId
	//ProducerId   string              `json:"producer_id"` //WebRTC Answer
	ProducerId   string              `json:"producer_id"`
	ProducerInfo SFUProducerUserInfo `json:"producer_info"`
}

type SFUConnectSessionResp struct {
	SessionId        string                `json:"session_id"` //RoomId
	ProducerId       string                `json:"producer_id"`
	SessionProducers []SFUProducerUserInfo `json:"session_producers"`
}

type SFUProducerUserInfo struct {
	ProducerUserId     string `json:"producer_user_id"`
	ProducerUserName   string `json:"producer_user_name"`
	ProducerUserAvatar string `json:"producer_user_avatar"`
}

type SfuConnectSessionResp struct {
	SessionId string `json:"session_id"` //RoomId
	SDPType   string `json:"SDPType"`    //WebRTC Answer
}

type SfuGetSessionProducerResp struct {
	SessionId    string   `json:"session_id"`    //RoomId
	ProducerList []string `json:"producer_list"` //A list of string
}

type SFUConsumeProducerResp struct {
	SessionId  string `json:"session_id"`
	ProducerId string `json:"producer_id"`
	SDPType    string `json:"SDPType"`
}

type SFUCloseConnectionResp struct {
	SessionId  string `json:"session_id"`
	ProducerId string `json:"producer_id"` //who close the connection
}
