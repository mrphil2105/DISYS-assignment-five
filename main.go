package main

import (
	"flag"
	"fmt"
	"io"
	lo "log"
	"os"
)

var (
	start = flag.String("start", "server",
		"Entrypoint for the application. Either client, server or frontend")
	name         = flag.String("name", "NoName", "Name of the instance")
	port         = flag.String("port", "50060", "Port to listen on")
	frontendPort = "50050"
	log          = lo.Default()
)

func main() {
	flag.Parse()

	prefix := fmt.Sprintf("%-8s: ", *name)
	logFileName := fmt.Sprintf("%s.log", *name)
	logFile, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)

	if err != nil {
		log.Fatalf("Unable to open or create file %s: %v", logFileName, err)
	}

	switch *start {
	case "server":
		log = lo.New(io.MultiWriter(os.Stdout, logFile), prefix, lo.Ltime)
		server()
	case "client":
		log = lo.New(logFile, prefix, lo.Ltime)
		// TODO: Add 'client()'
	case "frontend":
		log = lo.New(io.MultiWriter(os.Stdout, logFile), prefix, lo.Ltime)
		// TODO: Add 'frontend()'
	default:
		log.Fatalf("Invalid start value '%s'. Expected 'client', 'server' or 'frontend'.", *start)
	}
}
