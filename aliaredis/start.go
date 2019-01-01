package aliaredis

import "errors"

// Start - start the server at addr
func (s *Server) Start(addr string) error {
	listener, err := s.Listen("tcp", addr)
	if err != nil {
		return err
	}

	if listener != nil {
		s.Close = listener.Close
		s.Accept = listener.Accept
		s.Addr = listener.Addr
	} else {
		return errors.New("nil listener")
	}

	return nil
}
