package aliaredis

import (
	"bufio"
	"net"
)

// Handle - handle an incomming connection
func (s *Server) Handle(conn net.Conn) error {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		message := scanner.Text()
		if err := s.process(s, message); err != nil {
			return err
		}
	}
	return nil
}
