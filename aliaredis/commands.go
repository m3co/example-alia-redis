package aliaredis

func (s *Server) set(key, value string) *response {
	s.store.Store(key, value)
	return &response{value: "OK", ok: true}
}

func (s *Server) get(key string) *response {
	value, ok := s.store.Load(key)
	if value == nil {
		return &response{value: "", ok: ok}
	}
	return &response{value: value.(string), ok: ok}
}
