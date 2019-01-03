package aliaredis

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
