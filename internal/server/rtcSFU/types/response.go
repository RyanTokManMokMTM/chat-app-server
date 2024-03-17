package types

type SfuNewProducerResp struct {
	SessionId  string `json:"session_id"`  //RoomId
	ProducerId string `json:"producer_id"` //WebRTC Answer
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
