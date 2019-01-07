package main

import (
	"errors"
)

// process - handle an incomming connection
func process(s *Server, message string) (*response, error) {
	if message == "" {
		return nil, errors.New(errMessageInProcessIsNil)
	}
	if s.reSet.MatchString(message) {
		match := s.reSet.FindStringSubmatch(message)
		key := match[1]
		value := match[2]
		return s.set(key, value), nil
	}
	if s.reGet.MatchString(message) {
		match := s.reGet.FindStringSubmatch(message)
		key := match[1]
		return s.get(key), nil
	}
	if s.reDel.MatchString(message) {
		match := s.reDel.FindStringSubmatch(message)
		keys := match[1]
		return s.del(keys), nil
	}
	if s.reEnd.MatchString(message) {
		return nil, errors.New(errDisconnectClient)
	}
	return nil, errors.New(errMessageInProcessNotMatched)
}
