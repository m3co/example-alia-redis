package main

import (
	"log"
	"net"

	"./aliaredis"
)

func main() {
	s := aliaredis.Server{}

	// setup
	s.Listen = net.Listen

	if err := s.Start(":3000"); err != nil {
		log.Fatalln(err)
	}
	defer s.Close()

	log.Printf("s is running on %s\n", s.Addr())

	for {
		log.Println("accepting connections")
		conn, err := s.Accept()
		log.Println("accepted")
		if err != nil {
			log.Fatalln(err)
		}

		log.Println("going goroutine")

		go s.Handle(conn)
	}
}
