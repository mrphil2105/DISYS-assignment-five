package main

import (
	"flag"
	lo "log"
)

var (
	port = flag.String("port", "50050", "Port to listen on")
	log  = lo.Default()
)
