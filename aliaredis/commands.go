package aliaredis

import (
	"strconv"
	"strings"
)

func (s *Server) set(key, value string) *response {
	s.store.Store(key, value)
	ok := "OK"
	return &response{value: &ok}
}

func (s *Server) get(key string) *response {
	value, _ := s.store.Load(key)
	if value == nil {
		return &response{value: nil}
	}
	ok := value.(string)
	return &response{value: &ok}
}

func (s *Server) del(keys string) *response {
	i := 0
	for _, key := range strings.Split(keys, " ") {
		if _, ok := s.store.Load(key); ok == true {
			i = i + 1
			s.store.Delete(key)
		}
	}
	ok := strconv.Itoa(i)
	return &response{value: &ok}
}
