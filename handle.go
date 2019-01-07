package main

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
			conn.Write([]byte(fmt.Sprintln(err)))
			if fmt.Sprint(err) == errDisconnectClient {
				conn.Close()
				return err
			}
			continue
		}
		if response == nil {
			conn.Write([]byte(fmt.Sprintln("nil")))
		} else {
			if (*response).value == nil {
				conn.Write([]byte(fmt.Sprintln("nil")))
			} else {
				conn.Write([]byte(fmt.Sprintln(fmt.Sprintf("%q", *response.value))))
			}
		}
	}
	return nil
}
