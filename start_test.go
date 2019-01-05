package main

import (
	"errors"
	"fmt"
	"net"
	"testing"
	"time"
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

func Test_Serve_Start_send_one_command_and_exit(t *testing.T) {

	// setup
	s := Server{}
	go s.ListenAndServe("")
	time.Sleep(10 * time.Millisecond)

	expectedRes := fmt.Sprintf("%q", "OK")
	conn, err := net.Dial("tcp", ":9090")
	if err != nil {
		t.Error("something went wrong", err)
		return
	}

	conn.Write([]byte("set key value\n"))
	res := make([]byte, 4)
	conn.Read(res)

	if string(res) != expectedRes {
		t.Error("expecting OK from server", string(res), expectedRes)
	}
	conn.Close()
	s.Close()
}

func Test_Serve_restoreFromString_OK(t *testing.T) {

	// setup
	s := Server{}
	s.init()
	dump := "set key1 value1\nset key2 value2\n"
	s.restoreFromString(dump)

	value1, ok := s.store.Load("key1")
	if !ok {
		t.Error("expecing key1 to be in store")
	}
	if value1.(string) != "value1" {
		t.Error("expecting key1 to have as value value1")
	}
	value2, ok := s.store.Load("key2")
	if !ok {
		t.Error("expecing key2 to be in store")
	}
	if value2.(string) != "value2" {
		t.Error("expecting key2 to have as value value2")
	}
}
