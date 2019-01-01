package aliaredis

import (
	"io"
	"log"
	"net"
	"testing"
)

type dummyTestHandleConn struct {
	net.Conn
}

func (d dummyTestHandleConn) Read(s []byte) (int, error) {
	n := copy(s, "message")
	return n, io.EOF
}

func (d dummyTestHandleConn) Close() error {
	return nil
}

type dummyTestHandleServer struct {
	Server
}

func (s *dummyTestHandleServer) Process(message string) error {
	log.Println("dummy Process", message)
	return nil
}

func Test_Handle_normally(t *testing.T) {

	// setup
	s := dummyTestHandleServer{}

	/*
		// this is what I in fact want to do
		s.Process = func(message string) error {
			return nil
		}
		// how to proceed here?
	*/

	conn := dummyTestHandleConn{}
	s.Handle(conn)
}
