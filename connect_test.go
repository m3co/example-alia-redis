package main

import (
	"net"
	"testing"
)

func Test_Serve_redefine_Listen(t *testing.T) {

	// setup
	s := aliaRedis{}
	ListenCalled := false

	s.transport.Listen = func(_, _ string) (net.Listener, error) {
		ListenCalled = true
		return nil, nil
	}

	// excercise
	s.Serve("")

	// verify
	if !ListenCalled {
		t.Error("can't call Listen function")
	}
}
