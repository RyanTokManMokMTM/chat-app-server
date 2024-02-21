package server

import (
	"errors"
	"github.com/pion/webrtc/v3"
	"sync"
)

type SFURoom struct {
	//A map of connected Clinet
	sync.Mutex
	RoomUUID    string
	producerMap map[string]*SFUPeer
	consumerMap map[string]*SFUPeer
}

func NewSFURoom(roomUUID string) *SFURoom {
	return &SFURoom{
		RoomUUID:    roomUUID,
		producerMap: make(map[string]*SFUPeer),
		consumerMap: make(map[string]*SFUPeer),
	}
}

func (sf *SFURoom) NewConnection(
	iceServerURls []string,
	client *SocketClient,
	remoteSDP string,
	onTrackFunc func(peerConn *webrtc.PeerConnection, remote *webrtc.TrackRemote, receiver *webrtc.RTPReceiver)) (*webrtc.SessionDescription, error) {

	peerConn, err := webrtc.NewPeerConnection(webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{
				URLs: iceServerURls,
			},
		},
	})

	if err != nil {
		return nil, err
	}
	//Set up peer connection.
	removeDesc := webrtc.SessionDescription{
		SDP: remoteSDP,
	}

	if err := peerConn.SetRemoteDescription(removeDesc); err != nil {
		return nil, err
	}

	peerConn.OnTrack(func(remote *webrtc.TrackRemote, receiver *webrtc.RTPReceiver) {
		onTrackFunc(peerConn, remote, receiver)
	})

	//Create a answers and return to client.
	ans, err := peerConn.CreateAnswer(&webrtc.AnswerOptions{
		//TODO: some option here..
	})

	if err != nil {
		return nil, err
	}

	if err := peerConn.SetLocalDescription(ans); err != nil {
		return nil, err
	}

	//What about the track? -> set on the upper layer?

	//MARK : Added current peer to server map
	peerClient := NewSFUPeer(client, peerConn)
	sf.addNewClient(client.UUID, peerClient)
	return &ans, nil
}

func (sf *SFURoom) CloseConnection(clientId string) {
	for _, peer := range sf.producerMap {
		if peer.client.UUID == clientId {
			sf.remoteClient(clientId)
			return
		}
	}
}

func (sf *SFURoom) AddIceCandidate(clientId string, iceCandidate string) error {
	if peer, ok := sf.producerMap[clientId]; ok {
		return peer.AddIceCandidate(iceCandidate)
	}
	return errors.New("peer connection not found")
}

// GetAllPeers return all peers expect client himself/herself
func (sf *SFURoom) GetAllPeers(clientId string) []*SFUPeer {
	peersIds := make([]*SFUPeer, 0)
	for _, peer := range sf.producerMap {
		if peer.client.UUID != clientId {
			peersIds = append(peersIds, peer)
		}
	}
	return peersIds
}

func (sf *SFURoom) GetPeerById(clientId string) (*SFUPeer, error) {
	for _, peer := range sf.producerMap {
		if peer.client.UUID == clientId {
			return peer, nil
		}
	}
	return nil, errors.New("peer not found")
}

func (sf *SFURoom) Consume(clientId string) {
	//Create a channel uses for receiving data from producers
}

func (sf *SFURoom) addNewClient(clientId string, Client *SFUPeer) {
	sf.Lock()
	defer sf.Unlock()
	sf.producerMap[clientId] = Client
}

func (sf *SFURoom) remoteClient(clientId string) {
	sf.Lock()
	defer sf.Unlock()
	if existPeer, ok := sf.producerMap[clientId]; ok {
		existPeer.peerConn.Close() //closed the existing connection.
		existPeer = nil
	}
}
