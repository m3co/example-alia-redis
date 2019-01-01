package aliaredis

import "net"

// Server server struct
type Server struct {
	Addr     string
	Listener net.Listener
	Listen   func(network, address string) (net.Listener, error)
}
