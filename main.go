package main

import (
	"bufio"
	"log"
	"net"
)

func main() {
	addr := ":3000"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln(err)
	}
	defer (func() {
		log.Println("closing listener")
		listener.Close()
	})()

	log.Printf("Server is running on %s\n", addr)

	for {
		log.Println("accepting connections")
		conn, err := listener.Accept()
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
