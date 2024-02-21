package sfuType

type SFUJoinEventDataReq struct {
	RoomUUID string `json:"room_uuid"`
	Offer    string `json:"offer"`
}

type SFUJoinEventDataResp struct {
	Answer string `json:"answer"`
}
