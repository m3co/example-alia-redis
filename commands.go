package main

import (
	"log"
	"strconv"
	"strings"
)

func (s *Server) set(key, value string) *response {
	s.store.Store(key, value)
	if s.verbose {
		log.Println("request command set", key, value)
	}
	ok := "OK"
	if s.verbose {
		log.Println("result command set", ok)
	}
	return &response{value: &ok}
}

func (s *Server) get(key string) *response {
	value, _ := s.store.Load(key)
	if s.verbose {
		log.Println("request command get", key, value)
	}
	if value == nil {
		if s.verbose {
			log.Println("result command get nil")
		}
		return &response{value: nil}
	}
	ok := value.(string)
	if s.verbose {
		log.Println("result command get", ok)
	}
	return &response{value: &ok}
}

func (s *Server) del(keys string) *response {
	i := 0
	if s.verbose {
		log.Println("request command del", keys)
	}
	for _, key := range strings.Split(keys, " ") {
		if _, ok := s.store.Load(key); ok == true {
			i = i + 1
			s.store.Delete(key)
		}
	}
	ok := strconv.Itoa(i)
	if s.verbose {
		log.Println("result command del", ok)
	}
	return &response{value: &ok}
}
