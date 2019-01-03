package aliaredis

import (
	"testing"
)

func Test_Commands_message_set(t *testing.T) {

	// setup
	s := Server{}
	s.init()

	_, err := s.process(&s, "set key value")
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

func Test_Commands_message_get(t *testing.T) {

	// setup
	s := Server{}
	s.init()
	s.store.Store("key", "value")

	_, err := s.process(&s, "get key value")
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
