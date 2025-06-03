package types

type SfuNewProducerResp struct {
	SessionID string `json:"session_id"` //RoomID
	//ProducerID   string              `json:"producer_id"` //WebRTC Answer
	ProducerID   string              `json:"producer_id"`
	ProducerInfo SFUProducerUserInfo `json:"producer_info"`
}

type SFUConnectSessionResp struct {
	SessionID        string                `json:"session_id"` //RoomID
	ProducerID       string                `json:"producer_id"`
	SessionProducers []SFUProducerUserInfo `json:"session_producers"`
}

type SFUProducerUserInfo struct {
	ProduceruserID     string `json:"producer_user_id"`
	ProducerUserName   string `json:"producer_user_name"`
	ProducerUserAvatar string `json:"producer_user_avatar"`
}

type SfuConnectSessionResp struct {
	SessionID string `json:"session_id"` //RoomID
	SDPType   string `json:"SDPType"`    //WebRTC Answer
}

type SfuGetSessionProducerResp struct {
	SessionID    string   `json:"session_id"`    //RoomID
	ProducerList []string `json:"producer_list"` //A list of string
}

type SFUConsumeProducerResp struct {
	SessionID  string `json:"session_id"`
	ProducerID string `json:"producer_id"`
	SDPType    string `json:"SDPType"`
}

type SFUCloseConnectionResp struct {
	SessionID  string `json:"session_id"`
	ProducerID string `json:"producer_id"` //who close the connection
}
