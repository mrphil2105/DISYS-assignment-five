package main

import (
	"flag"
	"fmt"
	"io"
	lo "log"
	"os"
	"strconv"
)

var (
	start = flag.String("start", "server",
		"Entrypoint for the application. Either client, server or frontend")
	name         = flag.String("name", "NoName", "Name of the instance")
	node         = flag.String("node", "0", "A number to assign the node. Used to determine port")
	countingPort = uint16(50051)
	port         = ""
	frontendPort = "50050"
	log          = lo.Default()
)

func main() {
	flag.Parse()
	port = strconv.Itoa(int(countingPort + ParsePort(*node)))

	prefix := fmt.Sprintf("%-8s: ", *name)
	logFileName := fmt.Sprintf("%s.log", *name)
	logFile, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)

	if err != nil {
		log.Fatalf("Unable to open or create file %s: %v", logFileName, err)
	}

	log = lo.New(io.MultiWriter(os.Stdout, logFile), prefix, lo.Ltime)

	switch *start {
	case "server":
		server()
	case "client":
		client()
	case "frontend":
		frontend()
	default:
		log.Fatalf("Invalid start value '%s'. Expected 'client', 'server' or 'frontend'", *start)
	}
}
