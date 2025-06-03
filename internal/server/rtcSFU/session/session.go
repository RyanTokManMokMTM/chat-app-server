package session

import (
	"errors"
	"sync"

	"github.com/pion/webrtc/v3"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/server/rtcSFU/transportClient"
	"github.com/zeromicro/go-zero/core/logx"
)

type Session struct {
	sync.Mutex
	SessionID            string //roomID
	CallType             string
	sessionClients       map[string]*transportClient.TransportClient
	newTrackLoadReceived chan struct {
		clientID string
		track    *webrtc.TrackLocalStaticRTP
	}
}

func NewSession(SessionID, callType string) *Session {
	return &Session{
		SessionID:      SessionID,
		sessionClients: make(map[string]*transportClient.TransportClient),
		newTrackLoadReceived: make(chan struct {
			clientID string
			track    *webrtc.TrackLocalStaticRTP
		}),
		CallType: callType,
	}
}

func (s *Session) OnListingNewTrack() {
	for {
		select {
		case info, ok := <-s.newTrackLoadReceived:
			if !ok {
				logx.Error("newTrackLoadReceived closed")
				return
			}

			if info.track == nil {
				logx.Error("Track is nil")
				return
			}
			logx.Infof("On received a new track.from client %s............", info.clientID)
			for _, id := range s.GetSessionClients() {
				if id == info.clientID {
					//Track from the client id
					continue
				}

				//Get Client Info
				tc, err := s.GetTransportClient(id)
				if err != nil {
					logx.Error(err)
					continue
				}

				//Get Client current consumer info
				c, err := tc.GetConsumerByID(info.clientID)
				if err != nil {
					logx.Error(err)
					continue
				}

				//Check consumer existing current track.

				//for _, trans := range c.GetPeerConnection().GetTransceivers() {
				//	logx.Info("Current c : trans")
				//	s := trans.Sender()
				//	if s == nil {
				//		logx.Info("Sender is nil")
				//	}
				//	t := trans.Receiver()
				//	if t == nil {
				//		logx.Info("Reveiver is nil")
				//	}
				//
				//}

				if err := c.AddLocalTrack(info.track); err != nil {
					logx.Error(err)
					continue
				}

			}
		}
	}
}

func (s *Session) AddNewSessionClient(clientID string, client *transportClient.TransportClient) {
	s.Lock()
	defer s.Unlock()
	s.sessionClients[clientID] = client
	logx.Infof("Added %s to session", clientID)
}

func (s *Session) GetSessionClients() []string {
	sessionClient := make([]string, 0)
	for id, _ := range s.sessionClients {
		sessionClient = append(sessionClient, id)
	}

	return sessionClient
}

func (s *Session) IsEmpty() bool {
	if len(s.sessionClients) == 0 {
		return true
	}
	return false
}

func (s *Session) RemoveSessionClient(clientID string) {
	s.Lock()
	defer s.Unlock()
	_, ok := s.sessionClients[clientID]
	if ok {
		delete(s.sessionClients, clientID)
	}
}

func (s *Session) GetTransportClient(clientID string) (*transportClient.TransportClient, error) {
	client, ok := s.sessionClients[clientID]
	if !ok {
		return nil, errors.New("client not in the session")
	}

	return client, nil
}

func (s *Session) OnNewTrack(clientID string, track *webrtc.TrackLocalStaticRTP) {
	s.newTrackLoadReceived <- struct {
		clientID string
		track    *webrtc.TrackLocalStaticRTP
	}{clientID: clientID, track: track}
}
