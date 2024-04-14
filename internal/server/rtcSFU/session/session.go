package session

import (
	"errors"
	"github.com/pion/webrtc/v3"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/server/rtcSFU/transportClient"
	"github.com/zeromicro/go-zero/core/logx"
	"sync"
)

type Session struct {
	sync.Mutex
	SessionId            string //roomId
	sessionClients       map[string]*transportClient.TransportClient
	newTrackLoadReceived chan struct {
		clientId string
		track    *webrtc.TrackLocalStaticRTP
	}
}

func NewSession(SessionId string) *Session {
	return &Session{
		SessionId:      SessionId,
		sessionClients: make(map[string]*transportClient.TransportClient),
		newTrackLoadReceived: make(chan struct {
			clientId string
			track    *webrtc.TrackLocalStaticRTP
		}),
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
			logx.Infof("On received a new track.from client %s............", info.clientId)
			for _, id := range s.GetSessionClients() {
				if id == info.clientId {
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
				c, err := tc.GetConsumerById(info.clientId)
				if err != nil {
					logx.Error(err)
					continue
				}

				//Check consumer existing current track.
				for _, receiver := range c.GetPeerConnection().GetSenders() {
					t := receiver.Track()

					if t == nil {
						logx.Error("(receiver)Track is nil")
						continue
					}
					logx.Infof("(receiver)Current track with Kind %s, Id : %s, streamId :%s", t.Kind(), t.ID(), t.StreamID())
					if info.track.StreamID() == t.StreamID() {
						return
					}
				}

				if err := c.AddLocalTrack(info.track); err != nil {
					logx.Error(err)
					continue
				}

			}
		}
	}
}

func (s *Session) AddNewSessionClient(clientId string, client *transportClient.TransportClient) {
	s.Lock()
	defer s.Unlock()
	s.sessionClients[clientId] = client
	logx.Infof("Added %s to session", clientId)
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

func (s *Session) RemoveSessionClient(clientId string) {
	s.Lock()
	defer s.Unlock()
	_, ok := s.sessionClients[clientId]
	if ok {
		delete(s.sessionClients, clientId)
	}
}

func (s *Session) GetTransportClient(clientId string) (*transportClient.TransportClient, error) {
	client, ok := s.sessionClients[clientId]
	if !ok {
		return nil, errors.New("client not in the session")
	}

	return client, nil
}

func (s *Session) OnNewTrack(clientId string, track *webrtc.TrackLocalStaticRTP) {
	s.newTrackLoadReceived <- struct {
		clientId string
		track    *webrtc.TrackLocalStaticRTP
	}{clientId: clientId, track: track}
}
