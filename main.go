package main

import (
	"flag"
	"log"

	"github.com/rianby64/example-alia-redis/aliaredis"
)

func parseArgs() (string, int) {
	// not sure if the following has sense...
	// out there should be libs that allow to define short and long flags

	// I'm not going to test this function

	modePtrLong := flag.String("mode", "mem", "mem - write into memory\ndisk - write into disk\n")
	modePtr := flag.String("m", "", "-mode")
	portPtrLong := flag.Int("port", 9090, "port to listen\n")
	portPtr := flag.Int("p", 0, "-port")

	flag.Parse()

	modeShort := *modePtr
	modeLong := *modePtrLong
	portShort := *portPtr
	portLong := *portPtrLong

	var mode string
	var port int

	if modeShort == "" {
		mode = modeLong
	} else {
		mode = modeShort
		if modeLong != "mem" {
			log.Panicln("--mode and -m defined. Use only one value")
		}
	}

	if portShort == 0 {
		port = portLong
	} else {
		port = portShort
		if portLong != 9090 {
			log.Panicln("--port and -p defined. Use only one value")
		}
	}

	return mode, port
}

func main() {

	parseArgs()

	s := aliaredis.Server{}
	s.ListenAndServe("")
}
