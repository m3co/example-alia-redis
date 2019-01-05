package main

import (
	"flag"
	"log"
)

func parseArgs() (bool, string, int) {
	// not sure if the following has sense...
	// out there should be libs that allow to define short and long flags

	// I'm not going to test this function

	modePtrLong := flag.String("mode", "mem", "mem - write into memory\ndisk - write into disk\n")
	modePtr := flag.String("m", "", "-mode")
	portPtrLong := flag.Int("port", 9090, "port to listen\n")
	portPtr := flag.Int("p", 0, "-port")
	verbosePtrLong := flag.Bool("verbose", false, "verbose to listen\n")
	verbosePtr := flag.Bool("v", false, "-verbose")

	flag.Parse()

	modeShort := *modePtr
	modeLong := *modePtrLong
	portShort := *portPtr
	portLong := *portPtrLong
	verboseShort := *verbosePtr
	verboseLong := *verbosePtrLong

	var verbose bool
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

	if verboseShort == true {
		verbose = verboseShort
	}
	if verboseLong == true {
		verbose = verboseLong
	}

	return verbose, mode, port
}
