package main

import (
	"errors"
	"fmt"
	"io"
	"net"
	"testing"
)

// TestHandle_normally
var expectedMessageTestHandle = "message"
var actualMessageTestHandle = ""

type dummyTestHandleConn struct {
	net.Conn
}

func (d dummyTestHandleConn) Read(s []byte) (int, error) {
	n := copy(s, expectedMessageTestHandle)
	return n, io.EOF
}
func (d dummyTestHandleConn) Write(s []byte) (int, error) {
	actualMessageTestHandle = fmt.Sprintf("%s", s)
	return 0, nil
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
	if actualMessageTestHandle != fmt.Sprintln("nil") {
		t.Error("unexpected error")
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
		if actualMessageTestHandle != fmt.Sprintln(fmt.Sprintf(
			"%s", fmt.Sprint(expectedError))) {
			t.Error("Actual error message differs from expected")
		}
	} else {
		t.Error("unexpected normal execution")
	}
}

// Test_Handle_Command_get
var actualMessageTestHandleCommandGet = ""

type dummyTestHandleCommandGetConn struct {
	net.Conn
}

func (d dummyTestHandleCommandGetConn) Read(s []byte) (int, error) {
	n := copy(s, "get key")
	return n, io.EOF
}

func (d dummyTestHandleCommandGetConn) Write(s []byte) (int, error) {
	actualMessageTestHandleCommandGet = fmt.Sprintf("%s", s)
	return 0, nil
}

func (d dummyTestHandleCommandGetConn) Close() error {
	return nil
}

func Test_Handle_Command_get(t *testing.T) {

	// setup
	s := Server{}
	expectedValue := "value"
	s.store.Store("key", expectedValue)
	s.init()

	conn := dummyTestHandleCommandGetConn{}
	err := s.Handle(conn)

	if err != nil {
		t.Error("unexpected error")
	}
	if actualMessageTestHandleCommandGet != fmt.Sprintln(
		fmt.Sprintf("%q", expectedValue)) {
		t.Error("expected value differs from actual value")
	}
}

// Test_Handle_Command_set
var actualMessageTestHandleCommandSet = ""
var expectedValueToStoreTestHandleCommandSet = "value"

type dummyTestHandleCommandSetConn struct {
	net.Conn
}

func (d dummyTestHandleCommandSetConn) Read(s []byte) (int, error) {
	n := copy(s, fmt.Sprintf("set key %s", expectedValueToStoreTestHandleCommandSet))
	return n, io.EOF
}

func (d dummyTestHandleCommandSetConn) Write(s []byte) (int, error) {
	actualMessageTestHandleCommandSet = fmt.Sprintf("%s", s)
	return 0, nil
}

func (d dummyTestHandleCommandSetConn) Close() error {
	return nil
}

func Test_Handle_Command_set(t *testing.T) {

	// setup
	s := Server{}
	s.init()

	conn := dummyTestHandleCommandSetConn{}
	err := s.Handle(conn)

	actualValue, ok := s.store.Load("key")

	if !ok {
		t.Error("unexpected error while storing the pair key/value")
	}
	if err != nil {
		t.Error("unexpected error")
	}
	if fmt.Sprintf("%q", actualValue) != fmt.Sprintf(
		"%q", expectedValueToStoreTestHandleCommandSet) {
		t.Error("expected value differs from actual value")
	}
}

// Test_Handle_Command_get_nonexisting_key
var actualMessageTestHandleCommandGetNonexistingKey = ""

type dummyTestHandleCommandGetNonexistingKeyConn struct {
	net.Conn
}

func (d dummyTestHandleCommandGetNonexistingKeyConn) Read(s []byte) (int, error) {
	n := copy(s, "get key")
	return n, io.EOF
}

func (d dummyTestHandleCommandGetNonexistingKeyConn) Write(s []byte) (int, error) {
	actualMessageTestHandleCommandGetNonexistingKey = fmt.Sprintf("%s", s)
	return 0, nil
}

func (d dummyTestHandleCommandGetNonexistingKeyConn) Close() error {
	return nil
}

func Test_Handle_Command_get_nonexisting_key(t *testing.T) {

	// setup
	s := Server{}
	expectedValue := fmt.Sprintln("nil")
	s.init()

	conn := dummyTestHandleCommandGetNonexistingKeyConn{}
	err := s.Handle(conn)

	if err != nil {
		t.Error("unexpected error")
	}
	if actualMessageTestHandleCommandGetNonexistingKey != expectedValue {
		t.Error("expected value differs from actual value")
	}
}

// Test_Handle_Command_end
var actualMessageTestHandleCommandEnd = ""
var closeConnectionTestHandleCommandEnd = false

type dummyTestHandleCommandEndConn struct {
	net.Conn
}

func (d dummyTestHandleCommandEndConn) Read(s []byte) (int, error) {
	n := copy(s, "bye")
	return n, io.EOF
}

func (d dummyTestHandleCommandEndConn) Write(s []byte) (int, error) {
	actualMessageTestHandleCommandEnd = fmt.Sprintf("%s", s)
	return 0, nil
}

func (d dummyTestHandleCommandEndConn) Close() error {
	closeConnectionTestHandleCommandEnd = true
	return nil
}

func Test_Handle_Command_end(t *testing.T) {

	// setup
	s := Server{}
	expectedValue := fmt.Sprintln(errDisconnectClient)
	s.init()

	conn := dummyTestHandleCommandEndConn{}
	err := s.Handle(conn)

	if err == nil {
		t.Error("unexpected normal execution")
	}
	if actualMessageTestHandleCommandEnd != expectedValue {
		t.Error("expected value differs from actual value")
	}
	if !closeConnectionTestHandleCommandEnd {
		t.Error("expected Close function to be called")
	}
}
