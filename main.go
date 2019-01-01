package main

import (
	"bufio"
	"log"
	"net"
)

type transportAdapter struct {
	Listen func(network, address string) (net.Listener, error)
}

type aliaRedis struct {
	addr      string
	listener  net.Listener
	transport transportAdapter
}

// Serve - serves at addr
func (s *aliaRedis) Serve(addr string) error {
	s.addr = addr
	listener, err := s.transport.Listen("tcp", s.addr)
	if err != nil {
		return err
	}
	s.listener = listener
	return nil
}

func main() {
	s := aliaRedis{}

	// setup
	s.transport.Listen = net.Listen

	if err := s.Serve(":3000"); err != nil {
		log.Fatalln(err)
	}
	defer s.listener.Close()

	log.Printf("s is running on %s\n", s.addr)

	for {
		log.Println("accepting connections")
		conn, err := s.listener.Accept()
		log.Println("accepted")
		if err != nil {
			log.Fatalln(err)
		}

		log.Println("going goroutine")
		go (func() {
			defer (func() {
				log.Println("closing connection")
				conn.Close()
			})()
			log.Println("scanning")
			scanner := bufio.NewScanner(conn)
			log.Println("scanned")
			for scanner.Scan() {
				log.Println("reading")
				text := scanner.Text()
				log.Println("read", text)
			}
			log.Println("finished goroutine")
		})()
	}
}
