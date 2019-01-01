package aliaredis

import "net"

// ERRLISTENERISNIL whatever
var ERRLISTENERISNIL = "nil listener"

// Server server struct
type Server struct {
	Addr   func() net.Addr
	Close  func() error
	Accept func() (net.Conn, error)
	Listen func(network, address string) (net.Listener, error)
}
