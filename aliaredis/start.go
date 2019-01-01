package aliaredis

// Start - start the server at addr
func (s *Server) Start(addr string) error {
	s.Addr = addr
	listener, err := s.Listen("tcp", s.Addr)
	if err != nil {
		return err
	}
	s.Listener = listener
	return nil
}
