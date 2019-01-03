package aliaredis

import (
	"net"
	"regexp"
)

var errListenerIsNil = "nil listener"
var errMessageInProcessIsNil = "message is empty"
var errMessageInProcessNotMatched = "invalid message"

// Server server struct
type Server struct {
	Addr    func() net.Addr
	Close   func() error
	Accept  func() (net.Conn, error)
	Listen  func(network, address string) (net.Listener, error)
	process func(s *Server, message string) error // oh hell no!

	store map[string]string
	reSet *regexp.Regexp
	reGet *regexp.Regexp
	reDel *regexp.Regexp
}
