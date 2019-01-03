package aliaredis

import (
	"testing"
)

func Test_Commands_message_set(t *testing.T) {

	// setup
	s := Server{}
	s.init()

	res, err := s.process(&s, "set key value")
	if res == nil {
		t.Error("awaiting for a response")
	}
	if res.value != "OK" {
		t.Error("expecting to see OK as result")
	}
	if !res.ok {
		t.Error("expecting to see ok as true")
	}

	var key interface{} = "key"
	actual, ok := s.store.Load(key)

	if !ok {
		t.Error("expecting key value")
	} else {
		if actual.(string) != "value" {
			t.Error("expecting correct value")
		}
	}
	if err != nil {
		t.Error("unexpected error", err)
	}
}

func Test_Commands_message_get_OK(t *testing.T) {

	// setup
	s := Server{}
	s.init()
	s.store.Store("key", "value")

	res, err := s.process(&s, "get key value")
	if res == nil {
		t.Error("awaiting for a response")
	}
	if res.value != "value" {
		t.Error("expecting to see OK as result")
	}
	if !res.ok {
		t.Error("expecting to see ok as true")
	}

	var key interface{} = "key"
	actual, ok := s.store.Load(key)

	if !ok {
		t.Error("expecting key value")
	} else {
		if actual.(string) != "value" {
			t.Error("expecting correct value")
		}
	}
	if err != nil {
		t.Error("unexpected error", err)
	}
}

func Test_Commands_message_get_Error(t *testing.T) {

	// setup
	s := Server{}
	s.init()

	res, err := s.process(&s, "get key value")
	if res == nil {
		t.Error("awaiting for a response")
	}
	if res.value != "" {
		t.Error("expecting to see '' as result")
	}
	if res.ok {
		t.Error("expecting to see ok as false")
	}
	if err != nil {
		t.Error("unexpected error", err)
	}
}
