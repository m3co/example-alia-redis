package aliaredis

import (
	"bufio"
	"fmt"
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
			conn.Write([]byte(fmt.Sprintf("%s, closing...", err)))
			if fmt.Sprint(err) == errServerEnd {
				s.Close()
			}
			return err
		}
		if response == nil {
			conn.Write([]byte("nil"))
		} else {
			if (*response).value == nil {
				conn.Write([]byte("nil"))
			} else {
				conn.Write([]byte(fmt.Sprintf("%q", *response.value)))
			}
		}
	}
	return nil
}
