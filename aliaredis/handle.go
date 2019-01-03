package aliaredis

import (
	"bufio"
	"log"
	"net"
)

// Handle - handle an incomming connection
func (s *Server) Handle(conn net.Conn) error {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		message := scanner.Text()
		response, err := s.process(s, message)
		if err != nil {
			return err
		}
		log.Println("response:", response)
	}
	return nil
}
