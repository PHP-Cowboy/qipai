package net

import "sync"

type Session struct {
	sync.RWMutex
	Uid  int
	Cid  string
	data map[string]any
}

func NewSession(cid string) *Session {
	return &Session{
		Cid: cid,
	}
}

func (s *Session) Put(k string, v any) {
	s.Lock()
	defer s.Unlock()
	s.data[k] = v
}

func (s *Session) Get(k string) (any, bool) {
	s.RLock()
	defer s.RUnlock()

	v, ok := s.data[k]
	return v, ok
}

func (s *Session) SetData(uid int, data map[string]any) {
	s.Lock()
	defer s.Unlock()

	if s.Uid == uid {
		for k, v := range data {
			s.data[k] = v
		}
	}
}
