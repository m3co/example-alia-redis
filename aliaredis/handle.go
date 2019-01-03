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
		if response, err := s.process(s, message); err != nil {
			if response != "" {
				log.Println("response:", response)
			}
			return err
		}
	}
	return nil
}
