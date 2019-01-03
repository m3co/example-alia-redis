package aliaredis

import (
	"bufio"
	"fmt"
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
			log.Printf("%s, closing...", err)
			return err
		}
		if response == nil {
			log.Printf("nil")
		} else {
			if (*response).value == nil {
				log.Println("nil")
			} else {
				conn.Write([]byte(fmt.Sprintf("%q", *response.value)))
			}
		}
	}
	return nil
}
