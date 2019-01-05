package main

import (
	"net"
	"testing"
)

type dummyTestServeCloseAcceptAddr struct{}

func (d dummyTestServeCloseAcceptAddr) Close() error {
	return nil
}
func (d dummyTestServeCloseAcceptAddr) Accept() (net.Conn, error) {
	return nil, nil
}
func (d dummyTestServeCloseAcceptAddr) Addr() net.Addr {
	return nil
}

func Test_Serve_Close(t *testing.T) {

	// setup
	s := Server{}
	CloseCalled := false

	s.Listen = func(_, _ string) (net.Listener, error) {
		var d net.Listener = dummyTestServeCloseAcceptAddr{}
		return d, nil
	}
	s.Start("")

	s.Close = func() error {
		CloseCalled = true
		return nil
	}

	// excercise
	s.Close()

	// verify
	if !CloseCalled {
		t.Error("can't call Close")
	}
}

func Test_Serve_Accept(t *testing.T) {

	// setup
	s := Server{}
	AcceptCalled := false

	s.Listen = func(_, _ string) (net.Listener, error) {
		var d net.Listener = dummyTestServeCloseAcceptAddr{}
		return d, nil
	}
	s.Start("")

	s.Accept = func() (net.Conn, error) {
		AcceptCalled = true
		return nil, nil
	}

	// excercise
	s.Accept()

	// verify
	if !AcceptCalled {
		t.Error("can't call Accept")
	}
}

func Test_Serve_Addr(t *testing.T) {

	// setup
	s := Server{}
	AddrCalled := false

	s.Listen = func(_, _ string) (net.Listener, error) {
		var d net.Listener = dummyTestServeCloseAcceptAddr{}
		return d, nil
	}
	s.Start("")

	s.Addr = func() net.Addr {
		AddrCalled = true
		return nil
	}

	// excercise
	s.Addr()

	// verify
	if !AddrCalled {
		t.Error("can't call Addr")
	}
}
