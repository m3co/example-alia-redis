package aliaredis

import "log"

// Process - handle an incomming connection
func (s *Server) Process(message string) error {
	log.Println("message", message)
	return nil
}
