package aliaredis

import "net"

var errListenerIsNil = "nil listener"
var errMessageInProcessIsNil = "message is empty"

// Server server struct
type Server struct {
	Addr    func() net.Addr
	Close   func() error
	Accept  func() (net.Conn, error)
	Listen  func(network, address string) (net.Listener, error)
	process func(s *Server, message string) error // oh hell no!
}
