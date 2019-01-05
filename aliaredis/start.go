package aliaredis

import (
	"fmt"
	"log"
	"net"
	"os"
	"regexp"
	"strings"
)

func (s *Server) commit() {
	s.store.Range(func(key interface{}, value interface{}) bool {
		row := fmt.Sprintln(fmt.Sprintf("set %s %s", key.(string), value.(string)))
		s.storeFile.Write([]byte(row))
		return true
	})
}

func (s *Server) init() {
	s.process = process
	if s.Listen == nil {
		s.Listen = net.Listen
	}

	s.reSet = regexp.MustCompile("^(?i)set ([a-z0-9-]+)(?-i) (.*)")
	s.reGet = regexp.MustCompile("^(?i)get ([a-z0-9-]+)(?-i)")
	s.reDel = regexp.MustCompile("^(?i)del ([0-9a-z ]+)(?-i)")
	s.reEnd = regexp.MustCompile("^(?i)end$(?-i)")

	if s.mode == modeDisk {
		path := fmt.Sprintf("./db_port_%v", s.port)
		var storeFile *os.File
		fi, err := os.Stat(path)
		if err == nil {
			storeFile, err = os.OpenFile(path, os.O_RDWR|os.O_EXCL, 0600)
			if err != nil {
				log.Panicln(err)
			}
			dumpBytes := make([]byte, fi.Size())
			_, err := storeFile.Read(dumpBytes)
			if err != nil {
				log.Panic(err)
			}
			dump := string(dumpBytes)
			rows := strings.Split(dump, fmt.Sprintln(""))
			for _, value := range rows {
				if value != "" {
					_, err := s.process(s, value)
					if err != nil {
						log.Panic(err)
					}
				}
			}
		} else if os.IsNotExist(err) {
			storeFile, err = os.Create(path)
			if err != nil {
				log.Panicln(err)
			}
		} else {
			log.Panicln(err)
		}
		s.storeFile = storeFile
		s.storeFile.Seek(0, 0)
	}
}

// Start - start the server at addr
func (s *Server) Start(addr string) error {
	s.init()
	listener, err := s.Listen("tcp", addr)
	if err != nil {
		log.Println(err)
		return err
	}

	s.Close = listener.Close
	s.Accept = listener.Accept
	s.Addr = listener.Addr

	return nil
}

// ListenAndServe whatever
func (s *Server) ListenAndServe(addr string) error {
	mode, port := parseArgs()
	s.mode = mode
	s.port = port

	addrInternal := fmt.Sprintf(":%v", s.port)
	if addr != "" {
		addrInternal = addr
	}
	if err := s.Start(addrInternal); err != nil {
		return err
	}
	defer (func() {
		s.Close()
		if s.mode == modeDisk {
			s.commit()
			s.storeFile.Close()
		}
	})()

	log.Printf("server is running on %s\n", s.Addr())

	for {
		conn, err := s.Accept()
		if err != nil {
			return err
		}
		go s.Handle(conn)
	}
}
