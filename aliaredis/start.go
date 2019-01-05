package aliaredis

import (
	"fmt"
	"log"
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

// ListenAndServe whatever
func (s *Server) ListenAndServe(addr string) error {
	mode, port := parseArgs()
	s.mode = mode
	s.port = port

	addrInternal := fmt.Sprintf(":%v", s.port)
	if addr != "" {
		addrInternal = addr
	}
	if err := s.Start(addrInternal); err != nil {
		return err
	}
	defer s.Close()

	log.Printf("server is running on %s\n", s.Addr())

	for {
		conn, err := s.Accept()
		if err != nil {
			return err
		}
		go s.Handle(conn)
	}
}
