package aliaredis

// Serve - serves at addr
func (s *Server) Serve(addr string) error {
	s.Addr = addr
	listener, err := s.Listen("tcp", s.Addr)
	if err != nil {
		return err
	}
	s.Listener = listener
	return nil
}
