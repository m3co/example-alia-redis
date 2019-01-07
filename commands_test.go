package main

import (
	"fmt"
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
	if *res.value != "OK" {
		t.Error("expecting to see OK as result")
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
	if *res.value != "value" {
		t.Error("expecting to see OK as result")
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
	if res.value != nil {
		t.Error("expecting to see '' as result", res)
	}
	if err != nil {
		t.Error("unexpected error", err)
	}
}

func Test_Commands_message_Delete_OK(t *testing.T) {

	// setup
	s := Server{}
	s.init()
	s.store.Store("key1", "value1")
	s.store.Store("key2", "value2")
	s.store.Store("key3", "value3")

	res, err := s.process(&s, "del key1 key2 key3")
	if res != nil {
		t.Error("awaiting for an 'error'")
	}

	// it's a bit strange, but the result of deletion is an error
	if fmt.Sprint(err) != "3" {
		t.Error("expecting to see 3 as result")
	}
}
