package aliaredis

import (
	"errors"
	"fmt"
	"testing"
)

func Test_Process_empty_message_returns_error(t *testing.T) {

	// setup
	s := Server{}
	s.process = process
	expectedError := errors.New(errMessageInProcessIsNil)

	_, err := s.process(&s, "")

	if err != nil {
		if fmt.Sprint(err) != fmt.Sprint(expectedError) {
			t.Errorf("Error should be %v", errMessageInProcessIsNil)
		}
	} else {
		t.Error("unexpected normal execution")
	}
}

func Test_Process_message_not_matched_error(t *testing.T) {

	// setup
	s := Server{}
	expectedError := errors.New(errMessageInProcessNotMatched)
	s.init()

	_, err := s.process(&s, "whatever")

	if err != nil {
		if fmt.Sprint(err) != fmt.Sprint(expectedError) {
			t.Errorf("Error should be %v", errMessageInProcessNotMatched)
		}
	} else {
		t.Error("unexpected normal execution")
	}
}

func Test_Process_message_set_OK(t *testing.T) {

	// setup
	s := Server{}
	s.init()

	_, err := s.process(&s, "set key value")

	if err != nil {
		t.Error("unexpected error", err)
	}
}
