package aliaredis

import (
	"errors"
)

// process - handle an incomming connection
func process(s *Server, message string) error {
	if message == "" {
		return errors.New(errMessageInProcessIsNil)
	}
	if s.reSet.MatchString(message) {
		match := s.reSet.FindStringSubmatch(message)
		key := match[2]
		value := match[3]
		s.set(key, value)
		return nil
	}
	return errors.New(errMessageInProcessNotMatched)
}
