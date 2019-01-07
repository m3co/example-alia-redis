package main

import (
	"net"
	"os"
	"regexp"
	"sync"
)

var modeDisk = "disk"
var modeMem = "memory"

var errDisconnectClient = "bye"
var errMessageInProcessIsNil = "message is empty"
var errMessageInProcessNotMatched = "invalid message"

type response struct {
	value *string
}

// Server server struct
type Server struct {
	Addr    func() net.Addr
	Close   func() error
	Accept  func() (net.Conn, error)
	Listen  func(network, address string) (net.Listener, error)
	process func(s *Server, message string) (*response, error) // oh hell no!

	verbose bool
	port    int
	mode    string

	store     sync.Map
	storeFile *os.File

	reSet *regexp.Regexp
	reGet *regexp.Regexp
	reDel *regexp.Regexp
	reEnd *regexp.Regexp
}
