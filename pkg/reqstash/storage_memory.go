package reqstash

import (
	"container/list"
	"sync"
	"time"
)

type MemoryStorage struct {
	sync.RWMutex
	Capacity int
	TTL      time.Duration

	requests list.List
}

var (
	_ Storage  = (*MemoryStorage)(nil)
	_ Purgable = (*MemoryStorage)(nil)
)

func (s *MemoryStorage) Put(req Request) error {
	s.Lock()
	defer s.Unlock()

	s.requests.PushBack(req)
	if s.Capacity > 0 && s.requests.Len() > s.Capacity {
		s.requests.Remove(s.requests.Front())
	}

	return nil
}

func (s *MemoryStorage) ListAll() ([]Request, error) {
	s.RLock()
	defer s.RUnlock()

	requests := []Request{}
	for e := s.requests.Front(); e != nil; e = e.Next() {
		requests = append(requests, e.Value.(Request))
	}

	return requests, nil
}

func (s *MemoryStorage) Purge() (int, error) {
	if s.TTL == 0 {
		return 0, nil
	}

	s.Lock()
	defer s.Unlock()

	return 0, nil
}
