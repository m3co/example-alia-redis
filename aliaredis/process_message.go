package aliaredis

import "log"

// process - handle an incomming connection
func process(message string) error {
	log.Println("message", message)
	return nil
}
