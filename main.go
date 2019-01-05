package main

import (
	"github.com/rianby64/example-alia-redis/aliaredis"
)

func main() {
	s := aliaredis.Server{}
	s.ListenAndServe("")
}
