package main

import (
	"flag"
	lo "log"
)

var (
	port         = flag.String("port", "50060", "Port to listen on")
	frontendPort = "50050"
	log          = lo.Default()
)
