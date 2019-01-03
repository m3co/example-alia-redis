package aliaredis

import (
	"errors"
	"fmt"
	"io"
	"net"
	"testing"
)

type dummyTestHandleConn struct {
	net.Conn
}

var expectedMessageTestHandle = "message"

func (d dummyTestHandleConn) Read(s []byte) (int, error) {
	n := copy(s, expectedMessageTestHandle)
	return n, io.EOF
}

func (d dummyTestHandleConn) Close() error {
	return nil
}

func Test_Handle_normally(t *testing.T) {

	// setup
	s := Server{}
	actualMessage := ""

	s.process = func(s *Server, message string) (*response, error) {
		actualMessage = message
		return nil, nil
	}

	conn := dummyTestHandleConn{}
	err := s.Handle(conn)

	if err != nil {
		t.Error("unexpected error")
	}
	if actualMessage != expectedMessageTestHandle {
		t.Error("Handle is not calling Process method")
	}
}

func Test_Handle_Process_returns_error(t *testing.T) {

	// setup
	s := Server{}
	expectedError := errors.New("expected error")

	s.process = func(s *Server, message string) (*response, error) {
		return nil, expectedError
	}

	conn := dummyTestHandleConn{}
	err := s.Handle(conn)

	if err != nil {
		if fmt.Sprint(err) != fmt.Sprint(expectedError) {
			t.Error("Handle is not calling Process method")
		}
	} else {
		t.Error("unexpected normal execution")
	}

}
