package sessionManager

import (
	"errors"
	"github.com/pion/webrtc/v3"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/server/rtcSFU/session"
	"sync"
)

type SessionManager struct {
	sync.Mutex
	sessionMap           map[string]*session.Session
	newTrackLoadReceived chan struct {
		clientId  string
		sessionId string
		track     *webrtc.TrackLocalStaticRTP
	}
}

func NewSessionManager() *SessionManager {
	return &SessionManager{
		sessionMap: make(map[string]*session.Session),
	}
}

func (sm *SessionManager) CreateOneSession(sid string, callType string) *session.Session {
	s := session.NewSession(sid, callType)
	go func() {
		sm.onListingSessionNewTrack(s)
	}()
	sm.Lock()
	defer sm.Unlock()
	sm.sessionMap[sid] = s
	return s
}

func (sm *SessionManager) onListingSessionNewTrack(s *session.Session) {
	s.OnListingNewTrack()
}

func (sm *SessionManager) RemoveOneSession(sid string) {
	sm.Lock()
	defer sm.Unlock()
	if _, ok := sm.sessionMap[sid]; ok {
		delete(sm.sessionMap, sid)
	}
}

func (sm *SessionManager) GetOneSession(sid string) (*session.Session, error) {
	if s, ok := sm.sessionMap[sid]; ok {
		return s, nil
	}
	return nil, errors.New("session not found")
}
