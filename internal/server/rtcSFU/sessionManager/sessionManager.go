package sessionManager

import (
	"errors"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/server/rtcSFU/session"
	"sync"
)

type SessionManager struct {
	sync.Mutex
	sessionMap map[string]*session.Session
}

func NewSessionManager() *SessionManager {
	return &SessionManager{
		sessionMap: make(map[string]*session.Session),
	}
}
func (sm *SessionManager) CreateOneSession(sid string) *session.Session {
	s := session.NewSession(sid)
	sm.addSession(sid, s)
	return s
}
func (sm *SessionManager) GetOneSession(sid string) (*session.Session, error) {
	if s, ok := sm.sessionMap[sid]; ok {
		return s, nil
	}
	return nil, errors.New("session not found")
}

func (sm *SessionManager) addSession(sid string, s *session.Session) {
	sm.Lock()
	defer sm.Unlock()
	sm.sessionMap[sid] = s
}

func (sm *SessionManager) removeSession(sid string) {
	sm.Lock()
	defer sm.Unlock()
	if _, ok := sm.sessionMap[sid]; ok {
		delete(sm.sessionMap, sid)
	}
}
