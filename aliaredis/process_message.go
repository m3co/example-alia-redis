package aliaredis

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
		key := match[2]
		value := match[3]
		return s.set(key, value), nil
	}
	if s.reGet.MatchString(message) {
		match := s.reGet.FindStringSubmatch(message)
		key := match[2]
		return s.get(key), nil
	}
	return nil, errors.New(errMessageInProcessNotMatched)
}
