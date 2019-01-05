package aliaredis

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
	conn, err := net.Dial("tcp", ":3000")
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
