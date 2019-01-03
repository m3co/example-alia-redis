package aliaredis

import (
	"net"
	"regexp"
	"sync"
)

var errListenerIsNil = "nil listener"
var errMessageInProcessIsNil = "message is empty"
var errMessageInProcessNotMatched = "invalid message"

type response struct {
	value string
	ok    bool
}

// Server server struct
type Server struct {
	Addr    func() net.Addr
	Close   func() error
	Accept  func() (net.Conn, error)
	Listen  func(network, address string) (net.Listener, error)
	process func(s *Server, message string) (*response, error) // oh hell no!

	store sync.Map
	reSet *regexp.Regexp
	reGet *regexp.Regexp
	reDel *regexp.Regexp
}
