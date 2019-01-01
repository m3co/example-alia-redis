package main

import (
	"bufio"
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
	defer s.Listener.Close()

	log.Printf("s is running on %s\n", s.Addr)

	for {
		log.Println("accepting connections")
		conn, err := s.Listener.Accept()
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
