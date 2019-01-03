package aliaredis

import (
	"errors"
	"fmt"
	"net"
	"testing"
)

func Test_Serve_Start_with_error(t *testing.T) {

	// setup
	s := Server{}
	s.init()
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
