package server

import (
	"github.com/ryantokmanmokmtm/chat-app-server/common/errx"
	"sync"
)

type SFURooms struct {
	sync.Mutex
	rooms   map[string]*SFURoom
	IceUrls []string
}

func NewSFURooms(iceURls []string) *SFURooms {
	return &SFURooms{
		rooms:   make(map[string]*SFURoom),
		IceUrls: iceURls,
	}
}

// CreateNewRoom create a non-existing svc
func (sr *SFURooms) CreateNewRoom(roomUUId string) *SFURoom {
	room := NewSFURoom(roomUUId)
	sr.AddRoom(roomUUId, room)
	return room
}

func (sr *SFURooms) AddRoom(roomUUId string, room *SFURoom) {
	sr.Lock()
	defer sr.Lock()
	if _, ok := sr.rooms[roomUUId]; !ok {
		sr.rooms[roomUUId] = room
	}
}

// FindOneRoom check a svc exists
func (sr *SFURooms) FindOneRoom(roomUUId string) (*SFURoom, error) {
	for _, room := range sr.rooms {
		if room.RoomUUID == roomUUId {
			return room, nil
		}
	}
	return nil, errx.NewCustomErrCode(errx.SFU_ROOM_NOT_FOUND)
}

// FindOneRoomPeer find one peer in a svc
func (sr *SFURooms) FindOneRoomPeer(roomUUId, peerUUId string) (*SFUPeer, error) {
	room, err := sr.FindOneRoom(roomUUId)
	if err != nil {
		return nil, err
	}
	return room.GetPeerById(peerUUId)
}
