package aliaredis

import (
	"net"
	"testing"
)

func Test_Serve_redefine_Listen(t *testing.T) {

	// setup
	s := Server{}
	ListenCalled := false

	s.Listen = func(_, _ string) (net.Listener, error) {
		ListenCalled = true
		return nil, nil
	}

	// excercise
	s.Start("")

	// verify
	if !ListenCalled {
		t.Error("can't call Listen function")
	}
}
