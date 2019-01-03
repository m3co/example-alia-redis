package aliaredis

func (s *Server) set(key, value string) {
	s.store.Store(key, value)
}
