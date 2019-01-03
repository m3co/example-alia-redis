package aliaredis

import (
	"errors"
	"regexp"
)

func (s *Server) init() {
	s.process = process
	s.reSet = regexp.MustCompile("^(?i)set ([a-z0-9-]+)(?-i) (.*)")
	s.reGet = regexp.MustCompile("^(?i)get ([a-z0-9-]+)(?-i)")
	s.reDel = regexp.MustCompile("^(?i)del ([a-z0-9-]+)(?-i)")
}

// Start - start the server at addr
func (s *Server) Start(addr string) error {
	s.init()
	if s.Listen == nil {
		return errors.New(errListenerIsNil)
	}
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
