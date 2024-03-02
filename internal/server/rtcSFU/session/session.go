package session

import (
	"github.com/ryantokmanmokmtm/chat-app-server/internal/server/rtcSFU/transportClient"
	"sync"
)

type Session struct {
	sync.Mutex
	SessionId      string //roomId
	sessionClients map[string]*transportClient.TransportClient
}

func NewSession(SessionId string) *Session {
	return &Session{
		SessionId:      SessionId,
		sessionClients: make(map[string]*transportClient.TransportClient),
	}
}

func (s *Session) AddNewSessionClient(clientId string, client *transportClient.TransportClient) {
	s.Lock()
	defer s.Unlock()
	s.sessionClients[clientId] = client
}

func (s *Session) GetSessionClients() []string {
	sessionClient := make([]string, 0)
	for id, _ := range s.sessionClients {
		sessionClient = append(sessionClient, id)
	}

	return sessionClient
}

func (s *Session) RemoveSessionClient(clientId string) {
	s.Lock()
	defer s.Unlock()
	_, ok := s.sessionClients[clientId]
	if ok {
		delete(s.sessionClients, clientId)
	}
}
