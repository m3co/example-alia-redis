package aliaredis

func (s *Server) set(key, value string) {
	s.store.Store(key, value)
}

func (s *Server) get(key string) *response {
	value, ok := s.store.Load(key)
	return &response{value: value.(string), ok: ok}
}
