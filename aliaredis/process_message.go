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
	log.Println("message", message)
	return nil
}
