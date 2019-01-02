package aliaredis

import "net"

// errListenerIsNil whatever
var errListenerIsNil = "nil listener"

// Server server struct
type Server struct {
	Addr    func() net.Addr
	Close   func() error
	Accept  func() (net.Conn, error)
	Listen  func(network, address string) (net.Listener, error)
	Process func(message string) error
}
