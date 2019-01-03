package aliaredis

import (
	"errors"
	"fmt"
	"net"
	"testing"
)

func Test_Serve_Start_normally(t *testing.T) {

	// setup
	s := Server{}
	ListenCalled := false

	s.Listen = func(_, _ string) (net.Listener, error) {
		ListenCalled = true
		return nil, nil
	}

	// excercise
	err := s.Start("")

	// verify
	if !ListenCalled {
		t.Error("can't call Listen function")
	}
	if fmt.Sprint(err) != errListenerIsNil {
		t.Error("Expecting errListenerIsNil")
	}
}

func Test_Serve_Start_without_Listen(t *testing.T) {

	// setup
	s := Server{}

	// excercise
	err := s.Start("")

	// verify
	if fmt.Sprint(err) != errListenerIsNil {
		t.Error("Expecting errListenerIsNil")
	}
}

func Test_Serve_Start_with_error(t *testing.T) {

	// setup
	s := Server{}
	ListenCalled := false
	expectedError := "expected error"

	s.Listen = func(_, _ string) (net.Listener, error) {
		ListenCalled = true
		return nil, errors.New(expectedError)
	}

	// excercise
	err := s.Start("")

	// verify
	if !ListenCalled {
		t.Error("can't call Listen function")
	}
	if err == nil {
		t.Error("expecting an error")
	}
	if fmt.Sprint(err) != expectedError {
		t.Error("they differ")
	}
}
