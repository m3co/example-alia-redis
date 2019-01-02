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

var expectedMessage = "message"

func (d dummyTestHandleConn) Read(s []byte) (int, error) {
	n := copy(s, expectedMessage)
	return n, io.EOF
}

func (d dummyTestHandleConn) Close() error {
	return nil
}

func Test_Handle_normally(t *testing.T) {

	// setup
	s := Server{}
	actualMessage := ""

	s.Process = func(message string) error {
		actualMessage = message
		return nil
	}

	conn := dummyTestHandleConn{}
	err := s.Handle(conn)

	if err != nil {
		t.Error("unexpected error")
	}
	if actualMessage != expectedMessage {
		t.Error("Handle is not calling Process method")
	}
}

func Test_Handle_Process_returns_error(t *testing.T) {

	// setup
	s := Server{}
	expectedError := errors.New("expected error")

	s.Process = func(message string) error {
		return expectedError
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
