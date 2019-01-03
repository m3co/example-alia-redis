package aliaredis

import (
	"net"
	"regexp"
)

func (s *Server) init() {
	s.process = process
	if s.Listen == nil {
		s.Listen = net.Listen
	}

	s.reSet = regexp.MustCompile("^(?i)set ([a-z0-9-]+)(?-i) (.*)")
	s.reGet = regexp.MustCompile("^(?i)get ([a-z0-9-]+)(?-i)")
	s.reDel = regexp.MustCompile("^(?i)del ([0-9a-z ]+)(?-i)")
}

// Start - start the server at addr
func (s *Server) Start(addr string) error {
	s.init()
	listener, err := s.Listen("tcp", addr)
	if err != nil {
		return err
	}

	s.Close = listener.Close
	s.Accept = listener.Accept
	s.Addr = listener.Addr

	return nil
}
