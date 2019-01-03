package aliaredis

import (
	"errors"
	"log"
)

// process - handle an incomming connection
func process(s *Server, message string) error {
	if message == "" {
		return errors.New(errMessageInProcessIsNil)
	}
	log.Println(s.reSet.FindAllString(message, -1))
	return nil
}
